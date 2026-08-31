//go:build !cgo

package main

import "snapshot/hxrt"

func PureGoFeature_selected() *string {
	return hxrt.StringFromLiteral("pure-go")
}
