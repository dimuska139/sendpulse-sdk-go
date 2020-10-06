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

func TestEmails_UpdateAddressbook(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookId := 1
	responseBody, _ := ioutil.ReadFile("./testdata/updatedAddressBook.json")
	httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%s/addressbooks/%d", client.ApiBaseUrl, bookId),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	assert.NoError(t, spClient.UpdateAddressbook(bookId, fake.Word()))
}

func TestEmails_UpdateAddressBook_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookId := 1
	httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%s/addressbooks/%d", client.ApiBaseUrl, bookId),
		httpmock.NewStringResponder(http.StatusBadRequest, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.UpdateAddressbook(bookId, fake.Word())
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_UpdateAddressBook_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookId := 1
	httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%s/addressbooks/%d", client.ApiBaseUrl, bookId),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.UpdateAddressbook(bookId, fake.Word())
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_UpdateAddressBook_FalseResult(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookId := 1
	httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%s/addressbooks/%d", client.ApiBaseUrl, bookId),
		httpmock.NewStringResponder(http.StatusOK, "{\"result\":false}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.UpdateAddressbook(bookId, fake.Word())
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}
