package main

import "snapshot/hxrt"

func errValue() *go___Result {
	return go___Go_fail(hxrt.StringFromLiteral("bad"))
}

func main() {
	ok := okValue()
	_ = ok
	hxrt.Println(ok.__hx_this.isOk())
	hxrt.Println(ok.__hx_this.unwrap())
	err := errValue()
	_ = err
	hxrt.Println(err.__hx_this.isErr())
	hxrt.Println(err.__hx_this.error())
}

func okValue() *go___Result {
	return go___Go_ok(12)
}

type I_go___Chan interface {
	send(value any)
	recv() any
}

type go___Chan struct {
	__hx_this I_go___Chan
	queue     []any
	readIndex int
}

func New_go___Chan() *go___Chan {
	self := &go___Chan{}
	self.__hx_this = self
	self.queue = []any{}
	self.readIndex = 0
	return self
}

func (self *go___Chan) send(value any) {
	self.queue = append(self.queue, value)
}

func (self *go___Chan) recv() any {
	if self.readIndex >= len(self.queue) {
		return nil
	}
	var value any = self.queue[self.readIndex]
	_ = value
	self.readIndex = int(int32((self.readIndex + 1)))
	return value
}

type I_go___Error interface {
	toString() *string
}

type go___Error struct {
	__hx_this I_go___Error
	message   *string
}

func New_go___Error(message *string) *go___Error {
	self := &go___Error{}
	self.__hx_this = self
	self.message = message
	return self
}

func (self *go___Error) toString() *string {
	return self.message
}

func go___Go_fail(message *string) *go___Result {
	return go___Result_failure(message)
}

func go___Go_newChan() *go___Chan {
	return New_go___Chan()
}

func go___Go_newMap() *go___Map {
	return New_go___Map()
}

func go___Go_newSlice() *go___Slice {
	return New_go___Slice()
}

func go___Go_ok(value any) *go___Result {
	return go___Result_ok(value)
}

func go___Go_spawn(fn func()) {
	fn()
}

type I_go___Map interface {
	set(key any, value any)
	get(key any) any
	exists(key any) bool
}

type go___Map struct {
	__hx_this I_go___Map
	inner     *haxe__ds__StringMap
}

func New_go___Map() *go___Map {
	self := &go___Map{}
	self.__hx_this = self
	self.inner = New_haxe__ds__StringMap()
	return self
}

func (self *go___Map) set(key any, value any) {
	self.inner.set(hxrt.StdString(key), value)
}

func (self *go___Map) get(key any) any {
	return self.inner.get(hxrt.StdString(key))
}

func (self *go___Map) exists(key any) bool {
	return self.inner.exists(hxrt.StdString(key))
}

type I_go___Result interface {
	isOk() bool
	isErr() bool
	unwrap() any
	error() *string
}

type go___Result struct {
	__hx_this  I_go___Result
	value      any
	errorValue *go___Error
}

func New_go___Result(value any, errorValue *go___Error) *go___Result {
	self := &go___Result{}
	self.__hx_this = self
	self.value = value
	self.errorValue = errorValue
	return self
}

func (self *go___Result) isOk() bool {
	return hxrt.StringEqualAny(self.errorValue, nil)
}

func (self *go___Result) isErr() bool {
	return !hxrt.StringEqualAny(self.errorValue, nil)
}

func (self *go___Result) unwrap() any {
	if !hxrt.StringEqualAny(self.errorValue, nil) {
		hxrt.Throw(self.errorValue.__hx_this.toString())
		var hx_throw_zero_1 any
		return hx_throw_zero_1
	}
	return self.value
}

func (self *go___Result) error() *string {
	if hxrt.StringEqualAny(self.errorValue, nil) {
		return nil
	}
	return self.errorValue.__hx_this.toString()
}

func go___Result_failure(message *string) *go___Result {
	return New_go___Result(nil, New_go___Error(message))
}

func go___Result_ok(value any) *go___Result {
	return New_go___Result(value, nil)
}

type I_go___Slice interface {
	get_length() int
	push(value any)
	get(index int) any
	set(index int, value any)
	toArray() []any
}

type go___Slice struct {
	__hx_this I_go___Slice
	data      []any
	length    int
}

func New_go___Slice() *go___Slice {
	self := &go___Slice{}
	self.__hx_this = self
	self.data = []any{}
	return self
}

func (self *go___Slice) get_length() int {
	return len(self.data)
}

func (self *go___Slice) push(value any) {
	self.data = append(self.data, value)
}

func (self *go___Slice) get(index int) any {
	return self.data[index]
}

func (self *go___Slice) set(index int, value any) {
	self.data[index] = value
}

func (self *go___Slice) toArray() []any {
	return self.data
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
