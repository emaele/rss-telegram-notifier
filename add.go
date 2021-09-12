package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/emaele/rss-telegram-notifier/entities"
	"github.com/mmcdole/gofeed"
)

func (b *Backstore) addFeed(writer http.ResponseWriter, request *http.Request) {

	defer func() {
		err := request.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var addRequest entities.AddFeedRequest

	// decoding request in struct
	err := json.NewDecoder(request.Body).Decode(&addRequest)
	if err != nil {
		log.Printf("Decode failed due to: %v", err)
		writeHTTPResponse(http.StatusInternalServerError, "", writer)
		return
	}

	// parsing url from body
	feed, err := b.feedparser.ParseURL(addRequest.URL)
	if err != nil {
		log.Printf("Parse URL failed due to: %v", err)
		writeHTTPResponse(http.StatusUnprocessableEntity, "", writer)
		return
	}

	parseYoutubeFeeds(feed)

	// parsing the regex
	reg, err := regexp.Compile(addRequest.Filter)
	if err != nil {
		log.Printf("Regex Compile failed due to: %v", err)
		writeHTTPResponse(http.StatusUnprocessableEntity, "", writer)
		return
	}

	// creating new feed
	rssfeed := entities.RssFeed{
		Title:       feed.Title,
		Description: feed.Description,
		URL:         feed.FeedLink,
		Filter:      reg.String(),
	}

	// if there are rows affected we have a duplicate
	if feedExists(b.db, feed.Link) {
		writeHTTPResponse(http.StatusUnprocessableEntity, "duplicate!", writer)
		log.Printf("rss feed %s is a duplicate\n", addRequest.URL)
		return
	}

	// Adding feed to db
	err = createFeed(b.db, &rssfeed)
	if err != nil {
		writeHTTPResponse(http.StatusInternalServerError, "unable to add feed", writer)
		log.Printf("error creating feed %s, %v\n", rssfeed.Title, err)
		return
	}

	feedID := retrieveFeedID(b.db, rssfeed.URL)

	// fetching and filtering initial elements
	filteredItems := make([]*gofeed.Item, 0, len(feed.Items))

	for index, itm := range feed.Items {

		if reg.MatchString(itm.Title) {
			filteredItems = append(filteredItems, feed.Items[index])
		}
	}

	// setting them to true so we don't get spammed
	addItems(b.db, feedID, filteredItems, true)

	_, err = writer.Write([]byte("added"))
	if err != nil {
		log.Printf("error writing response to %s, %v", request.RemoteAddr, err)
	}
}
