package main

import (
	"flag"
	"log"
	"time"

	"github.com/emaele/rss-telegram-notifier/entities"
	"github.com/mmcdole/gofeed"
	tb "gopkg.in/tucnak/telebot.v2"
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

	db.AutoMigrate(&entities.RssFeed{})
	db.AutoMigrate(&entities.RssItem{})

	// initializing telegram bot
	bot, err = tb.NewBot(tb.Settings{
		Token:  telegramToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	// starting fetch routine
	go fetchElements()

	// starting notify routine
	go notificationRoutine()
}

func setCliParams() {
	flag.StringVar(&dbpath, "db", "rss.db", "database file path")
	flag.Parse()
}
