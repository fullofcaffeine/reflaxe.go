package main

import "snapshot/hxrt"

type I_haxe__ds___List__GoListKeyValueIterator interface {
	hasNext() bool
	next() map[string]any
}

type haxe__ds___List__GoListKeyValueIterator struct {
	__hx_this I_haxe__ds___List__GoListKeyValueIterator
	items     []any
	index     int
}

func New_haxe__ds___List__GoListKeyValueIterator(items []any) *haxe__ds___List__GoListKeyValueIterator {
	self := &haxe__ds___List__GoListKeyValueIterator{}
	self.__hx_this = self
	self.items = items
	self.index = 0
	return self
}

func (self *haxe__ds___List__GoListKeyValueIterator) hasNext() bool {
	return (self.index < len(self.items))
}

func (self *haxe__ds___List__GoListKeyValueIterator) next() map[string]any {
	hx_post_30 := self.index
	self.index = int(int32((self.index + 1)))
	key := hx_post_30
	hx_obj_31 := map[string]any{}
	hx_obj_31["key"] = key
	hx_obj_31["value"] = self.items[key]
	return hx_obj_31
}

type I_haxe__ds___List__GoListIterator interface {
	hasNext() bool
	next() any
}

type haxe__ds___List__GoListIterator struct {
	__hx_this I_haxe__ds___List__GoListIterator
	items     []any
	index     int
}

func New_haxe__ds___List__GoListIterator(items []any) *haxe__ds___List__GoListIterator {
	self := &haxe__ds___List__GoListIterator{}
	self.__hx_this = self
	self.items = items
	self.index = 0
	return self
}

func (self *haxe__ds___List__GoListIterator) hasNext() bool {
	return (self.index < len(self.items))
}

func (self *haxe__ds___List__GoListIterator) next() any {
	hx_post_32 := self.index
	self.index = int(int32((self.index + 1)))
	return self.items[hx_post_32]
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
	items     []any
	length    int
}

func New_haxe__ds__List() *haxe__ds__List {
	self := &haxe__ds__List{}
	self.__hx_this = self
	self.items = []any{}
	self.length = 0
	return self
}

func (self *haxe__ds__List) add(item any) {
	hx_arr_33 := self.items
	hx_arr_33 = append(hx_arr_33, item)
	self.items = hx_arr_33
	self.length = int(int32((self.length + 1)))
}

func (self *haxe__ds__List) push(item any) {
	next := []any{item}
	_g := 0
	_g1 := self.items
	for _g < len(_g1) {
		var existing any = _g1[_g]
		_g = int(int32((_g + 1)))
		next = append(next, existing)
	}
	self.items = next
	self.length = int(int32((self.length + 1)))
}

func (self *haxe__ds__List) first() any {
	var hx_if_35 any
	if self.length == 0 {
		hx_if_35 = nil
	} else {
		hx_if_35 = self.items[0]
	}
	return hx_if_35
}

func (self *haxe__ds__List) last() any {
	var hx_if_36 any
	if self.length == 0 {
		hx_if_36 = nil
	} else {
		hx_if_36 = self.items[int(int32((hxrt.Int32Wrap(self.length) - hxrt.Int32Wrap(1))))]
	}
	return hx_if_36
}

func (self *haxe__ds__List) pop() any {
	if self.length == 0 {
		return nil
	}
	var first any = self.items[0]
	remaining := []any{}
	_g := 1
	_g1 := len(self.items)
	for _g < _g1 {
		hx_post_37 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_37
		remaining = append(remaining, self.items[index])
	}
	self.items = remaining
	self.length = int(int32((self.length - 1)))
	return first
}

func (self *haxe__ds__List) isEmpty() bool {
	return (self.length == 0)
}

func (self *haxe__ds__List) clear() {
	self.items = []any{}
	self.length = 0
}

func (self *haxe__ds__List) remove(value any) bool {
	_g := 0
	_g1 := len(self.items)
	for _g < _g1 {
		hx_post_39 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_39
		if haxe__ds__List_sameValue(self.items[index], value) {
			remaining := []any{}
			_g_1 := 0
			_g1_1 := len(self.items)
			for _g_1 < _g1_1 {
				hx_post_40 := _g_1
				_g_1 = int(int32((_g_1 + 1)))
				copyIndex := hx_post_40
				if copyIndex != index {
					remaining = append(remaining, self.items[copyIndex])
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
	rendered := []*string{}
	_g := 0
	_g1 := len(self.items)
	for _g < _g1 {
		hx_post_42 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_42
		rendered = append(rendered, hxrt.StdString(self.items[index]))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{"), hxrt.StringJoinAny(func(hx_sort_src_44 []*string) []any {
		hx_sort_out_46 := make([]any, 0, len(hx_sort_src_44))
		for _, hx_sort_item_45 := range hx_sort_src_44 {
			hx_sort_out_46 = append(hx_sort_out_46, hx_sort_item_45)
		}
		return hx_sort_out_46
	}(rendered), hxrt.StringFromLiteral(", "))), hxrt.StringFromLiteral("}"))
}

func (self *haxe__ds__List) join(separator *string) *string {
	rendered := []*string{}
	_g := 0
	_g1 := len(self.items)
	for _g < _g1 {
		hx_post_47 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_47
		rendered = append(rendered, hxrt.StdString(self.items[index]))
	}
	return hxrt.StringJoinAny(func(hx_sort_src_49 []*string) []any {
		hx_sort_out_51 := make([]any, 0, len(hx_sort_src_49))
		for _, hx_sort_item_50 := range hx_sort_src_49 {
			hx_sort_out_51 = append(hx_sort_out_51, hx_sort_item_50)
		}
		return hx_sort_out_51
	}(rendered), separator)
}

func (self *haxe__ds__List) filter(predicate func(any) bool) *haxe__ds__List {
	filtered := New_haxe__ds__List()
	source := func(hx_value_52 any) *haxe__ds___List__GoListIterator {
		if hx_value_52 == nil {
			var hx_zero_53 *haxe__ds___List__GoListIterator
			return hx_zero_53
		}
		return hx_value_52.(*haxe__ds___List__GoListIterator)
	}(self.iterator())
	for func(hx_value_54 any) bool {
		if hx_value_54 == nil {
			var hx_zero_55 bool
			return hx_zero_55
		}
		return hx_value_54.(bool)
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
	source := func(hx_value_56 any) *haxe__ds___List__GoListIterator {
		if hx_value_56 == nil {
			var hx_zero_57 *haxe__ds___List__GoListIterator
			return hx_zero_57
		}
		return hx_value_56.(*haxe__ds___List__GoListIterator)
	}(self.iterator())
	for func(hx_value_58 any) bool {
		if hx_value_58 == nil {
			var hx_zero_59 bool
			return hx_zero_59
		}
		return hx_value_58.(bool)
	}(source.hasNext()) {
		var item any = source.next()
		mapped.add(transform(item))
	}
	return mapped
}

func haxe__ds__List_sameValue(left any, right any) bool {
	if hxrt.AnyEqualsNull(left) || hxrt.AnyEqualsNull(right) {
		return (left == right)
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
	return (left == right)
}
