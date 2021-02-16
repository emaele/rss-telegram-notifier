package entities

import (
	"time"
)

type RssItem struct {
	ID          uint      `json:"-" gorm:"primarykey"`
	CreatedAt   time.Time `json:"-"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	URL         string    `json:"URL" gorm:"not null,unique"`
	Sent        bool      `json:"-"`
	Feed        uint      `json:"-" gorm:"not null"`
}

type RssFeed struct {
	ID          uint      `json:"ID" gorm:"primarykey"`
	CreatedAt   time.Time `json:"-"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	URL         string    `json:"URL" gorm:"not null,unique"`
}

type AddFeedRequest struct {
	URL string `json:"URL"`
}
