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

func TestBooks_Get_Success(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	books := []Book{
		{
			ID:               1,
			Name:             fake.CharactersN(10),
			AllEmailQty:      1,
			ActiveEmailQty:   0,
			InactiveEmailQty: 10,
			CreationDate:     "2018-12-28 10:13:51",
			Status:           0,
			StatusExplain:    "Active",
		},
	}
	encoded, _ := json.Marshal(books)

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/addressbooks/%d", apiBaseUrl, books[0].ID),
		httpmock.NewStringResponder(http.StatusOK, string(encoded)))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	book, err := spClient.Emails.Books.Get(books[0].ID)
	assert.NoError(t, err)

	assert.Equal(t, books[0], *book)
}

func TestBooks_Get_BadJson(t *testing.T) {
	respBody := `Invalid json`

	notExistingBookID := 1

	path := fmt.Sprintf("/addressbooks/%d", notExistingBookID)
	url := apiBaseUrl + path

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Books.Get(uint(notExistingBookID))
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}

func TestBooks_Get_Error(t *testing.T) {
	notExistingBookID := 1

	path := fmt.Sprintf("/addressbooks/%d", notExistingBookID)
	url := apiBaseUrl + path

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusBadRequest, ""))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Books.Get(uint(notExistingBookID))
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
}
