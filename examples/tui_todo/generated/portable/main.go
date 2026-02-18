package main

import "examples_tui_todo_portable/hxrt"

func Harness_assertContract(runtime profile__TodoRuntime) *string {
	app := New_app__TodoApp(runtime)
	_ = app
	Harness_runBaseline(app)
	baseline := app.__hx_this.baselineSignature()
	_ = baseline
	if !hxrt.StringEqualAny(baseline, hxrt.StringFromLiteral("open=1,done=1,total=2")) {
		hxrt.Throw(hxrt.StringConcatAny(hxrt.StringFromLiteral("baseline drift: "), baseline))
	}
	if runtime.supportsBatchAdd() {
		added := app.__hx_this.addMany(Harness_batchTitles(), 3)
		_ = added
		if (added != 2) || (app.__hx_this.totalCount() != 4) {
			hxrt.Throw(hxrt.StringFromLiteral("batch add drift"))
		}
	} else {
		if app.__hx_this.totalCount() != 2 {
			hxrt.Throw(hxrt.StringFromLiteral("portable total drift"))
		}
	}
	if runtime.supportsDiagnostics() {
		diag := app.__hx_this.diagnostics()
		_ = diag
		if !hxrt.StringEqualAny(diag, hxrt.StringFromLiteral("p1=1,completed=1")) {
			hxrt.Throw(hxrt.StringFromLiteral("missing diagnostics"))
		}
	}
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("OK "), runtime.profileId())
}

func Harness_batchTitles() *haxe__ds__List {
	out := New_haxe__ds__List()
	_ = out
	out.add(hxrt.StringFromLiteral("Ship generated-go sync"))
	out.add(hxrt.StringFromLiteral("Add binary matrix"))
	return out
}

func Harness_run(runtime profile__TodoRuntime) *string {
	app := New_app__TodoApp(runtime)
	_ = app
	baselineView := Harness_runBaseline(app)
	_ = baselineView
	baseline := app.__hx_this.baselineSignature()
	_ = baseline
	extras := hxrt.StringFromLiteral("batch_add=0")
	_ = extras
	if runtime.supportsBatchAdd() {
		added := app.__hx_this.addMany(Harness_batchTitles(), 3)
		_ = added
		extras = hxrt.StringConcatAny(hxrt.StringFromLiteral("batch_add="), added)
	}
	extras = hxrt.StringConcatAny(extras, hxrt.StringConcatAny(hxrt.StringFromLiteral(",diag="), app.__hx_this.diagnostics()))
	return hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("profile="), runtime.profileId()), hxrt.StringFromLiteral("\nbaseline=")), baseline), hxrt.StringFromLiteral("\nbaseline_view:\n")), baselineView), hxrt.StringFromLiteral("\nfinal_view:\n")), app.__hx_this.render()), hxrt.StringFromLiteral("\nextras=")), extras)
}

func Harness_runBaseline(app *app__TodoApp) *string {
	app.__hx_this.add(hxrt.StringFromLiteral("Write profile docs"), 2)
	app.__hx_this.add(hxrt.StringFromLiteral("Backfill regression snapshots"), 1)
	app.__hx_this.toggle(2)
	app.__hx_this.tag(1, hxrt.StringFromLiteral("docs"))
	app.__hx_this.tag(2, hxrt.StringFromLiteral("tests"))
	return app.__hx_this.render()
}

func main() {
	var runtime profile__TodoRuntime = profile__RuntimeFactory_create()
	_ = runtime
	hxrt.Println(Harness_run(runtime))
}

type I_app__TodoApp interface {
	add(title *string, priority int)
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

func (self *app__TodoApp) add(title *string, priority int) {
	self.store.__hx_this.add(self.runtime.normalizeTitle(title), priority)
}

func (self *app__TodoApp) addMany(titles *haxe__ds__List, priority int) int {
	if !self.runtime.supportsBatchAdd() {
		return 0
	}
	added := 0
	_ = added
	count := titles.length
	_ = count
	i := 0
	_ = i
	for i < count {
		raw := titles.pop().(*string)
		_ = raw
		if hxrt.StringEqualAny(raw, nil) {
			break
		}
		title := raw
		_ = title
		self.__hx_this.add(title, priority)
		titles.add(title)
		added = (added + 1)
		i = (i + 1)
	}
	return added
}

func (self *app__TodoApp) toggle(id int) bool {
	return self.store.__hx_this.toggle(id)
}

func (self *app__TodoApp) tag(id int, tag *string) bool {
	return self.store.__hx_this.addTag(id, self.runtime.normalizeTag(tag))
}

func (self *app__TodoApp) baselineSignature() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("open="), self.__hx_this.openCount()), hxrt.StringFromLiteral(",done=")), self.__hx_this.doneCount()), hxrt.StringFromLiteral(",total=")), self.__hx_this.totalCount())
}

