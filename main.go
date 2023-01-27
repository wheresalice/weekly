package main

import (
	"github.com/urfave/cli/v2"
	"github.com/wheresalice/weekly/cmd"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "pinboard",
				Aliases: []string{"a"},
				Usage:   "fetch links from pinboard",
				Action: func(cCtx *cli.Context) error {
					cmd.Pinboard()
					return nil
				},
			},
			{
				Name:    "githubstars",
				Aliases: []string{"a"},
				Usage:   "fetch links from github stars",
				Action: func(cCtx *cli.Context) error {
					cmd.GitHubStars()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
