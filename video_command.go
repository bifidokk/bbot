package main

import (
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/kkdai/youtube/v2"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	CommandPrefix = "/download"
	DownloadDir   = "/tmp"
	FilePattern   = "*.mp4"
)

type videoCommand struct {
	bot *tgbotapi.BotAPI
}

func (c videoCommand) canRun(update tgbotapi.Update) bool {
	ln := strings.ToLower(update.Message.Text)

	return strings.HasPrefix(ln, CommandPrefix) && strings.Contains(ln, "youtube.com/watch")
}

func (c videoCommand) run(update tgbotapi.Update) {
	videoId := extractVideoID(update.Message.Text)
	log.Printf("Check download video with ID %s\n", videoId)

	if videoId != "" {
		go downloadAndSendVideo(c, videoId, update)
	}
}

func extractVideoID(str string) string {
	re := regexp.MustCompile(`(?i)youtube\.com/watch\?v=([a-zA-Z0-9_-]+)`)
	match := re.FindStringSubmatch(str)

	if len(match) >= 2 {
		return match[1]
	}

	return ""
}

func downloadAndSendVideo(c videoCommand, videoId string, update tgbotapi.Update) {
	log.Printf("Start download video with ID %s\n", videoId)
	client := youtube.Client{}

	video, err := client.GetVideo(videoId)
	if err != nil {
		log.Println(err)
		return
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		return
	}

	file, err := os.CreateTemp(DownloadDir, FilePattern)
	defer file.Close()

	if err != nil {
		log.Println(err)
		return
	}

	_, err = io.Copy(file, stream)
	if err != nil {
		log.Println(err)
		return
	}

	videoMsg := tgbotapi.NewVideoUpload(update.Message.Chat.ID, file.Name())
	_, err = c.bot.Send(videoMsg)
	os.Remove(file.Name())

	if err != nil {
		errorMsg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		c.bot.Send(errorMsg)
		log.Println(err)
	}
}
