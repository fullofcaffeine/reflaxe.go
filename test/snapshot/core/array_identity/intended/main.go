package main

import "snapshot/hxrt"

func applyGenericCallback(values *hxrt.Array, value any, callback *hxrt.Function[func(*hxrt.Array, any)]) {
	func() {
		var hx_callable_1 *hxrt.Function[func(*hxrt.Array, any)] = callback
		var hx_argument_2 *hxrt.Array = values
		var hx_argument_3 any = value
		hxrt.FunctionCall(hx_callable_1)(hx_argument_2, hx_argument_3)
	}()
}

var events *hxrt.Array = hxrt.NewArray()

var holderEvaluations int = 0

func main() {
	local := hxrt.NewArray(1, 2)
	localAlias := local
	localAlias.Push(3)
	func() {
		hx_insert_position_6 := 1
		var hx_insert_value_7 any = 9
		localAlias.Insert(hx_insert_position_6, hx_insert_value_7)
	}()
	func() bool {
		var hx_remove_value_9 any = 2
		hx_remove_index_10 := 0
		for hx_remove_index_10 < localAlias.Len() {
			hx_remove_element_11 := localAlias.Get(hx_remove_index_10)
			if hxrt.HaxeEqual(hx_remove_element_11, hx_remove_value_9) {
				localAlias.RemoveAt(hx_remove_index_10)
				return true
			}
			hx_remove_index_10 = (hx_remove_index_10 + 1)
		}
		return false
	}()
	localAlias.Pop()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("local="), show(local)))
	hxrt.Println(v)
	parameter := hxrt.NewArray(4)
	pushParameter(parameter, 5)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("parameter="), show(parameter)))
	hxrt.Println(v_1)
	returnedSource := hxrt.NewArray(6)
	returnedAlias := returnAlias(returnedSource)
	returnedAlias.Push(7)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("return="), show(returnedSource)))
	hxrt.Println(v_2)
	genericInts := hxrt.NewArray(8)
	pushGeneric(genericInts, 9)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic="), show(genericInts)))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic.has="), hxrt.StdString(Lambda_has(func() hxrt.Anonymous {
		hx_lambda_source_14 := hxrt.NewArray(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"))
		hx_lambda_wrapped_15 := map[string]any{}
		hx_lambda_wrapped_15["iterator"] = hxrt.NewFunction(func() hxrt.Anonymous {
			hx_lambda_index_18 := 0
			hx_lambda_iterator_map_16 := map[string]any{}
			hx_lambda_iterator_map_16["hasNext"] = hxrt.NewFunction(func() bool {
				return (hx_lambda_index_18 < hx_lambda_source_14.Len())
			})
			hx_lambda_iterator_map_16["next"] = hxrt.NewFunction(func() any {
				hx_lambda_value_17 := hx_lambda_source_14.Get(hx_lambda_index_18)
				hx_lambda_index_18 = (hx_lambda_index_18 + 1)
				return hx_lambda_value_17
			})
			return hxrt.AnonymousFromMap(hx_lambda_iterator_map_16)
		})
		return hxrt.AnonymousFromMap(hx_lambda_wrapped_15)
	}(), hxrt.StringFromLiteral("a")))))
	hxrt.Println(v_4)
	stringElements := hxrt.ArrayFromValues(func(hx_sort_src_19 []*string) []any {
		hx_sort_out_21 := make([]any, 0, len(hx_sort_src_19))
		for _, hx_sort_item_20 := range hx_sort_src_19 {
			hx_sort_out_21 = append(hx_sort_out_21, hx_sort_item_20)
		}
		return hx_sort_out_21
	}(hxrt.StringSplitStringPtr(hxrt.StringFromLiteral("left,right"), hxrt.StringFromLiteral(";"))))
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("string.split="), hxrt.ArrayFromValues(func(hx_sort_src_24 []*string) []any {
		hx_sort_out_26 := make([]any, 0, len(hx_sort_src_24))
		for _, hx_sort_item_25 := range hx_sort_src_24 {
			hx_sort_out_26 = append(hx_sort_out_26, hx_sort_item_25)
		}
		return hx_sort_out_26
	}(hxrt.StringSplitStringPtr(func(hx_value_22 any) *string {
		if hx_value_22 == nil {
			var hx_zero_23 *string
			return hx_zero_23
		}
		return hx_value_22.(*string)
	}(stringElements.Get(0)), hxrt.StringFromLiteral(",")))).Len()))
	hxrt.Println(v_5)
	retainedSource := hxrt.NewArray(10)
	retainErased(retainedSource)
	hx_arr_27 := retainedGeneric
	hx_arr_27.Push(11)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("retained="), show(retainedSource)))
	hxrt.Println(v_6)
	callbackSource := hxrt.NewArray(12)
	applyGenericCallback(callbackSource, 13, hx_recover_callable_3(hxrt.NewFunction(func(values *hxrt.Array, value int) {
		values.Push(value)
	})))
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("callback="), show(callbackSource)))
	hxrt.Println(v_7)
	sparse := hxrt.NewArray()
	sparseAlias := sparse
	hx_array_target_29 := sparse
	hx_array_index_30 := 2
	hx_array_target_29.Set(hx_array_index_30, 14)
	sparseAlias.Push(15)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sparse="), show(sparse)))
	hxrt.Println(v_8)
	identity := hxrt.NewArray(16)
	identityAlias := identity
	identityCopy := identity.Copy()
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("identity="), hxrt.StdString((identity == identityAlias))), hxrt.StringFromLiteral(":")), hxrt.StdString((identity == identityCopy))))
	hxrt.Println(v_9)
	holderEvaluations = 0
	events = hxrt.NewArray()
	hx_array_target_34 := makeHolder().values
	hx_arr_32 := events
	hx_arr_32.Push(hxrt.StringFromLiteral("index"))
	hx_array_index_35 := 2
	hx_arr_33 := events
	hx_arr_33.Push(hxrt.StringFromLiteral("value"))
	hx_array_target_34.Set(hx_array_index_35, 23)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("order="), holderEvaluations), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_10)
	operationValues = hxrt.NewArray(1, 2, 3)
	events = hxrt.NewArray()
	base := markedOperationValues()
	index := markedIndex()
	compoundResult := func() int {
		hx_array_target_36 := base
		hx_array_index_37 := index
		var hx_array_current_38 int = hxrt.IntFromNullableAny(hx_array_target_36.Get(hx_array_index_37))
		var hx_array_assigned_39 int = int((hxrt.Int32Wrap(hx_array_current_38) + hxrt.Int32Wrap(markedValue())))
		hx_array_target_36.Set(hx_array_index_37, hx_array_assigned_39)
		return hx_array_assigned_39
	}()
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("compound="), compoundResult), hxrt.StringFromLiteral(":")), show(operationValues)), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_11)
	operationValues = hxrt.NewArray(4, 5, 6)
	events = hxrt.NewArray()
	postResult := func() int {
		hx_array_target_40 := markedOperationValues()
		hx_array_index_41 := markedIndex()
		var hx_array_current_42 int = hxrt.IntFromNullableAny(hx_array_target_40.Get(hx_array_index_41))
		var hx_array_next_43 int = int(int32((hx_array_current_42 + 1)))
		hx_array_target_40.Set(hx_array_index_41, hx_array_next_43)
		return hx_array_current_42
	}()
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("post="), postResult), hxrt.StringFromLiteral(":")), show(operationValues)), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_12)
}

