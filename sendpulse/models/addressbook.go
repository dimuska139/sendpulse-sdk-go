package models

type Book struct {
	ID               int          `json:"id"`
	Name             string       `json:"name"`
	AllEmailQty      int          `json:"all_email_qty"`
	ActiveEmailQty   int          `json:"active_email_qty"`
	InactiveEmailQty int          `json:"inactive_email_qty"`
	CreationDate     DateTimeType `json:"creationdate"`
	Status           int          `json:"status"`
	StatusExplain    string       `json:"status_explain"`
}
