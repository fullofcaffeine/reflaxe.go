package main

import "snapshot/hxrt"

func main() {
	if (!hxrt.StringEqualStringPtr(_GuardedTry__GuardedTry_Fields__readValue(true, false), hxrt.StringFromLiteral("")) || !hxrt.StringEqualStringPtr(_GuardedTry__GuardedTry_Fields__readValue(false, false), hxrt.StringFromLiteral("value"))) || !hxrt.StringEqualStringPtr(_GuardedTry__GuardedTry_Fields__readValue(false, true), hxrt.StringFromLiteral("fallback")) {
		hxrt.Throw(hxrt.StringFromLiteral("unexpected result"))
	}
}
