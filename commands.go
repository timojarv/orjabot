package main

import (
	"fmt"
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/timojarv/orjabot/data"
)


func sendMsg(bot *tgbotapi.BotAPI, id int64, msg string) {
	newMsg := tgbotapi.NewMessage(id, msg)
	newMsg.ParseMode = "markdown"
	bot.Send(newMsg)
}

func handleHattu(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	name := msg.CommandArguments()

	if name == "orja" || name == "Orja" {
		sendMsg(bot, msg.Chat.ID, "Voi juku, kiitosta.")
		return
	}

	if !data.IsMember(msg.Chat.ID, name){
		sendMsg(bot, msg.Chat.ID, "Ketä?")
		return
	}

	if name == "@" + msg.From.UserName {
		sendMsg(bot, msg.Chat.ID, "Juu ei :D")
		return
	}

	data.AddNosto(msg.Chat.ID, name)

	sendMsg(bot, msg.Chat.ID, "Ei voi muuta sanoa kuin hattua nostaa!")
}

func handleNostot(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	newMessage := ""

	for k, v := range data.GetNostot(msg.Chat.ID) {
		newMessage += fmt.Sprintf("%s: %dkpl\n", k, v)
	}

	if newMessage != "" {
		newMessage = "*Nostoja: on (viimeiset 7 päivää)*\n" +newMessage
		sendMsg(bot, msg.Chat.ID, newMessage)
		return
	} else {
		newMessage = "*Nostoja: ei ole.*"
		sendMsg(bot, msg.Chat.ID, newMessage)
		return
	}

	sendMsg(bot, msg.Chat.ID, newMessage)
}


func handleSafkaa(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	rs, err := FetchRestaurants()
	if err != nil {
		return
	}

	log.Println(rs)

	rs.Filter([]string{"Reaktori", "Newton", "Hertsi", "Café Konehuone Såås Bar"})

	sendMsg(bot, msg.Chat.ID, rs.String())
}

