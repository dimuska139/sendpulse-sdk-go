package sendpulse

import (
	"fmt"
	"net/http"
)

type AddressService struct {
	client *Client
}

func newAddressService(cl *Client) *AddressService {
	return &AddressService{client: cl}
}

type Variable struct {
	Name  string      `json:"name"`
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value"`
}

type EmailInfo struct {
	BookID    int         `json:"book_id"`
	Status    int         `json:"status"`
	Variables []*Variable `json:"variables"`
}

func (service *AddressService) GetEmailInfo(email string) ([]*EmailInfo, error) {
	path := fmt.Sprintf("/emails/%s", email)
	var response []*EmailInfo
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &response, true)
	return response, err
}

func (service *AddressService) GetEmailsInfo(emails []string) (map[string][]*EmailInfo, error) {
	path := "/emails"
	type data struct {
		Emails []string `json:"emails"`
	}

	params := data{Emails: emails}
	respData := make(map[string][]*EmailInfo)
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &respData, true)
	return respData, err
}

type EmailInfoList struct {
	ListName string       `json:"list_name"`
	ListID   int          `json:"list_id"`
	AddDate  DateTimeType `json:"add_date"`
	Source   string       `json:"source"`
}

func (service *AddressService) GetDetails(email string) ([]*EmailInfoList, error) {
	path := fmt.Sprintf("/emails/%s/details", email)
	var response []*EmailInfoList
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &response, true)
	return response, err
}

func (service *AddressService) GetStatisticsByCampaign(campaignID int, email string) (*CampaignEmailStatistics, error) {
	path := fmt.Sprintf("/campaigns/%d/email/%s", campaignID, email)
	var respData *CampaignEmailStatistics
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}

type AddressBookEmailStatistics struct {
	Email         string      `json:"email"`
	AddressBookID int         `json:"abook_id,string"`
	Status        int         `json:"status"`
	StatusExplain string      `json:"status_explain"`
	Variables     []*Variable `json:"variables"`
}

type CampaignEmailStatistics struct {
	SentDate            DateTimeType `json:"sent_date"`
	GlobalStatus        int          `json:"global_status"`
	GlobalStatusExplain string       `json:"global_status_explain"`
	DetailStatus        int          `json:"detail_status"`
	DetailStatusExplain string       `json:"detail_status_explain"`
}

func (service *AddressService) GetStatisticsByAddressBook(addressBookID int, email string) (*AddressBookEmailStatistics, error) {
	path := fmt.Sprintf("/addressbooks/%d/emails/%s", addressBookID, email)
	var respData AddressBookEmailStatistics
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

type CampaignsEmailStatistics struct {
	Statistic *struct {
		Sent int `json:"sent"`
		Open int `json:"open"`
		Link int `json:"link"`
	} `json:"statistic"`
	Addressbooks []*struct {
		Id   int    `json:"id"`
		Name string `json:"address_book_name"`
	}
	Blacklist bool
}

func (service *AddressService) GetEmailStatisticsByCampaignsAndAddressBooks(email string) (*CampaignsEmailStatistics, error) {
	path := fmt.Sprintf("/emails/%s/campaigns", email)
	var respData *CampaignsEmailStatistics
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}

type CampaignsAndAddressBooksEmailStatistics struct {
	Sent         int `json:"sent"`
	Open         int `json:"open"`
	Link         int `json:"link"`
	Addressbooks []*struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"adressbooks"`
	Blacklist bool
}

func (service *AddressService) GetEmailsStatisticsByCampaignsAndAddressBooks(emails []string) (map[string]*CampaignsAndAddressBooksEmailStatistics, error) {
	path := "/emails/campaigns"
	respData := make(map[string]*CampaignsAndAddressBooksEmailStatistics)

	type data struct {
		Emails []string `json:"emails"`
	}

	params := data{Emails: emails}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &respData, true)
	return respData, err
}

func (service *AddressService) ChangeVariables(addressBookID int, email string, variables []*Variable) error {
	path := fmt.Sprintf("/addressbooks/%d/emails/variable", addressBookID)

	type data struct {
		Email     string      `json:"email"`
		Variables []*Variable `json:"variables"`
	}

	params := data{Email: email, Variables: variables}
	var respData struct {
		Result bool `json:"result"`
	}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &respData, true)
	return err
}
