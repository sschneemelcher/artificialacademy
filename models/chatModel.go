package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	// TODO
	Owner uint
}
