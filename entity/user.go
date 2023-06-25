package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string
}
