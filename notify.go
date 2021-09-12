package main

import (
	"log"
	"time"
)

func (b *Backstore) notificationRoutine() {

	for range time.NewTicker(15 * time.Second).C {
		elements, err := retrieveItemsToSend(b.db)
		if err != nil {
			continue
		}

		for _, element := range elements {

			message := createTelegramMessage(b.db, element, b.conf.TelegramChatID)

			_, err = b.bot.Send(message)
			if err != nil {
				log.Printf("Send \"%s\" to Telegram failed due to: %v", message.Text, err)
				continue
			}

			// setting as sent
			err = setItemAsSent(b.db, &element)
			if err != nil {
				log.Printf("Set Item As Sent failed due to: %v", err)
			}
		}
	}
}
