package main

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/timojarv/orjabot/data"
)

// RunBot runs the bot
func RunBot() {
	token := os.Getenv("TG_API_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on Telegram account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		name := update.Message.From.FirstName + " " + update.Message.From.LastName
		chatID := strconv.Itoa(int(update.Message.Chat.ID))

		log.Printf("#%s (%s) [@%s - %s] %s", chatID, update.Message.Chat.Title, update.Message.From.UserName, name, update.Message.Text)

		// This bot automatically keeps a list of group members
		if update.Message.From.UserName != "" {
			data.AddMember(update.Message.Chat.ID, "@"+update.Message.From.UserName)
		}

		// Command handling
		switch update.Message.Command() {
		case "safkaa":
			handleSafkaa(bot, update.Message)
		case "hattu":
			handleHattu(bot, update.Message)
		case "nostot":
			handleNostot(bot, update.Message)
		case "uusikeitto":
			handleKeitto(bot, update.Message)
		case "keitot":
			handleKeitot(bot, update.Message)
		default:
		}
	}
}
