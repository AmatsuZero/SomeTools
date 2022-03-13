package main

import (
	"SomeTools/comic"
	"SomeTools/video"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			video.Command(),
			comic.Command(),
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
