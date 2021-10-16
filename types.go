package main

import (
	"github.com/emaele/rss-telegram-notifier/types"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
)

type Backstore struct {
	conf       *types.ConfigurationParameters
	bot        *tg.BotAPI
	feedparser *gofeed.Parser
	db         *gorm.DB
}

var (
	tgMarkdownReservedChars = []string{".", "-", "(", ")", "#", "!", "|", "[", "]", "_", "*", "`", "~", "+"}
)
