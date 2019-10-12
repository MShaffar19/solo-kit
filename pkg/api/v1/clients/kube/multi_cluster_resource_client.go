package kube

import (
	"sync"
	"time"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/wrapper"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/multicluster"
	"github.com/solo-io/solo-kit/pkg/multicluster/handler"
	"k8s.io/client-go/rest"
)

var (
	NoClientForClusterError = func(version, name, cluster string) error {
		return errors.Errorf("%v.%v client does not exist for %v", version, name, cluster)
	}
)

type MultiClusterResourceClient interface {
	clients.ResourceClient
	handler.ClusterHandler
}

type multiClusterResourceClient struct {
	crd                crd.Crd
	skipCrdCreation    bool
	namespaceWhitelist []string
	resyncPeriod       time.Duration
	params             factory.NewResourceClientParams

	resourceType    resources.InputResource
	clients         map[string]*ResourceClient
	clientAccess    sync.RWMutex
	cacheGetter     multicluster.KubeSharedCacheGetter
	watchAggregator wrapper.WatchAggregator
}

var _ MultiClusterResourceClient = &multiClusterResourceClient{}

func NewMultiClusterResourceClient(
	cacheGetter multicluster.KubeSharedCacheGetter,
	watchAggregator wrapper.WatchAggregator,
	crd crd.Crd,
	skipCrdCreation bool,
	namespaceWhitelist []string,
	resyncPeriod time.Duration,
	resourceType resources.InputResource,
	params factory.NewResourceClientParams,
) *multiClusterResourceClient {
	return &multiClusterResourceClient{
		cacheGetter:        cacheGetter,
		watchAggregator:    watchAggregator,
		crd:                crd,
		skipCrdCreation:    skipCrdCreation,
		namespaceWhitelist: namespaceWhitelist,
		resyncPeriod:       resyncPeriod,
		resourceType:       resourceType,
		params:             params,
	}
}

func (rc *multiClusterResourceClient) Kind() string {
	return resources.Kind(rc.resourceType)
}

func (rc *multiClusterResourceClient) NewResource() resources.Resource {
	return resources.Clone(rc.resourceType)
}

func (rc *multiClusterResourceClient) Register() error {
	// not implemented
	// per-cluster clients are registered on ClusterAdded
	return nil
}

func (rc *multiClusterResourceClient) Read(namespace, name string, opts clients.ReadOpts) (resources.Resource, error) {
	client, err := rc.clientFor(opts.Cluster)
	if err != nil {
		return nil, err
	}
	return client.Read(namespace, name, opts)
}

func (rc *multiClusterResourceClient) Write(resource resources.Resource, opts clients.WriteOpts) (resources.Resource, error) {
	client, err := rc.clientFor(resource.GetMetadata().Cluster)
	if err != nil {
		return nil, err
	}
	return client.Write(resource, opts)
}

func (rc *multiClusterResourceClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	client, err := rc.clientFor(opts.Cluster)
	if err != nil {
		return err
	}
	return client.Delete(namespace, name, opts)
}

func (rc *multiClusterResourceClient) List(namespace string, opts clients.ListOpts) (resources.ResourceList, error) {
	client, err := rc.clientFor(opts.Cluster)
	if err != nil {
		return nil, err
	}

	return client.List(namespace, opts)
}

func (rc *multiClusterResourceClient) Watch(namespace string, opts clients.WatchOpts) (<-chan resources.ResourceList, <-chan error, error) {
	client, err := rc.clientFor(opts.Cluster)
	if err != nil {
		return nil, nil, err
	}

	return client.Watch(namespace, opts)
}

func (rc *multiClusterResourceClient) ClusterAdded(cluster string, restConfig *rest.Config) {
	kubeResourceClientFactory := &factory.KubeResourceClientFactory{
		Crd:                rc.crd,
		Cfg:                restConfig,
		SharedCache:        rc.cacheGetter.GetCache(cluster),
		SkipCrdCreation:    rc.skipCrdCreation,
		NamespaceWhitelist: rc.namespaceWhitelist,
		ResyncPeriod:       rc.resyncPeriod,
		Cluster:            cluster,
	}

	client, err := kubeResourceClientFactory.NewResourceClient(rc.params)
	if err != nil {
		return
	}
	kubeClient, ok := client.(*ResourceClient)
	if !ok {
		return
	}
	if err := client.Register(); err != nil {
		return
	}
	rc.clientAccess.Lock()
	defer rc.clientAccess.Unlock()
	rc.clients[cluster] = kubeClient
	if rc.watchAggregator != nil {
		rc.watchAggregator.AddWatch(client)
	}
}

func (rc *multiClusterResourceClient) ClusterRemoved(cluster string, restConfig *rest.Config) {
	rc.clientAccess.Lock()
	defer rc.clientAccess.Unlock()
	if client, ok := rc.clients[cluster]; ok {
		delete(rc.clients, cluster)
		if rc.watchAggregator != nil {
			rc.watchAggregator.RemoveWatch(client)
		}
	}
}

func (rc *multiClusterResourceClient) clientFor(cluster string) (*ResourceClient, error) {
	rc.clientAccess.RLock()
	defer rc.clientAccess.RUnlock()
	if client, ok := rc.clients[cluster]; ok {
		return client, nil
	}
	return nil, NoClientForClusterError(rc.crd.Version.Version, rc.crd.KindName, cluster)
}
