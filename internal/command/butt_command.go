package command

import (
	"github.com/bifidokk/bbot/internal/service"
	"log"
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	ButtApiURL   = "http://api.obutts.ru/butts/"
	ButtMediaURL = "http://media.obutts.ru/"
)

type ButtCommand struct {
	Bot   *tgbotapi.BotAPI
	Photo *service.PhotoApi
}

func (c ButtCommand) CanRun(update tgbotapi.Update) bool {
	ln := strings.ToLower(update.Message.Text)

	return strings.Contains(ln, "жоп")
}

func (c ButtCommand) Run(update tgbotapi.Update) {
	feed, err := c.Photo.GetRandomItem(ButtApiURL)

	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range feed.Items {
		filePath, err := service.DownloadFileFromURL(ButtMediaURL + item.Preview)

		if err != nil {
			log.Println(err)
			return
		}

		msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, filePath)
		_, err = c.Bot.Send(msg)

		if err != nil {
			return
		}
	}
}
