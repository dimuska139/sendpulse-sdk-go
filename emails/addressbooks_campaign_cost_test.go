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

func TestEmails_GetAddressbookCampaignCost(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/addressBookCampaignCost.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/cost", client.ApiBaseUrl, bookID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	book, err := spClient.GetAddressbookCampaignCost(bookID)
	assert.NoError(t, err)
	assert.NotNil(t, book)
}

func TestEmails_GetAddressbookCampaignCost_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/cost", client.ApiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	book, err := spClient.GetAddressbookCampaignCost(bookID)
	assert.Error(t, err)
	assert.Nil(t, book)
}

func TestEmails_GetAddressbookCampaignCost_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/cost", client.ApiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	book, err := spClient.GetAddressbookCampaignCost(bookID)
	assert.Error(t, err)
	assert.Nil(t, book)
}
