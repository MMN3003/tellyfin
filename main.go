package main

import (
	"fmt"
	"log"

	"github.com/MMN3003/tellyfin.git/config"
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type state string

const (
	begin        state = "begin"
	moviesLink   state = "moviesLink"
	moviesYear   state = "moviesYear"
	moviesTMDBID state = "moviesTMDBID"
	moviesIMDBID state = "moviesIMDBID"
)

func main() {

	c := config.Read(nil)
	bot, err := api.NewBotAPI(c.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := api.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		fmt.Println("111")

		if update.CallbackQuery != nil {

		}
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		if update.Message.IsCommand() {
			fmt.Println("222")
			processCommand(update, bot)
		}

		// if !update.Message.IsCommand() { // ignore any non-command Messages
		// 	continue
		// }
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := api.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		// bot.Send(msg)
	}
}
