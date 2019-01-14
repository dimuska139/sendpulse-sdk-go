package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestResponseErrorData(t *testing.T) {
	e := SendpulseError{http.StatusInternalServerError, "http://test.com", "Something went wrong", "Test message"}
	assert.Equal(t, fmt.Sprintf("Http code: %d, url: %s, body: %s, message: %s", e.HttpCode, e.Url, e.Body, e.Message), e.Error())
}

func TestGetTokenResponseError(t *testing.T) {
	url := apiBaseUrl + "/oauth/access_token"
	respBody := `Invalid json`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	apiUserId := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	_, err := ApiClient(apiUserId, apiSecret, 5)
	assert.Error(t, err)
	responseError, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)

	assert.Equal(t, http.StatusOK, responseError.HttpCode)
	assert.Equal(t, url, responseError.Url)
	assert.Equal(t, respBody, responseError.Body)
}

func TestGetTokenNoAccessToken(t *testing.T) {
	url := apiBaseUrl + "/oauth/access_token"
	respBody := `{"token_type": "Bearer","expires_in": 3600}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	apiUserId := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	_, err := ApiClient(apiUserId, apiSecret, 5)
	assert.Error(t, err)
	responseError, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)

	assert.Equal(t, http.StatusOK, responseError.HttpCode)
	assert.Equal(t, url, responseError.Url)
	assert.Equal(t, respBody, responseError.Body)
}

func TestGetTokenSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	apiUserId := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	_, err := ApiClient(apiUserId, apiSecret, 5)
	assert.NoError(t, err)
}
