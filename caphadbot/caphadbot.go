package main

import (
	"log"
	"sync"
	"telegram/GeneralBot"

	"github.com/icelain/jokeapi"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	joke := jokeapi.New()

	bot, cfg := GeneralBot.LoadBot(".")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	var wg sync.WaitGroup
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := handleUpdate(update, bot, cfg, joke)
			if err != nil {
				log.Printf("An error occured: %s", err.Error())
			}
		}()
	}

	wg.Wait()
}
