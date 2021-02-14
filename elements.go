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

	rows := db.Where("Feed = ?", feed).Find(&elements).RowsAffected

	if rows == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(elements)
	return
}
