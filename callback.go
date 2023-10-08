package main

// import (
// 	"fmt"

// 	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// type Callback string

// const (
// 	MoviesLink   Callback = "moviesLink"
// 	MoviesName   Callback = "moviesName"
// 	MoviesYear   Callback = "moviesYear"
// 	MoviesIMDBID Callback = "moviesIMDBID"
// 	MoviesTMDBID Callback = "moviesTMDBID"
// )

// func processCallback(update api.Update, bot *api.BotAPI) state {

// 	// Respond to the callback query, telling Telegram to show the user
// 	// a message with the data received.
// 	callback := api.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
// 	if _, err := bot.Request(callback); err != nil {
// 		panic(err)
// 	}

// 	// And finally, send a message containing the data received.

// 	fmt.Println("cmd")
// 	// msg := api.NewMessage(update.Message.Chat.ID, InitMessage)

// 	switch Callback(update.CallbackQuery.Data) {
// 	case MoviesLink:
// 		msg := api.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
// 		bot.Send(msg)
// 		return moviesYear
// 	case MoviesName:
// 		fmt.Println("start")
// 		msg := api.NewMessage(update.Message.Chat.ID, InitMessage)
// 		msg.ReplyMarkup = InitCommandKeyboard
// 		bot.Send(msg)
// 		return begin
// 	case MoviesYear:
// 		fmt.Println("start")
// 		msg := api.NewMessage(update.Message.Chat.ID, InitMessage)
// 		msg.ReplyMarkup = InitCommandKeyboard
// 		bot.Send(msg)
// 		return begin
// 	case MoviesIMDBID:
// 		fmt.Println("start")
// 		msg := api.NewMessage(update.Message.Chat.ID, InitMessage)
// 		msg.ReplyMarkup = InitCommandKeyboard
// 		bot.Send(msg)
// 		return begin
// 	case MoviesTMDBID:
// 		msg := api.NewMessage(update.Message.Chat.ID, MoviesLinkMessage)
// 		bot.Send(msg)
// 	default:
// 		msg.Text = "I don't know that command"
// 	}
// 	return begin
// }
