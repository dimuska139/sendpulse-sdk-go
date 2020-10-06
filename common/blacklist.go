package common

import (
	b64 "encoding/base64"
	"encoding/json"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"net/http"
	"strings"
)

func (api *Common) GetBlacklist() ([]*string, error) {
	path := "/blacklist"

	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var emails []*string
	if err := json.Unmarshal(body, &emails); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return emails, nil
}

func (api *Common) AddEmailsToBlacklist(emails []string, comment string) error {
	path := "/blacklist"

	data := map[string]interface{}{
		"emails": b64.StdEncoding.EncodeToString([]byte(strings.Join(emails, ","))),
	}

	if comment != "" {
		data["comment"] = comment
	}

	body, err := api.Client.NewRequest(path, http.MethodPost, nil, true)
	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	result, resultExists := respData["result"]
	if !resultExists || !result.(bool) {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return nil
}

func (api *Common) DeleteEmailsFromBlacklist(emails []string) error {
	path := "/blacklist"

	data := map[string]interface{}{
		"emails": b64.StdEncoding.EncodeToString([]byte(strings.Join(emails, ","))),
	}

	body, err := api.Client.NewRequest(path, http.MethodDelete, data, true)
	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	result, resultExists := respData["result"]
	if !resultExists || !result.(bool) {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return nil
}
