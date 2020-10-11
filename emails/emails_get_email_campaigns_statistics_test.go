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

func TestEmails_GetEmailCampaignsStatistics(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	email := fake.EmailAddress()
	responseBody, _ := ioutil.ReadFile("./testdata/emailCampaignsStatistics.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/emails/%s/campaigns", client.ApiBaseUrl, email),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetEmailCampaignsStatistics(email)
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, 152, int(info.Statistic.Sent))
}

func TestEmails_GetEmailCampaignsStatistics_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	email := fake.EmailAddress()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/emails/%s/campaigns", client.ApiBaseUrl, email),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetEmailCampaignsStatistics(email)
	assert.Error(t, err)
	assert.Nil(t, info)
}

func TestEmails_GetEmailCampaignsStatistics_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	email := fake.EmailAddress()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/emails/%s/campaigns", client.ApiBaseUrl, email),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetEmailCampaignsStatistics(email)
	assert.Error(t, err)
	assert.Nil(t, info)
}
