package main

import "snapshot/hxrt"

func describe(mode any) *string {
	var hx_switch_1 *string
	switch *hxrt.StdString(mode) {
	case *hxrt.StdString(hxrt.StringFromLiteral("human")):
		hx_switch_1 = hxrt.StringFromLiteral("human")
	case *hxrt.StdString(hxrt.StringFromLiteral("json")):
		hx_switch_1 = hxrt.StringFromLiteral("json")
	}
	return hx_switch_1
}

func main() {
	if !hxrt.StringEqualStringPtr(describe(any(hxrt.StringFromLiteral("human"))), hxrt.StringFromLiteral("human")) {
		hxrt.Throw(hxrt.StringFromLiteral("human enum-abstract switch mismatch"))
	}
	if !hxrt.StringEqualStringPtr(describe(any(hxrt.StringFromLiteral("json"))), hxrt.StringFromLiteral("json")) {
		hxrt.Throw(hxrt.StringFromLiteral("json enum-abstract switch mismatch"))
	}
	statementMatched := false
	var _g any = any(hxrt.StringFromLiteral("json"))
	switch *hxrt.StdString(_g) {
	case *hxrt.StdString(hxrt.StringFromLiteral("human")):
		hxrt.Throw(hxrt.StringFromLiteral("statement switch matched the wrong string value"))
	case *hxrt.StdString(hxrt.StringFromLiteral("json")):
		statementMatched = true
	}
	if !statementMatched {
		hxrt.Throw(hxrt.StringFromLiteral("statement enum-abstract switch missed"))
	}
}
