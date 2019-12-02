package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
	"time"
)

func TestCampaigns_Update_BadJson(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/campaigns", apiBaseUrl)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
    	result": true
	}`

	httpmock.RegisterResponder("PATCH", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	data := UpdateCampaignData{
		ID:          1,
		Name:        fake.Word(),
		SenderName:  fake.MaleFirstName(),
		SenderEmail: fake.EmailAddress(),
		Subject:     fake.Word(),
		Body:        fake.Word(),
		TemplateID:  1,
		SendDate:    time.Now(),
	}
	err := spClient.Emails.Campaigns.Update(data)

	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestCampaigns_Update_InvalidResponse(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/campaigns", apiBaseUrl)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
    	"strange_result": true
	}`

	httpmock.RegisterResponder("PATCH", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	data := UpdateCampaignData{
		ID:          1,
		Name:        fake.Word(),
		SenderName:  fake.MaleFirstName(),
		SenderEmail: fake.EmailAddress(),
		Subject:     fake.Word(),
		Body:        fake.Word(),
		TemplateID:  1,
		SendDate:    time.Now(),
	}
	err := spClient.Emails.Campaigns.Update(data)

	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestCampaigns_Update_Error(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/campaigns", apiBaseUrl)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
    	"strange_result": true
	}`

	httpmock.RegisterResponder("PATCH", url,
		httpmock.NewStringResponder(http.StatusBadRequest, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	data := UpdateCampaignData{
		ID:          1,
		Name:        fake.Word(),
		SenderName:  fake.MaleFirstName(),
		SenderEmail: fake.EmailAddress(),
		Subject:     fake.Word(),
		Body:        fake.Word(),
		TemplateID:  1,
		SendDate:    time.Now(),
	}
	err := spClient.Emails.Campaigns.Update(data)

	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestCampaigns_Update_Success(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/campaigns", apiBaseUrl)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
    	"result": true
	}`

	httpmock.RegisterResponder("PATCH", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	data := UpdateCampaignData{
		ID:          1,
		Name:        fake.Word(),
		SenderName:  fake.MaleFirstName(),
		SenderEmail: fake.EmailAddress(),
		Subject:     fake.Word(),
		Body:        fake.Word(),
		TemplateID:  1,
		SendDate:    time.Now(),
	}
	assert.NoError(t, spClient.Emails.Campaigns.Update(data))
}
