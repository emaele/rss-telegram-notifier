package main

import "net/http"

func writeHTTPResponse(statusCode int, body string, writer http.ResponseWriter) {

	writer.WriteHeader(statusCode)
	writer.Write([]byte(body))
	return
}
