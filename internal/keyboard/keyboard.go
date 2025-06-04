package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetMainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Рассчитать цену"),
			tgbotapi.NewKeyboardButton("Курс юаня"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("О боте"),
		),
	)
}
