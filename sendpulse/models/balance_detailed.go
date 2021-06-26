package models

type BalanceDetailed struct {
	Balance *struct {
		Main     float32 `json:"main,string"`
		Bonus    float32 `json:"bonus,string"`
		Currency string  `json:"currency"`
	} `json:"balance"`
	Email *struct {
		TariffName         string       `json:"tariff_name"`
		FinishedTime       DateTimeType `json:"finished_time"`
		EmailsLeft         int          `json:"emails_left"`
		MaximumSubscribers int          `json:"maximum_subscribers"`
		CurrentSubscribers int          `json:"current_subscribers"`
	} `json:"email"`
	Smtp *struct {
		TariffName string       `json:"tariff_name"`
		EndDate    DateTimeType `json:"end_date"`
		AutoRenew  int          `json:"auto_renew"`
	} `json:"smtp"`
	Push *struct {
		TariffName string       `json:"tariff_name"`
		EndDate    DateTimeType `json:"end_date"`
		AutoRenew  int          `json:"auto_renew"`
	} `json:"push"`
}
