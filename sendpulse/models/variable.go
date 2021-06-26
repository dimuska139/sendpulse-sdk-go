package models

type Variable struct {
	Name  string      `json:"name"`
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value"`
}
