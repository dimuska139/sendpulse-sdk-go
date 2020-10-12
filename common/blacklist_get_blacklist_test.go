package common

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

func TestCommon_GetBlacklist(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responseBody, _ := ioutil.ReadFile("./testdata/blacklist.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/blacklist", client.ApiBaseUrl),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	blacklist, err := spClient.GetBlacklist()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(blacklist))
	assert.Equal(t, "JohnDoe@test.com", *blacklist[0])
}

func TestCommon_GetBlacklist_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/blacklist", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	blacklist, err := spClient.GetBlacklist()
	assert.Error(t, err)
	assert.Equal(t, 0, len(blacklist))
}

func TestCommon_GetBlacklist_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/blacklist", client.ApiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	blacklist, err := spClient.GetBlacklist()
	assert.Error(t, err)
	assert.Equal(t, 0, len(blacklist))
}
