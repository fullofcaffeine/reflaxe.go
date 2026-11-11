//go:build go1.26

package main

import (
	"errors"
	"io/fs"
	"snapshot/hxrt"
)

func NativeErrorProbe_check() {
	match := func() *ErrorMatch {
		hx_tuple_1, hx_tuple_2 := errors.AsType[*fs.PathError](nil)
		return New_ErrorMatch(hx_tuple_1, hx_tuple_2)
	}()
	var v any = any(match.matched)
	hxrt.Println(v)
	var v_1 any = any(hxrt.AnyEqualsNull(func(hx_value_3 any) *fs.PathError {
		if hx_value_3 == nil {
			var hx_zero_4 *fs.PathError
			return hx_zero_4
		}
		return hx_value_3.(*fs.PathError)
	}(match.value)))
	hxrt.Println(v_1)
}
