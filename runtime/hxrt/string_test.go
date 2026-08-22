package hxrt

import (
	"math"
	"testing"
)

func TestStringSliceCodePointsStringPtrUsesLogicalRuneBounds(t *testing.T) {
	value := "a😀bé"
	got := StringSliceCodePointsStringPtr(&value, 1, 3)
	if got == nil || *got != "😀b" {
		t.Fatalf("StringSliceCodePointsStringPtr() = %v, want %q", got, "😀b")
	}
}

func TestStringEqualStringPtrDistinguishesNullFromLiteral(t *testing.T) {
	nullLiteral := "null"
	if StringEqualStringPtr(nil, &nullLiteral) {
		t.Fatal("StringEqualStringPtr treated null as the literal string \"null\"")
	}
	if !StringEqualStringPtr(nil, nil) {
		t.Fatal("StringEqualStringPtr rejected two null strings")
	}
}

func TestStringCompareStringPtrOrdersValuesAndNull(t *testing.T) {
	alpha := "alpha"
	beta := "beta"
	if StringCompareStringPtr(&alpha, &beta) >= 0 {
		t.Fatal("alpha must sort before beta")
	}
	if StringCompareStringPtr(&beta, &alpha) <= 0 {
		t.Fatal("beta must sort after alpha")
	}
	if StringCompareStringPtr(&alpha, &alpha) != 0 {
		t.Fatal("equal string values must compare equally")
	}
	if StringCompareStringPtr(nil, &alpha) >= 0 || StringCompareStringPtr(&alpha, nil) <= 0 || StringCompareStringPtr(nil, nil) != 0 {
		t.Fatal("null string ordering must be deterministic and distinct")
	}
}

func TestStringIndexOfStringPtrUsesLogicalRuneIndexes(t *testing.T) {
	value := "a😀bé😀"
	needle := "😀"
	empty := ""
	missing := "missing"

	tests := []struct {
		name     string
		search   *string
		start    int
		hasStart bool
		want     int
	}{
		{name: "first unicode match", search: &needle, want: 1},
		{name: "match after start", search: &needle, start: 2, hasStart: true, want: 4},
		{name: "empty at start", search: &empty, start: 3, hasStart: true, want: 3},
		{name: "empty clamps past end", search: &empty, start: 99, hasStart: true, want: 5},
		{name: "negative start is relative to end", search: &needle, start: -1, hasStart: true, want: 4},
		{name: "empty clamps negative start to zero", search: &empty, start: -1, hasStart: true, want: 0},
		{name: "large negative start clamps to zero", search: &needle, start: -99, hasStart: true, want: 1},
		{name: "missing", search: &missing, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := StringIndexOfStringPtr(&value, test.search, test.start, test.hasStart); got != test.want {
				t.Fatalf("StringIndexOfStringPtr() = %d, want %d", got, test.want)
			}
		})
	}
}

func TestStringParseFloatExactAcceptsOnlyValidatedCompleteTokens(t *testing.T) {
	valid := "-125.5e-1"
	if got := StringParseFloatExact(&valid); got != -12.55 {
		t.Fatalf("StringParseFloatExact(%q) = %v, want -12.55", valid, got)
	}

	invalid := "12tail"
	if got := StringParseFloatExact(&invalid); !math.IsNaN(got) {
		t.Fatalf("StringParseFloatExact(%q) = %v, want NaN", invalid, got)
	}
	if got := StringParseFloatExact(nil); !math.IsNaN(got) {
		t.Fatalf("StringParseFloatExact(nil) = %v, want NaN", got)
	}
}
