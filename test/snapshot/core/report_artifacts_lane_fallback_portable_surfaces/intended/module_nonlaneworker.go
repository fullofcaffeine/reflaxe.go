package main

import "snapshot/hxrt"

func NonLaneWorker_produce() {
	hxrt.Println(PortableSurfaceDigest_compute(13))
	nonLaneResult := NonLaneWorker_unresolvedFail(hxrt.StringFromLiteral("non-lane"))
	hxrt.Println((nonLaneResult == nil))
}

func NonLaneWorker_unresolvedFail(message *string) *go___Result {
	return go___Go_fail(message)
}
