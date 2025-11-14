package service

import (
	"bot/keyboard/internal/models"
	"fmt"
)

func (s *Service) HandleUpdate(update models.Update) {
	if update.Message != nil {
		s.handleMessage(update.Message)
	} else if update.CallbackQuery != nil {
		handleCallbackQuery(update.CallbackQuery)
	}
}

func (s *Service) handleMessage(message *models.Message) {
	// TODO Realize me!
	chatID, text, msgID := message.Chat.ID, message.Text, message.MessageID

	fmt.Printf("ChatID: %d; MessageID: %d\n=====Text - %s\n", chatID, msgID, text) // TODO Delete before finish

	switch text {
	case "/personal":
		s.startPersonalForm(chatID)
	case "/team":
		s.startTeamForm(chatID)
	case "/list":
		sendCompetitors(chatID)
	case "/send":
		mailList(chatID)
	default:
		s.checkState(chatID, text)
	}
}

func handleCallbackQuery(query *models.CallbackQuery) {

}
