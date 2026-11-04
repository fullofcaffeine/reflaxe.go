package main

import "snapshot/hxrt"

func discard(callback *hxrt.Function[func(int) int]) *hxrt.Function[func(int)] {
	return func(hx_callback_source_2 *hxrt.Function[func(int) int]) *hxrt.Function[func(int)] {
		return hxrt.AdaptFunction(hx_callback_source_2, func(hx_erased_callback_arg_1 int) {
			hxrt.FunctionCall(hx_callback_source_2)(hx_erased_callback_arg_1)
		})
	}(callback)
}

func main() {
	original := hxrt.NewFunction(func(value int) int {
		hxrt.Println(any(value))
		return value
	})
	returned := discard(original)
	stored := func(hx_callback_source_4 *hxrt.Function[func(int) int]) *hxrt.Function[func(int)] {
		return hxrt.AdaptFunction(hx_callback_source_4, func(hx_erased_callback_arg_3 int) {
			hxrt.FunctionCall(hx_callback_source_4)(hx_erased_callback_arg_3)
		})
	}(original)
	func() {
		var hx_callable_5 *hxrt.Function[func(int)] = returned
		var hx_argument_6 int = 3
		hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()
	func() {
		var hx_callable_7 *hxrt.Function[func(int)] = stored
		var hx_argument_8 int = 5
		hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
	}()
	var v any = any(hxrt.ReflectCompareMethods(any(original), any(returned)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.ReflectCompareMethods(any(original), any(stored)))
	hxrt.Println(v_1)
	holder := New__Main__Holder(func(hx_callback_source_10 *hxrt.Function[func(int) int]) *hxrt.Function[func(int)] {
		return hxrt.AdaptFunction(hx_callback_source_10, func(hx_erased_callback_arg_9 int) {
			hxrt.FunctionCall(hx_callback_source_10)(hx_erased_callback_arg_9)
		})
	}(original))
	func() {
		var hx_callable_11 *hxrt.Function[func(int)] = holder.callback
		var hx_argument_12 int = 7
		hxrt.FunctionCall(hx_callable_11)(hx_argument_12)
	}()
	holder.callback = func(hx_callback_source_14 *hxrt.Function[func(int) int]) *hxrt.Function[func(int)] {
		return hxrt.AdaptFunction(hx_callback_source_14, func(hx_erased_callback_arg_13 int) {
			hxrt.FunctionCall(hx_callback_source_14)(hx_erased_callback_arg_13)
		})
	}(original)
	func() {
		var hx_callable_15 *hxrt.Function[func(int)] = holder.callback
		var hx_argument_16 int = 8
		hxrt.FunctionCall(hx_callable_15)(hx_argument_16)
	}()
	choose := true
	var hx_if_19 *hxrt.Function[func(int)]
	if choose {
		hx_if_19 = func(hx_callback_source_18 *hxrt.Function[func(int) int]) *hxrt.Function[func(int)] {
			return hxrt.AdaptFunction(hx_callback_source_18, func(hx_erased_callback_arg_17 int) {
				hxrt.FunctionCall(hx_callback_source_18)(hx_erased_callback_arg_17)
			})
		}(original)
	} else {
		hx_if_19 = returned
	}
	branch := hx_if_19
	func() {
		var hx_callable_20 *hxrt.Function[func(int)] = branch
		var hx_argument_21 int = 9
		hxrt.FunctionCall(hx_callable_20)(hx_argument_21)
	}()
	var v_2 any = any(func() bool {
		var f2 any = any(holder.callback)
		return hxrt.ReflectCompareMethods(any(original), f2)
	}())
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.ReflectCompareMethods(any(original), any(branch)))
	hxrt.Println(v_3)
	var absent *hxrt.Function[func(int) int] = nil
	absentDiscarded := func(hx_callback_source_23 *hxrt.Function[func(int) int]) *hxrt.Function[func(int)] {
		return hxrt.AdaptFunction(hx_callback_source_23, func(hx_erased_callback_arg_22 int) {
			hxrt.FunctionCall(hx_callback_source_23)(hx_erased_callback_arg_22)
		})
	}(absent)
	hxrt.Println(any((absentDiscarded == nil)))
}

type I__Main__Holder interface {
}

type _Main__Holder struct {
	__hx_this I__Main__Holder
	callback  *hxrt.Function[func(int)]
}

func (self *_Main__Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New__Main__Holder(callback *hxrt.Function[func(int)]) *_Main__Holder {
	self := &_Main__Holder{}
	self.__hx_this = self
	self.callback = callback
	return self
}
