package main

import (
	"flag"
	"log"

	"github.com/emaele/rss-telegram-notifier/entities"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mmcdole/gofeed"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	readVars()
	setCliParams()

	feedParser = gofeed.NewParser()

	var err error

	db, err = gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
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

	// starting fetch routine
	go fetchElements()

	// starting notify routine
	go notificationRoutine()
}

func setCliParams() {
	flag.StringVar(&dbpath, "db", "rss.db", "database file path")
	flag.Parse()
}
