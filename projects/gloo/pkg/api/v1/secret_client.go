package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type SecretClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*Secret, error)
	Write(resource *Secret, opts clients.WriteOpts) (*Secret, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (SecretList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan SecretList, <-chan error, error)
}

type secretClient struct {
	rc clients.ResourceClient
}

func NewSecretClient(rcFactory factory.ResourceClientFactory) (SecretClient, error) {
	return NewSecretClientWithToken(rcFactory, "")
}

func NewSecretClientWithToken(rcFactory factory.ResourceClientFactory, token string) (SecretClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &Secret{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base Secret resource client")
	}
	return &secretClient{
		rc: rc,
	}, nil
}

func (client *secretClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *secretClient) Register() error {
	return client.rc.Register()
}

func (client *secretClient) Read(namespace, name string, opts clients.ReadOpts) (*Secret, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Secret), nil
}

func (client *secretClient) Write(secret *Secret, opts clients.WriteOpts) (*Secret, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(secret, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Secret), nil
}

func (client *secretClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()
	return client.rc.Delete(namespace, name, opts)
}

func (client *secretClient) List(namespace string, opts clients.ListOpts) (SecretList, error) {
	opts = opts.WithDefaults()
	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToSecret(resourceList), nil
}

func (client *secretClient) Watch(namespace string, opts clients.WatchOpts) (<-chan SecretList, <-chan error, error) {
	opts = opts.WithDefaults()
	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	secretsChan := make(chan SecretList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				secretsChan <- convertToSecret(resourceList)
			case <-opts.Ctx.Done():
				close(secretsChan)
				return
			}
		}
	}()
	return secretsChan, errs, nil
}

func convertToSecret(resources resources.ResourceList) SecretList {
	var secretList SecretList
	for _, resource := range resources {
		secretList = append(secretList, resource.(*Secret))
	}
	return secretList
}
