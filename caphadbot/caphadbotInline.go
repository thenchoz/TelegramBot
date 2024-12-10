package main

import (
	"context"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleInline(
	ctx context.Context,
	inline *tgbotapi.InlineQuery,
	bot *myBot,
) (
	inlineConf tgbotapi.InlineConfig,
	err error,
) {
	user := inline.From

	if user == nil {
		return
	}

	ctx = context.WithValue(ctx, "lang", user.LanguageCode)

	var msg, title, description string
	switch {
	case strings.HasPrefix("insult", inline.Query):
		title = "Insult"
		msg, _, err = handleCommand(ctx, "insult", user, bot)
		description = "Random insult"

	case strings.HasPrefix("joke", inline.Query):
		title = "Joke"
		msg, _, err = handleCommand(ctx, "joke", user, bot)
		description = "Random joke"

	case strings.HasPrefix("spell", inline.Query):
		title = "Spell"
		msg, _, err = handleCommand(ctx, "spell", user, bot)
		description = "Random spell"

	case strings.HasPrefix("spell_explained", inline.Query):
		title = "Spell"
		msg, _, err = handleCommand(ctx, "spell_explained", user, bot)
		description = "Random spell explained"

	case strings.HasPrefix("help", inline.Query):
		title = "Helper"
		msg, _, err = handleCommand(ctx, "help", user, bot)
		description = "Helper text"

	default:
		msg = "unknown"
		title = "Unknown"
		description = "Unknown command"
	}

	if err != nil {
		return
	}

	article := tgbotapi.NewInlineQueryResultArticle(inline.ID, title, msg)
	article.Description = description

	inlineConf = tgbotapi.InlineConfig{
		InlineQueryID: inline.ID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       []interface{}{article},
	}

	return
}
