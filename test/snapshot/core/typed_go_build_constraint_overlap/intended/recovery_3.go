//go:build !((!gamma && (alpha || beta)) || alpha || beta) && beta

package main

import "snapshot/hxrt"

func hx_recover_callable_3__probe_1(value any) (*hxrt.Function[func(*BetaFeature) *string], bool) {
	var zero *hxrt.Function[func(*BetaFeature) *string]
	return zero, false
}
