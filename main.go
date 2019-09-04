package main

import (
	"flag"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/emaele/kind-ol/utility"
)

var telegramToken string

func main() {

	setCLIParams()

	bot, err := tgbotapi.NewBotAPI(telegramToken)
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

func setCLIParams() {
	flag.StringVar(&telegramToken, "token", "", "Telegram BOTApi token")
	flag.Parse()
}
