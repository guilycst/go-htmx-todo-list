package repo

import (
	"log"

	"github.com/guilycst/go-htmx/internal/adapters/repositories/pg"
	"github.com/guilycst/go-htmx/internal/adapters/repositories/turso"
	"github.com/guilycst/go-htmx/internal/core/ports"
)

func GetRepo(storage string, connStr string, repo *ports.TodoRepository) {
	switch storage {
	case "TURSO":

		tursoRepo, err := turso.NewTodoDBRepository(connStr)
		if err != nil {
			log.Fatal(err)
		}
		*repo = tursoRepo
	default:
		pgRepo, err := pg.NewTodoDBRepository(connStr)
		if err != nil {
			log.Fatal(err)
		}
		*repo = pgRepo
	}
}
