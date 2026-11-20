package main

import "snapshot/hxrt"

type I_CounterIterable interface {
	iterator() *CounterIterator
}

type CounterIterable struct {
	__hx_this I_CounterIterable
	data      *hxrt.Array
}

func (self *CounterIterable) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_CounterIterable(data *hxrt.Array) *CounterIterable {
	return hx_init_CounterIterable(hx_alloc_CounterIterable(nil), data)
}

func hx_alloc_CounterIterable(owner any) *CounterIterable {
	self := &CounterIterable{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_CounterIterable)
	}
	return self
}

func hx_init_CounterIterable(self *CounterIterable, data *hxrt.Array) *CounterIterable {
	self.data = data
	return self
}

func (self *CounterIterable) iterator() *CounterIterator {
	return New_CounterIterator(self.data)
}

func (self *CounterIterable) HxClassName() string {
	return "CounterIterable"
}

type I_CounterIterator interface {
	hasNext() bool
	next() int
}

type CounterIterator struct {
	__hx_this I_CounterIterator
	data      *hxrt.Array
	index     int
}

func (self *CounterIterator) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_CounterIterator(data *hxrt.Array) *CounterIterator {
	return hx_init_CounterIterator(hx_alloc_CounterIterator(nil), data)
}

func hx_alloc_CounterIterator(owner any) *CounterIterator {
	self := &CounterIterator{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_CounterIterator)
	}
	return self
}

func hx_init_CounterIterator(self *CounterIterator, data *hxrt.Array) *CounterIterator {
	self.data = data
	self.index = 0
	return self
}

func (self *CounterIterator) hasNext() bool {
	return (self.index < self.data.Len())
}

func (self *CounterIterator) next() int {
	hx_post_1 := self.index
	self.index = int(int32((self.index + 1)))
	return hxrt.IntFromNullableAny(self.data.Get(hx_post_1))
}

func (self *CounterIterator) HxClassName() string {
	return "CounterIterator"
}

