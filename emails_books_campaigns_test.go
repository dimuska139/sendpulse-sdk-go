package sendpulse

import (
	"encoding/json"
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestBooks_Campaigns_Success(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data := []Task{
		{
			ID:     1,
			Name:   fake.FullName(),
			Status: 3,
		},
		{
			ID:     2,
			Name:   fake.FullName(),
			Status: 3,
		},
	}
	encoded, _ := json.Marshal(data)

	bookID := 1
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/addressbooks/%d/campaigns?limit=100&offset=0", apiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusOK, string(encoded)))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	campaignsList, err := spClient.Emails.Books.Campaigns(1, 100, 0)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(campaignsList))
}

func TestBooks_Campaigns_BadJson(t *testing.T) {
	bookID := 1

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/addressbooks/%d/campaigns?limit=100&offset=0", apiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusOK, `Invalid json`))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Books.Campaigns(bookID, 100, 0)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestBooks_Campaigns_Error(t *testing.T) {
	bookID := 1

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/addressbooks/%d/campaigns?limit=100&offset=0", apiBaseUrl, bookID),
		httpmock.NewStringResponder(http.StatusBadRequest, ""))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Books.Campaigns(bookID, 100, 0)
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}
