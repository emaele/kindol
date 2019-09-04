package utility

import (
	"log"

	"github.com/zpnk/go-bitly"
)

// ShortenURL gives you a shorten url with an unshorten url in input
func ShortenURL(urlvar string, b *bitly.Client) string {

	shortURL, err := b.Links.Shorten(urlvar)
	if err != nil {
		log.Panic(err)
	}
	return shortURL.URL
}
