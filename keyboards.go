package main

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var InitCommandKeyboard = api.NewInlineKeyboardMarkup(
	api.NewInlineKeyboardRow(
		api.NewInlineKeyboardButtonData("فیلم", string(moviesLink)),
		api.NewInlineKeyboardButtonData("سریال", "/"),
	),
	api.NewInlineKeyboardRow(
		api.NewInlineKeyboardButtonData("انیمه", "/"),
	),
)

// var MovieCommandKeyboard = api.NewInlineKeyboardMarkup(
// 	api.NewInlineKeyboardRow(
// 		api.NewInlineKeyboardButtonData("فیلم", string(moviesLink)),
// 		api.NewInlineKeyboardButtonData("سریال", "/"),
// 	),
// 	api.NewInlineKeyboardRow(
// 		api.NewInlineKeyboardButtonData("انیمه", "/"+string(animes)),
// 	),
// )
