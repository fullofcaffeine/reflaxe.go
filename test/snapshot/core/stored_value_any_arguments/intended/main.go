package main

import "snapshot/hxrt"

type I_Defaults interface {
	increment(value int) int
	enabled(value bool) bool
}

type Defaults struct {
	__hx_this I_Defaults
}

func (self *Defaults) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Defaults() *Defaults {
	return hx_init_Defaults(hx_alloc_Defaults(nil))
}

func hx_alloc_Defaults(owner any) *Defaults {
	self := &Defaults{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Defaults)
	}
	return self
}

func hx_init_Defaults(self *Defaults) *Defaults {
	return self
}

func (self *Defaults) increment(value int) int {
	return int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(1)))
}

func (self *Defaults) enabled(value bool) bool {
	return value
}

func main() {
	defaults := New_Defaults()
	hx_callable_4 := func() *hxrt.Function[func(any) int] {
		hx_receiver_1 := defaults
		return hxrt.BoundFunction(func(value any) int {
			if value == nil {
				value = 5
			}
			return hx_receiver_1.__hx_this.increment(func(hx_value_2 any) int {
				if hx_value_2 == nil {
					var hx_zero_3 int
					return hx_zero_3
				}
				return hx_value_2.(int)
			}(value))
		}, func() any {
			return hx_receiver_1.__hx_this
		}, "increment")
	}()
	hx_normalized_callable_5 := hxrt.AdaptFunction(hx_callable_4, func(value any) int {
		if value == nil {
			value = 5
		}
		return hxrt.FunctionCall(hx_callable_4)(hxrt.IntFromNullableAny(value))
	})
	var increment *hxrt.Function[func(any) int] = hx_normalized_callable_5
	hx_callable_9 := func() *hxrt.Function[func(any) bool] {
		hx_receiver_6 := defaults
		return hxrt.BoundFunction(func(value any) bool {
			if value == nil {
				value = true
			}
			return hx_receiver_6.__hx_this.enabled(func(hx_value_7 any) bool {
				if hx_value_7 == nil {
					var hx_zero_8 bool
					return hx_zero_8
				}
				return hx_value_7.(bool)
			}(value))
		}, func() any {
			return hx_receiver_6.__hx_this
		}, "enabled")
	}()
	hx_normalized_callable_12 := hxrt.AdaptFunction(hx_callable_9, func(value any) bool {
		if value == nil {
			value = true
		}
		return hxrt.FunctionCall(hx_callable_9)(func(hx_value_10 any) bool {
			if hx_value_10 == nil {
				var hx_zero_11 bool
				return hx_zero_11
			}
			return hx_value_10.(bool)
		}(value))
	})
	var enabled *hxrt.Function[func(any) bool] = hx_normalized_callable_12
	var v any = any(func() int {
		var hx_callable_13 *hxrt.Function[func(any) int] = increment
		var hx_argument_14 any = nil
		return hxrt.FunctionCall(hx_callable_13)(hx_argument_14)
	}())
	hxrt.Println(v)
	var v_1 any = any(func() int {
		var hx_callable_15 *hxrt.Function[func(any) int] = increment
		var hx_argument_16 any = 8
		return hxrt.FunctionCall(hx_callable_15)(hx_argument_16)
	}())
	hxrt.Println(v_1)
	var v_2 any = any(func() bool {
		var hx_callable_17 *hxrt.Function[func(any) bool] = enabled
		var hx_argument_18 any = nil
		return hxrt.FunctionCall(hx_callable_17)(hx_argument_18)
	}())
	hxrt.Println(v_2)
	var v_3 any = any(func() bool {
		var hx_callable_19 *hxrt.Function[func(any) bool] = enabled
		var hx_argument_20 any = false
		return hxrt.FunctionCall(hx_callable_19)(hx_argument_20)
	}())
	hxrt.Println(v_3)
	var erased any = hxrt.NewArray(4, 7)
	var v_4 any = any(total(func(hx_value_23 any) *hxrt.Array {
		if hx_value_23 == nil {
			var hx_zero_24 *hxrt.Array
			return hx_zero_24
		}
		return hx_value_23.(*hxrt.Array)
	}(func(hx_value_21 any) *hxrt.Array {
		if hx_value_21 == nil {
			var hx_zero_22 *hxrt.Array
			return hx_zero_22
		}
		return hx_value_21.(*hxrt.Array)
	}(erased))))
	hxrt.Println(v_4)
}

func total(values *hxrt.Array) int {
	return int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(values.Get(0))) + hxrt.Int32Wrap(hxrt.IntFromNullableAny(values.Get(1)))))
}
