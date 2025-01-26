// package datetime contains convenience functions for working with
// JavaScript date time strings

package datetime

import (
	"fmt"
	"regexp"
	"time"
)

// Format is the layout string for a JavaScript date time string.
// Reference infomation:
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date#date_time_string_format
const Format = "2006-01-02T15:04:05.999Z"

// Regexp is a pre-initialized regular expression for checking if a
// string is a valid JavaScript date time string.
var Regexp = regexp.MustCompile(
	`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{1,3})?Z$`,
)

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
