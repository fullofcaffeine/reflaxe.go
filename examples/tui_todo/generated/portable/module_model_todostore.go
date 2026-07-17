package main

import "examples_tui_todo_portable/hxrt"

type I_model__TodoStore interface {
	add(title *string, priority int) *model__TodoItem
	toggle(id int) bool
	addTag(id int, tag *string) bool
	list() *hxrt.Array
	totalCount() int
	openCount() int
	doneCount() int
	findById(id int) *model__TodoItem
}

type model__TodoStore struct {
	__hx_this I_model__TodoStore
	nextId    int
	entries   *hxrt.Array
}

func New_model__TodoStore() *model__TodoStore {
	self := &model__TodoStore{}
	self.__hx_this = self
	self.nextId = 1
	self.entries = hxrt.NewArray()
	return self
}

func (self *model__TodoStore) add(title *string, priority int) *model__TodoItem {
	item := New_model__TodoItem(self.nextId, title, priority)
	self.nextId = int(int32((self.nextId + 1)))
	hx_arr_86 := self.entries
	hx_arr_86.Push(item)
	return item
}

func (self *model__TodoStore) toggle(id int) bool {
	item := self.findById(id)
	if item == nil {
		return false
	}
	item.set_done(!item.done)
	return true
}

func (self *model__TodoStore) addTag(id int, tag *string) bool {
	item := self.findById(id)
	if item == nil {
		return false
	}
	hx_arr_87 := item.tags
	hx_arr_87.Push(tag)
	return true
}

func (self *model__TodoStore) list() *hxrt.Array {
	return self.entries
}

func (self *model__TodoStore) totalCount() int {
	return self.entries.Len()
}

func (self *model__TodoStore) openCount() int {
	total := 0
	_g := 0
	_g1 := self.entries
	for _g < _g1.Len() {
		item := func(hx_value_88 any) *model__TodoItem {
			if hx_value_88 == nil {
				var hx_zero_89 *model__TodoItem
				return hx_zero_89
			}
			return hx_value_88.(*model__TodoItem)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if !item.done {
			total = int(int32((total + 1)))
		}
	}
	return total
}

func (self *model__TodoStore) doneCount() int {
	total := 0
	_g := 0
	_g1 := self.entries
	for _g < _g1.Len() {
		item := func(hx_value_90 any) *model__TodoItem {
			if hx_value_90 == nil {
				var hx_zero_91 *model__TodoItem
				return hx_zero_91
			}
			return hx_value_90.(*model__TodoItem)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if item.done {
			total = int(int32((total + 1)))
		}
	}
	return total
}

func (self *model__TodoStore) findById(id int) *model__TodoItem {
	_g := 0
	_g1 := self.entries
	for _g < _g1.Len() {
		item := func(hx_value_92 any) *model__TodoItem {
			if hx_value_92 == nil {
				var hx_zero_93 *model__TodoItem
				return hx_zero_93
			}
			return hx_value_92.(*model__TodoItem)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if item.id == id {
			return item
		}
	}
	return nil
}
