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

type Variable struct {
	Name  string
	Type  string
	Value string
}

type Contact struct {
	Email         string
	Status        int
	StatusExplain string
	Variables     map[string]interface{}
}

type Email struct {
	Email     string                 `json:"email"`
	Variables map[string]interface{} `json:"variables"`
}

type CompaignCost struct {
	Cur                       string
	SentEmailsQty             int
	OverdraftAllEmailsPrice   int
	AddressesDeltaFromBalance int
	AddressesDeltaFromTariff  int
	MaxEmailsPerTask          int
	Result                    bool
}

func (b *books) Create(addressBookName string) (*uint, error) {
	path := "/addressbooks"

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

func (b *books) Update(addressBookId uint, name string) error {
	path := fmt.Sprintf("/addressbooks/%d", addressBookId)

	data := map[string]interface{}{
		"name": name,
	}

	body, err := b.Client.makeRequest(fmt.Sprintf(path), "PUT", data, true)
	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	result, resultExists := respData["result"]
	if !resultExists || !result.(bool) {
		return &SendpulseError{http.StatusOK, path, string(body), "invalid response"}
	}

	return nil
}

func (b *books) List(limit uint, offset uint) (*[]Book, error) {
	path := "/addressbooks"
	data := map[string]interface{}{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
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

func (b *books) Variables(addressBookId uint) ([]Variable, error) {
	path := fmt.Sprintf("/addressbooks/%d/variables", addressBookId)
	body, err := b.Client.makeRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var variables []Variable
	if err := json.Unmarshal(body, &variables); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	return variables, err
}

func (b *books) Emails(addressBookId uint, limit uint, offset uint) ([]Contact, error) {
	path := fmt.Sprintf("/addressbooks/%d/emails?limit=%d&offset=%d", addressBookId, limit, offset)

	body, err := b.Client.makeRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var contacts []Contact
	if err := json.Unmarshal(body, &contacts); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	return contacts, err
}

func (b *books) EmailsTotal(addressBookId uint) (uint, error) {
	path := fmt.Sprintf("/addressbooks/%d/emails/total", addressBookId)

	body, err := b.Client.makeRequest(fmt.Sprintf(path), "GET", nil, true)
	if err != nil {
		return 0, err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return 0, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	total, totalExists := respData["total"]
	if !totalExists {
		return 0, &SendpulseError{http.StatusOK, path, string(body), "'total' not found in response"}
	}

	return uint(total.(float64)), nil
}

func (b *books) AddEmails(addressBookId uint, notifications []Email, additionalParams map[string]string, senderEmail string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookId)

	encoded, err := json.Marshal(notifications)

	if err != nil {
		return errors.New("could not to encode emails list")
	}

	data := map[string]interface{}{
		"emails": string(encoded),
	}

	if senderEmail != "" { // double-opt-in method
		data["confirmation"] = "force"
		data["sender_email"] = senderEmail
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
	if !resultExists || !result.(bool) {
		return &SendpulseError{http.StatusOK, path, string(body), "invalid response"}
	}
	return nil
}

func (b *books) DeleteEmails(addressBookId uint, emailsList []string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookId)

	encoded, err := json.Marshal(emailsList)
	if err != nil {
		return errors.New("could not to encode emails list")
	}

	data := map[string]interface{}{
		"emails": string(encoded),
	}
	_, err = b.Client.makeRequest(path, "DELETE", data, true)
	return err
}

func (b *books) Delete(addressBookId uint) error {
	path := fmt.Sprintf("/addressbooks/%d", addressBookId)
	_, err := b.Client.makeRequest(path, "DELETE", nil, true)
	return err
}

func (b *books) CampaignCost(addressBookId uint) (*CompaignCost, error) {
	path := fmt.Sprintf("/addressbooks/%d/cost", addressBookId)

	body, err := b.Client.makeRequest(fmt.Sprintf(path), "GET", nil, true)
	if err != nil {
		return nil, err
	}

	var respData CompaignCost
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	return &respData, err
}
