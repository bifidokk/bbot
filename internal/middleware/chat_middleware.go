package middleware

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"gorm.io/gorm"
	"log"
)

type chatMiddleware struct {
	database *gorm.DB
}

type ChatMiddleware interface {
	Handle(update tgbotapi.Update)
}

func NewChatMiddleware(db *gorm.DB) ChatMiddleware {
	return &chatMiddleware{
		database: db,
	}
}

func (m *chatMiddleware) Handle(update tgbotapi.Update) {
	log.Println(update.UpdateID)
}
