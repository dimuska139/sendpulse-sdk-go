package emails

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"net/http"
)

func (api *Emails) GetEmailInfo(email string) ([]EmailInfo, error) {
	path := fmt.Sprintf("/emails/%s", email)
	body, err := api.Client.NewRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var info []EmailInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return info, nil
}

func (api *Emails) GetEmailsInfo(emails ...string) (map[string][]EmailInfo, error) {
	path := "/emails"

	type emailItem struct {
		Email string `json:"email"`
	}

	var emailsData []emailItem

	for _, email := range emails {
		emailsData = append(emailsData, emailItem{Email: email})
	}

	emailsJson, err := json.Marshal(emailsData)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"emails": string(emailsJson),
	}

	body, err := api.Client.NewRequest(path, "POST", data, true)

	if err != nil {
		return nil, err
	}

	results := make(map[string][]EmailInfo)
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return results, nil
}

func (api *Emails) GetEmailInfoDetails(email string) ([]EmailInfoDetails, error) {
	path := fmt.Sprintf("/emails/%s/details", email)
	body, err := api.Client.NewRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var info []EmailInfoDetails
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return info, nil
}

func (api *Emails) GetEmailCampaignStatistics(campaignID int, email string) (*EmailCampaignStatistics, error) {
	path := fmt.Sprintf("/campaigns/%d/email/%s", campaignID, email)
	body, err := api.Client.NewRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var info EmailCampaignStatistics
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &info, nil
}

func (api *Emails) GetEmailCampaignsStatistics(email string) (*EmailCampaignsStatistics, error) {
	path := fmt.Sprintf("/emails/%s/campaigns", email)
	body, err := api.Client.NewRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var info EmailCampaignsStatistics
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &info, nil
}

func (api *Emails) GetEmailsCampaignsStatistics(emailsList []string) (map[string]*EmailCampaignsStatisticsDetails, error) {
	path := "/emails/campaigns"

	encoded, err := json.Marshal(emailsList)
	if err != nil {
		return nil, errors.New("could not to encode emails list")
	}
	data := map[string]interface{}{
		"emails": string(encoded),
	}

	body, err := api.Client.NewRequest(path, "POST", data, true)
	if err != nil {
		return nil, err
	}

	results := make(map[string]*EmailCampaignsStatisticsDetails)
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return results, nil
}

func (api *Emails) GetEmailAddressbookStatistics(addressBookID int, email string) (*EmailAddressbookStatistics, error) {
	path := fmt.Sprintf("/addressbooks/%d/emails/%s", addressBookID, email)
	body, err := api.Client.NewRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var info EmailAddressbookStatistics
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &info, nil
}

func (api *Emails) UpdateVariable(addressBookID int, email string, variables map[string]interface{}) error {
	path := fmt.Sprintf("/addressbooks/%d/emails/variable", addressBookID)

	var variablesData []Variable
	for key, value := range variables {
		variablesData = append(variablesData, Variable{
			Name:  &key,
			Value: &value,
		})
	}

	variablesJson, err := json.Marshal(variablesData)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"email":     email,
		"variables": string(variablesJson),
	}

	body, err := api.Client.NewRequest(fmt.Sprintf(path), "POST", data, true)
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
