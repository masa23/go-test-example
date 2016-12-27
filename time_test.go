package main

import (
	"testing"
	"time"
)

// This code is copied from
// https://github.com/golang/go/blob/master/src/time/time_test.go

type TimeFormatTest struct {
	time           time.Time
	formattedValue string
}

var rfc3339Formats = []TimeFormatTest{
	{time.Date(2008, 9, 17, 20, 4, 26, 0, time.UTC), "2008-09-17T20:04:26Z"},
	{time.Date(1994, 9, 17, 20, 4, 26, 0, time.FixedZone("EST", -18000)), "1994-09-17T20:04:26-05:00"},
	{time.Date(2000, 12, 26, 1, 15, 6, 0, time.FixedZone("OTO", 15600)), "2000-12-26T01:15:06+04:20"},
}

func TestRFC3339Conversion(t *testing.T) {
	for _, f := range rfc3339Formats {
		if f.time.Format(time.RFC3339) != f.formattedValue {
			t.Error("RFC3339:")
			t.Errorf("  want=%+v", f.formattedValue)
			t.Errorf("  have=%+v", f.time.Format(time.RFC3339))
		}
	}
}

type FormatTest struct {
	name   string
	format string
	result string
}

var formatTests = []FormatTest{
	{"ANSIC", time.ANSIC, "Wed Feb  4 21:00:57 2009"},
	{"UnixDate", time.UnixDate, "Wed Feb  4 21:00:57 PST 2009"},
	{"RubyDate", time.RubyDate, "Wed Feb 04 21:00:57 -0800 2009"},
	{"RFC822", time.RFC822, "04 Feb 09 21:00 PST"},
	{"RFC850", time.RFC850, "Wednesday, 04-Feb-09 21:00:57 PST"},
	{"RFC1123", time.RFC1123, "Wed, 04 Feb 2009 21:00:57 PST"},
	{"RFC1123Z", time.RFC1123Z, "Wed, 04 Feb 2009 21:00:57 -0800"},
	{"RFC3339", time.RFC3339, "2009-02-04T21:00:57-08:00"},
	{"RFC3339Nano", time.RFC3339Nano, "2009-02-04T21:00:57.0123456-08:00"},
	{"Kitchen", time.Kitchen, "9:00PM"},
	{"am/pm", "3pm", "9pm"},
	{"AM/PM", "3PM", "9PM"},
	{"two-digit year", "06 01 02", "09 02 04"},
	// Three-letter months and days must not be followed by lower-case letter.
	{"Janet", "Hi Janet, the Month is January", "Hi Janet, the Month is February"},
	// Time stamps, Fractional seconds.
	{"Stamp", time.Stamp, "Feb  4 21:00:57"},
	{"StampMilli", time.StampMilli, "Feb  4 21:00:57.012"},
	{"StampMicro", time.StampMicro, "Feb  4 21:00:57.012345"},
	{"StampNano", time.StampNano, "Feb  4 21:00:57.012345600"},
}

func TestFormat(t *testing.T) {
	// The numeric time represents Thu Feb  4 21:00:57.012345600 PST 2010
	time_ := time.Unix(0, 1233810057012345600)
	for _, test := range formatTests {
		result := time_.Format(test.format)
		if result != test.result {
			t.Errorf("%s expected %q got %q", test.name, test.result, result)
		}
	}
}
