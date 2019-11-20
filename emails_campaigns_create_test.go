package sendpulse

import (
	"encoding/json"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
	"time"
)

func TestCampaigns_Create_IncorrectJson(t *testing.T) {
	data := CreateCampaignData{
		SenderName:   fake.Word(),
		SenderEmail:  fake.EmailAddress(),
		Subject:      fake.Word(),
		Body:         fake.Word(),
		TemplateID:   0,
		BodyAMP:      fake.Word(),
		ListID:       0,
		SegmentID:    0,
		SendTestOnly: nil,
		SendDate:     time.Now(),
		Name:         fake.Word(),
		Attachments:  nil,
		IsDraft:      true,
	}

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/campaigns"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respBody := `Incorrect json`

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	createdCampaignData, err := spClient.Emails.Campaigns.Create(data)
	assert.Error(t, err)
	spErr, isSpError := err.(*SendpulseError)
	assert.True(t, isSpError)
	assert.Nil(t, createdCampaignData)

	assert.Equal(t, http.StatusOK, spErr.HttpCode)
	assert.Equal(t, path, spErr.Url)
	assert.Equal(t, respBody, spErr.Body)
}

func TestCampaigns_Create_Error(t *testing.T) {
	data := CreateCampaignData{
		SenderName:   fake.Word(),
		SenderEmail:  fake.EmailAddress(),
		Subject:      fake.Word(),
		Body:         fake.Word(),
		TemplateID:   0,
		BodyAMP:      fake.Word(),
		ListID:       0,
		SegmentID:    0,
		SendTestOnly: nil,
		SendDate:     time.Now(),
		Name:         fake.Word(),
		Attachments:  nil,
		IsDraft:      true,
	}

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/campaigns"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respBody := `{}`

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusBadRequest, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	createdCampaignData, err := spClient.Emails.Campaigns.Create(data)
	assert.Error(t, err)
	spErr, isSpError := err.(*SendpulseError)
	assert.True(t, isSpError)
	assert.Nil(t, createdCampaignData)

	assert.Equal(t, http.StatusBadRequest, spErr.HttpCode)
	assert.Equal(t, path, spErr.Url)
	assert.Equal(t, respBody, spErr.Body)
}

func TestCampaigns_Create_Success(t *testing.T) {
	data := CreateCampaignData{
		SenderName:   fake.Word(),
		SenderEmail:  fake.EmailAddress(),
		Subject:      fake.Word(),
		Body:         fake.Word(),
		TemplateID:   0,
		BodyAMP:      fake.Word(),
		ListID:       0,
		SegmentID:    0,
		SendTestOnly: nil,
		SendDate:     time.Now(),
		Name:         fake.Word(),
		Attachments:  nil,
		IsDraft:      true,
	}

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/campaigns"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, `{"id":"27","status":0,"count":"0","tariff_email_qty":"1","paid_email_qty":"0","overdraft_price":"0","ovedraft_currency":"RUR"}`))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Campaigns.Create(data)
	assert.NoError(t, err)
}

func TestCampaigns_Create_Success_WithTestEmails(t *testing.T) {
	data := CreateCampaignData{
		SenderName:   fake.Word(),
		SenderEmail:  fake.EmailAddress(),
		Subject:      fake.Word(),
		Body:         fake.Word(),
		TemplateID:   0,
		BodyAMP:      fake.Word(),
		ListID:       0,
		SegmentID:    0,
		SendTestOnly: []string{fake.EmailAddress(), fake.EmailAddress()},
		SendDate:     time.Now(),
		Name:         fake.Word(),
		Attachments:  nil,
		IsDraft:      true,
	}

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/campaigns"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respData := createdCampaignDataRaw{
		ID:                "27",
		Status:            "0",
		Count:             "0",
		TariffEmailQty:    "1",
		PaidEmailQty:      "0",
		OverdraftPrice:    "0",
		OverdraftCurrency: "RUR",
	}
	encoded, _ := json.Marshal(respData)

	httpmock.RegisterResponder("PATCH", url,
		httpmock.NewStringResponder(http.StatusOK, string(encoded)))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Campaigns.Create(data)
	assert.NoError(t, err)
}
