package models

type Webhook struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Url    string `json:"url"`
	Action string `json:"action"`
}
