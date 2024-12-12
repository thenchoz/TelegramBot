package main

import (
	"context"
	"math/rand/v2"
	"strings"
)

func insulted() (insult string, err error) {
	l := len(insults)
	insult = insults[rand.IntN(l)]
	return
}

func joking(ctx context.Context, bot *myBot) (res string, err error) {
	response, err := bot.joke.FetchWithContext(ctx)
	if err != nil {
		return
	}
	for _, r := range response.Joke {
		res += r + "\n"
	}
	return
}

func spell(ctx context.Context, bot *myBot, explained bool) (spell string, err error) {
	sp, err := bot.hpAPI.FetchSpell(ctx)
	if err != nil {
		return
	}
	spell = sp.Spell
	if explained {
		spell += ":\n" + sp.Use
	} else {
		spell = "\u2728 " + spell + "! \u2728"
	}
	return
}

func quote(ctx context.Context, bot *myBot) (quote string, url string, err error) {
	q, err := bot.hpAPI.FetchQuote(ctx)
	if err != nil {
		return
	}
	quote = "\"" + q.Quote + "\"\n\n" + "In the " + q.Source + " \"" + q.Story + "\""

	// get character picture - looking only on surname
	bot.hpAPI.SetSearch(strings.Split(q.Speaker, " ")[0])
	c, err := bot.hpAPI.FetchCharacters(ctx)
	if err != nil || len(c) == 0 {
		quote += ", " + q.Speaker
		return quote, url, nil
	}

	url = c[0].Image
	return
}
