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
