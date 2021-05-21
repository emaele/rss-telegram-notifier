package entities

import (
	"time"

	"gorm.io/gorm"
)

// RssItem is a rss element
type RssItem struct {
	ID          int64          `json:"-" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Title       string         `json:"Title"`
	Description string         `json:"Description"`
	URL         string         `json:"URL" gorm:"not null,unique"`
	Sent        bool           `json:"-"`
	Feed        int64          `json:"-" gorm:"not null"`
}

// RssFeed represents a rss feed
type RssFeed struct {
	ID          int64          `json:"ID" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Title       string         `json:"Title"`
	Description string         `json:"Description"`
	URL         string         `json:"URL" gorm:"not null,unique"`
	Filter      string         `json:"Filter" gorm:"not null,unique"`
}

// AddFeedRequest is the struct for the add request
type AddFeedRequest struct {
	URL    string `json:"URL"`
	Filter string `json:"Filter"`
}
