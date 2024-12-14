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
	end bool,
	err error,
) {
	// default value -> non ending call
	end = false
	user := message.From

	if user == nil {
		return
	}

	var msg, url string

	if message.IsCommand() {
		msg, url, end, err = handleCommand(ctx, message.Command(), user, bot)
		if err != nil {
			return
		}
	} else {
		msg = "Sorry not sorry :" + message.Text
	}

	if url != "" {
		response := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FileURL(url))
		response.Caption = msg
		response.ParseMode = "markdown"
		_, err = bot.tgBot.Send(response)
	} else {
		response := tgbotapi.NewMessage(message.Chat.ID, msg)
		response.ParseMode = "markdown"
		_, err = bot.tgBot.Send(response)
	}

	return
}
