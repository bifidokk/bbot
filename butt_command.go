package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

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
	feed, _ := c.getRandomItem()

	for _, item := range feed.Items {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, ButtMediaURL+item.Preview)
		c.bot.Send(msg)
	}
}

func (c buttCommand) getRandomItem() (*Feed, error) {
	url := ButtApiURL + "0/1/random"
	feed, error := c.requestItems(url)

	return feed, error
}

func (c buttCommand) requestItems(url string) (*Feed, error) {
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
