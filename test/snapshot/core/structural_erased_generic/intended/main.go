package main

import "snapshot/hxrt"

type I_Holder interface {
	read() hxrt.Anonymous
	erase() any
}

type Holder struct {
	__hx_this I_Holder
	value     hxrt.Anonymous
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder(value hxrt.Anonymous) *Holder {
	return hx_init_Holder(hx_alloc_Holder(nil), value)
}

func hx_alloc_Holder(owner any) *Holder {
	self := &Holder{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Holder)
	}
	return self
}

func hx_init_Holder(self *Holder, value hxrt.Anonymous) *Holder {
	self.value = value
	return self
}

func (self *Holder) read() hxrt.Anonymous {
	return self.value
}

func (self *Holder) erase() any {
	return self.value
}

func main() {
	hx_obj_1 := map[string]any{}
	hx_obj_1["hasNext"] = hxrt.NewFunction(func() bool {
		return true
	})
	hx_obj_1["next"] = hxrt.NewFunction(func() int {
		return 3
	})
	source := hxrt.AnonymousFromMap(hx_obj_1)
	recovered := hx_structural_recover(hx_structural_recover(New_Holder(source).erase()))
	var v any = any(func() int {
		var hx_callable_2 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(recovered, "next"))
		return hxrt.FunctionCall(hx_callable_2)()
	}())
	hxrt.Println(v)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *Holder:
		return func(hx_structural_source_15 *Holder) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_15, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_15.value, Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_16 := hx_structural_source_15
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_16.__hx_this.read()
						}, func() any {
							return hx_receiver_16.__hx_this
						}, "read")
					}(), Found: true}
				case "erase":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_17 := hx_structural_source_15
						return hxrt.BoundFunction(func() any {
							return hx_receiver_17.__hx_this.erase()
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "erase")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "value":
					hx_structural_source_15.value = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_18 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_18, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_18.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_18.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_18.high = func(hx_value_19 any) int {
						if hx_value_19 == nil {
							var hx_zero_20 int
							return hx_zero_20
						}
						return hx_value_19.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_18.low = func(hx_value_21 any) int {
						if hx_value_21 == nil {
							var hx_zero_22 int
							return hx_zero_22
						}
						return hx_value_21.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func() int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_23 any) int {
				if hx_value_23 == nil {
					var hx_zero_24 int
					return hx_zero_24
				}
				return hx_value_23.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_25 any) int {
				if hx_value_25 == nil {
					var hx_zero_26 int
					return hx_zero_26
				}
				return hx_value_25.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_27 any) int {
				if hx_value_27 == nil {
					var hx_zero_28 int
					return hx_zero_28
				}
				return hx_value_27.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
