package orm

import (
	"log"
	"os"
	"time"

	"github.com/guilycst/go-htmx/internal/core/domain"
	"github.com/guilycst/go-htmx/internal/core/ports"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormTodoDBRepository struct {
	db *gorm.DB
}

func (r *GormTodoDBRepository) FindById(id any) (*domain.TodoItem, error) {
	var data domain.TodoItem
	rs := r.db.First(&data, id)
	return &data, rs.Error
}

func (r *GormTodoDBRepository) All() ([]domain.TodoItem, error) {
	var data []domain.TodoItem
	rs := r.db.Where("deleted_at is null").Order("done asc").Find(&data)
	return data, rs.Error
}

func (r *GormTodoDBRepository) Save(data *domain.TodoItem) error {
	rs := r.db.Save(&data)
	return rs.Error
}

func (r *GormTodoDBRepository) Delete(data *domain.TodoItem) error {
	rs := r.db.Delete(&data)
	return rs.Error
}

func (r *GormTodoDBRepository) SaveBatch(data []*domain.TodoItem) error {
	rs := r.db.Create(data)
	return rs.Error
}
func (r *GormTodoDBRepository) Create(data *domain.TodoItem) error {
	return r.Save(data)
}

func NewTodoDBRepository(dialector gorm.Dialector) (ports.TodoRepository, error) {
	db, err := gorm.Open(dialector, &gorm.Config{
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

	return &GormTodoDBRepository{
		db: db,
	}, nil
}
