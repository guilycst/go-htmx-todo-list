package htmx

import (
	"fmt"
	"net/http"
)

func (hx *HTMXHandler) Edit(w http.ResponseWriter, r *http.Request) {
	raw, id, err := getIdFromPath(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("id \"%s\" is invalid", raw), http.StatusBadRequest)
		return
	}

	found, err := hx.srv.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if found == nil {
		http.Error(w, fmt.Sprintf("id \"%d\" not found", id), http.StatusNotFound)
		return
	}

	err = hx.tmpl.ExecuteTemplate(w, "list_item_edit.html", ToView(*found))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
