package sendpulse

import (
	"encoding/json"
	"errors"
	"fmt"
)

type automation360 struct {
	Client *client
}

func (a *automation360) StartEvent(eventName string, variables map[string]string) error {
	if len(eventName) == 0 {
		return errors.New("event name is empty")
	}

	_, emailExists := variables["email"]
	_, phoneExists := variables["phone"]

	if !emailExists && !phoneExists {
		return errors.New("email and phone are empty")
	}

	body, err := a.Client.makeRequest(fmt.Sprintf("/events/name/%s", eventName), "POST", variables, true)

	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return errors.New(string(body))
	}
	_, resultExists := respData["result"]
	if !resultExists || !respData["result"].(bool) {
		return errors.New(string(body))
	}

	return nil
}
