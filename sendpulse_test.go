package sendpulse

import (
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestApiClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	timeout := 10
	config := Config{
		UserID:  fake.CharactersN(10),
		Secret:  fake.CharactersN(10),
		Timeout: timeout,
	}
	client, err := ApiClient(config)
	assert.NoError(t, err)
	assert.Equal(t, timeout, client.client.config.Timeout)
}

func TestApiClient_NullTimeout(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	timeout := 0
	config := Config{
		UserID:  fake.CharactersN(10),
		Secret:  fake.CharactersN(10),
		Timeout: timeout,
	}
	client, err := ApiClient(config)
	assert.NoError(t, err)
	assert.Equal(t, 5, client.client.config.Timeout)
}