func makeHolder() *SnapshotArrayIdentityHolder {
	holderEvaluations = int(int32((holderEvaluations + 1)))
	hx_arr_44 := events
	hx_arr_44.Push(hxrt.StringFromLiteral("target"))
	return New_SnapshotArrayIdentityHolder(hxrt.NewArray(20))
}

func markedIndex() int {
	hx_arr_45 := events
	hx_arr_45.Push(hxrt.StringFromLiteral("index"))
	return 2
}

func markedOperationValues() *hxrt.Array {
	hx_arr_46 := events
	hx_arr_46.Push(hxrt.StringFromLiteral("target"))
	return operationValues
}

func markedValue() int {
	hx_arr_47 := events
	hx_arr_47.Push(hxrt.StringFromLiteral("value"))
	return 23
}

var operationValues *hxrt.Array = hxrt.NewArray()

func pushGeneric(values *hxrt.Array, value any) {
	values.Push(value)
}

func pushParameter(values *hxrt.Array, value int) {
	values.Push(value)
}

func retainErased(values *hxrt.Array) {
	retainedGeneric = values
}

var retainedGeneric *hxrt.Array

func returnAlias(values *hxrt.Array) *hxrt.Array {
	return values
}

func show(values *hxrt.Array) *string {
	return hxrt.StringJoinAny(func() *hxrt.Array {
		_g := hxrt.NewArray()
		_g1 := 0
		for _g1 < values.Len() {
			var value any = values.Get(_g1)
			_g1 = int(int32((_g1 + 1)))
			_g.Push(hxrt.StdString(value))
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(","))
}

type I_SnapshotArrayIdentityBox interface {
}

type SnapshotArrayIdentityBox struct {
	__hx_this I_SnapshotArrayIdentityBox
	id        int
}

func (self *SnapshotArrayIdentityBox) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotArrayIdentityBox(id int) *SnapshotArrayIdentityBox {
	return hx_init_SnapshotArrayIdentityBox(hx_alloc_SnapshotArrayIdentityBox(nil), id)
}

func hx_alloc_SnapshotArrayIdentityBox(owner any) *SnapshotArrayIdentityBox {
	self := &SnapshotArrayIdentityBox{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SnapshotArrayIdentityBox)
	}
	return self
}

func hx_init_SnapshotArrayIdentityBox(self *SnapshotArrayIdentityBox, id int) *SnapshotArrayIdentityBox {
	self.id = id
	return self
}

type I_SnapshotArrayIdentityHolder interface {
}

type SnapshotArrayIdentityHolder struct {
	__hx_this I_SnapshotArrayIdentityHolder
	values    *hxrt.Array
}

func (self *SnapshotArrayIdentityHolder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotArrayIdentityHolder(values *hxrt.Array) *SnapshotArrayIdentityHolder {
	return hx_init_SnapshotArrayIdentityHolder(hx_alloc_SnapshotArrayIdentityHolder(nil), values)
}

func hx_alloc_SnapshotArrayIdentityHolder(owner any) *SnapshotArrayIdentityHolder {
	self := &SnapshotArrayIdentityHolder{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SnapshotArrayIdentityHolder)
	}
	return self
}

func hx_init_SnapshotArrayIdentityHolder(self *SnapshotArrayIdentityHolder, values *hxrt.Array) *SnapshotArrayIdentityHolder {
	self.values = values
	return self
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *SnapshotArrayIdentityBox:
		return func(hx_structural_source_7 *SnapshotArrayIdentityBox) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_7, func(name string) hxrt.AnonymousField {
				switch name {
				case "id":
					return hxrt.AnonymousField{Value: hx_structural_source_7.id, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *SnapshotArrayIdentityHolder:
		return func(hx_structural_source_8 *SnapshotArrayIdentityHolder) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_8, func(name string) hxrt.AnonymousField {
				switch name {
				case "values":
					return hxrt.AnonymousField{Value: hx_structural_source_8.values, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "values":
					hx_structural_source_8.values = func(hx_value_9 any) *hxrt.Array {
						if hx_value_9 == nil {
							var hx_zero_10 *hxrt.Array
							return hx_zero_10
						}
						return hx_value_9.(*hxrt.Array)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_11 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_11, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_11.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_11.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_11.high = func(hx_value_12 any) int {
						if hx_value_12 == nil {
							var hx_zero_13 int
							return hx_zero_13
						}
						return hx_value_12.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_11.low = func(hx_value_14 any) int {
						if hx_value_14 == nil {
							var hx_zero_15 int
							return hx_zero_15
						}
						return hx_value_14.(int)
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

func hx_recover_callable_0(value any) *hxrt.Function[func() hxrt.Anonymous] {
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
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() bool] {
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
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_16 any) bool {
				if hx_value_16 == nil {
					var hx_zero_17 bool
					return hx_zero_17
				}
				return hx_value_16.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_18 any) bool {
				if hx_value_18 == nil {
					var hx_zero_19 bool
					return hx_zero_19
				}
				return hx_value_18.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() any] {
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
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_3(value any) *hxrt.Function[func(*hxrt.Array, any)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*hxrt.Array, any)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*hxrt.Array, int)]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Array, arg_1 any) {
			hxrt.FunctionCall(source)(arg_0, func(hx_value_20 any) int {
				if hx_value_20 == nil {
					var hx_zero_21 int
					return hx_zero_21
				}
				return hx_value_20.(int)
			}(any(arg_1)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
