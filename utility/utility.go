package utility

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//URL è il link all'offerta lampo per kindle store
const URL = "https://www.amazon.it/Offerta-Lampo-Kindle/b?ie=UTF8&node=5689487031"

//Deal è la struttura del libro in offerta lampo
type Deal struct {
	Title  string
	Author string
	Price  string
	Link   string
}

func getWebpage() io.ReadCloser {
	resp, err := http.Get(URL)
	if err != nil {
		return nil
	}

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	return resp.Body
}

func RetrieveDeals() ([]Deal, error) {

	doc, err := goquery.NewDocumentFromReader(getWebpage())
	if err != nil {
		log.Fatal(err)
	}

	test := doc.Find(".a-box-group").First()

	var deals []Deal

	test.Find(".a-carousel-col").Each(func(i int, a *goquery.Selection) {

		a.Find("li").Each(func(i int, b *goquery.Selection) {

			var deal Deal

			linksuffix, _ := b.Find("a").Attr("href")
			deal.Link = "https://amazon.it" + cleanLink(linksuffix)

			deal.Title, _ = b.Find("a").Attr("title")

			deal.Author = b.Find(".acs_product-metadata__contributors").Text()
			if deal.Author == "" { // L'autore è ficcato in uno spazio a caso
				deal.Author = strings.TrimSpace(b.Clone().Children().Remove().End().Text())
			}

			price := b.Find(".a-color-price").Text()
			deal.Price = strings.TrimSpace(price)

			deals = append(deals, deal)
		})

	})

	return deals, err
}

func cleanLink(link string) string {
	slices := strings.Split(link, "/")

	newSuffix := ""
	for _, slice := range slices {
		if strings.HasPrefix(slice, "ref") {
			return newSuffix
		}
		newSuffix += slice + "/"
	}

	return newSuffix
}
