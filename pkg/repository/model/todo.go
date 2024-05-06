package model

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Title       string
	Done        bool
	Description string
	UserId      uint
}

type ToDos []ToDo
