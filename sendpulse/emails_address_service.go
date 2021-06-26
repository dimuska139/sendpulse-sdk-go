package sendpulse

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse/models"
	"net/http"
)

type AddressService struct {
	client *Client
}

func newAddressService(cl *Client) *AddressService {
	return &AddressService{client: cl}
}

func (service *AddressService) GetEmailInfo(email string) ([]*models.EmailInfo, error) {
	path := fmt.Sprintf("/emails/%s", email)
	var response []*models.EmailInfo
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &response, true)
	return response, err
}

func (service *AddressService) GetEmailsInfo(emails []string) (map[string][]*models.EmailInfo, error) {
	path := "/emails"
	type data struct {
		Emails []string `json:"emails"`
	}

	params := data{Emails: emails}
	respData := make(map[string][]*models.EmailInfo)
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &respData, true)
	return respData, err
}

func (service *AddressService) GetDetails(email string) ([]*models.EmailInfoList, error) {
	path := fmt.Sprintf("/emails/%s/details", email)
	var response []*models.EmailInfoList
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &response, true)
	return response, err
}

func (service *AddressService) GetStatisticsByCampaign(campaignID int, email string) (*models.CampaignEmailStatistics, error) {
	path := fmt.Sprintf("/campaigns/%d/email/%s", campaignID, email)
	var respData *models.CampaignEmailStatistics
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}

func (service *AddressService) GetStatisticsByAddressBook(addressBookID int, email string) (*models.AddressBookEmailStatistics, error) {
	path := fmt.Sprintf("/addressbooks/%d/emails/%s", addressBookID, email)
	var respData models.AddressBookEmailStatistics
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return &respData, err
}

func (service *AddressService) DeleteFromAllAddressBooks(email string) error {
	path := fmt.Sprintf("/emails/%s", email)
	var respData struct {
		Result bool
	}
	_, err := service.client.NewRequest(http.MethodDelete, fmt.Sprintf(path), nil, &respData, true)
	return err
}

func (service *AddressService) GetEmailStatisticsByCampaignsAndAddressBooks(email string) (*models.CampaignsEmailStatistics, error) {
	path := fmt.Sprintf("/emails/%s/campaigns", email)
	var respData *models.CampaignsEmailStatistics
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}

func (service *AddressService) GetEmailsStatisticsByCampaignsAndAddressBooks(emails []string) (map[string]*models.CampaignsAndAddressBooksEmailStatistics, error) {
	path := "/emails/campaigns"
	respData := make(map[string]*models.CampaignsAndAddressBooksEmailStatistics)

	type data struct {
		Emails []string `json:"emails"`
	}

	params := data{Emails: emails}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &respData, true)
	return respData, err
}

func (service *AddressService) ChangeVariables(addressBookID int, email string, variables []*models.Variable) error {
	path := fmt.Sprintf("/addressbooks/%d/emails/variable", addressBookID)

	type data struct {
		Email     string             `json:"email"`
		Variables []*models.Variable `json:"variables"`
	}

	params := data{Email: email, Variables: variables}
	var respData struct {
		Result bool `json:"result"`
	}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &respData, true)
	return err
}
