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

func TestEmails_GetAddressbookCampaigns(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/addressBookCampaigns.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/campaigns?limit=0&offset=10", client.ApiBaseUrl, bookID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	campaigns, err := spClient.GetAddressbookCampaigns(bookID, 0, 10)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(campaigns))
	assert.Equal(t, 9147533, int(campaigns[0].TaskID))
}

func TestEmails_GetAddressbookCampaigns_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/campaigns?limit=0&offset=10", client.ApiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	campaigns, err := spClient.GetAddressbookCampaigns(bookID, 0, 10)
	assert.Error(t, err)
	assert.Equal(t, 0, len(campaigns))
}

func TestEmails_GetAddressbookCampaigns_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/campaigns?limit=0&offset=10", client.ApiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	campaigns, err := spClient.GetAddressbookCampaigns(bookID, 0, 10)
	assert.Error(t, err)
	assert.Equal(t, 0, len(campaigns))
}
