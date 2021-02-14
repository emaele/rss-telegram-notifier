package main

import (
	"log"
	"os"
	"strconv"
)

func readVars() {
	var ok bool

	bindAddress, ok = os.LookupEnv("RSS_SERVER_BIND_ADDRESS")
	if !ok {
		bindAddress = "localhost:26009"
	}

	telegramToken, ok = os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		log.Fatal("telegram token not found")
	}

	chatidTemp, ok := os.LookupEnv("TELEGRAM_CHAT")
	if !ok {
		log.Fatal("telegram chatid not found")
	}

	var err error
	telegramChatID, err = strconv.ParseInt(chatidTemp, 10, 64)
	if err != nil {
		log.Fatal("telegram chatid is not valid")
	}

	// If not auth token provided we're running in no auth mode
	authToken, ok = os.LookupEnv("AUTHORIZATION_TOKEN")
	if !ok {
		authToken = ""
	}
}
