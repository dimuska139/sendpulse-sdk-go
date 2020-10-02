package emails

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"net/http"
)

func (api *Emails) CreateCampaign(campaignData CreateCampaignDto) (*Campaign, error) {
	path := "/campaigns"

	data := map[string]interface{}{
		"sender_name":  campaignData.SenderName,
		"sender_email": campaignData.SenderEmail,
		"subject":      campaignData.Subject,
		"body":         b64.StdEncoding.EncodeToString([]byte(campaignData.Body)),
		"template_id":  campaignData.TemplateID,
		"list_id":      campaignData.ListID,
		"attachments":  campaignData.Attachments,
	}

	if campaignData.IsDraft {
		data["type"] = "draft"
	}

	if campaignData.BodyAMP != "" {
		data["body_amp"] = b64.StdEncoding.EncodeToString([]byte(campaignData.BodyAMP))
	}

	if !campaignData.SendDate.IsZero() {
		data["send_date"] = campaignData.SendDate.Format("2006-01-02 15:04:05")
	}

	if campaignData.Name != "" {
		data["name"] = campaignData.Name
	}

	if campaignData.SegmentID != 0 {
		data["segment_id"] = campaignData.SegmentID
	}

	method := "POST"
	if len(campaignData.SendTestOnly) != 0 {
		method = "PATCH"
		encoded, _ := json.Marshal(campaignData.SendTestOnly)
		data["send_test_only"] = encoded
	}

	body, err := api.Client.NewRequest(fmt.Sprintf(path), method, data, true)
	if err != nil {
		return nil, err
	}

	var campaign Campaign
	if err := json.Unmarshal(body, &campaign); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &campaign, err
}

func (api *Emails) UpdateCampaign(campaignData UpdateCampaignDto) error {
	path := "/campaigns"

	data := map[string]interface{}{
		"id":           campaignData.ID,
		"name":         campaignData.Name,
		"sender_name":  campaignData.SenderName,
		"sender_email": campaignData.SenderEmail,
		"subject":      campaignData.Subject,
		"body":         b64.StdEncoding.EncodeToString([]byte(campaignData.Body)),
		"template_od":  campaignData.TemplateID,
		"send_date":    campaignData.SendDate.Format("2006-01-02 15:04:05"),
	}

	body, err := api.Client.NewRequest(fmt.Sprintf(path), "PATCH", data, true)
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

func (api *Emails) GetCampaign(id int) (*CampaignFullInfo, error) {
	path := fmt.Sprintf("/campaigns/%d", id)
	body, err := api.Client.NewRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var fullInfo CampaignFullInfo
	if err := json.Unmarshal(body, &fullInfo); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &fullInfo, err
}

func (api *Emails) GetCampaigns(limit int, offset int) ([]CampaignInfo, error) {
	path := "/campaigns"
	data := map[string]interface{}{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
	}
	body, err := api.Client.NewRequest(path, "GET", data, true)

	if err != nil {
		return nil, err
	}

	var campaigns []CampaignInfo
	if err := json.Unmarshal(body, &campaigns); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return campaigns, err
}

func (api *Emails) GetCampaignsByBook(addressbookId int, limit int, offset int) ([]Task, error) {
	path := fmt.Sprintf("/addressbooks/%d/campaigns", addressbookId)

	data := map[string]interface{}{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
	}
	body, err := api.Client.NewRequest(path, "GET", data, true)

	if err != nil {
		return nil, err
	}

	var campaigns []Task
	if err := json.Unmarshal(body, &campaigns); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return campaigns, err
}

func (api *Emails) GetCampaignStatisticsByCountry(campaignID int) (map[string]int, error) {
	path := fmt.Sprintf("/campaigns/%d/countries", campaignID)

	body, err := api.Client.NewRequest(path, "GET", nil, true)
	if err != nil {
		return nil, err
	}

	var respData map[string]int

	if err := json.Unmarshal(body, &respData); err != nil {
		if string(body) == "[]" {
			return respData, nil
		}

		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return respData, nil
}

func (api *Emails) GetCampaignStatisticsByReferrals(campaignID int) ([]ReferralsStatistics, error) {
	path := fmt.Sprintf("/campaigns/%d/referrals", campaignID)

	body, err := api.Client.NewRequest(path, "GET", nil, true)
	if err != nil {
		return nil, err
	}

	var respData []ReferralsStatistics
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return respData, nil
}

func (api *Emails) DeleteCampaign(campaignID int) error {
	path := fmt.Sprintf("/campaigns/%d", campaignID)

	body, err := api.Client.NewRequest(path, "DELETE", nil, true)
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
