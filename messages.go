package main

import (
	"fmt"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	InitMessage         = "جهت استفاده یکی رو انتخاب کن"
	MoviesLinkMessage   = "لینک فیلم را وارد کنید"
	MoviesNameMessage   = "نام فیلم را وارد کنید"
	MoviesYearMessage   = "سال ساخت فیلم را وارد کنید"
	MoviesIMDBIDMessage = "شناسه IMDB فیلم را وارد کنید"
	MoviesTMDBIDMessage = "شناسه TMDB فیلم را وارد کنید"
)

func processMessage(update api.Update, state state, bot *api.BotAPI) state {
	msg := api.NewMessage(update.Message.Chat.ID, InitMessage)

	switch state(state) {
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
