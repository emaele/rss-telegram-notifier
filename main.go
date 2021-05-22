package main

import (
	"log"
)

func main() {

	backstore := initBackstore()

	// starting fetch routine
	go backstore.fetchElements()

	// starting notify routine
	go backstore.notificationRoutine()

	server := setup(backstore.conf.BindAddress, backstore.conf.AuthorizationToken)

	log.Println("server listening...")
	log.Panic(server.ListenAndServe())
}
