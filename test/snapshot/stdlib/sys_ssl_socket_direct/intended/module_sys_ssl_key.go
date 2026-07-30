package main

import "snapshot/hxrt"

type I_sys__ssl__Key interface {
}

type sys__ssl__Key struct {
	__hx_this I_sys__ssl__Key
	handle    *hxrt.SslKey
}

func New_sys__ssl__Key(handle *hxrt.SslKey) *sys__ssl__Key {
	self := &sys__ssl__Key{}
	self.__hx_this = self
	self.handle = handle
	return self
}

func sys__ssl__Key_loadFile(file *string, isPublic any, pass *string) *sys__ssl__Key {
	return New_sys__ssl__Key(hxrt.SslKeyLoadFile(file, (isPublic.(bool) == true), pass))
}

func sys__ssl__Key_readDER(data *haxe__io__Bytes, isPublic bool) *sys__ssl__Key {
	values := hxrt.NewArray()
	_g := 0
	_g1 := data.length
	for _g < _g1 {
		hx_post_22 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_22
		values.Push(data.b[index])
	}
	return New_sys__ssl__Key(hxrt.SslKeyReadDERValues(func(hx_lambda_raw_24 []any) []int {
		hx_lambda_out_25 := make([]int, 0, len(hx_lambda_raw_24))
		for _, hx_lambda_item_26 := range hx_lambda_raw_24 {
			hx_lambda_out_25 = append(hx_lambda_out_25, func(hx_value_27 any) int {
				if hx_value_27 == nil {
					var hx_zero_28 int
					return hx_zero_28
				}
				return hx_value_27.(int)
			}(hx_lambda_item_26))
		}
		return hx_lambda_out_25
	}(values.Values()), isPublic))
}

func sys__ssl__Key_readPEM(data *string, isPublic bool, pass *string) *sys__ssl__Key {
	return New_sys__ssl__Key(hxrt.SslKeyReadPEM(data, isPublic, pass))
}
