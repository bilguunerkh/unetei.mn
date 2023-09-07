package models

import "gorm.io/gorm"

type Gift struct {
	gorm.Model
	Image string
	Body  string
}
