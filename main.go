package main

import (
	"log"
)

func main() {
	server := setup(bindAddress)

	log.Println("server listening...")
	log.Panic(server.ListenAndServe())
}
