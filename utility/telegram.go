package utility

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Thank you BomBonio!

// SetupInlineKeyboard is an utility function that makes it easier to build an appropriate set of buttons for reports
func SetupInlineKeyboard(url string) (keyboard tgbotapi.InlineKeyboardMarkup) {

	row := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Acquista", url))
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)

	//We finally append the lower row to the keyboard
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard)
	return
}
