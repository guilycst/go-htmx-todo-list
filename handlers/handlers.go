package handlers

import (
	"html/template"
	"net/http"

	"github.com/guilycst/go-htmx/services"
)

var (
	service *services.TodoService
	tmpl    *template.Template
)

func Start(srv *services.TodoService, tmp *template.Template) {
	service = srv
	tmpl = tmp
	http.HandleFunc("/", indexHandleFunc)
	http.HandleFunc("/add", addHandleFunc)
	http.HandleFunc("/list", listHandleFunc)
	http.HandleFunc("/done/", doneHandleFunc(true))
	http.HandleFunc("/undone/", doneHandleFunc(false))
	http.HandleFunc("/delete/", delete)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/update/", update)
	http.HandleFunc("/dist/", fileServerHandler)
	http.HandleFunc("/assets/", fileServerHandler)
}
