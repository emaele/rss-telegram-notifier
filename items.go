package main

import (
	"encoding/json"
	"net/http"

	"github.com/emaele/rss-telegram-notifier/entities"
	"github.com/gorilla/mux"
)

func getItems(writer http.ResponseWriter, request *http.Request) {
	var items []entities.RssItem

	vars := mux.Vars(request)

	feedID, ok := vars["id"]
	if !ok {
		writeHTTPResponse(http.StatusNotFound, "requested feed is not found", writer)
		return
	}

	items, err := retriveItemsByFeedID(feedID)
	if err != nil {
		writeHTTPResponse(http.StatusInternalServerError, "unable to retrieve items", writer)
		return
	}

	json.NewEncoder(writer).Encode(items)
	return
}
