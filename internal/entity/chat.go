package entity

import "time"

type Chat struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	TelegramID uint      `gorm:"not null" json:"telegram_id"`
	Title      string    `gorm:"type:VARCHAR(255)" json:"title"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Chat) TableName() string {
	return "chat"
}
