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

	feed, ok := vars["id"]
	if !ok {
		writeHTTPResponse(http.StatusNotFound, "requested feed is not found", writer)
		return
	}

	rows := db.Where("Feed = ?", feed).Find(&items).RowsAffected

	if rows == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(items)
	return
}
