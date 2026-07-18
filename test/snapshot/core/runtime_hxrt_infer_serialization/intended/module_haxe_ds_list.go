package main

import "snapshot/hxrt"

type I_haxe__ds___List__GoListKeyValueIterator interface {
	hasNext() bool
	next() map[string]any
}

type haxe__ds___List__GoListKeyValueIterator struct {
	__hx_this I_haxe__ds___List__GoListKeyValueIterator
	items     *hxrt.Array
	index     int
}

func New_haxe__ds___List__GoListKeyValueIterator(items *hxrt.Array) *haxe__ds___List__GoListKeyValueIterator {
	self := &haxe__ds___List__GoListKeyValueIterator{}
	self.__hx_this = self
	self.items = items
	self.index = 0
	return self
}

func (self *haxe__ds___List__GoListKeyValueIterator) hasNext() bool {
	return (self.index < self.items.Len())
}

func (self *haxe__ds___List__GoListKeyValueIterator) next() map[string]any {
	hx_post_2 := self.index
	self.index = int(int32((self.index + 1)))
	key := hx_post_2
	hx_obj_3 := map[string]any{}
	hx_obj_3["key"] = key
	hx_obj_3["value"] = self.items.Get(key)
	return hx_obj_3
}

type I_haxe__ds___List__GoListIterator interface {
	hasNext() bool
	next() any
}

type haxe__ds___List__GoListIterator struct {
	__hx_this I_haxe__ds___List__GoListIterator
	items     *hxrt.Array
	index     int
}

func New_haxe__ds___List__GoListIterator(items *hxrt.Array) *haxe__ds___List__GoListIterator {
	self := &haxe__ds___List__GoListIterator{}
	self.__hx_this = self
	self.items = items
	self.index = 0
	return self
}

func (self *haxe__ds___List__GoListIterator) hasNext() bool {
	return (self.index < self.items.Len())
}

func (self *haxe__ds___List__GoListIterator) next() any {
	hx_post_4 := self.index
	self.index = int(int32((self.index + 1)))
	return self.items.Get(hx_post_4)
}

type I_haxe__ds__List interface {
	add(item any)
	push(item any)
	first() any
	last() any
	pop() any
	isEmpty() bool
	clear()
	remove(value any) bool
	iterator() *haxe__ds___List__GoListIterator
	keyValueIterator() *haxe__ds___List__GoListKeyValueIterator
	toString() *string
	join(separator *string) *string
	filter(predicate func(any) bool) *haxe__ds__List
	map_(transform func(any) any) *haxe__ds__List
}

type haxe__ds__List struct {
	__hx_this I_haxe__ds__List
	items     *hxrt.Array
	length    int
}

func New_haxe__ds__List() *haxe__ds__List {
	self := &haxe__ds__List{}
	self.__hx_this = self
	self.items = hxrt.NewArray()
	self.length = 0
	return self
}

func (self *haxe__ds__List) add(item any) {
	hx_arr_91 := self.items
	hx_arr_91.Push(item)
	self.length = int(int32((self.length + 1)))
}

func (self *haxe__ds__List) push(item any) {
	next := hxrt.NewArray(item)
	_g := 0
	_g1 := self.items
	for _g < _g1.Len() {
		var existing any = _g1.Get(_g)
		_g = int(int32((_g + 1)))
		next.Push(existing)
	}
	self.items = next
	self.length = int(int32((self.length + 1)))
}

func (self *haxe__ds__List) first() any {
	var hx_if_93 any
	if self.length == 0 {
		hx_if_93 = nil
	} else {
		hx_if_93 = self.items.Get(0)
	}
	return hx_if_93
}

func (self *haxe__ds__List) last() any {
	var hx_if_94 any
	if self.length == 0 {
		hx_if_94 = nil
	} else {
		hx_if_94 = self.items.Get(int(int32((hxrt.Int32Wrap(self.length) - hxrt.Int32Wrap(1)))))
	}
	return hx_if_94
}

func (self *haxe__ds__List) pop() any {
	if self.length == 0 {
		return nil
	}
	var first any = self.items.Get(0)
	remaining := hxrt.NewArray()
	_g := 1
	_g1 := self.items.Len()
	for _g < _g1 {
		hx_post_95 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_95
		remaining.Push(self.items.Get(index))
	}
	self.items = remaining
	self.length = int(int32((self.length - 1)))
	return first
}

