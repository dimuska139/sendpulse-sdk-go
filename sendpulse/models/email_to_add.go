package models

type EmailToAdd struct {
	Email     string                 `json:"email"`
	Variables map[string]interface{} `json:"variables"`
}
