package model

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Title     string `gorm:"title"`
	Completed bool   `gorm:"completed"`
}
