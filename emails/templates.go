package emails

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"net/http"
)

func (api *Emails) CreateTemplate(name string, body string, lang string) (*int, error) {
	path := "/template"

	data := map[string]interface{}{
		"body": b64.StdEncoding.EncodeToString([]byte(body)),
		"lang": lang,
	}

	if name != "" {
		data["name"] = name
	}

	response, err := api.Client.NewRequest(fmt.Sprintf(path), http.MethodPost, data, true)
	if err != nil {
		return nil, err
	}

	type tplResp struct {
		Result bool
		RealID int `json:"real_id"`
	}

	var respData tplResp
	if err := json.Unmarshal(response, &respData); err != nil {
		return nil, &client.SendpulseError{
			HttpCode: http.StatusOK,
			Url:      path,
			Body:     body,
			Message:  err.Error(),
		}
	}

	if !respData.Result {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return &respData.RealID, err
}

func (api *Emails) UpdateTemplate(templateID int, body string, lang string) error {
	path := fmt.Sprintf("/template/edit/%d", templateID)

	data := map[string]interface{}{
		"body": b64.StdEncoding.EncodeToString([]byte(body)),
		"lang": lang,
	}

	response, err := api.Client.NewRequest(fmt.Sprintf(path), http.MethodPost, data, true)
	if err != nil {
		return err
	}

	type tplResp struct {
		Result bool
	}

	var respData tplResp
	if err := json.Unmarshal(response, &respData); err != nil {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: body, Message: err.Error()}
	}

	if !respData.Result {
		return &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return nil
}

func (api *Emails) GetTemplate(templateID int) (*Template, error) {
	path := fmt.Sprintf("/template/%d", templateID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var template *Template
	if err := json.Unmarshal(body, &template); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return template, err
}

func (api *Emails) GetTemplates(limit int, offset int, owner string) ([]*Template, error) {
	path := "/templates"
	data := map[string]interface{}{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
	}
	if owner != "" {
		data["owner"] = owner
	}
	body, err := api.Client.NewRequest(path, http.MethodGet, data, true)

	if err != nil {
		return nil, err
	}

	var templates []*Template
	if err := json.Unmarshal(body, &templates); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return templates, err
}
