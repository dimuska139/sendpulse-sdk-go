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

func TestEmails_GetWebhooks(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/webhooks.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/v2/email-service/webhook", client.ApiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	webhooks, err := spClient.GetWebhooks()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(webhooks))
	assert.Equal(t, 162242, int(webhooks[0].ID))
}

func TestEmails_GetWebhooks_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/v2/email-service/webhook", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, "{\"success:\": false, \"data\": []}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	webhooks, err := spClient.GetWebhooks()
	assert.Error(t, err)
	assert.Equal(t, 0, len(webhooks))
}

func TestEmails_GetWebhooks_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/v2/email-service/webhook", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	webhooks, err := spClient.GetWebhooks()
	assert.Error(t, err)
	assert.Equal(t, 0, len(webhooks))
}

func TestEmails_GetWebhooks_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/v2/email-service/webhook", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	webhooks, err := spClient.GetWebhooks()
	assert.Error(t, err)
	assert.Equal(t, 0, len(webhooks))
}
