package date

import (
	"errors"
	"time"
)

var (
	ErrFormatNotMatch = errors.New("date didn't match any of the supported format")
)

// TryParse tries to parse a date string into a time.Time object using a list of supported date/datetime format.
func TryParse(date string) (time.Time, error) {
	for _, f := range []string{
		"2006-01-02",
		"20060102",
		"January 02, 2006",
		"02 January 2006",
		"02-Jan-2006",
		"Jan-02-2006",
		"Jan-02-06",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05-0700",
		"2 Jan 2006 15:04:05",
		"2 Jan 2006 15:04",
		"Mon, 2 Jan 2006 15:04:05 MST",
		"Jan-06",
	} {
		t, err := time.Parse(f, date)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, ErrFormatNotMatch
}
