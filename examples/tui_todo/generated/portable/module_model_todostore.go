package main

type I_model__TodoStore interface {
	add(title *string, priority int) *model__TodoItem
	toggle(id int) bool
	addTag(id int, tag *string) bool
	list() []*model__TodoItem
	totalCount() int
	openCount() int
	doneCount() int
	findById(id int) *model__TodoItem
}

type model__TodoStore struct {
	__hx_this I_model__TodoStore
	nextId    int
	entries   []*model__TodoItem
}

func New_model__TodoStore() *model__TodoStore {
	self := &model__TodoStore{}
	self.__hx_this = self
	self.nextId = 1
	self.entries = []*model__TodoItem{}
	return self
}

func (self *model__TodoStore) add(title *string, priority int) *model__TodoItem {
	item := New_model__TodoItem(self.nextId, title, priority)
	self.nextId = int(int32((self.nextId + 1)))
	hx_arr_32 := self.entries
	hx_arr_32 = append(hx_arr_32, item)
	self.entries = hx_arr_32
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
	hx_arr_33 := item.tags
	hx_arr_33 = append(hx_arr_33, tag)
	item.tags = hx_arr_33
	return true
}

func (self *model__TodoStore) list() []*model__TodoItem {
	return self.entries
}

func (self *model__TodoStore) totalCount() int {
	return len(self.entries)
}

func (self *model__TodoStore) openCount() int {
	total := 0
	_g := 0
	_g1 := self.entries
	for _g < len(_g1) {
		item := _g1[_g]
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
	for _g < len(_g1) {
		item := _g1[_g]
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
	for _g < len(_g1) {
		item := _g1[_g]
		_g = int(int32((_g + 1)))
		if item.id == id {
			return item
		}
	}
	return nil
}
