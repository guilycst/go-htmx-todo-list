package handlers

import (
	"fmt"
	"net/http"
)

func update(w http.ResponseWriter, r *http.Request) {
	raw, id, err := getIdFromPath(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("id \"%s\" is invalid", raw), http.StatusBadRequest)
		return
	}

	found := service.FindById(id)

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
	service.Save(found)

	err = tmpl.ExecuteTemplate(w, "list_item.html", *found)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
