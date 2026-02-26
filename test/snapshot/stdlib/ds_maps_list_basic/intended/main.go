package main

import "snapshot/hxrt"

type EKey struct {
	tag    int
	params []any
}

var EKey_A *EKey = &EKey{tag: 0}

func EKey_B(v int) *EKey {
	enumValue := &EKey{tag: 1}
	enumValue.params = []any{v}
	return enumValue
}

type I_Box interface {
}

type Box struct {
	__hx_this I_Box
	id        int
}

func New_Box(id int) *Box {
	self := &Box{}
	self.__hx_this = self
	self.id = id
	return self
}

func main() {
	sm := New_haxe__ds__StringMap()
	sm.set(hxrt.StringFromLiteral("a"), 1)
	av := func(hx_value_1 any) any {
		if hx_value_1 == nil {
			return nil
		}
		return hx_value_1.(int)
	}(sm.get(hxrt.StringFromLiteral("a")))
	hxrt.Println(av)
	om := New_haxe__ds__ObjectMap()
	box := New_Box(7)
	om.set(box, hxrt.StringFromLiteral("box"))
	ov := func(hx_value_2 any) *string {
		if hx_value_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_value_2.(*string)
	}(om.get(box))
	hxrt.Println(ov)
	em := New_haxe__ds__EnumValueMap()
	em.set(EKey_A, hxrt.StringFromLiteral("enum"))
	ev := func(hx_value_4 any) *string {
		if hx_value_4 == nil {
			var hx_zero_5 *string
			return hx_zero_5
		}
		return hx_value_4.(*string)
	}(em.get(EKey_A))
	hxrt.Println(ev)
	list := New_haxe__ds__List()
	list.add(4)
	list.add(5)
	hxrt.Println(list.length)
	hxrt.Println(func(hx_value_6 any) any {
		if hx_value_6 == nil {
			return nil
		}
		return hx_value_6.(int)
	}(list.first()))
	hxrt.Println(func(hx_value_7 any) any {
		if hx_value_7 == nil {
			return nil
		}
		return hx_value_7.(int)
	}(list.last()))
	hxrt.Println(func(hx_value_8 any) any {
		if hx_value_8 == nil {
			return nil
		}
		return hx_value_8.(int)
	}(list.pop()))
	hxrt.Println(list.length)
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
