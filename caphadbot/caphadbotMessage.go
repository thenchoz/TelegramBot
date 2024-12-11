package main

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMessage(
	ctx context.Context,
	message *tgbotapi.Message,
	bot *myBot,
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

	if message.IsCommand() {
		msg.Text, end, err = handleCommand(ctx, message.Command(), user, bot)
	} else {
		text := message.Text
		msg.Text = "Sorry not sorry :" + text
	}

	return
}
