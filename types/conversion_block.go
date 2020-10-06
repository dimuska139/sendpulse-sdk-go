package types

import (
	"encoding/json"
	"errors"
)

type block struct {
	ID          *json.Number `json:"id,omitempty"`
	Name        *string      `json:"name,omitempty"`
	MainID      *json.Number `json:"main_id,omitempty"`
	AfType      *string      `json:"af_type,omitempty"`
	Created     *DateTime    `json:"created,omitempty"`
	LastSend    *DateTime    `json:"last_send,omitempty"`
	Conversions *int         `json:"conversions,omitempty"`
}

type ConversionBlock struct {
	block
}

func (d *ConversionBlock) UnmarshalJSON(data []byte) error {
	if string(data) == "[]" {
		return nil
	}
	var bl block

	err := json.Unmarshal(data, &bl)
	if err != nil {
		return errors.New("ConversionBlock: UnmarshalJSON:  " + err.Error())
	}

	*d = ConversionBlock{
		bl,
	}
	return nil
}
