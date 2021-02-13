package main

import (
	"log"

	"github.com/mmcdole/gofeed"
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

	db.AutoMigrate(&rssfeed{})
	db.AutoMigrate(&rsselement{})

	go fetchElements()
}
