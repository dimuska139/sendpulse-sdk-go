package common

import "github.com/dimuska139/sendpulse-sdk-go/types"

type Balance struct {
	Currency        *string  `json:"currency,omitempty"`
	BalanceCurrency *float32 `json:"balance_currency,omitempty"`
}

type BalanceDetailed struct {
	Balance *struct {
		Main     *types.Float32 `json:"main,omitempty"`
		Bonus    *types.Float32 `json:"bonus,omitempty"`
		Currency *string        `json:"currency,omitempty"`
	} `json:"balance,omitempty"`
	Email *struct {
		TariffName         *string         `json:"tariff_name,omitempty"`
		FinishedTime       *types.DateTime `json:"finished_time,omitempty"`
		EmailsLeft         *int            `json:"emails_left,omitempty"`
		MeximumSubscribers *int            `json:"maximum_subscribers,omitempty"`
		CurrentSubscribers *int            `json:"current_subscribers,omitempty"`
	} `json:"email,omitempty"`
	Smtp *struct {
		TariffName *string         `json:"tariff_name,omitempty"`
		EndDate    *types.DateTime `json:"end_date,omitempty"`
		AutoRenew  *int            `json:"auto_renew,omitempty"`
	} `json:"smtp,omitempty"`
	Push *struct {
		TariffName *string     `json:"tariff_name,omitempty"`
		EndDate    *types.Date `json:"end_date,omitempty"`
		AutoRenew  *int        `json:"auto_renew,omitempty"`
	} `json:"push,omitempty"`
}
