package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/guilycst/go-htmx/internal/core/domain"
	"github.com/guilycst/go-htmx/internal/core/ports"
	"github.com/guilycst/go-htmx/pkg/loadenv"
	"github.com/guilycst/go-htmx/pkg/repo"
)

var population = []*domain.TodoItem{}
var repository ports.TodoRepository

func init() {

	//Parse flags
	populationFile := flag.String("file", "", "JSON file containing population")
	env := flag.String("env", "", ".env file")
	flag.Parse()

	// Try to load .env file if any
	loadenv.LoadEnv(env)

	if populationFile == nil {
		log.Fatal("No population file provided (flag -file)")
	}

	//Read population file to memory
	data, err := os.ReadFile(*populationFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &population)

	if len(population) == 0 {
		log.Fatal("File is empty or in incorrect format")
	}

	//Create new repository
	var connStr string = os.Getenv("CONN_STR")
	var storage repo.Storage = repo.StorageFromString(os.Getenv("STORAGE"))
	//Create new repository
	pRepository, err := repo.GetRepo(storage, connStr)
	if err != nil {
		log.Fatal(err)
	}
	repository = *pRepository
}

func main() {
	err := repository.SaveBatch(population)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("💾✔️ - Database populated!")
}
