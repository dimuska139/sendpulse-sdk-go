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
	TemplateID   int
	BodyAMP      string
	ListID       int
	SegmentID    int
	SendTestOnly []string
	SendDate     time.Time
	Name         string
	Attachments  map[string]string
	Type         string
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

type campaignStatisticsCountsRaw struct {
	Code    interface{}
	Count   interface{}
	Explain string
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

type campaignFullInfoRaw struct {
	campaignInfoRaw
	Statistics []campaignStatisticsCountsRaw
	SendDate   time.Time
	Permalink  string
}

type Task struct {
	ID     int    `json:"task_id"`
	Name   string `json:"task_name"`
	Status int    `json:"task_status"`
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

	type createdCampaignDataRaw struct {
		ID                interface{}
		Status            interface{}
		Count             interface{}
		TariffEmailQty    interface{}
		PaidEmailQty      interface{}
		OverdraftPrice    interface{}
		OverdraftCurrency string
	}

	var raw createdCampaignDataRaw
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	id, ok := raw.ID.(int)
	if !ok {
		return nil, fmt.Errorf("can not convert %s (id) to int", raw.ID)
	}

	status, ok := raw.Status.(int)
	if !ok {
		return nil, fmt.Errorf("can not convert %s (status) to int", raw.Status)
	}

	count, ok := raw.Count.(int)
	if !ok {
		return nil, fmt.Errorf("can not convert %s (count) to int", raw.Count)
	}

	tariffEmailQty, ok := raw.TariffEmailQty.(int)
	if !ok {
		return nil, fmt.Errorf("can not convert %s (tariff_email_qty) to int", raw.TariffEmailQty)
	}

	paidEmailQty, ok := raw.PaidEmailQty.(int)
	if !ok {
		return nil, fmt.Errorf("can not convert %s (paid_email_qty) to int", raw.PaidEmailQty)
	}

	overdraftPrice, ok := raw.OverdraftPrice.(int)
	if !ok {
		return nil, fmt.Errorf("can not convert %s (overdraft_price) to int", raw.OverdraftPrice)
	}

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
		id, ok := raw.ID.(int)
		if !ok {
			return nil, fmt.Errorf("can not convert %s (id) to int", raw.ID)
		}

		status, ok := raw.Status.(int)
		if !ok {
			return nil, fmt.Errorf("can not convert %s (status) to int", raw.Status)
		}

		allEmailQty, ok := raw.AllEmailQty.(int)
		if !ok {
			return nil, fmt.Errorf("can not convert %s (all_email_qty) to int", raw.AllEmailQty)
		}

		tariffEmailQty, ok := raw.TariffEmailQty.(int)
		if !ok {
			return nil, fmt.Errorf("can not convert %s (tariff_email_qty) to int", raw.TariffEmailQty)
		}

		paidEmailQty, ok := raw.PaidEmailQty.(int)
		if !ok {
			return nil, fmt.Errorf("can not convert %s (paid_email_qty) to int", raw.PaidEmailQty)
		}

		overdraftPrice, ok := raw.OverdraftPrice.(int)
		if !ok {
			return nil, fmt.Errorf("can not convert %s (overdraft_price) to int", raw.OverdraftPrice)
		}

		listID, ok := raw.Message.ListID.(int)
		if !ok {
			return nil, fmt.Errorf("can not convert %s (list_id) to int", raw.Message.ListID)
		}

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

func (c *campaigns) ListByBook(bookID int, limit int, offset int) ([]Task, error) {
	path := fmt.Sprintf("/addressbooks/%d/campaigns", bookID)
	data := map[string]interface{}{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
	}

	body, err := c.Client.makeRequest(path, "GET", data, true)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(body, &tasks); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	return tasks, nil
}

func (c *campaigns) Countries(campaignID uint) (map[string]int, error) {
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
