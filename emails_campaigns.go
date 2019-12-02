package sendpulse

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type campaigns struct {
	Client *client
}

type createdCampaignDataRaw struct {
	ID                interface{} `json:"id"`
	Status            interface{} `json:"status"`
	Count             interface{} `json:"count"`
	TariffEmailQty    interface{} `json:"tariff_email_qty"`
	PaidEmailQty      interface{} `json:"paid_email_qty"`
	OverdraftPrice    interface{} `json:"overdraft_price"`
	OverdraftCurrency string      `json:"ovedraft_currency"`
}

type CreateCampaignData struct {
	SenderName   string
	SenderEmail  string
	Subject      string
	Body         string
	TemplateID   int
	BodyAMP      string
	ListID       int
	SegmentID    int
	SendTestOnly []string
	SendDate     time.Time
	Name         string
	Attachments  map[string]string
	IsDraft      bool
}

type CreatedCampaignData struct {
	ID                int
	Status            int
	Count             int
	TariffEmailQty    int
	PaidEmailQty      int
	OverdraftPrice    int
	OverdraftCurrency string
}

type UpdateCampaignData struct {
	ID          int
	Name        string
	SenderName  string
	SenderEmail string
	Subject     string
	Body        string
	TemplateID  int
	SendDate    time.Time
}

type CampaignStatisticsCounts struct {
	Code    int
	Count   int
	Explain string
}

type MessageInfo struct {
	SenderName  string
	SenderEmail string
	Subject     string
	Body        string
	Attachments string
	ListID      int
}

type CampaignInfo struct {
	ID                int
	Name              string
	Message           MessageInfo
	Status            int
	AllEmailQty       int
	TariffEmailQty    int
	PaidEmailQty      int
	OverdraftPrice    int
	OverdraftCurrency string
}

type CampaignFullInfo struct {
	CampaignInfo
	Statistics []CampaignStatisticsCounts
	SendDate   time.Time
	Permalink  string
}

type messageInfoRaw struct {
	SenderName  string
	SenderEmail string
	Subject     string
	Body        string
	Attachments string
	ListID      interface{}
}

type campaignInfoRaw struct {
	ID                interface{}
	Name              string
	Message           messageInfoRaw
	Status            interface{}
	AllEmailQty       interface{}
	TariffEmailQty    interface{}
	PaidEmailQty      interface{}
	OverdraftPrice    interface{}
	OverdraftCurrency string
}

type ReferralsStatistics struct {
	Link  string
	Count int
}

// Limit: 4 mailing per hour
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

	method := "POST"
	if len(campaignData.SendTestOnly) != 0 {
		method = "PATCH"
		encoded, _ := json.Marshal(campaignData.SendTestOnly)
		data["send_test_only"] = encoded
	}

	body, err := c.Client.makeRequest(fmt.Sprintf(path), method, data, true)
	if err != nil {
		return nil, err
	}

	var raw createdCampaignDataRaw
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	id, _ := strconv.Atoi(fmt.Sprint(raw.ID))
	status, _ := strconv.Atoi(fmt.Sprint(raw.Status))
	count, _ := strconv.Atoi(fmt.Sprint(raw.Count))
	tariffEmailQty, _ := strconv.Atoi(fmt.Sprint(raw.TariffEmailQty))
	paidEmailQty, _ := strconv.Atoi(fmt.Sprint(raw.PaidEmailQty))
	overdraftPrice, _ := strconv.Atoi(fmt.Sprint(raw.OverdraftPrice))

	createdCampaign := CreatedCampaignData{
		ID:                id,
		Status:            status,
		Count:             count,
		TariffEmailQty:    tariffEmailQty,
		PaidEmailQty:      paidEmailQty,
		OverdraftPrice:    overdraftPrice,
		OverdraftCurrency: raw.OverdraftCurrency,
	}
	return &createdCampaign, err
}

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

	body, err := c.Client.makeRequest(fmt.Sprintf(path), "PATCH", data, true)
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

func (c *campaigns) Get(campaignID int) (*CampaignFullInfo, error) {
	path := fmt.Sprintf("/campaigns/%d", campaignID)
	body, err := c.Client.makeRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var fullInfo CampaignFullInfo
	if err := json.Unmarshal(body, &fullInfo); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	return &fullInfo, err
}

func (c *campaigns) List(limit int, offset int) ([]CampaignInfo, error) {
	path := "/campaigns"
	data := map[string]interface{}{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
	}
	body, err := c.Client.makeRequest(path, "GET", data, true)

	if err != nil {
		return nil, err
	}

	var respData []campaignInfoRaw
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	var campaignsList []CampaignInfo
	for _, raw := range respData {
		id, _ := strconv.Atoi(fmt.Sprint(raw.ID))
		status, _ := strconv.Atoi(fmt.Sprint(raw.Status))
		allEmailQty, _ := strconv.Atoi(fmt.Sprint(raw.AllEmailQty))
		tariffEmailQty, _ := strconv.Atoi(fmt.Sprint(raw.TariffEmailQty))
		paidEmailQty, _ := strconv.Atoi(fmt.Sprint(raw.PaidEmailQty))
		overdraftPrice, _ := strconv.Atoi(fmt.Sprint(raw.OverdraftPrice))
		listID, _ := strconv.Atoi(fmt.Sprint(raw.Message.ListID))

		campaignsList = append(campaignsList, CampaignInfo{
			ID:   id,
			Name: raw.Name,
			Message: MessageInfo{
				SenderName:  raw.Message.SenderName,
				SenderEmail: raw.Message.SenderEmail,
				Subject:     raw.Message.Subject,
				Body:        raw.Message.Body,
				Attachments: raw.Message.Attachments,
				ListID:      listID,
			},
			Status:            status,
			AllEmailQty:       allEmailQty,
			TariffEmailQty:    tariffEmailQty,
			PaidEmailQty:      paidEmailQty,
			OverdraftPrice:    overdraftPrice,
			OverdraftCurrency: raw.OverdraftCurrency,
		})
	}

	return campaignsList, nil
}

func (c *campaigns) Countries(campaignID int) (map[string]int, error) {
	path := fmt.Sprintf("/campaigns/%d/countries", campaignID)

	body, err := c.Client.makeRequest(path, "GET", nil, true)
	if err != nil {
		return nil, err
	}

	var respData map[string]int
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	return respData, nil
}

func (c *campaigns) Referrals(campaignID int) ([]ReferralsStatistics, error) {
	path := fmt.Sprintf("/campaigns/%d/referrals", campaignID)

	body, err := c.Client.makeRequest(path, "GET", nil, true)
	if err != nil {
		return nil, err
	}

	var respData []ReferralsStatistics
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	return respData, nil
}

func (c *campaigns) Cancel(campaignID int) error {
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
