package main

import "net/http"

func checkTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		// Verifying provided token
		providedToken := request.Header.Get("Authorization")
		if providedToken != authToken {
			writeHTTPResponse(http.StatusUnauthorized, "invalid token", writer)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
