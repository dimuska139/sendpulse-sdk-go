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

func TestEmails_UnsubscribeEmailsFromAddressbook(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/unsubscribeEmails.json")
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/addressbooks/%d/emails/unsubscribe", client.ApiBaseUrl, bookID),
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
	err := spClient.UnsubscribeEmailsFromAddressbook(bookID, []*string{&email, &email1})
	assert.NoError(t, err)
}

func TestEmails_UnsubscribeEmailsFromAddressbook_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/addressbooks/%d/emails/unsubscribe", client.ApiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)
	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	email := fake.EmailAddress()
	email1 := fake.EmailAddress()
	err := spClient.UnsubscribeEmailsFromAddressbook(bookID, []*string{&email, &email1})
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_UnsubscribeEmailsFromAddressbook_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/addressbooks/%d/emails/unsubscribe", client.ApiBaseUrl, bookID),
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
	err := spClient.UnsubscribeEmailsFromAddressbook(bookID, []*string{&email, &email1})
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_UnsubscribeEmailsFromAddressbook_NoResult(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/addressbooks/%d/emails/unsubscribe", client.ApiBaseUrl, bookID),
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
	err := spClient.UnsubscribeEmailsFromAddressbook(bookID, []*string{&email, &email1})
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}
