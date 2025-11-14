package service

import (
	"bot/keyboard/internal/controller"
)

type Service struct {
	tgc controller.TelegramController
}

func NewTelegramService(controller *controller.TelegramController) *Service {
	return &Service{
		tgc: *controller,
	}
}
