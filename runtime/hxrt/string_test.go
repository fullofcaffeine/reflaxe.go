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
