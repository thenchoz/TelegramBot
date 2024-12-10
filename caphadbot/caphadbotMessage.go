package main

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/icelain/jokeapi"
)

func handleMessage(
	ctx context.Context,
	message *tgbotapi.Message,
	bot *tgbotapi.BotAPI,
	joke *jokeapi.JokeAPI,
) (
	msg tgbotapi.MessageConfig,
	end bool,
	err error,
) {
	// default value -> non ending call
	end = false
	user := message.From

	if user == nil {
		return
	}

	msg = tgbotapi.NewMessage(message.Chat.ID, "")
	ctx = context.WithValue(ctx, "lang", user.LanguageCode)

	if message.IsCommand() {
		msg.Text, end, err = handleCommand(ctx, message.Command(), user, bot, joke)
	} else {
		text := message.Text
		msg.Text = "Sorry not sorry :" + text
	}

	return
}
