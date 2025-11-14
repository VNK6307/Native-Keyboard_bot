package controller

import (
	"bot/keyboard/internal/config"
	"bot/keyboard/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TelegramController struct {
	baseURL string
	client  *http.Client
}

func NewTelegramController(cfg *config.Config) *TelegramController {
	return &TelegramController{
		baseURL: cfg.BotURL,
		client:  &http.Client{},
	}
}

func (tc *TelegramController) makeRequest(method string, payload interface{}) ([]byte, error) {
	data, _ := json.Marshal(payload)
	resp, err := tc.client.Post(fmt.Sprintf("%s/%s", tc.baseURL, method), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("telegram error: %s", string(body))
	}
	return body, nil
}

func (tc *TelegramController) GetUpdates(offset int) ([]models.Update, error) {
	result, err := tc.makeRequest("getUpdates", map[string]interface{}{
		"offset":  offset,
		"timeout": 30 * time.Second,
	})

	if err != nil {
		return nil, err
	}

	var res struct {
		Result []models.Update `json:"result"`
	}

	err = json.Unmarshal(result, &res)
	if err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (tc *TelegramController) SendMessage(chatID uint64, text string) (int, error) {
	payload := map[string]interface{}{
		"chat_id":    chatID,
		"text":       text,
		"parse_mode": "HTML",
	}
	res, err := tc.makeRequest("sendMessage", payload)
	if err != nil {
		return 0, err
	}
	var result struct {
		Result struct {
			MessageID int `json:"message_id"`
		}
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return 0, err
	}
	return result.Result.MessageID, nil
}

func (tc *TelegramController) SendMessageWithTeamKeyboard(chatID uint64, s string, keyboard [][]models.KeyboardButton) (int, error) {
	payload := map[string]interface{}{
		"chat_id":    chatID,
		"text":       s,
		"parse_mode": "HTML",
		"keyboard":   keyboard,
	}

	fmt.Println("Вошел в отправку клавиатуры") // TODO Delete before finish
	fmt.Printf("Клавиатура - %+v\n", keyboard) // TODO Delete before finish

	body, err := tc.makeRequest("sendMessage", payload)
	if err != nil {
		return 0, fmt.Errorf("telegram error: %w", err)
	}

	fmt.Println("Запрос на отправку клавиатуры выполнен")

	var res struct {
		Result struct {
			MessageID int `json:"message_id"`
		} `json:"result"`
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return 0, fmt.Errorf("unmarshal error: %w", err)
	}

	return res.Result.MessageID, nil
}
