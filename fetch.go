package main

import (
	"log"
	"regexp"
	"time"

	"github.com/mmcdole/gofeed"
)

func fetchElements() {
	// starting new ticker, we're going to check for new feed items every 15 minutes
	for range time.NewTicker(15 * time.Minute).C {

		// get feeds
		feeds, err := retrieveFeeds()
		if err != nil {
			log.Printf("unable to retrieve feeds, %v\n", err)
			continue
		}

		log.Printf("found %d feeds to check for\n", len(feeds))

		for _, f := range feeds {

			// fetching elements
			feed, parserr := feedParser.ParseURL(f.URL)
			if parserr != nil {
				log.Printf("unable to fetch items for %s\n", f.Title)
				continue
			}

			// filtering elements
			reg := regexp.MustCompile(f.Filter)

			filteredItems := make([]*gofeed.Item, len(feed.Items), 0)
			
			for _, itm := range feed.Items {
				if reg.MatchString(itm.Title) {
					filteredItems = append(filteredItems, itm)
				}
			}

			log.Printf("found %d elements for %s\n", len(filteredItems), feed.Title)

			// adding to db
			addItems(f.ID, filteredItems, false)
		}
	}
}
