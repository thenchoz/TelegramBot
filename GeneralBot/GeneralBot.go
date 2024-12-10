package GeneralBot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func LoadBot(path string) (bot *tgbotapi.BotAPI, cfg Config, err error) {
	cfg, err = LoadConfig(path)
	if err != nil {
		return
	}

	bot, err = tgbotapi.NewBotAPI(cfg.Telegram_Token)
	if err != nil {
		return
	}

	bot.Debug = cfg.Bot_Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return
}

func BotHelper(ctx context.Context, bot *tgbotapi.BotAPI) (str string, err error) {
	commands, err := bot.GetMyCommands()
	if err != nil {
		return
	}

	helper := bot.Self.UserName + ".\n"
	if cfg := ctx.Value("cfg"); cfg != nil {
		helper += cfg.(Config).Bot_Helper + "\n"
	}

	lang := ctx.Value("lang")
	switch lang {
	case "fr":
		str = "Bonjour, mon nom est "
		str += helper
		str += "Je connais les commandes suivantes :"

	default:
		str = "Hi, my name is "
		str += helper
		str += "I'm aware of the following command:"
	}

	for _, c := range commands {
		str += "\n\t- " + c.Command + " : " + c.Description
	}

	return
}
