package models

type CampaignCost struct {
	Cur                       string `json:"email"`
	SentEmailsQty             int    `json:"sent_emails_qty"`
	OverdraftAllEmailsPrice   int    `json:"overdraft_all_emails_price"`
	AddressesDeltaFromBalance int    `json:"address_delta_from_balance"`
	AddressesDeltaFromTariff  int    `json:"address_delta_from_tariff"`
	MaxEmailsPerTask          int    `json:"max_emails_per_task"`
	Result                    bool   `json:"result"`
}
