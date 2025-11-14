package service

import (
	"bot/keyboard/internal/models"
	"bot/keyboard/internal/repositories"
	"fmt"
	"log"
)

func (s *Service) askTeamName(chatID uint64) {
	//TODO Realize me!
	State[chatID] = WaitingTeamNameState

	_, err := s.tgc.SendMessage(chatID, "<b>Введите название вашей команды:</b>\n")
	if err != nil {
		log.Printf("SendMessage mistake: %v", err)
		return
	}
}

func (s *Service) saveTeamName(chatID uint64, name string) {
	s.teamRepo[chatID] = repositories.NewTeamRepository()

	s.teamRepo[chatID].TeamName = name

	_, err := s.tgc.SendMessage(chatID, "<b>Выберите следующий шаг из вариантов внизу:</b>\n")
	if err != nil {
		return
	}

	fmt.Println("Дошел до отправки клавиатуры") // TODO Delete before finish

	row1 := []models.KeyboardButton{
		{Text: "Фамилия и имя участника"},
	}
	row2 := []models.KeyboardButton{
		{Text: "Проверить ввод"},
		{Text: "Отправить заявку"},
	}

	teamKeyboard := [][]models.KeyboardButton{row1, row2}
	_, err = s.tgc.SendMessageWithTeamKeyboard(chatID, "", teamKeyboard)
	if err != nil {
		return
	}

	State[chatID] = WaitingUserChoiceState
}
