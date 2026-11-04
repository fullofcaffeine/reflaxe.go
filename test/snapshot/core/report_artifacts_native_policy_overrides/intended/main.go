package main

import (
	"reflect"
	"snapshot/hxrt"
)

func main() {
	channel := go__concurrency_newChan__int_95e97e5e(1)
	go__concurrency_send__int_95e97e5e(channel.__hx_native, 7)
	var v any = any(go__concurrency_recv__int_95e97e5e(channel.__hx_native))
	hxrt.Println(v)
	go__concurrency_close__int_95e97e5e(channel.__hx_native)
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

func hx_recover_callable_0(value any) *hxrt.Function[func() bool] {
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
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() any] {
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
	case *hxrt.Function[func() *string]:
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

func hx_recover_callable_2(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *go___Result {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *go___Result:
		return source
	default:
		return value.(*go___Result)
	}
}

func hx_recover_class_1(value any) *haxe__ds__StringMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__StringMap:
		return source
	default:
		return value.(*haxe__ds__StringMap)
	}
}

func hx_recover_class_2(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_3(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}
