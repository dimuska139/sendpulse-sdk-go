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

func TestCampaigns_Referrals_Success(t *testing.T) {
	campaignID := 1

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data := []ReferralsStatistics{
		{
			Link:  "http://first_link.com",
			Count: 123454,
		},
		{
			Link:  "http://second_link.com",
			Count: 5463,
		},
	}

	encoded, _ := json.Marshal(data)

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/campaigns/%d/referrals", apiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusOK, string(encoded)))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	stat, err := spClient.Emails.Campaigns.Referrals(campaignID)
	assert.NoError(t, err)
	assert.Equal(t, data, stat)
}

func TestCampaigns_Referrals_BadJson(t *testing.T) {
	campaignID := 1

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/campaigns/%d/referrals", apiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusOK, `Invalid json`))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Campaigns.Referrals(campaignID)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestCampaigns_Referrals_Error(t *testing.T) {
	campaignID := 1

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/campaigns/%d/referrals", apiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusBadRequest, ""))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Campaigns.Referrals(campaignID)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}
