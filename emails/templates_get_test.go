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

func TestEmails_GetTemplate(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	templateID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/template.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/template/%d", client.ApiBaseUrl, templateID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	template, err := spClient.GetTemplate(templateID)
	assert.NoError(t, err)
	assert.NotNil(t, template)
	assert.Equal(t, "c7a94d4f8395ae5a4183423309d5e99b", template.ID)
}

func TestEmails_GetTemplate_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	templateID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/template/%d", client.ApiBaseUrl, templateID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	template, err := spClient.GetTemplate(templateID)
	assert.Error(t, err)
	assert.Nil(t, template)
}

func TestEmails_GetTemplate_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	templateID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/template/%d", client.ApiBaseUrl, templateID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	template, err := spClient.GetTemplate(templateID)
	assert.Error(t, err)
	assert.Nil(t, template)
}
