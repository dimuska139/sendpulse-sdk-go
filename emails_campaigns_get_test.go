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

func TestCampaigns_Get_Success(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignData := CampaignFullInfo{
		CampaignInfo: CampaignInfo{
			ID:   10113867,
			Name: fake.Word(),
			Message: MessageInfo{
				SenderName:  fake.DomainName(),
				SenderEmail: fake.EmailAddress(),
				Subject:     fake.Word(),
				Body:        fake.Word(),
				Attachments: fake.Word(),
				ListID:      2128929,
			},
			Status:            3,
			AllEmailQty:       1000,
			TariffEmailQty:    1000,
			PaidEmailQty:      0,
			OverdraftPrice:    0,
			OverdraftCurrency: "RUR",
		},
		Statistics: []CampaignStatisticsCounts{{
			Code:    1,
			Count:   1000,
			Explain: "Sent",
		}},
		Permalink: fake.DomainName(),
	}
	encoded, _ := json.Marshal(campaignData)

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/campaigns/%d", apiBaseUrl, campaignData.CampaignInfo.ID),
		httpmock.NewStringResponder(http.StatusOK, string(encoded)))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	campaign, err := spClient.Emails.Campaigns.Get(campaignData.CampaignInfo.ID)
	assert.NoError(t, err)

	assert.Equal(t, campaignData, *campaign)
}

func TestCampaigns_Get_BadJson(t *testing.T) {
	campaignID := 1

	path := fmt.Sprintf("/campaigns/%d", campaignID)
	url := apiBaseUrl + path

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusOK, `Invalid json`))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Campaigns.Get(campaignID)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestCampaigns_Get_Error(t *testing.T) {
	campaignID := 1

	path := fmt.Sprintf("/campaigns/%d", campaignID)
	url := apiBaseUrl + path

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusBadRequest, ""))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Campaigns.Get(campaignID)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}
