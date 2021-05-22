package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func setup(bindAddress, authToken string) *http.Server {

	router := mux.NewRouter()

	if authToken != "" {
		router.Use(checkTokenMiddleware)
	}

	// setting up routes
	router.HandleFunc("/health", healthCheck).Methods(http.MethodGet)

	// feed router
	feedRouter := router.PathPrefix("/feed").Subrouter()
	feedRouter.HandleFunc("", getFeeds).Methods(http.MethodGet)
	feedRouter.HandleFunc("/add", addFeed).Methods(http.MethodPost)
	feedRouter.HandleFunc("/{id}", getItems).Methods(http.MethodGet)
	feedRouter.HandleFunc("/delete/{id}", deleteFeed).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    bindAddress,
		Handler: router,

		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	return srv
}
