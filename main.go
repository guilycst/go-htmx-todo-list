package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	http.HandleFunc("/", indexHandleFunc)
	http.HandleFunc("/add", addHandleFunc)
	http.HandleFunc("/list", listHandleFunc)
	http.HandleFunc("/done/", doneHandleFunc(true))
	http.HandleFunc("/undone/", doneHandleFunc(false))
	http.HandleFunc("/dist/", fileServerHandler)
	http.HandleFunc("/assets/", fileServerHandler)

	// Start the HTTP server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandleFunc(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func addHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	title := r.Form.Get("title")
	description := r.Form.Get("description")

	if title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
	}

	todos = append(todos, TodoItem{
		Id:          atomic.AddUint32(&idc, 1),
		Title:       title,
		Description: description,
	})

	// Execute the template with the provided data
	err = tmpl.ExecuteTemplate(w, "list.html", todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func listHandleFunc(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "list.html", todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func doneHandleFunc(done bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pathSegments := strings.Split(r.URL.Path, "/")
		raw := pathSegments[2]
		id, err := strconv.ParseUint(raw, 10, 32)
		if err != nil {
			http.Error(w, fmt.Sprintf("id \"%s\" is invalid", raw), http.StatusBadRequest)
			return
		}

		idx, found := findTodoItem(uint32(id))

		if found == nil {
			http.Error(w, fmt.Sprintf("id \"%d\" not found", id), http.StatusNotFound)
			return
		}

		found.Done = done
		todos[idx] = *found

		err = tmpl.ExecuteTemplate(w, "list.html", todos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func findTodoItem(id uint32) (int, *TodoItem) {
	var found *TodoItem
	var idx int
	for i, v := range todos {
		if v.Id == id {
			idx = i
			found = &v
			break
		}
	}
	return idx, found
}

func fileServerHandler(w http.ResponseWriter, r *http.Request) {
	// Get the path of the file requested by the client
	filePath := r.URL.Path

	// Open the file
	file, err := os.Open("." + filePath)
	if err != nil {
		// Return a 404 Not Found status if the file doesn't exist
		http.NotFound(w, r)
		return
	}
	defer file.Close()

	// Get file information
	fileInfo, err := file.Stat()
	if err != nil {
		// Return a 500 Internal Server Error status if there's an error getting file info
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Serve the file with its proper content type
	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}
