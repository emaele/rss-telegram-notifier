package main

import (
	"log"
	"net/http"
)

func healthCheck(writer http.ResponseWriter, _ *http.Request) {
	_, err := writer.Write([]byte("healthy!"))
	if err != nil {
		log.Println(err)
	}
}
