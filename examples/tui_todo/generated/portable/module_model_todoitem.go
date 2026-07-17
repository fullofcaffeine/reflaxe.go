package main

import "examples_tui_todo_portable/hxrt"

type I_model__TodoItem interface {
	set_title(value *string) *string
	set_done(value bool) bool
	set_priority(value int) int
}

type model__TodoItem struct {
	__hx_this I_model__TodoItem
	id        int
	title     *string
	done      bool
	priority  int
	tags      *hxrt.Array
}

func New_model__TodoItem(id int, title *string, priority int) *model__TodoItem {
	self := &model__TodoItem{}
	self.__hx_this = self
	self.id = id
	self.set_title(title)
	self.set_done(false)
	self.set_priority(priority)
	self.tags = hxrt.NewArray()
	return self
}

func (self *model__TodoItem) set_title(value *string) *string {
	self.title = value
	return value
}

func (self *model__TodoItem) set_done(value bool) bool {
	self.done = value
	return value
}

func (self *model__TodoItem) set_priority(value int) int {
	self.priority = value
	return value
}
