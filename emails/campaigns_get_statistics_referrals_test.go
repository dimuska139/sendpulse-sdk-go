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

func TestEmails_GetCampaignStatisticsByReferrals(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/campaignStatisticsByReferrals.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d/referrals", client.ApiBaseUrl, campaignID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	statistics, err := spClient.GetCampaignStatisticsByReferrals(campaignID)
	assert.NoError(t, err)
	assert.Equal(t, "http://first_link.com", statistics[0].Link)
}

func TestEmails_GetCampaignStatisticsByReferrals_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d/referrals", client.ApiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	statistics, err := spClient.GetCampaignStatisticsByReferrals(campaignID)
	assert.Error(t, err)
	assert.Nil(t, statistics)
}

func TestEmails_GetCampaignStatisticsByReferrals_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d/referrals", client.ApiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	statistics, err := spClient.GetCampaignStatisticsByReferrals(campaignID)
	assert.Error(t, err)
	assert.Nil(t, statistics)
}
