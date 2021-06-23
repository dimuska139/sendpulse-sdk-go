package sendpulse

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse/models"
	"net/http"
)

type WebhooksService struct {
	client *Client
}

func newWebhooksService(cl *Client) *WebhooksService {
	return &WebhooksService{client: cl}
}

func (service *WebhooksService) List() ([]*models.Webhook, error) {
	path := "/v2/email-service/webhook"

	var respData struct {
		Success bool              `json:"success"`
		Data    []*models.Webhook `json:"data"`
	}

	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Data, err
}

func (service *WebhooksService) Get(id int) (*models.Webhook, error) {
	path := fmt.Sprintf("/v2/email-service/webhook/%d", id)

	var respData struct {
		Success bool            `json:"success"`
		Data    *models.Webhook `json:"data"`
	}

	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Data, err
}

func (service *WebhooksService) Create(actions []string, url string) ([]*models.Webhook, error) {
	path := "/v2/email-service/webhook/"

	type data struct {
		Actions []string `json:"actions"`
		Url     string   `json:"url"`
	}

	var respData struct {
		Success bool              `json:"success"`
		Data    []*models.Webhook `json:"data"`
	}
	params := data{Actions: actions, Url: url}
	_, err := service.client.NewRequest(http.MethodPost, path, params, &respData, true)
	return respData.Data, err
}

func (service *WebhooksService) Update(id int, url string) error {
	path := fmt.Sprintf("/v2/email-service/webhook/%d", id)

	type data struct {
		Url string `json:"url"`
	}

	var respData struct {
		Success bool   `json:"success"`
		Data    []bool `json:"data"`
	}
	params := data{Url: url}
	_, err := service.client.NewRequest(http.MethodPut, path, params, &respData, true)
	return err
}

func (service *WebhooksService) Delete(id int) error {
	path := fmt.Sprintf("/v2/email-service/webhook/%d", id)

	var respData struct {
		Success bool   `json:"success"`
		Data    []bool `json:"data"`
	}
	_, err := service.client.NewRequest(http.MethodDelete, path, nil, &respData, true)
	return err
}
