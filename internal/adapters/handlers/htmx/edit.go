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

	found := hx.srv.FindById(id)

	if found == nil {
		http.Error(w, fmt.Sprintf("id \"%d\" not found", id), http.StatusNotFound)
		return
	}

	err = hx.tmpl.ExecuteTemplate(w, "list_item_edit.html", *found)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
