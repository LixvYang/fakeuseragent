package download

import (
	"net/http"
	"log"
	"errors"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrRequest = errors.New("request err")
	ErrResponse = errors.New("response err")
)

// Get html resource by http
func ResponseDownload(url string) (*goquery.Document, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, ErrRequest
	}

	request.Header.Add("User-Agent", "User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	client := http.DefaultClient

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, ErrResponse
	}

	return goquery.NewDocumentFromReader(response.Body)
}

