package todosrv

import (
	"errors"
	"time"

	"github.com/guilycst/go-htmx/internal/core/domain"
	"github.com/guilycst/go-htmx/internal/core/ports"
)

var ErrorTitleRequired = errors.New("title is required")

type service struct {
	repository ports.Repository[domain.TodoItem]
}

func (s *service) FindById(id uint) *domain.TodoItem {
	return s.repository.FindById(id)
}

func (s *service) All() []domain.TodoItem {
	return s.repository.All()
}

func (s *service) Save(item *domain.TodoItem) error {
	if item.Title == "" {
		return ErrorTitleRequired
	}
	s.repository.Save(item)
	return nil
}

func (s *service) Delete(item *domain.TodoItem) {
	now := time.Now()
	item.DeletedAt = &now
	s.repository.Save(item)
}

func New(repository ports.Repository[domain.TodoItem]) *service {
	return &service{
		repository: repository,
	}
}
