package repositories

import (
	"log"
	"os"
	"time"

	"github.com/guilycst/go-htmx/internal/core/domain"
	"github.com/guilycst/go-htmx/internal/core/ports"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TodoDBRepository[T domain.TodoItem] struct {
	db *gorm.DB
}

func (r *TodoDBRepository[T]) FindById(id any) *T {
	var data T
	r.db.First(&data, id)
	return &data
}

func (r *TodoDBRepository[T]) All() []T {
	var data []T
	r.db.Where("deleted_at is null").Order("done asc").Find(&data)
	return data
}

func (r *TodoDBRepository[T]) Save(data *T) {
	r.db.Save(&data)
}

func (r *TodoDBRepository[T]) Delete(data *T) {
	r.db.Delete(&data)
}

func (r *TodoDBRepository[T]) SaveBatch(data []*T) {
	r.db.Create(data)
}

func NewTodoDBRepository[T domain.TodoItem](connStr string) (ports.Repository[T], error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&domain.TodoItem{})

	return &TodoDBRepository[T]{
		db: db,
	}, nil
}
