package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func setup(bindAddress string) *http.Server {

	router := mux.NewRouter()

	// setting up routes
	router.HandleFunc("/health", healthCheck).Methods(http.MethodGet)
	router.HandleFunc("/add", addFeed).Methods(http.MethodPost)
	router.HandleFunc("/feeds", getFeeds).Methods(http.MethodGet)
	router.HandleFunc("/feeds/{id}", getElements).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    bindAddress,
		Handler: router,

		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	return srv
}
