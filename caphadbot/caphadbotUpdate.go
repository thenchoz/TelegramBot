package main

import (
	"telegram/GeneralBot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/icelain/jokeapi"
)

func handleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI, cfg GeneralBot.Config, joke *jokeapi.JokeAPI) (err error) {
	switch {
	case update.Message != nil:
		msg := handleMessage(update.Message, bot, cfg, joke)
		_, err = bot.Send(msg)
		return err

	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
		return err

	case update.InlineQuery != nil:
		inline := handleInline(update.InlineQuery, bot, cfg, joke)
		_, err = bot.Request(inline)
		return err

	default:
		return err
	}
}
