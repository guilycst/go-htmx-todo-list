package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/guilycst/go-htmx/internal/adapters/repositories"
	"github.com/guilycst/go-htmx/internal/core/domain"
	"github.com/guilycst/go-htmx/internal/core/ports"
	"github.com/guilycst/go-htmx/pkg/loadenv"
)

var population = []*domain.TodoItem{}
var repository ports.Repository[domain.TodoItem]

func init() {

	// Try to load .env file if any
	loadenv.LoadEnv()

	//Parse flags
	populationFile := flag.String("file", "", "JSON file containing population")
	flag.Parse()

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
	var pgConnStr string = os.Getenv("PG_CONN_STR")
	repository, err = repositories.NewTodoDBRepository[domain.TodoItem](pgConnStr)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	repository.SaveBatch(population)
	log.Print("💾✔️ - Database populated!")
}
