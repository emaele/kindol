package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/emaele/kind-ol/utility"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("810364826:AAENsdTQ1hlDUl1ZTZgxMdeTc4wsXOwvB_g")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	deals, err := utility.RetrieveDeals()
	if err != nil {
		log.Fatal(err)
	}

	for _, deal := range deals {
		bot.Send(tgbotapi.NewMessage(-1001263015029, fmt.Sprintf("%s.\n%s.\nPrezzo: %s.\n\n%s", deal.Title, deal.Author, deal.Price, deal.Link)))
	}

}
