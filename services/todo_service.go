package services

import (
	"github.com/guilycst/go-htmx/models"
	"gorm.io/gorm"
)

type TodoService struct {
	db *gorm.DB
}

func (s *TodoService) FindById(id uint) *models.TodoItem {
	var item models.TodoItem
	s.db.First(&item, id)
	return &item
}

func (s *TodoService) All() []models.TodoItem {
	var items []models.TodoItem
	s.db.Find(&items)
	return items
}

func (s *TodoService) Save(item *models.TodoItem) {
	s.db.Save(&item)
}

func (s *TodoService) Delete(item *models.TodoItem) {
	s.db.Delete(&item)
}

func StartTodo(database *gorm.DB) *TodoService {
	var s = TodoService{
		db: database,
	}
	return &s
}
