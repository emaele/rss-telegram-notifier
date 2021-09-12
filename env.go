package main

import (
	"log"
	"os"
	"strconv"

	"github.com/emaele/rss-telegram-notifier/types"
)

func readVars() (conf types.ConfigurationParameters) {
	var ok bool

	conf.BindAddress, ok = os.LookupEnv("RSS_SERVER_BIND_ADDRESS")
	if !ok {
		conf.BindAddress = "0.0.0.0:26009"
	}

	conf.TelegramToken, ok = os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		log.Fatal("telegram token not found")
	}

	chatidTemp, ok := os.LookupEnv("TELEGRAM_CHAT")
	if !ok {
		log.Fatal("telegram chatid not found")
	}

	var err error
	conf.TelegramChatID, err = strconv.ParseInt(chatidTemp, 10, 64)
	if err != nil {
		log.Fatal("telegram chatid is not valid")
	}

	// If not auth token provided we're running in no auth mode
	conf.AuthorizationToken, ok = os.LookupEnv("AUTHORIZATION_TOKEN")
	if !ok {
		conf.AuthorizationToken = ""
	}

	conf.DBUser, ok = os.LookupEnv("DB_USER")
	if !ok {
		log.Fatal("database user not found")
	}

	conf.DBPassword, ok = os.LookupEnv("DB_PASSWORD")
	if !ok {
		log.Fatal("database password not found")
	}

	conf.DBName, ok = os.LookupEnv("DB_NAME")
	if !ok {
		log.Fatal("database name not found")
	}

	conf.DBHost, ok = os.LookupEnv("DB_HOST")
	if !ok {
		log.Fatal("database host not found")
	}

	conf.DBPort, ok = os.LookupEnv("DB_PORT")
	if !ok {
		log.Fatal("database port not found")
	}

	return
}
