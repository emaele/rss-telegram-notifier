package main

import (
	"gorm.io/gorm"
)

type rssItem struct {
	gorm.Model
	Title       string `json:"Title"`
	Description string `json:"Description"`
	URL         string `gorm:"not null,unique"`
	Sent        bool
	Feed        uint `gorm:"not null"`
}

type rssFeed struct {
	gorm.Model
	Title       string `json:"Title"`
	Description string `json:"Description"`
	URL         string `json:"URL" gorm:"not null,unique"`
}
