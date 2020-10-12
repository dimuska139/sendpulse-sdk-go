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

func TestEmails_GetAddressbookEmailsTotal(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/emailsTotal.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/emails/total", client.ApiBaseUrl, bookID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	total, err := spClient.GetAddressbookEmailsTotal(bookID)
	assert.NoError(t, err)
	assert.Equal(t, 1, total)
}

func TestEmails_GetAddressbookEmailsTotal_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/emails/total", client.ApiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	total, err := spClient.GetAddressbookEmailsTotal(bookID)
	assert.Error(t, err)
	assert.Equal(t, 0, total)
}

func TestEmails_GetAddressbookEmailsTotal_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/emails/total", client.ApiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	total, err := spClient.GetAddressbookEmailsTotal(bookID)
	assert.Error(t, err)
	assert.Equal(t, 0, total)
}

func TestEmails_GetAddressbookEmailsTotal_NoTotal(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/emails/total", client.ApiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusOK, "{}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	total, err := spClient.GetAddressbookEmailsTotal(bookID)
	assert.Error(t, err)
	assert.Equal(t, 0, total)
}
