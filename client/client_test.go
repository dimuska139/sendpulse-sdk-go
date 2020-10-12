package client

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go"
	"github.com/icrowley/fake"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSendpulseError_Error(t *testing.T) {
	e := SendpulseError{http.StatusInternalServerError, "http://test.com", "Something went wrong", "Test message"}
	assert.Equal(t, fmt.Sprintf("Http code: %d, url: %s, body: %s, message: %s", e.HttpCode, e.Url, e.Body, e.Message), e.Error())
}

func TestClient_ClearToken(t *testing.T) {
	config := sendpulse.Config{
		UserID: fake.Word(),
		Secret: fake.Word(),
	}

	c := NewClient(http.DefaultClient, &config)
	token := fake.Word()
	c.token = token

	tok, _ := c.getToken()
	assert.Equal(t, token, tok)

	c.clearToken()

	emptyTok, _ := c.getToken()
	assert.Equal(t, "", emptyTok)
}

func TestClient_GetToken_Stored(t *testing.T) {
	config := sendpulse.Config{
		UserID: fake.Word(),
		Secret: fake.Word(),
	}
	token := fake.Word()
	c := NewClient(http.DefaultClient, &config)
	c.token = token
	result, err := c.getToken()
	assert.NoError(t, err)
	assert.Equal(t, token, result)
}

func TestClient_GetToken_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/oauth/access_token", ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError,
			`Something went wrong`))

	config := sendpulse.Config{
		UserID: fake.Word(),
		Secret: fake.Word(),
	}

	c := NewClient(http.DefaultClient, &config)
	token, err := c.getToken()
	assert.Error(t, err)
	assert.Equal(t, "", token)
}

func TestClient_GetToken_BadJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/oauth/access_token", ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK,
			`{access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	config := sendpulse.Config{
		UserID: fake.Word(),
		Secret: fake.Word(),
	}

	c := NewClient(http.DefaultClient, &config)
	token, err := c.getToken()
	assert.Error(t, err)
	assert.Equal(t, "", token)
}

func TestClient_GetToken_NoTokenProperty(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/oauth/access_token", ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK,
			`{"token_type": "Bearer","expires_in": 3600}`))

	config := sendpulse.Config{
		UserID: fake.Word(),
		Secret: fake.Word(),
	}

	c := NewClient(http.DefaultClient, &config)
	token, err := c.getToken()
	assert.Error(t, err)
	assert.Equal(t, "", token)
}

func TestClient_GetToken_Success(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	token := fake.Word()
	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s/oauth/access_token", ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "`+token+`","token_type": "Bearer","expires_in": 3600}`))

	config := sendpulse.Config{
		UserID: fake.Word(),
		Secret: fake.Word(),
	}

	c := NewClient(http.DefaultClient, &config)
	newToken, err := c.getToken()
	assert.NoError(t, err)
	assert.Equal(t, token, newToken)
}
