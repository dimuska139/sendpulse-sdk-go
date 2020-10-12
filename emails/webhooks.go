package emails

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"net/http"
)

func (api *Emails) GetWebhooks() ([]*Webhook, error) {
	path := "/v2/email-service/webhook"

	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	type response struct {
		Success bool
		Data    []*Webhook
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if !resp.Success {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return resp.Data, nil
}

func (api *Emails) GetWebhook(webhookID int) (*Webhook, error) {
	path := fmt.Sprintf("/v2/email-service/webhook/%d", webhookID)

	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)
	if err != nil {
		return nil, err
	}

	type response struct {
		Success bool
		Data    *Webhook
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if !resp.Success {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return resp.Data, nil
}

func (api *Emails) CreateWebhook(url string, actions []string) ([]*Webhook, error) {
	path := "/v2/email-service/webhook/"

	encodedActions, err := json.Marshal(actions)
	if err != nil {
		return nil, errors.New("could not to encode actions list")
	}

	data := map[string]interface{}{
		"actions": string(encodedActions),
		"url":     url,
	}

	body, err := api.Client.NewRequest(path, http.MethodPost, data, true)
	if err != nil {
		return nil, err
	}

	type response struct {
		Success bool
		Data    []*Webhook
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if !resp.Success {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return resp.Data, nil
}

func (api *Emails) DeleteWebhook(webhookID int) error {
	path := fmt.Sprintf("/v2/email-service/webhook/%d", webhookID)

	body, err := api.Client.NewRequest(path, http.MethodDelete, nil, true)
	if err != nil {
		return err
	}

	type response struct {
		Success bool
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if !resp.Success {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return nil
}
