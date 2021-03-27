package main

import (
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type yesCommand struct {
	bot *tgbotapi.BotAPI
}

func (c yesCommand) canRun(update tgbotapi.Update) bool {
	ln := strings.ToLower(update.Message.Text)

	return ln == "да"
}

func (c yesCommand) run(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "пизда")
	c.bot.Send(msg)
}
