package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/guilycst/go-htmx/internal/adapters/handlers/fileserver"
	"github.com/guilycst/go-htmx/internal/adapters/handlers/htmx"

	_ "time/tzdata"

	"github.com/guilycst/go-htmx/internal/core/services/todosrv"
	"github.com/guilycst/go-htmx/pkg/loadenv"
	"github.com/guilycst/go-htmx/pkg/repo"
)

func init() {

	//Parse flags
	env := flag.String("env", "", ".env file")
	flag.Parse()

	// Try to load .env file if any
	loadenv.LoadEnv(env)

	var (
		storage repo.Storage = repo.StorageFromString(os.Getenv("STORAGE"))
		connStr string       = os.Getenv("CONN_STR")
	)

	//Create new repository
	repository, err := repo.GetRepo(storage, connStr)
	if err != nil {
		log.Fatal(err)
	}

	//Create service
	srv := todosrv.New(repository)

	//Initialize http handlers
	handler, err := htmx.NewHTMXHandler(srv)
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

	distFsh, err := fileserver.NewFileServerHandler("./dist")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/dist/", distFsh)

	assetsFsh, err := fileserver.NewFileServerHandler("./assets")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/assets/", assetsFsh)
}

func main() {
	// Start the HTTP server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
