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

func TestEmails_GetWebhook(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	webhookID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/webhook.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/v2/email-service/webhook/%d", client.ApiBaseUrl, webhookID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	webhook, err := spClient.GetWebhook(webhookID)
	assert.NoError(t, err)
	assert.Equal(t, 162242, int(webhook.ID))
}

func TestEmails_GetWebhook_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	webhookID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/v2/email-service/webhook/%d", client.ApiBaseUrl, webhookID),
		httpmock.NewStringResponder(http.StatusOK, "{\"success:\": false, \"data\": null}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	webhook, err := spClient.GetWebhook(webhookID)
	assert.Error(t, err)
	assert.Nil(t, webhook)
}

func TestEmails_GetWebhook_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	webhookID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/v2/email-service/webhook/%d", client.ApiBaseUrl, webhookID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	webhook, err := spClient.GetWebhook(webhookID)
	assert.Error(t, err)
	assert.Nil(t, webhook)
}

func TestEmails_GetWebhook_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	webhookID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/v2/email-service/webhook/%d", client.ApiBaseUrl, webhookID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	webhook, err := spClient.GetWebhook(webhookID)
	assert.Error(t, err)
	assert.Nil(t, webhook)
}
