package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zpnk/go-bitly"
	"github.com/emaele/kindol/utility"
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

	chatID, err := strconv.ParseInt(os.Getenv("CHATID"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	for _, deal := range deals {
		msg := tgbotapi.NewPhotoShare(chatID, deal.Cover)
		msg.Caption = fmt.Sprintf("📚 %s\n✍️ %s\n\n💶 %s", deal.Title, deal.Author, deal.Price)
		msg.ReplyMarkup = utility.SetupInlineKeyboard(deal.Link)
		bot.Send(msg)
	}
}
