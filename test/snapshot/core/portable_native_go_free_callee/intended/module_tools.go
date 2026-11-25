package main

import "snapshot/hxrt"

func Tools_describe(value any) *string {
	var hx_if_1 *string
	if hxrt.AnyEqualsNull(value) {
		hx_if_1 = hxrt.StringFromLiteral("nil")
	} else {
		hx_if_1 = hxrt.StringFromLiteral("class value")
	}
	return hx_if_1
}

func Tools_plain() *string {
	return Tools_describe("class")
}
