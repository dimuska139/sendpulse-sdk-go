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

func TestBooks_CampaignCost_Success(t *testing.T) {
	var bookId int = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaign := CampaignCost{
		Cur:                       fake.Word(),
		SentEmailsQty:             10,
		OverdraftAllEmailsPrice:   20,
		AddressesDeltaFromBalance: 30,
		AddressesDeltaFromTariff:  40,
		MaxEmailsPerTask:          50,
		Result:                    false,
	}
	encoded, _ := json.Marshal(campaign)

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/addressbooks/%d/cost", apiBaseUrl, bookId),
		httpmock.NewStringResponder(http.StatusOK, string(encoded)))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	camp, err := spClient.Emails.Books.CampaignCost(bookId)
	assert.NoError(t, err)

	assert.Equal(t, campaign, *camp)
}

func TestBooks_CampaignCost_BadJson(t *testing.T) {
	var bookId int = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/addressbooks/%d/cost", apiBaseUrl, bookId),
		httpmock.NewStringResponder(http.StatusOK, "{bad json"))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	camp, err := spClient.Emails.Books.CampaignCost(bookId)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Nil(t, camp)
}

func TestBooks_CampaignCost_Error(t *testing.T) {
	var bookId int = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/addressbooks/%d/cost", apiBaseUrl, bookId),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	camp, err := spClient.Emails.Books.CampaignCost(bookId)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Nil(t, camp)
}
