package sendpulse

import (
	b64 "encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

type BlacklistService struct {
	client *Client
}

func newBlacklistService(cl *Client) *BlacklistService {
	return &BlacklistService{client: cl}
}

func (service *BlacklistService) Add(emails []string, comment string) error {
	path := "/blacklist"

	type paramsFormat struct {
		Emails  string `json:"emails"`
		Comment string `json:"comment,omitempty"`
	}

	params := paramsFormat{
		Emails: b64.StdEncoding.EncodeToString([]byte(strings.Join(emails, ","))),
	}

	if comment != "" {
		params.Comment = comment
	}

	type response struct {
		Result bool
	}

	var respData response
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &respData, true)
	return err
}

func (service *BlacklistService) Remove(emails []string) error {
	path := "/blacklist"

	type paramsFormat struct {
		Emails string `json:"emails"`
	}

	params := paramsFormat{
		Emails: b64.StdEncoding.EncodeToString([]byte(strings.Join(emails, ","))),
	}

	type response struct {
		Result bool
	}

	var respData response
	_, err := service.client.NewRequest(http.MethodDelete, fmt.Sprintf(path), params, &respData, true)
	return err
}

func (service *BlacklistService) List() ([]string, error) {
	path := "/blacklist"

	var respData []string
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}
