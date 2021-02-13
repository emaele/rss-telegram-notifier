package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getElements(writer http.ResponseWriter, request *http.Request) {
	var elements []rsselement

	vars := mux.Vars(request)

	feed, ok := vars["id"]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("request feed is not found"))
		return
	}

	db.Where("Feed = ?", feed).Find(&elements)

	json.NewEncoder(writer).Encode(elements)
	return
}
