package command

import (
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type YesCommand struct {
	Bot *tgbotapi.BotAPI
}

func (c YesCommand) CanRun(update tgbotapi.Update) bool {
	ln := strings.ToLower(update.Message.Text)

	return ln == "да"
}

func (c YesCommand) Run(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "пизда")
	_, err := c.Bot.Send(msg)

	if err != nil {
		return
	}
}
