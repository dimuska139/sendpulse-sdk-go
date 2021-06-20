package models

type Variable struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
