package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type SettingsClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*Settings, error)
	Write(resource *Settings, opts clients.WriteOpts) (*Settings, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (SettingsList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan SettingsList, <-chan error, error)
}

type settingsClient struct {
	rc clients.ResourceClient
}

func NewSettingsClient(rcFactory factory.ResourceClientFactory) (SettingsClient, error) {
	return NewSettingsClientWithToken(rcFactory, "")
}

func NewSettingsClientWithToken(rcFactory factory.ResourceClientFactory, token string) (SettingsClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &Settings{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base Settings resource client")
	}
	return &settingsClient{
		rc: rc,
	}, nil
}

func (client *settingsClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *settingsClient) Register() error {
	return client.rc.Register()
}

func (client *settingsClient) Read(namespace, name string, opts clients.ReadOpts) (*Settings, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Settings), nil
}

func (client *settingsClient) Write(settings *Settings, opts clients.WriteOpts) (*Settings, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(settings, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Settings), nil
}

func (client *settingsClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()
	return client.rc.Delete(namespace, name, opts)
}

func (client *settingsClient) List(namespace string, opts clients.ListOpts) (SettingsList, error) {
	opts = opts.WithDefaults()
	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToSettings(resourceList), nil
}

func (client *settingsClient) Watch(namespace string, opts clients.WatchOpts) (<-chan SettingsList, <-chan error, error) {
	opts = opts.WithDefaults()
	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	settingsChan := make(chan SettingsList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				settingsChan <- convertToSettings(resourceList)
			case <-opts.Ctx.Done():
				close(settingsChan)
				return
			}
		}
	}()
	return settingsChan, errs, nil
}

func convertToSettings(resources resources.ResourceList) SettingsList {
	var settingsList SettingsList
	for _, resource := range resources {
		settingsList = append(settingsList, resource.(*Settings))
	}
	return settingsList
}
