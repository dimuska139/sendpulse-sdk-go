package sendpulse

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type templates struct {
	Client *client
}

type CategoryInfo struct {
	ID              int
	Name            string
	MetaDescription string
	FullDescription string
	Code            string
	Sort            int
}

type TemplateInfo struct {
	ID              string
	RealID          int
	Lang            string
	Name            string
	NameSlug        string
	Created         string
	FullDescription string
	Category        string
	CategoryInfo    CategoryInfo
	Tags            interface{} //[]map[string]interface{}
	Owner           string
	Preview         string
}

func (tpl *templates) Create(name string, tplBody string, lang string) (int, error) {
	path := "/template"
	data := map[string]interface{}{
		"body": b64.StdEncoding.EncodeToString([]byte(tplBody)),
		"lang": lang,
	}

	if name != "" {
		data["name"] = name
	}
	body, err := tpl.Client.makeRequest(path, "POST", data, true)
	if err != nil {
		return 0, err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return 0, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	result, resultExists := respData["result"]
	if !resultExists || !result.(bool) {
		return 0, &SendpulseError{http.StatusOK, path, string(body), "invalid response"}
	}

	rawRealID, realIDExists := respData["real_id"]
	if !realIDExists {
		return 0, &SendpulseError{http.StatusOK, path, string(body), "invalid response"}
	}

	realID, _ := strconv.Atoi(fmt.Sprint(rawRealID))
	return realID, nil
}

func (tpl *templates) Update(templateID int, templateRealID string, tplBody string, lang string) error {
	path := fmt.Sprintf("/template/edit/%d", templateID)

	data := map[string]interface{}{
		"id":   templateRealID,
		"body": b64.StdEncoding.EncodeToString([]byte(tplBody)),
		"lang": lang,
	}

	body, err := tpl.Client.makeRequest(fmt.Sprintf(path), "POST", data, true)
	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	result, resultExists := respData["result"]
	if !resultExists || !result.(bool) {
		return &SendpulseError{http.StatusOK, path, string(body), "invalid response"}
	}

	return nil
}

func (tpl *templates) List() ([]TemplateInfo, error) {
	path := "/templates"
	data := map[string]interface{}{}
	body, err := tpl.Client.makeRequest(path, "GET", data, true)

	if err != nil {
		return nil, err
	}

	var respData []TemplateInfo
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	return respData, nil
}
