package templates

import (
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

var funcs = template.FuncMap{
	"join":        strings.Join,
	"lowercase":   strings.ToLower,
	"lower_camel": strcase.ToLowerCamel,
	"upper_camel": strcase.ToCamel,
	"snake":       strcase.ToSnake,
}

var ResourceTemplate = template.Must(template.New("resource").Funcs(funcs).Parse(`package {{ .Project.PackageName }}

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func New{{ .Name }}(namespace, name string) *{{ .Name }} {
	return &{{ .Name }}{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

{{- if $.HasStatus }}

func (r *{{ .Name }}) SetStatus(status core.Status) {
	r.Status = status
}
{{- end }}

func (r *{{ .Name }}) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

{{- if .HasData}}

func (r *{{ .Name }}) SetData(data map[string]string) {
	r.Data = data
}
{{- end}}

type {{ .Name }}List []*{{ .Name }}
type {{ .PluralName }}ByNamespace map[string]{{ .Name }}List

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list {{ .Name }}List) Find(namespace, name string) (*{{ .Name }}, error) {
	for _, {{ lower_camel .Name }} := range list {
		if {{ lower_camel .Name }}.Metadata.Name == name {
			if namespace == "" || {{ lower_camel .Name }}.Metadata.Namespace == namespace {
				return {{ lower_camel .Name }}, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find {{ lower_camel .Name }} %v.%v", namespace, name)
}

func (list {{ .Name }}List) AsResources() resources.ResourceList {
	var ress resources.ResourceList 
	for _, {{ lower_camel .Name }} := range list {
		ress = append(ress, {{ lower_camel .Name }})
	}
	return ress
}

{{ if $.HasStatus -}}
func (list {{ .Name }}List) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, {{ lower_camel .Name }} := range list {
		ress = append(ress, {{ lower_camel .Name }})
	}
	return ress
}
{{- end}}

func (list {{ .Name }}List) Names() []string {
	var names []string
	for _, {{ lower_camel .Name }} := range list {
		names = append(names, {{ lower_camel .Name }}.Metadata.Name)
	}
	return names
}

func (list {{ .Name }}List) NamespacesDotNames() []string {
	var names []string
	for _, {{ lower_camel .Name }} := range list {
		names = append(names, {{ lower_camel .Name }}.Metadata.Namespace + "." + {{ lower_camel .Name }}.Metadata.Name)
	}
	return names
}

func (list {{ .Name }}List) Sort() {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
}

func (list {{ .Name }}List) Clone() {{ .Name }}List {
	var {{ lower_camel .Name }}List {{ .Name }}List
	for _, {{ lower_camel .Name }} := range list {
		{{ lower_camel .Name }}List = append({{ lower_camel .Name }}List, proto.Clone({{ lower_camel .Name }}).(*{{ .Name }}))
	}
	return {{ lower_camel .Name }}List 
}

func (list {{ .Name }}List) ByNamespace() {{ .PluralName }}ByNamespace {
	byNamespace := make({{ .PluralName }}ByNamespace)
	for _, {{ lower_camel .Name }} := range list {
		byNamespace.Add({{ lower_camel .Name }})
	}
	return byNamespace
}

func (byNamespace {{ .PluralName }}ByNamespace) Add({{ lower_camel .Name }} ... *{{ .Name }}) {
	for _, item := range {{ lower_camel .Name }} {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace {{ .PluralName }}ByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace {{ .PluralName }}ByNamespace) List() {{ .Name }}List {
	var list {{ .Name }}List
	for _, {{ lower_camel .Name }}List := range byNamespace {
		list = append(list, {{ lower_camel .Name }}List...)
	}
	list.Sort()
	return list
}

func (byNamespace {{ .PluralName }}ByNamespace) Clone() {{ .PluralName }}ByNamespace {
	return byNamespace.List().Clone().ByNamespace()
}

var _ resources.Resource = &{{ .Name }}{}

// Kubernetes Adapter for {{ .Name }}

func (o *{{ .Name }}) GetObjectKind() schema.ObjectKind {
	t := {{ .Name }}Crd.TypeMeta()
	return &t
}

func (o *{{ .Name }}) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*{{ .Name }})
}

var {{ .Name }}Crd = crd.NewCrd("{{ .Project.GroupName }}",
	"{{ upper_camel .PluralName }}",
	"{{ .Project.GroupName }}",
	"{{ .Project.Version }}",
	"{{ .Name }}",
	"{{ .ShortName }}",
	&{{ .Name }}{})
`))