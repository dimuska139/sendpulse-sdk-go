package emails

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"github.com/icrowley/fake"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEmails_GetCampaign(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/campaign.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d", client.ApiBaseUrl, campaignID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	campaign, err := spClient.GetCampaign(campaignID)
	assert.NoError(t, err)
	assert.NotNil(t, campaign)
}

func TestEmails_GetCampaign_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d", client.ApiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	campaign, err := spClient.GetCampaign(campaignID)
	assert.Error(t, err)
	assert.Nil(t, campaign)
}

func TestEmails_GetCampaign_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d", client.ApiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	campaign, err := spClient.GetCampaign(campaignID)
	assert.Error(t, err)
	assert.Nil(t, campaign)
}
