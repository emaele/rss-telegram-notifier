package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/emaele/rss-telegram-notifier/entities"
)

func addFeed(writer http.ResponseWriter, request *http.Request) {

	defer func() {
		err := request.Body.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	var addRequest entities.AddFeedRequest

	err := json.NewDecoder(request.Body).Decode(&addRequest)
	if err != nil {
		log.Panic(err)
		writeHTTPResponse(http.StatusInternalServerError, "", writer)
		return
	}

	// parsing url from body
	feed, err := feedParser.ParseURL(addRequest.URL)
	if err != nil {
		log.Panic(err)
		writeHTTPResponse(http.StatusUnprocessableEntity, "", writer)
		return
	}

	rssfeed := entities.RssFeed{
		Title:       feed.Title,
		Description: feed.Description,
		URL:         feed.FeedLink,
	}

	var f entities.RssFeed
	rows := db.Where(entities.RssFeed{URL: rssfeed.URL}).Find(&f).RowsAffected

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

	db.Select("ID").Where(entities.RssFeed{URL: rssfeed.URL}).Find(&f)

	// fetching initial elements
	// setting them to true so we don't get spammed
	addItems(f.ID, feed.Items, true)

	writer.Write([]byte("added"))
	return
}
