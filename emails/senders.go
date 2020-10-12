package emails

import (
	"encoding/json"
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"net/http"
)

func (api *Emails) GetSenders() ([]*Sender, error) {
	path := "/senders"

	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var senders []*Sender
	if err := json.Unmarshal(body, &senders); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return senders, nil
}

func (api *Emails) AppendSender(email string, name string) error {
	path := "/senders"

	data := map[string]interface{}{
		"email": email,
		"name":  name,
	}
	body, err := api.Client.NewRequest(fmt.Sprintf(path), http.MethodPost, data, true)
	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	result, resultExists := respData["result"]
	if !resultExists || !result.(bool) {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return nil
}

func (api *Emails) DeleteSender(email string) error {
	path := fmt.Sprintf("/senders")
	data := map[string]interface{}{
		"email": email,
	}

	body, err := api.Client.NewRequest(path, http.MethodDelete, data, true)
	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}
	result, resultExists := respData["result"]
	if !resultExists || !result.(bool) {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}
	return nil
}

func (api *Emails) ActivateSender(email string) error {
	path := fmt.Sprintf("/senders/%s/code", email)

	body, err := api.Client.NewRequest(path, http.MethodPost, nil, true)
	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}
	result, resultExists := respData["result"]
	if !resultExists || !result.(bool) {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}
	return nil
}

func (api *Emails) ActivateSenderViaEmail(email string) error {
	path := fmt.Sprintf("/senders/%s/code", email)

	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)
	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}
	result, resultExists := respData["result"]
	if !resultExists || !result.(bool) {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}
	return nil
}
