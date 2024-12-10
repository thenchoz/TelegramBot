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
	end bool,
	err error,
) {
	// default value -> non ending call
	end = false

	switch command {
	case "start":
		msg, err = GeneralBot.BotHelper(ctx, bot.tgBot)
		msg = "Welcome " + user.UserName + "!\n" + msg

	case "help":
		msg, err = GeneralBot.BotHelper(ctx, bot.tgBot)

	case "insult":
		msg, err = insulted(ctx)

	case "joke":
		msg, err = joking(ctx, bot)

	case "spell":
		msg, err = spell(ctx, bot, false)

	case "spell_explained":
		msg, err = spell(ctx, bot, true)

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
