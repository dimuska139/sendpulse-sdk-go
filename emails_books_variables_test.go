package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestBooks_Variables_BadJson(t *testing.T) {
	respBody := `Invalid json`

	bookID := 1
	path := fmt.Sprintf("/addressbooks/%d/variables", bookID)
	url := apiBaseUrl + path

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Books.Variables(bookID)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestBooks_Variables_Error(t *testing.T) {
	respBody := `Invalid json`

	bookID := 1
	path := fmt.Sprintf("/addressbooks/%d/variables", bookID)
	url := apiBaseUrl + path

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusBadRequest, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Books.Variables(bookID)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestBooks_Variables_Success(t *testing.T) {
	respBody := `[
		{
			"name": "емейл",
			"type": "string"
		},
		{
			"name": "имя",
			"type": "string"
		},
		{
			"name": "дата",
			"type": "data"
		},
		{
			"name": "код",
			"type": "number"
		}
	]`

	bookID := 1
	path := fmt.Sprintf("/addressbooks/%d/variables", bookID)
	url := apiBaseUrl + path

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	variables, err := spClient.Emails.Books.Variables(bookID)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(variables))
}
