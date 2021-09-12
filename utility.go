package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/emaele/rss-telegram-notifier/entities"
	"gorm.io/gorm"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func writeHTTPResponse(statusCode int, body string, writer http.ResponseWriter) {

	writer.WriteHeader(statusCode)
	_, err := writer.Write([]byte(body))
	if err != nil {
		log.Println(err)
	}
}

func createTelegramKeyboard(URL string) tg.InlineKeyboardMarkup {
	var keyboard tg.InlineKeyboardMarkup
	row := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonURL("🔗 Link", URL))
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)

	//We finally append the lower row to the keyboard
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard)

	return keyboard
}

func createTelegramMessage(db *gorm.DB, element entities.RssItem, telegramChatID int64) tg.MessageConfig {

	feedTitle := retrieveFeedTitle(db, element.Feed)

	if feedTitle == "" {
		feedTitle = "New Feed!"
	}

	tgMarkdownReservedChars := []string{".", "-", "(", ")", "#", "!", "|", "[", "]", "_", "*", "`", "~"}

	var text string

	// Pre-parsing our elements for markdown Reserved Chars
	for _, char := range tgMarkdownReservedChars {
		element.Title = strings.ReplaceAll(element.Title, char, fmt.Sprintf(`\%s`, char))
		element.ImageURL = strings.ReplaceAll(element.ImageURL, char, fmt.Sprintf(`\%s`, char))
		feedTitle = strings.ReplaceAll(feedTitle, char, fmt.Sprintf(`\%s`, char))
	}

	// Creating the message with pre-parsed items
	if element.ImageURL != "" {
		text = fmt.Sprintf("📣 *%s*\n\n[➡️](%s) %s", feedTitle, element.ImageURL, element.Title)
	} else {
		text = fmt.Sprintf("📣 *%s*\n\n➡️ %s", feedTitle, element.Title)
	}

	message := tg.NewMessage(telegramChatID, text)
	message.ParseMode = "MarkdownV2"
	message.ReplyMarkup = createTelegramKeyboard(element.URL)

	return message
}
