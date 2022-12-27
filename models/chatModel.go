package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Title    string
	Messages []Message
	Users    []UserChat
}
