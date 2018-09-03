package rest

/*
if this destination spec has rest service spec
this will grab the parameters from the route extention
*/
import (
	"context"
	"regexp"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"

	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	glooplugins "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins"
	restapi "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins/rest"
	transformapi "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins/transformation"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/plugins"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/plugins/pluginutils"

	"github.com/solo-io/solo-kit/projects/gloo/pkg/plugins/transformation"
)

type UpstreamWithServiceSpec interface {
	GetServiceSpec() *glooplugins.ServiceSpec
}

type plugin struct {
	transformsAdded   *bool
	recordedUpstreams map[string]UpstreamWithServiceSpec
	ctx               context.Context
}

func NewPlugin(transformsAdded *bool) plugins.Plugin {
	return &plugin{transformsAdded: transformsAdded}
}

func (p *plugin) Init(params plugins.InitParams) error {
	p.ctx = params.Ctx
	p.recordedUpstreams = make(map[string]UpstreamWithServiceSpec)
	return nil
}

func (p *plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, _ *envoyapi.Cluster) error {
	if withServiceSpec, ok := in.UpstreamSpec.UpstreamType.(UpstreamWithServiceSpec); ok {
		p.recordedUpstreams[in.Metadata.Name] = withServiceSpec
	}
	return nil
}

func (p *plugin) ProcessRoute(params plugins.Params, in *v1.Route, out *envoyroute.Route) error {
	return pluginutils.MarkPerFilterConfig(p.ctx, in, out, transformation.FilterName, func(spec *v1.Destination) (proto.Message, error) {
		// check if it's aws destination
		if spec.DestinationSpec == nil {
			return nil, nil
		}
		restDestinationSpec, ok := spec.DestinationSpec.DestinationType.(*v1.DestinationSpec_Rest)
		if !ok {
			return nil, nil
		}
		// get upstream
		getservicespec, ok := p.recordedUpstreams[spec.UpstreamName]
		if !ok {
			// TODO(yuval-k): panic in debug
			return nil, errors.Errorf("%v does not have a service spec", spec.UpstreamName)
		}

		servicesoec := getservicespec.GetServiceSpec()
		if servicesoec == nil {
			return nil, errors.Errorf("%v does has an empty service spec", spec.UpstreamName)

		}
		restservicespec, ok := servicesoec.PluginType.(*glooplugins.ServiceSpec_Rest)
		if restservicespec == nil || !ok {
			return nil, errors.Errorf("%v does not have a REST service spec", spec.UpstreamName)
		}
		funcname := restDestinationSpec.Rest.FunctionName
		transform := restservicespec.Rest.Transformation[funcname]
		if transform == nil {
			return nil, errors.Errorf("unknown function %v", funcname)
		}

		// add extentions from the destination spec
		var err error
		transform.Extractors, err = p.createRequestExtractors(restDestinationSpec.Rest.Parameters)
		if err != nil {
			return nil, err
		}
		// should be aws upstream

		// get function
		ret := &transformapi.RouteTransformations{
			RequestTransformation: &transformapi.Transformation{
				TransformationType: &transformapi.Transformation_TransformationTemplate{
					TransformationTemplate: transform,
				},
			},
		}

		*p.transformsAdded = true

		return ret, nil
	})
}

func (p *plugin) createRequestExtractors(params *restapi.DestinationSpec_Parameters) (map[string]*transformapi.Extraction, error) {
	extractors := make(map[string]*transformapi.Extraction)
	if params == nil {
		return extractors, nil
	}

	// special http2 headers, get the whole thing for free
	// as a convenience to the user
	// TODO: add more
	for _, header := range []string{
		"path",
		"method",
	} {
		p.addHeaderExtractorFromParam(":"+header, "{"+header+"}", extractors)
	}
	// headers we support submatching on
	// custom as well as the path and authority/host header
	if params.Path != nil {
		if err := p.addHeaderExtractorFromParam(":path", params.Path.Value, extractors); err != nil {
			return nil, errors.Wrapf(err, "error processing parameter")
		}
	}
	for headerName, headerValue := range params.Headers {
		if err := p.addHeaderExtractorFromParam(headerName, headerValue, extractors); err != nil {
			return nil, errors.Wrapf(err, "error processing parameter")
		}
	}
	return extractors, nil
}

func (p *plugin) addHeaderExtractorFromParam(header, parameter string, extractors map[string]*transformapi.Extraction) error {
	if parameter == "" {
		return nil
	}
	// remember that the order of the param names correlates with their order in the regex
	paramNames, regexMatcher := getNamesAndRegexFromParamString(parameter)
	contextutils.LoggerFrom(p.ctx).Debugf("transformation pluginN: extraction for header %v: parameters: %v regex matcher: %v", header, paramNames, regexMatcher)
	// if no regex, this is a "default variable" that the user gets for free
	if len(paramNames) == 0 {
		// extract everything
		// TODO(yuval): create a special extractor that doesn't use regex when we just want the whole thing
		extract := &transformapi.Extraction{
			Header:   header,
			Regex:    "(.*)",
			Subgroup: uint32(1),
		}
		extractors[strings.TrimPrefix(header, ":")] = extract
	}

	// count the number of open braces,
	// if they are not equal to the # of counted params,
	// the user gave us bad variable names or unterminated braces and we should error
	expectedParameterCount := strings.Count(parameter, "{")
	if len(paramNames) != expectedParameterCount {
		return errors.Errorf("%v is not valid syntax. {} braces must be closed and variable names must satisfy regex "+
			`([\-._[:alnum:]]+)`, parameter)
	}

	// otherwise it's regex, and we need to create an extraction for each variable name they defined
	for i, name := range paramNames {
		extract := &transformapi.Extraction{
			Header:   header,
			Regex:    regexMatcher,
			Subgroup: uint32(i + 1),
		}
		extractors[name] = extract
	}
	return nil
}

var rxp = regexp.MustCompile(`\{([\.\-_[:word:]]+)\}`)

func getNamesAndRegexFromParamString(paramString string) ([]string, string) {
	// escape regex
	// TODO: make sure all envoy regex is being escaped here
	parameterNames := rxp.FindAllString(paramString, -1)
	for i, name := range parameterNames {
		parameterNames[i] = strings.TrimSuffix(strings.TrimPrefix(name, "{"), "}")
	}

	return parameterNames, buildRegexString(rxp, paramString)
}

func buildRegexString(rxp *regexp.Regexp, paramString string) string {
	var regexString string
	var prevEnd int
	for _, startStop := range rxp.FindAllStringIndex(paramString, -1) {
		start := startStop[0]
		end := startStop[1]
		subStr := regexp.QuoteMeta(paramString[prevEnd:start]) + `([\-._[:alnum:]]+)`
		regexString += subStr
		prevEnd = end
	}

	return regexString + regexp.QuoteMeta(paramString[prevEnd:])
}
