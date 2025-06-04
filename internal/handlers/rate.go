package handlers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alexbirbirdev/go-poison-bot/internal/exchange"
	"github.com/alexbirbirdev/go-poison-bot/internal/reply"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Rate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	reply.WithTyping(bot, update.Message.Chat.ID)
	rate, err := exchange.GetCNYRate()
	if err != nil {
		reply.ReplyWithError(bot, update.Message.Chat.ID, "Ошибка при получении курса юаня: "+err.Error())
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

	reply.SendReply(bot, update.Message.Chat.ID, response)
}
