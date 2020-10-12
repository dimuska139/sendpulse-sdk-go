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

func TestEmails_DeleteSender(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/deleteSender.json")
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/senders", client.ApiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	assert.NoError(t, spClient.DeleteSender(fake.EmailAddress()))
}

func TestEmails_DeleteSender_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/senders", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusBadRequest, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.DeleteSender(fake.EmailAddress())
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_DeleteSender_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/senders", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.DeleteSender(fake.EmailAddress())
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_DeleteSender_NoResult(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/senders", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, "{}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.DeleteSender(fake.EmailAddress())
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}
