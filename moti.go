package main

import (
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/timojarv/orjabot/data"
)

type motiFunc func(*tgbotapi.Message) tgbotapi.Chattable

var motiFuncs = []motiFunc{motiLause, motiLause, motiLause, motiKeitto}

func moti(msg *tgbotapi.Message) tgbotapi.Chattable {
	return motiFuncs[rand.Intn(len(motiFuncs))](msg)
}

var lauseet = []string{
	"Vaikka yrittäisit parhaasi, voit silti olla ihan paska.",
	"Yrittäminen on ensimmäinen askel epäonnistumiseen.",
	"Pidä pää ylhäällä ja muista että loppujen lopuksi elämä on täysin merkityksetöntä.",
	"Meillä jokaisella on vain yksi elämä! Harmi että olet tuhlannut omasi.",
	"Mietin milloin olit viimeksi onnellinen. Se tunne ei enää koskaan palaa.",
	"Huomenna on uusi päivä :) valitettavasti",
	"Ihminen ratkaisee.",
	"Jokainen aamu on uusi mahdollisuus kusta asiat totaalisesti.",
	"Rauhoitu, elämäsi huonoin asia ei ole tapahtunut vielä.",
	"Motivaatio on tärkeää. Jokainen ruumis Mt. Everestin rinteillä oli joskus motivoitunut henkilö.",
	"se että olet uniikki ei tarkoita, että olisit millään tavalla hyödyllinen.",
	"Virheistään ei aina opi, joskus vaan yksinkertaisesti epäonnistuu täydellisesti.",
	"Itsensä haastaminen on erinomainen tapa epäonnistua.",
	"Ehkä elämäsi tarkoitus on toimia varoittavana esimerkkinä muille?",
	"Toivo on ensimmäinen askel kohti pettymystä.",
	"Kaikelle on syynsä. Joskus se syy on että olet vitun tyhmä ja teet vääriä päätöksiä.",
	"Älä ota elämää turhan vakavasti, kuolet kuitenkin.",
}

func motiLause(msg *tgbotapi.Message) tgbotapi.Chattable {
	l := len(lauseet)
	lause := lauseet[rand.Intn(l)]

	return tgbotapi.NewMessage(msg.Chat.ID, lause)
}

func motiKeitto(msg *tgbotapi.Message) tgbotapi.Chattable {
	response := "*Alkoholi on ratkaisu.*\n"

	keitot, err := data.GetKeittoList(msg.Chat.ID)
	if err != nil {
		return motiLause(msg)
	}

	l := len(*keitot)
	if l == 0 {
		return motiLause(msg)
	}

	keitto := (*keitot)[rand.Intn(l)]
	response += keitto.String()

	m := tgbotapi.NewMessage(msg.Chat.ID, response)
	m.ParseMode = "markdown"
	return m
}