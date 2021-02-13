package main

import (
	"fmt"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func notificationRoutine() {

	for range time.NewTicker(1 * time.Minute).C {
		var elements []rsselement
		rows := db.Debug().Where("sent = ?", false).Find(&elements).RowsAffected

		if rows == 0 {
			continue
		}

		for _, element := range elements {
			message := fmt.Sprintf("%s \n %s", element.Title, element.URL)

			_, err := bot.Send(&tb.Chat{ID: telegramChatID}, message)
			if err != nil {
				log.Println(err)
				continue
			}

			// setting as sent
			db.Model(&element).Update("sent", true)
		}
	}
}
