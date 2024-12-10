package main

import (
	"context"
	"log"
	"sync"

	"telegram/GeneralBot"

	"github.com/icelain/jokeapi"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	joke := jokeapi.New()

	bot, cfg, err := GeneralBot.LoadBot(".")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	var wg sync.WaitGroup
	updates := bot.GetUpdatesChan(u)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "cfg", cfg)
	ctx, cancel := context.WithCancel(ctx)

	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			wg.Add(1)
			go func() {
				defer wg.Done()
				end, err := handleUpdate(ctx, update, bot, joke)
				if err != nil {
					log.Printf("An error occured: %s", err.Error())
				}
				if end {
					cancel()
					wg.Wait()
				}
			}()
		}
	}
}
