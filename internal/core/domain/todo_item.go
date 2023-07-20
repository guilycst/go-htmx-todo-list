package domain

import (
	"time"
)

type TodoItem struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Title       string
	Description string
	Done        bool
}

func (t TodoItem) GetTitle() string {
	return t.Title
}

func (t *TodoItem) pGetTitle() string {
	return t.Title
}
