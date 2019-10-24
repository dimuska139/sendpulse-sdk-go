package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestBooks_Update_BadJson(t *testing.T) {
	var bookId uint = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/addressbooks/%d", apiBaseUrl, bookId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
    	result": true
	}`

	httpmock.RegisterResponder("PUT", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	err := spClient.Emails.Books.Update(bookId, fake.Word())

	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestBooks_Update_InvalidResponse(t *testing.T) {
	var bookId uint = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/addressbooks/%d", apiBaseUrl, bookId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
    	"strange_result": true
	}`

	httpmock.RegisterResponder("PUT", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	err := spClient.Emails.Books.Update(bookId, fake.Word())

	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestBooks_Update_Error(t *testing.T) {
	var bookId uint = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/addressbooks/%d", apiBaseUrl, bookId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
    	"strange_result": true
	}`

	httpmock.RegisterResponder("PUT", url,
		httpmock.NewStringResponder(http.StatusBadRequest, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	err := spClient.Emails.Books.Update(bookId, fake.Word())

	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestBooks_Update_Success(t *testing.T) {
	var bookId uint = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	url := fmt.Sprintf("%s/addressbooks/%d", apiBaseUrl, bookId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{
    	"result": true
	}`

	httpmock.RegisterResponder("PUT", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	assert.NoError(t, spClient.Emails.Books.Update(bookId, fake.Word()))
}
