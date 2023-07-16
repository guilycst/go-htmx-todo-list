package main

import (
	"html/template"
	"log"
	"net/http"
	"sync/atomic"
)

const templateDir = "templates/"

var (
	idc  uint32
	tmpl *template.Template
)

type TodoItem struct {
	Id          uint32
	Title       string
	Description string
	Done        bool
}

var todos []TodoItem = []TodoItem{
	{
		Id:          atomic.AddUint32(&idc, 1),
		Title:       "Finish project proposal",
		Description: "Due on 4/1/23",
	},
	{
		Id:          atomic.AddUint32(&idc, 1),
		Title:       "Buy groceries",
		Description: "Bananas, milk, bread",
		Done:        true,
	},
	{
		Id:          atomic.AddUint32(&idc, 1),
		Title:       "Go for a run",
		Description: "3 miles",
	},
}

func init() {
	templatePattern := templateDir + "*.html"

	templates, err := template.ParseGlob(templatePattern)
	if err != nil {
		log.Fatal(err)
	}
	tmpl = templates
}

func main() {

	//Initialize http handlers
	initHandlers()

	// Start the HTTP server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
