package webpage

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ParseDocFromURL(url string, handler func(doc *goquery.Document, err error) error) error {
	if handler == nil {
		return fmt.Errorf("没有指定 Handler")
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	// Request the HTML page.
	res, err := client.Get(url)
	if err != nil {
		handler(nil, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		handler(nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status))
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	return handler(doc, err)
}
