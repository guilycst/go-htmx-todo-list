package handlers

import (
	"net/http"
	"os"
)

func fileServerHandler(w http.ResponseWriter, r *http.Request) {
	// Get the path of the file requested by the client
	filePath := r.URL.Path

	// Open the file
	file, err := os.Open("." + filePath)
	if err != nil {
		// Return a 404 Not Found status if the file doesn't exist
		http.NotFound(w, r)
		return
	}
	defer file.Close()

	// Get file information
	fileInfo, err := file.Stat()
	if err != nil {
		// Return a 500 Internal Server Error status if there's an error getting file info
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Serve the file with its proper content type
	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}
