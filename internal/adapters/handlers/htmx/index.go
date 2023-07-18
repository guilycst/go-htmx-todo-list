package htmx

import "net/http"

func (hx *HTMXHandler) IndexHandleFunc(w http.ResponseWriter, r *http.Request) {
	err := hx.tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
