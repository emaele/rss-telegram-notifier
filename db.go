package main

import (
	"errors"
	"log"

	"github.com/emaele/rss-telegram-notifier/entities"
	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
)

func retrieveFeedByID(ID string) (entities.RssFeed, error) {
	var feed entities.RssFeed

	// search for the requested feed
	res := db.Where("ID = ?", ID).Find(&feed)

	if res.RowsAffected == 0 {
		return feed, errors.New("no feed found")
	}

	if res.Error != nil {
		return feed, res.Error
	}

	return feed, nil
}

func retrieveFeeds(db *gorm.DB) ([]entities.RssFeed, error) {
	var feeds []entities.RssFeed

	res := db.Find(&feeds)
	if res.RowsAffected == 0 {
		return feeds, errors.New("no record found")
	}

	if res.Error != nil {
		return feeds, res.Error
	}

	return feeds, nil
}

func retriveItemsByFeedID(ID string) ([]entities.RssItem, error) {
	var items []entities.RssItem

	res := db.Where("Feed = ?", ID).Find(&items)
	if res.RowsAffected == 0 {
		return items, errors.New("no items found")
	}

	if res.Error != nil {
		return items, res.Error
	}

	return items, nil
}

func addItems(db *gorm.DB, feedID int64, items []*gofeed.Item, markAsSent bool) {
	log.Printf("adding %d feed elements\n", len(items))

	for _, feedelement := range items {
		element := entities.RssItem{
			Title:       feedelement.Title,
			Description: feedelement.Description,
			URL:         feedelement.Link,
			Sent:        markAsSent,
			Feed:        feedID,
		}

		if feedelement.Image != nil {
			element.ImageURL = feedelement.Image.URL
		}

		err := db.Where(entities.RssItem{URL: element.URL}).FirstOrCreate(&element).Error
		if err != nil {
			log.Printf("there was an error adding %s, %v", element.Title, err)
		}
	}
}

func retrieveFeedTitle(feedID int64) string {
	var feedTitle string

	db.Table("rss_feeds").Where("ID = ?", feedID).Pluck("Title", &feedTitle)

	return feedTitle
}

func retrieveItemsToSend() ([]entities.RssItem, error) {
	var elements []entities.RssItem
	err := db.Where("sent = ?", false).Find(&elements).Error

	return elements, err
}

func setItemAsSent(element *entities.RssItem) error {
	return db.Model(&element).Update("sent", true).Error
}

func deleteItemFromDB(element *entities.RssItem) error {
	return db.Delete(element).Error
}

func deleteFeedFromDB(element *entities.RssFeed) error {
	return db.Delete(element).Error
}

func feedExists(URL string) bool {
	// checking if the feed is duplicate
	var f entities.RssFeed
	rows := db.Where(entities.RssFeed{URL: URL}).Find(&f).RowsAffected

	return rows != 0
}

func createFeed(rssfeed *entities.RssFeed) error {
	return db.Create(rssfeed).Error
}

func retrieveFeedID(URL string) (feedID int64) {
	db.Table("rss_feeds").Where(entities.RssFeed{URL: URL}).Pluck("id", &feedID)

	return
}
