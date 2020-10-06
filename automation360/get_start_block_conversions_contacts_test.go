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

func TestAutomation360_GetStartBlockConversionsContacts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/conversionsContacts.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d/conversions/list/maintrigger", client.ApiBaseUrl, autoresponderID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	contacts, total, err := spClient.GetStartBlockConversionsContacts(autoresponderID)
	assert.NoError(t, err)
	assert.Equal(t, 40941, int(*contacts[0].ID))
	assert.Equal(t, 5, *total)
}

func TestAutomation360_GetStartBlockConversionsContacts_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d/conversions/list/all", client.ApiBaseUrl, autoresponderID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	contacts, total, err := spClient.GetStartBlockConversionsContacts(autoresponderID)
	assert.Error(t, err)
	assert.Equal(t, 0, len(contacts))
	assert.Nil(t, total)
}

func TestAutomation360_GetStartBlockConversionsContacts_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d/conversions/list/maintrigger", client.ApiBaseUrl, autoresponderID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	contacts, total, err := spClient.GetStartBlockConversionsContacts(autoresponderID)
	assert.Error(t, err)
	assert.Equal(t, 0, len(contacts))
	assert.Nil(t, total)
}

func TestAutomation360_GetStartBlockConversionsContacts_NoData(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	autoresponderID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/autoresponders/%d/conversions/list/maintrigger", client.ApiBaseUrl, autoresponderID),
		httpmock.NewStringResponder(http.StatusOK, "{}"),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	contacts, total, err := spClient.GetStartBlockConversionsContacts(autoresponderID)
	assert.Error(t, err)
	assert.Equal(t, 0, len(contacts))
	assert.Nil(t, total)
}
