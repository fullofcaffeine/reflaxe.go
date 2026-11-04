//go:build !cgo

package main

func hx_is_nominal_0__probe_0(value any) (bool, bool) {
	switch source := value.(type) {
	case *PureGoFeature:
		return (source != nil), true
	default:
		var zero bool
		return zero, false
	}
}

func hx_is_nominal_0__probe_1(value any) (bool, bool) {
	var zero bool
	return zero, false
}
