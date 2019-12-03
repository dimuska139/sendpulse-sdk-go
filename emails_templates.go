package sendpulse

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
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

type rawTemplateInfo struct {
	ID              string
	RealID          int
	Lang            string
	Name            string
	NameSlug        string
	Created         string
	FullDescription string
	Category        string
	CategoryInfo    CategoryInfo
	Tags            interface{}
	Owner           string
	Preview         string
}

type TemplateInfo struct {
	ID              string
	RealID          int
	Lang            string
	Name            string
	NameSlug        string
	Created         time.Time
	FullDescription string
	Category        string
	CategoryInfo    CategoryInfo
	Tags            map[string]interface{}
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

func (tpl *templates) Get(templateID interface{}) (*TemplateInfo, error) {
	path := fmt.Sprintf("/template/%v", templateID)
	body, err := tpl.Client.makeRequest(path, "GET", nil, true)

	if err != nil {
		return nil, err
	}

	var respData rawTemplateInfo
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	created, err := time.Parse("2006-01-02 15:04:05", respData.Created)
	if err != nil {
		return nil, err
	}

	_, tagsAreEmpty := respData.Tags.([]interface{})
	tags := make(map[string]interface{})
	if !tagsAreEmpty {
		tags = respData.Tags.(map[string]interface{})
	}

	template := TemplateInfo{
		ID:              respData.ID,
		RealID:          respData.RealID,
		Lang:            respData.Lang,
		Name:            respData.Name,
		NameSlug:        respData.NameSlug,
		Created:         created,
		FullDescription: respData.FullDescription,
		Category:        respData.Category,
		CategoryInfo:    respData.CategoryInfo,
		Tags:            tags,
		Owner:           respData.Owner,
		Preview:         respData.Preview,
	}
	return &template, err
}

func (tpl *templates) Update(templateID interface{}, tplBody string, lang string) error {
	path := fmt.Sprintf("/template/edit/%v", templateID)

	data := map[string]interface{}{
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

	var respData []rawTemplateInfo
	if err := json.Unmarshal(body, &respData); err != nil {
		return nil, &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	var templates []TemplateInfo
	for _, raw := range respData {
		created, err := time.Parse("2006-01-02 15:04:05", raw.Created)
		if err != nil {
			return nil, err
		}

		_, tagsAreEmpty := raw.Tags.([]interface{})
		tags := make(map[string]interface{})
		if !tagsAreEmpty {
			tags = raw.Tags.(map[string]interface{})
		}

		templates = append(templates, TemplateInfo{
			ID:              raw.ID,
			RealID:          raw.RealID,
			Lang:            raw.Lang,
			Name:            raw.Name,
			NameSlug:        raw.NameSlug,
			Created:         created,
			FullDescription: raw.FullDescription,
			Category:        raw.Category,
			CategoryInfo:    raw.CategoryInfo,
			Tags:            tags,
			Owner:           raw.Owner,
			Preview:         raw.Preview,
		})
	}

	return templates, nil
}
