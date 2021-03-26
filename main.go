package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
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

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("bbot")

	botToken := viperEnvVariable("token")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		ln := strings.ToLower(update.Message.Text)

		if strings.Contains(ln, "сиськ") {
			feed, _ := getRandomItem()

			for _, item := range feed.Items {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, MediaURL+item.Preview)
				bot.Send(msg)
			}
		} else if ln == "да" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "пизда")
			bot.Send(msg)
		}
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

func viperEnvVariable(key string) string {
	value, ok := viper.Get(key).(string)

	if !ok {
		fmt.Printf("Invalid type assertion")
	}

	return value
}
