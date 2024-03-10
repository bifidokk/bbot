package command

import (
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const sberSticker = "CAACAgIAAxkBAAN0YiPGHCMjFixKheQZK4K6XsvVV2IAAvYNAAJtvrFInmUZw4n6my4jBA"

type StickerCommand struct {
	Bot *tgbotapi.BotAPI
}

func (c StickerCommand) CanRun(update tgbotapi.Update) bool {
	ln := strings.ToLower(update.Message.Text)

	return strings.Contains(ln, "сбербанк") || strings.Contains(ln, "сбер")
}

func (c StickerCommand) Run(update tgbotapi.Update) {
	msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, sberSticker)
	_, err := c.Bot.Send(msg)

	if err != nil {
		return
	}
}
