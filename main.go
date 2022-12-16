package main

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/kirsle/configdir"
	"github.com/mmcdole/gofeed"
	"log"
	"os"
	"path/filepath"
)

func lastUpdated(u string) string {
	configPath := configdir.LocalConfig("pinboard2markdown")
	file, err := os.ReadFile(filepath.Join(configPath, slug.Make(u)))
	if os.IsNotExist(err) {
		return "0"
	} else {
		return string(file)
	}
}

func setLastUpdated(u string, d string) {
	configPath := configdir.LocalConfig("pinboard2markdown")
	err := configdir.MakePath(configPath) // Ensure it exists.
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filepath.Join(configPath, slug.Make(u)), []byte(d), 0600)
	if err != nil {
		log.Fatalf("Failed writing last updated data: %s", err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: pinboard2markdown <feed url>")
		return
	}
	feedUrl := os.Args[1]
	log.Printf("loading data from %s\n", feedUrl)
	//	https://feeds.pinboard.in/rss/u:wheresalice/

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedUrl)

	latest := lastUpdated(feedUrl)
	newLatest := lastUpdated(feedUrl)
	for _, item := range feed.Items {
		if item.Published > latest {
			fmt.Printf("## [%s](%s)\n", item.Title, item.Link)
			fmt.Println()
			fmt.Printf("> %s\n", item.Description)

			if item.Published > newLatest {
				newLatest = item.Published
			}
		}
	}
	setLastUpdated(feedUrl, newLatest)
}
