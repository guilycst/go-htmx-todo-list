package main

import (
	"net/http"
	"strconv"
	"strings"
)

func remove(slice []TodoItem, s int) []TodoItem {
	return append(slice[:s], slice[s+1:]...)
}

func getIdFromPath(r *http.Request) (string, uint64, error) {
	pathSegments := strings.Split(r.URL.Path, "/")
	raw := pathSegments[2]
	id, err := strconv.ParseUint(raw, 10, 32)
	return raw, id, err
}

func findTodoItem(id uint32) (int, *TodoItem) {
	var found *TodoItem
	var idx int
	for i, v := range todos {
		if v.Id == id {
			idx = i
			found = &v
			break
		}
	}
	return idx, found
}
