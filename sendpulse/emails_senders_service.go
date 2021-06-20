package sendpulse

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse/models"
	"net/http"
)

type SendersService struct {
	client *Client
}

func newSendersService(cl *Client) *SendersService {
	return &SendersService{client: cl}
}

func (service *SendersService) Create(name string, email string) error {
	path := "/senders"

	type paramsFormat struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	params := paramsFormat{
		Name:  name,
		Email: email,
	}

	type response struct {
		Result bool
	}

	var respData response
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &respData, true)
	return err
}

func (service *SendersService) GetActivationCode(email string) error {
	path := fmt.Sprintf("/senders/%s/code", email)

	type response struct {
		Result bool
	}
	var respData response
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return err
}

func (service *SendersService) Activate(email, code string) error {
	path := fmt.Sprintf("/senders/%s/code", email)

	type paramsFormat struct {
		Code string `json:"code"`
	}

	params := paramsFormat{
		Code: code,
	}

	type response struct {
		Result bool
	}
	var respData response
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &respData, true)
	return err
}

func (service *SendersService) List() ([]*models.Sender, error) {
	path := "/senders"

	var respData []*models.Sender
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}

func (service *SendersService) Delete(email string) error {
	path := "/senders"

	type paramsFormat struct {
		Email string `json:"email"`
	}

	params := paramsFormat{
		Email: email,
	}

	type response struct {
		Result bool
	}
	var respData response
	_, err := service.client.NewRequest(http.MethodDelete, fmt.Sprintf(path), params, &respData, true)
	return err
}
