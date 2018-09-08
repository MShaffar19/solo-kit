package grpc

import (
	"context"
	"crypto/sha1"
	"fmt"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	envoytranscoder "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/transcoder/v2"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"github.com/gogo/googleapis/google/api"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

	"github.com/envoyproxy/go-control-plane/pkg/util"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
	glooplugins "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins"
	grpcapi "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins/grpc"
	transformapi "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins/transformation"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/plugins"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/plugins/pluginutils"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/plugins/transformation"
)

type ServicesAndDescriptor struct {
	Spec        *grpcapi.ServiceSpec
	Descriptors *descriptor.FileDescriptorSet
}

func NewPlugin(transformsAdded *bool) plugins.Plugin {
	return &plugin{
		recordedUpstreams: make(map[core.ResourceRef]*v1.Upstream),
		upstreamServices:  make(map[string]ServicesAndDescriptor),
		transformsAdded:   transformsAdded,
	}
}

type plugin struct {
	transformsAdded   *bool
	recordedUpstreams map[core.ResourceRef]*v1.Upstream
	upstreamServices  map[string]ServicesAndDescriptor

	ctx context.Context
}

const (
	filterName  = "envoy.grpc_json_transcoder"
	pluginStage = plugins.PreOutAuth

	ServiceTypeGRPC = "gRPC"
)

func (p *plugin) Init(params plugins.InitParams) error {
	p.ctx = params.Ctx
	return nil
}

func (p *plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, out *envoyapi.Cluster) error {
	servicespec, ok := in.UpstreamSpec.UpstreamType.(v1.ServiceSpecGetter)
	if !ok {
		return nil
	}

	if servicespec.GetServiceSpec() == nil {
		return nil
	}
	if servicespec.GetServiceSpec().PluginType == nil {
		return nil
	}

	grpcwrapper, ok := servicespec.GetServiceSpec().PluginType.(*glooplugins.ServiceSpec_Grpc)
	if !ok {
		return nil
	}
	grpcspec := grpcwrapper.Grpc

	if len(grpcspec.GrpcServices) == 0 {
		return errors.New("service_info.properties.service_names cannot be empty")
	}
	descriptors, err := convertProto(grpcspec.Descriptors)
	if err != nil {
		return errors.Wrapf(err, "parsing grpc spec as a proto descriptor set")
	}

	for _, svc := range grpcspec.GrpcServices {

		// find the relevant service

		err := addHttpRulesToProto(in, svc, descriptors)
		if err != nil {
			return errors.Wrapf(err, "failed to generate http rules for service %s in proto descriptors", svc.Name)
		}
	}

	addWellKnownProtos(descriptors)

	p.recordedUpstreams[in.Metadata.Ref()] = in
	p.upstreamServices[in.Metadata.Name] = ServicesAndDescriptor{
		Descriptors: descriptors,
		Spec:        grpcspec,
	}

	out.Http2ProtocolOptions = &envoycore.Http2ProtocolOptions{}

	return nil
}

func genFullServiceName(packageName, serviceName string) string {
	return packageName + "." + serviceName
}

func convertProto(b []byte) (*descriptor.FileDescriptorSet, error) {
	var fileDescriptor descriptor.FileDescriptorSet
	err := proto.Unmarshal(b, &fileDescriptor)
	return &fileDescriptor, err
}

func getPath(matcher *v1.Matcher) string {
	switch path := matcher.PathSpecifier.(type) {
	case *v1.Matcher_Prefix:
		return path.Prefix
	case *v1.Matcher_Exact:
		return path.Exact
	case *v1.Matcher_Regex:
		return path.Regex
	}
	panic("invalid matcher")
}

func (p *plugin) ProcessRoute(params plugins.Params, in *v1.Route, out *envoyroute.Route) error {

	return pluginutils.MarkPerFilterConfig(p.ctx, in, out, transformation.FilterName, func(spec *v1.Destination) (proto.Message, error) {
		// check if it's grpc destination
		if spec.DestinationSpec == nil {
			return nil, nil
		}
		grpcDestinationSpecWrapper, ok := spec.DestinationSpec.DestinationType.(*v1.DestinationSpec_Grpc)
		if !ok {
			return nil, nil
		}
		// copy as it might be modified
		grpcDestinationSpec := *grpcDestinationSpecWrapper.Grpc

		if grpcDestinationSpec.Parameters == nil {
			path := getPath(in.Matcher) + "?{query_string}"

			grpcDestinationSpec.Parameters = &transformapi.Parameters{
				Path: &types.StringValue{Value: path},
			}
		}

		// get the package_name.service_name to generate the path that envoy wants
		fullServiceName := genFullServiceName(grpcDestinationSpec.Package, grpcDestinationSpec.Service)
		methodName := grpcDestinationSpec.Function

		upstream := p.recordedUpstreams[spec.Upstream]
		if upstream == nil {
			return nil, errors.New("upstream was not recorded for grpc route")
		}

		// create the transformation for the route
		outPath := httpPath(upstream, fullServiceName, methodName)

		// add query matcher to out path. kombina for now
		// TODO: support query for matching
		outPath += `?{{ default(query_string), "")}}`

		// we always choose post
		httpMethod := "POST"
		return &transformapi.TransformationTemplate{
			Headers: map[string]*transformapi.InjaTemplate{
				":method": {Text: httpMethod},
				":path":   {Text: outPath},
			},
			BodyTransformation: &transformapi.TransformationTemplate_MergeExtractorsToBody{
				MergeExtractorsToBody: &transformapi.MergeExtractorsToBody{},
			},
		}, nil

	})

}

