package sendpulse_sdk_go

import (
	"fmt"
	"strings"
	"time"
)

type DateTime time.Time

const dtFormat = "2006-01-02 15:04:05"

func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		*d = DateTime(time.Time{})
		return nil
	}
	t, err := time.Parse(dtFormat, s)
	if err != nil {
		return err
	}

	*d = DateTime(t)
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *DateTime) String() string {
	t := time.Time(*d)
	return fmt.Sprintf("%q", t.Format(dtFormat))
}

type Float32 float32

func (f *Float32) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		*f = Float32(0)
		return nil
	}

	v, err := fmt.Sscanf(s, "%f", f)
	if err != nil {
		return fmt.Errorf("sscanf: %w", err)
	}
	if v != 1 {
		return fmt.Errorf("failed to parse float32")
	}
	return nil
}
