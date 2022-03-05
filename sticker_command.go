package main

import (
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type stickerCommand struct {
	bot *tgbotapi.BotAPI
}

func (c stickerCommand) canRun(update tgbotapi.Update) bool {
	ln := strings.ToLower(update.Message.Text)

	return strings.Contains(ln, "сбербанк") || strings.Contains(ln, "сбер")
}

func (c stickerCommand) run(update tgbotapi.Update) {
	msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, "CAACAgIAAxkBAAN0YiPGHCMjFixKheQZK4K6XsvVV2IAAvYNAAJtvrFInmUZw4n6my4jBA")
	c.bot.Send(msg)
}
