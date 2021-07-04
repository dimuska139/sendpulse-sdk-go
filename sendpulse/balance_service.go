package sendpulse

import (
	"fmt"
	"net/http"
	"strings"
)

type BalanceService struct {
	client *Client
}

func newBalanceService(cl *Client) *BalanceService {
	return &BalanceService{client: cl}
}

type Balance struct {
	Currency        string  `json:"currency"`
	BalanceCurrency float32 `json:"balance_currency"`
}

type BalanceDetailed struct {
	Balance struct {
		Main     float32 `json:"main,string"`
		Bonus    float32 `json:"bonus,string"`
		Currency string  `json:"currency"`
	} `json:"balance"`
	Email struct {
		TariffName         string       `json:"tariff_name"`
		FinishedTime       DateTimeType `json:"finished_time"`
		EmailsLeft         int          `json:"emails_left"`
		MaximumSubscribers int          `json:"maximum_subscribers"`
		CurrentSubscribers int          `json:"current_subscribers"`
	} `json:"email"`
	Smtp struct {
		TariffName string       `json:"tariff_name"`
		EndDate    DateTimeType `json:"end_date"`
		AutoRenew  int          `json:"auto_renew"`
	} `json:"smtp"`
	Push struct {
		TariffName string       `json:"tariff_name"`
		EndDate    DateTimeType `json:"end_date"`
		AutoRenew  int          `json:"auto_renew"`
	} `json:"push"`
}

func (service *BalanceService) GetBalance(currency string) (*Balance, error) {
	path := "/balance"
	if currency != "" {
		path += "/" + strings.ToLower(currency)
	}

	var respData Balance
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return &respData, err
}

func (service *BalanceService) GetDetailedBalance() (*BalanceDetailed, error) {
	path := "/user/balance/detail"

	var respData BalanceDetailed
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return &respData, err
}
