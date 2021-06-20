package sendpulse

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse/models"
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
func (service *AddressBooksService) Create(name string) (int, error) {
	path := "/addressbooks"

	type data struct {
		BookName string `json:"bookName"`
	}

	type response struct {
		ID int `json:"id"`
	}
	var respData response
	params := data{BookName: name}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &respData, true)
	return respData.ID, err
}

// Update method makes request to update the name of address book.
func (service *AddressBooksService) Update(id int, name string) error {
	path := fmt.Sprintf("/addressbooks/%d", id)

	type data struct {
		Name string `json:"name"`
	}

	var response struct {
		Result bool
	}
	params := data{Name: name}
	_, err := service.client.NewRequest(http.MethodPut, path, params, &response, true)
	return err
}

// List method returns address books collection
func (service *AddressBooksService) List(limit int, offset int) ([]*models.Book, error) {
	path := fmt.Sprintf("/addressbooks?limit=%d&offset=%d", limit, offset)
	var books []*models.Book
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &books, true)
	return books, err
}

// Get method returns address book information by its ID
func (service *AddressBooksService) Get(addressBookID int) (*models.Book, error) {
	path := fmt.Sprintf("/addressbooks/%d", addressBookID)
	var books []*models.Book
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &books, true)
	var book *models.Book
	if len(books) != 0 {
		book = books[0]
	}
	return book, err
}

// Variables method returns variables of the address book
func (service *AddressBooksService) Variables(addressBookID int) ([]*models.VariableMeta, error) {
	path := fmt.Sprintf("/addressbooks/%d/variables", addressBookID)
	var variables []*models.VariableMeta
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &variables, true)
	return variables, err
}

func (service *AddressBooksService) Emails(id, limit, offset int) ([]*models.Email, error) {
	path := fmt.Sprintf("/addressbooks/%d/emails?limit=%d&offset=%d", id, limit, offset)
	var emails []*models.Email
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &emails, true)
	return emails, err
}

func (service *AddressBooksService) EmailsTotal(addressBookID int) (int, error) {
	path := fmt.Sprintf("/addressbooks/%d/emails/total", addressBookID)
	var response struct {
		Total int
	}
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &response, true)
	return response.Total, err
}

func (service *AddressBooksService) EmailsByVariable(addressBookID int, variable string, value interface{}) ([]*models.Email, error) {
	path := fmt.Sprintf("/addressbooks/%d/variables/%s/%v", addressBookID, variable, value)
	var emails []*models.Email
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &emails, true)
	return emails, err
}

func (service *AddressBooksService) SingleOptIn(addressBookID int, emails []*models.EmailToAdd) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookID)
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Emails []*models.EmailToAdd `json:"emails"`
	}

	body := bodyFormat{Emails: emails}
	_, err := service.client.NewRequest(http.MethodPost, path, body, &response, true)
	return err
}

func (service *AddressBooksService) DoubleOptIn(addressBookID int, emails []*models.EmailToAdd, senderEmail string, messageLang string, templateID string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookID)
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Emails       []*models.EmailToAdd `json:"emails"`
		Confirmation string               `json:"confirmation"`
		SenderEmail  string               `json:"sender_email"`
		MessageLang  string               `json:"message_lang"`
		TemplateID   string               `json:"template_id,omitempty"`
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
	_, err := service.client.NewRequest(http.MethodPost, path, body, &response, true)
	return err
}

func (service *AddressBooksService) DeleteEmails(addressBookID int, emails []string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookID)
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Emails []string `json:"emails"`
	}
	body := bodyFormat{Emails: emails}

	_, err := service.client.NewRequest(http.MethodDelete, path, body, &response, true)
	return err
}

func (service *AddressBooksService) Delete(addressBookID int) error {
	path := fmt.Sprintf("/addressbooks/%d", addressBookID)
	var response struct {
		Result bool `json:"result"`
	}
	_, err := service.client.NewRequest(http.MethodDelete, path, nil, &response, true)
	return err
}

func (service *AddressBooksService) CampaignCost(addressBookID int) (*models.CampaignCost, error) {
	path := fmt.Sprintf("/addressbooks/%d/cost", addressBookID)
	var cost models.CampaignCost

	_, err := service.client.NewRequest(http.MethodGet, path, nil, &cost, true)
	return &cost, err
}

func (service *AddressBooksService) Unsubscribe(addressBookID int, emails []string) error {
	path := fmt.Sprintf("/addressbooks/%d/emails/unsubscribe", addressBookID)
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Emails []string `json:"emails"`
	}
	body := bodyFormat{Emails: emails}

	_, err := service.client.NewRequest(http.MethodPost, path, body, &response, true)
	return err
}

func (service *AddressBooksService) CampaignsList(addressBookID, limit, offset int) ([]*models.Task, error) {
	path := fmt.Sprintf("/addressbooks/%d/campaigns?limit=%d&offset=%d", addressBookID, limit, offset)
	var tasks []*models.Task
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &tasks, true)
	return tasks, err
}

func (service *AddressBooksService) UpdateEmailVariables(addressBookID int, email string, variables []*models.Variable) error {
	path := fmt.Sprintf("/addressbooks/%d/emails/variable", addressBookID)
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Email     string             `json:"email"`
		Variables []*models.Variable `json:"variables"`
	}

	body := bodyFormat{Email: email, Variables: variables}
	_, err := service.client.NewRequest(http.MethodPost, path, body, &response, true)
	return err
}
