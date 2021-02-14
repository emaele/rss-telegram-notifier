package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getItems(writer http.ResponseWriter, request *http.Request) {
	var items []rssItem

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
