package main

import (
	"github.com/cookiesvanilli/TelegramBot_golang/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("MyToken")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("ConsumerKEy")
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot, pocketClient, "http://localhost:8000")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

//cmd - для сборки бинарных файлов, main
//pkg - вся логика приложения
//go mod init github.com/...
//git init
//ls -> cat go.mod
