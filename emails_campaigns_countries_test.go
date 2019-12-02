package sendpulse

import (
	"encoding/json"
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestCampaigns_Countries_Success(t *testing.T) {
	campaignID := 1

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data := make(map[string]int)
	data["UA"] = 23
	data["RU"] = 34567

	encoded, _ := json.Marshal(data)

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/campaigns/%d/countries", apiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusOK, string(encoded)))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	stat, err := spClient.Emails.Campaigns.Countries(campaignID)
	assert.NoError(t, err)
	assert.Equal(t, data, stat)
}

func TestCampaigns_Countries_BadJson(t *testing.T) {
	campaignID := 1

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/campaigns/%d/countries", apiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusOK, `Invalid json`))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Campaigns.Countries(campaignID)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestCampaigns_Countries_Error(t *testing.T) {
	campaignID := 1

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/campaigns/%d/countries", apiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusBadRequest, ""))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Campaigns.Countries(campaignID)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}
