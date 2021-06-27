package sendpulse

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse/models"
	"net/http"
	"strconv"
)

type MailingsService struct {
	client *Client
}

func newMailingsService(cl *Client) *MailingsService {
	return &MailingsService{client: cl}
}

func (service *MailingsService) CreateMailing(data models.MailingDto) (*models.Mailing, error) {
	path := "/campaigns"
	var innerMailing struct {
		models.Mailing
		OverdraftPrice string `json:"overdraft_price"`
	}

	if data.Body != nil {
		*data.Body = b64.StdEncoding.EncodeToString([]byte(*data.Body))
	}

	if data.BodyAMP != nil {
		*data.BodyAMP = b64.StdEncoding.EncodeToString([]byte(*data.BodyAMP))
	}

	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), data, &innerMailing, true)
	if err != nil {
		return nil, err
	}

	f64, _ := strconv.ParseFloat(innerMailing.OverdraftPrice, 32)

	innerMailing.Mailing.OverdraftPrice = float32(f64)

	return &innerMailing.Mailing, err
}

func (service *MailingsService) UpdateMailing(id int, data models.MailingDto) error {
	path := fmt.Sprintf("/campaigns/%d", id)
	var respData struct {
		Result bool `json:"result"`
		Id     int  `json:"id"`
	}

	if data.Body != nil {
		*data.Body = b64.StdEncoding.EncodeToString([]byte(*data.Body))
	}

	if data.BodyAMP != nil {
		*data.BodyAMP = b64.StdEncoding.EncodeToString([]byte(*data.BodyAMP))
	}

	_, err := service.client.NewRequest(http.MethodPatch, fmt.Sprintf(path), data, &respData, true)
	return err
}

func (service *MailingsService) GetMailing(id int) (*models.Mailing, error) {
	path := fmt.Sprintf("/campaigns/%d", id)
	var respData models.Mailing
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return &respData, err
}

func (service *MailingsService) List(limit int, offset int) ([]*models.Mailing, error) {
	path := fmt.Sprintf("/campaigns?limit=%d&offset=%d", limit, offset)
	var items []*models.Mailing
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &items, true)
	return items, err
}

func (service *MailingsService) MailingsByAddressBook(addressBookID, limit, offset int) ([]*models.Task, error) {
	path := fmt.Sprintf("/addressbooks/%d/campaigns?limit=%d&offset=%d", addressBookID, limit, offset)
	var tasks []*models.Task
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &tasks, true)
	return tasks, err
}

func (service *MailingsService) CountriesStatistics(id int) (map[string]int, error) {
	path := fmt.Sprintf("/campaigns/%d/countries", id)
	response := make(map[string]int)
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &response, true)
	return response, err
}

func (service *MailingsService) ReferralsStatistics(id int) ([]*models.MailingRefStat, error) {
	path := fmt.Sprintf("/campaigns/%d/referrals", id)
	var response []*models.MailingRefStat
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &response, true)
	return response, err
}

func (service *MailingsService) Cancel(id int) error {
	path := fmt.Sprintf("/campaigns/%d", id)
	var response struct {
		Result bool `json:"result"`
	}
	_, err := service.client.NewRequest(http.MethodDelete, path, nil, &response, true)
	return err
}
