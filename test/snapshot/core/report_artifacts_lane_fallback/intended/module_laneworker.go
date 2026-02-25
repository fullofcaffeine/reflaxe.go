package main

import "snapshot/hxrt"

func LaneWorker_produce() {
	var laneResult any = go___Go_fail(hxrt.StringFromLiteral("lane"))
	hxrt.Println(hxrt.AnyEqualsNull(laneResult))
}
