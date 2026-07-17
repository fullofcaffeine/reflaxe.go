package hxrt

import "testing"

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
