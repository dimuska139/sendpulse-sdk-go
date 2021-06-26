package models

type Mailing struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message *struct {
		SenderName    string  `json:"sender_name"`
		SenderEmail   string  `json:"sender_email"`
		Subject       string  `json:"subject"`
		Body          *string `json:"body"`
		Attachments   string  `json:"attachments"`
		AddressBookID int     `json:"list_id"`
	}
	Status            int           `json:"status"`
	AllEmailQty       int           `json:"all_email_qty"`
	TariffEmailQty    int           `json:"tariff_email_qty"`
	PaidEmailQty      int           `json:"paid_email_qty"`
	OverdraftPrice    float32       `json:"overdraft_price"`
	OverdraftCurrency string        `json:"overdraft_currency"`
	SendDate          *DateTimeType `json:"send_date"`
}
