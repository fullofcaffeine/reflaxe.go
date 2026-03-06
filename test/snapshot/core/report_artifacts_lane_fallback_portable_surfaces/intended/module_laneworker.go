package main

import "snapshot/hxrt"

func LaneWorker_produce() {
	hxrt.Println(PortableSurfaceDigest_compute(11))
	laneResult := LaneWorker_unresolvedFail(hxrt.StringFromLiteral("lane"))
	hxrt.Println((laneResult == nil))
}

func LaneWorker_unresolvedFail(message *string) *go___Result {
	return go___Go_fail(message)
}
