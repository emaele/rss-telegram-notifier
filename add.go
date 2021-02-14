package main

import (
	"io/ioutil"
	"log"
	"net/http"
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
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	bodyString := string(bodyBytes)

	// parsing url from body
	feed, err := feedParser.ParseURL(bodyString)
	if err != nil {
		log.Panic(err)
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	rssFeed := rssfeed{
		Title:       feed.Title,
		Description: feed.Description,
		URL:         feed.FeedLink,
	}

	var f rssfeed
	rows := db.Where(rssfeed{URL: rssFeed.URL}).Find(&f).RowsAffected
	if rows != 0 { // if there are rows affected we have a duplicate
		writer.WriteHeader(http.StatusUnprocessableEntity)
		writer.Write([]byte("duplicate!"))
		return
	}

	// Adding feed to db
	err = db.Create(&rssFeed).Error
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("unable to add feed"))
		log.Panic(err)
		return
	}

	writer.Write([]byte("added"))
	return
}
