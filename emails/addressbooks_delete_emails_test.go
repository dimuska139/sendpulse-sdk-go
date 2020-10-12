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

func TestEmails_DeleteEmailsFromAddressbook(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookId := 1
	responseBody, _ := ioutil.ReadFile("./testdata/deleteEmail.json")
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/addressbooks/%d/emails", client.ApiBaseUrl, bookId),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	email := fake.EmailAddress()
	email1 := fake.EmailAddress()
	assert.NoError(t, spClient.DeleteEmailsFromAddressbook(bookId, []*string{&email, &email1}))
}

func TestEmails_DeleteEmailsFromAddressbook_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookId := 1
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/addressbooks/%d/emails", client.ApiBaseUrl, bookId),
		httpmock.NewStringResponder(http.StatusBadRequest, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	email := fake.EmailAddress()
	email1 := fake.EmailAddress()
	err := spClient.DeleteEmailsFromAddressbook(bookId, []*string{&email, &email1})
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_DeleteEmailsFromAddressbook_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookId := 1
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/addressbooks/%d/emails", client.ApiBaseUrl, bookId),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	email := fake.EmailAddress()
	email1 := fake.EmailAddress()
	err := spClient.DeleteEmailsFromAddressbook(bookId, []*string{&email, &email1})
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_DeleteEmailsFromAddressbook_NoResult(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookId := 1
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/addressbooks/%d/emails", client.ApiBaseUrl, bookId),
		httpmock.NewStringResponder(http.StatusOK, "{}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	email := fake.EmailAddress()
	email1 := fake.EmailAddress()
	err := spClient.DeleteEmailsFromAddressbook(bookId, []*string{&email, &email1})
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}
