package xcontrib

import (
	"embed"
	"entgo.io/ent/entc/gen"
	"text/template"
)

var (
	templates *template.Template

	//go:embed template/*
	templateFS embed.FS
)

func initTemplates() {
	var err error
	templates, err = template.New("").
		Funcs(TemplateFuncs).
		Funcs(gen.Funcs).
		ParseFS(templateFS, "template/*.tmpl", "template/edge/*.tmpl", "template/method/*.tmpl")

	if err != nil {
		panic(err)
	}
}
