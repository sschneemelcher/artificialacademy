package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Verified bool
}