func (self *app__TodoApp) totalCount() int {
	return self.store.__hx_this.totalCount()
}

func (self *app__TodoApp) openCount() int {
	return self.store.__hx_this.openCount()
}

func (self *app__TodoApp) doneCount() int {
	return self.store.__hx_this.doneCount()
}

func (self *app__TodoApp) diagnostics() *string {
	if !self.runtime.supportsDiagnostics() {
		return hxrt.StringFromLiteral("off")
	}
	return self.runtime.diagnostics(self.store.__hx_this.list())
}

func (self *app__TodoApp) render() *string {
	out := hxrt.StringFromLiteral("== TODO ==")
	_ = out
	items := self.store.__hx_this.list()
	_ = items
	count := items.length
	_ = count
	i := 0
	_ = i
	for i < count {
		raw := items.pop().(*model__TodoItem)
		_ = raw
		if hxrt.StringEqualAny(raw, nil) {
			break
		}
		item := raw
		_ = item
		state := hxrt.StringFromLiteral("[ ]")
		_ = state
		if item.done {
			state = hxrt.StringFromLiteral("[x]")
		}
		tags := hxrt.StringFromLiteral("-")
		_ = tags
		if item.tags.length != 0 {
			tags = app__TodoApp_joinStringList(item.tags, hxrt.StringFromLiteral(","))
		}
		out = hxrt.StringConcatAny(out, hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("\n"), state), hxrt.StringFromLiteral(" #")), item.id), hxrt.StringFromLiteral(" p")), item.priority), hxrt.StringFromLiteral(" ")), item.title), hxrt.StringFromLiteral(" tags:")), tags))
		items.add(item)
		i = (i + 1)
	}
	out = hxrt.StringConcatAny(out, hxrt.StringConcatAny(hxrt.StringFromLiteral("\nsummary "), self.__hx_this.baselineSignature()))
	return out
}

func (self *app__TodoApp) items() *haxe__ds__List {
	return self.store.__hx_this.list()
}

func app__TodoApp_joinStringList(values *haxe__ds__List, separator *string) *string {
	out := hxrt.StringFromLiteral("")
	_ = out
	first := true
	_ = first
	count := values.length
	_ = count
	i := 0
	_ = i
	for i < count {
		raw := values.pop().(*string)
		_ = raw
		if hxrt.StringEqualAny(raw, nil) {
			break
		}
		value := raw
		_ = value
		if !first {
			out = hxrt.StringConcatAny(out, separator)
		}
		out = hxrt.StringConcatAny(out, value)
		values.add(value)
		first = false
		i = (i + 1)
	}
	return out
}

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
	tags      *haxe__ds__List
}

func New_model__TodoItem(id int, title *string, priority int) *model__TodoItem {
	self := &model__TodoItem{}
	self.__hx_this = self
	self.id = id
	self.__hx_this.set_title(title)
	self.__hx_this.set_done(false)
	self.__hx_this.set_priority(priority)
	self.tags = New_haxe__ds__List()
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
	_ = item
	self.nextId = (self.nextId + 1)
	self.entries.add(item)
	return item
}

func (self *model__TodoStore) toggle(id int) bool {
	item := self.__hx_this.findById(id)
	_ = item
	if hxrt.StringEqualAny(item, nil) {
		return false
	}
	item.__hx_this.set_done(!item.done)
	return true
}

