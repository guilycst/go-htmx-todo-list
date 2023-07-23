package htmx

import (
	"embed"
	"html/template"

	"github.com/guilycst/go-htmx/internal/core/domain"
	"github.com/guilycst/go-htmx/internal/core/ports"
)

//go:embed templates
var tmplFs embed.FS

type HTMXHandler struct {
	srv  ports.TodoService
	tmpl *template.Template
}

type todoItemView struct {
	domain.TodoItem
	Order int64
}

func ToView(t domain.TodoItem) todoItemView {
	order := t.CreatedAt.Unix()
	if t.Done {
		order = t.UpdatedAt.Unix()
	}

	return todoItemView{
		t,
		order,
	}
}

func NewHTMXHandler(srv ports.TodoService) (*HTMXHandler, error) {
	//Parse templates
	funcs := template.FuncMap(template.FuncMap{
		"attr": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})

	tmpl, err := template.New("todo").Funcs(funcs).ParseFS(tmplFs, "templates/*.html")
	if err != nil {
		return nil, err
	}

	return &HTMXHandler{
		srv:  srv,
		tmpl: tmpl,
	}, nil
}
