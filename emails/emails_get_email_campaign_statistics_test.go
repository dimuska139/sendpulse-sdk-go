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

func TestEmails_GetEmailCampaignStatistics(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 14563
	email := fake.EmailAddress()
	responseBody, _ := ioutil.ReadFile("./testdata/emailCampaignStatistics.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d/email/%s", client.ApiBaseUrl, campaignID, email),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetCampaignEmailStatistics(campaignID, email)
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, "In queue", info.GlobalStatusExplain)
}

func TestEmails_GetEmailCampaignStatistics_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 14563
	email := fake.EmailAddress()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d/email/%s", client.ApiBaseUrl, campaignID, email),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetCampaignEmailStatistics(campaignID, email)
	assert.Error(t, err)
	assert.Nil(t, info)
}

func TestEmails_GetEmailCampaignStatistics_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 14563
	email := fake.EmailAddress()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/campaigns/%d/email/%s", client.ApiBaseUrl, campaignID, email),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetCampaignEmailStatistics(campaignID, email)
	assert.Error(t, err)
	assert.Nil(t, info)
}
