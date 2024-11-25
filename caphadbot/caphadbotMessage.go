package main

import (
	"strings"

	"telegram/GeneralBot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/icelain/jokeapi"
)

func handleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI, cfg GeneralBot.Config, joke *jokeapi.JokeAPI) (msg tgbotapi.MessageConfig) {
	user := message.From

	if user == nil {
		return
	}

	msg = tgbotapi.NewMessage(message.Chat.ID, "")

	if message.IsCommand() {
		msg.Text = handleCommand(message.Command(), bot, cfg, joke)
	} else {
		text := message.Text
		msg.Text = "Sorry not sorry :" + text
	}

	return msg
}

func handleButton(query *tgbotapi.CallbackQuery) {
	// unused, plan in the futur
}

func handleInline(inline *tgbotapi.InlineQuery, bot *tgbotapi.BotAPI, cfg GeneralBot.Config, joke *jokeapi.JokeAPI) (inlineConf tgbotapi.InlineConfig) {
	user := inline.From

	if user == nil {
		return
	}

	var article tgbotapi.InlineQueryResultArticle
	switch {
	case strings.HasPrefix("insult", inline.Query):
		article = tgbotapi.NewInlineQueryResultArticle(inline.ID, "Insult", insulted())
		article.Description = "Random insult"
		break

	case strings.HasPrefix("joke", inline.Query):
		article = tgbotapi.NewInlineQueryResultArticle(inline.ID, "Joke", joking(joke))
		article.Description = "Random joke"
		break

	case strings.HasPrefix("help", inline.Query):
		article = tgbotapi.NewInlineQueryResultArticle(inline.ID, "Helper", GeneralBot.BotHelper(bot, cfg))
		article.Description = "Helper text"
		break

	default:
		article = tgbotapi.NewInlineQueryResultArticle(inline.ID, "Unknown", "unknown")
		article.Description = "Unknown command"
		break
	}

	inlineConf = tgbotapi.InlineConfig{
		InlineQueryID: inline.ID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       []interface{}{article},
	}

	return inlineConf
}

func handleCommand(command string, bot *tgbotapi.BotAPI, cfg GeneralBot.Config, joke *jokeapi.JokeAPI) (msg string) {
	switch command {
	case "start":
		return "start msg"

	case "help":
		return GeneralBot.BotHelper(bot, cfg)

	case "insult":
		return insulted()

	case "joke":
		return joking(joke)

	default:
		return "Unknown command"
	}
}
