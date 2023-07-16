package handlers

import (
	"fmt"
	"net/http"
)

func doneHandleFunc(done bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

		found.Done = done
		service.Save(found)

		err = tmpl.ExecuteTemplate(w, "list_item.html", found)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
