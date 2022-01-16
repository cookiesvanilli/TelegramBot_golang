package main

import (
	"github.com/cookiesvanilli/TelegramBot_golang/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("MyToken")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

//cmd - для сборки бинарных файлов, main
//pkg - вся логика приложения
//go mod init github.com/...
//git init
//ls -> cat go.mod
