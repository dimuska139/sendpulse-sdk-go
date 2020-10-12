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

func TestEmails_DeleteCampaign(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignId := 1
	responseBody, _ := ioutil.ReadFile("./testdata/deleteCampaign.json")
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/campaigns/%d", client.ApiBaseUrl, campaignId),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	assert.NoError(t, spClient.DeleteCampaign(campaignId))
}

func TestEmails_DeleteCampaign_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignId := 1
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/campaigns/%d", client.ApiBaseUrl, campaignId),
		httpmock.NewStringResponder(http.StatusBadRequest, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.DeleteCampaign(campaignId)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_DeleteCampaign_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignId := 1
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/campaigns/%d", client.ApiBaseUrl, campaignId),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.DeleteCampaign(campaignId)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_DeleteCampaign_NoResult(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignId := 1
	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s/campaigns/%d", client.ApiBaseUrl, campaignId),
		httpmock.NewStringResponder(http.StatusOK, "{}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.DeleteCampaign(campaignId)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}
