package command

import (
	"github.com/bifidokk/bbot/internal/service"
	"log"
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	ApiURL   = "http://api.oboobs.ru/boobs/"
	MediaURL = "http://media.oboobs.ru/"
)

type BoobCommand struct {
	Bot   *tgbotapi.BotAPI
	Photo *service.PhotoApi
}

func (c BoobCommand) CanRun(update tgbotapi.Update) bool {
	ln := strings.ToLower(update.Message.Text)

	return strings.Contains(ln, "сиськ")
}

func (c BoobCommand) Run(update tgbotapi.Update) {
	feed, err := c.Photo.GetRandomItem(ApiURL)

	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range feed.Items {
		filePath, err := service.DownloadFileFromURL(MediaURL + item.Preview)

		if err != nil {
			log.Println(err)
			return
		}

		msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, filePath)
		c.Bot.Send(msg)
	}
}
