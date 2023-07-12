package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

const templateDir = "templates/"

type Index struct {
	Title   string
	Header  string
	Content string
	Test    string
}

func main() {
	templatePattern := templateDir + "*.html"

	tmpl, err := template.ParseGlob(templatePattern)
	if err != nil {
		log.Fatal(err)
	}

	// Define an HTTP handler function to render the templates
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Execute the template with the provided data
		err := tmpl.ExecuteTemplate(w, "index.html", Index{
			Title:   "A title",
			Header:  "A header",
			Content: "Some content",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Register the file server handler for dist files
	http.HandleFunc("/dist/", fileServerHandler)

	// Start the HTTP server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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
