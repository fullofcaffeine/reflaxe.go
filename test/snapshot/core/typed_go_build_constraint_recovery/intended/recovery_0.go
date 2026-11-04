//go:build !cgo

package main

import "snapshot/hxrt"

func hx_recover_callable_1(value any) *hxrt.Function[func(*PureGoFeature) *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*PureGoFeature) *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*PureGoFeature) *string]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_1(value any) *PureGoFeature {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *PureGoFeature:
		return source
	default:
		return value.(*PureGoFeature)
	}
}
