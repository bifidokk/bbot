package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	ApiURL   = "http://api.oboobs.ru/boobs/"
	MediaURL = "http://media.oboobs.ru/"
	Timeout  = 1000
)

type boobCommand struct {
	bot *tgbotapi.BotAPI
}

func (c boobCommand) canRun(update tgbotapi.Update) bool {
	ln := strings.ToLower(update.Message.Text)

	return strings.Contains(ln, "сиськ")
}

func (c boobCommand) run(update tgbotapi.Update) {
	feed, err := c.getRandomItem()

	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range feed.Items {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, MediaURL+item.Preview)
		c.bot.Send(msg)
	}
}

func (c boobCommand) getRandomItem() (*Feed, error) {
	url := ApiURL + "0/1/random"
	feed, error := c.requestItems(url)

	return feed, error
}

func (c boobCommand) requestItems(url string) (*Feed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*Timeout))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var items []Item
	err = json.Unmarshal(body, &items)

	feed := new(Feed)
	feed.Items = items

	return feed, nil
}
