package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (b *Backstore) getFeeds(writer http.ResponseWriter, _ *http.Request) {
	feeds, err := retrieveFeeds(b.db)
	if err != nil {
		writeHTTPResponse(http.StatusNotFound, "no feeds", writer)
		return
	}

	err = json.NewEncoder(writer).Encode(feeds)
	if err != nil {
		log.Println(err)
	}
}
