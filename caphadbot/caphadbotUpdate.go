package main

import (
	"context"
	"log"

	"telegram/GeneralBot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel, cfg GeneralBot.Config, bot *tgbotapi.BotAPI) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			err := handleUpdate(update, bot, cfg)
			if err != nil {
				log.Printf("An error occured: %s", err.Error())
			}
		}
	}
}

func handleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI, cfg GeneralBot.Config) (err error) {
	switch {
	case update.Message != nil:
		msg := handleMessage(update.Message, bot, cfg)
		_, err = bot.Send(msg)
		return err

	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
		return err

	case update.InlineQuery != nil:
		inline := handleInline(update.InlineQuery, bot, cfg)
		_, err = bot.Request(inline)
		return err

	default:
		return err
	}
}
