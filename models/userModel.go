package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	First_name   string
	Last_name    string
	Phone_number float64
	Email        string `gorm: "unique"`
	Password     string
}
