package main

import "snapshot/hxrt"

func discard(callback *hxrt.Function[func(int) int]) *hxrt.Function[func(int)] {
	return hx_recover_callable_0(callback)
}

func main() {
	original := hxrt.NewFunction(func(value int) int {
		hxrt.Println(any(value))
		return value
	})
	returned := discard(original)
	stored := hx_recover_callable_0(original)
	func() {
		var hx_callable_1 *hxrt.Function[func(int)] = returned
		var hx_argument_2 int = 3
		hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
	func() {
		var hx_callable_3 *hxrt.Function[func(int)] = stored
		var hx_argument_4 int = 5
		hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}()
	var v any = any(hxrt.ReflectCompareMethods(any(original), any(returned)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.ReflectCompareMethods(any(original), any(stored)))
	hxrt.Println(v_1)
	holder := New__Main__Holder(hx_recover_callable_0(original))
	func() {
		var hx_callable_5 *hxrt.Function[func(int)] = holder.callback
		var hx_argument_6 int = 7
		hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()
	holder.callback = hx_recover_callable_0(original)
	func() {
		var hx_callable_7 *hxrt.Function[func(int)] = holder.callback
		var hx_argument_8 int = 8
		hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
	}()
	choose := true
	var hx_if_9 *hxrt.Function[func(int)]
	if choose {
		hx_if_9 = hx_recover_callable_0(original)
	} else {
		hx_if_9 = returned
	}
	branch := hx_if_9
	func() {
		var hx_callable_10 *hxrt.Function[func(int)] = branch
		var hx_argument_11 int = 9
		hxrt.FunctionCall(hx_callable_10)(hx_argument_11)
	}()
	var v_2 any = any(func() bool {
		var f2 any = any(holder.callback)
		return hxrt.ReflectCompareMethods(any(original), f2)
	}())
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.ReflectCompareMethods(any(original), any(branch)))
	hxrt.Println(v_3)
	var absent *hxrt.Function[func(int) int] = nil
	absentDiscarded := hx_recover_callable_0(absent)
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
	return hx_init__Main__Holder(hx_alloc__Main__Holder(nil), callback)
}

func hx_alloc__Main__Holder(owner any) *_Main__Holder {
	self := &_Main__Holder{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I__Main__Holder)
	}
	return self
}

func hx_init__Main__Holder(self *_Main__Holder, callback *hxrt.Function[func(int)]) *_Main__Holder {
	self.callback = callback
	return self
}

func hx_recover_callable_0(value any) *hxrt.Function[func(int)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(int)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
