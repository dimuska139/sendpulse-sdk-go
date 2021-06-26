package models

type EmailInfoList struct {
	ListName string       `json:"list_name"`
	ListID   int          `json:"list_id"`
	AddDate  DateTimeType `json:"add_date"`
	Source   string       `json:"source"`
}
