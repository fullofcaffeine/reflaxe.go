//go:build cgo

package main

import "snapshot/hxrt"

func hx_recover_callable_0(value any) *hxrt.Function[func(*CgoFeature) *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*CgoFeature) *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*CgoFeature) *string]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *CgoFeature {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *CgoFeature:
		return source
	default:
		return value.(*CgoFeature)
	}
}
