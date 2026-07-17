package main

import (
	"errors"
	"reflect"
	"snapshot/hxrt"
)

func fallbackPath() *string {
	slice := go___Go_newSlice()
	slice.push(func(hx_value_1 any) any {
		if hx_value_1 == nil {
			return nil
		}
		return hx_value_1.(int)
	}(nil))
	slice.push(5)
	map_ := go___Go_newMap()
	key := hxrt.NewArray(1, 2)
	map_.set(key, 7)
	okResult := go___Go_ok(nil)
	errResult := go___Go_fail(hxrt.StringFromLiteral("lane"))
	errText := func(hx_value_2 any) *string {
		if hx_value_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_value_2.(*string)
	}(errResult.error())
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(func() int {
		var hx_if_5 int
		if func(hx_value_4 any) any {
			if hx_value_4 == nil {
				return nil
			}
			return hx_value_4.(int)
		}(slice.get(0)) == nil {
			hx_if_5 = 1
		} else {
			hx_if_5 = 0
		}
		return hx_if_5
	}(), hxrt.StringFromLiteral(",")), func(hx_value_6 any) any {
		if hx_value_6 == nil {
			return nil
		}
		return hx_value_6.(int)
	}(slice.get(1))), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_9 int
		if func(hx_value_7 any) bool {
			if hx_value_7 == nil {
				var hx_zero_8 bool
				return hx_zero_8
			}
			return hx_value_7.(bool)
		}(map_.exists(key)) {
			hx_if_9 = 1
		} else {
			hx_if_9 = 0
		}
		return hx_if_9
	}()), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_12 int
		if func(hx_value_10 any) bool {
			if hx_value_10 == nil {
				var hx_zero_11 bool
				return hx_zero_11
			}
			return hx_value_10.(bool)
		}(map_.exists(hxrt.NewArray(1, 2))) {
			hx_if_12 = 1
		} else {
			hx_if_12 = 0
		}
		return hx_if_12
	}()), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_14 int
		if func(hx_value_13 any) any {
			if hx_value_13 == nil {
				return nil
			}
			return hx_value_13.(int)
		}(map_.get(hxrt.NewArray(3, 4))) == nil {
			hx_if_14 = 1
		} else {
			hx_if_14 = 0
		}
		return hx_if_14
	}()), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_17 int
		if func(hx_value_15 any) bool {
			if hx_value_15 == nil {
				var hx_zero_16 bool
				return hx_zero_16
			}
			return hx_value_15.(bool)
		}(okResult.isOk()) {
			hx_if_17 = 1
		} else {
			hx_if_17 = 0
		}
		return hx_if_17
	}()), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_19 int
		if func(hx_value_18 any) any {
			if hx_value_18 == nil {
				return nil
			}
			return hx_value_18.(int)
		}(okResult.unwrap()) == nil {
			hx_if_19 = 1
		} else {
			hx_if_19 = 0
		}
		return hx_if_19
	}()), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_22 int
		if func(hx_value_20 any) bool {
			if hx_value_20 == nil {
				var hx_zero_21 bool
				return hx_zero_21
			}
			return hx_value_20.(bool)
		}(errResult.isErr()) {
			hx_if_22 = 1
		} else {
			hx_if_22 = 0
		}
		return hx_if_22
	}()), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_23 int
		if !hxrt.StringEqualStringPtr(errText, nil) && hxrt.StringEqualStringPtr(errText, hxrt.StringFromLiteral("lane")) {
			hx_if_23 = 1
		} else {
			hx_if_23 = 0
		}
		return hx_if_23
	}())
}

func main() {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typed="), typedPath()))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fallback="), fallbackPath()))
	hxrt.Println(v_1)
}

