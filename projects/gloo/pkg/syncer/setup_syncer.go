package syncer

import (
	"context"
	"net"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/types"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/bootstrap"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/defaults"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/discovery"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/plugins/registry"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/translator"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/xds"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
)

type RunFunc func(opts bootstrap.Opts) error

// for use by UDS, FDS, other v1.SetupSyncers
func NewSetupSyncerWithRunFunc(runFunc RunFunc) v1.SetupSyncer {
	return &settingsSyncer{
		grpcServer: func(ctx context.Context) *grpc.Server {
			return grpc.NewServer(grpc.StreamInterceptor(
				grpc_middleware.ChainStreamServer(
					grpc_ctxtags.StreamServerInterceptor(),
					grpc_zap.StreamServerInterceptor(zap.NewNop()),
					func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
						contextutils.LoggerFrom(ctx).Infof("gRPC call: %v", info.FullMethod)
						return handler(srv, ss)
					},
				)),
			)
		},
		runFunc: runFunc,
	}
}

func NewSetupSyncer() v1.SetupSyncer {
	return NewSetupSyncerWithRunFunc(RunGloo)
}

type settingsSyncer struct {
	runFunc    RunFunc
	grpcServer func(ctx context.Context) *grpc.Server
}

func (s *settingsSyncer) Sync(ctx context.Context, snap *v1.SetupSnapshot) error {
	switch {
	case len(snap.Settings.List()) == 0:
		return errors.Errorf("no settings files found")
	case len(snap.Settings.List()) > 1:
		return errors.Errorf("multiple settings files found")
	}
	settings := snap.Settings.List()[0]

	var (
		cfg       *rest.Config
		clientset kubernetes.Interface
	)
	cache := memory.NewInMemoryResourceCache()
	kubeCache := kube.NewKubeCache()

	upstreamFactory, err := bootstrap.ConfigFactoryForSettings(
		settings,
		cache,
		kubeCache,
		v1.UpstreamCrd,
		&cfg,
	)
	if err != nil {
		return err
	}

	proxyFactory, err := bootstrap.ConfigFactoryForSettings(
		settings,
		cache,
		kubeCache,
		v1.ProxyCrd,
		&cfg,
	)
	if err != nil {
		return err
	}

	secretFactory, err := bootstrap.SecretFactoryForSettings(
		settings,
		cache,
		v1.SecretCrd.Plural,
		&cfg,
		&clientset,
	)
	if err != nil {
		return err
	}

	artifactFactory, err := bootstrap.ArtifactFactoryForSettings(
		settings,
		cache,
		v1.ArtifactCrd.Plural,
		&cfg,
		&clientset,
	)
	if err != nil {
		return err
	}

	ipPort := strings.Split(settings.BindAddr, ":")
	if len(ipPort) != 2 {
		return errors.Errorf("invalid bind addr: %v", settings.BindAddr)
	}
	port, err := strconv.Atoi(ipPort[1])
	if err != nil {
		return errors.Wrapf(err, "invalid bind addr: %v", settings.BindAddr)
	}
	refreshRate, err := types.DurationFromProto(settings.RefreshRate)
	if err != nil {
		return err
	}

	writeNamespace := settings.DiscoveryNamespace
	if writeNamespace == "" {
		writeNamespace = defaults.GlooSystem
	}
	watchNamespaces := settings.WatchNamespaces
	var writeNamespaceProvided bool
	for _, ns := range watchNamespaces {
		if ns == writeNamespace {
			writeNamespaceProvided = true
			break
		}
	}
	if !writeNamespaceProvided {
		watchNamespaces = append(watchNamespaces, writeNamespace)
	}
	opts := bootstrap.Opts{
		WriteNamespace:  writeNamespace,
		WatchNamespaces: watchNamespaces,
		Upstreams:       upstreamFactory,
		Proxies:         proxyFactory,
		Secrets:         secretFactory,
		Artifacts:       artifactFactory,
		WatchOpts: clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: refreshRate,
		},
		BindAddr: &net.TCPAddr{
			IP:   net.ParseIP(ipPort[0]),
			Port: port,
		},
		GrpcServer: s.grpcServer(ctx),
		// if nil, kube plugin disabled
		KubeClient: clientset,
		DevMode:    true,
	}

	return s.runFunc(opts)
}

func RunGloo(opts bootstrap.Opts) error {
	watchOpts := opts.WatchOpts.WithDefaults()
	opts.WatchOpts.Ctx = contextutils.WithLogger(opts.WatchOpts.Ctx, "gloo")

	watchOpts.Ctx = contextutils.WithLogger(watchOpts.Ctx, "setup")
	endpointsFactory := &factory.MemoryResourceClientFactory{
		Cache: memory.NewInMemoryResourceCache(),
	}

	upstreamClient, err := v1.NewUpstreamClient(opts.Upstreams)
	if err != nil {
		return err
	}
	if err := upstreamClient.Register(); err != nil {
		return err
	}

	proxyClient, err := v1.NewProxyClient(opts.Proxies)
	if err != nil {
		return err
	}
	if err := proxyClient.Register(); err != nil {
		return err
	}

	endpointClient, err := v1.NewEndpointClient(endpointsFactory)
	if err != nil {
		return err
	}

	secretClient, err := v1.NewSecretClient(opts.Secrets)
	if err != nil {
		return err
	}

	artifactClient, err := v1.NewArtifactClient(opts.Artifacts)
	if err != nil {
		return err
	}

	cache := v1.NewApiEmitter(artifactClient, endpointClient, proxyClient, secretClient, upstreamClient)

	xdsHasher, xdsCache := xds.SetupEnvoyXds(opts.WatchOpts.Ctx, opts.GrpcServer, nil)

	rpt := reporter.NewReporter("gloo", upstreamClient.BaseClient(), proxyClient.BaseClient())

	plugins := registry.Plugins(opts)

	var discoveryPlugins []discovery.DiscoveryPlugin
	for _, plug := range plugins {
		disc, ok := plug.(discovery.DiscoveryPlugin)
		if ok {
			discoveryPlugins = append(discoveryPlugins, disc)
		}
	}

	sync := NewTranslatorSyncer(translator.NewTranslator(plugins), xdsCache, xdsHasher, rpt, opts.DevMode)
	eventLoop := v1.NewApiEventLoop(cache, sync)

	errs := make(chan error)

	eds := discovery.NewEndpointDiscovery(opts.WriteNamespace, endpointClient, discoveryPlugins)
	edsErrs, err := discovery.RunEds(upstreamClient, eds, opts.WriteNamespace, watchOpts)
	if err != nil {
		return err
	}
	go errutils.AggregateErrs(watchOpts.Ctx, errs, edsErrs, "eds.gloo")

	eventLoopErrs, err := eventLoop.Run(opts.WatchNamespaces, watchOpts)
	if err != nil {
		return err
	}
	go errutils.AggregateErrs(watchOpts.Ctx, errs, eventLoopErrs, "event_loop.gloo")

	logger := contextutils.LoggerFrom(watchOpts.Ctx)

	go func() {
		for {
			select {
			case <-watchOpts.Ctx.Done():
				logger.Debugf("context cancelled")
				return
			}
		}
	}()

	lis, err := net.Listen(opts.BindAddr.Network(), opts.BindAddr.String())
	if err != nil {
		return err
	}
	go func() {
		<-opts.WatchOpts.Ctx.Done()
		opts.GrpcServer.Stop()
	}()

	defer opts.GrpcServer.Stop()
	return opts.GrpcServer.Serve(lis)
}
