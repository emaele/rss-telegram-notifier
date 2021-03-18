package main

import (
	"log"
	"time"

	"github.com/emaele/rss-telegram-notifier/entities"
)

func notificationRoutine() {

	for range time.NewTicker(10 * time.Second).C {
		var elements []entities.RssItem
		rows := db.Where("sent = ?", false).Find(&elements).RowsAffected

		if rows == 0 {
			continue
		}

		for _, element := range elements {

			message := createTelegramMessage(element)

			_, err := bot.Send(message)
			if err != nil {
				log.Println(err)
				continue
			}

			// setting as sent
			db.Model(&element).Update("sent", true)
		}
	}
}
