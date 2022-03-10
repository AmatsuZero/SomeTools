package webpage

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func ParseDocFromURL(url string, handler func(doc *goquery.Document, err error) error) error {
	if handler == nil {
		return fmt.Errorf("没有指定 Handler")
	}
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		handler(nil, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		handler(nil, err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	return handler(doc, err)
}
