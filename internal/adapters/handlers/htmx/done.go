package htmx

import (
	"fmt"
	"net/http"
)

func (hx *HTMXHandler) DoneHandleFunc(done bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

		found.Done = done
		err = hx.srv.Save(found)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = hx.tmpl.ExecuteTemplate(w, "list_item.html", ToView(*found))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
