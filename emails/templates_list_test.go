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

func TestEmails_GetTemplates(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/templates.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/templates?limit=0&offset=10", client.ApiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	templates, err := spClient.GetTemplates(0, 10, "")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(templates))
	assert.Equal(t, "c7a94d4f8395ae5a4183423309d5e99b", *templates[0].ID)
}

func TestEmails_GetTemplatesByOwner(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/templates.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/templates?limit=0&offset=10&owner=sendpulse", client.ApiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	templates, err := spClient.GetTemplates(0, 10, "sendpulse")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(templates))
	assert.Equal(t, "c7a94d4f8395ae5a4183423309d5e99b", *templates[0].ID)
}

func TestEmails_GetTemplates_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/templates?limit=0&offset=10", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	templates, err := spClient.GetTemplates(0, 10, "")
	assert.Error(t, err)
	assert.Equal(t, 0, len(templates))
}

func TestEmails_GetTemplates_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/templates?limit=0&offset=10", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	templates, err := spClient.GetTemplates(0, 10, "")
	assert.Error(t, err)
	assert.Equal(t, 0, len(templates))
}
