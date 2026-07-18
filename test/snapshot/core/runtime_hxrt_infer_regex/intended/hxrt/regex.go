package hxrt

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// RegexHandle is the opaque native regular-expression capability retained by EReg.
//
// What: Own one compiled Go RE2 expression.
// Why: A compiled regexp is a host resource and cannot be expressed as portable
// Haxe data, while EReg state and public policy belong in staged Haxe source.
// How: Expose only compilation, match-index discovery, and quoting through
// NativeRegex; generated application code never sees regexp.Regexp.
type RegexHandle struct {
	expression *regexp.Regexp
}

// RegexMatch carries one complete match as Haxe string (code-point) offsets.
//
// Go's regexp package reports UTF-8 byte offsets, but haxe.go String indexing is
// code-point based. Converting here keeps the target-native fact out of EReg policy.
type RegexMatch struct {
	Indices []int
}

// RegexCompile validates EReg options and compiles one native expression.
func RegexCompile(pattern *string, options *string) *RegexHandle {
	rawPattern := *StdString(pattern)
	rawOptions := *StdString(options)
	inlineFlags := ""
	for _, option := range rawOptions {
		switch option {
		case 'g', 'u':
			// Global behavior is source-owned; RE2 is UTF-8 aware by default.
		case 'i', 'm', 's':
			if !strings.ContainsRune(inlineFlags, option) {
				inlineFlags += string(option)
			}
		default:
			Throw(StringFromLiteral("Unsupported regexp option '" + string(option) + "'"))
			return &RegexHandle{expression: regexp.MustCompile("a^")}
		}
	}
	if inlineFlags != "" {
		rawPattern = "(?" + inlineFlags + ")" + rawPattern
	}
	compiled, err := regexp.Compile(rawPattern)
	if err != nil {
		Throw(err)
		return &RegexHandle{expression: regexp.MustCompile("a^")}
	}
	return &RegexHandle{expression: compiled}
}

func regexRuneIndexToByteOffset(value string, runeIndex int) int {
	if runeIndex <= 0 {
		return 0
	}
	current := 0
	for byteOffset := range value {
		if current == runeIndex {
			return byteOffset
		}
		current++
	}
	return len(value)
}

func regexByteIndexToRuneOffset(value string, byteIndex int) int {
	if byteIndex < 0 {
		return -1
	}
	if byteIndex > len(value) {
		byteIndex = len(value)
	}
	return utf8.RuneCountInString(value[:byteIndex])
}

func regexMatchFromBytes(value string, indexes []int) *RegexMatch {
	converted := make([]int, len(indexes))
	for index, byteOffset := range indexes {
		converted[index] = regexByteIndexToRuneOffset(value, byteOffset)
	}
	return &RegexMatch{Indices: converted}
}

// RegexFind returns the first match at or after one Haxe code-point offset.
func RegexFind(handle *RegexHandle, source *string, position int) *RegexMatch {
	if handle == nil || handle.expression == nil {
		return nil
	}
	raw := *StdString(source)
	if position < 0 {
		position = 0
	}
	if position > utf8.RuneCountInString(raw) {
		return nil
	}
	byteStart := regexRuneIndexToByteOffset(raw, position)
	found := handle.expression.FindStringSubmatchIndex(raw[byteStart:])
	if found == nil {
		return nil
	}
	shifted := make([]int, len(found))
	for index, byteOffset := range found {
		if byteOffset < 0 {
			shifted[index] = -1
		} else {
			shifted[index] = byteOffset + byteStart
		}
	}
	return regexMatchFromBytes(raw, shifted)
}

// RegexEscape quotes one literal string for use as an RE2 pattern.
func RegexEscape(value *string) *string {
	return StringFromLiteral(regexp.QuoteMeta(*StdString(value)))
}
