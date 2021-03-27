package main

import (
	"fmt"

	"github.com/spf13/viper"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type command interface {
	canRun(update tgbotapi.Update) bool
	run(update tgbotapi.Update)
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("bbot")

	botToken := viperEnvVariable("token")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}

	var commands = []command{
		boobCommand{bot},
		yesCommand{bot},
	}

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		for _, c := range commands {
			if c.canRun(update) {
				c.run(update)
			}
		}
	}
}

func viperEnvVariable(key string) string {
	value, ok := viper.Get(key).(string)

	if !ok {
		fmt.Printf("Invalid type assertion")
	}

	return value
}
