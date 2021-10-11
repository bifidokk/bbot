package main

import (
	"fmt"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
)

type command interface {
	canRun(update tgbotapi.Update) bool
	run(update tgbotapi.Update)
}

func main() {
	botToken := getEnv("token")
	webhookUrl := getEnv("webhook_url")
	baseUrl := getEnv("base_url")

	webhookPath := getEnv("webhook_path")
	if len(webhookPath) == 0 {
		webhookPath = "/"
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}

	var commands = []command{
		boobCommand{bot},
		yesCommand{bot},
	}

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(webhookUrl + webhookPath + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook(webhookPath + bot.Token)
	go http.ListenAndServe(baseUrl, nil)

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
