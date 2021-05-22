package main

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
)

var (
	bot        *tg.BotAPI
	feedParser *gofeed.Parser

	db *gorm.DB
)
