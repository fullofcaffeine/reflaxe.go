//go:build exception_probe

package main

import "snapshot/hxrt"

func hx_recover_class_0(value any) *TaggedError {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *TaggedError:
		return source
	default:
		return value.(*TaggedError)
	}
}

func hx_recover_class_1__probe_0(value any) (*haxe__Exception, bool) {
	switch source := value.(type) {
	case *TaggedError:
		return source.haxe__Exception, true
	default:
		var zero *haxe__Exception
		return zero, false
	}
}

func hx_is_nominal_0__probe_0(value any) (bool, bool) {
	switch source := value.(type) {
	case *TaggedError:
		return (source != nil), true
	default:
		var zero bool
		return zero, false
	}
}

func hx_is_nominal_1__probe_0(value any) (bool, bool) {
	switch source := value.(type) {
	case *TaggedError:
		return (source != nil), true
	default:
		var zero bool
		return zero, false
	}
}

func hx_is_nominal_2__probe_0(value any) (bool, bool) {
	switch source := value.(type) {
	case *TaggedError:
		return (source != nil), true
	default:
		var zero bool
		return zero, false
	}
}
