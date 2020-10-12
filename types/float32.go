package types

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Float32 float32

func (d *Float32) UnmarshalJSON(data []byte) error {
	var customFloat float32
	if data[0] == '"' {
		f64, err := strconv.ParseFloat(string(data[1:len(data)-1]), 32)
		if err != nil {
			return errors.New("SendpulseFloat32: UnmarshalJSON: " + err.Error())
		}
		customFloat = float32(f64)
	} else {
		if err := json.Unmarshal(data, &customFloat); err != nil {
			return errors.New("SendpulseFloat32: UnmarshalJSON: " + err.Error())
		}
	}
	*d = Float32(customFloat)
	return nil
}
