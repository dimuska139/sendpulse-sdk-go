package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestBooks_Emails_Total_BadJson(t *testing.T) {
	var bookId int = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/addressbooks/%d/emails/total", apiBaseUrl, bookId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `
		"total": 1
	}`

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	total, err := spClient.Emails.Books.EmailsTotal(bookId)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Equal(t, 0, total)
}

func TestBooks_Emails_Total_Error(t *testing.T) {
	var bookId int = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/addressbooks/%d/emails/total", apiBaseUrl, bookId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
		"total": 1
	}`

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusInternalServerError, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	total, err := spClient.Emails.Books.EmailsTotal(bookId)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Equal(t, 0, total)
}

func TestBooks_Emails_Total_InvalidResponse(t *testing.T) {
	var bookId int = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/addressbooks/%d/emails/total", apiBaseUrl, bookId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{}`

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	total, err := spClient.Emails.Books.EmailsTotal(bookId)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Equal(t, 0, total)
}

func TestBooks_Emails_Total_Success(t *testing.T) {
	var bookId int = 1
	var expectedTotal int = 12345
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/addressbooks/%d/emails/total", apiBaseUrl, bookId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{"total": ` + fmt.Sprintf("%d", expectedTotal) + `}`

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	total, err := spClient.Emails.Books.EmailsTotal(bookId)
	assert.NoError(t, err)
	assert.Equal(t, expectedTotal, total)
}
