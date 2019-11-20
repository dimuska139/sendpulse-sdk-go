package sendpulse

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type books struct {
	Client *client
}

type bookRaw struct {
	ID               interface{} `json:"id"`
	Name             string      `json:"name"`
	AllEmailQty      interface{} `json:"all_email_qty"`
	ActiveEmailQty   interface{} `json:"active_email_qty"`
	InactiveEmailQty interface{} `json:"inactive_email_qty"`
	CreationDate     time.Time   `json:"creationdate"`
	Status           interface{} `json:"status"`
	StatusExplain    string      `json:"status_explain"`
}

type Book struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	AllEmailQty      int       `json:"all_email_qty"`
	ActiveEmailQty   int       `json:"active_email_qty"`
	InactiveEmailQty int       `json:"inactive_email_qty"`
	CreationDate     time.Time `json:"creationdate"`
	Status           int       `json:"status"`
	StatusExplain    string    `json:"status_explain"`
}

type Variable struct {
	Name  string
	Type  string
	Value interface{}
}

type contactRaw struct {
	Email         string
	Status        interface{}
	StatusExplain string
	Variables     []Variable
}

type Contact struct {
	Email         string
	Status        int
	StatusExplain string
	Variables     []Variable
}

type Email struct {
	Email     string                 `json:"email"`
	Variables map[string]interface{} `json:"variables"`
}

type campaignCostRaw struct {
	Cur                       string
	SentEmailsQty             interface{}
	OverdraftAllEmailsPrice   interface{}
	AddressesDeltaFromBalance interface{}
	AddressesDeltaFromTariff  interface{}
	MaxEmailsPerTask          interface{}
	Result                    bool
}

type CampaignCost struct {
	Cur                       string
	SentEmailsQty             int
	OverdraftAllEmailsPrice   int
	AddressesDeltaFromBalance int
	AddressesDeltaFromTariff  int
	MaxEmailsPerTask          int
	Result                    bool
}

func (b *books) Create(addressBookName string) (*int, error) {
	path := "/addressbooks"

	data := map[string]interface{}{
		"bookName": addressBookName,
	}
	body, err := b.Client.makeRequest(fmt.Sprintf(path), "POST", data, true)
	if err != nil {
		return nil, err
	}

	var respData map[string]int
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	createdBookId, idExists := respData["id"]
	if !idExists {
		return nil, &SendpulseError{http.StatusOK, path, string(body), "'id' not found in response"}
	}

	return &createdBookId, err
}

