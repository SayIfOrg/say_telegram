package main

import (
	"gopkg.in/telebot.v3/middleware"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Use(middleware.Logger())

	IncludeRoutes(b)

	b.Start()
}
