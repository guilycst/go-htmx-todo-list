package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/guilycst/go-htmx/handlers"
	"github.com/guilycst/go-htmx/models"
	"github.com/guilycst/go-htmx/services"
	"gorm.io/gorm"
)

const templateDir = "templates/"

var (
	tmpl    *template.Template
	db      *gorm.DB
	service *services.TodoService
)

func init() {
	//Parse templates
	templatePattern := templateDir + "*.html"

	var templates = template.New("")
	funcMap := template.FuncMap{
		"attr": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}
	templates.Funcs(funcMap)
	templates, err := templates.ParseGlob(templatePattern)
	if err != nil {
		log.Fatal(err)
	}
	tmpl = templates

	tmpl.Funcs(funcMap)

	//Open database conn and run migration
	db = openDBConn()
	db.AutoMigrate(&models.TodoItem{})

	//Initialize todo service
	service = services.StartTodo(db)

	//Initialize http handlers
	handlers.Start(service, tmpl)
}

func main() {

	// Start the HTTP server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
