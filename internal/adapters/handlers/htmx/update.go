package htmx

import (
	"fmt"
	"net/http"

	"github.com/guilycst/go-htmx/internal/core/services/todosrv"
)

func (hx *HTMXHandler) Update(w http.ResponseWriter, r *http.Request) {
	raw, id, err := getIdFromPath(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("id \"%s\" is invalid", raw), http.StatusBadRequest)
		return
	}

	found, err := hx.srv.FindById(id)
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

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

	found.Title = r.Form.Get("title")
	found.Description = r.Form.Get("description")

	err = hx.srv.Save(found)
	if err != nil {
		if err == todosrv.ErrorTitleRequired {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	err = hx.tmpl.ExecuteTemplate(w, "list_item.html", ToView(*found))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
