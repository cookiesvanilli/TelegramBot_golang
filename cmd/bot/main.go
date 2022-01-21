package main

import (
	"github.com/boltdb/bolt"
	"github.com/cookiesvanilli/TelegramBot_golang/pkg/config"
	"github.com/cookiesvanilli/TelegramBot_golang/pkg/repository"
	"github.com/cookiesvanilli/TelegramBot_golang/pkg/repository/boltdb"
	"github.com/cookiesvanilli/TelegramBot_golang/pkg/server"
	"github.com/cookiesvanilli/TelegramBot_golang/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient(cfg.PocketConsumerKey)
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, cfg.AuthServerURL, cfg.Messages)

	authorizationServer := server.NewAuthorizationServer(pocketClient, tokenRepository, cfg.TelegramBotURL)
	go func() {
		if err := authorizationServer.Start(); err != nil { //метод старт блокирующая операция, но горутина разрешает эту проблему
			log.Fatal(err)
		}
	}()

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB() (*bolt.DB, error) {
	//create DB
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	//transaction
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return db, nil
}

//cmd - для сборки бинарных файлов, main
//pkg - вся логика приложения
//go mod init github.com/...
//git init
//ls -> cat go.mod
