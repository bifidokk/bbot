package main

import (
	"log"
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	ButtApiURL   = "http://api.obutts.ru/butts/"
	ButtMediaURL = "http://media.obutts.ru/"
)

type buttCommand struct {
	bot *tgbotapi.BotAPI
	photo *PhotoApi
}

func (c buttCommand) canRun(update tgbotapi.Update) bool {
	ln := strings.ToLower(update.Message.Text)

	return strings.Contains(ln, "жоп")
}

func (c buttCommand) run(update tgbotapi.Update) {
	feed, err := c.photo.GetRandomItem(ButtApiURL )

	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range feed.Items {
		filePath, err := downloadFileFromURL(ButtMediaURL + item.Preview)

		if err != nil {
			log.Println(err)
			return
		}

		msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, filePath)
		c.bot.Send(msg)
	}
}