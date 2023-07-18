package ports

import "github.com/guilycst/go-htmx/internal/core/domain"

type Repository[T any] interface {
	FindById(id any) *T
	All() []T
	Save(data *T)
	Delete(data *T)
	SaveBatch(data []*T)
}

type TodoService interface {
	FindById(id uint) *domain.TodoItem
	All() []domain.TodoItem
	Save(item *domain.TodoItem) error
	Delete(item *domain.TodoItem)
}
