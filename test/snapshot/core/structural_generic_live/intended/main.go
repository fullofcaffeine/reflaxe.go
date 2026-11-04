package main

import "snapshot/hxrt"

type I_Holder interface {
	read() hxrt.Anonymous
}

type Holder struct {
	__hx_this I_Holder
	value     hxrt.Anonymous
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder(value hxrt.Anonymous) *Holder {
	self := &Holder{}
	self.__hx_this = self
	self.value = value
	return self
}

func (self *Holder) read() hxrt.Anonymous {
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
	holder := New_Holder(func(hx_structural_origin_2 hxrt.Anonymous) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_origin_2, func(name string) hxrt.AnonymousField {
			switch name {
			case "next":
				return hxrt.AnonymousField{Value: func(hx_callback_source_6 *hxrt.Function[func() int]) *hxrt.Function[func() any] {
					return hxrt.AdaptFunction(hx_callback_source_6, func() any {
						return hxrt.FunctionCall(hx_callback_source_6)()
					})
				}(func(hx_obj_3 hxrt.Anonymous) *hxrt.Function[func() int] {
					hx_field_4 := hxrt.AnonymousGet(hx_obj_3, "next")
					if hx_field_4 == nil {
						var hx_zero_5 *hxrt.Function[func() int]
						return hx_zero_5
					}
					return hx_field_4.(*hxrt.Function[func() int])
				}(hx_structural_origin_2)), Found: true}
			}
			return hxrt.AnonymousLookup(hx_structural_origin_2, name)
		}, func(name string, incoming any) bool {
			switch name {
			case "next":
				hxrt.AnonymousSet(hx_structural_origin_2, "next", func(hx_callback_source_9 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
					return hxrt.AdaptFunction(hx_callback_source_9, func() int {
						return func(hx_value_10 any) int {
							if hx_value_10 == nil {
								var hx_zero_11 int
								return hx_zero_11
							}
							return hx_value_10.(int)
						}(hxrt.FunctionCall(hx_callback_source_9)())
					})
				}(func(hx_value_7 any) *hxrt.Function[func() any] {
					if hx_value_7 == nil {
						var hx_zero_8 *hxrt.Function[func() any]
						return hx_zero_8
					}
					return hx_value_7.(*hxrt.Function[func() any])
				}(incoming)))
				return true
			}
			hxrt.AnonymousSet(hx_structural_origin_2, name, incoming)
			return true
		})
	}(source))
	view := func(hx_structural_origin_12 hxrt.Anonymous) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_origin_12, func(name string) hxrt.AnonymousField {
			switch name {
			case "next":
				return hxrt.AnonymousField{Value: func(hx_callback_source_16 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
					return hxrt.AdaptFunction(hx_callback_source_16, func() int {
						return func(hx_value_17 any) int {
							if hx_value_17 == nil {
								var hx_zero_18 int
								return hx_zero_18
							}
							return hx_value_17.(int)
						}(hxrt.FunctionCall(hx_callback_source_16)())
					})
				}(func(hx_obj_13 hxrt.Anonymous) *hxrt.Function[func() any] {
					hx_field_14 := hxrt.AnonymousGet(hx_obj_13, "next")
					if hx_field_14 == nil {
						var hx_zero_15 *hxrt.Function[func() any]
						return hx_zero_15
					}
					return hx_field_14.(*hxrt.Function[func() any])
				}(hx_structural_origin_12)), Found: true}
			}
			return hxrt.AnonymousLookup(hx_structural_origin_12, name)
		}, func(name string, incoming any) bool {
			switch name {
			case "next":
				hxrt.AnonymousSet(hx_structural_origin_12, "next", func(hx_callback_source_21 *hxrt.Function[func() int]) *hxrt.Function[func() any] {
					return hxrt.AdaptFunction(hx_callback_source_21, func() any {
						return hxrt.FunctionCall(hx_callback_source_21)()
					})
				}(func(hx_value_19 any) *hxrt.Function[func() int] {
					if hx_value_19 == nil {
						var hx_zero_20 *hxrt.Function[func() int]
						return hx_zero_20
					}
					return hx_value_19.(*hxrt.Function[func() int])
				}(incoming)))
				return true
			}
			hxrt.AnonymousSet(hx_structural_origin_12, name, incoming)
			return true
		})
	}(holder.read())
	saved := func(hx_obj_22 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_23 := hxrt.AnonymousGet(hx_obj_22, "next")
		if hx_field_23 == nil {
			var hx_zero_24 *hxrt.Function[func() int]
			return hx_zero_24
		}
		return hx_field_23.(*hxrt.Function[func() int])
	}(view)
	var v any = any(same(view, source))
	hxrt.Println(v)
	hxrt.Println(any(hxrt.HaxeEqual(saved, func(hx_obj_25 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_26 := hxrt.AnonymousGet(hx_obj_25, "next")
		if hx_field_26 == nil {
			var hx_zero_27 *hxrt.Function[func() int]
			return hx_zero_27
		}
		return hx_field_26.(*hxrt.Function[func() int])
	}(view))))
	var v_1 any = any(func() int {
		var hx_callable_28 *hxrt.Function[func() int] = saved
		return hxrt.FunctionCall(hx_callable_28)()
	}())
	hxrt.Println(v_1)
	func() *hxrt.Function[func() int] {
		hx_structural_target_29 := source
		hx_structural_value_31 := hxrt.NewFunction(func() int {
			return 8
		})
		hxrt.AnonymousSet(hx_structural_target_29, "next", hx_structural_value_31)
		return hx_structural_value_31
	}()
	var v_2 any = any(func() int {
		var hx_callable_35 *hxrt.Function[func() int] = func(hx_obj_32 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_33 := hxrt.AnonymousGet(hx_obj_32, "next")
			if hx_field_33 == nil {
				var hx_zero_34 *hxrt.Function[func() int]
				return hx_zero_34
			}
			return hx_field_33.(*hxrt.Function[func() int])
		}(view)
		return hxrt.FunctionCall(hx_callable_35)()
	}())
	hxrt.Println(v_2)
	var v_3 any = any(func() int {
		var hx_callable_36 *hxrt.Function[func() int] = saved
		return hxrt.FunctionCall(hx_callable_36)()
	}())
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.HaxeEqual(func(hx_obj_37 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_38 := hxrt.AnonymousGet(hx_obj_37, "next")
		if hx_field_38 == nil {
			var hx_zero_39 *hxrt.Function[func() int]
			return hx_zero_39
		}
		return hx_field_38.(*hxrt.Function[func() int])
	}(view), func(hx_obj_40 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_41 := hxrt.AnonymousGet(hx_obj_40, "next")
		if hx_field_41 == nil {
			var hx_zero_42 *hxrt.Function[func() int]
			return hx_zero_42
		}
		return hx_field_41.(*hxrt.Function[func() int])
	}(source)))
	hxrt.Println(v_4)
	var missing hxrt.Anonymous = nil
	var v_5 any = any((func(hx_structural_origin_53 hxrt.Anonymous) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_origin_53, func(name string) hxrt.AnonymousField {
			switch name {
			case "next":
				return hxrt.AnonymousField{Value: func(hx_callback_source_57 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
					return hxrt.AdaptFunction(hx_callback_source_57, func() int {
						return func(hx_value_58 any) int {
							if hx_value_58 == nil {
								var hx_zero_59 int
								return hx_zero_59
							}
							return hx_value_58.(int)
						}(hxrt.FunctionCall(hx_callback_source_57)())
					})
				}(func(hx_obj_54 hxrt.Anonymous) *hxrt.Function[func() any] {
					hx_field_55 := hxrt.AnonymousGet(hx_obj_54, "next")
					if hx_field_55 == nil {
						var hx_zero_56 *hxrt.Function[func() any]
						return hx_zero_56
					}
					return hx_field_55.(*hxrt.Function[func() any])
				}(hx_structural_origin_53)), Found: true}
			}
			return hxrt.AnonymousLookup(hx_structural_origin_53, name)
		}, func(name string, incoming any) bool {
			switch name {
			case "next":
				hxrt.AnonymousSet(hx_structural_origin_53, "next", func(hx_callback_source_62 *hxrt.Function[func() int]) *hxrt.Function[func() any] {
					return hxrt.AdaptFunction(hx_callback_source_62, func() any {
						return hxrt.FunctionCall(hx_callback_source_62)()
					})
				}(func(hx_value_60 any) *hxrt.Function[func() int] {
					if hx_value_60 == nil {
						var hx_zero_61 *hxrt.Function[func() int]
						return hx_zero_61
					}
					return hx_value_60.(*hxrt.Function[func() int])
				}(incoming)))
				return true
			}
			hxrt.AnonymousSet(hx_structural_origin_53, name, incoming)
			return true
		})
	}(New_Holder(func(hx_structural_origin_43 hxrt.Anonymous) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_origin_43, func(name string) hxrt.AnonymousField {
			switch name {
			case "next":
				return hxrt.AnonymousField{Value: func(hx_callback_source_47 *hxrt.Function[func() int]) *hxrt.Function[func() any] {
					return hxrt.AdaptFunction(hx_callback_source_47, func() any {
						return hxrt.FunctionCall(hx_callback_source_47)()
					})
				}(func(hx_obj_44 hxrt.Anonymous) *hxrt.Function[func() int] {
					hx_field_45 := hxrt.AnonymousGet(hx_obj_44, "next")
					if hx_field_45 == nil {
						var hx_zero_46 *hxrt.Function[func() int]
						return hx_zero_46
					}
					return hx_field_45.(*hxrt.Function[func() int])
				}(hx_structural_origin_43)), Found: true}
			}
			return hxrt.AnonymousLookup(hx_structural_origin_43, name)
		}, func(name string, incoming any) bool {
			switch name {
			case "next":
				hxrt.AnonymousSet(hx_structural_origin_43, "next", func(hx_callback_source_50 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
					return hxrt.AdaptFunction(hx_callback_source_50, func() int {
						return func(hx_value_51 any) int {
							if hx_value_51 == nil {
								var hx_zero_52 int
								return hx_zero_52
							}
							return hx_value_51.(int)
						}(hxrt.FunctionCall(hx_callback_source_50)())
					})
				}(func(hx_value_48 any) *hxrt.Function[func() any] {
					if hx_value_48 == nil {
						var hx_zero_49 *hxrt.Function[func() any]
						return hx_zero_49
					}
					return hx_value_48.(*hxrt.Function[func() any])
				}(incoming)))
				return true
			}
			hxrt.AnonymousSet(hx_structural_origin_43, name, incoming)
			return true
		})
	}(missing)).read()) == nil))
	hxrt.Println(v_5)
	hx_obj_63 := map[string]any{}
	hx_obj_63["hasNext"] = hxrt.NewFunction(func() bool {
		return true
	})
	hx_obj_63["next"] = hxrt.NewFunction(func() any {
		return nil
	})
	nullable := hxrt.AnonymousFromMap(hx_obj_63)
	nullableView := New_Holder(nullable).read()
	var v_6 any = any((func() any {
		var hx_callable_67 *hxrt.Function[func() any] = func(hx_obj_64 hxrt.Anonymous) *hxrt.Function[func() any] {
			hx_field_65 := hxrt.AnonymousGet(hx_obj_64, "next")
			if hx_field_65 == nil {
				var hx_zero_66 *hxrt.Function[func() any]
				return hx_zero_66
			}
			return hx_field_65.(*hxrt.Function[func() any])
		}(nullableView)
		return hxrt.FunctionCall(hx_callable_67)()
	}() == nil))
	hxrt.Println(v_6)
	func() *hxrt.Function[func() any] {
		hx_structural_target_68 := nullable
		hx_structural_value_70 := hxrt.NewFunction(func() any {
			return 5
		})
		hxrt.AnonymousSet(hx_structural_target_68, "next", hx_structural_value_70)
		return hx_structural_value_70
	}()
	var v_7 any = any(func() any {
		var hx_callable_74 *hxrt.Function[func() any] = func(hx_obj_71 hxrt.Anonymous) *hxrt.Function[func() any] {
			hx_field_72 := hxrt.AnonymousGet(hx_obj_71, "next")
			if hx_field_72 == nil {
				var hx_zero_73 *hxrt.Function[func() any]
				return hx_zero_73
			}
			return hx_field_72.(*hxrt.Function[func() any])
		}(nullableView)
		return hxrt.FunctionCall(hx_callable_74)()
	}())
	hxrt.Println(v_7)
}

func same(left any, right any) bool {
	return hxrt.HaxeEqual(left, right)
}
