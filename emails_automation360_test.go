package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestAutomation360_StartEvent_NoPhoneAndEmail(t *testing.T) {
	eventName := fake.Word()
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)

	variables := make(map[string]interface{})
	variables["name"] = fake.FullName()
	err := spClient.Emails.Automation360.StartEvent(eventName, variables)
	assert.Error(t, err)
	_, isSPError := err.(*SendpulseError)
	assert.False(t, isSPError)
}

func TestAutomation360_StartEvent_Success(t *testing.T) {
	eventName := fake.Word()

	url := fmt.Sprintf("%s/events/name/%s", apiBaseUrl, eventName)

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK,
			`{"result": true}`))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	variables := make(map[string]interface{})
	variables["phone"] = fake.Phone()
	variables["name"] = fake.FullName()
	err := spClient.Emails.Automation360.StartEvent(eventName, variables)
	assert.NoError(t, err)
}

func TestAutomation360_StartEvent_Error(t *testing.T) {
	eventName := fake.Word()

	url := fmt.Sprintf("%s/events/name/%s", apiBaseUrl, eventName)

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusNotFound, ""))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	variables := make(map[string]interface{})
	variables["phone"] = fake.Phone()
	variables["name"] = fake.FullName()
	err := spClient.Emails.Automation360.StartEvent(eventName, variables)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestAutomation360_StartEvent_BadJson(t *testing.T) {
	eventName := fake.Word()

	url := fmt.Sprintf("%s/events/name/%s", apiBaseUrl, eventName)

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK,
			`{result": true}`))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	variables := make(map[string]interface{})
	variables["email"] = fake.EmailAddress()
	variables["name"] = fake.FullName()
	err := spClient.Emails.Automation360.StartEvent(eventName, variables)
	assert.Error(t, err)

	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestAutomation360_StartEvent_BadResult(t *testing.T) {
	eventName := fake.Word()

	url := fmt.Sprintf("%s/events/name/%s", apiBaseUrl, eventName)

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK,
			`{"result": false}`))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	variables := make(map[string]interface{})
	variables["email"] = fake.EmailAddress()
	variables["name"] = fake.FullName()
	err := spClient.Emails.Automation360.StartEvent(eventName, variables)
	assert.Error(t, err)

	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}
