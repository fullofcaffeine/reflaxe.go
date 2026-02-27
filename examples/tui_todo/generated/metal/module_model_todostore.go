package main

type I_model__TodoStore interface {
	add(title *string, priority int) *model__TodoItem
	toggle(id int) bool
	addTag(id int, tag *string) bool
	list() *haxe__ds__List
	totalCount() int
	openCount() int
	doneCount() int
	findById(id int) *model__TodoItem
}

type model__TodoStore struct {
	__hx_this I_model__TodoStore
	nextId    int
	entries   *haxe__ds__List
}

func New_model__TodoStore() *model__TodoStore {
	self := &model__TodoStore{}
	self.__hx_this = self
	self.nextId = 1
	self.entries = New_haxe__ds__List()
	return self
}

func (self *model__TodoStore) add(title *string, priority int) *model__TodoItem {
	item := New_model__TodoItem(self.nextId, title, priority)
	self.nextId = int(int32((self.nextId + 1)))
	self.entries.add(item)
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
	item.tags.add(tag)
	return true
}

func (self *model__TodoStore) list() *haxe__ds__List {
	return self.entries
}

func (self *model__TodoStore) totalCount() int {
	return self.entries.length
}

func (self *model__TodoStore) openCount() int {
	total := 0
	count := self.entries.length
	i := 0
	for i < count {
		value := func(hx_value_31 any) *model__TodoItem {
			if hx_value_31 == nil {
				var hx_zero_32 *model__TodoItem
				return hx_zero_32
			}
			return hx_value_31.(*model__TodoItem)
		}(self.entries.pop())
		if value == nil {
			break
		}
		item := value
		if !item.done {
			total = int(int32((total + 1)))
		}
		self.entries.add(item)
		i = int(int32((i + 1)))
	}
	return total
}

func (self *model__TodoStore) doneCount() int {
	total := 0
	count := self.entries.length
	i := 0
	for i < count {
		value := func(hx_value_33 any) *model__TodoItem {
			if hx_value_33 == nil {
				var hx_zero_34 *model__TodoItem
				return hx_zero_34
			}
			return hx_value_33.(*model__TodoItem)
		}(self.entries.pop())
		if value == nil {
			break
		}
		item := value
		if item.done {
			total = int(int32((total + 1)))
		}
		self.entries.add(item)
		i = int(int32((i + 1)))
	}
	return total
}

func (self *model__TodoStore) findById(id int) *model__TodoItem {
	var found *model__TodoItem = nil
	count := self.entries.length
	i := 0
	for i < count {
		value := func(hx_value_35 any) *model__TodoItem {
			if hx_value_35 == nil {
				var hx_zero_36 *model__TodoItem
				return hx_zero_36
			}
			return hx_value_35.(*model__TodoItem)
		}(self.entries.pop())
		if value == nil {
			break
		}
		item := value
		if item.id == id {
			found = item
		}
		self.entries.add(item)
		i = int(int32((i + 1)))
	}
	return found
}
