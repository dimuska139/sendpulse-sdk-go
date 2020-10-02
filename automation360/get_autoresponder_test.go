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

func TestAutomation360_GetAutoresponder(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/autoresponder.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d", client.ApiBaseUrl, autoresponderID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	autoresponder, err := spClient.GetAutoresponder(autoresponderID)
	assert.NoError(t, err)
	assert.NotNil(t, autoresponder)
}

func TestAutomation360_GetAutoresponder_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d", client.ApiBaseUrl, autoresponderID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	autoresponder, err := spClient.GetAutoresponder(autoresponderID)
	assert.Error(t, err)
	assert.Nil(t, autoresponder)
}

func TestAutomation360_GetAutoresponder_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d", client.ApiBaseUrl, autoresponderID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	autoresponder, err := spClient.GetAutoresponder(autoresponderID)
	assert.Error(t, err)
	assert.Nil(t, autoresponder)
}
