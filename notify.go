package main

import (
	"log"
	"time"
)

func notificationRoutine() {

	for range time.NewTicker(15 * time.Second).C {
		elements, err := retrieveItemsToSend()
		if err != nil {
			continue
		}

		for _, element := range elements {

			message := createTelegramMessage(element)

			_, err = bot.Send(message)
			if err != nil {
				log.Printf("Send to Telegram failed due to: %v", err)
				continue
			}

			// setting as sent
			err = setItemAsSent(&element)
			if err != nil {
				log.Printf("Set Item As Sent failed due to: %v", err)
			}
		}
	}
}
