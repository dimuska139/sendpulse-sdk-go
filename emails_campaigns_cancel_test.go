package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestCampaigns_Cancel_BadJson(t *testing.T) {
	campaignID := 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/campaigns/%d", apiBaseUrl, campaignID)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
    	result": true
	}`

	httpmock.RegisterResponder("DELETE", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	err := spClient.Emails.Campaigns.Cancel(campaignID)

	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestCampaigns_Cancel_InvalidResponse(t *testing.T) {
	campaignID := 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/campaigns/%d", apiBaseUrl, campaignID)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
    	"strange_result": true
	}`

	httpmock.RegisterResponder("DELETE", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	err := spClient.Emails.Campaigns.Cancel(campaignID)

	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestCampaigns_Cancel_Error(t *testing.T) {
	campaignID := 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/campaigns/%d", apiBaseUrl, campaignID)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
    	"result": true
	}`

	httpmock.RegisterResponder("DELETE", url,
		httpmock.NewStringResponder(http.StatusInternalServerError, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	err := spClient.Emails.Campaigns.Cancel(campaignID)

	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestCampaigns_Cancel_Success(t *testing.T) {
	campaignID := 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/campaigns/%d", apiBaseUrl, campaignID)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
    	"result": true
	}`

	httpmock.RegisterResponder("DELETE", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	assert.NoError(t, spClient.Emails.Campaigns.Cancel(campaignID))
}