func main() {
	values := func(hx_structural_source_2 *CounterIterable) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_2, func(name string) hxrt.AnonymousField {
			switch name {
			case "data":
				return hxrt.AnonymousField{Value: hx_structural_source_2.data, Found: true}
			case "iterator":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *CounterIterator] {
					hx_receiver_3 := hx_structural_source_2
					return hxrt.BoundFunction(func() *CounterIterator {
						return hx_receiver_3.__hx_this.iterator()
					}, func() any {
						return hx_receiver_3.__hx_this
					}, "iterator")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(New_CounterIterable(hxrt.NewArray(1, 2, 3)))
	emptyValues := func(hx_structural_source_4 *CounterIterable) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_4, func(name string) hxrt.AnonymousField {
			switch name {
			case "data":
				return hxrt.AnonymousField{Value: hx_structural_source_4.data, Found: true}
			case "iterator":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *CounterIterator] {
					hx_receiver_5 := hx_structural_source_4
					return hxrt.BoundFunction(func() *CounterIterator {
						return hx_receiver_5.__hx_this.iterator()
					}, func() any {
						return hx_receiver_5.__hx_this
					}, "iterator")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(New_CounterIterable(hxrt.NewArray()))
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("count.values="), Lambda_count(func() hxrt.Anonymous {
		hx_lambda_source_6 := values
		hx_lambda_wrapped_7 := map[string]any{}
		hx_lambda_wrapped_7["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
			hx_lambda_iterator_10 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_6, "iterator")))()
			hx_lambda_iterator_map_8 := map[string]any{}
			hx_lambda_iterator_map_8["hasNext"] = hxrt.NewFunction(func() bool {
				return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_10, "hasNext")))()
			})
			hx_lambda_iterator_map_8["next"] = hxrt.NewFunction(func() any {
				hx_lambda_value_9 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_10, "next")))()
				return hx_lambda_value_9
			})
			return hxrt.AnonymousFromMap(hx_lambda_iterator_map_8)
		})
		return hxrt.AnonymousFromMap(hx_lambda_wrapped_7)
	}(), nil)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("count.empty="), Lambda_count(func() hxrt.Anonymous {
		hx_lambda_source_11 := emptyValues
		hx_lambda_wrapped_12 := map[string]any{}
		hx_lambda_wrapped_12["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
			hx_lambda_iterator_15 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_11, "iterator")))()
			hx_lambda_iterator_map_13 := map[string]any{}
			hx_lambda_iterator_map_13["hasNext"] = hxrt.NewFunction(func() bool {
				return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_15, "hasNext")))()
			})
			hx_lambda_iterator_map_13["next"] = hxrt.NewFunction(func() any {
				hx_lambda_value_14 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_15, "next")))()
				return hx_lambda_value_14
			})
			return hxrt.AnonymousFromMap(hx_lambda_iterator_map_13)
		})
		return hxrt.AnonymousFromMap(hx_lambda_wrapped_12)
	}(), nil)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("count.values.odd="), Lambda_count(func() hxrt.Anonymous {
		hx_lambda_source_20 := values
		hx_lambda_wrapped_21 := map[string]any{}
		hx_lambda_wrapped_21["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
			hx_lambda_iterator_24 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_20, "iterator")))()
			hx_lambda_iterator_map_22 := map[string]any{}
			hx_lambda_iterator_map_22["hasNext"] = hxrt.NewFunction(func() bool {
				return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_24, "hasNext")))()
			})
			hx_lambda_iterator_map_22["next"] = hxrt.NewFunction(func() any {
				hx_lambda_value_23 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_24, "next")))()
				return hx_lambda_value_23
			})
			return hxrt.AnonymousFromMap(hx_lambda_iterator_map_22)
		})
		return hxrt.AnonymousFromMap(hx_lambda_wrapped_21)
	}(), func(hx_lambda_callback_19 *hxrt.Function[func(int) bool]) *hxrt.Function[func(any) bool] {
		return hxrt.AdaptFunction(hx_lambda_callback_19, func(hx_lambda_arg_16 any) bool {
			return hxrt.FunctionCall(hx_lambda_callback_19)(func(hx_value_17 any) int {
				if hx_value_17 == nil {
					var hx_zero_18 int
					return hx_zero_18
				}
				return hx_value_17.(int)
			}(hx_lambda_arg_16))
		})
	}(hxrt.NewFunction(func(v_3 int) bool {
		return (int((hxrt.Int32Wrap(v_3) % hxrt.Int32Wrap(2))) == 1)
	})))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("empty.values="), hxrt.StdString(Lambda_empty(func() hxrt.Anonymous {
		hx_lambda_source_25 := values
		hx_lambda_wrapped_26 := map[string]any{}
		hx_lambda_wrapped_26["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
			hx_lambda_iterator_29 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_25, "iterator")))()
			hx_lambda_iterator_map_27 := map[string]any{}
			hx_lambda_iterator_map_27["hasNext"] = hxrt.NewFunction(func() bool {
				return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_29, "hasNext")))()
			})
			hx_lambda_iterator_map_27["next"] = hxrt.NewFunction(func() any {
				hx_lambda_value_28 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_29, "next")))()
				return hx_lambda_value_28
			})
			return hxrt.AnonymousFromMap(hx_lambda_iterator_map_27)
		})
		return hxrt.AnonymousFromMap(hx_lambda_wrapped_26)
	}()))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("empty.empty="), hxrt.StdString(Lambda_empty(func() hxrt.Anonymous {
		hx_lambda_source_30 := emptyValues
		hx_lambda_wrapped_31 := map[string]any{}
		hx_lambda_wrapped_31["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
			hx_lambda_iterator_34 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_30, "iterator")))()
			hx_lambda_iterator_map_32 := map[string]any{}
			hx_lambda_iterator_map_32["hasNext"] = hxrt.NewFunction(func() bool {
				return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_34, "hasNext")))()
			})
			hx_lambda_iterator_map_32["next"] = hxrt.NewFunction(func() any {
				hx_lambda_value_33 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_34, "next")))()
				return hx_lambda_value_33
			})
			return hxrt.AnonymousFromMap(hx_lambda_iterator_map_32)
		})
		return hxrt.AnonymousFromMap(hx_lambda_wrapped_31)
	}()))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("exists.values.gt2="), hxrt.StdString(Lambda_exists(func() hxrt.Anonymous {
		hx_lambda_source_39 := values
		hx_lambda_wrapped_40 := map[string]any{}
		hx_lambda_wrapped_40["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
			hx_lambda_iterator_43 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_39, "iterator")))()
			hx_lambda_iterator_map_41 := map[string]any{}
			hx_lambda_iterator_map_41["hasNext"] = hxrt.NewFunction(func() bool {
				return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_43, "hasNext")))()
			})
			hx_lambda_iterator_map_41["next"] = hxrt.NewFunction(func() any {
				hx_lambda_value_42 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_43, "next")))()
				return hx_lambda_value_42
			})
			return hxrt.AnonymousFromMap(hx_lambda_iterator_map_41)
		})
		return hxrt.AnonymousFromMap(hx_lambda_wrapped_40)
	}(), func(hx_lambda_callback_38 *hxrt.Function[func(int) bool]) *hxrt.Function[func(any) bool] {
		return hxrt.AdaptFunction(hx_lambda_callback_38, func(hx_lambda_arg_35 any) bool {
			return hxrt.FunctionCall(hx_lambda_callback_38)(func(hx_value_36 any) int {
				if hx_value_36 == nil {
					var hx_zero_37 int
					return hx_zero_37
				}
				return hx_value_36.(int)
			}(hx_lambda_arg_35))
		})
	}(hxrt.NewFunction(func(v_6 int) bool {
		return (v_6 > 2)
	}))))))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("exists.values.gt10="), hxrt.StdString(Lambda_exists(func() hxrt.Anonymous {
		hx_lambda_source_48 := values
		hx_lambda_wrapped_49 := map[string]any{}
		hx_lambda_wrapped_49["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
			hx_lambda_iterator_52 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_48, "iterator")))()
			hx_lambda_iterator_map_50 := map[string]any{}
			hx_lambda_iterator_map_50["hasNext"] = hxrt.NewFunction(func() bool {
				return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_52, "hasNext")))()
			})
			hx_lambda_iterator_map_50["next"] = hxrt.NewFunction(func() any {
				hx_lambda_value_51 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_52, "next")))()
				return hx_lambda_value_51
			})
			return hxrt.AnonymousFromMap(hx_lambda_iterator_map_50)
		})
		return hxrt.AnonymousFromMap(hx_lambda_wrapped_49)
	}(), func(hx_lambda_callback_47 *hxrt.Function[func(int) bool]) *hxrt.Function[func(any) bool] {
		return hxrt.AdaptFunction(hx_lambda_callback_47, func(hx_lambda_arg_44 any) bool {
			return hxrt.FunctionCall(hx_lambda_callback_47)(func(hx_value_45 any) int {
				if hx_value_45 == nil {
					var hx_zero_46 int
					return hx_zero_46
				}
				return hx_value_45.(int)
			}(hx_lambda_arg_44))
		})
	}(hxrt.NewFunction(func(v_7 int) bool {
		return (v_7 > 10)
	}))))))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("has.values.2="), hxrt.StdString(Lambda_has(func() hxrt.Anonymous {
		hx_lambda_source_53 := values
		hx_lambda_wrapped_54 := map[string]any{}
		hx_lambda_wrapped_54["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
			hx_lambda_iterator_57 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_53, "iterator")))()
			hx_lambda_iterator_map_55 := map[string]any{}
			hx_lambda_iterator_map_55["hasNext"] = hxrt.NewFunction(func() bool {
				return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_57, "hasNext")))()
			})
			hx_lambda_iterator_map_55["next"] = hxrt.NewFunction(func() any {
				hx_lambda_value_56 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_57, "next")))()
				return hx_lambda_value_56
			})
			return hxrt.AnonymousFromMap(hx_lambda_iterator_map_55)
		})
		return hxrt.AnonymousFromMap(hx_lambda_wrapped_54)
	}(), 2))))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("has.values.9="), hxrt.StdString(Lambda_has(func() hxrt.Anonymous {
		hx_lambda_source_58 := values
		hx_lambda_wrapped_59 := map[string]any{}
		hx_lambda_wrapped_59["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
			hx_lambda_iterator_62 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_58, "iterator")))()
			hx_lambda_iterator_map_60 := map[string]any{}
			hx_lambda_iterator_map_60["hasNext"] = hxrt.NewFunction(func() bool {
				return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_62, "hasNext")))()
			})
			hx_lambda_iterator_map_60["next"] = hxrt.NewFunction(func() any {
				hx_lambda_value_61 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_62, "next")))()
				return hx_lambda_value_61
			})
			return hxrt.AnonymousFromMap(hx_lambda_iterator_map_60)
		})
		return hxrt.AnonymousFromMap(hx_lambda_wrapped_59)
	}(), 9))))
	hxrt.Println(v_8)
	filtered := Lambda_filter(func() hxrt.Anonymous {
		hx_lambda_source_67 := values
		hx_lambda_wrapped_68 := map[string]any{}
		hx_lambda_wrapped_68["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
			hx_lambda_iterator_71 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_67, "iterator")))()
			hx_lambda_iterator_map_69 := map[string]any{}
			hx_lambda_iterator_map_69["hasNext"] = hxrt.NewFunction(func() bool {
				return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_71, "hasNext")))()
			})
			hx_lambda_iterator_map_69["next"] = hxrt.NewFunction(func() any {
				hx_lambda_value_70 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_71, "next")))()
				return hx_lambda_value_70
			})
			return hxrt.AnonymousFromMap(hx_lambda_iterator_map_69)
		})
		return hxrt.AnonymousFromMap(hx_lambda_wrapped_68)
	}(), func(hx_lambda_callback_66 *hxrt.Function[func(int) bool]) *hxrt.Function[func(any) bool] {
		return hxrt.AdaptFunction(hx_lambda_callback_66, func(hx_lambda_arg_63 any) bool {
			return hxrt.FunctionCall(hx_lambda_callback_66)(func(hx_value_64 any) int {
				if hx_value_64 == nil {
					var hx_zero_65 int
					return hx_zero_65
				}
				return hx_value_64.(int)
			}(hx_lambda_arg_63))
		})
	}(hxrt.NewFunction(func(v_9 int) bool {
		return (int((hxrt.Int32Wrap(v_9) % hxrt.Int32Wrap(2))) == 1)
	})))
	mapFn := hx_recover_callable_6824eed0(hxrt.NamedFunction(Lambda_map_, "Lambda.map"))
	mapped := func() *hxrt.Array {
		var hx_callable_77 *hxrt.Function[func(hxrt.Anonymous, *hxrt.Function[func(int) int]) *hxrt.Array] = mapFn
		var hx_argument_78 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_lambda_source_72 := values
			hx_lambda_wrapped_73 := map[string]any{}
			hx_lambda_wrapped_73["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
				hx_lambda_iterator_76 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_72, "iterator")))()
				hx_lambda_iterator_map_74 := map[string]any{}
				hx_lambda_iterator_map_74["hasNext"] = hxrt.NewFunction(func() bool {
					return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_76, "hasNext")))()
				})
				hx_lambda_iterator_map_74["next"] = hxrt.NewFunction(func() any {
					hx_lambda_value_75 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_76, "next")))()
					return hx_lambda_value_75
				})
				return hxrt.AnonymousFromMap(hx_lambda_iterator_map_74)
			})
			return hxrt.AnonymousFromMap(hx_lambda_wrapped_73)
		}()
		var hx_argument_79 *hxrt.Function[func(int) int] = hxrt.NewFunction(func(v_9 int) int {
			return int((hxrt.Int32Wrap(v_9) * hxrt.Int32Wrap(10)))
		})
		return hxrt.FunctionCall(hx_callable_77)(hx_argument_78, hx_argument_79)
	}()
	mappedArray := func() *hxrt.Array {
		var hx_callable_85 *hxrt.Function[func(hxrt.Anonymous, *hxrt.Function[func(int) int]) *hxrt.Array] = mapFn
		var hx_argument_86 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_lambda_source_80 := hxrt.NewArray(4, 5)
			hx_lambda_wrapped_81 := map[string]any{}
			hx_lambda_wrapped_81["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
				hx_lambda_index_84 := 0
				hx_lambda_iterator_map_82 := map[string]any{}
				hx_lambda_iterator_map_82["hasNext"] = hxrt.NewFunction(func() bool {
					return (hx_lambda_index_84 < hx_lambda_source_80.Len())
				})
				hx_lambda_iterator_map_82["next"] = hxrt.NewFunction(func() any {
					hx_lambda_value_83 := hx_lambda_source_80.Get(hx_lambda_index_84)
					hx_lambda_index_84 = (hx_lambda_index_84 + 1)
					return hx_lambda_value_83
				})
				return hxrt.AnonymousFromMap(hx_lambda_iterator_map_82)
			})
			return hxrt.AnonymousFromMap(hx_lambda_wrapped_81)
		}()
		var hx_argument_87 *hxrt.Function[func(int) int] = hxrt.NewFunction(func(v_9 int) int {
			return int((hxrt.Int32Wrap(v_9) + hxrt.Int32Wrap(1)))
		})
		return hxrt.FunctionCall(hx_callable_85)(hx_argument_86, hx_argument_87)
	}()
	folded := func(hx_value_100 any) int {
		if hx_value_100 == nil {
			var hx_zero_101 int
			return hx_zero_101
		}
		return hx_value_100.(int)
	}(Lambda_fold(func() hxrt.Anonymous {
		hx_lambda_source_95 := values
		hx_lambda_wrapped_96 := map[string]any{}
		hx_lambda_wrapped_96["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
			hx_lambda_iterator_99 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_95, "iterator")))()
			hx_lambda_iterator_map_97 := map[string]any{}
			hx_lambda_iterator_map_97["hasNext"] = hxrt.NewFunction(func() bool {
				return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_99, "hasNext")))()
			})
			hx_lambda_iterator_map_97["next"] = hxrt.NewFunction(func() any {
				hx_lambda_value_98 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_99, "next")))()
				return hx_lambda_value_98
			})
			return hxrt.AnonymousFromMap(hx_lambda_iterator_map_97)
		})
		return hxrt.AnonymousFromMap(hx_lambda_wrapped_96)
	}(), func(hx_lambda_callback_94 *hxrt.Function[func(int, int) int]) *hxrt.Function[func(any, any) any] {
		return hxrt.AdaptFunction(hx_lambda_callback_94, func(hx_lambda_value_88 any, hx_lambda_acc_89 any) any {
			return hxrt.FunctionCall(hx_lambda_callback_94)(func(hx_value_90 any) int {
				if hx_value_90 == nil {
					var hx_zero_91 int
					return hx_zero_91
				}
				return hx_value_90.(int)
			}(hx_lambda_value_88), func(hx_value_92 any) int {
				if hx_value_92 == nil {
					var hx_zero_93 int
					return hx_zero_93
				}
				return hx_value_92.(int)
			}(hx_lambda_acc_89))
		})
	}(hxrt.NewFunction(func(v_9 int, acc int) int {
		return int((hxrt.Int32Wrap(acc) + hxrt.Int32Wrap(v_9)))
	})), 0))
	foldFn := hx_recover_callable_bd34b669(hxrt.NamedFunction(Lambda_fold, "Lambda.fold"))
	foldedByFunction := func() int {
		var hx_callable_107 *hxrt.Function[func(hxrt.Anonymous, *hxrt.Function[func(int, int) int], int) int] = foldFn
		var hx_argument_108 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_lambda_source_102 := values
			hx_lambda_wrapped_103 := map[string]any{}
			hx_lambda_wrapped_103["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
				hx_lambda_iterator_106 := hxrt.FunctionCall(hx_recover_callable_7838dc94(hxrt.AnonymousGet(hx_lambda_source_102, "iterator")))()
				hx_lambda_iterator_map_104 := map[string]any{}
				hx_lambda_iterator_map_104["hasNext"] = hxrt.NewFunction(func() bool {
					return hxrt.FunctionCall(hx_recover_callable_cf6928b1(hxrt.AnonymousGet(hx_lambda_iterator_106, "hasNext")))()
				})
				hx_lambda_iterator_map_104["next"] = hxrt.NewFunction(func() any {
					hx_lambda_value_105 := hxrt.FunctionCall(hx_recover_callable_e273723c(hxrt.AnonymousGet(hx_lambda_iterator_106, "next")))()
					return hx_lambda_value_105
				})
				return hxrt.AnonymousFromMap(hx_lambda_iterator_map_104)
			})
			return hxrt.AnonymousFromMap(hx_lambda_wrapped_103)
		}()
		var hx_argument_109 *hxrt.Function[func(int, int) int] = hxrt.NewFunction(func(v_9 int, acc int) int {
			return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(acc) * hxrt.Int32Wrap(10)))) + hxrt.Int32Wrap(v_9)))
		})
		var hx_argument_110 int = 0
		return hxrt.FunctionCall(hx_callable_107)(hx_argument_108, hx_argument_109, hx_argument_110)
	}()
	out := hxrt.NewArray()
	index := 0
	var _g_current int
	var _g_array *hxrt.Array
	_g_current = 0
	_g_array = hxrt.NewArray(1, 2, 3)
	for _g_current < _g_array.Len() {
		hx_post_111 := _g_current
		_g_current = int(int32((_g_current + 1)))
		value := hxrt.IntFromNullableAny(_g_array.Get(hx_post_111))
		out.Push(int((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(value))))
		index = int(int32((index + 1)))
	}
	mappedWithIndex := out
	var v_9 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("filter.values.length="), filtered.Len()))
	hxrt.Println(v_9)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("filter.values.first="), filtered.Get(0))))
	var v_10 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("filter.values.last="), filtered.Get(int((hxrt.Int32Wrap(filtered.Len()) - hxrt.Int32Wrap(1))))))
	hxrt.Println(v_10)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("map.values.first="), mapped.Get(0))))
	var v_11 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("map.values.last="), mapped.Get(int((hxrt.Int32Wrap(mapped.Len()) - hxrt.Int32Wrap(1))))))
	hxrt.Println(v_11)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("map.array="), mappedArray.Get(0)), hxrt.StringFromLiteral(",")), mappedArray.Get(1))))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("fold.values.sum="), folded)))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("fold.function="), foldedByFunction)))
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mapi.values="), hxrt.StringJoinAny(mappedWithIndex.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_12)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *CounterIterable:
		return func(hx_structural_source_1 *CounterIterable) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "data":
					return hxrt.AnonymousField{Value: hx_structural_source_1.data, Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *CounterIterator] {
						hx_receiver_2 := hx_structural_source_1
						return hxrt.BoundFunction(func() *CounterIterator {
							return hx_receiver_2.__hx_this.iterator()
						}, func() any {
							return hx_receiver_2.__hx_this
						}, "iterator")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *CounterIterator:
		return func(hx_structural_source_3 *CounterIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_3, func(name string) hxrt.AnonymousField {
				switch name {
				case "data":
					return hxrt.AnonymousField{Value: hx_structural_source_3.data, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_3.index, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_6 := hx_structural_source_3
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_6.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_6.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_7 := hx_structural_source_3
						return hxrt.BoundFunction(func() int {
							return hx_receiver_7.__hx_this.next()
						}, func() any {
							return hx_receiver_7.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "index":
					hx_structural_source_3.index = func(hx_value_4 any) int {
						if hx_value_4 == nil {
							var hx_zero_5 int
							return hx_zero_5
						}
						return hx_value_4.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_8 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_8, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_8.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_8.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_8.high = func(hx_value_9 any) int {
						if hx_value_9 == nil {
							var hx_zero_10 int
							return hx_zero_10
						}
						return hx_value_9.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_8.low = func(hx_value_11 any) int {
						if hx_value_11 == nil {
							var hx_zero_12 int
							return hx_zero_12
						}
						return hx_value_11.(int)
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

func hx_recover_callable_4a4f19d9(value any) *hxrt.Function[func(any, any) any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any, any) any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(hxrt.Anonymous, *hxrt.Function[func(any) any]) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 any, arg_1 any) any {
			return hxrt.FunctionCall(source)(hx_structural_recover(arg_0), hx_recover_callable_f4ee5029(arg_1))
		})
	case *hxrt.Function[func(int, int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any, arg_1 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(arg_0), func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(arg_1))
		})
	case *hxrt.Function[func(any, any) any]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_6824eed0(value any) *hxrt.Function[func(hxrt.Anonymous, *hxrt.Function[func(int) int]) *hxrt.Array] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(hxrt.Anonymous, *hxrt.Function[func(int) int]) *hxrt.Array]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(hxrt.Anonymous, *hxrt.Function[func(any) any]) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 *hxrt.Function[func(int) int]) *hxrt.Array {
			return hxrt.FunctionCall(source)(arg_0, hx_recover_callable_f4ee5029(arg_1))
		})
	case *hxrt.Function[func(int, int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 *hxrt.Function[func(int) int]) *hxrt.Array {
			return func(hx_value_5 any) *hxrt.Array {
				if hx_value_5 == nil {
					var hx_zero_6 *hxrt.Array
					return hx_zero_6
				}
				return hx_value_5.(*hxrt.Array)
			}(hxrt.FunctionCall(source)(func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(arg_0), func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(arg_1)))
		})
	case *hxrt.Function[func(any, any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 *hxrt.Function[func(int) int]) *hxrt.Array {
			return func(hx_value_7 any) *hxrt.Array {
				if hx_value_7 == nil {
					var hx_zero_8 *hxrt.Array
					return hx_zero_8
				}
				return hx_value_7.(*hxrt.Array)
			}(hxrt.FunctionCall(source)(arg_0, arg_1))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_7838dc94(value any) *hxrt.Function[func() hxrt.Anonymous] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() hxrt.Anonymous]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() *CounterIterator]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return source
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_bd34b669(value any) *hxrt.Function[func(hxrt.Anonymous, *hxrt.Function[func(int, int) int], int) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(hxrt.Anonymous, *hxrt.Function[func(int, int) int], int) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(hxrt.Anonymous, *hxrt.Function[func(any, any) any], any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 *hxrt.Function[func(int, int) int], arg_2 int) int {
			return func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(hxrt.FunctionCall(source)(arg_0, hx_recover_callable_4a4f19d9(arg_1), arg_2))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_c5309b43(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() *CounterIterator]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_cf6928b1(value any) *hxrt.Function[func() bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() *CounterIterator]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_7 any) bool {
				if hx_value_7 == nil {
					var hx_zero_8 bool
					return hx_zero_8
				}
				return hx_value_7.(bool)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_e273723c(value any) *hxrt.Function[func() int] {
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
	case *hxrt.Function[func() *CounterIterator]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_f4ee5029(value any) *hxrt.Function[func(any) any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any) any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
