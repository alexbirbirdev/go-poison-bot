package main

import (
	"log"

	"github.com/alexbirbirdev/go-poison-bot/internal/bot"
)

func main() {
	if err := bot.Start(); err != nil {
		log.Fatal("Боту плохо", err)
	}
}
