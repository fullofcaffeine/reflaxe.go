package main

import "snapshot/hxrt"

func NativeSlicer_middle(values *hxrt.Array) *hxrt.Array {
	return values.SliceOptional(1, -1)
}
