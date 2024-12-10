package main

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/icelain/jokeapi"
)

func handleUpdate(
	ctx context.Context,
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
	joke *jokeapi.JokeAPI,
) (
	ending bool,
	err error,
) {
	// default value -> not ending call
	ending = false

	switch {
	case update.Message != nil:
		var msg tgbotapi.MessageConfig
		msg, ending, err = handleMessage(ctx, update.Message, bot, joke)
		if err != nil {
			return
		}
		_, err = bot.Send(msg)

	case update.CallbackQuery != nil:
		handleButton(ctx, update.CallbackQuery)

	case update.InlineQuery != nil:
		var inline tgbotapi.InlineConfig
		inline, err = handleInline(ctx, update.InlineQuery, bot, joke)
		if err != nil {
			return
		}
		_, err = bot.Request(inline)
	}

	return
}
