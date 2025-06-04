package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexbirbirdev/go-poison-bot/internal/calc"
	"github.com/alexbirbirdev/go-poison-bot/internal/exchange"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Price(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	input := strings.TrimSpace(update.Message.Text)
	yuanAmount, err := strconv.ParseFloat(input, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, введите только число — стоимость в юанях.")
		bot.Send(msg)
		return
	}
	if yuanAmount <= 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Цена не может быть отрицательной… :)")
		bot.Send(msg)
		return
	}

	rate, err := exchange.GetCNYRate()
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при получении курса юаня: "+err.Error())
		bot.Send(msg)
		return
	}

	rubPrice := calc.YuanToRub(yuanAmount, rate)

	response := fmt.Sprintf("Цена в рублях: %.0f₽", rubPrice)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	bot.Send(msg)
}
