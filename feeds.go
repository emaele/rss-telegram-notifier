package main

import (
	"encoding/json"
	"net/http"
)

func getFeeds(writer http.ResponseWriter, request *http.Request) {
	feeds, err := retrieveFeeds()

	if err != nil {
		writeHTTPResponse(http.StatusNotFound, "no feeds", writer)
		return
	}

	json.NewEncoder(writer).Encode(feeds)
	return
}
