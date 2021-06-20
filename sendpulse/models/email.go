package models

type Email struct {
	Email         string                 `json:"email"`
	Phone         int                    `json:"phone"`
	Status        int                    `json:"status"`
	StatusExplain string                 `json:"status_explain"`
	Variables     map[string]interface{} `json:"variables"`
}
