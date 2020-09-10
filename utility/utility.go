package utility

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/emaele/kindol/types"
	"github.com/zpnk/go-bitly"
)

// RetrieveDeals ottiene le offerte lampo del giorno e le restituisce in un vettore di Deal
func RetrieveDeals(bitly *bitly.Client) ([]types.Deal, error) {

	body := getWebpage(types.OffersURL)
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	var deals []types.Deal

	doc.Find(".a-carousel").First().Find("li").Each(func(i int, b *goquery.Selection) {

		var deal types.Deal

		fmt.Printf("Processing book number: %d\n", i+1)

		linksuffix, _ := b.Find("a").Attr("href")

		deal.Link = ShortenURL(fmt.Sprintf("https://amazon.it%s?&tag=shitposting-21", cleanLink(linksuffix)), bitly)
		deal.Cover = getCover(deal.Link)

		deal.Title = strings.Trim(b.Find(".acs-product-block__product-title").Text(), "\n")
		deal.Title = strings.TrimLeft(deal.Title, " ")

		deal.Author = strings.Trim(b.Find(".acs-product-block__contributor").Text(), "\n")
		deal.Author = strings.TrimLeft(deal.Author, " ")

		price := b.Find(".acs-product-block__price .a-offscreen").First().Text()
		deal.Price = strings.TrimSpace(price)

		deals = append(deals, deal)
	})

	return deals, err
}

func getCover(url string) string {
	body := getWebpage(url)
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	d := doc.Find(".a-dynamic-image.frontImage")
	cover, _ := d.Attr("src")
	return cover
}

func cleanLink(link string) string {
	slices := strings.Split(link, "/")

	var newSuffix string

	for _, slice := range slices {
		if strings.HasPrefix(slice, "ref") {
			return newSuffix
		}
		newSuffix += slice + "/"
	}

	return newSuffix
}
