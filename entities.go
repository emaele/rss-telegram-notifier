package main

import (
	"gorm.io/gorm"
)

type rssItem struct {
	gorm.Model  `json:"-"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	URL         string `json:"URL" gorm:"not null,unique"`
	Sent        bool   `json:"-"`
	Feed        uint   `json:"-" gorm:"not null"`
}

type rssFeed struct {
	gorm.Model  `json:"-"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	URL         string `json:"URL" gorm:"not null,unique"`
}

type addFeedRequest struct {
	URL string `json:"URL"`
}
