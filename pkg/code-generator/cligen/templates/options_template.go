package templates

import (
	"github.com/solo-io/solo-kit/pkg/code-generator/templateutils"
	"text/template"
)

var OptionsTemplate = template.Must(template.New("p").Funcs(templateutils.Funcs).Parse(`
package options


type Options struct {
	Name string
	Config Config
	{{range .Resources}}{{.Name}} {{.Name}}
	{{end}}
}

type Config struct {
}

{{range .Resources}}type {{.Name}} struct {
}
{{end}}


`))
