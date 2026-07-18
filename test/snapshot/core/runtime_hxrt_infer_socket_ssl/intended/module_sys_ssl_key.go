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
		hx_post_9 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_9
		values.Push(data.b[index])
	}
	return New_sys__ssl__Key(hxrt.SslKeyReadDERValues(func(hx_lambda_raw_11 []any) []int {
		hx_lambda_out_12 := make([]int, 0, len(hx_lambda_raw_11))
		for _, hx_lambda_item_13 := range hx_lambda_raw_11 {
			hx_lambda_out_12 = append(hx_lambda_out_12, func(hx_value_14 any) int {
				if hx_value_14 == nil {
					var hx_zero_15 int
					return hx_zero_15
				}
				return hx_value_14.(int)
			}(hx_lambda_item_13))
		}
		return hx_lambda_out_12
	}(values.Values()), isPublic))
}

func sys__ssl__Key_readPEM(data *string, isPublic bool, pass *string) *sys__ssl__Key {
	return New_sys__ssl__Key(hxrt.SslKeyReadPEM(data, isPublic, pass))
}
