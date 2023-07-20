package todosrv

import (
	"errors"
	"log"
	"time"

	"github.com/guilycst/go-htmx/internal/core/domain"
	"github.com/guilycst/go-htmx/internal/core/ports"
)

var ErrorTitleRequired = errors.New("title is required")
var ErrorInternal = errors.New("internal error")

type service struct {
	repository ports.TodoRepository
}

func (s *service) FindById(id uint) (*domain.TodoItem, error) {
	item, err := s.repository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, ErrorInternal
	}
	return item, nil
}

func (s *service) All() ([]domain.TodoItem, error) {
	items, err := s.repository.All()
	if err != nil {
		log.Println(err)
		return nil, ErrorInternal
	}
	return items, nil
}

func (s *service) Save(item *domain.TodoItem) error {
	if item.Title == "" {
		return ErrorTitleRequired
	}

	if err := s.repository.Save(item); err != nil {
		log.Println(err)
		return ErrorInternal
	}

	return nil
}

func (s *service) Create(item *domain.TodoItem) error {
	if item.Title == "" {
		return ErrorTitleRequired
	}

	if err := s.repository.Create(item); err != nil {
		log.Println(err)
		return ErrorInternal
	}

	return nil
}

func (s *service) Delete(item *domain.TodoItem) error {
	now := time.Now()
	item.DeletedAt = &now

	if err := s.repository.Save(item); err != nil {
		log.Println(err)
		return ErrorInternal
	}

	return nil
}

func New(repository ports.TodoRepository) *service {
	return &service{
		repository: repository,
	}
}
