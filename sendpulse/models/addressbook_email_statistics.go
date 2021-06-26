package models

type AddressBookEmailStatistics struct {
	Email         string      `json:"email"`
	AddressBookID int         `json:"abook_id,string"`
	Status        int         `json:"status"`
	StatusExplain string      `json:"status_explain"`
	Variables     []*Variable `json:"variables"`
}
