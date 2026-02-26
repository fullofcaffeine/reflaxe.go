package main

import "snapshot/hxrt"

func NonLaneFallback_run() {
	nonLaneResult := NonLaneFallback_unresolvedFail(hxrt.StringFromLiteral("non-lane"))
	hxrt.Println((nonLaneResult == nil))
}

func NonLaneFallback_unresolvedFail(message *string) *go___Result {
	return go___Go_fail(message)
}
