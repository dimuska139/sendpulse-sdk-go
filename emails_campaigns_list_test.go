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

func TestCampaigns_List_Success(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignData := []campaignInfoRaw{
		{
			ID:   1,
			Name: fake.FullName(),
			Message: messageInfoRaw{
				SenderName:  fake.FullName(),
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
		}, {
			ID:   2,
			Name: fake.FullName(),
			Message: messageInfoRaw{
				SenderName:  fake.FullName(),
				SenderEmail: fake.EmailAddress(),
				Subject:     fake.Word(),
				Body:        fake.Word(),
				Attachments: fake.Word(),
				ListID:      2128930,
			},
			Status:            3,
			AllEmailQty:       1000,
			TariffEmailQty:    1000,
			PaidEmailQty:      0,
			OverdraftPrice:    0,
			OverdraftCurrency: "RUR",
		},
	}
	encoded, _ := json.Marshal(campaignData)

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/campaigns?limit=100&offset=0", apiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, string(encoded)))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	campaignsList, err := spClient.Emails.Campaigns.List(100, 0)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(campaignsList))
}

func TestCampaigns_List_BadJson(t *testing.T) {
	path := fmt.Sprintf("/campaigns?limit=100&offset=0")
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

	_, err := spClient.Emails.Campaigns.List(100, 0)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestCampaigns_List_Error(t *testing.T) {
	path := fmt.Sprintf("/campaigns?limit=100&offset=0")
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

	_, err := spClient.Emails.Campaigns.List(100, 0)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}
