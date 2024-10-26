package main

import (
	"flag"
	"log"
	"read-adviser-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	// token = flafs.Get(token)
	tgClient := telegram.New(tgBotHost, mustToken())
	_ = tgClient

	// tgclient = telegram.New(token)

	// fetcher = fetcher.New(tgclient)

	// processor = processor.New(tgclient)

	// consumer.start(fetcher, processor)

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
