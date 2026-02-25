package main

import "snapshot/hxrt"

func NonLaneWorker_produce() {
	var nonLaneResult any = go___Go_fail(hxrt.StringFromLiteral("non-lane"))
	hxrt.Println(hxrt.AnyEqualsNull(nonLaneResult))
}
