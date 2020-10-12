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

func TestEmails_AddEmailsToAddressbookSingleOptIn(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/singleOptIn.json")
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/addressbooks/%d/emails", client.ApiBaseUrl, bookID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	emails := []*Email{
		{
			Email:     fake.EmailAddress(),
			Variables: map[string]interface{}{},
		},
		{
			Email: fake.EmailAddress(),
			Variables: map[string]interface{}{
				"name": fake.FirstName(),
			},
		},
	}

	err := spClient.AddEmailsToAddressbookSingleOptIn(bookID, emails)
	assert.NoError(t, err)
}

func TestEmails_AddEmailsToAddressbookSingleOptIn_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/addressbooks/%d/emails", client.ApiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)
	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	emails := []*Email{
		{
			Email:     fake.EmailAddress(),
			Variables: map[string]interface{}{},
		},
	}
	err := spClient.AddEmailsToAddressbookSingleOptIn(bookID, emails)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_AddEmailsToAddressbookSingleOptIn_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/addressbooks/%d/emails", client.ApiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}
	spClient := New(http.DefaultClient, &config)
	emails := []*Email{
		{
			Email:     fake.EmailAddress(),
			Variables: map[string]interface{}{},
		},
	}
	err := spClient.AddEmailsToAddressbookSingleOptIn(bookID, emails)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_AddEmailsToAddressbookSingleOptIn_NoResult(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/addressbooks/%d/emails", client.ApiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusOK, "{}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}
	spClient := New(http.DefaultClient, &config)
	emails := []*Email{
		{
			Email:     fake.EmailAddress(),
			Variables: map[string]interface{}{},
		},
	}
	err := spClient.AddEmailsToAddressbookSingleOptIn(bookID, emails)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}
