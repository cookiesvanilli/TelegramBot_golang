package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"net/url"
)

const (
	commandStart           = "start"
	replyStart             = "Привет! Чтобы вставить ссылки в своём Pocket аккаунте, необходимо дать разрешение на доступ этому боту. Для этого переходи по ссылке: \n%s"
	replyAlreadyAuthorized = "Вы уже авторизированны. Пришлите ссылку, а я её сохраню"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	//log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Ссылка успешно сохранена")
	//msg.ReplyToMessageID = message.MessageID //bot отвечает на конкретное сообщение

	//проверка на валидный формат URL
	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		msg.Text = "Это не валидная ссылка"
		_, err := b.bot.Send(msg)
		return err
	}

	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		msg.Text = "Вы не авторизированны. Используйте команду /start"
		_, err := b.bot.Send(msg)
		return err
	}

	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		msg.Text = "Увы, не удалось сохранить ссылку. Попробуйте ещё раз позже."
		_, err := b.bot.Send(msg)
		return err
	}

	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthProcess(message)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, replyAlreadyAuthorized)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")
	_, err := b.bot.Send(msg)
	return err
}