func typedPath() *string {
	slice := go___Go_newSlice()
	go__slice_push__int_95e97e5e(slice, 3)
	go__slice_push__int_95e97e5e(slice, 4)
	go__slice_set__int_95e97e5e(slice, 1, 9)
	map_ := go___Go_newMap()
	go__map_set___string__int_e8ed7ec7(map_, hxrt.StringFromLiteral("alpha"), 7)
	go__map_set___string__int_e8ed7ec7(map_, hxrt.StringFromLiteral("beta"), 5)
	ok := go__result_ok__int_95e97e5e(42)
	err := go__result_failure__int_95e97e5e(hxrt.StringFromLiteral("typed"))
	errText := go__result_error__int_95e97e5e(err)
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(go__slice_length__int_95e97e5e(slice), hxrt.StringFromLiteral(",")), go__slice_get__int_95e97e5e(slice, 1)), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_24 int
		if go__map_exists___string__int_e8ed7ec7(map_, hxrt.StringFromLiteral("alpha")) {
			hx_if_24 = 1
		} else {
			hx_if_24 = 0
		}
		return hx_if_24
	}()), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_25 int
		if go__map_exists___string__int_e8ed7ec7(map_, hxrt.StringFromLiteral("gamma")) {
			hx_if_25 = 1
		} else {
			hx_if_25 = 0
		}
		return hx_if_25
	}()), hxrt.StringFromLiteral(",")), func(hx_value_26 any) any {
		if hx_value_26 == nil {
			return nil
		}
		return hx_value_26.(int)
	}(go__map_get___string__int_e8ed7ec7(map_, hxrt.StringFromLiteral("beta")))), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_28 int
		if func(hx_value_27 any) any {
			if hx_value_27 == nil {
				return nil
			}
			return hx_value_27.(int)
		}(go__map_get___string__int_e8ed7ec7(map_, hxrt.StringFromLiteral("gamma"))) == nil {
			hx_if_28 = 1
		} else {
			hx_if_28 = 0
		}
		return hx_if_28
	}()), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_29 int
		if go__result_isOk__int_95e97e5e(ok) {
			hx_if_29 = 1
		} else {
			hx_if_29 = 0
		}
		return hx_if_29
	}()), hxrt.StringFromLiteral(",")), go__result_unwrap__int_95e97e5e(ok)), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_30 int
		if go__result_isErr__int_95e97e5e(err) {
			hx_if_30 = 1
		} else {
			hx_if_30 = 0
		}
		return hx_if_30
	}()), hxrt.StringFromLiteral(",")), func() int {
		var hx_if_31 int
		if !hxrt.StringEqualStringPtr(errText, nil) && hxrt.StringEqualStringPtr(errText, hxrt.StringFromLiteral("typed")) {
			hx_if_31 = 1
		} else {
			hx_if_31 = 0
		}
		return hx_if_31
	}())
}

func go__concurrency_makeChan(buffer int) any {
	if buffer > 0 {
		return make(chan any, buffer)
	}
	return make(chan any)
}

func go__concurrency_send(channel any, value any) {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return
	}
	sendValue := reflect.ValueOf(value)
	if !sendValue.IsValid() {
		sendValue = reflect.Zero(chanValue.Type().Elem())
	} else if !sendValue.Type().AssignableTo(chanValue.Type().Elem()) {
		if sendValue.Type().ConvertibleTo(chanValue.Type().Elem()) {
			sendValue = sendValue.Convert(chanValue.Type().Elem())
		} else {
			return
		}
	}
	chanValue.Send(sendValue)
}

func go__concurrency_trySend(channel any, value any) bool {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return false
	}
	sendValue := reflect.ValueOf(value)
	if !sendValue.IsValid() {
		sendValue = reflect.Zero(chanValue.Type().Elem())
	} else if !sendValue.Type().AssignableTo(chanValue.Type().Elem()) {
		if sendValue.Type().ConvertibleTo(chanValue.Type().Elem()) {
			sendValue = sendValue.Convert(chanValue.Type().Elem())
		} else {
			return false
		}
	}
	cases := []reflect.SelectCase{
		{Dir: reflect.SelectSend, Chan: chanValue, Send: sendValue},
		{Dir: reflect.SelectDefault},
	}
	chosen, _, _ := reflect.Select(cases)
	return chosen == 0
}

func go__concurrency_recv(channel any) any {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return nil
	}
	recvValue, _ := chanValue.Recv()
	if !recvValue.IsValid() {
		return nil
	}
	return recvValue.Interface()
}

