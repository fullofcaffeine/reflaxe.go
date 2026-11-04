package main

import "examples_tui_todo_portable/hxrt"

type I_app__TodoApp interface {
	add(title *string, priority int) int
	addMany(titles *hxrt.Array, priority int) int
	toggle(id int) bool
	tag(id int, tag *string) bool
	baselineSignature() *string
	totalCount() int
	openCount() int
	doneCount() int
	diagnostics() *string
	buildRuntimeMetrics() *profile__TodoRuntimeMetrics
	render() *string
	items() *hxrt.Array
}

type app__TodoApp struct {
	__hx_this I_app__TodoApp
	runtime   profile__TodoRuntime
	store     *model__TodoStore
}

func (self *app__TodoApp) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__TodoApp(runtime profile__TodoRuntime) *app__TodoApp {
	return hx_init_app__TodoApp(hx_alloc_app__TodoApp(nil), runtime)
}

func hx_alloc_app__TodoApp(owner any) *app__TodoApp {
	self := &app__TodoApp{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__TodoApp)
	}
	return self
}

func hx_init_app__TodoApp(self *app__TodoApp, runtime profile__TodoRuntime) *app__TodoApp {
	self.runtime = runtime
	self.store = New_model__TodoStore()
	return self
}

func (self *app__TodoApp) add(title *string, priority int) int {
	item := self.store.add(self.runtime.normalizeTitle(title), priority)
	return item.id
}

func (self *app__TodoApp) addMany(titles *hxrt.Array, priority int) int {
	added := 0
	_g := 0
	for _g < titles.Len() {
		title := func(hx_value_1 any) *string {
			if hx_value_1 == nil {
				var hx_zero_2 *string
				return hx_zero_2
			}
			return hx_value_1.(*string)
		}(titles.Get(_g))
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
	return self.runtime.diagnostics(self.buildRuntimeMetrics())
}

func (self *app__TodoApp) buildRuntimeMetrics() *profile__TodoRuntimeMetrics {
	items := self.store.list()
	total := items.Len()
	done := 0
	p1 := 0
	_g := 0
	for _g < items.Len() {
		item := hx_recover_class_1(items.Get(_g))
		_g = int(int32((_g + 1)))
		if item.done {
			done = int(int32((done + 1)))
		}
		if item.priority == 1 {
			p1 = int(int32((p1 + 1)))
		}
	}
	return New_profile__TodoRuntimeMetrics(total, int((hxrt.Int32Wrap(total) - hxrt.Int32Wrap(done))), done, p1)
}

func (self *app__TodoApp) render() *string {
	out := hxrt.StringFromLiteral("== TODO ==")
	items := self.store.list()
	_g := 0
	for _g < items.Len() {
		item := hx_recover_class_1(items.Get(_g))
		_g = int(int32((_g + 1)))
		state := hxrt.StringFromLiteral("[ ]")
		if item.done {
			state = hxrt.StringFromLiteral("[x]")
		}
		tags := hxrt.StringFromLiteral("-")
		if item.tags.Len() != 0 {
			tags = app__TodoApp_joinStringList(item.tags, hxrt.StringFromLiteral(","))
		}
		out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("\n"), state), hxrt.StringFromLiteral(" #")), item.id), hxrt.StringFromLiteral(" p")), item.priority), hxrt.StringFromLiteral(" ")), item.title), hxrt.StringFromLiteral(" tags:")), tags))
	}
	out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("\nsummary "), self.baselineSignature()))
	return out
}

func (self *app__TodoApp) items() *hxrt.Array {
	return self.store.list()
}

func app__TodoApp_joinStringList(values *hxrt.Array, separator *string) *string {
	out := hxrt.StringFromLiteral("")
	first := true
	_g := 0
	for _g < values.Len() {
		value := func(hx_value_3 any) *string {
			if hx_value_3 == nil {
				var hx_zero_4 *string
				return hx_zero_4
			}
			return hx_value_3.(*string)
		}(values.Get(_g))
		_g = int(int32((_g + 1)))
		if !first {
			out = hxrt.StringConcatStringPtr(out, separator)
		}
		out = hxrt.StringConcatStringPtr(out, value)
		first = false
	}
	return out
}
