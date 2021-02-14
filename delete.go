package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func deleteFeed(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	feed, ok := vars["id"]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("request feed is not found"))
		return
	}

	var f rssFeed

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
	}

	// deleting feed elements
	var elements []rssItem
	db.Where("ID = ?", f.ID).Find(&elements)

	db.Delete(elements)

	return
}
