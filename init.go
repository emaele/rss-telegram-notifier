package main

import (
	"flag"
	"fmt"
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mmcdole/gofeed"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/emaele/rss-telegram-notifier/entities"
	"github.com/emaele/rss-telegram-notifier/types"
)

func init() {
	configuration := readVars()

	feedParser = gofeed.NewParser()

	var err error

	connstring := fmt.Sprintf(types.Mariadbdsn, configuration.DBUser, configuration.DBPassword, configuration.DBHost, configuration.DBName)
	db, err = gorm.Open(mysql.Open(connstring), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	err = db.AutoMigrate(&entities.RssFeed{})
	if err != nil {
		log.Panic(err)
	}

	err = db.AutoMigrate(&entities.RssItem{})
	if err != nil {
		log.Panic(err)
	}

	// initializing telegram bot
	bot, err = tg.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}
}

func setCliParams(dbpath *string) {
	flag.StringVar(dbpath, "db", "rss.db", "database file path")
	flag.Parse()
}
