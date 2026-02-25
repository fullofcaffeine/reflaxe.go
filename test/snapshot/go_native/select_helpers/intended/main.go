package main

import (
	"reflect"
	"snapshot/hxrt"
)

func main() {
	gate := go___Go_newChan(1)
	_g := go___Select_send_Int(gate, 7)
	var hx_switch_1 *string
	switch _g.tag {
	case 0:
		hx_switch_1 = hxrt.StringFromLiteral("sent")
	case 1:
		hx_switch_1 = hxrt.StringFromLiteral("default")
	}
	sendFirst := hx_switch_1
	_g_1 := go___Select_send_Int(gate, 8)
	var hx_switch_2 *string
	switch _g_1.tag {
	case 0:
		hx_switch_2 = hxrt.StringFromLiteral("sent")
	case 1:
		hx_switch_2 = hxrt.StringFromLiteral("default")
	}
	sendSecond := hx_switch_2
	_g_2 := go___Select_recv_Int(gate)
	var hx_switch_3 *string
	switch _g_2.tag {
	case 0:
		_g_3 := _g_2.params[0].(int)
		value := _g_3
		hx_switch_3 = hxrt.StringConcatAny(hxrt.StringFromLiteral("recv:"), value)
	case 1:
		hx_switch_3 = hxrt.StringFromLiteral("empty")
	}
	recvFirst := hx_switch_3
	_g_4 := go___Select_recv_Int(gate)
	var hx_switch_4 *string
	switch _g_4.tag {
	case 0:
		_g_5 := _g_4.params[0].(int)
		value_1 := _g_5
		hx_switch_4 = hxrt.StringConcatAny(hxrt.StringFromLiteral("recv:"), value_1)
	case 1:
		hx_switch_4 = hxrt.StringFromLiteral("empty")
	}
	recvSecond := hx_switch_4
	hxrt.Println(sendFirst)
	hxrt.Println(sendSecond)
	hxrt.Println(recvFirst)
	hxrt.Println(recvSecond)
	left := go___Go_newChan(1)
	right := go___Go_newChan(1)
	right.send(hxrt.StringFromLiteral("beta"))
	_g_6 := go___Select_recv2_String_String(left, right)
	var hx_switch_5 *string
	switch _g_6.tag {
	case 0:
		_g_7 := _g_6.params[0].(*string)
		value_2 := _g_7
		hx_switch_5 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("left:"), value_2)
	case 1:
		_g_8 := _g_6.params[0].(*string)
		value_3 := _g_8
		hx_switch_5 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("right:"), value_3)
	case 2:
		hx_switch_5 = hxrt.StringFromLiteral("none")
	}
	recvTwo := hx_switch_5
	hxrt.Println(recvTwo)
	sendTwoA := go___Go_newChan(1)
	sendTwoB := go___Go_newChan(1)
	_g_9 := go___Select_send2_Int_Int(sendTwoA, 11, sendTwoB, 22)
	var hx_switch_6 *string
	switch _g_9.tag {
	case 0:
		hx_switch_6 = hxrt.StringFromLiteral("a")
	case 1:
		hx_switch_6 = hxrt.StringFromLiteral("b")
	case 2:
		hx_switch_6 = hxrt.StringFromLiteral("none")
	}
	sendTwo := hx_switch_6
	sendTwoValues := hxrt.StringConcatAny(hxrt.StringConcatAny(func(hx_value_7 any) int {
		if hx_value_7 == nil {
			var hx_zero_8 int
			return hx_zero_8
		}
		return hx_value_7.(int)
	}(sendTwoA.recvOr(-1)), hxrt.StringFromLiteral(",")), func(hx_value_9 any) int {
		if hx_value_9 == nil {
			var hx_zero_10 int
			return hx_zero_10
		}
		return hx_value_9.(int)
	}(sendTwoB.recvOr(-1)))
	hxrt.Println(sendTwo)
	hxrt.Println(sendTwoValues)
}

type haxe__ds__IntMap struct {
	h map[int]any
}

type haxe__ds__StringMap struct {
	h map[string]any
}

type haxe__ds__ObjectMap struct {
	h map[any]any
}

type haxe__ds__EnumValueMap struct {
	h map[any]any
}

type haxe__ds__List struct {
	items  []any
	length int
}

func New_haxe__ds__IntMap() *haxe__ds__IntMap {
	return &haxe__ds__IntMap{h: map[int]any{}}
}

func (self *haxe__ds__IntMap) set(key int, value any) {
	self.h[key] = value
}

func (self *haxe__ds__IntMap) get(key int) any {
	value := self.h[key]
	return value
}

func (self *haxe__ds__IntMap) exists(key int) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__IntMap) remove(key int) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	return &haxe__ds__StringMap{h: map[string]any{}}
}

func (self *haxe__ds__StringMap) set(key *string, value any) {
	self.h[*hxrt.StdString(key)] = value
}

func (self *haxe__ds__StringMap) get(key *string) any {
	value := self.h[*hxrt.StdString(key)]
	return value
}

func (self *haxe__ds__StringMap) exists(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	return ok
}

func (self *haxe__ds__StringMap) remove(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	delete(self.h, *hxrt.StdString(key))
	return ok
}

func New_haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	return &haxe__ds__ObjectMap{h: map[any]any{}}
}

func (self *haxe__ds__ObjectMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__ObjectMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__ObjectMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__ObjectMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__EnumValueMap() *haxe__ds__EnumValueMap {
	return &haxe__ds__EnumValueMap{h: map[any]any{}}
}

func (self *haxe__ds__EnumValueMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__EnumValueMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__EnumValueMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__EnumValueMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__List() *haxe__ds__List {
	return &haxe__ds__List{items: []any{}, length: 0}
}

func (self *haxe__ds__List) add(item any) {
	self.items = append(self.items, item)
	self.length = len(self.items)
}

func (self *haxe__ds__List) push(item any) {
	self.items = append([]any{item}, self.items...)
	self.length = len(self.items)
}

func (self *haxe__ds__List) pop() any {
	if len(self.items) == 0 {
		return nil
	}
	head := self.items[0]
	self.items = self.items[1:]
	self.length = len(self.items)
	return head
}

func (self *haxe__ds__List) first() any {
	if len(self.items) == 0 {
		return nil
	}
	return self.items[0]
}

func (self *haxe__ds__List) last() any {
	size := len(self.items)
	if size == 0 {
		return nil
	}
	return self.items[(size - 1)]
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
	chosen, recvValue, _ := reflect.Select(cases)
	if chosen == 0 {
		if !recvValue.IsValid() {
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
	chosen, recvValue, _ := reflect.Select(cases)
	if chosen == 0 {
		if !recvValue.IsValid() {
			return New_go___Result(nil, nil)
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
