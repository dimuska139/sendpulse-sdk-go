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

func TestAutomation360_GetConversions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/conversions.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d/conversions", client.ApiBaseUrl, autoresponderID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	conversions, err := spClient.GetConversions(autoresponderID)
	assert.NoError(t, err)
	assert.NotNil(t, conversions)
	assert.Equal(t, 5, int(conversions.TotalConversions))
}

func TestAutomation360_GetConversions_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d/conversions", client.ApiBaseUrl, autoresponderID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	conversions, err := spClient.GetConversions(autoresponderID)
	assert.Error(t, err)
	assert.Nil(t, conversions)
}

func TestAutomation360_GetConversions_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d/conversions", client.ApiBaseUrl, autoresponderID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	conversions, err := spClient.GetConversions(autoresponderID)
	assert.Error(t, err)
	assert.Nil(t, conversions)
}
