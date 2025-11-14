package controller

import (
	"bot/keyboard/internal/models"
	"fmt"
)

func (tc *TelegramController) HandleUpdates(update models.Update) {
	if update.Message != nil {
		handleMessage(update.Message)
	} else if update.CallbackQuery != nil {
		handleCallbackQuery(update.CallbackQuery)
	}
}

func handleMessage(message *models.Message) {
	chatID, text, msgID := message.Chat.ID, message.Text, message.MessageID

	fmt.Printf("ChatID: %d; MessageID: %d\n=====Text - %s\n", chatID, msgID, text)
}

func handleCallbackQuery(callbackQuery *models.CallbackQuery) {
	fmt.Printf("Здесь будет обработана inline кнопка.\n")
}
