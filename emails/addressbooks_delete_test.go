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

func TestEmails_DeleteAddressBook(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookId := 1
	responseBody, _ := ioutil.ReadFile("./testdata/deleteAddressBook.json")
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/addressbooks/%d", client.ApiBaseUrl, bookId),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	assert.NoError(t, spClient.DeleteAddressBook(bookId))
}

func TestEmails_DeleteAddressBook_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookId := 1
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/addressbooks/%d", client.ApiBaseUrl, bookId),
		httpmock.NewStringResponder(http.StatusBadRequest, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.DeleteAddressBook(bookId)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_DeleteAddressBook_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookId := 1
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/addressbooks/%d", client.ApiBaseUrl, bookId),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.DeleteAddressBook(bookId)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_DeleteAddressBook_NoResult(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookId := 1
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/addressbooks/%d", client.ApiBaseUrl, bookId),
		httpmock.NewStringResponder(http.StatusOK, "{}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.DeleteAddressBook(bookId)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}
