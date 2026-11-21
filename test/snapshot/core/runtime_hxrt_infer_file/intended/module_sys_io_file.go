package main

import "snapshot/hxrt"

func sys__io__File_read(path *string, binary bool) *sys__io__FileInput {
	return New_sys__io__FileInput(hxrt.FileOpenInput(path))
}

func sys__io__File_saveBytes(path *string, bytes *haxe__io__Bytes) {
	values := hxrt.NewArray()
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_1 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_1
		values.Push(bytes.b[index])
	}
	hxrt.FileWriteByteValues(path, func(hx_lambda_raw_3 []any) []int {
		hx_lambda_out_4 := make([]int, 0, len(hx_lambda_raw_3))
		for _, hx_lambda_item_5 := range hx_lambda_raw_3 {
			hx_lambda_out_4 = append(hx_lambda_out_4, func(hx_value_6 any) int {
				if hx_value_6 == nil {
					var hx_zero_7 int
					return hx_zero_7
				}
				return hx_value_6.(int)
			}(hx_lambda_item_5))
		}
		return hx_lambda_out_4
	}(values.Values()))
}

func sys__io__File_update(path *string, binary bool) *sys__io__FileOutput {
	return New_sys__io__FileOutput(hxrt.FileOpenUpdate(path))
}
