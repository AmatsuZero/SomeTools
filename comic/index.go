package comic

import (
	"SomeTools/webpage"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "comic",
		Usage: "下载漫画",
		Action: func(c *cli.Context) error {
			uri := ""
			if c.NArg() > 0 {
				uri = c.Args().First()
			}
			u, err := url.Parse(uri)
			if err != nil {
				return err
			}
			var handler func(doc *goquery.Document, err error) error
			switch u.Host {
			case "e-hentai.org":
				handler = eHentaiParser
			default:
				handler = defaultHandler
			}
			return webpage.ParseDocFromURL(uri, handler)
		},
	}
}

func eHentaiParser(doc *goquery.Document, err error) error {
	if err != nil {
		return err
	}

	return nil
}

func defaultHandler(doc *goquery.Document, err error) error {
	if err != nil {
		return err
	}

	return nil
}
