package htmx

import (
	"log"
	"net/http"
)

func (hx *HTMXHandler) ListHandleFunc(w http.ResponseWriter, r *http.Request) {
	items, err := hx.srv.All()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = hx.tmpl.ExecuteTemplate(w, "list.html", items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
