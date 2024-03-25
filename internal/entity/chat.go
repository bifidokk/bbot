package entity

import "time"

type Chat struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	TelegramID string    `gorm:"type:VARCHAR(255) not null" json:"telegram_id"`
	Title      string    `gorm:"type:VARCHAR(255)" json:"title"`
	Type       string    `gorm:"type:VARCHAR(64) not null" json:"type"`
	CreatedAt  time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (Chat) TableName() string {
	return "chat"
}
