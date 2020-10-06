package emails

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"github.com/dimuska139/sendpulse-sdk-go/types"
	"net/http"
)

// CreateAddressbook makes request to create new address book.
// It returns the pointer to an ID of the new address boook and any error
func (api *Emails) CreateAddressbook(name string) (*int, error) {
	path := "/addressbooks"

	data := map[string]interface{}{
		"bookName": name,
	}
	body, err := api.Client.NewRequest(fmt.Sprintf(path), http.MethodPost, data, true)
	if err != nil {
		return nil, err
	}

	type response struct {
		ID types.Int `json:"id"`
	}

	var respData response
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &client.SendpulseError{
			HttpCode: http.StatusOK,
			Url:      path,
			Body:     string(body),
			Message:  err.Error(),
		}
	}

	bookID := int(respData.ID)

	return &bookID, err
}

func (api *Emails) UpdateAddressbook(id int, name string) error {
	path := fmt.Sprintf("/addressbooks/%d", id)

	data := map[string]interface{}{
		"name": name,
	}

	body, err := api.Client.NewRequest(fmt.Sprintf(path), http.MethodPut, data, true)
	if err != nil {
		return err
	}

	type response struct {
		Result bool `json:"result"`
	}

	var respData response
	if err := json.Unmarshal(body, &respData); err != nil {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if !respData.Result {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return nil
}

func (api *Emails) GetAddressbooks(limit int, offset int) ([]*Book, error) {
	path := "/addressbooks"
	data := map[string]interface{}{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
	}
	body, err := api.Client.NewRequest(path, http.MethodGet, data, true)

	if err != nil {
		return nil, err
	}

	var books []*Book
	if err := json.Unmarshal(body, &books); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return books, nil
}

func (api *Emails) GetAddressbook(id int) (*Book, error) {
	path := fmt.Sprintf("/addressbooks/%d", id)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var books []Book
	if err := json.Unmarshal(body, &books); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if len(books) == 0 {
		return nil, nil
	}

	return &books[0], err
}

func (api *Emails) GetAddressbookVariables(id int) ([]*Variable, error) {
	path := fmt.Sprintf("/addressbooks/%d/variables", id)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var variables []*Variable
	if err := json.Unmarshal(body, &variables); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return variables, err
}

func (api *Emails) GetAddressbookEmails(id int, limit int, offset int) ([]*EmailDetailed, error) {
	path := fmt.Sprintf("/addressbooks/%d/emails", id)

	data := map[string]interface{}{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
	}
	body, err := api.Client.NewRequest(path, http.MethodGet, data, true)

	if err != nil {
		return nil, err
	}

	var emails []*EmailDetailed
	if err := json.Unmarshal(body, &emails); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return emails, nil
}

func (api *Emails) GetAddressbookEmailsTotal(id int) (int, error) {
	path := fmt.Sprintf("/addressbooks/%d/emails/total", id)

	body, err := api.Client.NewRequest(fmt.Sprintf(path), http.MethodGet, nil, true)
	if err != nil {
		return 0, err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return 0, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	total, totalExists := respData["total"]
	if !totalExists {
		return 0, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return int(total.(float64)), nil
}

func (api *Emails) GetAddressbookEmailsByVariable(id int, variable string, value interface{}) ([]*EmailDetailed, error) {
	path := fmt.Sprintf("/addressbooks/%d/variables/%s/%v", id, variable, value)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var emails []*EmailDetailed
	if err := json.Unmarshal(body, &emails); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return emails, nil
}

/**
Known limitations:
-- Max 10 rps allowed
-- Max 255 chars per variable
-- Sendpulse calls "trim" function to every variable
-- Sendpulse rejects requests with html tags an \r symbols
-- Sendpulse don't remove previous user variables if user already added to address book before
*/
func (api *Emails) AddEmailsToAddressbookSingleOptIn(id int, emails []*Email) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", id)

	encoded, err := json.Marshal(emails)

	if err != nil {
		return errors.New("could not to encode emails list")
	}

	data := map[string]interface{}{
		"emails": string(encoded),
	}

	body, err := api.Client.NewRequest(path, http.MethodPost, data, true)

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

/**
Known limitations:
-- Max 10 rps allowed
-- Max 255 chars per variable
-- Sendpulse calls "trim" function to every variable
-- Sendpulse rejects requests with html tags an \r symbols
-- Sendpulse don't remove previous user variables if user already added to address book before
*/
func (api *Emails) AddEmailsToAddressbookDoubleOptIn(id int, emails []*Email, senderEmail string, templateId string, messageLang string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", id)

	encoded, err := json.Marshal(emails)

	if err != nil {
		return errors.New("could not to encode emails list")
	}

	data := map[string]interface{}{
		"emails":       string(encoded),
		"confirmation": "force",
		"sender_email": senderEmail,
		"message_lang": messageLang,
	}

	if templateId != "" {
		data["template_id"] = templateId
	}

	body, err := api.Client.NewRequest(path, http.MethodPost, data, true)

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

func (api *Emails) DeleteEmailsFromAddressbook(id int, emails []*string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", id)

	encoded, err := json.Marshal(emails)
	if err != nil {
		return errors.New("could not to encode emails list")
	}

	data := map[string]interface{}{
		"emails": string(encoded),
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

func (api *Emails) DeleteAddressBook(addressBookID int) error {
	path := fmt.Sprintf("/addressbooks/%d", addressBookID)

	body, err := api.Client.NewRequest(path, http.MethodDelete, nil, true)
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

func (api *Emails) GetAddressbookCampaignCost(id int) (*CampaignCost, error) {
	path := fmt.Sprintf("/addressbooks/%d/cost", id)

	body, err := api.Client.NewRequest(fmt.Sprintf(path), http.MethodGet, nil, true)
	if err != nil {
		return nil, err
	}

	var cost CampaignCost
	if err := json.Unmarshal(body, &cost); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &cost, err
}

func (api *Emails) UnsubscribeEmailsFromAddressbook(addressBookID int, emails []*string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails/unsubscribe", addressBookID)

	encoded, err := json.Marshal(emails)
	if err != nil {
		return errors.New("could not to encode emails list")
	}

	data := map[string]interface{}{
		"emails": string(encoded),
	}
	body, err := api.Client.NewRequest(path, http.MethodPost, data, true)
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
