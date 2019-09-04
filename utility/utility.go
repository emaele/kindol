package utility

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/zpnk/go-bitly"
)

//URL è il link all'offerta lampo per kindle store
const URL = "https://www.amazon.it/Offerta-Lampo-Kindle/b?ie=UTF8&node=5689487031"

//Deal è la struttura del libro in offerta lampo
type Deal struct {
	Title  string
	Author string
	Cover  string
	Price  string
	Link   string
}

// RetrieveDeals ottiene le offerte lampo del giorno e le restituisce in un vettore di Deal
func RetrieveDeals(bitly *bitly.Client) ([]Deal, error) {

	body := getWebpage(URL)
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	var deals []Deal

	doc.Find(".a-carousel").First().Find("li").Each(func(i int, b *goquery.Selection) {

		var deal Deal

		linksuffix, _ := b.Find("a").Attr("href")

		deal.Link = ShortenURL("https://amazon.it"+cleanLink(linksuffix)+"?&tag=shitposting-21", bitly)

		deal.Cover = getCover(deal.Link)

		deal.Title, _ = b.Find("a").Attr("title")

		deal.Author = b.Find(".acs_product-metadata__contributors").Text()
		if deal.Author == "" { // L'autore è ficcato in uno spazio a caso
			deal.Author = strings.TrimSpace(b.Clone().Children().Remove().End().Text())
		}

		price := b.Find(".a-color-price").Text()
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
