package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"👋 Привет! Я — Poison калькулятор.\n\n"+
			"💡 Просто отправь мне цену товара в юанях, а я посчитаю итоговую цену в рублях по формуле:\n"+
			"`(стоимость товара в юанях * курc юаня) + доставка + комиссия`\n\n"+
			"Отправь, например: `799`\n"+
			"и я скажу точную цену.")

	msg.ParseMode = "Markdown"
	bot.Send(msg)
}
