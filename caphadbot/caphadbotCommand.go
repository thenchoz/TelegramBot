package main

import (
	"context"
	"errors"

	"telegram/GeneralBot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleCommand(
	ctx context.Context,
	command string,
	user *tgbotapi.User,
	bot *myBot,
) (
	msg string,
	url string,
	end bool,
	err error,
) {
	// default value -> non ending call
	end = false

	switch command {
	case "start":
		ctx := context.WithValue(ctx, "lang", user.LanguageCode)
		msg, err = GeneralBot.BotHelper(ctx, bot.tgBot)
		msg = "Welcome " + user.UserName + "!\n" + msg

	case "help":
		ctx := context.WithValue(ctx, "lang", user.LanguageCode)
		msg, err = GeneralBot.BotHelper(ctx, bot.tgBot)

	case "insult":
		msg, err = insulting()

	case "joke":
		msg, err = joking(ctx, bot)

	case "spell":
		msg, err = spell(ctx, bot, false)

	case "spell_explained":
		bot.hpAPI.SetLang(user.LanguageCode)
		msg, err = spell(ctx, bot, true)
		bot.hpAPI.Reset()

	case "hpquote":
		bot.hpAPI.SetLang(user.LanguageCode)
		msg, url, err = hpquote(ctx, bot)
		bot.hpAPI.Reset()

	case "quote":
		msg, err = quote(ctx)

	case "stop":
		if cfg := ctx.Value("cfg"); cfg != nil {
			if cfg.(GeneralBot.Config).Admin_Id == user.ID {
				msg = "Goodby"
				end = true
			} else {
				msg = "Unauthorise"
			}
		} else {
			err = errors.New("admin id not found")
		}

	default:
		msg = "Unknown command"
	}

	return
}
