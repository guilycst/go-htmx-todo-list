package htmx

import (
	"html/template"
	"path/filepath"

	"github.com/guilycst/go-htmx/internal/core/ports"
	"github.com/guilycst/go-htmx/pkg/dirutil"
)

type HTMXHandler struct {
	srv  ports.TodoService
	tmpl *template.Template
}

func NewHTMXHandler(srv ports.TodoService, templatesDir string) (*HTMXHandler, error) {

	err := dirutil.IsDir(templatesDir)
	if err != nil {
		return nil, err
	}

	//Parse templates
	templatePattern := filepath.Join(templatesDir, "*.html")
	tmpl, err := template.ParseGlob(templatePattern)
	if err != nil {
		return nil, err
	}

	return &HTMXHandler{
		srv:  srv,
		tmpl: tmpl,
	}, nil
}
