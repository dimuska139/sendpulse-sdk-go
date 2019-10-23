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

func TestBooks_Create_EmptyName(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	_, err := spClient.Emails.Books.Create("")
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.False(t, isResponseError)
}

func TestBooks_Create_Existing(t *testing.T) {
	bookName := fake.Word()
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/addressbooks"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	respBody := `{
    	"error_code": 203,
    	"message": "Book name already in use"
	}`

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusBadRequest, respBody))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	bookId, err := spClient.Emails.Books.Create(bookName)
	assert.Error(t, err)
	spErr, isSPError := err.(*SendpulseError)
	assert.True(t, isSPError)
	assert.Nil(t, bookId)

	assert.Equal(t, http.StatusBadRequest, spErr.HttpCode)
	assert.Equal(t, path, spErr.Url)
	assert.Equal(t, respBody, spErr.Body)
	assert.Equal(t, "", spErr.Message)
}

func TestBooks_Create_IncorrectJson(t *testing.T) {
	bookName := fake.Word()
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/addressbooks"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	respBody := `Incorrect json`

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	bookId, err := spClient.Emails.Books.Create(bookName)
	assert.Error(t, err)
	spErr, isSpError := err.(*SendpulseError)
	assert.True(t, isSpError)
	assert.Nil(t, bookId)

	assert.Equal(t, http.StatusOK, spErr.HttpCode)
	assert.Equal(t, path, spErr.Url)
	assert.Equal(t, respBody, spErr.Body)
}

func TestBooks_Create_NoIdInResponse(t *testing.T) {
	bookName := fake.Word()
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/addressbooks"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	respBody := `{
    	"no_id": "Error"
	}`

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusBadRequest, respBody))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	bookId, err := spClient.Emails.Books.Create(bookName)
	assert.Error(t, err)
	httpErr, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Nil(t, bookId)

	assert.Equal(t, http.StatusBadRequest, httpErr.HttpCode)
	assert.Equal(t, path, httpErr.Url)
	assert.Equal(t, respBody, httpErr.Body)
	assert.Equal(t, "", httpErr.Message)
}

func TestBooks_Create_Success(t *testing.T) {
	bookName := fake.Word()
	var newBookId uint = 1
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/addressbooks"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	respBody := fmt.Sprintf(`{
    	"id": %d
	}`, newBookId)

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	bookId, err := spClient.Emails.Books.Create(bookName)
	assert.NoError(t, err)
	assert.Equal(t, newBookId, *bookId)
}

func TestBooks_Get_Success(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	books := []Book{
		Book{
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

	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/addressbooks/%d", apiBaseUrl, books[0].ID),
		httpmock.NewStringResponder(http.StatusOK, string(encoded)))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)
	book, err := spClient.Emails.Books.Get(books[0].ID)
	assert.NoError(t, err)

	assert.Equal(t, books[0], *book)
}

func TestBooks_Get_NotFound(t *testing.T) {
	respBody := `{
     		"error_code": 213,
   			"message": "Book not found"
		}`

	notExistingBookID := 1

	path := fmt.Sprintf("/addressbooks/%d", notExistingBookID)
	url := apiBaseUrl + path

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusBadRequest, respBody))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	_, err := spClient.Emails.Books.Get(uint(notExistingBookID))
	assert.Error(t, err)
	httpErr, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Equal(t, http.StatusBadRequest, httpErr.HttpCode)
	assert.Equal(t, path, httpErr.Url)
	assert.Equal(t, respBody, httpErr.Body)
	assert.Equal(t, "", httpErr.Message)
}

func TestBooks_Get_InvalidJson(t *testing.T) {
	respBody := `Invalid json`

	notExistingBookID := 1

	path := fmt.Sprintf("/addressbooks/%d", notExistingBookID)
	url := apiBaseUrl + path

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusBadRequest, respBody))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	_, err := spClient.Emails.Books.Get(uint(notExistingBookID))
	assert.Error(t, err)
	httpErr, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Equal(t, http.StatusBadRequest, httpErr.HttpCode)
	assert.Equal(t, path, httpErr.Url)
	assert.Equal(t, respBody, httpErr.Body)
	assert.Equal(t, "", httpErr.Message)
}

func TestBooks_List_Success(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	books := []Book{
		Book{
			ID:               1,
			Name:             fake.CharactersN(10),
			AllEmailQty:      1,
			ActiveEmailQty:   0,
			InactiveEmailQty: 10,
			CreationDate:     "2018-12-28 10:13:51",
			Status:           0,
			StatusExplain:    "Active",
		},
		Book{
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
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	responseBooks, err := spClient.Emails.Books.List(0, 10)
	assert.NoError(t, err)

	assert.Equal(t, books, *responseBooks)
}

func TestBooks_List_InvalidJson(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	path := "/addressbooks"
	url := apiBaseUrl + path

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	respBody := "Invalid json"

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	_, err := spClient.Emails.Books.List(0, 10)
	assert.Error(t, err)

	spError, isSpError := err.(*SendpulseError)
	assert.True(t, isSpError)
	assert.Equal(t, http.StatusOK, spError.HttpCode)
	assert.Equal(t, path, spError.Url)
	assert.Equal(t, respBody, spError.Body)
}

func TestBooks_AddEmails_EmptyList(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	emailsList := make([]Email, 0)
	params := make(map[string]string)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))
	spClient, _ := ApiClient(apiUid, apiSecret, 5)
	err := spClient.Emails.Books.AddEmails(1, emailsList, params, "")
	assert.Error(t, err)
	_, isResponseError := err.(*SendpulseError)
	assert.False(t, isResponseError)
}

func TestBooks_AddEmails_BookNotFound(t *testing.T) {
	addressBookId := 1

	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookId)
	url := apiBaseUrl + path

	respBody := `{
   			"error_code": 404,
    		"message": "Not Found"
		}`

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	emails := []Email{
		Email{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
		Email{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
	}

	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusBadRequest, respBody))
	spClient, _ := ApiClient(apiUid, apiSecret, 5)
	err := spClient.Emails.Books.AddEmails(uint(addressBookId), emails, make(map[string]string), "")
	assert.Error(t, err)

	httpErr, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Equal(t, http.StatusBadRequest, httpErr.HttpCode)
	assert.Equal(t, path, httpErr.Url)
	assert.Equal(t, respBody, httpErr.Body)
}

func TestBooks_AddEmails_InvalidJson(t *testing.T) {
	addressBookId := 1

	path := fmt.Sprintf("/addressbooks/%d/emails", addressBookId)
	url := apiBaseUrl + path
	respBody := `Invalid json`

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	emails := []Email{
		Email{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
		Email{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
	}

	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusBadRequest, respBody))
	spClient, _ := ApiClient(apiUid, apiSecret, 5)
	err := spClient.Emails.Books.AddEmails(uint(addressBookId), emails, make(map[string]string), "")
	assert.Error(t, err)

	spErr, isResponseError := err.(*SendpulseError)
	assert.True(t, isResponseError)
	assert.Equal(t, http.StatusBadRequest, spErr.HttpCode)
	assert.Equal(t, path, spErr.Url)
	assert.Equal(t, respBody, spErr.Body)
}

func TestBooks_AddEmails_Success(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	emails := []Email{
		{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
		{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
	}

	addressBookId := 1

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/addressbooks/%d/emails", apiBaseUrl, addressBookId),
		httpmock.NewStringResponder(http.StatusOK, `{
    		"result": true
		}`))
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))
	spClient, _ := ApiClient(apiUid, apiSecret, 5)
	err := spClient.Emails.Books.AddEmails(uint(addressBookId), emails, make(map[string]string), "")
	assert.NoError(t, err)
}

func TestBooks_AddEmails_WithParamsSuccess(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	emails := []Email{
		{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
		{
			Email:     fake.EmailAddress(),
			Variables: make(map[string]interface{}),
		},
	}

	addressBookId := 1

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/addressbooks/%d/emails", apiBaseUrl, addressBookId),
		httpmock.NewStringResponder(http.StatusOK, `{
    		"result": true
		}`))
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))
	spClient, _ := ApiClient(apiUid, apiSecret, 5)
	extraParams := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}
	err := spClient.Emails.Books.AddEmails(uint(addressBookId), emails, extraParams, fake.EmailAddress())
	assert.NoError(t, err)
}
