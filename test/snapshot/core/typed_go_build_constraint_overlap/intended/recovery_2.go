//go:build !((!gamma && (alpha || beta)) || alpha || beta) && alpha

package main

import "snapshot/hxrt"

func hx_recover_callable_2__probe_1(value any) (*hxrt.Function[func(*AlphaFeature) *string], bool) {
	var zero *hxrt.Function[func(*AlphaFeature) *string]
	return zero, false
}
