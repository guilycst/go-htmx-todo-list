package handlers

import (
	"net/http"

	"github.com/guilycst/go-htmx/models"
)

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

	item := models.TodoItem{
		Title:       title,
		Description: description,
	}

	service.Save(&item)

	err = tmpl.ExecuteTemplate(w, "list_item.html", item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