func (b *books) Update(addressBookId int, name string) error {
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

func (b *books) List(limit int, offset int) ([]Book, error) {
	path := "/addressbooks"
	data := map[string]interface{}{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
	}
	body, err := b.Client.makeRequest(path, "GET", data, true)

	if err != nil {
		return nil, err
	}

	var respData []bookRaw
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	var books []Book
	for _, raw := range respData {
		id, _ := strconv.Atoi(fmt.Sprint(raw.ID))
		allEmailQty, _ := strconv.Atoi(fmt.Sprint(raw.AllEmailQty))
		activeEmailQty, _ := strconv.Atoi(fmt.Sprint(raw.ActiveEmailQty))
		inactiveEmailQty, _ := strconv.Atoi(fmt.Sprint(raw.InactiveEmailQty))
		status, _ := strconv.Atoi(fmt.Sprint(raw.Status))
		books = append(books, Book{
			ID:               id,
			Name:             raw.Name,
			AllEmailQty:      allEmailQty,
			ActiveEmailQty:   activeEmailQty,
			InactiveEmailQty: inactiveEmailQty,
			CreationDate:     raw.CreationDate,
			Status:           status,
			StatusExplain:    raw.StatusExplain,
		})
	}

	return books, nil
}

func (b *books) Get(addressBookId int) (*Book, error) {
	path := fmt.Sprintf("/addressbooks/%d", addressBookId)
	body, err := b.Client.makeRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var respData []bookRaw
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	id, _ := strconv.Atoi(fmt.Sprint(respData[0].ID))
	allEmailQty, _ := strconv.Atoi(fmt.Sprint(respData[0].AllEmailQty))
	activeEmailQty, _ := strconv.Atoi(fmt.Sprint(respData[0].ActiveEmailQty))
	inactiveEmailQty, _ := strconv.Atoi(fmt.Sprint(respData[0].InactiveEmailQty))
	status, _ := strconv.Atoi(fmt.Sprint(respData[0].Status))
	book := Book{
		ID:               id,
		Name:             respData[0].Name,
		AllEmailQty:      allEmailQty,
		ActiveEmailQty:   activeEmailQty,
		InactiveEmailQty: inactiveEmailQty,
		CreationDate:     respData[0].CreationDate,
		Status:           status,
		StatusExplain:    respData[0].StatusExplain,
	}

	return &book, err
}

func (b *books) Variables(addressBookId int) ([]Variable, error) {
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

func (b *books) Emails(addressBookId int, limit int, offset int) ([]Contact, error) {
	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookId)

	data := map[string]interface{}{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
	}
	body, err := b.Client.makeRequest(path, "GET", data, true)

	if err != nil {
		return nil, err
	}

	var contactsRaw []contactRaw
	if err := json.Unmarshal(body, &contactsRaw); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	var contacts []Contact
	for _, raw := range contactsRaw {
		status, _ := strconv.Atoi(fmt.Sprint(raw.Status))
		contacts = append(contacts, Contact{
			Email:         raw.Email,
			Status:        status,
			StatusExplain: raw.StatusExplain,
			Variables:     raw.Variables,
		})
	}
	return contacts, err
}

func (b *books) EmailsTotal(addressBookId int) (int, error) {
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

	return int(total.(float64)), nil
}

/**
Known limitations:
-- Max 10 rps allowed
-- Max 255 chars per variable
-- Sendpulse calls trim function to every variable
-- Sendpulse rejects requests with html tags an \r symbols
-- Sendpulse don't remove previous user variables if user already added to address book before
*/
func (b *books) AddEmails(addressBookId int, notifications []Email, additionalParams map[string]string, senderEmail string) error {
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

func (b *books) DeleteEmails(addressBookId int, emailsList []string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookId)

	encoded, err := json.Marshal(emailsList)
	if err != nil {
		return errors.New("could not to encode emails list")
	}

	data := map[string]interface{}{
		"emails": string(encoded),
	}
	body, err := b.Client.makeRequest(path, "DELETE", data, true)
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

func (b *books) Delete(addressBookId int) error {
	path := fmt.Sprintf("/addressbooks/%d", addressBookId)
	body, err := b.Client.makeRequest(path, "DELETE", nil, true)
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

func (b *books) CampaignCost(addressBookId int) (*CampaignCost, error) {
	path := fmt.Sprintf("/addressbooks/%d/cost", addressBookId)

	body, err := b.Client.makeRequest(fmt.Sprintf(path), "GET", nil, true)
	if err != nil {
		return nil, err
	}

	var respData campaignCostRaw
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	sentEmailsQty, _ := strconv.Atoi(fmt.Sprint(respData.SentEmailsQty))
	overdraftAllEmailsPrice, _ := strconv.Atoi(fmt.Sprint(respData.OverdraftAllEmailsPrice))
	addressesDeltaFromBalance, _ := strconv.Atoi(fmt.Sprint(respData.AddressesDeltaFromBalance))
	addressesDeltaFromTariff, _ := strconv.Atoi(fmt.Sprint(respData.AddressesDeltaFromTariff))
	maxEmailsPerTask, _ := strconv.Atoi(fmt.Sprint(respData.MaxEmailsPerTask))
	cost := CampaignCost{
		Cur:                       respData.Cur,
		SentEmailsQty:             sentEmailsQty,
		OverdraftAllEmailsPrice:   overdraftAllEmailsPrice,
		AddressesDeltaFromBalance: addressesDeltaFromBalance,
		AddressesDeltaFromTariff:  addressesDeltaFromTariff,
		MaxEmailsPerTask:          maxEmailsPerTask,
		Result:                    false,
	}

	return &cost, err
}
