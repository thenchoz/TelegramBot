package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleInline(
	ctx context.Context,
	inline *tgbotapi.InlineQuery,
	bot *myBot,
) (
	err error,
) {
	user := inline.From

	if user == nil {
		return
	}

	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: inline.ID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       []interface{}{},
	}

	for i := range 2 {

		var msg, title, description, url string
		switch {
		case strings.HasPrefix("insult", inline.Query):
			title = "Insult"
			msg, _, _, err = handleCommand(ctx, "insult", user, bot)
			description = "Random insult"

		case strings.HasPrefix("joke", inline.Query):
			title = "Joke"
			msg, _, _, err = handleCommand(ctx, "joke", user, bot)
			description = "Random joke"

		case strings.HasPrefix("quote", inline.Query):
			title = "Quote"
			msg, url, _, err = handleCommand(ctx, "quote", user, bot)
			description = "Random quote"

		case strings.HasPrefix("spell", inline.Query):
			title = "Spell"
			msg, _, _, err = handleCommand(ctx, "spell", user, bot)
			description = "Random spell"

		case strings.HasPrefix("spell_explained", inline.Query):
			title = "Spell"
			msg, _, _, err = handleCommand(ctx, "spell_explained", user, bot)
			description = "Random spell explained"

		case strings.HasPrefix("help", inline.Query):
			title = "Helper"
			msg, _, _, err = handleCommand(ctx, "help", user, bot)
			description = "Helper text"

		default:
			msg = "unknown"
			title = "Unknown"
			description = "Unknown command"
		}

		if err != nil {
			return
		}

		id := inline.ID + strconv.Itoa(i)
		article := tgbotapi.NewInlineQueryResultArticle(id, title, msg)
		article.Description = description

		inlineConf.Results = append(inlineConf.Results, article)

		if url != "" {
			photo := tgbotapi.NewInlineQueryResultPhotoWithThumb(id+"pic", url, url)
			photo.Caption = msg
			photo.Description = description + " with picture"
			photo.Title = title

			fmt.Println(photo)

			inlineConf.Results = append(inlineConf.Results, photo)
		}
	}

	_, err = bot.tgBot.Request(inlineConf)

	return
}
