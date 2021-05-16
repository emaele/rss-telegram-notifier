package main

import (
	"log"
)

func main() {
	// starting fetch routine
	go fetchElements()

	// starting notify routine
	go notificationRoutine()

	server := setup(bindAddress)

	log.Println("server listening...")
	log.Panic(server.ListenAndServe())
}
