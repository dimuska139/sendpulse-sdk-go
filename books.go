package sendpulse

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
	Email     string                 `json:"email"`
	Variables map[string]interface{} `json:"variables"`
}

func (b *books) Create(addressBookName string) (*uint, error) {
	path := "/addressbooks"

	if len(addressBookName) == 0 {
		return nil, errors.New("could not to create address book with empty name")
	}
	data := map[string]interface{}{
		"bookName": addressBookName,
	}
	body, err := b.Client.makeRequest(fmt.Sprintf(path), "POST", data, true)
	if err != nil {
		return nil, err
	}

	var respData map[string]uint
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	createdBookId, idExists := respData["id"]
	if !idExists {
		return nil, &SendpulseError{http.StatusOK, path, string(body), "'id' not found in response"}
	}

	return &createdBookId, err
}

func (b *books) Get(addressBookId uint) (*Book, error) {
	path := fmt.Sprintf("/addressbooks/%d", addressBookId)
	body, err := b.Client.makeRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var respData []Book
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	return &respData[0], err
}

func (b *books) List(limit uint, offset uint) (*[]Book, error) {
	path := "/addressbooks"
	data := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}
	body, err := b.Client.makeRequest(path, "GET", data, true)

	if err != nil {
		return nil, err
	}

	var respData []Book
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	return &respData, nil
}

func (b *books) AddEmails(addressBookId uint, notifications []Email, additionalParams map[string]string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookId)

	if len(notifications) == 0 {
		return errors.New("empty emails list")
	}

	encoded, err := json.Marshal(notifications)

	if err != nil {
		return errors.New("could not to encode emails list")
	}

	data := map[string]interface{}{
		"emails": string(encoded),
	}

	if len(additionalParams) != 0 {
		for k, v := range additionalParams {
			data[k] = v
		}
	}

	body, err := b.Client.makeRequest(path, "POST", data, true)

	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}
	result, resultExists := respData["result"]
	if !resultExists {
		return &SendpulseError{http.StatusOK, path, string(body), "'result' not found in response"}
	}

	if !result.(bool) {
		return &SendpulseError{http.StatusOK, path, string(body), "'result' is false"}
	}

	return nil
}