func (self *model__TodoStore) addTag(id int, tag *string) bool {
	item := self.__hx_this.findById(id)
	_ = item
	if hxrt.StringEqualAny(item, nil) {
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
	_ = total
	count := self.entries.length
	_ = count
	i := 0
	_ = i
	for i < count {
		value := self.entries.pop().(*model__TodoItem)
		_ = value
		if hxrt.StringEqualAny(value, nil) {
			break
		}
		item := value
		_ = item
		if !item.done {
			total = (total + 1)
		}
		self.entries.add(item)
		i = (i + 1)
	}
	return total
}

func (self *model__TodoStore) doneCount() int {
	total := 0
	_ = total
	count := self.entries.length
	_ = count
	i := 0
	_ = i
	for i < count {
		value := self.entries.pop().(*model__TodoItem)
		_ = value
		if hxrt.StringEqualAny(value, nil) {
			break
		}
		item := value
		_ = item
		if item.done {
			total = (total + 1)
		}
		self.entries.add(item)
		i = (i + 1)
	}
	return total
}

func (self *model__TodoStore) findById(id int) *model__TodoItem {
	var found *model__TodoItem = nil
	_ = found
	count := self.entries.length
	_ = count
	i := 0
	_ = i
	for i < count {
		value := self.entries.pop().(*model__TodoItem)
		_ = value
		if hxrt.StringEqualAny(value, nil) {
			break
		}
		item := value
		_ = item
		if item.id == id {
			found = item
		}
		self.entries.add(item)
		i = (i + 1)
	}
	return found
}

type I_profile__PortableRuntime interface {
	profileId() *string
	normalizeTitle(title *string) *string
	normalizeTag(tag *string) *string
	supportsBatchAdd() bool
	supportsDiagnostics() bool
	diagnostics(items *haxe__ds__List) *string
}

type profile__PortableRuntime struct {
	__hx_this I_profile__PortableRuntime
}

func New_profile__PortableRuntime() *profile__PortableRuntime {
	self := &profile__PortableRuntime{}
	self.__hx_this = self
	return self
}

func (self *profile__PortableRuntime) profileId() *string {
	return hxrt.StringFromLiteral("portable")
}

func (self *profile__PortableRuntime) normalizeTitle(title *string) *string {
	return title
}

func (self *profile__PortableRuntime) normalizeTag(tag *string) *string {
	return tag
}

func (self *profile__PortableRuntime) supportsBatchAdd() bool {
	return false
}

func (self *profile__PortableRuntime) supportsDiagnostics() bool {
	return false
}

func (self *profile__PortableRuntime) diagnostics(items *haxe__ds__List) *string {
	return hxrt.StringFromLiteral("off")
}

func profile__RuntimeFactory_create() profile__TodoRuntime {
	return New_profile__PortableRuntime()
}

type profile__TodoRuntime interface {
	profileId() *string
	normalizeTitle(title *string) *string
	normalizeTag(tag *string) *string
	supportsBatchAdd() bool
	supportsDiagnostics() bool
	diagnostics(items *haxe__ds__List) *string
}

type haxe__ds__IntMap struct {
	h map[int]any
}

type haxe__ds__StringMap struct {
	h map[string]any
}

type haxe__ds__ObjectMap struct {
	h map[any]any
}

type haxe__ds__EnumValueMap struct {
	h map[any]any
}

type haxe__ds__List struct {
	items  []any
	length int
}

func New_haxe__ds__IntMap() *haxe__ds__IntMap {
	return &haxe__ds__IntMap{h: map[int]any{}}
}

func (self *haxe__ds__IntMap) set(key int, value any) {
	self.h[key] = value
}

func (self *haxe__ds__IntMap) get(key int) any {
	value := self.h[key]
	return value
}

func (self *haxe__ds__IntMap) exists(key int) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__IntMap) remove(key int) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	return &haxe__ds__StringMap{h: map[string]any{}}
}

func (self *haxe__ds__StringMap) set(key *string, value any) {
	self.h[*hxrt.StdString(key)] = value
}

func (self *haxe__ds__StringMap) get(key *string) any {
	value := self.h[*hxrt.StdString(key)]
	return value
}

func (self *haxe__ds__StringMap) exists(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	return ok
}

func (self *haxe__ds__StringMap) remove(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	delete(self.h, *hxrt.StdString(key))
	return ok
}

func New_haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	return &haxe__ds__ObjectMap{h: map[any]any{}}
}

func (self *haxe__ds__ObjectMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__ObjectMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__ObjectMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__ObjectMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__EnumValueMap() *haxe__ds__EnumValueMap {
	return &haxe__ds__EnumValueMap{h: map[any]any{}}
}

func (self *haxe__ds__EnumValueMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__EnumValueMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__EnumValueMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__EnumValueMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__List() *haxe__ds__List {
	return &haxe__ds__List{items: []any{}, length: 0}
}

func (self *haxe__ds__List) add(item any) {
	self.items = append(self.items, item)
	self.length = len(self.items)
}

func (self *haxe__ds__List) push(item any) {
	self.add(item)
}

func (self *haxe__ds__List) pop() any {
	if len(self.items) == 0 {
		return nil
	}
	head := self.items[0]
	self.items = self.items[1:]
	self.length = len(self.items)
	return head
}

func (self *haxe__ds__List) first() any {
	if len(self.items) == 0 {
		return nil
	}
	return self.items[0]
}

func (self *haxe__ds__List) last() any {
	size := len(self.items)
	if size == 0 {
		return nil
	}
	return self.items[(size - 1)]
}
