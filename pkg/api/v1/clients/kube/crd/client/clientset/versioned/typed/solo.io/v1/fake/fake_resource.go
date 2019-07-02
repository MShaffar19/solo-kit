/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	soloiov1 "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd/solo.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"
)

// FakeResources implements ResourceInterface
type FakeResources struct {
	Fake *FakeResourcesV1
	ns   string
}

//
// NOTE(marco): the occurrences of these two variables have been replaced respectively with
//	c.Fake.GroupVersion().WithResource(c.Fake.Plural) and
//	c.Fake.GroupVersion().WithKind(c.Fake.KindName)
// to allow for the fake clientset to work with our CRD registration logic.
//
//var resourcesResource = schema.GroupVersionResource{Group: "resources.solo.io", Version: "v1", Resource: "resources"}
//var resourcesKind = schema.GroupVersionKind{Group: "resources.solo.io", Version: "v1", Kind: "Resource"}

// Get takes name of the resource, and returns the corresponding resource object, and an error if there is any.
func (c *FakeResources) Get(name string, options v1.GetOptions) (result *soloiov1.Resource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(c.Fake.GroupVersion().WithResource(c.Fake.Plural), c.ns, name), &soloiov1.Resource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*soloiov1.Resource), err
}

// List takes label and field selectors, and returns the list of Resources that match those selectors.
func (c *FakeResources) List(opts v1.ListOptions) (result *soloiov1.ResourceList, err error) {
	obj, err := c.Fake.Invokes(
		testing.NewListAction(
			c.Fake.GroupVersion().WithResource(c.Fake.Plural),
			c.Fake.GroupVersion().WithKind(c.Fake.KindName),
			c.ns, opts,
		), &soloiov1.ResourceList{},
	)

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &soloiov1.ResourceList{ListMeta: obj.(*soloiov1.ResourceList).ListMeta}
	for _, item := range obj.(*soloiov1.ResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resources.
func (c *FakeResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(c.Fake.GroupVersion().WithResource(c.Fake.Plural), c.ns, opts))

}

// Create takes the representation of a resource and creates it.  Returns the server's representation of the resource, and an error, if there is any.
func (c *FakeResources) Create(resource *soloiov1.Resource) (result *soloiov1.Resource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(c.Fake.GroupVersion().WithResource(c.Fake.Plural), c.ns, resource), &soloiov1.Resource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*soloiov1.Resource), err
}

// Update takes the representation of a resource and updates it. Returns the server's representation of the resource, and an error, if there is any.
func (c *FakeResources) Update(resource *soloiov1.Resource) (result *soloiov1.Resource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(c.Fake.GroupVersion().WithResource(c.Fake.Plural), c.ns, resource), &soloiov1.Resource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*soloiov1.Resource), err
}

// Delete takes name of the resource and deletes it. Returns an error if one occurs.
func (c *FakeResources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(c.Fake.GroupVersion().WithResource(c.Fake.Plural), c.ns, name), &soloiov1.Resource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(c.Fake.GroupVersion().WithResource(c.Fake.Plural), c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &soloiov1.ResourceList{})
	return err
}

// Patch applies the patch and returns the patched resource.
func (c *FakeResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *soloiov1.Resource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(c.Fake.GroupVersion().WithResource(c.Fake.Plural), c.ns, name, pt, data, subresources...), &soloiov1.Resource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*soloiov1.Resource), err
}
