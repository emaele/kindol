package main

import (
	"flag"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zpnk/go-bitly"
	"gitlab.com/emaele/kind-ol/utility"
)

var (
	telegramToken string
	bitlyToken    string
)

func main() {

	setCLIParams()

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// BitLy authentication
	bitly := bitly.New(bitlyToken)

	deals, err := utility.RetrieveDeals(bitly)
	if err != nil {
		log.Fatal(err)
	}

	for _, deal := range deals {
		bot.Send(tgbotapi.NewMessage(-1001263015029, fmt.Sprintf("%s.\n%s.\nPrezzo: %s.\n\n%s", deal.Title, deal.Author, deal.Price, deal.Link)))
	}

}

func setCLIParams() {
	flag.StringVar(&telegramToken, "telegram", "", "Telegram BOTApi token")
	flag.StringVar(&bitlyToken, "bitly", "", "BitLy token")
	flag.Parse()
}
