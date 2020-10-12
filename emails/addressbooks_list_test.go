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

func TestEmails_GetAddressbooks(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/addressBooksList.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks?limit=0&offset=10", client.ApiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	books, err := spClient.GetAddressbooks(0, 10)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(books))
}

func TestEmails_GetAddressbooks_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks?limit=0&offset=10", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	books, err := spClient.GetAddressbooks(0, 10)
	assert.Error(t, err)
	assert.Equal(t, 0, len(books))
}

func TestEmails_GetAddressbooks_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/addressbooks?limit=0&offset=10", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	books, err := spClient.GetAddressbooks(0, 10)
	assert.Error(t, err)
	assert.Equal(t, 0, len(books))
}
