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

	config := Config{
		UserID:  fake.CharactersN(10),
		Secret:  fake.CharactersN(10),
		Timeout: 5,
	}
	_, err := ApiClient(config)
	assert.NoError(t, err)
}
