package main

import (
	"log"
	"net/http"

	"github.com/emaele/rss-telegram-notifier/entities"
	"github.com/gorilla/mux"
)

func deleteFeed(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	feed, ok := vars["id"]
	if !ok {
		writeHTTPResponse(http.StatusNotFound, "request feed is not found", writer)
		return
	}

	var f entities.RssFeed

	rows := db.Where("ID = ?", feed).Find(&f).RowsAffected

	if rows == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	// deleting feed
	err := db.Delete(&f).Error
	if err != nil {
		log.Panic(err)
		writeHTTPResponse(http.StatusInternalServerError, "unable to delete", writer)
		return
	}

	// deleting feed elements
	var elements []entities.RssItem
	db.Where("Feed = ?", f.ID).Find(&elements)

	for _, element := range elements {
		db.Delete(&element)
	}

	return
}
