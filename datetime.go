package datetime

import (
	"fmt"
	"time"
)

// Format is the layout string for a JavaScript date time string.
// Reference infomation:
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date#date_time_string_format
const Format = "2006-01-02T15:04:05.999Z"

// Parse parses a JavaScript date time string and returns the time value
// it represents.
func Parse(dateTime string) (time.Time, error) {
	t, err := time.Parse(Format, dateTime)
	if err != nil {
		err = fmt.Errorf("datetime.Parse: %w", err)
	}
	return t, err
}

// MustParse is like Parse but panics if the expression cannot be parsed.
func MustParse(dateTime string) time.Time {
	t, err := Parse(dateTime)
	if err != nil {
		panic(err)
	}
	return t
}

// String returns the time formatted as a JavaScript date time string.
func String(t time.Time) string {
	return t.UTC().Format(Format)
}
