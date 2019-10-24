package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestBooks_Emails_BadJson(t *testing.T) {
	var bookId uint = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	var limit uint = 10
	var offset uint = 0
	url := fmt.Sprintf("%s/addressbooks/%d/emails", apiBaseUrl, bookId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `Very bad json]`

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	contacts, err := spClient.Emails.Books.Emails(bookId, limit, offset)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Nil(t, contacts)
}

func TestBooks_Emails_Error(t *testing.T) {
	var bookId uint = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	var limit uint = 10
	var offset uint = 0
	url := fmt.Sprintf("%s/addressbooks/%d/emails?limit=%d&offset=%d", apiBaseUrl, bookId, limit, offset)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `[
	  {
		"email": "test@test.com",
		"status": "0",
		"status_explain": "New",
		"variables": [
		  {
		"name": "имя переменной",
		"type": "string",
		"value": "значение переменной"
		  },
		  {
		"name": "имя переменной",
		"type": "string",
		"value": "значение переменной"
		  }
		]
	  },
	  {
		"email": "test2@test.com",
		"status": "0",
		"status_explain": "New",
		"variables": [
		  {
		"name": "имя переменной",
		"type": "string",
		"value": "значение переменной"
		  },
		  {
		"name": "имя переменной",
		"type": "string",
		"value": "значение переменной"
		  }
		]
	  }
	]`

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusInternalServerError, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	contacts, err := spClient.Emails.Books.Emails(bookId, limit, offset)

	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Nil(t, contacts)
}

func TestBooks_Emails_Success(t *testing.T) {
	var bookId uint = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	var limit uint = 10
	var offset uint = 0
	url := fmt.Sprintf("%s/addressbooks/%d/emails?limit=%d&offset=%d", apiBaseUrl, bookId, limit, offset)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := `[
	  {
		"email": "test@test.com",
		"status": "0",
		"status_explain": "New",
		"variables": [
		  {
		"name": "имя переменной",
		"type": "string",
		"value": "значение переменной"
		  },
		  {
		"name": "имя переменной",
		"type": "string",
		"value": "значение переменной"
		  }
		]
	  },
	  {
		"email": "test2@test.com",
		"status": "0",
		"status_explain": "New",
		"variables": [
		  {
		"name": "имя переменной",
		"type": "string",
		"value": "значение переменной"
		  },
		  {
		"name": "имя переменной",
		"type": "string",
		"value": "значение переменной"
		  }
		]
	  }
	]`

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Books.Emails(bookId, limit, offset)
	assert.NoError(t, err)
}
