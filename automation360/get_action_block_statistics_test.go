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

func TestAutomation360_GetActionBlockStatistics(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	blockID := 1
	responseBody, _ := ioutil.ReadFile("./testdata/actionBlockStatistics.json")
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/stats/action/%d/group-stat", client.ApiBaseUrl, blockID),
		httpmock.NewBytesResponder(http.StatusOK, responseBody),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	statistics, err := spClient.GetActionBlockStatistics(blockID)
	assert.NoError(t, err)
	assert.NotNil(t, statistics)
}

func TestAutomation360_GetActionBlockStatistics_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	blockID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/stats/action/%d/group-stat", client.ApiBaseUrl, blockID),
		httpmock.NewStringResponder(http.StatusInternalServerError, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	statistics, err := spClient.GetActionBlockStatistics(blockID)
	assert.Error(t, err)
	assert.Nil(t, statistics)
}

func TestAutomation360_GetActionBlockStatistics_InvalidJson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	blockID := 1
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s/a360/stats/action/%d/group-stat", client.ApiBaseUrl, blockID),
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	config := sendpulse.Config{
		UserID: fake.CharactersN(50),
		Secret: fake.CharactersN(50),
		Token:  fake.Word(),
	}

	spClient := New(http.DefaultClient, &config)

	statistics, err := spClient.GetActionBlockStatistics(blockID)
	assert.Error(t, err)
	assert.Nil(t, statistics)
}
