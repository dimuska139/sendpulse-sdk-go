package sendpulse

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type ValidatorService struct {
	client *Client
}

func newValidatorService(cl *Client) *ValidatorService {
	return &ValidatorService{client: cl}
}

func (service *ValidatorService) ValidateAddressBook(addressBookID int) error {
	path := "/verifier-service/send-list-to-verify/"
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		ID int `json:"id"`
	}
	body := bodyFormat{ID: addressBookID}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), body, &response, true)
	return err
}

// ValidationProgress is a results of email validation progress.
type ValidationProgress struct {
	Total     int `json:"total"`
	Processed int `json:"processed"`
}

func (service *ValidatorService) GetAddressBookValidationProgress(addressBookID int) (*ValidationProgress, error) {
	path := fmt.Sprintf("/verifier-service/get-progress/?id=%d", addressBookID)
	var response struct {
		Result bool                `json:"result"`
		Data   *ValidationProgress `json:"data"`
	}
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &response, true)
	return response.Data, err
}

// AddressBookValidationResult is a results of email validation progress.
type AddressBookValidationResult struct {
	ID                int          `json:"id"`
	Name              string       `json:"address_book_name"`
	AllEmailsQuantity int          `json:"all_emails_quantity"`
	Status            int          `json:"status"`
	CheckDate         DateTimeType `json:"check_date"`
	Data              *struct {
		Unverified  int `json:"0"`
		Valid       int `json:"1"`
		Unconfirmed int `json:"2"`
		Invalid     int `json:"3"`
	} `json:"data"`
	IsUpdated       int    `json:"is_updated"`
	StatusText      string `json:"status_text"`
	IsGarbageInBook bool   `json:"is_garbage_in_book"`
}

// AddressBookValidationResultDetailed is a detailed result of email validation progress.
type AddressBookValidationResultDetailed struct {
	AddressBookValidationResult
	EmailAddresses []struct {
		ID           int          `json:"id"`
		EmailAddress string       `json:"email_address"`
		CheckDate    DateTimeType `json:"check_date"`
		Status       int          `json:"status"`
		StatusText   string       `json:"status_text"`
	} `json:"email_addresses"`
	EmailAddressesTotal int `json:"email_addresses_total"`
}

func (service *ValidatorService) GetAddressBookValidationResult(addressBookID int) (*AddressBookValidationResultDetailed, error) {
	path := fmt.Sprintf("/verifier-service/check/?id=%d", addressBookID)
	var response *AddressBookValidationResultDetailed
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &response, true)
	return response, err
}

func (service *ValidatorService) GetValidatedAddressBooksList(limit, offset int) ([]*AddressBookValidationResult, error) {
	path := fmt.Sprintf("/verifier-service/check-list?start=%d&count=%d", offset, limit)
	var response struct {
		Total int                            `json:"total"`
		List  []*AddressBookValidationResult `json:"list"`
	}
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &response, true)
	return response.List, err
}

func (service *ValidatorService) ValidateEmail(email string) error {
	path := "/verifier-service/send-single-to-verify/"
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Email string `json:"email"`
	}
	body := bodyFormat{Email: email}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), body, &response, true)
	return err
}

type EmailValidationResult struct {
	Email  string `json:"email"`
	Checks struct {
		Status      int    `json:"status"`
		ValidFormat int    `json:"valid_format"`
		Disposable  int    `json:"disposable"`
		Webmail     int    `json:"webmail"`
		Gibberish   int    `json:"gibberish"`
		StatusText  string `json:"status_text"`
	} `json:"checks"`
}

func (service *ValidatorService) GetEmailValidationResult(email string) (*EmailValidationResult, error) {
	path := fmt.Sprintf("/verifier-service/get-single-result/?email=%s", email)
	var response struct {
		Result bool                   `json:"result"`
		Data   *EmailValidationResult `json:"data"`
	}
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &response, true)
	return response.Data, err
}

func (service *ValidatorService) DeleteEmailValidationResult(email string) error {
	path := "/verifier-service/delete-single-result"
	var response struct {
		Result bool `json:"result"`
	}
	type bodyFormat struct {
		Email string `json:"email"`
	}
	body := bodyFormat{Email: email}
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), body, &response, true)
	return err
}

type AddressBookReportParams struct {
	ID       int    `json:"id"`
	Format   int    `json:"format,omitempty"`
	Statuses []int  `json:"status,omitempty"`
	Lang     string `json:"lang,omitempty"`
}

func (service *ValidatorService) CreateAddressBookValidationReport(params AddressBookReportParams) error {
	path := "/verifier-service/make-report"
	var response struct {
		Result bool `json:"result"`
	}

	type bodyFormat struct {
		ID       int    `json:"id"`
		Format   int    `json:"format,omitempty"`
		Statuses string `json:"status"`
		Lang     string `json:"lang,omitempty"`
	}

	body := bodyFormat{
		ID:     params.ID,
		Format: params.Format,
		Lang:   params.Lang,
	}

	if len(params.Statuses) != 0 {
		strStatuses := make([]string, 0)
		for _, status := range params.Statuses {
			strStatuses = append(strStatuses, strconv.Itoa(status))
		}
		body.Statuses = "[" + strings.Join(strStatuses, ",") + "]"
	}

	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &response, true)
	return err
}

func (service *ValidatorService) GetAddressBookValidationReport(addressBookID int) (*AddressBookValidationResultDetailed, error) {
	path := fmt.Sprintf("/verifier-service/check-report?id=%d", addressBookID)
	var response *AddressBookValidationResultDetailed
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &response, true)
	return response, err
}