func go__concurrency_recvOr(channel any, defaultValue any) any {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return defaultValue
	}
	cases := []reflect.SelectCase{
		{Dir: reflect.SelectRecv, Chan: chanValue},
		{Dir: reflect.SelectDefault},
	}
	chosen, recvValue, received := reflect.Select(cases)
	if chosen == 0 {
		if !received {
			return defaultValue
		}
		return recvValue.Interface()
	}
	return defaultValue
}

func go__concurrency_tryRecv(channel any) *go___Result {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("empty")))
	}
	cases := []reflect.SelectCase{
		{Dir: reflect.SelectRecv, Chan: chanValue},
		{Dir: reflect.SelectDefault},
	}
	chosen, recvValue, received := reflect.Select(cases)
	if chosen == 0 {
		if !received {
			return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("closed")))
		}
		return New_go___Result(recvValue.Interface(), nil)
	}
	return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("empty")))
}

func go__concurrency_close(channel any) {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return
	}
	chanValue.Close()
}

func go__concurrency_spawn(fn func()) {
	go fn()
}

func go__slice_push__int_95e97e5e(slice *go___Slice, value int) {
	slice.data = append(slice.data, value)
}

func go__slice_set__int_95e97e5e(slice *go___Slice, index int, value int) {
	slice.data[index] = value
}

func go__slice_get__int_95e97e5e(slice *go___Slice, index int) int {
	raw := slice.data[index]
	if raw == nil {
		var zero int
		return zero
	}
	return raw.(int)
}

func go__slice_length__int_95e97e5e(slice *go___Slice) int {
	return len(slice.data)
}

func go__slice_toArray__int_95e97e5e(slice *go___Slice) []int {
	raw := slice.data
	out := make([]int, len(raw))
	for idx, value := range raw {
		if value == nil {
			continue
		}
		out[idx] = value.(int)
	}
	return out
}

func go__map_set___string__int_e8ed7ec7(mapValue *go___Map, key *string, value int) {
	mapValue.inner.set(hxrt.StdString(any(key)), value)
}

func go__map_get___string__int_e8ed7ec7(mapValue *go___Map, key *string) any {
	return mapValue.inner.get(hxrt.StdString(any(key)))
}

func go__map_exists___string__int_e8ed7ec7(mapValue *go___Map, key *string) bool {
	return mapValue.inner.exists(hxrt.StdString(any(key)))
}

func go__result_fromValueError(value any, err error) *go___Result {
	if err != nil {
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral(err.Error())))
	}
	return New_go___Result(value, nil)
}

func go__result_ok__int_95e97e5e(value int) *go___Result {
	return New_go___Result(value, nil)
}

func go__result_failure__int_95e97e5e(message *string) *go___Result {
	return New_go___Result(nil, New_go___Error(message))
}

func go__result_valueError__int_95e97e5e(result *go___Result) (int, error) {
	var zero int
	if result == nil {
		return zero, errors.New("nil go.Result")
	}
	if result.errorValue != nil {
		return zero, errors.New(*hxrt.StdString(result.errorValue.message))
	}
	if result.value == nil {
		return zero, nil
	}
	return result.value.(int), nil
}

func go__result_isOk__int_95e97e5e(result *go___Result) bool {
	_, err := go__result_valueError__int_95e97e5e(result)
	return (err == nil)
}

func go__result_isErr__int_95e97e5e(result *go___Result) bool {
	_, err := go__result_valueError__int_95e97e5e(result)
	return (err != nil)
}

func go__result_unwrap__int_95e97e5e(result *go___Result) int {
	value, err := go__result_valueError__int_95e97e5e(result)
	if err != nil {
		hxrt.Throw(hxrt.StringFromLiteral(err.Error()))
		var zero int
		return zero
	}
	return value
}

func go__result_error__int_95e97e5e(result *go___Result) *string {
	_, err := go__result_valueError__int_95e97e5e(result)
	if err == nil {
		return nil
	}
	return hxrt.StringFromLiteral(err.Error())
}
