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
	hx_post_279 := self.index
	self.index = int(int32((self.index + 1)))
	key := hx_post_279
	hx_obj_280 := map[string]any{}
	hx_obj_280["key"] = key
	hx_obj_280["value"] = self.items.Get(key)
	return hx_obj_280
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
	hx_post_281 := self.index
	self.index = int(int32((self.index + 1)))
	return self.items.Get(hx_post_281)
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
	hx_arr_368 := self.items
	hx_arr_368.Push(item)
	self.length = int(int32((self.length + 1)))
}

func (self *haxe__ds__List) push(item any) {
	next := hxrt.NewArray(item)
	hx_tmp := 0
	hx_tmp_1 := self.items
	for hx_tmp < hx_tmp_1.Len() {
		var existing any = hx_tmp_1.Get(hx_tmp)
		hx_tmp = int(int32((hx_tmp + 1)))
		next.Push(existing)
	}
	self.items = next
	self.length = int(int32((self.length + 1)))
}

func (self *haxe__ds__List) first() any {
	var hx_if_370 any
	if self.length == 0 {
		hx_if_370 = nil
	} else {
		hx_if_370 = self.items.Get(0)
	}
	return hx_if_370
}

func (self *haxe__ds__List) last() any {
	var hx_if_371 any
	if self.length == 0 {
		hx_if_371 = nil
	} else {
		hx_if_371 = self.items.Get(int(int32((hxrt.Int32Wrap(self.length) - hxrt.Int32Wrap(1)))))
	}
	return hx_if_371
}

func (self *haxe__ds__List) pop() any {
	if self.length == 0 {
		return nil
	}
	var first any = self.items.Get(0)
	remaining := hxrt.NewArray()
	hx_tmp := 1
	hx_tmp_1 := self.items.Len()
	for hx_tmp < hx_tmp_1 {
		hx_post_372 := hx_tmp
		hx_tmp = int(int32((hx_tmp + 1)))
		index := hx_post_372
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
	hx_tmp := 0
	hx_tmp_1 := self.items.Len()
	for hx_tmp < hx_tmp_1 {
		hx_post_374 := hx_tmp
		hx_tmp = int(int32((hx_tmp + 1)))
		index := hx_post_374
		if haxe__ds__List_sameValue(self.items.Get(index), value) {
			remaining := hxrt.NewArray()
			hx_tmp_2 := 0
			hx_tmp_3 := self.items.Len()
			for hx_tmp_2 < hx_tmp_3 {
				hx_post_375 := hx_tmp_2
				hx_tmp_2 = int(int32((hx_tmp_2 + 1)))
				copyIndex := hx_post_375
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
	hx_tmp := 0
	hx_tmp_1 := self.items.Len()
	for hx_tmp < hx_tmp_1 {
		hx_post_377 := hx_tmp
		hx_tmp = int(int32((hx_tmp + 1)))
		index := hx_post_377
		rendered.Push(hxrt.StdString(self.items.Get(index)))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{"), hxrt.StringJoinAny(rendered.Values(), hxrt.StringFromLiteral(", "))), hxrt.StringFromLiteral("}"))
}

func (self *haxe__ds__List) join(separator *string) *string {
	rendered := hxrt.NewArray()
	hx_tmp := 0
	hx_tmp_1 := self.items.Len()
	for hx_tmp < hx_tmp_1 {
		hx_post_379 := hx_tmp
		hx_tmp = int(int32((hx_tmp + 1)))
		index := hx_post_379
		rendered.Push(hxrt.StdString(self.items.Get(index)))
	}
	return hxrt.StringJoinAny(rendered.Values(), separator)
}

func (self *haxe__ds__List) filter(predicate func(any) bool) *haxe__ds__List {
	filtered := New_haxe__ds__List()
	source := func(hx_value_381 any) *haxe__ds___List__GoListIterator {
		if hx_value_381 == nil {
			var hx_zero_382 *haxe__ds___List__GoListIterator
			return hx_zero_382
		}
		return hx_value_381.(*haxe__ds___List__GoListIterator)
	}(self.iterator())
	for func(hx_value_383 any) bool {
		if hx_value_383 == nil {
			var hx_zero_384 bool
			return hx_zero_384
		}
		return hx_value_383.(bool)
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
	source := func(hx_value_385 any) *haxe__ds___List__GoListIterator {
		if hx_value_385 == nil {
			var hx_zero_386 *haxe__ds___List__GoListIterator
			return hx_zero_386
		}
		return hx_value_385.(*haxe__ds___List__GoListIterator)
	}(self.iterator())
	for func(hx_value_387 any) bool {
		if hx_value_387 == nil {
			var hx_zero_388 bool
			return hx_zero_388
		}
		return hx_value_387.(bool)
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
