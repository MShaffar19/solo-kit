// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewFakeResource(namespace, name string) *FakeResource {
	fakeresource := &FakeResource{}
	fakeresource.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return fakeresource
}

func (r *FakeResource) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *FakeResource) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.Count,
	)
}

type FakeResourceList []*FakeResource

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list FakeResourceList) Find(namespace, name string) (*FakeResource, error) {
	for _, fakeResource := range list {
		if fakeResource.GetMetadata().Name == name {
			if namespace == "" || fakeResource.GetMetadata().Namespace == namespace {
				return fakeResource, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find fakeResource %v.%v", namespace, name)
}

func (list FakeResourceList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, fakeResource := range list {
		ress = append(ress, fakeResource)
	}
	return ress
}

func (list FakeResourceList) Names() []string {
	var names []string
	for _, fakeResource := range list {
		names = append(names, fakeResource.GetMetadata().Name)
	}
	return names
}

func (list FakeResourceList) NamespacesDotNames() []string {
	var names []string
	for _, fakeResource := range list {
		names = append(names, fakeResource.GetMetadata().Namespace+"."+fakeResource.GetMetadata().Name)
	}
	return names
}

func (list FakeResourceList) Sort() FakeResourceList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list FakeResourceList) Clone() FakeResourceList {
	var fakeResourceList FakeResourceList
	for _, fakeResource := range list {
		fakeResourceList = append(fakeResourceList, resources.Clone(fakeResource).(*FakeResource))
	}
	return fakeResourceList
}

func (list FakeResourceList) Each(f func(element *FakeResource)) {
	for _, fakeResource := range list {
		f(fakeResource)
	}
}

func (list FakeResourceList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *FakeResource) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

var _ resources.Resource = &FakeResource{}

// Kubernetes Adapter for FakeResource

func (o *FakeResource) GetObjectKind() schema.ObjectKind {
	t := FakeResourceCrd.TypeMeta()
	return &t
}

func (o *FakeResource) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*FakeResource)
}

var FakeResourceCrd = crd.NewCrd("crds.testing.solo.io",
	"fakes",
	"crds.testing.solo.io",
	"v1",
	"FakeResource",
	"fk",
	false,
	&FakeResource{})
