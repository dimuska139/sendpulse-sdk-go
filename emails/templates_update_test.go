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

func TestEmails_UpdateTemplate(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	templateID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/updateTemplate.json")
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/template/edit/%d", client.ApiBaseUrl, templateID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	assert.NoError(t, spClient.UpdateTemplate(templateID, fake.Word(), fake.Word()))
}

func TestEmails_UpdateTemplate_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	templateID := 1
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/template/edit/%d", client.ApiBaseUrl, templateID),
		httpmock.NewStringResponder(http.StatusBadRequest, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.UpdateTemplate(templateID, fake.Word(), fake.Word())
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_UpdateTemplate_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	templateID := 1
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/template/edit/%d", client.ApiBaseUrl, templateID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.UpdateTemplate(templateID, fake.Word(), fake.Word())
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}

func TestEmails_UpdateTemplate_FalseResult(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	templateID := 1
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/template/edit/%d", client.ApiBaseUrl, templateID),
		httpmock.NewStringResponder(http.StatusOK, "{\"result\":false}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)
	err := spClient.UpdateTemplate(templateID, fake.Word(), fake.Word())
	assert.Error(t, err)
	_, isResponseError := err.(*client.SendpulseError)
	assert.True(t, isResponseError)
}
