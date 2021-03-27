package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	ApiURL   = "http://api.oboobs.ru/boobs/"
	MediaURL = "http://media.oboobs.ru/"
)

type Feed struct {
	Items []Item
}

type Item struct {
	Preview string
}

type boobCommand struct {
	bot *tgbotapi.BotAPI
}

func (c boobCommand) canRun(update tgbotapi.Update) bool {
	ln := strings.ToLower(update.Message.Text)

	return strings.Contains(ln, "сиськ")
}

func (c boobCommand) run(update tgbotapi.Update) {
	feed, _ := getRandomItem()

	for _, item := range feed.Items {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, MediaURL+item.Preview)
		c.bot.Send(msg)
	}
}

func getRandomItem() (*Feed, error) {
	url := ApiURL + "0/1/random"
	feed, error := requestItems(url)

	return feed, error
}

func requestItems(url string) (*Feed, error) {
	resp, err := http.Get(url)

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
