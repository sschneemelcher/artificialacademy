package models

import "gorm.io/gorm"

type UserChat struct {
	gorm.Model
	UserID uint `gorm:"foreignKey:UserID;references:ID"`
	User   User
	ChatID uint `gorm:"foreignKey:ChatID;references:ID"`
	Chat   Chat
}
