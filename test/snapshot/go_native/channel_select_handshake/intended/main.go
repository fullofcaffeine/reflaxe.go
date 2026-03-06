package main

import (
	"reflect"
	"snapshot/hxrt"
)

func main() {
	requests := go__concurrency_newChan__int_95e97e5e(0)
	responses := go__concurrency_newChan__int_95e97e5e(0)
	go___Go_spawn(func() {
		value := go__concurrency_recv__int_95e97e5e(requests.__hx_native)
		go__concurrency_send__int_95e97e5e(responses.__hx_native, value)
	})
	go__concurrency_send__int_95e97e5e(requests.__hx_native, 41)
	hxrt.Println(go__concurrency_recv__int_95e97e5e(responses.__hx_native))
	buffered := go__concurrency_newChan__int_95e97e5e(1)
	hxrt.Println(go__concurrency_trySend__int_95e97e5e(buffered.__hx_native, 7))
	hxrt.Println(go__concurrency_trySend__int_95e97e5e(buffered.__hx_native, 8))
	hxrt.Println(go__concurrency_recvOr__int_95e97e5e(buffered.__hx_native, -1))
	hxrt.Println(go__concurrency_recvOr__int_95e97e5e(buffered.__hx_native, -1))
	go__concurrency_close__int_95e97e5e(buffered.__hx_native)
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

func (self *haxe__ds__IntMap) set(key any, value any) {
	resolvedKey := hxrt.IntFromNullableAny(key)
	self.h[resolvedKey] = value
}

func (self *haxe__ds__IntMap) get(key any) any {
	resolvedKey := hxrt.IntFromNullableAny(key)
	value := self.h[resolvedKey]
	return value
}

func (self *haxe__ds__IntMap) exists(key any) bool {
	resolvedKey := hxrt.IntFromNullableAny(key)
	_, ok := self.h[resolvedKey]
	return ok
}

func (self *haxe__ds__IntMap) remove(key any) bool {
	resolvedKey := hxrt.IntFromNullableAny(key)
	_, ok := self.h[resolvedKey]
	delete(self.h, resolvedKey)
	return ok
}

func (self *haxe__ds__IntMap) keys() map[string]any {
	keys := make([]int, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() int { key := keys[index]; index++; return key }
	return iter
}

func (self *haxe__ds__IntMap) iterator() map[string]any {
	keys := make([]int, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return self.h[key] }
	return iter
}

func (self *haxe__ds__IntMap) keyValueIterator() map[string]any {
	keys := make([]int, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() map[string]any {
		key := keys[index]
		index++
		return map[string]any{"key": key, "value": self.h[key]}
	}
	return iter
}

func (self *haxe__ds__IntMap) copyIMap() haxe__IMap {
	copied := New_haxe__ds__IntMap()
	for key, value := range self.h {
		copied.h[key] = value
	}
	return copied
}

func (self *haxe__ds__IntMap) toString() *string {
	return hxrt.StringFromLiteral("{}")
}

func (self *haxe__ds__IntMap) clear() {
	self.h = map[int]any{}
}

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	return &haxe__ds__StringMap{h: map[string]any{}}
}

func (self *haxe__ds__StringMap) set(key any, value any) {
	self.h[*hxrt.StdString(key)] = value
}

func (self *haxe__ds__StringMap) get(key any) any {
	value := self.h[*hxrt.StdString(key)]
	return value
}

func (self *haxe__ds__StringMap) exists(key any) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	return ok
}

func (self *haxe__ds__StringMap) remove(key any) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	delete(self.h, *hxrt.StdString(key))
	return ok
}

func (self *haxe__ds__StringMap) keys() map[string]any {
	keys := make([]string, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() *string { key := keys[index]; index++; return hxrt.StringFromLiteral(key) }
	return iter
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	keys := make([]string, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return self.h[key] }
	return iter
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	keys := make([]string, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() map[string]any {
		key := keys[index]
		index++
		return map[string]any{"key": hxrt.StringFromLiteral(key), "value": self.h[key]}
	}
	return iter
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	copied := New_haxe__ds__StringMap()
	for key, value := range self.h {
		copied.h[key] = value
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	return hxrt.StringFromLiteral("{}")
}

func (self *haxe__ds__StringMap) clear() {
	self.h = map[string]any{}
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

func (self *haxe__ds__ObjectMap) keys() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return key }
	return iter
}

func (self *haxe__ds__ObjectMap) iterator() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return self.h[key] }
	return iter
}

func (self *haxe__ds__ObjectMap) keyValueIterator() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() map[string]any {
		key := keys[index]
		index++
		return map[string]any{"key": key, "value": self.h[key]}
	}
	return iter
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	copied := New_haxe__ds__ObjectMap()
	for key, value := range self.h {
		copied.h[key] = value
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	return hxrt.StringFromLiteral("{}")
}

func (self *haxe__ds__ObjectMap) clear() {
	self.h = map[any]any{}
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

func (self *haxe__ds__EnumValueMap) keys() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return key }
	return iter
}

func (self *haxe__ds__EnumValueMap) iterator() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return self.h[key] }
	return iter
}

func (self *haxe__ds__EnumValueMap) keyValueIterator() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() map[string]any {
		key := keys[index]
		index++
		return map[string]any{"key": key, "value": self.h[key]}
	}
	return iter
}

func (self *haxe__ds__EnumValueMap) copyIMap() haxe__IMap {
	copied := New_haxe__ds__EnumValueMap()
	for key, value := range self.h {
		copied.h[key] = value
	}
	return copied
}

func (self *haxe__ds__EnumValueMap) toString() *string {
	return hxrt.StringFromLiteral("{}")
}

func (self *haxe__ds__EnumValueMap) clear() {
	self.h = map[any]any{}
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
	case value := <-channel.(chan int):
		return value
	default:
		return defaultValue
	}
}

func go__concurrency_tryRecv__int_95e97e5e(channel any) *go___Result {
	select {
	case value := <-channel.(chan int):
		return New_go___Result(value, nil)
	default:
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("empty")))
	}
}

func go__concurrency_close__int_95e97e5e(channel any) {
	close(channel.(chan int))
}
