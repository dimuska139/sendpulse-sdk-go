package types

import (
	"encoding/json"
	"errors"
)

type Float32 float32

func (d *Float32) UnmarshalJSON(b []byte) error {
	var customFloat float32
	if b[0] == 34 {
		if err := json.Unmarshal(b[1:len(b)-1], &customFloat); err != nil {
			return errors.New("SendpulseFloat32: UnmarshalJSON: " + err.Error())
		}
	} else {
		if err := json.Unmarshal(b, &customFloat); err != nil {
			return errors.New("SendpulseFloat32: UnmarshalJSON: " + err.Error())
		}
	}
	*d = Float32(customFloat)
	return nil
}
