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

	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Panic(err)
	}
	bodyString := string(bodyBytes)

	feed, err := feedParser.ParseURL(bodyString)
	if err != nil {
		log.Panic(err)
	}

	rssFeed := rssfeed{
		Title:       feed.Title,
		Description: feed.Description,
		URL:         feed.FeedLink,
	}

	db.Debug().Create(&rssFeed)

	writer.Write([]byte("added"))
}
