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

func TestEmails_GetAddressbookEmail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	email := fake.EmailAddress()
	responseBody, _ := ioutil.ReadFile("./testdata/addressBookEmailInfo.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/emails/%s", client.ApiBaseUrl, bookID, email),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetEmailAddressbookEmail(bookID, email)
	assert.NoError(t, err)
	assert.Equal(t, "test@test.com", info.Email)
}

func TestEmails_GetAddressbookEmail_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	email := fake.EmailAddress()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/emails/%s", client.ApiBaseUrl, bookID, email),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetEmailAddressbookEmail(bookID, email)
	assert.Error(t, err)
	assert.Nil(t, info)
}

func TestEmails_GetAddressbookEmail_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	email := fake.EmailAddress()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/emails/%s", client.ApiBaseUrl, bookID, email),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetEmailAddressbookEmail(bookID, email)
	assert.Error(t, err)
	assert.Nil(t, info)
}
