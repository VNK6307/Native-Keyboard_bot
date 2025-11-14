package service

import (
	"bot/keyboard/models"
	"fmt"
)

func (s *Service) HandleUpdate(update models.Update) {
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

func handleCallbackQuery(query *models.CallbackQuery) {

}
