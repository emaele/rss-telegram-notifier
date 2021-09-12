package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (b *Backstore) deleteFeed(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	feedID, ok := vars["id"]
	if !ok {
		writeHTTPResponse(http.StatusInternalServerError, "there was an error deleting the feed", writer)
		return
	}

	feed, err := retrieveFeedByID(b.db, feedID)
	if err != nil {
		log.Println(err)
		writeHTTPResponse(http.StatusNotFound, "unable to delete feed", writer)
		return
	}

	// retrieve feed items first
	items, err := retriveItemsByFeedID(b.db, feedID)
	if err != nil {
		log.Printf("unable to retrieve feed item, %v\n", err)
	} else {
		// deleting feed elements
		for _, item := range items {
			err = deleteItemFromDB(b.db, &item)
			if err != nil {

				// exit delete function if we're unable to delete an item
				log.Printf("unable to delete feed item, %v\n", err)
				writeHTTPResponse(http.StatusInternalServerError, "unable to delete feed", writer)
				return
			}
		}
	}

	// now deleting feed
	err = deleteFeedFromDB(b.db, &feed)
	if err != nil {
		log.Printf("unable to delete feed, %v\n", err)
		writeHTTPResponse(http.StatusInternalServerError, "unable to delete", writer)
		return
	}
}
