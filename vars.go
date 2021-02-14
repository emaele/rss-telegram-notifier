package main

import (
	"github.com/mmcdole/gofeed"
	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
)

var (
	// Config params
	bindAddress    string
	telegramToken  string
	telegramChatID int64

	bot        *tb.Bot
	feedParser *gofeed.Parser

	db      *gorm.DB
	dbpath  string
	dbdebug bool
)
