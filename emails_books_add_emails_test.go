package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestBooks_AddEmails_BadJson(t *testing.T) {
	addressBookId := 1

	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookId)
	url := apiBaseUrl + path
	respBody := `Invalid json`

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	emails := []Email{
		{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
		{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
	}

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	err := spClient.Emails.Books.AddEmails(addressBookId, emails, make(map[string]string), "")
	assert.Error(t, err)

	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestBooks_AddEmails_Success(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	emails := []Email{
		{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
		{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
	}

	addressBookId := 1
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/addressbooks/%d/emails", apiBaseUrl, addressBookId),
		httpmock.NewStringResponder(http.StatusOK, `{
    		"result": true
		}`))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	extraParams := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}
	err := spClient.Emails.Books.AddEmails(addressBookId, emails, extraParams, fake.EmailAddress())
	assert.NoError(t, err)
}

func TestBooks_AddEmails_InvalidResponse(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	emails := []Email{
		{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
		{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
	}

	addressBookId := 1
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/addressbooks/%d/emails", apiBaseUrl, addressBookId),
		httpmock.NewStringResponder(http.StatusOK, `{
    		"foo": true
		}`))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	extraParams := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}
	err := spClient.Emails.Books.AddEmails(addressBookId, emails, extraParams, fake.EmailAddress())
	assert.Error(t, err)

	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestBooks_AddEmails_Error(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	emails := []Email{
		{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
		{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
	}

	addressBookId := 1
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/addressbooks/%d/emails", apiBaseUrl, addressBookId),
		httpmock.NewStringResponder(http.StatusInternalServerError, `{
    		"result": true
		}`))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	extraParams := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}
	err := spClient.Emails.Books.AddEmails(addressBookId, emails, extraParams, fake.EmailAddress())
	assert.Error(t, err)

	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}