func (self *haxe__ds__List) isEmpty() bool {
	return (self.length == 0)
}

func (self *haxe__ds__List) clear() {
	self.items = hxrt.NewArray()
	self.length = 0
}

func (self *haxe__ds__List) remove(value any) bool {
	_g := 0
	_g1 := self.items.Len()
	for _g < _g1 {
		hx_post_97 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_97
		if haxe__ds__List_sameValue(self.items.Get(index), value) {
			remaining := hxrt.NewArray()
			_g_1 := 0
			_g1_1 := self.items.Len()
			for _g_1 < _g1_1 {
				hx_post_98 := _g_1
				_g_1 = int(int32((_g_1 + 1)))
				copyIndex := hx_post_98
				if copyIndex != index {
					remaining.Push(self.items.Get(copyIndex))
				}
			}
			self.items = remaining
			self.length = int(int32((self.length - 1)))
			return true
		}
	}
	return false
}

func (self *haxe__ds__List) iterator() *haxe__ds___List__GoListIterator {
	return New_haxe__ds___List__GoListIterator(self.items)
}

func (self *haxe__ds__List) keyValueIterator() *haxe__ds___List__GoListKeyValueIterator {
	return New_haxe__ds___List__GoListKeyValueIterator(self.items)
}

func (self *haxe__ds__List) toString() *string {
	rendered := hxrt.NewArray()
	_g := 0
	_g1 := self.items.Len()
	for _g < _g1 {
		hx_post_100 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_100
		rendered.Push(hxrt.StdString(self.items.Get(index)))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{"), hxrt.StringJoinAny(rendered.Values(), hxrt.StringFromLiteral(", "))), hxrt.StringFromLiteral("}"))
}

func (self *haxe__ds__List) join(separator *string) *string {
	rendered := hxrt.NewArray()
	_g := 0
	_g1 := self.items.Len()
	for _g < _g1 {
		hx_post_102 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_102
		rendered.Push(hxrt.StdString(self.items.Get(index)))
	}
	return hxrt.StringJoinAny(rendered.Values(), separator)
}

func (self *haxe__ds__List) filter(predicate func(any) bool) *haxe__ds__List {
	filtered := New_haxe__ds__List()
	source := func(hx_value_104 any) *haxe__ds___List__GoListIterator {
		if hx_value_104 == nil {
			var hx_zero_105 *haxe__ds___List__GoListIterator
			return hx_zero_105
		}
		return hx_value_104.(*haxe__ds___List__GoListIterator)
	}(self.iterator())
	for func(hx_value_106 any) bool {
		if hx_value_106 == nil {
			var hx_zero_107 bool
			return hx_zero_107
		}
		return hx_value_106.(bool)
	}(source.hasNext()) {
		var item any = source.next()
		if predicate(item) {
			filtered.add(item)
		}
	}
	return filtered
}

func (self *haxe__ds__List) map_(transform func(any) any) *haxe__ds__List {
	mapped := New_haxe__ds__List()
	source := func(hx_value_108 any) *haxe__ds___List__GoListIterator {
		if hx_value_108 == nil {
			var hx_zero_109 *haxe__ds___List__GoListIterator
			return hx_zero_109
		}
		return hx_value_108.(*haxe__ds___List__GoListIterator)
	}(self.iterator())
	for func(hx_value_110 any) bool {
		if hx_value_110 == nil {
			var hx_zero_111 bool
			return hx_zero_111
		}
		return hx_value_110.(bool)
	}(source.hasNext()) {
		var item any = source.next()
		mapped.add(transform(item))
	}
	return mapped
}

func haxe__ds__List_sameValue(left any, right any) bool {
	if hxrt.AnyEqualsNull(left) || hxrt.AnyEqualsNull(right) {
		return hxrt.HaxeEqual(left, right)
	}
	if func(hx_value any) bool {
		switch hx_value.(type) {
		case *string:
			return true
		case string:
			return true
		default:
			return false
		}
	}(any(left)) || func(hx_value any) bool {
		switch hx_value.(type) {
		case *string:
			return true
		case string:
			return true
		default:
			return false
		}
	}(any(right)) {
		return hxrt.StringEqualStringPtr(hxrt.StdString(left), hxrt.StdString(right))
	}
	return hxrt.HaxeEqual(left, right)
}
