package main

import (
	"log"
	"os"
	"reflect"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/joho/godotenv"
)

//Bot create logic for work with telegram api.
func Bot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

			switch update.Message.Text {

			case "/start":
				msg := tgbotapi.NewMessage(
					update.Message.Chat.ID,
					"Привет, я бот для закупок, че ты хочешь...")
				bot.Send(msg)

			case "/help":
				msg := tgbotapi.NewMessage(
					update.Message.Chat.ID,
					"Это помощь!")
				bot.Send(msg)
			}
		}
	}
}

// init is invoked before main()
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(" No .env file found!")
	}

	if err := CreateTable(); err != nil {
		panic(err)
	}
}

// main function
func main() {
	// Call Bot
	Bot()
}
