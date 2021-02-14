package main

import (
	"encoding/json"
	"net/http"
)

func getFeeds(writer http.ResponseWriter, request *http.Request) {
	var feeds []rssfeed

	rows := db.Find(&feeds).RowsAffected

	if rows == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(feeds)
	return
}
