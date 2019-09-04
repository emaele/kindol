package main

import (
	"fmt"
	"log"
	"os"

	//"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambda"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zpnk/go-bitly"
	"gitlab.com/emaele/kind-ol/utility"
)

func main() {
	lambda.Start(mainBot)
}

func mainBot() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// BitLy authentication
	bitly := bitly.New(os.Getenv("BITLY"))

	deals, err := utility.RetrieveDeals(bitly)
	if err != nil {
		log.Fatal(err)
	}

	for _, deal := range deals {
		msg := tgbotapi.NewPhotoShare(-1001263015029, deal.Cover)
		msg.Caption = fmt.Sprintf("📚 %s\n✍️ %s\n\n💶 %s", deal.Title, deal.Author, deal.Price)
		msg.ReplyMarkup = utility.SetupInlineKeyboard(deal.Link)
		bot.Send(msg)
	}

}
