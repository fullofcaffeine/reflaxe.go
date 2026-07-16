package hxrt

import (
	"fmt"
	"time"
)

// DateParts is the typed, representation-neutral calendar carrier consumed by
// staged Date source.
//
// What: Carries one local or UTC decomposition using Haxe field conventions.
// Why: Go time.Time must not become part of an application-generated Date layout.
// How: Store only scalar values; Month and Day are already zero-based where the
// Haxe API requires it.
type DateParts struct {
	FullYear int
	Month    int
	Date     int
	Day      int
	Hours    int
	Minutes  int
	Seconds  int
}

// DateLocalTime constructs epoch milliseconds from local calendar components.
//
// What: Supplies the native portion of the Haxe Date constructor.
// Why: Local timezone and daylight-saving rules come from the host Go runtime.
// How: Convert Haxe's zero-based month and delegate normalization to time.Date.
func DateLocalTime(year int, month int, day int, hour int, min int, sec int) float64 {
	value := time.Date(year, time.Month(month+1), day, hour, min, sec, 0, time.Local)
	return dateMilliseconds(value)
}

// DateNow returns the current Unix timestamp in milliseconds.
//
// What: Supplies Date.now with wall-clock time.
// Why: Reading the host clock is a native capability rather than Haxe library policy.
// How: Convert time.Now directly to the portable epoch-millisecond carrier.
func DateNow() float64 {
	return dateMilliseconds(time.Now())
}

// DateParse accepts the three string forms in the Haxe 4.3.7 Date contract.
//
// What: Parses local datetime/date strings and epoch-relative UTC time strings.
// Why: Go owns timezone-aware parsing, while the accepted formats remain fixed by
// the staged Haxe API.
// How: Try the two local layouts, then compute milliseconds since UTC epoch for
// the time-only layout; invalid input crosses hxrt.Throw.
func DateParse(value *string) float64 {
	raw := *StdString(value)
	if parsed, err := time.ParseInLocation("2006-01-02 15:04:05", raw, time.Local); err == nil {
		return dateMilliseconds(parsed)
	}
	if parsed, err := time.ParseInLocation("2006-01-02", raw, time.Local); err == nil {
		return dateMilliseconds(parsed)
	}
	if parsed, err := time.Parse("15:04:05", raw); err == nil {
		return float64((parsed.Hour()*60*60 + parsed.Minute()*60 + parsed.Second()) * 1000)
	}
	Throw(fmt.Errorf("invalid date format: %s", raw))
	return 0
}

// DateLocalParts decomposes epoch milliseconds in the host local timezone.
//
// What: Supplies all local Date component getters through one typed carrier.
// Why: Repeating conversion per field would duplicate timezone work and API surface.
// How: Convert once through time.Local and map Go conventions to Haxe conventions.
func DateLocalParts(milliseconds float64) *DateParts {
	return dateParts(dateTime(milliseconds).In(time.Local))
}

// DateUTCParts decomposes epoch milliseconds in UTC.
//
// What: Supplies all UTC Date component getters through one typed carrier.
// Why: UTC behavior is part of the public Haxe API but conversion is native time work.
// How: Convert the portable timestamp once and map the fields into DateParts.
func DateUTCParts(milliseconds float64) *DateParts {
	return dateParts(dateTime(milliseconds).UTC())
}

// DateTimezoneOffset returns Haxe's local-to-UTC difference in minutes.
//
// What: Implements Date.getTimezoneOffset for the represented instant.
// Why: Offsets can vary with daylight-saving transitions and must come from host
// timezone data.
// How: Negate Go's local east-of-UTC offset to match Haxe and JavaScript signs.
func DateTimezoneOffset(milliseconds float64) int {
	_, secondsEastOfUTC := dateTime(milliseconds).In(time.Local).Zone()
	return -secondsEastOfUTC / 60
}

// DateFormatLocal emits the canonical Haxe local date string.
//
// What: Formats YYYY-MM-DD HH:MM:SS in the local timezone.
// Why: The format is Haxe policy, but locale/timezone conversion is native.
// How: Use Go's fixed reference layout and return an hxrt string pointer.
func DateFormatLocal(milliseconds float64) *string {
	formatted := dateTime(milliseconds).In(time.Local).Format("2006-01-02 15:04:05")
	return StringFromLiteral(formatted)
}

func dateParts(value time.Time) *DateParts {
	return &DateParts{
		FullYear: value.Year(),
		Month:    int(value.Month()) - 1,
		Date:     value.Day(),
		Day:      int(value.Weekday()),
		Hours:    value.Hour(),
		Minutes:  value.Minute(),
		Seconds:  value.Second(),
	}
}

// Use the millisecond API so valid calendar years do not inherit UnixNano's
// roughly 1678-to-2262 representable range. The staged Date retains any
// fractional millisecond in its own Float carrier; calendar access is defined
// at whole-millisecond precision here.
func dateTime(milliseconds float64) time.Time {
	return time.UnixMilli(int64(milliseconds))
}

// UnixMilli preserves the same wide time.Time range when a native calendar
// value enters the portable epoch-millisecond carrier.
func dateMilliseconds(value time.Time) float64 {
	return float64(value.UnixMilli())
}
