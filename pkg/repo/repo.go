package repo

import (
	"errors"
	"log"
	"strings"

	"github.com/guilycst/go-htmx/internal/adapters/repositories/libsql"
	"github.com/guilycst/go-htmx/internal/adapters/repositories/orm"
	"github.com/guilycst/go-htmx/internal/core/ports"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage string

const (
	POSTGRESQL Storage = "POSTGRESQL"
	LIBSQL     Storage = "LIBSQL"
	SQLITE     Storage = "SQLITE"
	UNKNOWN    Storage = "UNKNOWN"
)

func StorageFromString(str string) Storage {
	str = strings.ToUpper(str)
	switch str {
	case "POSTGRESQL":
		return POSTGRESQL
	case "LIBSQL":
		return LIBSQL
	case "SQLITE":
		return SQLITE
	default:
		log.Printf("No storage type %s!!!\n", str)
		return UNKNOWN
	}
}

var (
	ErrNoDialector = errors.New("Storage type not supported")
)

func GetRepo(stg Storage, connStr string) (*ports.TodoRepository, error) {
	switch stg {
	case LIBSQL:

		libsqlRepo, err := libsql.NewTodoDBRepository(connStr)
		if err != nil {
			return nil, err
		}
		return &libsqlRepo, nil
	default:
		dialector, err := getGormDialector(stg, connStr)
		if err != nil {
			return nil, err
		}

		ormRepo, err := orm.NewTodoDBRepository(dialector)
		if err != nil {
			return nil, err
		}
		return &ormRepo, nil
	}
}

func getGormDialector(stg Storage, connStr string) (gorm.Dialector, error) {
	switch stg {
	case POSTGRESQL:
		return postgres.Open(connStr), nil
	case SQLITE:
		return sqlite.Open(connStr), nil
	default:
		return nil, ErrNoDialector
	}
}
