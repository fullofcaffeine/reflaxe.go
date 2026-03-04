package main

import "examples_tui_todo_portable/hxrt"

type I_app__TodoApp interface {
	add(title *string, priority int) int
	addMany(titles []*string, priority int) int
	toggle(id int) bool
	tag(id int, tag *string) bool
	baselineSignature() *string
	totalCount() int
	openCount() int
	doneCount() int
	diagnostics() *string
	buildRuntimeMetrics() *profile__TodoRuntimeMetrics
	render() *string
	items() []*model__TodoItem
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

func (self *app__TodoApp) addMany(titles []*string, priority int) int {
	if !self.runtime.supportsBatchAdd() {
		return 0
	}
	added := 0
	_g := 0
	for _g < len(titles) {
		title := titles[_g]
		_g = int(int32((_g + 1)))
		self.add(title, priority)
		added = int(int32((added + 1)))
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
	return self.runtime.diagnostics(self.buildRuntimeMetrics())
}

func (self *app__TodoApp) buildRuntimeMetrics() *profile__TodoRuntimeMetrics {
	items := self.store.list()
	total := len(items)
	done := 0
	p1 := 0
	_g := 0
	for _g < len(items) {
		item := items[_g]
		_g = int(int32((_g + 1)))
		if item.done {
			done = int(int32((done + 1)))
		}
		if item.priority == 1 {
			p1 = int(int32((p1 + 1)))
		}
	}
	return New_profile__TodoRuntimeMetrics(total, int(int32((hxrt.Int32Wrap(total) - hxrt.Int32Wrap(done)))), done, p1)
}

func (self *app__TodoApp) render() *string {
	out := hxrt.StringFromLiteral("== TODO ==")
	items := self.store.list()
	_g := 0
	for _g < len(items) {
		item := items[_g]
		_g = int(int32((_g + 1)))
		state := hxrt.StringFromLiteral("[ ]")
		if item.done {
			state = hxrt.StringFromLiteral("[x]")
		}
		tags := hxrt.StringFromLiteral("-")
		if len(item.tags) != 0 {
			tags = app__TodoApp_joinStringList(item.tags, hxrt.StringFromLiteral(","))
		}
		out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("\n"), state), hxrt.StringFromLiteral(" #")), item.id), hxrt.StringFromLiteral(" p")), item.priority), hxrt.StringFromLiteral(" ")), item.title), hxrt.StringFromLiteral(" tags:")), tags))
	}
	out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("\nsummary "), self.baselineSignature()))
	return out
}

func (self *app__TodoApp) items() []*model__TodoItem {
	return self.store.list()
}

func app__TodoApp_joinStringList(values []*string, separator *string) *string {
	out := hxrt.StringFromLiteral("")
	first := true
	_g := 0
	for _g < len(values) {
		value := values[_g]
		_g = int(int32((_g + 1)))
		if !first {
			out = hxrt.StringConcatStringPtr(out, separator)
		}
		out = hxrt.StringConcatStringPtr(out, value)
		first = false
	}
	return out
}
