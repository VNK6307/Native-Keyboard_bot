package service

import (
	"log"
)

func (s *Service) startPersonalForm(chatID uint64) {
	//TODO Realize me!

	_, err := s.tgc.SendMessage(chatID, "Здесь будет заполнение личной заявки.")
	if err != nil {
		log.Printf("SendMessage mistake: %v", err)
		return
	}
}

func (s *Service) startTeamForm(chatID uint64) {
	//TODO Realize me!

	s.askTeamName(chatID)
}

func mailList(id uint64) {

}

func sendCompetitors(id uint64) {

}
