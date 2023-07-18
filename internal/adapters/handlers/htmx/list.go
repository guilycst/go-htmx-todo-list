package htmx

import "net/http"

func (hx *HTMXHandler) ListHandleFunc(w http.ResponseWriter, r *http.Request) {
	var items = hx.srv.All()
	err := hx.tmpl.ExecuteTemplate(w, "list.html", items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
