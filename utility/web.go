package utility

import (
	"io"
	"log"
	"net/http"
)

func getWebpage(url string) io.ReadCloser {
	var client http.Client

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Golang_scraper/1.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	return resp.Body
}
