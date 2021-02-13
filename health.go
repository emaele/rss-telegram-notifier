package main

import (
	"net/http"
)

func healthCheck(writer http.ResponseWriter, _ *http.Request) {
	writer.Write([]byte("healthy!"))
}
