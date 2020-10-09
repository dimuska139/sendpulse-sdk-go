package types

import (
	"encoding/json"
	"errors"
)

type Int int

func (d *Int) UnmarshalJSON(data []byte) error {
	var customInt int
	if data[0] == '"' {
		if err := json.Unmarshal(data[1:len(data)-1], &customInt); err != nil {
			return errors.New("SendpulseInt: UnmarshalJSON: " + err.Error())
		}
	} else {
		if err := json.Unmarshal(data, &customInt); err != nil {
			return errors.New("SendpulseInt: UnmarshalJSON: " + err.Error())
		}
	}
	*d = Int(customInt)
	return nil
}
