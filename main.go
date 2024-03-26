package main

import (
	"log"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	token = "6644069287:AAGZ-CH0jzO6RRWT0Z_ayXH8kxybLz3nO0U"
)

func main() {
	bot, err := tgapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
