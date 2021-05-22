package main

import (
	"fmt"
	"log"

	"github.com/emaele/rss-telegram-notifier/entities"
	"github.com/emaele/rss-telegram-notifier/types"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mmcdole/gofeed"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initBackstore() (b Backstore) {
	configuration := readVars()

	feedParser := gofeed.NewParser()

	var err error

	connstring := fmt.Sprintf(types.Mariadbdsn, configuration.DBUser, configuration.DBPassword, configuration.DBHost, configuration.DBPort, configuration.DBName)
	db, err := gorm.Open(mysql.Open(connstring), &gorm.Config{})
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
	bot, err := tg.NewBotAPI(configuration.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	return Backstore{
		conf:       &configuration,
		bot:        bot,
		feedparser: feedParser,
		db:         db,
	}
}
