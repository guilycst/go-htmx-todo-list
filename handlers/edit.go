package handlers

import (
	"fmt"
	"net/http"
)

func edit(w http.ResponseWriter, r *http.Request) {
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

	err = tmpl.ExecuteTemplate(w, "list_item_edit.html", *found)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
