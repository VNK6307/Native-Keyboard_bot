package main

import (
	"bot/keyboard/internal/config"
	"bot/keyboard/internal/controller"
	"bot/keyboard/internal/service"
	"fmt"
	"log"
)

type Bot struct {
	controller controller.TelegramController
	service    service.Service
}

func main() {

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load configuration", err)
	}

	fmt.Printf("Bot token: %+v\n", cfg.BotToken)
	telegramController := controller.NewTelegramController(cfg)
	telegramService := service.NewTelegramService(telegramController)

	bot := &Bot{controller: *telegramController, service: *telegramService}
	bot.Start()

}

func (bot *Bot) Start() {
	offset := 0
	for {
		updates, err := bot.controller.GetUpdates(offset)
		if err != nil {
			log.Println("getUpdates error:", err)
			continue
		}
		for _, update := range updates {
			bot.service.HandleUpdate(update)
			offset = update.UpdateID + 1
		}
	}
}
