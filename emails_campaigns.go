package sendpulse

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type campaigns struct {
	Client *client
}

type CreateCampaignData struct {
	SenderName   string
	SenderEmail  string
	Subject      string
	Body         string
	TemplateID   string
	BodyAMP      string
	ListID       uint
	SegmentID    uint
	SendTestOnly []string
	SendDate     time.Time
	Name         string
	Attachments  map[string]string
	Type         string
}

type CreatedCampaignData struct {
	ID                string
	Status            string
	Count             string
	TariffEmailQty    string
	PaidEmailQty      string
	OverdraftPrice    string
	OverdraftCurrency string
}

type UpdateCampaignData struct {
	ID          string
	Name        string
	SenderName  string
	SenderEmail string
	Subject     string
	Body        string
	TemplateID  string
	SendDate    time.Time
}

type CampaignStatisticsCounts struct {
	Code    uint
	Count   uint
	Explain string
}

type CampaignFullInfo struct {
	ID                uint
	Name              string
	Message           map[string]interface{}
	Status            uint
	AllEmailQty       uint
	TariffEmailQty    uint
	PaidEmailQty      uint
	OverdraftPrice    uint
	OverdraftCurrency string
	Statistics        []CampaignStatisticsCounts
	SendDate          time.Time
	Permalink         string
}

// Only 4 mailing per hour
func (c *campaigns) Create(campaignData CreateCampaignData) (*CreatedCampaignData, error) {
	path := "/campaigns"

	data := map[string]interface{}{
		"sender_name":  campaignData.SenderName,
		"sender_email": campaignData.SenderEmail,
		"subject":      campaignData.Subject,
		"body":         b64.StdEncoding.EncodeToString([]byte(campaignData.Body)),
		"template_id":  campaignData.TemplateID,
		"list_id":      campaignData.ListID,
		"segment_id":   campaignData.SegmentID,
		"attachments":  campaignData.Attachments,
		"type":         campaignData.Type,
	}

	if campaignData.BodyAMP != "" {
		data["body_amp"] = b64.StdEncoding.EncodeToString([]byte(campaignData.BodyAMP))
	}

	if len(campaignData.SendTestOnly) != 0 {
		encoded, err := json.Marshal(campaignData.SendTestOnly)
		if err != nil {
			return nil, err
		}
		data["send_test_only"] = encoded
	}

	if !campaignData.SendDate.IsZero() {
		data["send_date"] = campaignData.SendDate.Format("2006-01-02 15:04:05")
	}

	if campaignData.Name != "" {
		data["name"] = campaignData.Name
	}

	method := "POST"
	if len(campaignData.SendTestOnly) != 0 {
		method = "PATCH"
	}

	body, err := c.Client.makeRequest(fmt.Sprintf(path), method, data, true)
	if err != nil {
		return nil, err
	}

	var createdCampaign CreatedCampaignData
	if err := json.Unmarshal(body, &createdCampaign); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}
	return &createdCampaign, err
}

// TODO: Unknown response format
func (c *campaigns) Update(campaignData UpdateCampaignData) error {
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

	_, err := c.Client.makeRequest(fmt.Sprintf(path), "PATCH", data, true)
	if err != nil {
		return err
	}
	return nil
}

func (c *campaigns) Get(campaignID string) (*CampaignFullInfo, error) {
	path := fmt.Sprintf("/campaigns/%s", campaignID)
	body, err := c.Client.makeRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var info CampaignFullInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	return &info, err
}

func (c *campaigns) List(limit uint, offset uint) (*[]Book, error) {

}

func (c *campaigns) ListByBook(bookID uint, limit uint, offset uint) (*[]Book, error) {

}

func (c *campaigns) Countries(campaignID uint) (map[string]int, error) {

}

func (c *campaigns) Referrals(campaignID uint) (*[]Book, error) {

}

func (c *campaigns) Delete(campaignID uint) error {
	path := fmt.Sprintf("/campaigns/%d", campaignID)
	body, err := c.Client.makeRequest(path, "DELETE", nil, true)
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
