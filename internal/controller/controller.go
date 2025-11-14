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
