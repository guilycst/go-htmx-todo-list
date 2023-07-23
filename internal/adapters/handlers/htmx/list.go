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

	viewItems := []todoItemView{}
	for _, v := range items {
		viewItems = append(viewItems, ToView(v))
	}

	err = hx.tmpl.ExecuteTemplate(w, "list.html", viewItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
