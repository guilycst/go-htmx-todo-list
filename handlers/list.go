package handlers

import "net/http"

func listHandleFunc(w http.ResponseWriter, r *http.Request) {
	var items = service.All()
	err := tmpl.ExecuteTemplate(w, "list.html", items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
