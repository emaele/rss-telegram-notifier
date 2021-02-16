package main

import (
	"log"
	"time"
)

func fetchElements() {
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
			feed, err := feedParser.ParseURL(f.URL)
			if err != nil {
				log.Panic(err)
			}

			log.Printf("found %d elements for %s\n", len(feed.Items), feed.Title)

			// adding to db
			addItems(f.ID, feed.Items, false)
		}
	}
}
