package main

import (
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var b, err = tb.NewBot(tb.Settings{
	// You can also set custom API URL.
	// If field is empty it equals to "https://api.telegram.org".

	Token:  "1849565820:AAH1--d2jzrkmOkF251s91RDXu6WbljAC0U",
	Poller: &tb.LongPoller{Timeout: 10 * time.Second},
})

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic(port)
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", hello)

	b.Start()
}

func hello(m *tb.Message) {
	b.Send(m.Sender, "Hello World!")
}
