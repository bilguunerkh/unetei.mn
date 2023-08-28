package models

import "gorm.io/gorm"

type Litter struct {
	gorm.Model
	Title string
	Body  string
}
