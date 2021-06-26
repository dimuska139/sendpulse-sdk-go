package sendpulse

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse/models"
	"net/http"
)

type TemplatesService struct {
	client *Client
}

func newTemplatesService(cl *Client) *TemplatesService {
	return &TemplatesService{client: cl}
}

func (service *TemplatesService) Create(name string, body string, lang string) (int, error) {
	path := "/template"

	type paramsFormat struct {
		Name string `json:"name,omitempty"`
		Lang string `json:"lang"`
		Body string `json:"body"`
	}

	params := paramsFormat{
		Body: b64.StdEncoding.EncodeToString([]byte(body)),
		Lang: lang,
	}

	if name != "" {
		params.Name = name
	}

	var response struct {
		Result bool `json:"result"`
		RealID int  `json:"real_id"`
	}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &response, true)
	return response.RealID, err
}

func (service *TemplatesService) Update(templateID int, body string, lang string) error {
	path := fmt.Sprintf("/template/edit/%d", templateID)

	type paramsFormat struct {
		Lang string `json:"lang"`
		Body string `json:"body"`
	}

	params := paramsFormat{
		Body: b64.StdEncoding.EncodeToString([]byte(body)),
		Lang: lang,
	}

	var response struct {
		Result bool `json:"result"`
	}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &response, true)
	return err
}

func (service *TemplatesService) Get(templateID int) (*models.Template, error) {
	path := fmt.Sprintf("/template/%d", templateID)
	var respData models.Template
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return &respData, err
}

func (service *TemplatesService) List(limit, offset int, owner string) ([]*models.Template, error) {
	path := fmt.Sprintf("/templates?limit=%d&offset=%d", limit, offset)
	if owner != "" {
		path += fmt.Sprintf("&owner=%s", owner)
	}

	var respData []*models.Template
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}
