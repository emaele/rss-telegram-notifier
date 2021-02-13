package main

import (
	"gorm.io/gorm"
)

type rsselement struct {
	gorm.Model
	Title       string
	Description string
	URL         string `gorm:"not null,unique"`
	Sent        bool
	Feed        uint `gorm:"not null"`
}

type rssfeed struct {
	gorm.Model
	Title       string
	Description string
	URL         string `gorm:"not null,unique"`
}
