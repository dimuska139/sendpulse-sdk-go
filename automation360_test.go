package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestStartEventEmptyEventName(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))
	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	variables := make(map[string]interface{})
	variables["phone"] = fake.Phone()
	variables["name"] = fake.FullName()
	err := spClient.Automation360.StartEvent("", variables)
	assert.Error(t, err)
	_, isSPError := err.(*SendpulseError)
	assert.False(t, isSPError)
}

func TestStartEventNoPhoneAndEmail(t *testing.T) {
	eventName := fake.Word()
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))
	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	variables := make(map[string]interface{})
	variables["name"] = fake.FullName()
	err := spClient.Automation360.StartEvent(eventName, variables)
	assert.Error(t, err)
	_, isSPError := err.(*SendpulseError)
	assert.False(t, isSPError)
}

func TestStartEventNotExists(t *testing.T) {
	eventName := fake.Word()

	path := fmt.Sprintf("/events/name/%s", eventName)
	url := apiBaseUrl + path

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	respBody := `{"error_code": 102,"message": "Event not exists"}`
	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusBadRequest,
			respBody))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	variables := make(map[string]interface{})
	variables["email"] = fake.EmailAddress()
	variables["name"] = fake.FullName()
	err := spClient.Automation360.StartEvent(eventName, variables)
	assert.Error(t, err)
	ResponseError, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Equal(t, http.StatusBadRequest, ResponseError.HttpCode)
	assert.Equal(t, path, ResponseError.Url)
	assert.Equal(t, respBody, ResponseError.Body)
	assert.Equal(t, "", ResponseError.Message)
}

func TestStartEventDublicateData(t *testing.T) {
	eventName := fake.Word()

	path := fmt.Sprintf("/events/name/%s", eventName)
	url := apiBaseUrl + path

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	respBody := `{"result": false,"message": "Dublicate data"}`

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	variables := make(map[string]interface{})
	variables["phone"] = fake.Phone()
	variables["name"] = fake.FullName()
	err := spClient.Automation360.StartEvent(eventName, variables)
	assert.Error(t, err)
	spError, isSPError := err.(*SendpulseError)
	assert.True(t, isSPError)
	assert.Equal(t, http.StatusOK, spError.HttpCode)
	assert.Equal(t, path, spError.Url)
	assert.Equal(t, respBody, spError.Body)
	assert.Equal(t, "'result' is false", spError.Message)
}

func TestStartEventWithPhone(t *testing.T) {
	eventName := fake.Word()

	url := fmt.Sprintf("%s/events/name/%s", apiBaseUrl, eventName)

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK,
			`{"result": true}`))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	variables := make(map[string]interface{})
	variables["phone"] = fake.Phone()
	variables["name"] = fake.FullName()
	err := spClient.Automation360.StartEvent(eventName, variables)
	assert.NoError(t, err)
}

func TestStartEventWithEmail(t *testing.T) {
	eventName := fake.Word()

	url := fmt.Sprintf("%s/events/name/%s", apiBaseUrl, eventName)

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK,
			`{"result": true}`))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	variables := make(map[string]interface{})
	variables["email"] = fake.EmailAddress()
	variables["name"] = fake.FullName()
	err := spClient.Automation360.StartEvent(eventName, variables)
	assert.NoError(t, err)
}
