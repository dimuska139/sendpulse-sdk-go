package sendpulse

import (
	"fmt"
	"net/http"
)

type SendersService struct {
	client *Client
}

func newSendersService(cl *Client) *SendersService {
	return &SendersService{client: cl}
}

func (service *SendersService) CreateSender(name string, email string) error {
	path := "/senders"

	type paramsFormat struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	params := paramsFormat{
		Name:  name,
		Email: email,
	}

	var response struct {
		Result bool `json:"result"`
	}

	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &response, true)
	return err
}

func (service *SendersService) GetSenderActivationCode(email string) error {
	path := fmt.Sprintf("/senders/%s/code", email)

	var response struct {
		Result bool `json:"result"`
	}
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &response, true)
	return err
}

func (service *SendersService) ActivateSender(email, code string) error {
	path := fmt.Sprintf("/senders/%s/code", email)

	type paramsFormat struct {
		Code string `json:"code"`
	}

	params := paramsFormat{
		Code: code,
	}

	var response struct {
		Result bool `json:"result"`
	}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &response, true)
	return err
}

type Sender struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func (service *SendersService) GetSenders() ([]*Sender, error) {
	path := "/senders"

	var respData []*Sender
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}

func (service *SendersService) DeleteSender(email string) error {
	path := "/senders"

	type paramsFormat struct {
		Email string `json:"email"`
	}

	params := paramsFormat{
		Email: email,
	}

	var response struct {
		Result bool `json:"result"`
	}
	_, err := service.client.NewRequest(http.MethodDelete, fmt.Sprintf(path), params, &response, true)
	return err
}
