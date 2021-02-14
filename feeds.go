package main

import (
	"encoding/json"
	"net/http"
)

func getFeeds(writer http.ResponseWriter, request *http.Request) {
	var feeds []rssFeed

	rows := db.Find(&feeds).RowsAffected

	if rows == 0 {
		writeHTTPResponse(http.StatusNotFound, "no feeds", writer)
		return
	}

	json.NewEncoder(writer).Encode(feeds)
	return
}
