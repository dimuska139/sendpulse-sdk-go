package automation360

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

func TestAutomation360_StartEvent(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	eventName := fake.Word()
	responseBody, _ := ioutil.ReadFile("./testdata/startEvent.json")
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/events/name/%s", client.ApiBaseUrl, eventName),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	err := spClient.StartEvent(eventName, map[string]interface{}{
		"email": fake.EmailAddress(),
		"name":  fake.FirstName(),
	})
	assert.NoError(t, err)
}

func TestAutomation360_StartEvent_EmptyEmailAndPhone(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	eventName := fake.Word()

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	err := spClient.StartEvent(eventName, map[string]interface{}{
		"name": fake.FirstName(),
	})
	assert.Error(t, err)
}

func TestAutomation360_StartEvent_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	eventName := fake.Word()
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/events/name/%s", client.ApiBaseUrl, eventName),
		httpmock.NewStringResponder(http.StatusBadRequest, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.StartEvent(eventName, map[string]interface{}{
		"phone": fake.Phone(),
		"name":  fake.FirstName(),
	})
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestAutomation360_StartEvent_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	eventName := fake.Word()
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/events/name/%s", client.ApiBaseUrl, eventName),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.StartEvent(eventName, map[string]interface{}{
		"phone": fake.Phone(),
		"name":  fake.FirstName(),
	})
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestAutomation360_StartEvent_NoResult(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	eventName := fake.Word()
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/events/name/%s", client.ApiBaseUrl, eventName),
		httpmock.NewStringResponder(http.StatusOK, "{}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.StartEvent(eventName, map[string]interface{}{
		"phone": fake.Phone(),
		"name":  fake.FirstName(),
	})
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}
