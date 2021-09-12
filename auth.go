package main

import (
	"log"
	"net/http"
)

func (b *Backstore) checkTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		// Verifying provided token
		providedToken := request.Header.Get("Authorization")
		if providedToken != b.conf.AuthorizationToken {
			writeHTTPResponse(http.StatusUnauthorized, "invalid token", writer)
			log.Printf("%s provided invalid token\n", request.RemoteAddr)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
