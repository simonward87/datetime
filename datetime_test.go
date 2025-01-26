package datetime_test

import (
	"testing"
	"time"

	"github.com/simonward87/datetime"
)

func TestString(t *testing.T) {
	testCases := []struct {
		expect string
		layout string
		name   string
		value  string
	}{
		{
			expect: "2006-01-02T22:04:05Z",
			layout: time.Layout,
			name:   "without_subseconds",
			value:  time.Layout,
		},
		{
			expect: "2025-01-26T12:18:15.722Z",
			layout: "2006/01/02 15:04:05.999",
			name:   "with_subseconds",
			value:  "2025/01/26 12:18:15.722",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			timestamp, err := time.Parse(tc.layout, tc.value)
			if err != nil {
				t.Fatal(err)
			}

			dt := datetime.String(timestamp)
			if !datetime.Regexp.MatchString(dt) {
				t.Errorf("regexp.Match: got %s, want %s", dt, tc.expect)
			}
		})
	}

}

func TestParse(t *testing.T) {
	timestamp, err := datetime.Parse("2025-01-26T12:18:15.722Z")
	if err != nil {
		t.Fatal(err)
	}

	expect, err := time.Parse(
		"Jan _2 2006 15:04:05.000 -0700",
		"Jan 26 2025 12:18:15.722 -0000",
	)
	if err != nil {
		t.Fatal(err)
	}

	if out := expect.Sub(timestamp); out != 0 {
		t.Errorf("time.Sub: got %v, want %v", 0, out)
	}
}
