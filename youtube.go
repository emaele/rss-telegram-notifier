package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/mmcdole/gofeed"
)

func parseYoutubeFeeds(feed *gofeed.Feed) {

	parsedUrl, err := url.Parse(feed.Link)
	if err != nil {
		log.Printf("Error parsing URL due to: %v", err)
		return
	}

	// I trim the prefix www. because I don't know if youtube will use always the www. CNAME
	if strings.TrimPrefix(parsedUrl.Hostname(), "www.") != "youtube.com" {
		return
	}

	for index := range feed.Items {
		addThumbnailToYoutubeItems(feed.Items[index])
	}
}

func addThumbnailToYoutubeItems(item *gofeed.Item) {

	// If I already have an image, I don't add it
	if item.Image != nil && item.Image.URL != "" {
		return
	}

	var url string

	// Checking if there are Extensions since it's there that I have the thumbnail
	if item.Extensions == nil {
		return
	}

	// Checking media -> group, then since I have multiple []Extensions I use for and map access to retrieve (finally) the thumbnail
	media, ok := item.Extensions["media"]
	if !ok {
		return
	}

	group, ok := media["group"]
	if !ok {
		return
	}

	for _, groupElem := range group {
		if groupElem.Name != "group" {
			continue
		}

		if groupElem.Children == nil {
			continue
		}

		thumbnail, ok := groupElem.Children["thumbnail"]
		if !ok {
			continue
		}

		for _, thumbElem := range thumbnail {
			if thumbElem.Name != "thumbnail" {
				continue
			}

			if thumbElem.Attrs == nil {
				continue
			}

			url = thumbElem.Attrs["url"]
		}
	}

	// When finally I have an URL, I add the Image directly into the feed

	item.Image = &gofeed.Image{
		URL:   url,
		Title: "thumbnail",
	}

}
