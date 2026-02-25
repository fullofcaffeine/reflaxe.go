package main

import "examples_tui_todo_portable/hxrt"

type I_app__TodoApp interface {
	add(title *string, priority int) int
	addMany(titles *haxe__ds__List, priority int) int
	toggle(id int) bool
	tag(id int, tag *string) bool
	baselineSignature() *string
	totalCount() int
	openCount() int
	doneCount() int
	diagnostics() *string
	render() *string
	items() *haxe__ds__List
}

type app__TodoApp struct {
	__hx_this I_app__TodoApp
	runtime   profile__TodoRuntime
	store     *model__TodoStore
}

func New_app__TodoApp(runtime profile__TodoRuntime) *app__TodoApp {
	self := &app__TodoApp{}
	self.__hx_this = self
	self.runtime = runtime
	self.store = New_model__TodoStore()
	return self
}

func (self *app__TodoApp) add(title *string, priority int) int {
	item := self.store.add(self.runtime.normalizeTitle(title), priority)
	return item.id
}

func (self *app__TodoApp) addMany(titles *haxe__ds__List, priority int) int {
	if !self.runtime.supportsBatchAdd() {
		return 0
	}
	added := 0
	count := titles.length
	i := 0
	for i < count {
		raw := func(hx_value_23 any) *string {
			if hx_value_23 == nil {
				var hx_zero_24 *string
				return hx_zero_24
			}
			return hx_value_23.(*string)
		}(titles.pop())
		if hxrt.StringEqualStringPtr(raw, nil) {
			break
		}
		title := raw
		self.add(title, priority)
		titles.add(title)
		added = int(int32((added + 1)))
		i = int(int32((i + 1)))
	}
	return added
}

func (self *app__TodoApp) toggle(id int) bool {
	return self.store.toggle(id)
}

func (self *app__TodoApp) tag(id int, tag *string) bool {
	return self.store.addTag(id, self.runtime.normalizeTag(tag))
}

func (self *app__TodoApp) baselineSignature() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("open="), self.openCount()), hxrt.StringFromLiteral(",done=")), self.doneCount()), hxrt.StringFromLiteral(",total=")), self.totalCount())
}

func (self *app__TodoApp) totalCount() int {
	return self.store.totalCount()
}

func (self *app__TodoApp) openCount() int {
	return self.store.openCount()
}

func (self *app__TodoApp) doneCount() int {
	return self.store.doneCount()
}

func (self *app__TodoApp) diagnostics() *string {
	if !self.runtime.supportsDiagnostics() {
		return hxrt.StringFromLiteral("off")
	}
	return self.runtime.diagnostics(self.store.list())
}

func (self *app__TodoApp) render() *string {
	out := hxrt.StringFromLiteral("== TODO ==")
	items := self.store.list()
	count := items.length
	i := 0
	for i < count {
		raw := func(hx_value_25 any) *model__TodoItem {
			if hx_value_25 == nil {
				var hx_zero_26 *model__TodoItem
				return hx_zero_26
			}
			return hx_value_25.(*model__TodoItem)
		}(items.pop())
		if raw == nil {
			break
		}
		item := raw
		state := hxrt.StringFromLiteral("[ ]")
		if item.done {
			state = hxrt.StringFromLiteral("[x]")
		}
		tags := hxrt.StringFromLiteral("-")
		if item.tags.length != 0 {
			tags = app__TodoApp_joinStringList(item.tags, hxrt.StringFromLiteral(","))
		}
		out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("\n"), state), hxrt.StringFromLiteral(" #")), item.id), hxrt.StringFromLiteral(" p")), item.priority), hxrt.StringFromLiteral(" ")), item.title), hxrt.StringFromLiteral(" tags:")), tags))
		items.add(item)
		i = int(int32((i + 1)))
	}
	out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("\nsummary "), self.baselineSignature()))
	return out
}

func (self *app__TodoApp) items() *haxe__ds__List {
	return self.store.list()
}

func app__TodoApp_joinStringList(values *haxe__ds__List, separator *string) *string {
	out := hxrt.StringFromLiteral("")
	first := true
	count := values.length
	i := 0
	for i < count {
		raw := func(hx_value_27 any) *string {
			if hx_value_27 == nil {
				var hx_zero_28 *string
				return hx_zero_28
			}
			return hx_value_27.(*string)
		}(values.pop())
		if hxrt.StringEqualStringPtr(raw, nil) {
			break
		}
		value := raw
		if !first {
			out = hxrt.StringConcatStringPtr(out, separator)
		}
		out = hxrt.StringConcatStringPtr(out, value)
		values.add(value)
		first = false
		i = int(int32((i + 1)))
	}
	return out
}
