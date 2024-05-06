package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string
	PasswordHash string
	ToDos        []ToDo
}

type Users []User