// returns package name
func addHttpRulesToProto(upstream *v1.Upstream, currentsvc *grpcapi.ServiceSpec_GrpcService, set *descriptor.FileDescriptorSet) error {
	for _, file := range set.File {
		if file.Package == nil || *file.Package != currentsvc.PackageName {
			continue
		}
		for _, svc := range file.Service {
			if svc.Name == nil || *svc.Name != currentsvc.Name {
				continue
			}
			for _, method := range svc.Method {
				fullServiceName := genFullServiceName(currentsvc.PackageName, currentsvc.Name)
				if method.Options == nil {
					method.Options = &descriptor.MethodOptions{}
				}
				if err := proto.SetExtension(method.Options, api.E_Http, &api.HttpRule{
					Pattern: &api.HttpRule_Post{
						Post: httpPath(upstream, fullServiceName, *method.Name),
					},
					Body: "*",
				}); err != nil {
					return errors.Wrap(err, "setting http extensions for method.Options")
				}
				log.Debugf("method.options: %v", *method.Options)
			}
		}
	}

	return nil
}

func addWellKnownProtos(descriptors *descriptor.FileDescriptorSet) {
	var googleApiHttpFound, googleApiAnnotationsFound, googleApiDescriptorFound bool
	for _, file := range descriptors.File {
		log.Debugf("inspecting descriptor for proto file %v...", *file.Name)
		if *file.Name == "google/api/http.proto" {
			googleApiHttpFound = true
			continue
		}
		if *file.Name == "google/api/annotations.proto" {
			googleApiAnnotationsFound = true
			continue
		}
		if *file.Name == "google/protobuf/descriptor.proto" {
			googleApiDescriptorFound = true
			continue
		}
	}
	if !googleApiDescriptorFound {
		addGoogleApisDescriptor(descriptors)
	}

	if !googleApiHttpFound {
		addGoogleApisHttp(descriptors)
	}

	if !googleApiAnnotationsFound {
		//TODO: investigate if we need this
		//addGoogleApisAnnotations(packageName, set)
	}
}

func httpPath(upstream *v1.Upstream, serviceName, methodName string) string {
	h := sha1.New()
	h.Write([]byte(upstream.Metadata.Namespace + upstream.Metadata.Name + serviceName))
	return "/" + fmt.Sprintf("%x", h.Sum(nil))[:8] + "/" + upstream.Metadata.Name + "/" + serviceName + "/" + methodName
}

func (p *plugin) HttpFilters(params plugins.Params, listener *v1.HttpListener) ([]plugins.StagedHttpFilter, error) {

	if len(p.upstreamServices) == 0 {
		return nil, nil
	}

	var filters []plugins.StagedHttpFilter
	for _, serviceAndDescriptor := range p.upstreamServices {
		descriptorBytes, err := proto.Marshal(serviceAndDescriptor.Descriptors)
		if err != nil {
			return nil, errors.Wrapf(err, "marshaling proto descriptor")
		}
		var fullServiceNames []string
		for _, grpcsvc := range serviceAndDescriptor.Spec.GrpcServices {
			fullName := genFullServiceName(grpcsvc.PackageName, grpcsvc.Name)
			fullServiceNames = append(fullServiceNames, fullName)
		}
		filterConfig, err := util.MessageToStruct(&envoytranscoder.GrpcJsonTranscoder{
			DescriptorSet: &envoytranscoder.GrpcJsonTranscoder_ProtoDescriptorBin{
				ProtoDescriptorBin: descriptorBytes,
			},
			Services:                  fullServiceNames,
			MatchIncomingRequestRoute: true,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "ERROR: marshaling GrpcJsonTranscoder config")
		}
		filters = append(filters, plugins.StagedHttpFilter{
			HttpFilter: &envoyhttp.HttpFilter{
				Name:   filterName,
				Config: filterConfig,
			},
			Stage: pluginStage,
		})
	}

	if len(filters) == 0 {
		return nil, errors.Errorf("ERROR: no valid GrpcJsonTranscoder available")
	}

	return filters, nil
}