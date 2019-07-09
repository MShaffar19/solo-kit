package crd

import (
	"log"
	"sync"

	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd/client/clientset/versioned/scheme"
	v1 "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd/solo.io/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	apiexts "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO(ilackarms): evaluate this fix for concurrent map access in k8s.io/apimachinery/pkg/runtime.SchemaBuider
var (
	registerLock sync.Mutex

	VersionNotFoundError = func(version string) error {
		return errors.Errorf("could not find version %v", version)
	}
)

type Ladder interface {
	Climb(src SoloKitCrd, dst SoloKitCrd) error
	Descend(src SoloKitCrd, dst SoloKitCrd) error
}

type Converter interface {
	Convert(src SoloKitCrd, dst SoloKitCrd) error
}

type SoloKitCrd interface {
	runtime.Object
	resources.InputResource
}

type CrdMeta struct {
	Plural        string
	Group         string
	KindName      string
	ShortName     string
	ClusterScoped bool
}

type Version struct {
	Version string
	Type    SoloKitCrd
}

type Crd struct {
	CrdMeta
	Version Version
}

type MultiVersionCrd struct {
	CrdMeta
	Versions []Version
}

func (m *MultiVersionCrd) GetVersion(requested string) (*Version, error) {
	for _, version := range m.Versions {
		if version.Version == requested {
			return &version, nil
		}
	}
	return nil, VersionNotFoundError(requested)
}

func NewCrd(
	plural string,
	group string,
	version string,
	kindName string,
	shortName string,
	clusterScoped bool,
	objType SoloKitCrd) Crd {
	c := Crd{
		CrdMeta: CrdMeta{
			Plural:        plural,
			Group:         group,
			KindName:      kindName,
			ShortName:     shortName,
			ClusterScoped: clusterScoped,
		},
		Version: Version{
			Version: version,
			Type:    objType,
		},
	}
	if err := c.AddToScheme(scheme.Scheme); err != nil {
		log.Panicf("error while adding [%v] CRD to scheme: %v", c.FullName(), err)
	}
	// if res, ok := objType.(resources.Resource); ok {
	// 	c.Version.ProtoSpec = res
	// } else {
	// 	log.Panicf("error while creating crd for %v, must extend " +
	// 		"resources.Resource interface", c.FullName())
	// }
	return c
}

func (d Crd) Register(apiexts apiexts.Interface) error {
	return GetRegistry().registerCrd(d.GroupVersionKind(), apiexts)
}

func (d Crd) KubeResource(resource resources.InputResource) *v1.Resource {
	res := KubeResource(resource)
	res.TypeMeta = d.TypeMeta()
	return res
}

func (d CrdMeta) FullName() string {
	return d.Plural + "." + d.Group
}

func (d Crd) TypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       d.KindName,
		APIVersion: d.Group + "/" + d.Version.Version,
	}
}

// GroupVersion is group version used to register these objects
func (d Crd) GroupVersion() schema.GroupVersion {
	return schema.GroupVersion{Group: d.Group, Version: d.Version.Version}
}

// GroupVersionKing is the unique id of this crd
func (d Crd) GroupVersionKind() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: d.Group, Version: d.Version.Version, Kind: d.KindName}
}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func (d Crd) Kind(kind string) schema.GroupKind {
	return d.GroupVersion().WithKind(kind).GroupKind()
}

func (d CrdMeta) GroupKind() schema.GroupKind {
	return schema.GroupKind{Group: d.Group, Kind: d.KindName}
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func (d Crd) Resource(resource string) schema.GroupResource {
	return d.GroupVersion().WithResource(resource).GroupResource()
}

func (d Crd) SchemeBuilder() runtime.SchemeBuilder {
	return runtime.NewSchemeBuilder(func(scheme *runtime.Scheme) error {
		scheme.AddKnownTypeWithName(d.GroupVersion().WithKind(d.KindName), &v1.Resource{})
		scheme.AddKnownTypeWithName(d.GroupVersion().WithKind(d.KindName+"List"), &v1.ResourceList{})

		metav1.AddToGroupVersion(scheme, d.GroupVersion())
		return nil
	})
}

func (d Crd) AddToScheme(s *runtime.Scheme) error {
	registerLock.Lock()
	defer registerLock.Unlock()
	builder := d.SchemeBuilder()
	return (&builder).AddToScheme(s)
}
