package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestBooks_Create_IncorrectJson(t *testing.T) {
	bookName := fake.Word()
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/addressbooks"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respBody := `Incorrect json`

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	bookId, err := spClient.Emails.Books.Create(bookName)
	assert.Error(t, err)
	spErr, isSpError := err.(*SendpulseError)
	assert.True(t, isSpError)
	assert.Nil(t, bookId)

	assert.Equal(t, http.StatusOK, spErr.HttpCode)
	assert.Equal(t, path, spErr.Url)
	assert.Equal(t, respBody, spErr.Body)
}

func TestBooks_Create_NoIdInResponse(t *testing.T) {
	bookName := fake.Word()
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/addressbooks"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `{}`

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	bookId, err := spClient.Emails.Books.Create(bookName)
	assert.Error(t, err)
	httpErr, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Nil(t, bookId)

	assert.Equal(t, http.StatusOK, httpErr.HttpCode)
	assert.Equal(t, path, httpErr.Url)
	assert.Equal(t, respBody, httpErr.Body)
}

func TestBooks_Create_Error(t *testing.T) {
	bookName := fake.Word()
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/addressbooks"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusBadRequest, ""))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	bookId, err := spClient.Emails.Books.Create(bookName)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Nil(t, bookId)
}

func TestBooks_Create_Success(t *testing.T) {
	bookName := fake.Word()
	var newBookId uint = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/addressbooks"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := fmt.Sprintf(`{
    	"id": %d
	}`, newBookId)

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	bookId, err := spClient.Emails.Books.Create(bookName)
	assert.NoError(t, err)
	assert.Equal(t, newBookId, *bookId)
}
