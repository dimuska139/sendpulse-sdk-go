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

func TestBooks_List_Success(t *testing.T) {
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
		{
			ID:               2,
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

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/addressbooks", apiBaseUrl),
		httpmock.NewStringResponder(http.StatusOK, string(encoded)))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	responseBooks, err := spClient.Emails.Books.List(0, 10)
	assert.NoError(t, err)

	assert.Equal(t, books, responseBooks)
}

func TestBooks_List_BadJson(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/addressbooks"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := "Invalid json"

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	config := Config{
		UserID:  apiUid,
		Secret:  apiSecret,
		Timeout: 5,
	}
	spClient, _ := ApiClient(config)
	spClient.client.token = fake.Word()

	_, err := spClient.Emails.Books.List(0, 10)
	assert.Error(t, err)

	_, isSpError := err.(*SendpulseError)
	assert.True(t, isSpError)
}

func TestBooks_List_Error(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/addressbooks"
	url := apiBaseUrl + path

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

	_, err := spClient.Emails.Books.List(0, 10)
	assert.Error(t, err)

	_, isSpError := err.(*SendpulseError)
	assert.True(t, isSpError)
}
