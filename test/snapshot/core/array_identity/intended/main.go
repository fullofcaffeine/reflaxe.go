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
		hx_insert_position_12 := 1
		var hx_insert_value_13 any = 9
		localAlias.Insert(hx_insert_position_12, hx_insert_value_13)
	}()
	func() bool {
		var hx_remove_value_15 any = 2
		hx_remove_index_16 := 0
		for hx_remove_index_16 < localAlias.Len() {
			hx_remove_element_17 := localAlias.Get(hx_remove_index_16)
			if hxrt.HaxeEqual(hx_remove_element_17, hx_remove_value_15) {
				localAlias.RemoveAt(hx_remove_index_16)
				return true
			}
			hx_remove_index_16 = (hx_remove_index_16 + 1)
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
		hx_lambda_source_20 := hxrt.NewArray(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"))
		hx_lambda_wrapped_21 := map[string]any{}
		hx_lambda_wrapped_21["iterator"] = func() map[string]any {
			hx_lambda_index_24 := 0
			hx_lambda_iterator_map_22 := map[string]any{}
			hx_lambda_iterator_map_22["hasNext"] = func() bool {
				return (hx_lambda_index_24 < hx_lambda_source_20.Len())
			}
			hx_lambda_iterator_map_22["next"] = func() any {
				hx_lambda_value_23 := hx_lambda_source_20.Get(hx_lambda_index_24)
				hx_lambda_index_24 = (hx_lambda_index_24 + 1)
				return hx_lambda_value_23
			}
			return hx_lambda_iterator_map_22
		}
		return hx_lambda_wrapped_21
	}(), hxrt.StringFromLiteral("a")))))
	hxrt.Println(v_4)
	stringElements := hxrt.ArrayFromValues(func(hx_sort_src_25 []*string) []any {
		hx_sort_out_27 := make([]any, 0, len(hx_sort_src_25))
		for _, hx_sort_item_26 := range hx_sort_src_25 {
			hx_sort_out_27 = append(hx_sort_out_27, hx_sort_item_26)
		}
		return hx_sort_out_27
	}(hxrt.StringSplitStringPtr(hxrt.StringFromLiteral("left,right"), hxrt.StringFromLiteral(";"))))
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("string.split="), hxrt.ArrayFromValues(func(hx_sort_src_30 []*string) []any {
		hx_sort_out_32 := make([]any, 0, len(hx_sort_src_30))
		for _, hx_sort_item_31 := range hx_sort_src_30 {
			hx_sort_out_32 = append(hx_sort_out_32, hx_sort_item_31)
		}
		return hx_sort_out_32
	}(hxrt.StringSplitStringPtr(func(hx_value_28 any) *string {
		if hx_value_28 == nil {
			var hx_zero_29 *string
			return hx_zero_29
		}
		return hx_value_28.(*string)
	}(stringElements.Get(0)), hxrt.StringFromLiteral(",")))).Len()))
	hxrt.Println(v_5)
	retainedSource := hxrt.NewArray(10)
	retainErased(retainedSource)
	hx_arr_33 := retainedGeneric
	hx_arr_33.Push(11)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("retained="), show(retainedSource)))
	hxrt.Println(v_6)
	callbackSource := hxrt.NewArray(12)
	applyGenericCallback(callbackSource, 13, func(hx_erased_callback_arg_35 *hxrt.Array, hx_erased_callback_arg_36 any) {
		func(values *hxrt.Array, value int) {
			values.Push(value)
		}(hx_erased_callback_arg_35, func(hx_value_37 any) int {
			if hx_value_37 == nil {
				var hx_zero_38 int
				return hx_zero_38
			}
			return hx_value_37.(int)
		}(hx_erased_callback_arg_36))
	})
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("callback="), show(callbackSource)))
	hxrt.Println(v_7)
	sparse := hxrt.NewArray()
	sparseAlias := sparse
	hx_array_target_39 := sparse
	hx_array_index_40 := 2
	hx_array_target_39.Set(hx_array_index_40, 14)
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
	hx_array_target_44 := makeHolder().values
	hx_arr_42 := events
	hx_arr_42.Push(hxrt.StringFromLiteral("index"))
	hx_array_index_45 := 2
	hx_arr_43 := events
	hx_arr_43.Push(hxrt.StringFromLiteral("value"))
	hx_array_target_44.Set(hx_array_index_45, 23)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("order="), holderEvaluations), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_10)
	operationValues = hxrt.NewArray(1, 2, 3)
	events = hxrt.NewArray()
	base := markedOperationValues()
	index := markedIndex()
	compoundResult := func() int {
		hx_array_target_46 := base
		hx_array_index_47 := index
		var hx_array_current_48 int = hxrt.IntFromNullableAny(hx_array_target_46.Get(hx_array_index_47))
		var hx_array_assigned_49 int = int(int32((hxrt.Int32Wrap(hx_array_current_48) + hxrt.Int32Wrap(markedValue()))))
		hx_array_target_46.Set(hx_array_index_47, hx_array_assigned_49)
		return hx_array_assigned_49
	}()
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("compound="), compoundResult), hxrt.StringFromLiteral(":")), show(operationValues)), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_11)
	operationValues = hxrt.NewArray(4, 5, 6)
	events = hxrt.NewArray()
	postResult := func() int {
		hx_array_target_50 := markedOperationValues()
		hx_array_index_51 := markedIndex()
		var hx_array_current_52 int = hxrt.IntFromNullableAny(hx_array_target_50.Get(hx_array_index_51))
		var hx_array_next_53 int = int(int32((hx_array_current_52 + 1)))
		hx_array_target_50.Set(hx_array_index_51, hx_array_next_53)
		return hx_array_current_52
	}()
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("post="), postResult), hxrt.StringFromLiteral(":")), show(operationValues)), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(events.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_12)
}

func makeHolder() *SnapshotArrayIdentityHolder {
	holderEvaluations = int(int32((holderEvaluations + 1)))
	hx_arr_54 := events
	hx_arr_54.Push(hxrt.StringFromLiteral("target"))
	return New_SnapshotArrayIdentityHolder(hxrt.NewArray(20))
}

func markedIndex() int {
	hx_arr_55 := events
	hx_arr_55.Push(hxrt.StringFromLiteral("index"))
	return 2
}

func markedOperationValues() *hxrt.Array {
	hx_arr_56 := events
	hx_arr_56.Push(hxrt.StringFromLiteral("target"))
	return operationValues
}

func markedValue() int {
	hx_arr_57 := events
	hx_arr_57.Push(hxrt.StringFromLiteral("value"))
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
