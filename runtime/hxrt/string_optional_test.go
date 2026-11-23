package hxrt

import "testing"

// stringOptionalRow renders `substr(1, v)|substring(0, v)|indexOf("b", v)|
// lastIndexOf("b", v)` through the helpers for a `Null<Int>` argument.
func stringOptionalRow(text *string, value any) string {
	search := StringFromLiteral("b")
	typed := *StringSubstrOptionalStringPtr(text, 1, value) + "|" +
		*StringSubstringOptionalStringPtr(text, 0, value) + "|" +
		*StdString(StringIndexOfOptionalStringPtr(text, search, value)) + "|" +
		*StdString(StringLastIndexOfOptionalStringPtr(text, search, value))
	generic := *StringSubstrOptional(text, 1, value) + "|" +
		*StringSubstringOptional(text, 0, value) + "|" +
		*StdString(StringIndexOfOptional(text, search, value)) + "|" +
		*StdString(StringLastIndexOfOptional(text, search, value))
	if typed != generic {
		return typed + " (any receiver: " + generic + ")"
	}
	return typed
}

// A null optional argument reads as omitted, as in Haxe `--interp`; any other
// value reads as that Int, whatever integer type stores it.
func TestStringOptionalHelpersReadNullAsOmitted(t *testing.T) {
	text := StringFromLiteral("abcabc")
	cases := []struct {
		name  string
		value any
		want  string
	}{
		{"null", nil, "bcabc|abcabc|1|4"},
		{"value", 2, "bc|ab|4|1"},
		{"negative", -2, "bca||4|-1"},
		{"out of range", 99, "bcabc|abcabc|-1|4"},
		{"negative out of range", -99, "||1|-1"},
		{"zero", 0, "||1|-1"},
		{"int64 storage", int64(3), "bca|abc|4|1"},
	}
	for _, c := range cases {
		if got := stringOptionalRow(text, c.value); got != c.want {
			t.Fatalf("%s: got %q, want %q", c.name, got, c.want)
		}
	}
}

// A null argument gives the same result as the omitted form of each method,
// also for code points beyond ASCII.
func TestStringOptionalHelpersMatchOmittedForms(t *testing.T) {
	text := StringFromLiteral("héllo")
	search := StringFromLiteral("l")
	if got, want := *StringSubstrOptionalStringPtr(text, 1, nil), *StringSubstrStringPtr(text, 1, 0, false); got != want {
		t.Fatalf("substr: got %q, want %q", got, want)
	}
	if got, want := *StringSubstringOptionalStringPtr(text, 1, nil), *StringSubstringStringPtr(text, 1, StringLengthStringPtr(text)); got != want {
		t.Fatalf("substring: got %q, want %q", got, want)
	}
	if got, want := StringIndexOfOptionalStringPtr(text, search, nil), StringIndexOfStringPtr(text, search, 0, false); got != want {
		t.Fatalf("indexOf: got %d, want %d", got, want)
	}
	if got, want := StringLastIndexOfOptionalStringPtr(text, search, nil), StringLastIndexOfStringPtr(text, search, 0, false); got != want {
		t.Fatalf("lastIndexOf: got %d, want %d", got, want)
	}
}
