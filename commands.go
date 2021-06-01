package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

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

	if name == "orja" || name == "Orja" || name == "@unisafkabot" {
		newPhoto := tgbotapi.NewPhotoUpload(msg.Chat.ID, "./hattu.jpg")
		bot.Send(newPhoto)
		data.AddNosto(msg.Chat.ID, "orja")
		return
	}

	if !data.IsMember(msg.Chat.ID, name) {
		sendMsg(bot, msg.Chat.ID, "Ketä?")
		return
	}

	if name == "@"+msg.From.UserName {
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
		newMessage = "*Nostoja: on (viimeiset 7 päivää)*\n" + newMessage
		sendMsg(bot, msg.Chat.ID, newMessage)
		return
	} else {
		newMessage = "*Nostoja: ei ole.*"
		sendMsg(bot, msg.Chat.ID, newMessage)
		return
	}
}

func handleSafkaa(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	rs, err := FetchRestaurants()
	if err != nil {
		return
	}

	rs.Filter([]string{"Reaktori", "Newton", "Hertsi", "Café Konehuone Såås Bar", "Reaktori (iltaruoka)"})

	sendMsg(bot, msg.Chat.ID, rs.String())
}

func handleKeitto(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	args := strings.SplitN(msg.CommandArguments(), " - ", 3)

	formatError := "Homman idea ois et laitat päivämäärän, keittoindeksin ja nimen tällain:\n*6.9.2019 - 4 - Iso keitto ja vihree lonkero*\nen lähe tulkkailee mitää sun omia sähellyksiä."

	if len(args) < 3 {
		sendMsg(bot, msg.Chat.ID, formatError)
		return
	}

	date, err := time.Parse("2.1.2006", args[0])
	index, err2 := strconv.Atoi(args[1])

	if err != nil || err2 != nil {
		sendMsg(bot, msg.Chat.ID, formatError)
		return
	}

	if index > 5 {
		sendMsg(bot, msg.Chat.ID, "Turhan rankka keittoindeksi (max 5)")
		return
	}

	date = time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, time.Local)

	err = data.AddKeitto(msg.Chat.ID, data.Keitto{
		Name:        args[2],
		Date:        date.Unix(),
		KeittoIndex: index,
	})

	responses := []string{"Asia kunnossa!", "Homma okei.", "Juoma kirkas!", "Job is bueno!", "Käskystä", "Keittovuori kasvaa!"}

	if err != nil {
		sendMsg(bot, msg.Chat.ID, err.Error())
	} else {
		sendMsg(bot, msg.Chat.ID, responses[int(time.Now().Unix())%len(responses)])
	}

}

func handleKeitot(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	keitot, err := data.GetKeittoList(msg.Chat.ID)

	if err != nil {
		sendMsg(bot, msg.Chat.ID, err.Error())
		return
	}

	response := "*Tukevat keitot:*\n"

	for _, keitto := range *keitot {
		response += "• " + keitto.String() + "\n"
	}

	sendMsg(bot, msg.Chat.ID, response)
}

func handleMoti(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	bot.Send(moti(msg))
}

func handleKiitos(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	kiitokset, err := data.AddKiitos()
	if err != nil {
		sendMsg(bot, msg.Chat.ID, err.Error())
		return
	}

	sendMsg(bot, msg.Chat.ID, fmt.Sprintf("Kiitosta! Orja on koettu hyödylliseksi %d kertaa.", kiitokset))
}

func handle3Keitto(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	tyyppi := "Pikku"
	if rand.Intn(2) == 1 {
		tyyppi = "Iso"
	}

	keitto := []string{"keitto", "märkä", "märpqä", "soppa"}[rand.Intn(4)]

	sendMsg(bot, msg.Chat.ID, fmt.Sprintf("%s %s!", tyyppi, keitto))
}
