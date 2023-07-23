package ports

import "github.com/guilycst/go-htmx/internal/core/domain"

type TodoRepository interface {
	FindById(id any) (*domain.TodoItem, error)
	All() ([]domain.TodoItem, error)
	Save(data *domain.TodoItem) error
	Create(data *domain.TodoItem) error
	Delete(data *domain.TodoItem) error
	SaveBatch(data []*domain.TodoItem) error
}

type TodoService interface {
	FindById(id uint) (*domain.TodoItem, error)
	All() ([]domain.TodoItem, error)
	Save(item *domain.TodoItem) error
	Delete(item *domain.TodoItem) error
	Create(item *domain.TodoItem) error
}
