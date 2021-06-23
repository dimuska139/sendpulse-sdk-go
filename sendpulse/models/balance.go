package models

type Balance struct {
	Currency        string  `json:"currency"`
	BalanceCurrency float32 `json:"balance_currency"`
}
