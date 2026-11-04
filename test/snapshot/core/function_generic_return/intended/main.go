package main

import "snapshot/hxrt"

type I_Holder interface {
	reader() *hxrt.Function[func() any]
	savedReader() *hxrt.Function[func() any]
	echo() *hxrt.Function[func(any) any]
	missing() *hxrt.Function[func() any]
}

type Holder struct {
	__hx_this I_Holder
	value     any
	saved     *hxrt.Function[func() any]
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder(value any) *Holder {
	self := &Holder{}
	self.__hx_this = self
	self.value = value
	self.saved = hxrt.NewFunction(func() any {
		return value
	})
	return self
}

func (self *Holder) reader() *hxrt.Function[func() any] {
	_gthis := self
	return hxrt.NewFunction(func() any {
		return _gthis.value
	})
}

func (self *Holder) savedReader() *hxrt.Function[func() any] {
	return self.saved
}

func (self *Holder) echo() *hxrt.Function[func(any) any] {
	return hxrt.NewFunction(func(value any) any {
		return value
	})
}

func (self *Holder) missing() *hxrt.Function[func() any] {
	return nil
}

func main() {
	read := func(hx_callback_source_1 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
		return hxrt.AdaptFunction(hx_callback_source_1, func() int {
			return func(hx_value_2 any) int {
				if hx_value_2 == nil {
					var hx_zero_3 int
					return hx_zero_3
				}
				return hx_value_2.(int)
			}(hxrt.FunctionCall(hx_callback_source_1)())
		})
	}(New_Holder(19).reader())
	alias := read
	var v any = any(func() int {
		var hx_callable_4 *hxrt.Function[func() int] = read
		return hxrt.FunctionCall(hx_callable_4)()
	}())
	hxrt.Println(v)
	var v_1 any = any(hxrt.ReflectCompareMethods(any(read), any(alias)))
	hxrt.Println(v_1)
	holder := New_Holder(23)
	var v_2 any = any(func() int {
		var hx_callable_9 *hxrt.Function[func(int) int] = func(hx_callback_source_6 *hxrt.Function[func(any) any]) *hxrt.Function[func(int) int] {
			return hxrt.AdaptFunction(hx_callback_source_6, func(hx_erased_callback_arg_5 int) int {
				return func(hx_value_7 any) int {
					if hx_value_7 == nil {
						var hx_zero_8 int
						return hx_zero_8
					}
					return hx_value_7.(int)
				}(hxrt.FunctionCall(hx_callback_source_6)(hx_erased_callback_arg_5))
			})
		}(holder.echo())
		var hx_argument_10 int = 29
		return hxrt.FunctionCall(hx_callable_9)(hx_argument_10)
	}())
	hxrt.Println(v_2)
	var v_3 any = any(func() int {
		var hx_callable_14 *hxrt.Function[func() int] = func(hx_callback_source_11 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
			return hxrt.AdaptFunction(hx_callback_source_11, func() int {
				return func(hx_value_12 any) int {
					if hx_value_12 == nil {
						var hx_zero_13 int
						return hx_zero_13
					}
					return hx_value_12.(int)
				}(hxrt.FunctionCall(hx_callback_source_11)())
			})
		}(holder.savedReader())
		return hxrt.FunctionCall(hx_callable_14)()
	}())
	hxrt.Println(v_3)
	var v_4 any = any(func() bool {
		var f1 any = any(func(hx_callback_source_15 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
			return hxrt.AdaptFunction(hx_callback_source_15, func() int {
				return func(hx_value_16 any) int {
					if hx_value_16 == nil {
						var hx_zero_17 int
						return hx_zero_17
					}
					return hx_value_16.(int)
				}(hxrt.FunctionCall(hx_callback_source_15)())
			})
		}(holder.savedReader()))
		var f2 any = any(func(hx_callback_source_18 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
			return hxrt.AdaptFunction(hx_callback_source_18, func() int {
				return func(hx_value_19 any) int {
					if hx_value_19 == nil {
						var hx_zero_20 int
						return hx_zero_20
					}
					return hx_value_19.(int)
				}(hxrt.FunctionCall(hx_callback_source_18)())
			})
		}(holder.savedReader()))
		return hxrt.ReflectCompareMethods(f1, f2)
	}())
	hxrt.Println(v_4)
	var v_5 any = any(func() bool {
		var f1_1 any = any(func(hx_callback_source_21 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
			return hxrt.AdaptFunction(hx_callback_source_21, func() int {
				return func(hx_value_22 any) int {
					if hx_value_22 == nil {
						var hx_zero_23 int
						return hx_zero_23
					}
					return hx_value_22.(int)
				}(hxrt.FunctionCall(hx_callback_source_21)())
			})
		}(holder.reader()))
		var f2_1 any = any(func(hx_callback_source_24 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
			return hxrt.AdaptFunction(hx_callback_source_24, func() int {
				return func(hx_value_25 any) int {
					if hx_value_25 == nil {
						var hx_zero_26 int
						return hx_zero_26
					}
					return hx_value_25.(int)
				}(hxrt.FunctionCall(hx_callback_source_24)())
			})
		}(holder.reader()))
		return hxrt.ReflectCompareMethods(f1_1, f2_1)
	}())
	hxrt.Println(v_5)
	var v_6 any = any((func(hx_callback_source_27 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
		return hxrt.AdaptFunction(hx_callback_source_27, func() int {
			return func(hx_value_28 any) int {
				if hx_value_28 == nil {
					var hx_zero_29 int
					return hx_zero_29
				}
				return hx_value_28.(int)
			}(hxrt.FunctionCall(hx_callback_source_27)())
		})
	}(holder.missing()) == nil))
	hxrt.Println(v_6)
	var v_7 any = any(func() *string {
		var hx_callable_33 *hxrt.Function[func() *string] = func(hx_callback_source_30 *hxrt.Function[func() any]) *hxrt.Function[func() *string] {
			return hxrt.AdaptFunction(hx_callback_source_30, func() *string {
				return func(hx_value_31 any) *string {
					if hx_value_31 == nil {
						var hx_zero_32 *string
						return hx_zero_32
					}
					return hx_value_31.(*string)
				}(hxrt.FunctionCall(hx_callback_source_30)())
			})
		}(New_Holder(hxrt.StringFromLiteral("text")).reader())
		return hxrt.FunctionCall(hx_callable_33)()
	}())
	hxrt.Println(v_7)
	var v_8 any = any((func() any {
		var hx_callable_34 *hxrt.Function[func() any] = New_Holder(nil).reader()
		return hxrt.FunctionCall(hx_callable_34)()
	}() == nil))
	hxrt.Println(v_8)
}
