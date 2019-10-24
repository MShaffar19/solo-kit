// Code generated by solo-kit. DO NOT EDIT.

package kubernetes

import (
	"sort"

	github_com_solo_io_solo_kit_api_external_kubernetes_namespace "github.com/solo-io/solo-kit/api/external/kubernetes/namespace"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewKubeNamespace(namespace, name string) *KubeNamespace {
	kubenamespace := &KubeNamespace{}
	kubenamespace.KubeNamespace.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return kubenamespace
}

// require custom resource to implement Clone() as well as resources.Resource interface

type CloneableKubeNamespace interface {
	resources.Resource
	Clone() *github_com_solo_io_solo_kit_api_external_kubernetes_namespace.KubeNamespace
}

var _ CloneableKubeNamespace = &github_com_solo_io_solo_kit_api_external_kubernetes_namespace.KubeNamespace{}

type KubeNamespace struct {
	github_com_solo_io_solo_kit_api_external_kubernetes_namespace.KubeNamespace
}

func (r *KubeNamespace) Clone() resources.Resource {
	return &KubeNamespace{KubeNamespace: *r.KubeNamespace.Clone()}
}

func (r *KubeNamespace) Hash() uint64 {
	clone := r.KubeNamespace.Clone()

	resources.UpdateMetadata(clone, func(meta *core.Metadata) {
		meta.ResourceVersion = ""
	})

	return hashutils.HashAll(clone)
}

func (r *KubeNamespace) GroupVersionKind() schema.GroupVersionKind {
	return KubeNamespaceGVK
}

type KubeNamespaceList []*KubeNamespace

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list KubeNamespaceList) Find(namespace, name string) (*KubeNamespace, error) {
	for _, kubeNamespace := range list {
		if kubeNamespace.GetMetadata().Name == name {
			if namespace == "" || kubeNamespace.GetMetadata().Namespace == namespace {
				return kubeNamespace, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find kubeNamespace %v.%v", namespace, name)
}

func (list KubeNamespaceList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, kubeNamespace := range list {
		ress = append(ress, kubeNamespace)
	}
	return ress
}

func (list KubeNamespaceList) Names() []string {
	var names []string
	for _, kubeNamespace := range list {
		names = append(names, kubeNamespace.GetMetadata().Name)
	}
	return names
}

func (list KubeNamespaceList) NamespacesDotNames() []string {
	var names []string
	for _, kubeNamespace := range list {
		names = append(names, kubeNamespace.GetMetadata().Namespace+"."+kubeNamespace.GetMetadata().Name)
	}
	return names
}

func (list KubeNamespaceList) Sort() KubeNamespaceList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list KubeNamespaceList) Clone() KubeNamespaceList {
	var kubeNamespaceList KubeNamespaceList
	for _, kubeNamespace := range list {
		kubeNamespaceList = append(kubeNamespaceList, resources.Clone(kubeNamespace).(*KubeNamespace))
	}
	return kubeNamespaceList
}

func (list KubeNamespaceList) Each(f func(element *KubeNamespace)) {
	for _, kubeNamespace := range list {
		f(kubeNamespace)
	}
}

func (list KubeNamespaceList) EachResource(f func(element resources.Resource)) {
	for _, kubeNamespace := range list {
		f(kubeNamespace)
	}
}

func (list KubeNamespaceList) AsInterfaces() []interface{} {
	asInterfaces := make([]interface{}, 0, len(list))
	list.Each(func(element *KubeNamespace) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

var (
	KubeNamespaceGVK = schema.GroupVersionKind{
		Version: "kubernetes",
		Group:   "kubernetes.solo.io",
		Kind:    "KubeNamespace",
	}
)
