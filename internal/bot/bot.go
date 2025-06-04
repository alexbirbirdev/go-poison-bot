package bot

import (
	"log"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"github.com/alexbirbirdev/go-poison-bot/internal/handlers"
)

var userLastMsgTime = make(map[int64]time.Time)

func Start() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Не загрузился .env")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		return err
	}

	bot.Debug = false
	log.Printf("Бот запущен: @%s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		now := time.Now()

		lastTime, exists := userLastMsgTime[chatID]
		if exists && now.Sub(lastTime) < 2*time.Second {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "ОКАК🐈‍⬛: Вы отправляете запросы слишком часто!")
			bot.Send(msg)
			continue
		}
		userLastMsgTime[chatID] = now

		input := strings.TrimSpace(update.Message.Text)

		log.Printf("[%s]: %s", update.Message.From.UserName, update.Message.Text)

		switch input {
		case "/start":
			handlers.Start(bot, update)

		case "/rate":
			handlers.Rate(bot, update)
		default:
			handlers.Price(bot, update)
		}
	}

	return nil
}
