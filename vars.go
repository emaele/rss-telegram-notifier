package main

import (
	"gorm.io/gorm"

	"github.com/mmcdole/gofeed"
)

var (
	bindAddress string

	feedParser *gofeed.Parser

	db *gorm.DB
)
