package handlers

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alexbirbirdev/go-poison-bot/internal/calc"
	"github.com/alexbirbirdev/go-poison-bot/internal/exchange"
	"github.com/alexbirbirdev/go-poison-bot/internal/reply"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Price(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	reply.WithTyping(bot, update.Message.Chat.ID)

	input := strings.TrimSpace(update.Message.Text)
	yuanAmount, err := strconv.ParseFloat(input, 64)
	if err != nil {
		reply.SendReply(bot, update.Message.Chat.ID, "Пожалуйста, введите только число — стоимость в юанях.")
		return
	}
	if yuanAmount <= 0 {
		reply.SendReply(bot, update.Message.Chat.ID, "Цена не может быть отрицательной… :)")
		return
	}

	rate, err := exchange.GetCNYRate()
	if err != nil {
		reply.ReplyWithError(bot, update.Message.Chat.ID, "Ошибка при получении курса юаня: "+err.Error())
		return
	}

	rubPrice := calc.YuanToRub(yuanAmount, rate)

	response := fmt.Sprintf("Цена в рублях: %.0f₽", rubPrice)

	reply.SendReply(bot, update.Message.Chat.ID, response)

	shareBtn := tgbotapi.NewInlineKeyboardButtonURL("🤝 Поделиться ботом", os.Getenv("BOT_SHARE_URL"))
	keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(shareBtn))
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Поделись ботом с другом! просто перешли это сообщение :)")
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}
