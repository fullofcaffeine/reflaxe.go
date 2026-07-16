package main

import (
	"reflect"
	"snapshot/hxrt"
)

func main() {
	ints := New_go___Slice()
	go__slice_push__int_95e97e5e(ints, 3)
	go__slice_push__int_95e97e5e(ints, 5)
	go__slice_set__int_95e97e5e(ints, 1, 8)
	var v any = any(go__slice_length__int_95e97e5e(ints))
	hxrt.Println(v)
	var v_1 any = any(go__slice_get__int_95e97e5e(ints, 1))
	hxrt.Println(v_1)
	intsArray := go__slice_toArray__int_95e97e5e(ints)
	hxrt.Println(any(intsArray[0]))
	words := go___Go_newSlice()
	go__slice_push___string_f613ccd0(words, hxrt.StringFromLiteral("go"))
	go__slice_push___string_f613ccd0(words, hxrt.StringFromLiteral("haxe"))
	var v_2 any = any(go__slice_get___string_f613ccd0(words, 0))
	hxrt.Println(v_2)
	wordsArray := go__slice_toArray___string_f613ccd0(words)
	var v_3 any = any(len(wordsArray))
	hxrt.Println(v_3)
	scores := New_go___Map()
	go__map_set__int___string_d6952de3(scores, 7, hxrt.StringFromLiteral("seven"))
	var v_4 any = any(go__map_exists__int___string_d6952de3(scores, 7))
	hxrt.Println(v_4)
	var v_5 any = any(func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(go__map_get__int___string_d6952de3(scores, 7)))
	hxrt.Println(v_5)
	missing := func(hx_value_3 any) *string {
		if hx_value_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_value_3.(*string)
	}(go__map_get__int___string_d6952de3(scores, 99))
	hxrt.Println(func() any {
		var hx_if_5 any
		if hxrt.StringEqualStringPtr(missing, nil) {
			hx_if_5 = hxrt.StringFromLiteral("none")
		} else {
			hx_if_5 = missing
		}
		return hx_if_5
	}())
	byName := go___Go_newMap()
	go__map_set___string__int_e8ed7ec7(byName, hxrt.StringFromLiteral("alice"), 11)
	var v_6 any = any(go__map_exists___string__int_e8ed7ec7(byName, hxrt.StringFromLiteral("alice")))
	hxrt.Println(v_6)
	var v_7 any = any(func(hx_value_6 any) any {
		if hx_value_6 == nil {
			return nil
		}
		return hx_value_6.(int)
	}(go__map_get___string__int_e8ed7ec7(byName, hxrt.StringFromLiteral("alice"))))
	hxrt.Println(v_7)
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

func go__slice_push___string_f613ccd0(slice *go___Slice, value *string) {
	slice.data = append(slice.data, value)
}

func go__slice_set___string_f613ccd0(slice *go___Slice, index int, value *string) {
	slice.data[index] = value
}

func go__slice_get___string_f613ccd0(slice *go___Slice, index int) *string {
	raw := slice.data[index]
	if raw == nil {
		var zero *string
		return zero
	}
	return raw.(*string)
}

func go__slice_length___string_f613ccd0(slice *go___Slice) int {
	return len(slice.data)
}

func go__slice_toArray___string_f613ccd0(slice *go___Slice) []*string {
	raw := slice.data
	out := make([]*string, len(raw))
	for idx, value := range raw {
		if value == nil {
			continue
		}
		out[idx] = value.(*string)
	}
	return out
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

func go__map_set__int___string_d6952de3(mapValue *go___Map, key int, value *string) {
	mapValue.inner.set(hxrt.StdString(any(key)), value)
}

func go__map_get__int___string_d6952de3(mapValue *go___Map, key int) any {
	return mapValue.inner.get(hxrt.StdString(any(key)))
}

func go__map_exists__int___string_d6952de3(mapValue *go___Map, key int) bool {
	return mapValue.inner.exists(hxrt.StdString(any(key)))
}
