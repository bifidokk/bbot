package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	ButtApiURL   = "http://api.obutts.ru/butts/"
	ButtMediaURL = "http://media.obutts.ru/"
)

type buttCommand struct {
	bot *tgbotapi.BotAPI
}

func (c buttCommand) canRun(update tgbotapi.Update) bool {
	ln := strings.ToLower(update.Message.Text)

	return strings.Contains(ln, "жоп")
}

func (c buttCommand) run(update tgbotapi.Update) {
	feed, err := c.getRandomItem()

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

func (c buttCommand) getRandomItem() (*Feed, error) {
	url := ButtApiURL + "0/1/random"
	feed, err := c.requestItems(url)

	return feed, err
}

func (c buttCommand) requestItems(url string) (*Feed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*Timeout))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var items []Item
	err = json.Unmarshal(body, &items)

	if err != nil {
		return nil, err
	}

	feed := new(Feed)
	feed.Items = items

	return feed, nil
}
