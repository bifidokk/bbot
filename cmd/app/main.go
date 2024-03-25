package main

import (
	"github.com/bifidokk/bbot/internal/command"
	"github.com/bifidokk/bbot/internal/config"
	"github.com/bifidokk/bbot/internal/middleware"
	"github.com/bifidokk/bbot/internal/repository"
	"github.com/bifidokk/bbot/internal/service"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
)

type Command interface {
	CanRun(update tgbotapi.Update) bool
	Run(update tgbotapi.Update)
}

type Middleware interface {
	Handle(update tgbotapi.Update)
}

func main() {
	botToken := config.GetEnv("token")
	webhookUrl := config.GetEnv("webhook_url")
	baseUrl := config.GetEnv("base_url")

	webhookPath := config.GetEnv("webhook_path")
	if len(webhookPath) == 0 {
		webhookPath = "/"
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}

	db, err := config.ConnectDb()
	if err != nil {
		panic(err)
	}

	var apiService = service.NewPhotoApi()

	var commands = []Command{
		command.BoobCommand{Bot: bot, Photo: apiService},
		command.ButtCommand{Bot: bot, Photo: apiService},
		command.YesCommand{Bot: bot},
		command.StickerCommand{Bot: bot},
		command.VideoCommand{Bot: bot},
	}

	var chatRepository = repository.NewChatRepository(db)

	var middlewares = []Middleware{
		middleware.NewChatMiddleware(chatRepository),
	}

	log.Printf("Authorized on account %s\n", bot.Self.UserName)
	log.Println(webhookUrl + webhookPath)

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
	go func() {
		err := http.ListenAndServe(baseUrl, nil)
		if err != nil {
			panic(err)
		}
	}()

	for update := range updates {
		for _, m := range middlewares {
			m.Handle(update)
		}

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		for _, c := range commands {
			if c.CanRun(update) {
				c.Run(update)
			}
		}
	}
}
