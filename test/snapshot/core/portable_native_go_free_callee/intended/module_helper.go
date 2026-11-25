package main

import "snapshot/hxrt"

func _Helper__Helper_Fields__describe(value any) *string {
	var hx_if_1 *string
	if hxrt.AnyEqualsNull(value) {
		hx_if_1 = hxrt.StringFromLiteral("nil")
	} else {
		hx_if_1 = hxrt.StringFromLiteral("module value")
	}
	return hx_if_1
}

func _Helper__Helper_Fields__plain() *string {
	return _Helper__Helper_Fields__describe("module")
}
