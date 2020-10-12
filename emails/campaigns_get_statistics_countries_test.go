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

func TestEmails_GetCampaignStatisticsByCountries(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/campaignStatisticsByCountries.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d/countries", client.ApiBaseUrl, campaignID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	statistics, err := spClient.GetCampaignStatisticsByCountries(campaignID)
	assert.NoError(t, err)
	assert.Equal(t, 23, statistics["UA"])
}

func TestEmails_GetCampaignStatisticsByCountries_Empty(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d/countries", client.ApiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusOK, "[]"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	statistics, err := spClient.GetCampaignStatisticsByCountries(campaignID)
	assert.NoError(t, err)
	assert.Nil(t, statistics)
}

func TestEmails_GetCampaignStatisticsByCountries_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d/countries", client.ApiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	statistics, err := spClient.GetCampaignStatisticsByCountries(campaignID)
	assert.Error(t, err)
	assert.Nil(t, statistics)
}

func TestEmails_GetCampaignStatisticsByCountries_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d/countries", client.ApiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	statistics, err := spClient.GetCampaignStatisticsByCountries(campaignID)
	assert.Error(t, err)
	assert.Nil(t, statistics)
}
