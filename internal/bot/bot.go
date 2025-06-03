package bot

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"github.com/alexbirbirdev/go-poison-bot/internal/exchange"
)

func Start() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Не загрузился .env")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		return err
	}

	bot.Debug = true
	log.Printf("Бот запущен: @%s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		input := strings.TrimSpace(update.Message.Text)
		yuanAmount, err := strconv.ParseFloat(input, 64)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, введите только число — стоимость в юанях.")
			bot.Send(msg)
			continue
		}

		rate, err := exchange.GetCNYRate()
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при получении курса юаня"+err.Error())
			bot.Send(msg)
			continue
		}

		exchangeRate := rate + 0.8
		delivery := 2000.0
		comission := 1000.0

		rubPrice := yuanAmount*exchangeRate + delivery + comission

		response := fmt.Sprintf("Цена в рублях: %.0f₽", rubPrice)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		bot.Send(msg)
	}

	return nil
}
