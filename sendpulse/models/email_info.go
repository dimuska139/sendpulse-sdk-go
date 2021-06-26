package models

type EmailInfo struct {
	BookID    int         `json:"book_id"`
	Status    int         `json:"status"`
	Variables []*Variable `json:"variables"`
}
