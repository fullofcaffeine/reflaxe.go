//go:build cgo

package main

import "snapshot/hxrt"

func CgoFeature_selected() *string {
	return hxrt.StringFromLiteral("cgo")
}
