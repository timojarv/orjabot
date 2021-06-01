package main

import (
	"log"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/timojarv/orjabot/data"
)

const Sopqa int64 = -1001412227493

// RunBot runs the bot
func RunBot() {
	token := os.Getenv("TG_API_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on Telegram account %s", bot.Self.UserName)

	go RockYourDayEveryDay(bot)

	sendMsg(bot, Sopqa, "Kiitos ryhmään pääsystä")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		name := update.Message.From.FirstName + " " + update.Message.From.LastName
		chatID := strconv.Itoa(int(update.Message.Chat.ID))

		log.Printf("#%s (%s) [@%s - %s] %s", chatID, update.Message.Chat.Title, update.Message.From.UserName, name, update.Message.Text)

		// This bot persists messages
		msg := data.Message{
			Chat:      chatID,
			Message:   update.Message.Text,
			Username:  update.Message.From.UserName,
			Author:    name,
			Timestamp: time.Now(),
		}

		data.SaveMessage(msg)

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
		case "moti":
			handleMoti(bot, update.Message)
		case "kiitos":
			handleKiitos(bot, update.Message)
		default:
		}
	}
}

func RockYourDayEveryDay(bot *tgbotapi.BotAPI) {
	helsinki, _ := time.LoadLocation("Europe/Helsinki")
	for {
		if hour, min, _ := time.Now().In(helsinki).Clock(); hour == 6 && min == 0 {
			sendMsg(bot, Sopqa, "Rock your day!")
			time.Sleep(23 * time.Hour)
		}

		time.Sleep(30 * time.Second)
	}
}
