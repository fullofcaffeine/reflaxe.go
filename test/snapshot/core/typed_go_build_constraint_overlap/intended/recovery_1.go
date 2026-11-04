//go:build !((!gamma && (alpha || beta)) || alpha || beta) && !gamma && (alpha || beta)

package main

import "snapshot/hxrt"

func hx_recover_callable_4__probe_0(value any) (*hxrt.Function[func(*CompoundFeature) *string], bool) {
	var zero *hxrt.Function[func(*CompoundFeature) *string]
	return zero, false
}
