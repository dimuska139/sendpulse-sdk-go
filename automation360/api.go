package automation360

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"net/http"
)

func (api *Automation360) StartEvent(eventName string, variables map[string]interface{}) error {
	path := fmt.Sprintf("/events/name/%s", eventName)

	_, emailExists := variables["email"]
	_, phoneExists := variables["phone"]

	if !emailExists && !phoneExists {
		return errors.New("email and phone are empty")
	}

	body, err := api.Client.NewRequest(path, http.MethodPost, variables, true)
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

func (api *Automation360) GetAutoresponder(autoresponderID int) (*Autoresponder, error) {
	path := fmt.Sprintf("/a360/autoresponders/%d", autoresponderID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var autoresponder Autoresponder
	if err := json.Unmarshal(body, &autoresponder); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &autoresponder, err
}

func (api *Automation360) GetStartBlockStatistics(blockID int) (*MainTriggerBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/main-trigger/%d/group-stat", blockID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}
	type response struct {
		Data *MainTriggerBlockStat `json:"data,omitempty"`
	}
	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if resp.Data == nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return resp.Data, err
}

func (api *Automation360) GetStartBlockRecipients(blockID int, limit int, offset int, sortDirection string, sortField string) ([]*MainTriggerBlockRecipient, *int, error) {
	path := fmt.Sprintf("/a360/stats/main-trigger/%d/addresses", blockID)
	data := map[string]interface{}{
		"limit":         fmt.Sprint(limit),
		"offset":        fmt.Sprint(offset),
		"sortDirection": sortDirection,
		"sortField":     sortField,
	}
	body, err := api.Client.NewRequest(path, http.MethodGet, data, true)

	if err != nil {
		return nil, nil, err
	}

	type response struct {
		Data  *[]*MainTriggerBlockRecipient `json:"data,omitempty"`
		Total *int                          `json:"total,omitempty"`
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if resp.Data == nil || resp.Total == nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return *resp.Data, resp.Total, nil
}

func (api *Automation360) GetEmailBlockStatistics(blockID int) (*EmailBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/email/%d/group-stat", blockID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var stat EmailBlockStat
	if err := json.Unmarshal(body, &stat); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &stat, err
}

func (api *Automation360) GetEmailBlockRecipients(blockID int, limit int, offset int, sortDirection string, sortField string) ([]*EmailBlockRecipient, *int, error) {
	path := fmt.Sprintf("/a360/stats/email/%d/addresses", blockID)
	data := map[string]interface{}{
		"limit":         fmt.Sprint(limit),
		"offset":        fmt.Sprint(offset),
		"sortDirection": sortDirection,
		"sortField":     sortField,
	}
	body, err := api.Client.NewRequest(path, http.MethodGet, data, true)

	if err != nil {
		return nil, nil, err
	}

	type response struct {
		Data  *[]*EmailBlockRecipient `json:"data,omitempty"`
		Total *int                    `json:"total,omitempty"`
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if resp.Data == nil || resp.Total == nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return *resp.Data, resp.Total, nil
}

func (api *Automation360) GetPushBlockStatistics(blockID int) (*PushBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/push/%d/group-stat", blockID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var stat PushBlockStat
	if err := json.Unmarshal(body, &stat); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &stat, err
}

func (api *Automation360) GetPushBlockRecipients(blockID int, limit int, offset int, sortDirection string, sortField string) ([]*PushBlockRecipient, *int, error) {
	path := fmt.Sprintf("/a360/stats/push/%d/addresses", blockID)
	data := map[string]interface{}{
		"limit":         fmt.Sprint(limit),
		"offset":        fmt.Sprint(offset),
		"sortDirection": sortDirection,
		"sortField":     sortField,
	}
	body, err := api.Client.NewRequest(path, http.MethodGet, data, true)

	if err != nil {
		return nil, nil, err
	}

	type response struct {
		Data  *[]*PushBlockRecipient `json:"data,omitempty"`
		Total *int                   `json:"total,omitempty"`
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if resp.Data == nil || resp.Total == nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return *resp.Data, resp.Total, nil
}

func (api *Automation360) GetSmsBlockStatistics(blockID int) (*SmsBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/sms/%d/group-stat", blockID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var stat SmsBlockStat
	if err := json.Unmarshal(body, &stat); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &stat, err
}

func (api *Automation360) GetSmsBlockRecipients(blockID int, limit int, offset int, sortDirection string, sortField string) ([]*SmsBlockRecipient, *int, error) {
	path := fmt.Sprintf("/a360/stats/sms/%d/addresses", blockID)
	data := map[string]interface{}{
		"limit":         fmt.Sprint(limit),
		"offset":        fmt.Sprint(offset),
		"sortDirection": sortDirection,
		"sortField":     sortField,
	}
	body, err := api.Client.NewRequest(path, http.MethodGet, data, true)

	if err != nil {
		return nil, nil, err
	}

	type response struct {
		Data  *[]*SmsBlockRecipient `json:"data,omitempty"`
		Total *int                  `json:"total,omitempty"`
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if resp.Data == nil || resp.Total == nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return *resp.Data, resp.Total, nil
}

func (api *Automation360) GetFilterBlockStatistics(blockID int) (*FilterBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/filter/%d/group-stat", blockID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var stat FilterBlockStat
	if err := json.Unmarshal(body, &stat); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &stat, err
}

func (api *Automation360) GetFilterBlockRecipients(blockID int, limit int, offset int, sortDirection string, sortField string) ([]*FilterBlockRecipient, *int, error) {
	path := fmt.Sprintf("/a360/stats/flow-operator/%d/addresses", blockID)
	data := map[string]interface{}{
		"limit":         fmt.Sprint(limit),
		"offset":        fmt.Sprint(offset),
		"sortDirection": sortDirection,
		"sortField":     sortField,
	}
	body, err := api.Client.NewRequest(path, http.MethodGet, data, true)

	if err != nil {
		return nil, nil, err
	}

	type response struct {
		Data  *[]*FilterBlockRecipient `json:"data,omitempty"`
		Total *int                     `json:"total,omitempty"`
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if resp.Data == nil || resp.Total == nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return *resp.Data, resp.Total, nil
}

func (api *Automation360) GetConditionBlockStatistics(blockID int) (*ConditionBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/trigger/%d/group-stat", blockID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var stat ConditionBlockStat
	if err := json.Unmarshal(body, &stat); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &stat, err
}

func (api *Automation360) GetConditionBlockRecipients(blockID int, limit int, offset int, sortDirection string, sortField string) ([]*ConditionBlockRecipient, *int, error) {
	path := fmt.Sprintf("/a360/stats/flow-operator/%d/addresses", blockID)
	data := map[string]interface{}{
		"limit":         fmt.Sprint(limit),
		"offset":        fmt.Sprint(offset),
		"sortDirection": sortDirection,
		"sortField":     sortField,
	}
	body, err := api.Client.NewRequest(path, http.MethodGet, data, true)

	if err != nil {
		return nil, nil, err
	}

	type response struct {
		Data  *[]*ConditionBlockRecipient `json:"data,omitempty"`
		Total *int                        `json:"total,omitempty"`
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if resp.Data == nil || resp.Total == nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return *resp.Data, resp.Total, nil
}

func (api *Automation360) GetGoalBlockStatistics(blockID int) (*SmsBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/goal/%d/group-stat", blockID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var stat SmsBlockStat
	if err := json.Unmarshal(body, &stat); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &stat, err
}

func (api *Automation360) GetGoalBlockRecipients(blockID int, limit int, offset int, sortDirection string, sortField string) ([]*GoalBlockRecipient, *int, error) {
	path := fmt.Sprintf("/a360/stats/flow-operator/%d/addresses", blockID)
	data := map[string]interface{}{
		"limit":         fmt.Sprint(limit),
		"offset":        fmt.Sprint(offset),
		"sortDirection": sortDirection,
		"sortField":     sortField,
	}
	body, err := api.Client.NewRequest(path, http.MethodGet, data, true)

	if err != nil {
		return nil, nil, err
	}

	type response struct {
		Data  *[]*GoalBlockRecipient `json:"data,omitempty"`
		Total *int                   `json:"total,omitempty"`
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if resp.Data == nil || resp.Total == nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return *resp.Data, resp.Total, nil
}

func (api *Automation360) GetActionBlockStatistics(blockID int) (*ActionBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/action/%d/group-stat", blockID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var stat ActionBlockStat
	if err := json.Unmarshal(body, &stat); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &stat, err
}

func (api *Automation360) GetActionBlockRecipients(blockID int, limit int, offset int, sortDirection string, sortField string) ([]*ActionBlockRecipient, *int, error) {
	path := fmt.Sprintf("/a360/stats/flow-operator/%d/addresses", blockID)
	data := map[string]interface{}{
		"limit":         fmt.Sprint(limit),
		"offset":        fmt.Sprint(offset),
		"sortDirection": sortDirection,
		"sortField":     sortField,
	}
	body, err := api.Client.NewRequest(path, http.MethodGet, data, true)

	if err != nil {
		return nil, nil, err
	}

	type response struct {
		Data  *[]*ActionBlockRecipient `json:"data,omitempty"`
		Total *int                     `json:"total,omitempty"`
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if resp.Data == nil || resp.Total == nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return *resp.Data, resp.Total, nil
}

func (api *Automation360) GetConversions(autoresponderID int) (*Conversion, error) {
	path := fmt.Sprintf("/a360/autoresponders/%d/conversions", autoresponderID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	type response struct {
		Data *Conversion `json:"data,omitempty"`
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return resp.Data, nil
}

func (api *Automation360) GetConversionsContacts(autoresponderID int) ([]*ConversionContact, *int, error) {
	path := fmt.Sprintf("/a360/autoresponders/%d/conversions/list/all", autoresponderID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, nil, err
	}

	type response struct {
		Items *[]*ConversionContact `json:"items,omitempty"`
		Total *int                  `json:"total,omitempty"`
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if resp.Items == nil || resp.Total == nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return *resp.Items, resp.Total, nil
}

func (api *Automation360) GetStartBlockConversionsContacts(autoresponderID int) ([]*ConversionContact, *int, error) {
	path := fmt.Sprintf("/a360/autoresponders/%d/conversions/list/maintrigger", autoresponderID)
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, nil, err
	}

	type response struct {
		Items *[]*ConversionContact `json:"items,omitempty"`
		Total *int                  `json:"total,omitempty"`
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if resp.Items == nil || resp.Total == nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return *resp.Items, resp.Total, nil
}

func (api *Automation360) GetGoalBlockConversionsContacts(autoresponderID int, goalID int) ([]*ConversionContact, *int, error) {
	path := fmt.Sprintf("/a360/autoresponders/%d/conversions/list/goal", autoresponderID)
	if goalID != 0 {
		path = fmt.Sprintf("/a360/autoresponders/%d/conversions/list/goal/%d", autoresponderID, goalID)
	}
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, nil, err
	}

	type response struct {
		Items *[]*ConversionContact `json:"items,omitempty"`
		Total *int                  `json:"total,omitempty"`
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	if resp.Items == nil || resp.Total == nil {
		return nil, nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: "invalid response"}
	}

	return *resp.Items, resp.Total, nil
}
