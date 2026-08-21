package main

import "snapshot/hxrt"

func expect(actual int, expected int) {
	if actual != expected {
		hxrt.Throw(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("expected "), expected), hxrt.StringFromLiteral(", got ")), actual))
	}
}

func main() {
	value := hxrt.StringFromLiteral("a😀bé😀")
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral("😀"), 0, false), 1)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral("😀"), 2, true), 4)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral(""), 3, true), 3)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral(""), 99, true), 5)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral("😀"), -1, true), 4)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral(""), -1, true), 0)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral("😀"), -99, true), 1)
	expect(hxrt.StringIndexOfStringPtr(value, hxrt.StringFromLiteral("missing"), 0, false), -1)
}
