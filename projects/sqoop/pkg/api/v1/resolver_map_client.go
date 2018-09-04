package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type ResolverMapClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*ResolverMap, error)
	Write(resource *ResolverMap, opts clients.WriteOpts) (*ResolverMap, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (ResolverMapList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan ResolverMapList, <-chan error, error)
}

type resolverMapClient struct {
	rc clients.ResourceClient
}

func NewResolverMapClient(rcFactory factory.ResourceClientFactory) (ResolverMapClient, error) {
	return NewResolverMapClientWithToken(rcFactory, "")
}

func NewResolverMapClientWithToken(rcFactory factory.ResourceClientFactory, token string) (ResolverMapClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &ResolverMap{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base ResolverMap resource client")
	}
	return &resolverMapClient{
		rc: rc,
	}, nil
}

func (client *resolverMapClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *resolverMapClient) Register() error {
	return client.rc.Register()
}

func (client *resolverMapClient) Read(namespace, name string, opts clients.ReadOpts) (*ResolverMap, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*ResolverMap), nil
}

func (client *resolverMapClient) Write(resolverMap *ResolverMap, opts clients.WriteOpts) (*ResolverMap, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(resolverMap, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*ResolverMap), nil
}

func (client *resolverMapClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()
	return client.rc.Delete(namespace, name, opts)
}

func (client *resolverMapClient) List(namespace string, opts clients.ListOpts) (ResolverMapList, error) {
	opts = opts.WithDefaults()
	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToResolverMap(resourceList), nil
}

func (client *resolverMapClient) Watch(namespace string, opts clients.WatchOpts) (<-chan ResolverMapList, <-chan error, error) {
	opts = opts.WithDefaults()
	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	resolverMapsChan := make(chan ResolverMapList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				resolverMapsChan <- convertToResolverMap(resourceList)
			case <-opts.Ctx.Done():
				close(resolverMapsChan)
				return
			}
		}
	}()
	return resolverMapsChan, errs, nil
}

func convertToResolverMap(resources resources.ResourceList) ResolverMapList {
	var resolverMapList ResolverMapList
	for _, resource := range resources {
		resolverMapList = append(resolverMapList, resource.(*ResolverMap))
	}
	return resolverMapList
}
