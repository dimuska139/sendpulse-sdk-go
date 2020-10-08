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
	"time"
)

func TestEmails_UpdateCampaign(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/updatedCampaignBook.json")
	httpmock.RegisterResponder(http.MethodPatch, fmt.Sprintf("%s/campaigns/%d", client.ApiBaseUrl, campaignID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	assert.NoError(t, spClient.UpdateCampaign(campaignID, UpdateCampaignDto{
		ID:          campaignID,
		Name:        fake.Word(),
		SenderName:  fake.FirstName(),
		SenderEmail: fake.EmailAddress(),
		Subject:     fake.Word(),
		Body:        fake.Word(),
		TemplateID:  123456,
		SendDate:    time.Now(),
	}))
}

func TestEmails_UpdateCampaign_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	httpmock.RegisterResponder(http.MethodPatch, fmt.Sprintf("%s/campaigns/%d", client.ApiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusBadRequest, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.UpdateCampaign(campaignID, UpdateCampaignDto{
		ID:          campaignID,
		Name:        fake.Word(),
		SenderName:  fake.FirstName(),
		SenderEmail: fake.EmailAddress(),
		Subject:     fake.Word(),
		Body:        fake.Word(),
		TemplateID:  123456,
		SendDate:    time.Now(),
	})
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_UpdateCampaign_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	httpmock.RegisterResponder(http.MethodPatch, fmt.Sprintf("%s/campaigns/%d", client.ApiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.UpdateCampaign(campaignID, UpdateCampaignDto{
		ID:          campaignID,
		Name:        fake.Word(),
		SenderName:  fake.FirstName(),
		SenderEmail: fake.EmailAddress(),
		Subject:     fake.Word(),
		Body:        fake.Word(),
		TemplateID:  123456,
		SendDate:    time.Now(),
	})
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_UpdateCampaign_FalseResult(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	campaignID := 1
	httpmock.RegisterResponder(http.MethodPatch, fmt.Sprintf("%s/campaigns/%d", client.ApiBaseUrl, campaignID),
		httpmock.NewStringResponder(http.StatusOK, "{\"result\":false}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.UpdateCampaign(campaignID, UpdateCampaignDto{
		ID:          campaignID,
		Name:        fake.Word(),
		SenderName:  fake.FirstName(),
		SenderEmail: fake.EmailAddress(),
		Subject:     fake.Word(),
		Body:        fake.Word(),
		TemplateID:  123456,
		SendDate:    time.Now(),
	})
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}
