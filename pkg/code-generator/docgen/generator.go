package docgen

import (
	"bytes"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/ilackarms/protokit"
	code_generator "github.com/solo-io/solo-kit/pkg/code-generator"
	"github.com/solo-io/solo-kit/pkg/code-generator/docgen/options"
	md "github.com/solo-io/solo-kit/pkg/code-generator/docgen/templates/markdown"
	rst "github.com/solo-io/solo-kit/pkg/code-generator/docgen/templates/restructured"
	"github.com/solo-io/solo-kit/pkg/code-generator/model"
)

const fileHeader = `<!-- Code generated by solo-kit. DO NOT EDIT. -->
`

type DocsGen struct {
	DocsOptions options.DocsOptions
	Project     *model.Project
}

// must ignore validate.proto from lyft
// may need to add more here
var ignoredFiles = []string{
	"validate/validate.proto",
	"github.com/solo-io/solo-kit/api/external/validate/validate.proto",
}

func (d *DocsGen) protoSuffix() string {
	if d.DocsOptions.Output == Restructured {
		return ".sk.rst"
	}
	return ".sk.md"
}

func (d *DocsGen) protoFileTemplate() *template.Template {
	if d.DocsOptions.Output == Restructured {
		return rst.ProtoFileTemplate(d.Project)
	}
	return md.ProtoFileTemplate(d.Project)
}

func (d *DocsGen) GenerateFilesForProtoFiles(protoFiles []*protokit.FileDescriptor) (code_generator.Files, error) {

	// Collect names of files that contain resources for which doc gen has to be skipped
	skipMap := make(map[string]bool)
	for _, res := range d.Project.Resources {
		if res.SkipDocsGen {
			skipMap[res.Filename] = true
		}
	}

	var v code_generator.Files
	for suffix, tmpl := range map[string]*template.Template{
		d.protoSuffix(): d.protoFileTemplate(),
	} {
		for _, protoFile := range protoFiles {
			// Skip if file is to be ignored
			var ignore bool
			for _, ignoredFile := range ignoredFiles {
				if protoFile.GetName() == ignoredFile {
					ignore = true
					break
				}
			}
			if ignore {
				continue
			}

			// Skip if the file contains a top-level resource that has to be skipped
			if skipMap[protoFile.GetName()] {
				continue
			}

			content, err := generateProtoFileFile(protoFile, tmpl)
			if err != nil {
				return nil, err
			}
			fileName := protoFile.GetName() + suffix
			v = append(v, code_generator.File{
				Filename: fileName,
				Content:  content,
			})
		}
	}

	return v, nil
}

func GenerateFiles(project *model.Project, docsOptions *options.DocsOptions) (code_generator.Files, error) {
	protoFiles := protokit.ParseCodeGenRequest(project.Request)
	if docsOptions == nil {
		docsOptions = &DocsOptions{
			Output: Markdown,
		}
	}

	docGenerator := DocsGen{
		DocsOptions: *docsOptions,
		Project:     project,
	}

	files, err := docGenerator.GenerateFilesForProject()
	if err != nil {
		return nil, err
	}
	messageFiles, err := docGenerator.GenerateFilesForProtoFiles(protoFiles)
	if err != nil {
		return nil, err
	}
	files = append(files, messageFiles...)

	for i := range files {
		files[i].Content = fileHeader + files[i].Content
	}
	return files, nil
}

func (d *DocsGen) projectSuffix() string {
	if d.DocsOptions.Output == Restructured {
		return ".project.sk.rst"
	}
	return ".project.sk.md"
}

func (d *DocsGen) projectDocsRootTemplate() *template.Template {
	if d.DocsOptions.Output == Restructured {
		return rst.ProjectDocsRootTemplate(d.Project)
	}
	return md.ProjectDocsRootTemplate(d.Project)
}

func (d *DocsGen) GenerateFilesForProject() (code_generator.Files, error) {
	var v code_generator.Files
	for suffix, tmpl := range map[string]*template.Template{
		d.projectSuffix(): d.projectDocsRootTemplate(),
	} {
		content, err := generateProjectFile(d.Project, tmpl)
		if err != nil {
			return nil, err
		}
		v = append(v, code_generator.File{
			Filename: strcase.ToSnake(d.Project.ProjectConfig.Name) + suffix,
			Content:  content,
		})
	}
	return v, nil
}

func generateProjectFile(project *model.Project, tmpl *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, project); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func generateProtoFileFile(protoFile *protokit.FileDescriptor, tmpl *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, protoFile); err != nil {
		return "", err
	}
	return buf.String(), nil
}
