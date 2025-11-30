package float

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Price float64

func (p *Price) UnmarshalJSON(data []byte) error {
	var f float64
	var s string

	// Try float first
	if err := json.Unmarshal(data, &f); err == nil {
		*p = Price(f)
		return nil
	}

	// Try string
	if err := json.Unmarshal(data, &s); err == nil {
		parsed, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		*p = Price(parsed)
		return nil
	}

	return errors.New("invalid price format")
}
