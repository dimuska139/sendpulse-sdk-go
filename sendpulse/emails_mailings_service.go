package sendpulse

import (
	b64 "encoding/base64"
	"fmt"
	"net/http"
	"strconv"
)

type MailingsService struct {
	client *Client
}

func newMailingsService(cl *Client) *MailingsService {
	return &MailingsService{client: cl}
}

type CampaignParams struct {
	SenderName    string            `json:"sender_name"`
	SenderEmail   string            `json:"sender_email"`
	Subject       string            `json:"subject"`
	Body          string            `json:"body,omitempty"`
	TemplateID    string            `json:"template_id,omitempty"`
	AddressBookID int               `json:"list_id,omitempty"`
	SegmentID     int               `json:"segment_id,omitempty"`
	IsTest        bool              `json:"is_test,omitempty"`
	SendDate      DateTimeType      `json:"send_date,omitempty"`
	Name          string            `json:"name,omitempty"`
	Attachments   map[string]string `json:"attachments"`
	Type          string            `json:"type,omitempty"`
	BodyAMP       string            `json:"body_amp,omitempty"`
}

type Campaign struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message struct {
		SenderName    string `json:"sender_name"`
		SenderEmail   string `json:"sender_email"`
		Subject       string `json:"subject"`
		Body          string `json:"body"`
		Attachments   string `json:"attachments"`
		AddressBookID int    `json:"list_id"`
	}
	Status            int          `json:"status"`
	AllEmailQty       int          `json:"all_email_qty"`
	TariffEmailQty    int          `json:"tariff_email_qty"`
	PaidEmailQty      int          `json:"paid_email_qty"`
	OverdraftPrice    float32      `json:"overdraft_price"`
	OverdraftCurrency string       `json:"overdraft_currency"`
	SendDate          DateTimeType `json:"send_date"`
}

func (service *MailingsService) CreateCampaign(data CampaignParams) (*Campaign, error) {
	path := "/campaigns"
	var innerMailing struct {
		Campaign
		OverdraftPrice string `json:"overdraft_price"`
	}

	if data.Body != "" {
		data.Body = b64.StdEncoding.EncodeToString([]byte(data.Body))
	}

	if data.BodyAMP != "" {
		data.BodyAMP = b64.StdEncoding.EncodeToString([]byte(data.BodyAMP))
	}

	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), data, &innerMailing, true)
	if err != nil {
		return nil, err
	}

	f64, _ := strconv.ParseFloat(innerMailing.OverdraftPrice, 32)

	innerMailing.Campaign.OverdraftPrice = float32(f64)

	return &innerMailing.Campaign, err
}

func (service *MailingsService) UpdateCampaign(id int, data CampaignParams) error {
	path := fmt.Sprintf("/campaigns/%d", id)
	var respData struct {
		Result bool `json:"result"`
		Id     int  `json:"id"`
	}

	if data.Body != "" {
		data.Body = b64.StdEncoding.EncodeToString([]byte(data.Body))
	}

	if data.BodyAMP != "" {
		data.BodyAMP = b64.StdEncoding.EncodeToString([]byte(data.BodyAMP))
	}

	_, err := service.client.NewRequest(http.MethodPatch, fmt.Sprintf(path), data, &respData, true)
	return err
}

func (service *MailingsService) GetCampaign(id int) (*Campaign, error) {
	path := fmt.Sprintf("/campaigns/%d", id)
	var respData Campaign
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return &respData, err
}

func (service *MailingsService) GetCampaigns(limit int, offset int) ([]*Campaign, error) {
	path := fmt.Sprintf("/campaigns?limit=%d&offset=%d", limit, offset)
	var items []*Campaign
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &items, true)
	return items, err
}

type Task struct {
	ID     int    `json:"task_id"`
	Name   string `json:"task_name"`
	Status int    `json:"task_status"`
}

func (service *MailingsService) GetCampaignsByAddressBook(addressBookID, limit, offset int) ([]*Task, error) {
	path := fmt.Sprintf("/addressbooks/%d/campaigns?limit=%d&offset=%d", addressBookID, limit, offset)
	var tasks []*Task
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &tasks, true)
	return tasks, err
}

func (service *MailingsService) GetCampaignCountriesStatistics(id int) (map[string]int, error) {
	path := fmt.Sprintf("/campaigns/%d/countries", id)
	response := make(map[string]int)
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &response, true)
	return response, err
}

type MailingRefStat struct {
	Link  string `json:"link"`
	Count int    `json:"count"`
}

func (service *MailingsService) GetCampaignReferralsStatistics(id int) ([]*MailingRefStat, error) {
	path := fmt.Sprintf("/campaigns/%d/referrals", id)
	var response []*MailingRefStat
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &response, true)
	return response, err
}

func (service *MailingsService) CancelCampaign(id int) error {
	path := fmt.Sprintf("/campaigns/%d", id)
	var response struct {
		Result bool `json:"result"`
	}
	_, err := service.client.NewRequest(http.MethodDelete, path, nil, &response, true)
	return err
}
