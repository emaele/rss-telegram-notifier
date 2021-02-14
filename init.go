package main

import (
	"log"
	"time"

	"github.com/mmcdole/gofeed"
	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	readVars()

	feedParser = gofeed.NewParser()

	var err error
	db, err = gorm.Open(sqlite.Open("rss.db"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&rssFeed{})
	db.AutoMigrate(&rssItem{})

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
