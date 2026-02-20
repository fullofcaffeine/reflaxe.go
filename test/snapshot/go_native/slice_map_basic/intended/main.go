package main

import "snapshot/hxrt"

func main() {
	s := New_go___Slice()
	_ = s
	s.__hx_this.push(1)
	s.__hx_this.push(2)
	s.__hx_this.push(3)
	s.__hx_this.set(1, 7)
	hxrt.Println(s.__hx_this.get_length())
	hxrt.Println(s.__hx_this.get(1))
	m := New_go___Map()
	_ = m
	m.__hx_this.set(42, hxrt.StringFromLiteral("answer"))
	hxrt.Println(m.__hx_this.exists(42))
	hxrt.Println(m.__hx_this.get(42))
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
