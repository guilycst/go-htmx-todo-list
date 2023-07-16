package handlers

import (
	"fmt"
	"net/http"
)

func delete(w http.ResponseWriter, r *http.Request) {
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

	service.Delete(found)

	w.WriteHeader(http.StatusAccepted)
}
