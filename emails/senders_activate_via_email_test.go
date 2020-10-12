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

func TestEmails_ActivateSenderViaEmail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	email := fake.EmailAddress()
	responseBody, _ := ioutil.ReadFile("./testdata/activateSenderViaEmail.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/senders/%s/code", client.ApiBaseUrl, email),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	assert.NoError(t, spClient.ActivateSenderViaEmail(email))
}

func TestEmails_ActivateSenderViaEmail_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	email := fake.EmailAddress()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/senders/%s/code", client.ApiBaseUrl, email),
		httpmock.NewStringResponder(http.StatusBadRequest, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.ActivateSenderViaEmail(email)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_ActivateSenderViaEmail_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	email := fake.EmailAddress()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/senders/%s/code", client.ApiBaseUrl, email),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.ActivateSenderViaEmail(email)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_ActivateSenderViaEmail_NoResult(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	email := fake.EmailAddress()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/senders/%s/code", client.ApiBaseUrl, email),
		httpmock.NewStringResponder(http.StatusOK, "{}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.ActivateSenderViaEmail(email)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}
