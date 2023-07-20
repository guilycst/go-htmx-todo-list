package htmx

import (
	"embed"
	"html/template"

	"github.com/guilycst/go-htmx/internal/core/ports"
)

//go:embed templates
var tmplFs embed.FS

type HTMXHandler struct {
	srv  ports.TodoService
	tmpl *template.Template
}

func NewHTMXHandler(srv ports.TodoService, templatesDir string) (*HTMXHandler, error) {
	//Parse templates
	tmpl, err := template.ParseFS(tmplFs, "templates/*.html")
	if err != nil {
		return nil, err
	}

	return &HTMXHandler{
		srv:  srv,
		tmpl: tmpl,
	}, nil
}
