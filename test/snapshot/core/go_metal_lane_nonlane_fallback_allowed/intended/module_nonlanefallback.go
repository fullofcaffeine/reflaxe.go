package main

import "snapshot/hxrt"

func NonLaneFallback_run() {
	var maybe any = go___Go_fail(hxrt.StringFromLiteral("non-lane"))
	hxrt.Println(hxrt.AnyEqualsNull(maybe))
}
