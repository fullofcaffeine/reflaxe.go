package hxrt

import "testing"

func TestStringSliceCodePointsStringPtrUsesLogicalRuneBounds(t *testing.T) {
	value := "a😀bé"
	got := StringSliceCodePointsStringPtr(&value, 1, 3)
	if got == nil || *got != "😀b" {
		t.Fatalf("StringSliceCodePointsStringPtr() = %v, want %q", got, "😀b")
	}
}
