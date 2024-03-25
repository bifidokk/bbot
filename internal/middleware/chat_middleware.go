package middleware

import (
	"errors"
	"github.com/bifidokk/bbot/internal/entity"
	"github.com/bifidokk/bbot/internal/repository"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type chatMiddleware struct {
	chatRepository repository.ChatRepository
}

type ChatMiddleware interface {
	Handle(update tgbotapi.Update)
}

func NewChatMiddleware(repository repository.ChatRepository) ChatMiddleware {
	return &chatMiddleware{
		chatRepository: repository,
	}
}

func (m *chatMiddleware) Handle(update tgbotapi.Update) {
	chatID := strconv.FormatInt(update.Message.Chat.ID, 10)
	_, err := m.chatRepository.FindByTelegramId(chatID)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		chat := &entity.Chat{
			TelegramID: chatID,
			Title:      update.Message.Chat.Title,
			Type:       update.Message.Chat.Type,
		}

		_, err := m.chatRepository.Create(chat)

		if err != nil {
			log.Println(err)
		}
	}
}
