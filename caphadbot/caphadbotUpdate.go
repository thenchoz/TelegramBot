package main

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleUpdate(
	ctx context.Context,
	update tgbotapi.Update,
	bot *myBot,
) (
	ending bool,
	err error,
) {
	// default value -> not ending call
	ending = false

	switch {
	case update.Message != nil:
		var msg tgbotapi.MessageConfig
		msg, ending, err = handleMessage(ctx, update.Message, bot)
		if err != nil {
			return
		}
		_, err = bot.tgBot.Send(msg)

	case update.CallbackQuery != nil:
		handleButton(ctx, update.CallbackQuery)

	case update.InlineQuery != nil:
		var inline tgbotapi.InlineConfig
		inline, err = handleInline(ctx, update.InlineQuery, bot)
		if err != nil {
			return
		}
		_, err = bot.tgBot.Request(inline)
	}

	return
}
