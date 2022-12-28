package helpers

import (
	"sschneemelcher/artificialacademy/initializers"
	"sschneemelcher/artificialacademy/models"
)

type History struct {
	Name    string
	UserID  uint
	Content string
	ChatID  uint
}
type ChatList struct {
	Title  string
	ChatID uint
}

func GetHistory(chat models.UserChat) ([]History, error) {
	// get history from DB
	history := []History{}
	result := initializers.DB.Model(&models.User{}).
		Select("users.name, messages.user_id, messages.content, messages.chat_id").
		Where("messages.chat_id = ? AND messages.deleted_at IS NULL", chat.ChatID).
		Joins("right join messages on messages.user_id = users.id").
		Order("messages.created_at ASC").
		Find(&history)

	return history, result.Error
}

func GetChatList(user models.User) ([]ChatList, error) {
	chats := []ChatList{}
	result := initializers.DB.Model(&models.Chat{}).
		Select("chats.title, user_chats.chat_id").
		Where("user_chats.user_id = ? AND user_chats.deleted_at IS NULL AND chats.deleted_at IS NULL", user.ID).
		Joins("right join user_chats on user_chats.chat_id = chats.id").
		Order("chats.created_at ASC").
		Find(&chats)

	return chats, result.Error
}