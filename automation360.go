package sendpulse

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type automation360 struct {
	Client *client
}

func (a *automation360) StartEvent(eventName string, variables map[string]interface{}) error {
	path := fmt.Sprintf("/events/name/%s", eventName)

	_, emailExists := variables["email"]
	_, phoneExists := variables["phone"]

	if !emailExists && !phoneExists {
		return errors.New("email and phone are empty")
	}

	body, err := a.Client.makeRequest(path, "POST", variables, true)

	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return &SendpulseError{http.StatusOK, path, string(body), err.Error()}
	}

	result, resultExists := respData["result"]
	if !resultExists {
		return &SendpulseError{http.StatusOK, path, string(body), "'result' not found in response"}
	}

	if !result.(bool) {
		return &SendpulseError{http.StatusOK, path, string(body), "'result' is false"}
	}

	return nil
}
