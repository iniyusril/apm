package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title       string
	Description string
}
