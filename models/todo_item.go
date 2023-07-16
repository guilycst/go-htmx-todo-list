package models

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Title       string
	Description string
	Done        bool
}
