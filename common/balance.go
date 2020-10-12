package common

import (
	"encoding/json"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"net/http"
	"strings"
)

func (api *Common) GetBalance(currency string) (*Balance, error) {
	path := "/balance"
	if currency != "" {
		path += "/" + strings.ToLower(currency)
	}

	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var balance Balance
	if err := json.Unmarshal(body, &balance); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &balance, err
}

func (api *Common) GetBalanceDetailed() (*BalanceDetailed, error) {
	path := "/user/balance/detail"
	body, err := api.Client.NewRequest(path, http.MethodGet, nil, true)

	if err != nil {
		return nil, err
	}

	var balance BalanceDetailed
	if err := json.Unmarshal(body, &balance); err != nil {
		return nil, &client.SendpulseError{HttpCode: http.StatusOK, Url: path, Body: string(body), Message: err.Error()}
	}

	return &balance, err
}
