package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
)

func initHandlers() {
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

	item := TodoItem{
		Id:          atomic.AddUint32(&idc, 1),
		Title:       title,
		Description: description,
	}
	todos = append(todos, item)

	// Execute the template with the provided data
	err = tmpl.ExecuteTemplate(w, "list_item.html", item)
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
		raw, id, err := getIdFromPath(r)
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

func delete(w http.ResponseWriter, r *http.Request) {
	raw, id, err := getIdFromPath(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("id \"%s\" is invalid", raw), http.StatusBadRequest)
		return
	}

	idx, found := findTodoItem(uint32(id))

	if found == nil {
		http.Error(w, fmt.Sprintf("id \"%d\" not found", id), http.StatusNotFound)
		return
	}

	todos = remove(todos, idx)

	w.WriteHeader(http.StatusAccepted)
}

func edit(w http.ResponseWriter, r *http.Request) {
	raw, id, err := getIdFromPath(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("id \"%s\" is invalid", raw), http.StatusBadRequest)
		return
	}

	_, found := findTodoItem(uint32(id))

	if found == nil {
		http.Error(w, fmt.Sprintf("id \"%d\" not found", id), http.StatusNotFound)
		return
	}

	err = tmpl.ExecuteTemplate(w, "list_item_edit.html", *found)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	raw, id, err := getIdFromPath(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("id \"%s\" is invalid", raw), http.StatusBadRequest)
		return
	}

	idx, found := findTodoItem(uint32(id))

	if found == nil {
		http.Error(w, fmt.Sprintf("id \"%d\" not found", id), http.StatusNotFound)
		return
	}

	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	title := r.Form.Get("title")
	description := r.Form.Get("description")

	if title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
	}

	found.Title = title
	found.Description = description
	todos[idx] = *found

	err = tmpl.ExecuteTemplate(w, "list_item.html", *found)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
