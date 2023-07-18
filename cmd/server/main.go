package main

import (
	"log"
	"net/http"
	"os"

	"github.com/guilycst/go-htmx/internal/adapters/handlers/fileserver"
	"github.com/guilycst/go-htmx/internal/adapters/handlers/htmx"
	"github.com/guilycst/go-htmx/internal/adapters/repositories"
	"github.com/guilycst/go-htmx/internal/core/domain"
	"github.com/guilycst/go-htmx/internal/core/services/todosrv"
	"github.com/guilycst/go-htmx/pkg/loadenv"
)

func init() {

	// Try to load .env file if any
	loadenv.LoadEnv()

	var (
		pgConnStr string = os.Getenv("PG_CONN_STR")
		tmplDir   string = os.Getenv("TEMPLATES_DIR")
		distDir   string = os.Getenv("DIST_DIR")
		pubDir    string = os.Getenv("PUB_DIR")
	)

	//Create new repository
	repository, err := repositories.NewTodoDBRepository[domain.TodoItem](pgConnStr)
	if err != nil {
		log.Fatal(err)
	}

	//Create service
	srv := todosrv.New(repository)

	//Initialize http handlers
	handler, err := htmx.NewHTMXHandler(srv, tmplDir)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler.IndexHandleFunc)
	http.HandleFunc("/add", handler.AddHandleFunc)
	http.HandleFunc("/list", handler.ListHandleFunc)
	http.HandleFunc("/done/", handler.DoneHandleFunc(true))
	http.HandleFunc("/undone/", handler.DoneHandleFunc(false))
	http.HandleFunc("/delete/", handler.Delete)
	http.HandleFunc("/edit/", handler.Edit)
	http.HandleFunc("/update/", handler.Update)

	distFsh, err := fileserver.NewFileServerHandler(distDir)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/dist/", distFsh)

	pubFsh, err := fileserver.NewFileServerHandler(pubDir)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/public/", pubFsh)
}

func main() {
	// Start the HTTP server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
