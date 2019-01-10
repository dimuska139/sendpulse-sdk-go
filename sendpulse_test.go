package sendpulse

import (
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestNoUserId(t *testing.T) {
	_, err := ApiClient("", fake.CharactersN(50), 5)
	assert.Error(t, err)
}

func TestNoSecret(t *testing.T) {
	_, err := ApiClient(fake.CharactersN(50), "", 5)
	assert.Error(t, err)
}

func TestApiClient(t *testing.T) {
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

func TestApiClientError(t *testing.T) {
	url := apiBaseUrl + "/oauth/access_token"
	respBody := `{"error": "invalid_client","error_description": "Client authentication failed.","message": "Client authentication failed.","error_code": 1}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusUnauthorized,
			respBody))

	apiUserId := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	_, err := ApiClient(apiUserId, apiSecret, 5)
	assert.Error(t, err)
	httpError, isHttpError := err.(*HttpError)
	assert.True(t, isHttpError)
	assert.Equal(t, url, httpError.Url)
	assert.Equal(t, http.StatusUnauthorized, httpError.HttpCode)
	assert.Equal(t, respBody, httpError.Message)
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
	_, isHttpError := err.(*HttpError)
	assert.False(t, isHttpError)
	assert.Equal(t, respBody, err.Error())
}
