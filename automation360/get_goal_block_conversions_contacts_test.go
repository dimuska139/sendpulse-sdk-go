package automation360

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"github.com/icrowley/fake"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAutomation360_GetGoalBlockConversionsContacts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/conversionsContacts.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d/conversions/list/goal", client.ApiBaseUrl, autoresponderID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	contacts, total, err := spClient.GetGoalBlockConversionsContacts(autoresponderID, 0)
	assert.NoError(t, err)
	assert.Equal(t, 40941, int(*contacts[0].ID))
	assert.Equal(t, 5, *total)
}

func TestAutomation360_GetGoalBlockConversionsContactsWithGoalID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	goalID := 2
	responseBody, _ := ioutil.ReadFile("./testdata/conversionsContacts.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d/conversions/list/goal/%d", client.ApiBaseUrl, autoresponderID, goalID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	contacts, total, err := spClient.GetGoalBlockConversionsContacts(autoresponderID, goalID)
	assert.NoError(t, err)
	assert.Equal(t, 40941, int(*contacts[0].ID))
	assert.Equal(t, 5, *total)
}

func TestAutomation360_GetGoalBlockConversionsContacts_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d/conversions/list/goal", client.ApiBaseUrl, autoresponderID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	contacts, total, err := spClient.GetGoalBlockConversionsContacts(autoresponderID, 0)
	assert.Error(t, err)
	assert.Equal(t, 0, len(contacts))
	assert.Nil(t, total)
}

func TestAutomation360_GetGoalBlockConversionsContacts_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d/conversions/list/goal", client.ApiBaseUrl, autoresponderID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	contacts, total, err := spClient.GetGoalBlockConversionsContacts(autoresponderID, 0)
	assert.Error(t, err)
	assert.Equal(t, 0, len(contacts))
	assert.Nil(t, total)
}

func TestAutomation360_GetGoalBlockConversionsContacts_NoData(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d/conversions/list/goal", client.ApiBaseUrl, autoresponderID),
		httpmock.NewStringResponder(http.StatusOK, "{}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	contacts, total, err := spClient.GetGoalBlockConversionsContacts(autoresponderID, 0)
	assert.Error(t, err)
	assert.Equal(t, 0, len(contacts))
	assert.Nil(t, total)
}
