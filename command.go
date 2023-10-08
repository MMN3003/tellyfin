package main

import (
	"fmt"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type command string

const (
	startCommand        command = "start"
	helpCommand         command = "help"
	moviesCommand       command = "movies"
	moviesLinkCommand   command = "moviesLink"
	moviesYearCommand   command = "moviesYear"
	moviesIMDBIDCommand command = "moviesIMDBID"
	moviesTMDBIDCommand command = "moviesTMDBID"
	seriesCommand       command = "series"
	animesCommand       command = "animes"
	sayhiCommand        command = "sayhi"
	statusCommand       command = "status"
	openCommand         command = "open"
	closeCommand        command = "close"
)

func processCommand(update api.Update, bot *api.BotAPI) state {
	fmt.Println("cmd")
	msg := api.NewMessage(update.Message.Chat.ID, InitMessage)

	switch command(update.Message.Command()) {
	case startCommand:
		fmt.Println("start")
		msg := api.NewMessage(update.Message.Chat.ID, InitMessage)
		msg.ReplyMarkup = InitCommandKeyboard
		bot.Send(msg)
		return begin
	case moviesCommand:
		msg := api.NewMessage(update.Message.Chat.ID, MoviesLinkMessage)
		bot.Send(msg)
	case helpCommand:
		msg.Text = "I understand /sayhi and /status."
	case sayhiCommand:
		msg.Text = "Hi :)"
	case statusCommand:
		msg.Text = "I'm ok."
	// case openCommand:
	// 	msg.ReplyMarkup = numericKeyboard
	case closeCommand:
		msg.ReplyMarkup = api.NewRemoveKeyboard(true)
	default:
		msg.Text = "I don't know that command"
	}
	return begin
}
