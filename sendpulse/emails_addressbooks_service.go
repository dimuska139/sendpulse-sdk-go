package sendpulse

import (
	"fmt"
	"net/http"
)

type AddressBooksService struct {
	client *Client
}

func newAddressBooksService(cl *Client) *AddressBooksService {
	return &AddressBooksService{client: cl}
}

// Update method makes request to create new address book.
// It returns the pointer to an ID of the new address boook and any error
func (service *AddressBooksService) CreateAddressBook(name string) (int, error) {
	path := "/addressbooks"

	type data struct {
		BookName string `json:"bookName"`
	}

	var response struct {
		ID int `json:"id"`
	}
	params := data{BookName: name}
	_, err := service.client.newRequest(http.MethodPost, fmt.Sprintf(path), params, &response, true)
	return response.ID, err
}

// Update method makes request to update the name of address book.
func (service *AddressBooksService) UpdateAddressBook(id int, name string) error {
	path := fmt.Sprintf("/addressbooks/%d", id)

	type data struct {
		Name string `json:"name"`
	}

	var response struct {
		Result bool
	}
	params := data{Name: name}
	_, err := service.client.newRequest(http.MethodPut, path, params, &response, true)
	return err
}

type Book struct {
	ID               int          `json:"id"`
	Name             string       `json:"name"`
	AllEmailQty      int          `json:"all_email_qty"`
	ActiveEmailQty   int          `json:"active_email_qty"`
	InactiveEmailQty int          `json:"inactive_email_qty"`
	CreationDate     DateTimeType `json:"creationdate"`
	Status           int          `json:"status"`
	StatusExplain    string       `json:"status_explain"`
}

// List method returns address books collection
func (service *AddressBooksService) GetAddressbooks(limit int, offset int) ([]*Book, error) {
	path := fmt.Sprintf("/addressbooks?limit=%d&offset=%d", limit, offset)
	var books []*Book
	_, err := service.client.newRequest(http.MethodGet, path, nil, &books, true)
	return books, err
}

// Get method returns address book information by its ID
func (service *AddressBooksService) GetAddressbook(addressBookID int) (*Book, error) {
	path := fmt.Sprintf("/addressbooks/%d", addressBookID)
	var books []*Book
	_, err := service.client.newRequest(http.MethodGet, path, nil, &books, true)
	var book *Book
	if len(books) != 0 {
		book = books[0]
	}
	return book, err
}

type VariableMeta struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Variables method returns variables of the address book
func (service *AddressBooksService) GetAddressbookVariables(addressBookID int) ([]*VariableMeta, error) {
	path := fmt.Sprintf("/addressbooks/%d/variables", addressBookID)
	var variables []*VariableMeta
	_, err := service.client.newRequest(http.MethodGet, path, nil, &variables, true)
	return variables, err
}

type Email struct {
	Email         string                 `json:"email"`
	Phone         int                    `json:"phone"`
	Status        int                    `json:"status"`
	StatusExplain string                 `json:"status_explain"`
	Variables     map[string]interface{} `json:"variables"`
}

func (service *AddressBooksService) GetAddressBookEmails(id, limit, offset int) ([]*Email, error) {
	path := fmt.Sprintf("/addressbooks/%d/emails?limit=%d&offset=%d", id, limit, offset)
	var emails []*Email
	_, err := service.client.newRequest(http.MethodGet, path, nil, &emails, true)
	return emails, err
}

func (service *AddressBooksService) CountAddressBookEmails(addressBookID int) (int, error) {
	path := fmt.Sprintf("/addressbooks/%d/emails/total", addressBookID)
	var response struct {
		Total int
	}
	_, err := service.client.newRequest(http.MethodGet, path, nil, &response, true)
	return response.Total, err
}

type EmailToAdd struct {
	Email     string                 `json:"email"`
	Variables map[string]interface{} `json:"variables"`
}

func (service *AddressBooksService) GetAddressBookEmailsByVariable(addressBookID int, variable string, value interface{}) ([]*Email, error) {
	path := fmt.Sprintf("/addressbooks/%d/variables/%s/%v", addressBookID, variable, value)
	var emails []*Email
	_, err := service.client.newRequest(http.MethodGet, path, nil, &emails, true)
	return emails, err
}

func (service *AddressBooksService) SingleOptIn(addressBookID int, emails []*EmailToAdd) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookID)
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Emails []*EmailToAdd `json:"emails"`
	}

	body := bodyFormat{Emails: emails}
	_, err := service.client.newRequest(http.MethodPost, path, body, &response, true)
	return err
}

func (service *AddressBooksService) DoubleOptIn(addressBookID int, emails []*EmailToAdd, senderEmail string, messageLang string, templateID string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookID)
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Emails       []*EmailToAdd `json:"emails"`
		Confirmation string        `json:"confirmation"`
		SenderEmail  string        `json:"sender_email"`
		MessageLang  string        `json:"message_lang"`
		TemplateID   string        `json:"template_id,omitempty"`
	}

	body := bodyFormat{
		Emails:       emails,
		Confirmation: "force",
		SenderEmail:  senderEmail,
		MessageLang:  messageLang,
	}

	if templateID != "" {
		body.TemplateID = templateID
	}
	_, err := service.client.newRequest(http.MethodPost, path, body, &response, true)
	return err
}

func (service *AddressBooksService) DeleteAddressBookEmails(addressBookID int, emails []string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookID)
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Emails []string `json:"emails"`
	}
	body := bodyFormat{Emails: emails}

	_, err := service.client.newRequest(http.MethodDelete, path, body, &response, true)
	return err
}

func (service *AddressBooksService) DeleteAddressBook(addressBookID int) error {
	path := fmt.Sprintf("/addressbooks/%d", addressBookID)
	var response struct {
		Result bool `json:"result"`
	}
	_, err := service.client.newRequest(http.MethodDelete, path, nil, &response, true)
	return err
}

type CampaignCost struct {
	Cur                       string `json:"email"`
	SentEmailsQty             int    `json:"sent_emails_qty"`
	OverdraftAllEmailsPrice   int    `json:"overdraft_all_emails_price"`
	AddressesDeltaFromBalance int    `json:"address_delta_from_balance"`
	AddressesDeltaFromTariff  int    `json:"address_delta_from_tariff"`
	MaxEmailsPerTask          int    `json:"max_emails_per_task"`
	Result                    bool   `json:"result"`
}

func (service *AddressBooksService) CountCampaignCost(addressBookID int) (*CampaignCost, error) {
	path := fmt.Sprintf("/addressbooks/%d/cost", addressBookID)
	var cost CampaignCost

	_, err := service.client.newRequest(http.MethodGet, path, nil, &cost, true)
	return &cost, err
}

func (service *AddressBooksService) UnsubscribeEmails(addressBookID int, emails []string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails/unsubscribe", addressBookID)
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Emails []string `json:"emails"`
	}
	body := bodyFormat{Emails: emails}

	_, err := service.client.newRequest(http.MethodPost, path, body, &response, true)
	return err
}

func (service *AddressBooksService) UpdateEmailVariables(addressBookID int, email string, variables []*Variable) error {
	path := fmt.Sprintf("/addressbooks/%d/emails/variable", addressBookID)
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Email     string      `json:"email"`
		Variables []*Variable `json:"variables"`
	}

	body := bodyFormat{Email: email, Variables: variables}
	_, err := service.client.newRequest(http.MethodPost, path, body, &response, true)
	return err
}
