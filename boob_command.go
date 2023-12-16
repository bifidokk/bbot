package main

import (
	"log"
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	ApiURL   = "http://api.oboobs.ru/boobs/"
	MediaURL = "http://media.oboobs.ru/"
	Timeout  = 1000
)

type boobCommand struct {
	bot *tgbotapi.BotAPI
	photo *PhotoApi
}

func (c boobCommand) canRun(update tgbotapi.Update) bool {
	ln := strings.ToLower(update.Message.Text)

	return strings.Contains(ln, "сиськ")
}

func (c boobCommand) run(update tgbotapi.Update) {
	feed, err := c.photo.GetRandomItem(ApiURL)

	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range feed.Items {
		filePath, err := downloadFileFromURL(MediaURL + item.Preview)

		if err != nil {
			log.Println(err)
			return
		}

		msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, filePath)
		c.bot.Send(msg)
	}
}