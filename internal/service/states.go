package service

import "fmt"

const (
	StateNone              = 0
	WaitingTeamNameState   = 1
	WaitingUserChoiceState = 2
	//WaitingNextCompetitorState  = 4
)

var State = make(map[uint64]uint8)

var defaultText = "Неизвестная команда. Повторите ввод."

func (s *Service) checkState(chatID uint64, text string) {

	fmt.Printf("State = %d\n", State[chatID])

	switch State[chatID] {
	case WaitingTeamNameState:
		s.saveTeamName(chatID, text)
	case WaitingUserChoiceState:
		//	s.saveTeamMember(chatID, text)
		fmt.Println("Waiting user choice state.")
	// TODO Realize keyboard

	default:
		//s.SendDefault(chatID, defaultText)
	}

}
