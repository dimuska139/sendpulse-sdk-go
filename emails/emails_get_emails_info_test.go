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

func TestEmails_GetEmailsInfo(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/emails.json")
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/emails", client.ApiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	email1 := "example@yourdomain.com"
	email2 := "example2@yourdomain.com"
	info, err := spClient.GetEmailsInfo(email1, email2)

	details1, email1exists := info[email1]
	details2, email2exists := info[email2]

	assert.True(t, email1exists)
	assert.True(t, email2exists)

	fmt.Println(*details1[0])
	assert.Equal(t, email1, details1[0].Email)
	assert.Equal(t, 391141, int(details2[0].BookID))
	assert.NoError(t, err)
}

func TestEmails_GetEmailsInfo_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/emails", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetEmailsInfo(fake.EmailAddress(), fake.EmailAddress())
	assert.Error(t, err)
	assert.Nil(t, info)
}

func TestEmails_GetEmailsInfo_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/emails", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	info, err := spClient.GetEmailsInfo(fake.EmailAddress(), fake.EmailAddress())
	assert.Error(t, err)
	assert.Nil(t, info)
}
