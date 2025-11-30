package date

import (
	"errors"
	"strconv"
	"time"
)

var ErrInvalidDateOnlyFormat = errors.New("invalid date-only format")

type DateOnly time.Time

func (d *DateOnly) UnmarshalJSON(jsonValue []byte) error {
	unQuoteJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidDateOnlyFormat
	}

	parsedTime, err := time.Parse("2006-01-02", unQuoteJSONValue)
	if err != nil {
		return ErrInvalidDateOnlyFormat
	}

	// Set the DateOnly value to the parsed time
	*d = DateOnly(parsedTime)
	return nil
}


func (d DateOnly) ToTime() time.Time {
	return time.Time(d)
}


