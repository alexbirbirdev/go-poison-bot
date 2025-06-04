package reply

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendReply(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("ошибка при отправке сообщения: %v", err)
	}
}

func ReplyWithError(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, "⚠️ "+text)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка при отправке error-сообщения: %v", err)
	}
}

func WithTyping(bot *tgbotapi.BotAPI, chatID int64) {
	action := tgbotapi.NewChatAction(chatID, tgbotapi.ChatTyping)
	if _, err := bot.Send(action); err != nil {
		log.Printf("Ошибка при отправке typing-action: %v", err)
	}
}
