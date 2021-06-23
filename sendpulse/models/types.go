package models

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type DateTimeType struct {
	time.Time
}

func (d *DateTimeType) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		d.Time = time.Time{}
		return nil
	}
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

type Float32Type float32

func (d *Float32Type) UnmarshalJSON(b []byte) error {
	var value float32
	if b[0] == '"' {
		f64, err := strconv.ParseFloat(string(b[1:len(b)-1]), 32)
		if err != nil {
			return err
		}
		value = float32(f64)
	} else {
		if err := json.Unmarshal(b, &value); err != nil {
			return err
		}
	}
	*d = Float32Type(value)
	return nil
}
