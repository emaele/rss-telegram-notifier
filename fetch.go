package main

import (
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

func fetchElements() {
	for range time.NewTicker(30 * time.Minute).C {
		// get feeds
		var feeds []rssFeed

		rows := db.Find(&feeds).RowsAffected
		log.Printf("found %d feeds to check for\n", rows)

		for _, f := range feeds {
			// fetching elements
			feed, err := feedParser.ParseURL(f.URL)
			if err != nil {
				log.Panic(err)
			}

			log.Printf("found %d elements for %s", len(feed.Items), feed.Title)

			addItems(f.ID, feed.Items)
		}
	}
}

func addItems(feedID uint, items []*gofeed.Item) {
	for _, feedelement := range items {
		element := rssItem{
			Title:       feedelement.Title,
			Description: feedelement.Description,
			URL:         feedelement.Link,
			Sent:        false,
			Feed:        feedID,
		}

		err := db.Where(rssItem{URL: element.URL}).FirstOrCreate(&element).Error
		if err != nil {
			log.Panic(err)
		}
	}
}
