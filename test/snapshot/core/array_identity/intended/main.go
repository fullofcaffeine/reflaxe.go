package main

import "snapshot/hxrt"

func applyGenericCallback(values *hxrt.Array, value any, callback func(*hxrt.Array, any)) {
	callback(values, value)
}

var events *hxrt.Array = hxrt.NewArray()

var holderEvaluations int = 0

func main() {
	local := hxrt.NewArray(1, 2)
	localAlias := local
	localAlias.Push(3)
	func() {
		hx_insert_position_3 := 1
		var hx_insert_value_4 any = 9
		localAlias.Insert(hx_insert_position_3, hx_insert_value_4)
	}()
	func() bool {
		var hx_remove_value_6 any = 2
		hx_remove_index_7 := 0
		for hx_remove_index_7 < localAlias.Len() {
			hx_remove_element_8 := localAlias.Get(hx_remove_index_7)
			if hxrt.HaxeEqual(hx_remove_element_8, hx_remove_value_6) {
				localAlias.RemoveAt(hx_remove_index_7)
				return true
			}
			hx_remove_index_7 = (hx_remove_index_7 + 1)
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
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic.has="), hxrt.StdString(Lambda_has(func() map[string]any {
		hx_lambda_source_11 := hxrt.NewArray(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"))
		hx_lambda_wrapped_12 := map[string]any{}
		hx_lambda_wrapped_12["iterator"] = func() map[string]any {
			hx_lambda_index_15 := 0
			hx_lambda_iterator_map_13 := map[string]any{}
			hx_lambda_iterator_map_13["hasNext"] = func() bool {
				return (hx_lambda_index_15 < hx_lambda_source_11.Len())
			}
			hx_lambda_iterator_map_13["next"] = func() any {
				hx_lambda_value_14 := hx_lambda_source_11.Get(hx_lambda_index_15)
				hx_lambda_index_15 = (hx_lambda_index_15 + 1)
				return hx_lambda_value_14
			}
			return hx_lambda_iterator_map_13
		}
		return hx_lambda_wrapped_12
	}(), hxrt.StringFromLiteral("a")))))
	hxrt.Println(v_4)
	stringElements := hxrt.ArrayFromValues(func(hx_sort_src_16 []*string) []any {
		hx_sort_out_18 := make([]any, 0, len(hx_sort_src_16))
		for _, hx_sort_item_17 := range hx_sort_src_16 {
			hx_sort_out_18 = append(hx_sort_out_18, hx_sort_item_17)
		}
		return hx_sort_out_18
	}(hxrt.StringSplitStringPtr(hxrt.StringFromLiteral("left,right"), hxrt.StringFromLiteral(";"))))
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("string.split="), hxrt.ArrayFromValues(func(hx_sort_src_21 []*string) []any {
		hx_sort_out_23 := make([]any, 0, len(hx_sort_src_21))
		for _, hx_sort_item_22 := range hx_sort_src_21 {
			hx_sort_out_23 = append(hx_sort_out_23, hx_sort_item_22)
		}
		return hx_sort_out_23
	}(hxrt.StringSplitStringPtr(func(hx_value_19 any) *string {
		if hx_value_19 == nil {
			var hx_zero_20 *string
			return hx_zero_20
		}
		return hx_value_19.(*string)
	}(stringElements.Get(0)), hxrt.StringFromLiteral(",")))).Len()))
	hxrt.Println(v_5)
	retainedSource := hxrt.NewArray(10)
	retainErased(retainedSource)
	hx_arr_24 := retainedGeneric
	hx_arr_24.Push(11)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("retained="), show(retainedSource)))
	hxrt.Println(v_6)
	callbackSource := hxrt.NewArray(12)
	applyGenericCallback(callbackSource, 13, func(hx_erased_callback_arg_26 *hxrt.Array, hx_erased_callback_arg_27 any) {
		func(values *hxrt.Array, value int) {
			values.Push(value)
		}(hx_erased_callback_arg_26, func(hx_value_28 any) int {
			if hx_value_28 == nil {
				var hx_zero_29 int
				return hx_zero_29
			}
			return hx_value_28.(int)
		}(hx_erased_callback_arg_27))
	})
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("callback="), show(callbackSource)))
	hxrt.Println(v_7)
	sparse := hxrt.NewArray()
	sparseAlias := sparse
	hx_array_target_30 := sparse
	hx_array_index_31 := 2
	hx_array_target_30.Set(hx_array_index_31, 14)
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
	hx_array_target_35 := makeHolder().values
	hx_arr_33 := events
	hx_arr_33.Push(hxrt.StringFromLiteral("index"))
	hx_array_index_36 := 2
	hx_arr_34 := events
	hx_arr_34.Push(hxrt.StringFromLiteral("value"))
	hx_array_target_35.Set(hx_array_index_36, 23)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("order="), holderEvaluations), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_10)
	operationValues = hxrt.NewArray(1, 2, 3)
	events = hxrt.NewArray()
	base := markedOperationValues()
	index := markedIndex()
	compoundResult := func() int {
		hx_array_target_37 := base
		hx_array_index_38 := index
		var hx_array_current_39 int = hxrt.IntFromNullableAny(hx_array_target_37.Get(hx_array_index_38))
		var hx_array_assigned_40 int = int(int32((hxrt.Int32Wrap(hx_array_current_39) + hxrt.Int32Wrap(markedValue()))))
		hx_array_target_37.Set(hx_array_index_38, hx_array_assigned_40)
		return hx_array_assigned_40
	}()
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("compound="), compoundResult), hxrt.StringFromLiteral(":")), show(operationValues)), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_11)
	operationValues = hxrt.NewArray(4, 5, 6)
	events = hxrt.NewArray()
	postResult := func() int {
		hx_array_target_41 := markedOperationValues()
		hx_array_index_42 := markedIndex()
		var hx_array_current_43 int = hxrt.IntFromNullableAny(hx_array_target_41.Get(hx_array_index_42))
		var hx_array_next_44 int = int(int32((hx_array_current_43 + 1)))
		hx_array_target_41.Set(hx_array_index_42, hx_array_next_44)
		return hx_array_current_43
	}()
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("post="), postResult), hxrt.StringFromLiteral(":")), show(operationValues)), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_12)
}

func makeHolder() *SnapshotArrayIdentityHolder {
	holderEvaluations = int(int32((holderEvaluations + 1)))
	hx_arr_45 := events
	hx_arr_45.Push(hxrt.StringFromLiteral("target"))
	return New_SnapshotArrayIdentityHolder(hxrt.NewArray(20))
}

func markedIndex() int {
	hx_arr_46 := events
	hx_arr_46.Push(hxrt.StringFromLiteral("index"))
	return 2
}

func markedOperationValues() *hxrt.Array {
	hx_arr_47 := events
	hx_arr_47.Push(hxrt.StringFromLiteral("target"))
	return operationValues
}

func markedValue() int {
	hx_arr_48 := events
	hx_arr_48.Push(hxrt.StringFromLiteral("value"))
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

func New_SnapshotArrayIdentityBox(id int) *SnapshotArrayIdentityBox {
	self := &SnapshotArrayIdentityBox{}
	self.__hx_this = self
	self.id = id
	return self
}

type I_SnapshotArrayIdentityHolder interface {
}

type SnapshotArrayIdentityHolder struct {
	__hx_this I_SnapshotArrayIdentityHolder
	values    *hxrt.Array
}

func New_SnapshotArrayIdentityHolder(values *hxrt.Array) *SnapshotArrayIdentityHolder {
	self := &SnapshotArrayIdentityHolder{}
	self.__hx_this = self
	self.values = values
	return self
}
