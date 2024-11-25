package GeneralBot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func LoadBot(path string) (bot *tgbotapi.BotAPI, cfg Config) {
	var err error
	cfg, err = LoadConfig(path)
	if err != nil {
		log.Panic(err)
	}

	bot, err = tgbotapi.NewBotAPI(cfg.Telegram_Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = cfg.Bot_Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot, cfg
}

func BotHelper(bot *tgbotapi.BotAPI, cfg Config) string {
	return "Hi, my name is " + bot.Self.UserName + ".\n" + cfg.Bot_Helper
}
