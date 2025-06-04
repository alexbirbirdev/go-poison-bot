package handlers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alexbirbirdev/go-poison-bot/internal/exchange"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Rate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	rate, err := exchange.GetCNYRate()
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при получении курса юаня: "+err.Error())
		bot.Send(msg)

	}
	deltaStr := os.Getenv("EXCHANGE_YUAN_DELTA")
	delta, err := strconv.ParseFloat(deltaStr, 64)
	if err != nil {
		delta = 0
	}

	adjusted := rate + delta
	response := fmt.Sprintf(
		"📈 Курс ЦБ: %.2f₽\n"+
			"💰 Закупочный курс: %.2f₽",
		rate, adjusted,
	)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	bot.Send(msg)
}
