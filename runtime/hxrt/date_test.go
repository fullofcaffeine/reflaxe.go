package hxrt

import (
	"testing"
	"time"
	_ "time/tzdata"
)

func TestDateLocalAndUTCPartsRespectHostTimezone(t *testing.T) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("load deterministic timezone: %v", err)
	}
	withDateTestLocation(t, location)

	winter := DateLocalTime(2024, 0, 15, 12, 34, 56)
	winterLocal := DateLocalParts(winter)
	if winterLocal.FullYear != 2024 || winterLocal.Month != 0 || winterLocal.Date != 15 || winterLocal.Hours != 12 || winterLocal.Minutes != 34 || winterLocal.Seconds != 56 {
		t.Fatalf("winter local parts = %#v", winterLocal)
	}
	if offset := DateTimezoneOffset(winter); offset != 300 {
		t.Fatalf("winter timezone offset = %d, want 300", offset)
	}

	summer := DateLocalTime(2024, 6, 15, 12, 34, 56)
	if offset := DateTimezoneOffset(summer); offset != 240 {
		t.Fatalf("summer timezone offset = %d, want 240", offset)
	}
	utc := DateUTCParts(summer)
	if utc.FullYear != 2024 || utc.Month != 6 || utc.Date != 15 || utc.Hours != 16 || utc.Minutes != 34 || utc.Seconds != 56 {
		t.Fatalf("summer UTC parts = %#v", utc)
	}
	if got := *StdString(DateFormatLocal(summer)); got != "2024-07-15 12:34:56" {
		t.Fatalf("formatted summer date = %q", got)
	}
}

func TestDateParsingCoversEveryHaxeFormat(t *testing.T) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("load deterministic timezone: %v", err)
	}
	withDateTestLocation(t, location)

	full := DateParse(StringFromLiteral("2024-02-29 15:04:05"))
	if got := DateLocalParts(full); got.FullYear != 2024 || got.Month != 1 || got.Date != 29 || got.Hours != 15 || got.Minutes != 4 || got.Seconds != 5 {
		t.Fatalf("full datetime parts = %#v", got)
	}
	dateOnly := DateParse(StringFromLiteral("2024-02-29"))
	if got := DateLocalParts(dateOnly); got.FullYear != 2024 || got.Month != 1 || got.Date != 29 || got.Hours != 0 || got.Minutes != 0 || got.Seconds != 0 {
		t.Fatalf("date-only parts = %#v", got)
	}
	if got := DateParse(StringFromLiteral("03:04:05")); got != 11_045_000 {
		t.Fatalf("time-only milliseconds = %v, want 11045000", got)
	}
}

func TestDateInvalidInputCrossesExceptionCarrier(t *testing.T) {
	deferred := false
	func() {
		defer func() {
			deferred = recover() != nil
		}()
		DateParse(StringFromLiteral("not-a-date"))
	}()
	if !deferred {
		t.Fatal("invalid date did not cross the hxrt exception carrier")
	}
}

func TestDateNowUsesEpochMilliseconds(t *testing.T) {
	before := float64(time.Now().UnixMilli())
	current := DateNow()
	after := float64(time.Now().UnixMilli())
	if current < before-10 || current > after+10 {
		t.Fatalf("DateNow() = %v, expected [%v, %v]", current, before, after)
	}
}

func TestDateConversionsDoNotInheritUnixNanoRange(t *testing.T) {
	withDateTestLocation(t, time.UTC)

	future := DateLocalTime(2500, 0, 2, 3, 4, 5)
	if got := DateLocalParts(future); got.FullYear != 2500 || got.Month != 0 || got.Date != 2 || got.Hours != 3 || got.Minutes != 4 || got.Seconds != 5 {
		t.Fatalf("future local parts = %#v", got)
	}

	past := DateLocalTime(1600, 10, 12, 13, 14, 15)
	if got := DateUTCParts(past); got.FullYear != 1600 || got.Month != 10 || got.Date != 12 || got.Hours != 13 || got.Minutes != 14 || got.Seconds != 15 {
		t.Fatalf("past UTC parts = %#v", got)
	}
}

func withDateTestLocation(t *testing.T, location *time.Location) {
	t.Helper()
	previous := time.Local
	time.Local = location
	t.Cleanup(func() {
		time.Local = previous
	})
}
