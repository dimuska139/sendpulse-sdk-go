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

func TestEmails_CreateCampaign(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/createdCampaign.json")
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/campaigns", client.ApiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	created, err := spClient.CreateCampaign(CreateCampaignDto{
		SenderName:   fake.FirstName(),
		SenderEmail:  fake.EmailAddress(),
		Subject:      fake.Word(),
		Body:         fake.Word(),
		TemplateID:   12345,
		BodyAMP:      fake.Word(),
		ListID:       12345,
		SegmentID:    12345,
		SendTestOnly: nil,
		SendDate:     time.Now(),
		Name:         fake.Word(),
		Attachments: map[string]string{
			"test": fake.Word(),
			"name": fake.Word(),
		},
		IsDraft: true,
	})
	assert.NoError(t, err)
	assert.Equal(t, 27, int(*created.ID))
}

func TestEmails_CreateCampaignWithTestEmails(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/createdCampaign.json")
	httpmock.RegisterResponder(http.MethodPatch, fmt.Sprintf("%s/campaigns", client.ApiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	created, err := spClient.CreateCampaign(CreateCampaignDto{
		SenderName:   fake.FirstName(),
		SenderEmail:  fake.EmailAddress(),
		Subject:      fake.Word(),
		Body:         fake.Word(),
		TemplateID:   12345,
		BodyAMP:      fake.Word(),
		ListID:       12345,
		SegmentID:    12345,
		SendTestOnly: []string{fake.EmailAddress(), fake.EmailAddress(), fake.EmailAddress()},
		SendDate:     time.Now(),
		Name:         fake.Word(),
		Attachments: map[string]string{
			"test": fake.Word(),
			"name": fake.Word(),
		},
		IsDraft: true,
	})
	assert.NoError(t, err)
	assert.Equal(t, 27, int(*created.ID))
}

func TestEmails_CreateCampaign_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/campaigns", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	created, err := spClient.CreateCampaign(CreateCampaignDto{
		SenderName:   fake.FirstName(),
		SenderEmail:  fake.EmailAddress(),
		Subject:      fake.Word(),
		Body:         fake.Word(),
		TemplateID:   12345,
		BodyAMP:      fake.Word(),
		ListID:       12345,
		SegmentID:    12345,
		SendTestOnly: nil,
		SendDate:     time.Now(),
		Name:         fake.Word(),
		Attachments: map[string]string{
			"test": fake.Word(),
			"name": fake.Word(),
		},
		IsDraft: true,
	})
	assert.Nil(t, created)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_CreateCampaign_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/campaigns", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	created, err := spClient.CreateCampaign(CreateCampaignDto{
		SenderName:   fake.FirstName(),
		SenderEmail:  fake.EmailAddress(),
		Subject:      fake.Word(),
		Body:         fake.Word(),
		TemplateID:   12345,
		BodyAMP:      fake.Word(),
		ListID:       12345,
		SegmentID:    12345,
		SendTestOnly: nil,
		SendDate:     time.Now(),
		Name:         fake.Word(),
		Attachments: map[string]string{
			"test": fake.Word(),
			"name": fake.Word(),
		},
		IsDraft: true,
	})
	assert.Nil(t, created)
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}
