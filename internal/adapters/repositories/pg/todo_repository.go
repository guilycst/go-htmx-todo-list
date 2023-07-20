package pg

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

type TodoDBRepository struct {
	db *gorm.DB
}

func (r *TodoDBRepository) FindById(id any) (*domain.TodoItem, error) {
	var data domain.TodoItem
	rs := r.db.First(&data, id)
	return &data, rs.Error
}

func (r *TodoDBRepository) All() ([]domain.TodoItem, error) {
	var data []domain.TodoItem
	rs := r.db.Where("deleted_at is null").Order("done asc").Find(&data)
	return data, rs.Error
}

func (r *TodoDBRepository) Save(data *domain.TodoItem) error {
	rs := r.db.Save(&data)
	return rs.Error
}

func (r *TodoDBRepository) Delete(data *domain.TodoItem) error {
	rs := r.db.Delete(&data)
	return rs.Error
}

func (r *TodoDBRepository) SaveBatch(data []*domain.TodoItem) error {
	rs := r.db.Create(data)
	return rs.Error
}
func (r *TodoDBRepository) Create(data *domain.TodoItem) error {
	return r.Save(data)
}

func NewTodoDBRepository(connStr string) (ports.TodoRepository, error) {
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

	return &TodoDBRepository{
		db: db,
	}, nil
}
