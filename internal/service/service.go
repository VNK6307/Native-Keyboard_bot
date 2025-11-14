package service

import (
	"bot/keyboard/internal/controller"
	"bot/keyboard/internal/repositories"
)

type Service struct {
	tgc      controller.TelegramController
	teamRepo map[uint64]*repositories.Team
}

func NewTelegramService(controller *controller.TelegramController) *Service {
	return &Service{
		tgc:      *controller,
		teamRepo: make(map[uint64]*repositories.Team),
	}
}
