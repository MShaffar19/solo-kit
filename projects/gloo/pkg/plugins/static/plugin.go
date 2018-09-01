package static

import (
	"net"

	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoyauth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoyendpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/plugins"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/gogo/protobuf/types"
)

type plugin struct{	hostRewriteUpstreams map[core.ResourceRef]bool}


func NewPlugin() plugins.Plugin {
	return &plugin{}
}

func (p *plugin) Init(params plugins.InitParams) error {
	p.hostRewriteUpstreams = make(map[core.ResourceRef]bool)
	return nil
}

func (p *plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, out *envoyapi.Cluster) error {
	// not ours
	staticSpec, ok := in.UpstreamSpec.UpstreamType.(*v1.UpstreamSpec_Static)
	if !ok {
		return nil
	}

	spec := staticSpec.Static
	var foundSslPort bool
	var hostname string
	for _, host := range spec.Hosts {
		if host.Addr == "" {
			return errors.Errorf("addr cannot be empty for host")
		}
		if host.Port == 0 {
			return errors.Errorf("port cannot be empty for host")
		}
		if host.Port == 443 {
			foundSslPort = true
		}
		ip := net.ParseIP(host.Addr)
		if ip != nil {
			out.Type = envoyapi.Cluster_STATIC
		} else {
			out.Type = envoyapi.Cluster_STRICT_DNS
			// for sni
			hostname = host.Addr
		}

		if out.LoadAssignment == nil {
			out.LoadAssignment = &envoyapi.ClusterLoadAssignment{
				ClusterName: out.Name,
				Endpoints:   []envoyendpoint.LocalityLbEndpoints{{}},
			}
		}

		out.LoadAssignment.Endpoints[0].LbEndpoints = append(out.LoadAssignment.Endpoints[0].LbEndpoints,
			envoyendpoint.LbEndpoint{
				Endpoint: &envoyendpoint.Endpoint{
					Address: &envoycore.Address{
						Address: &envoycore.Address_SocketAddress{
							SocketAddress: &envoycore.SocketAddress{
								Protocol: envoycore.TCP,
								Address:  host.Addr,
								PortSpecifier: &envoycore.SocketAddress_PortValue{
									PortValue: host.Port,
								},
							},
						},
					},
				},
			})
		// fix issue where ipv6 addr cannot bind
		out.DnsLookupFamily = envoyapi.Cluster_V4_ONLY

		// if host port is 443 or if the user wants it, we will use TLS
		if spec.UseTls || foundSslPort {
			// tell envoy to use TLS to connect to this upstream
			// TODO: support client certificates
			out.TlsContext = &envoyauth.UpstreamTlsContext{
				Sni: hostname,
			}
		}

		// the upstream has a DNS name. to cover the case that it is an external service
		// that requires the host header, we will add host rewrite.
		if hostname != "" {
			// cache the name of this upstream, we need to enable automatic host rewrite on routes
			p.hostRewriteUpstreams[in.Metadata.Ref()] = true
		}

	}
	return nil

	// configure the cluster to use EDS:ADS and call it a day
	out.Type = envoyapi.Cluster_STATIC
	out.EdsClusterConfig = &envoyapi.Cluster_EdsClusterConfig{
		EdsConfig: &envoycore.ConfigSource{
			ConfigSourceSpecifier: &envoycore.ConfigSource_Ads{
				Ads: &envoycore.AggregatedConfigSource{},
			},
		},
	}
	return nil
}


func (p *plugin) ProcessRouteAction(params plugins.Params, in *v1.RouteAction, _ map[string]*plugins.RoutePlugin, out *envoyroute.RouteAction) error {
	upstreams := destinationUpstreams(in)
	for _, ref := range upstreams {
		if _, ok := p.hostRewriteUpstreams[ref]; !ok {
			continue
		}
		// this is a route to one of our ssl upstreams
		// enable auto host rewrite
		out.HostRewriteSpecifier = &envoyroute.RouteAction_AutoHostRewrite{
			AutoHostRewrite: &types.BoolValue{
				Value: true,
			},
		}
		// one is good enough
		break
	}
	return nil
}

func destinationUpstreams(in *v1.RouteAction) []core.ResourceRef {
	switch dest := in.Destination.(type) {
	case *v1.RouteAction_Single:
		return []core.ResourceRef{dest.Single.Upstream}
	case *v1.RouteAction_Multi:
		var upstreams []core.ResourceRef
		for _, dest := range dest.Multi.Destinations {
			upstreams = append(upstreams, dest.Destination.Upstream)
		}
		return upstreams
	}
	panic("invalid route")
}