package video

import (
	"SomeTools/webpage"
	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli/v2"
	"net/url"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "video",
		Usage: "下载视频",
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
			case "spankbang.com":
				handler = spankParser
			default:
				break
			}
			return webpage.ParseDocFromURL(uri, handler)
		},
	}
}

func spankParser(doc *goquery.Document, err error) error {
	if err != nil {
		return err
	}

	return nil
}
