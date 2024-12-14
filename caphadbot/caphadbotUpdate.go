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
		ending, err = handleMessage(ctx, update.Message, bot)

	case update.InlineQuery != nil:
		err = handleInline(ctx, update.InlineQuery, bot)
	}

	return
}
