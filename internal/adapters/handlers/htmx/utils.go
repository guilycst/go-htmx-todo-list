package htmx

import (
	"net/http"
	"strconv"
	"strings"
)

func getIdFromPath(r *http.Request) (string, uint, error) {
	pathSegments := strings.Split(r.URL.Path, "/")
	raw := pathSegments[2]
	id, err := strconv.ParseUint(raw, 10, 32)
	if err != nil {
		return raw, 0, err
	}

	return raw, uint(id), err
}
