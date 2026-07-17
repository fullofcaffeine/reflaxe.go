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
	hx_post_237 := self.index
	self.index = int(int32((self.index + 1)))
	key := hx_post_237
	hx_obj_238 := map[string]any{}
	hx_obj_238["key"] = key
	hx_obj_238["value"] = self.items[key]
	return hx_obj_238
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
	hx_post_239 := self.index
	self.index = int(int32((self.index + 1)))
	return self.items[hx_post_239]
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
	hx_arr_326 := self.items
	hx_arr_326 = append(hx_arr_326, item)
	self.items = hx_arr_326
	self.length = int(int32((self.length + 1)))
}

func (self *haxe__ds__List) push(item any) {
	next := []any{item}
	hx_tmp := 0
	hx_tmp_1 := self.items
	for hx_tmp < len(hx_tmp_1) {
		var existing any = hx_tmp_1[hx_tmp]
		hx_tmp = int(int32((hx_tmp + 1)))
		next = append(next, existing)
	}
	self.items = next
	self.length = int(int32((self.length + 1)))
}

func (self *haxe__ds__List) first() any {
	var hx_if_328 any
	if self.length == 0 {
		hx_if_328 = nil
	} else {
		hx_if_328 = self.items[0]
	}
	return hx_if_328
}

func (self *haxe__ds__List) last() any {
	var hx_if_329 any
	if self.length == 0 {
		hx_if_329 = nil
	} else {
		hx_if_329 = self.items[int(int32((hxrt.Int32Wrap(self.length) - hxrt.Int32Wrap(1))))]
	}
	return hx_if_329
}

func (self *haxe__ds__List) pop() any {
	if self.length == 0 {
		return nil
	}
	var first any = self.items[0]
	remaining := []any{}
	hx_tmp := 1
	hx_tmp_1 := len(self.items)
	for hx_tmp < hx_tmp_1 {
		hx_post_330 := hx_tmp
		hx_tmp = int(int32((hx_tmp + 1)))
		index := hx_post_330
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
	hx_tmp := 0
	hx_tmp_1 := len(self.items)
	for hx_tmp < hx_tmp_1 {
		hx_post_332 := hx_tmp
		hx_tmp = int(int32((hx_tmp + 1)))
		index := hx_post_332
		if haxe__ds__List_sameValue(self.items[index], value) {
			remaining := []any{}
			hx_tmp_2 := 0
			hx_tmp_3 := len(self.items)
			for hx_tmp_2 < hx_tmp_3 {
				hx_post_333 := hx_tmp_2
				hx_tmp_2 = int(int32((hx_tmp_2 + 1)))
				copyIndex := hx_post_333
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
	hx_tmp := 0
	hx_tmp_1 := len(self.items)
	for hx_tmp < hx_tmp_1 {
		hx_post_335 := hx_tmp
		hx_tmp = int(int32((hx_tmp + 1)))
		index := hx_post_335
		rendered = append(rendered, hxrt.StdString(self.items[index]))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{"), hxrt.StringJoinAny(func(hx_sort_src_337 []*string) []any {
		hx_sort_out_339 := make([]any, 0, len(hx_sort_src_337))
		for _, hx_sort_item_338 := range hx_sort_src_337 {
			hx_sort_out_339 = append(hx_sort_out_339, hx_sort_item_338)
		}
		return hx_sort_out_339
	}(rendered), hxrt.StringFromLiteral(", "))), hxrt.StringFromLiteral("}"))
}

func (self *haxe__ds__List) join(separator *string) *string {
	rendered := []*string{}
	hx_tmp := 0
	hx_tmp_1 := len(self.items)
	for hx_tmp < hx_tmp_1 {
		hx_post_340 := hx_tmp
		hx_tmp = int(int32((hx_tmp + 1)))
		index := hx_post_340
		rendered = append(rendered, hxrt.StdString(self.items[index]))
	}
	return hxrt.StringJoinAny(func(hx_sort_src_342 []*string) []any {
		hx_sort_out_344 := make([]any, 0, len(hx_sort_src_342))
		for _, hx_sort_item_343 := range hx_sort_src_342 {
			hx_sort_out_344 = append(hx_sort_out_344, hx_sort_item_343)
		}
		return hx_sort_out_344
	}(rendered), separator)
}

func (self *haxe__ds__List) filter(predicate func(any) bool) *haxe__ds__List {
	filtered := New_haxe__ds__List()
	source := func(hx_value_345 any) *haxe__ds___List__GoListIterator {
		if hx_value_345 == nil {
			var hx_zero_346 *haxe__ds___List__GoListIterator
			return hx_zero_346
		}
		return hx_value_345.(*haxe__ds___List__GoListIterator)
	}(self.iterator())
	for func(hx_value_347 any) bool {
		if hx_value_347 == nil {
			var hx_zero_348 bool
			return hx_zero_348
		}
		return hx_value_347.(bool)
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
	source := func(hx_value_349 any) *haxe__ds___List__GoListIterator {
		if hx_value_349 == nil {
			var hx_zero_350 *haxe__ds___List__GoListIterator
			return hx_zero_350
		}
		return hx_value_349.(*haxe__ds___List__GoListIterator)
	}(self.iterator())
	for func(hx_value_351 any) bool {
		if hx_value_351 == nil {
			var hx_zero_352 bool
			return hx_zero_352
		}
		return hx_value_351.(bool)
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
