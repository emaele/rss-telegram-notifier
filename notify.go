package main

import (
	"log"
	"time"
)

func notificationRoutine() {

	for range time.NewTicker(15 * time.Minute).C {
		elements, err := retrieveItemsToSend()
		if err != nil {
			continue
		}

		for _, element := range elements {

			message := createTelegramMessage(element)

			_, err = bot.Send(message)
			if err != nil {
				log.Println(err)
				continue
			}

			// setting as sent
			err = setItemAsSent(&element)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
