package main

import (
	"encoding/json"
	"net/http"
)

func getFeeds(writer http.ResponseWriter, request *http.Request) {
	var feeds []rssfeed

	db.Find(&feeds)

	json.NewEncoder(writer).Encode(feeds)
	return
}
