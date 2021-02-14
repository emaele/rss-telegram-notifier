package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mmcdole/gofeed"
)

func addFeed(writer http.ResponseWriter, request *http.Request) {

	defer func() {
		err := request.Body.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	// reading body as bytes stream
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Panic(err)
		writeHTTPResponse(http.StatusUnprocessableEntity, "", writer)
		return
	}
	bodyString := string(bodyBytes)

	// parsing url from body
	feed, err := feedParser.ParseURL(bodyString)
	if err != nil {
		log.Panic(err)
		writeHTTPResponse(http.StatusUnprocessableEntity, "", writer)
		return
	}

	rssfeed := rssFeed{
		Title:       feed.Title,
		Description: feed.Description,
		URL:         feed.FeedLink,
	}

	var f rssFeed
	rows := db.Where(rssFeed{URL: rssfeed.URL}).Find(&f).RowsAffected

	// if there are rows affected we have a duplicate
	if rows != 0 {
		writeHTTPResponse(http.StatusUnprocessableEntity, "duplicate!", writer)
		return
	}

	// Adding feed to db
	err = db.Create(&rssfeed).Error
	if err != nil {
		writeHTTPResponse(http.StatusInternalServerError, "unable to add feed", writer)
		log.Panic(err)
		return
	}

	// fetching initial elements
	// setting them to true so we don't get spammed
	addItems(f.ID, feed.Items, true)

	writer.Write([]byte("added"))
	return
}

func addItems(feedID uint, items []*gofeed.Item, markAsSent bool) {
	for _, feedelement := range items {
		element := rssItem{
			Title:       feedelement.Title,
			Description: feedelement.Description,
			URL:         feedelement.Link,
			Sent:        markAsSent,
			Feed:        feedID,
		}

		err := db.Where(rssItem{URL: element.URL}).FirstOrCreate(&element).Error
		if err != nil {
			log.Panic(err)
		}
	}
}
