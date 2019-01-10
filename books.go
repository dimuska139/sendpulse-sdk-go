package sendpulse

import (
	"encoding/json"
	"errors"
	"fmt"
)

type books struct {
	Client *client
}

type Book struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	AllEmailQty      uint   `json:"all_email_qty"`
	ActiveEmailQty   uint   `json:"active_email_qty"`
	InactiveEmailQty uint   `json:"inactive_email_qty"`
	CreationDate     string `json:"creationdate"`
	Status           uint   `json:"status"`
	StatusExplain    string `json:"status_explain"`
}

type Email struct {
	Email     string            `json:"email"`
	Variables map[string]string `json:"variables"`
}

func (b *books) Get(addressBookId uint) (*Book, error) {
	body, err := b.Client.makeRequest(fmt.Sprintf("/addressbooks/%d", addressBookId), "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var respData []Book
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, errors.New(string(body))
	}

	return &respData[0], err
}

func (b *books) List(limit uint, offset uint) (*[]Book, error) {
	data := map[string]string{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
	}
	body, err := b.Client.makeRequest("/addressbooks", "GET", data, true)

	if err != nil {
		return nil, err
	}

	var respData []Book
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, errors.New(string(body))
	}

	return &respData, nil
}

func (b *books) AddEmails(addressBookId uint, notifications []Email, additionalParams map[string]string) error {
	if len(notifications) == 0 {
		return errors.New("empty emails list")
	}

	encoded, err := json.Marshal(notifications)

	if err != nil {
		return errors.New("could not to encode emails list")
	}

	data := map[string]string{
		"emails": string(encoded),
	}

	if len(additionalParams) != 0 {
		for k, v := range additionalParams {
			data[k] = v
		}
	}

	body, err := b.Client.makeRequest(fmt.Sprintf("/addressbooks/%d/emails", addressBookId), "POST", data, true)

	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return errors.New(string(body))
	}
	_, resultExists := respData["result"]
	if !resultExists || !respData["result"].(bool) {
		return errors.New(string(body))
	}

	return nil
}
