package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestTemplates_Create_IncorrectJson(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/template"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respBody := `Incorrect json`

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 1,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	realID, err := spClient.Emails.Templates.Create(fake.Word(), fake.Word(), fake.Word())
	assert.Error(t, err)
	spErr, isSpError := err.(*SendpulseError)
	assert.True(t, isSpError)
	assert.Equal(t, 0, realID)

	assert.Equal(t, http.StatusOK, spErr.HttpCode)
	assert.Equal(t, path, spErr.Url)
	assert.Equal(t, respBody, spErr.Body)
}

func TestTemplates_Create_Error(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/template"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respBody := `{}`

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusBadRequest, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 1,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	realID, err := spClient.Emails.Templates.Create(fake.Word(), fake.Word(), fake.Word())
	assert.Error(t, err)
	spErr, isSpError := err.(*SendpulseError)
	assert.True(t, isSpError)
	assert.Equal(t, 0, realID)

	assert.Equal(t, http.StatusBadRequest, spErr.HttpCode)
	assert.Equal(t, path, spErr.Url)
	assert.Equal(t, respBody, spErr.Body)
}

func TestTemplates_Create_NoResult(t *testing.T) {
	expectedRealID := 1
	respBody := fmt.Sprintf(`{"real_id":%d}`, expectedRealID)
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/template"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 1,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	realID, err := spClient.Emails.Templates.Create(fake.Word(), fake.Word(), fake.Word())
	assert.Error(t, err)
	spErr, isSpError := err.(*SendpulseError)
	assert.True(t, isSpError)
	assert.Equal(t, 0, realID)

	assert.Equal(t, http.StatusOK, spErr.HttpCode)
	assert.Equal(t, path, spErr.Url)
	assert.Equal(t, respBody, spErr.Body)
}

func TestTemplates_Create_NoRealID(t *testing.T) {
	respBody := `{"result":true}`
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/template"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 1,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	realID, err := spClient.Emails.Templates.Create(fake.Word(), fake.Word(), fake.Word())
	assert.Error(t, err)
	spErr, isSpError := err.(*SendpulseError)
	assert.True(t, isSpError)
	assert.Equal(t, 0, realID)

	assert.Equal(t, http.StatusOK, spErr.HttpCode)
	assert.Equal(t, path, spErr.Url)
	assert.Equal(t, respBody, spErr.Body)
}

func TestTemplates_Create_Success(t *testing.T) {
	expectedRealID := 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/template"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(`{"result":true,"real_id":%d}`, expectedRealID)))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 1,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	realID, err := spClient.Emails.Templates.Create(fake.Word(), fake.Word(), fake.Word())
	assert.NoError(t, err)
	assert.Equal(t, expectedRealID, realID)
}
