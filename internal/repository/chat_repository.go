package repository

import (
	"github.com/bifidokk/bbot/internal/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

type ChatRepository interface {
	Create(chat *entity.Chat) (*entity.Chat, error)
	FindByTelegramId(id string) (*entity.Chat, error)
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &userRepository{
		database: db,
	}
}

func (r *userRepository) Create(chat *entity.Chat) (*entity.Chat, error) {
	result := r.database.Create(&chat)

	return chat, result.Error
}

func (r *userRepository) FindByTelegramId(id string) (*entity.Chat, error) {
	var chat *entity.Chat
	result := r.database.Where(&entity.Chat{TelegramID: id}).First(&chat)

	return chat, result.Error
}
