package main

import (
	"errors"
	"snapshot/hxrt"
)

func errResult() *go___Result {
	return go__result_failure__int_95e97e5e(hxrt.StringFromLiteral("boom"))
}

func main() {
	ok := okResult()
	hxrt.Println(go__result_isOk__int_95e97e5e(ok))
	hxrt.Println(go__result_unwrap__int_95e97e5e(ok))
	err := errResult()
	hxrt.Println(go__result_isErr__int_95e97e5e(err))
	hxrt.Println(go__result_error__int_95e97e5e(err))
	hxrt.TryCatch(func() {
		go__result_unwrap__int_95e97e5e(err)
		hxrt.Println(hxrt.StringFromLiteral("unexpected"))
	}, func(hx_caught_1 any) {
		e := hx_caught_1
		_ = e
		hxrt.Println(hxrt.StringFromLiteral("caught"))
		hxrt.Println(hxrt.StdString(e))
	})
	okViaGo := go__result_ok___string_f613ccd0(hxrt.StringFromLiteral("done"))
	hxrt.Println(go__result_unwrap___string_f613ccd0(okViaGo))
	errViaGo := go__result_failure___string_f613ccd0(hxrt.StringFromLiteral("bad"))
	hxrt.Println(go__result_error___string_f613ccd0(errViaGo))
}

func okResult() *go___Result {
	return go__result_ok__int_95e97e5e(7)
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
	channel.(chan any) <- value
}

func go__concurrency_trySend(channel any, value any) bool {
	select {
	case channel.(chan any) <- value:
		return true
	default:
		return false
	}
}

func go__concurrency_recv(channel any) any {
	return <-channel.(chan any)
}

func go__concurrency_recvOr(channel any, defaultValue any) any {
	select {
	case value := <-channel.(chan any):
		return value
	default:
		return defaultValue
	}
}

func go__concurrency_tryRecv(channel any) *go___Result {
	select {
	case value := <-channel.(chan any):
		return New_go___Result(value, nil)
	default:
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("empty")))
	}
}

func go__concurrency_close(channel any) {
	close(channel.(chan any))
}

func go__concurrency_spawn(fn func()) {
	go fn()
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
