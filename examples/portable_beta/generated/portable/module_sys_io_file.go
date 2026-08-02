package main

import "examples_portable_beta/hxrt"

func sys__io__File_append(path *string, binary bool) *sys__io__FileOutput {
	return New_sys__io__FileOutput(hxrt.FileOpenAppend(path))
}

func sys__io__File_copy(srcPath *string, dstPath *string) {
	hxrt.FileCopyContents(srcPath, dstPath)
}

func sys__io__File_getBytes(path *string) *haxe__io__Bytes {
	values := hxrt.FileReadByteValues(path)
	bytes := haxe__io__Bytes_alloc(len(values))
	_g := 0
	_g1 := len(values)
	for _g < _g1 {
		hx_post_33 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_33
		bytes.b[index] = int(int32((hxrt.Int32Wrap(values[index]) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
	}
	return bytes
}

func sys__io__File_getContent(path *string) *string {
	return hxrt.StdString(hxrt.FileReadContent(path))
}

func sys__io__File_read(path *string, binary bool) *sys__io__FileInput {
	return New_sys__io__FileInput(hxrt.FileOpenInput(path))
}

func sys__io__File_saveBytes(path *string, bytes *haxe__io__Bytes) {
	values := hxrt.NewArray()
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_34 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_34
		values.Push(bytes.b[index])
	}
	hxrt.FileWriteByteValues(path, func(hx_lambda_raw_36 []any) []int {
		hx_lambda_out_37 := make([]int, 0, len(hx_lambda_raw_36))
		for _, hx_lambda_item_38 := range hx_lambda_raw_36 {
			hx_lambda_out_37 = append(hx_lambda_out_37, func(hx_value_39 any) int {
				if hx_value_39 == nil {
					var hx_zero_40 int
					return hx_zero_40
				}
				return hx_value_39.(int)
			}(hx_lambda_item_38))
		}
		return hx_lambda_out_37
	}(values.Values()))
}

func sys__io__File_saveContent(path *string, content *string) {
	hxrt.FileWriteContent(path, content)
}

func sys__io__File_update(path *string, binary bool) *sys__io__FileOutput {
	return New_sys__io__FileOutput(hxrt.FileOpenUpdate(path))
}

func sys__io__File_write(path *string, binary bool) *sys__io__FileOutput {
	return New_sys__io__FileOutput(hxrt.FileOpenWrite(path))
}
