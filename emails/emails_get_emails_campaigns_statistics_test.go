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

func TestEmails_GetEmailsCampaignsStatistics(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/emailsCampaigns.json")
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/emails/campaigns", client.ApiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	email1 := "example@yourdomain.com"
	email2 := "example2@yourdomain.com"
	details, err := spClient.GetEmailsCampaignsStatistics(email1, email2)

	details1, email1exists := details[email1]
	details2, email2exists := details[email2]

	assert.True(t, email1exists)
	assert.True(t, email2exists)

	assert.Equal(t, 21, int(details1.Sent))
	assert.Equal(t, 1, int(details2.Open))
	assert.NoError(t, err)
}

func TestEmails_GetEmailsCampaignsStatistics_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/emails/campaigns", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	details, err := spClient.GetEmailsCampaignsStatistics(fake.EmailAddress(), fake.EmailAddress())
	assert.Error(t, err)
	assert.Nil(t, details)
}

func TestEmails_GetEmailsCampaignsStatistics_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/emails/campaigns", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	details, err := spClient.GetEmailsCampaignsStatistics(fake.EmailAddress(), fake.EmailAddress())
	assert.Error(t, err)
	assert.Nil(t, details)
}
