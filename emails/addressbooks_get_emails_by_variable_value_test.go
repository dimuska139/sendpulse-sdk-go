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

func TestEmails_GetAddressbookEmailsByVariable(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	variableName := fake.Word()
	variableValue := fake.Word()
	responseBody, _ := ioutil.ReadFile("./testdata/emailsList.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/variables/%s/%v", client.ApiBaseUrl, bookID, variableName, variableValue),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	emails, err := spClient.GetAddressbookEmailsByVariable(bookID, variableName, variableValue)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(emails))
}

func TestEmails_GetAddressbookEmailsByVariable_Empty(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	variableName := fake.Word()
	variableValue := fake.Word()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/variables/%s/%v", client.ApiBaseUrl, bookID, variableName, variableValue),
		httpmock.NewStringResponder(http.StatusOK, "[]"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	emails, err := spClient.GetAddressbookEmailsByVariable(bookID, variableName, variableValue)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(emails))
}

func TestEmails_GetAddressbookEmailsByVariable_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	variableName := fake.Word()
	variableValue := fake.Word()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/variables/%s/%v", client.ApiBaseUrl, bookID, variableName, variableValue),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	emails, err := spClient.GetAddressbookEmailsByVariable(bookID, variableName, variableValue)
	assert.Error(t, err)
	assert.Equal(t, 0, len(emails))
}

func TestEmails_GetAddressbookEmailsByVariable_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	bookID := 1
	variableName := fake.Word()
	variableValue := fake.Word()
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks/%d/variables/%s/%v", client.ApiBaseUrl, bookID, variableName, variableValue),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	emails, err := spClient.GetAddressbookEmailsByVariable(bookID, variableName, variableValue)
	assert.Error(t, err)
	assert.Equal(t, 0, len(emails))
}
