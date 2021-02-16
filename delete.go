package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func deleteFeed(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	feedID, ok := vars["id"]
	if !ok {
		writeHTTPResponse(http.StatusInternalServerError, "there was an error deleting the feed", writer)
		return
	}

	feed, err := retrieveFeedByID(feedID)
	if err != nil {
		log.Fatal(err)
		writeHTTPResponse(http.StatusNotFound, "unable to delete feed", writer)
		return
	}

	// retrieve feed items first
	items, err := retriveItemsByFeedID(feedID)
	if err != nil {
		log.Printf("unable to retrieve feed item, %v\n", err)
		writeHTTPResponse(http.StatusInternalServerError, "unable to delete feed", writer)
		return
	}

	// deleting feed elements
	for _, item := range items {
		err := db.Delete(&item).Error
		if err != nil {
			// exit delete function if we're unable to delete an item
			log.Printf("unable to delete feed item, %v\n", err)
			writeHTTPResponse(http.StatusInternalServerError, "unable to delete feed", writer)
			return
		}
	}

	// now deleting feed
	err = db.Delete(&feed).Error
	if err != nil {
		log.Printf("unable to delete feed, %v\n", err)
		writeHTTPResponse(http.StatusInternalServerError, "unable to delete", writer)
		return
	}

	return
}
