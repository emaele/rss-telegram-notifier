package main

import (
	"log"
	"time"
)

func fetchElements() {
	for range time.NewTicker(30 * time.Minute).C {
		// get feeds
		var feeds []rssfeed

		rows := db.Find(&feeds).RowsAffected
		log.Printf("found %d feeds to check for\n", rows)

		for _, f := range feeds {
			// fetching elements
			feed, err := feedParser.ParseURL(f.URL)
			if err != nil {
				log.Panic(err)
			}

			log.Printf("found %d elements for %s", len(feed.Items), feed.Title)

			for _, feedelement := range feed.Items {
				element := rsselement{
					Title:       feedelement.Title,
					Description: feedelement.Description,
					URL:         feedelement.Link,
					Sent:        false,
					Feed:        f.ID,
				}

				db.Where(rsselement{URL: element.URL}).FirstOrCreate(&element)
			}
		}
	}
}
