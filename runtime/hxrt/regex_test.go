package hxrt

import "testing"

func TestRegexFindConvertsUTF8OffsetsToHaxeCodePoints(t *testing.T) {
	handle := RegexCompile(StringFromLiteral("é(.)"), StringFromLiteral(""))
	match := RegexFind(handle, StringFromLiteral("aé🙂z"), 0)
	if match == nil {
		t.Fatal("expected a match")
	}
	want := []int{1, 3, 2, 3}
	if len(match.Indices) != len(want) {
		t.Fatalf("indices length = %d, want %d", len(match.Indices), len(want))
	}
	for index, value := range want {
		if match.Indices[index] != value {
			t.Fatalf("indices[%d] = %d, want %d", index, match.Indices[index], value)
		}
	}
}

func TestRegexFindUsesCodePointStartOffset(t *testing.T) {
	handle := RegexCompile(StringFromLiteral("🙂"), StringFromLiteral(""))
	match := RegexFind(handle, StringFromLiteral("é🙂"), 1)
	if match == nil || len(match.Indices) < 2 {
		t.Fatal("expected a match after the first code point")
	}
	if match.Indices[0] != 1 || match.Indices[1] != 2 {
		t.Fatalf("indices = %v, want [1 2]", match.Indices)
	}
}
