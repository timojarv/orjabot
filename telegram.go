package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func RunBot() {
	token := os.Getenv("TG_API_TOKEN")
	fmt.Println(token)
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "safkaa":
			handleSafkaa(bot, update.Message)
		default:
		}

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		//bot.Send(msg)
	}
}

func handleSafkaa(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	rs, err := FetchRestaurants()
	if err != nil {
		return
	}

	rs.Filter([]string{"Reaktori", "Newton", "Hertsi", "Café Konehuone Såås Bar"})

	newMsg := tgbotapi.NewMessage(msg.Chat.ID, rs.String())
	newMsg.ParseMode = "markdown"

	bot.Send(newMsg)
}
