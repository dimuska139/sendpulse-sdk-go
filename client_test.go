package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestSendpulseError_Error(t *testing.T) {
	e := SendpulseError{http.StatusInternalServerError, "http://test.com", "Something went wrong", "Test message"}
	assert.Equal(t, fmt.Sprintf("Http code: %d, url: %s, body: %s, message: %s", e.HttpCode, e.Url, e.Body, e.Message), e.Error())
}

func TestClient_ClearToken(t *testing.T) {
	config := Config{
		UserID:  fake.Word(),
		Secret:  fake.Word(),
		Timeout: 0,
	}

	c := NewClient(config)
	token := fake.Word()
	c.token = token

	tok, _ := c.getToken()
	assert.Equal(t, token, tok)

	c.clearToken()

	emptyTok, _ := c.getToken()
	assert.Equal(t, "", emptyTok)
}

func TestClient_GetToken_Stored(t *testing.T) {
	config := Config{
		UserID:  fake.Word(),
		Secret:  fake.Word(),
		Timeout: 0,
	}
	token := fake.Word()
	c := NewClient(config)
	c.token = token
	result, err := c.getToken()
	assert.NoError(t, err)
	assert.Equal(t, token, result)
}

func TestClient_GetToken_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusInternalServerError,
			`Something went wrong`))

	config := Config{
		UserID:  fake.Word(),
		Secret:  fake.Word(),
		Timeout: 0,
	}

	c := NewClient(config)
	token, err := c.getToken()
	assert.Error(t, err)
	assert.Equal(t, "", token)
}

func TestClient_GetToken_BadJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	config := Config{
		UserID:  fake.Word(),
		Secret:  fake.Word(),
		Timeout: 0,
	}

	c := NewClient(config)
	token, err := c.getToken()
	assert.Error(t, err)
	assert.Equal(t, "", token)
}

func TestClient_GetToken_NoTokenProperty(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"token_type": "Bearer","expires_in": 3600}`))

	config := Config{
		UserID:  fake.Word(),
		Secret:  fake.Word(),
		Timeout: 0,
	}

	c := NewClient(config)
	token, err := c.getToken()
	assert.Error(t, err)
	assert.Equal(t, "", token)
}

func TestClient_GetToken_Success(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	token := fake.Word()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "`+token+`","token_type": "Bearer","expires_in": 3600}`))

	config := Config{
		UserID:  fake.Word(),
		Secret:  fake.Word(),
		Timeout: 0,
	}

	c := NewClient(config)
	newToken, err := c.getToken()
	assert.NoError(t, err)
	assert.Equal(t, token, newToken)
}
