package main

import (
	"fmt"
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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			switch update.Message.Command() {

			case "start":
				msg.Text = "Привет, я бот для закупок, скажи мне, чего ты хочешь ?"

			case "help":
				msg.Text = "Тут будет помощь..."

			case "category":
				categories, err := GetAllCategory()
				if err != nil {
					log.Panic(err)
					msg.Text = "Произошла ошибка, повторите позже"
				}
				for _, category := range categories {
					msg.Text = fmt.Sprintln(category)
					bot.Send(msg)
				}

			default:
				msg.Text = "К сожалению, я не знаю такой команды"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}

// init is invoked before main()
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(" No .env file found!")
	}

	if err := CreateTables(); err != nil {
		panic(err)
	}
}

// main function
func main() {
	// Call Bot
	Bot()
}
