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

func TestEmails_GetEmailInfo(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	email := fake.EmailAddress()
	responseBody, _ := ioutil.ReadFile("./testdata/emailInfo.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/emails/%s", client.ApiBaseUrl, email),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetEmailInfo(email)
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, "test@test.com", info[0].Email)
}

func TestEmails_GetEmailInfo_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	email := fake.EmailAddress()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/emails/%s", client.ApiBaseUrl, email),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetEmailInfo(email)
	assert.Error(t, err)
	assert.Nil(t, info)
}

func TestEmails_GetEmailInfo_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	email := fake.EmailAddress()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/emails/%s", client.ApiBaseUrl, email),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetEmailInfo(email)
	assert.Error(t, err)
	assert.Nil(t, info)
}
