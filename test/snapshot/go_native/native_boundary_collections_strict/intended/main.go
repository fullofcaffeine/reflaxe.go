package main

import (
	"errors"
	"reflect"
	"snapshot/hxrt"
)

func main() {
	var v any = any(NativeCollectionOps_eval())
	hxrt.Println(v)
	var v_1 any = any(NativeResultOps_eval())
	hxrt.Println(v_1)
}

func NativeCollectionOps_eval() *string {
	channel := go__concurrency_newChan__int_95e97e5e(1)
	go__concurrency_send__int_95e97e5e(channel.__hx_native, 7)
	received := go__concurrency_recvOr__int_95e97e5e(channel.__hx_native, -1)
	slice := go___Go_newSlice()
	go__slice_push__int_95e97e5e(slice, received)
	go__slice_set__int_95e97e5e(slice, 0, int(int32((hxrt.Int32Wrap(go__slice_get__int_95e97e5e(slice, 0)) + hxrt.Int32Wrap(1)))))
	map_ := go___Go_newMap()
	go__map_set__int__int_b895c9bd(map_, 1, go__slice_get__int_95e97e5e(slice, 0))
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(func(hx_value_1 any) any {
		if hx_value_1 == nil {
			return nil
		}
		return hx_value_1.(int)
	}(go__map_get__int__int_b895c9bd(map_, 1))), hxrt.StringFromLiteral("|")), func() *string {
		var hx_if_2 *string
		if go__map_exists__int__int_b895c9bd(map_, 1) {
			hx_if_2 = hxrt.StringFromLiteral("1")
		} else {
			hx_if_2 = hxrt.StringFromLiteral("0")
		}
		return hx_if_2
	}())
}

func NativeResultOps_eval() *string {
	ok := go__result_ok___string_f613ccd0(hxrt.StringFromLiteral("done"))
	err := go__result_failure___string_f613ccd0(hxrt.StringFromLiteral("broken"))
	errValue := go__result_error___string_f613ccd0(err)
	var hx_if_3 *string
	if hxrt.StringEqualStringPtr(errValue, nil) {
		hx_if_3 = hxrt.StringFromLiteral("none")
	} else {
		hx_if_3 = errValue
	}
	errLabel := hx_if_3
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(go__result_unwrap___string_f613ccd0(ok), hxrt.StringFromLiteral("|")), func() *string {
		var hx_if_4 *string
		if go__result_isOk___string_f613ccd0(ok) {
			hx_if_4 = hxrt.StringFromLiteral("1")
		} else {
			hx_if_4 = hxrt.StringFromLiteral("0")
		}
		return hx_if_4
	}()), hxrt.StringFromLiteral("|")), func() *string {
		var hx_if_5 *string
		if go__result_isErr___string_f613ccd0(err) {
			hx_if_5 = hxrt.StringFromLiteral("1")
		} else {
			hx_if_5 = hxrt.StringFromLiteral("0")
		}
		return hx_if_5
	}()), hxrt.StringFromLiteral("|")), errLabel)
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

func go__concurrency_makeChan__int_95e97e5e(buffer int) any {
	if buffer > 0 {
		return make(chan int, buffer)
	}
	return make(chan int)
}

func go__concurrency_setBuffer__int_95e97e5e(channel *go___Chan, buffer int) {
	if channel == nil {
		return
	}
	channel.__hx_native = go__concurrency_makeChan__int_95e97e5e(buffer)
}

func go__concurrency_newChan__int_95e97e5e(buffer int) *go___Chan {
	channel := New_go___Chan()
	go__concurrency_setBuffer__int_95e97e5e(channel, buffer)
	return channel
}

func go__concurrency_send__int_95e97e5e(channel any, value int) {
	channel.(chan int) <- value
}

func go__concurrency_trySend__int_95e97e5e(channel any, value int) bool {
	select {
	case channel.(chan int) <- value:
		return true
	default:
		return false
	}
}

func go__concurrency_recv__int_95e97e5e(channel any) int {
	return <-channel.(chan int)
}

func go__concurrency_recvOr__int_95e97e5e(channel any, defaultValue int) int {
	select {
	case value, received := <-channel.(chan int):
		if !received {
			return defaultValue
		}
		return value
	default:
		return defaultValue
	}
}

func go__concurrency_tryRecv__int_95e97e5e(channel any) *go___Result {
	select {
	case value, received := <-channel.(chan int):
		if !received {
			return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("closed")))
		}
		return New_go___Result(value, nil)
	default:
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("empty")))
	}
}

func go__concurrency_close__int_95e97e5e(channel any) {
	close(channel.(chan int))
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

func go__map_set__int__int_b895c9bd(mapValue *go___Map, key int, value int) {
	mapValue.inner.set(hxrt.StdString(any(key)), value)
}

func go__map_get__int__int_b895c9bd(mapValue *go___Map, key int) any {
	return mapValue.inner.get(hxrt.StdString(any(key)))
}

func go__map_exists__int__int_b895c9bd(mapValue *go___Map, key int) bool {
	return mapValue.inner.exists(hxrt.StdString(any(key)))
}

func go__result_fromValueError(value any, err error) *go___Result {
	if err != nil {
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral(err.Error())))
	}
	return New_go___Result(value, nil)
}

func go__result_ok___string_f613ccd0(value *string) *go___Result {
	return New_go___Result(value, nil)
}

func go__result_failure___string_f613ccd0(message *string) *go___Result {
	return New_go___Result(nil, New_go___Error(message))
}

func go__result_valueError___string_f613ccd0(result *go___Result) (*string, error) {
	var zero *string
	if result == nil {
		return zero, errors.New("nil go.Result")
	}
	if result.errorValue != nil {
		return zero, errors.New(*hxrt.StdString(result.errorValue.message))
	}
	if result.value == nil {
		return zero, nil
	}
	return result.value.(*string), nil
}

func go__result_isOk___string_f613ccd0(result *go___Result) bool {
	_, err := go__result_valueError___string_f613ccd0(result)
	return (err == nil)
}

func go__result_isErr___string_f613ccd0(result *go___Result) bool {
	_, err := go__result_valueError___string_f613ccd0(result)
	return (err != nil)
}

func go__result_unwrap___string_f613ccd0(result *go___Result) *string {
	value, err := go__result_valueError___string_f613ccd0(result)
	if err != nil {
		hxrt.Throw(hxrt.StringFromLiteral(err.Error()))
		var zero *string
		return zero
	}
	return value
}

func go__result_error___string_f613ccd0(result *go___Result) *string {
	_, err := go__result_valueError___string_f613ccd0(result)
	if err == nil {
		return nil
	}
	return hxrt.StringFromLiteral(err.Error())
}
