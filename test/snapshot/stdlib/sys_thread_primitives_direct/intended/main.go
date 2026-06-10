package main

import "snapshot/hxrt"

func main() {
	lock := New_sys__thread__Lock()
	lock.release()
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("lock.release_before_wait="), hxrt.StdString(lock.wait(nil))))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("lock.timeout_empty="), hxrt.StdString(lock.wait(0.0))))
	mutex := New_sys__thread__Mutex()
	mutex.acquire()
	mutex.acquire()
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mutex.try_reentrant="), hxrt.StdString(mutex.tryAcquire())))
	mutex.release()
	mutex.release()
	mutex.release()
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mutex.try_after_release="), hxrt.StdString(mutex.tryAcquire())))
	mutex.release()
	condition := New_sys__thread__Condition()
	condition.acquire()
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("condition.try_reentrant="), hxrt.StdString(condition.tryAcquire())))
	condition.release()
	condition.release()
	sem := New_sys__thread__Semaphore(1)
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sem.try_first="), hxrt.StdString(sem.tryAcquire(nil))))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sem.try_empty="), hxrt.StdString(sem.tryAcquire(0.0))))
	sem.release()
	sem.acquire()
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sem.try_after_acquire="), hxrt.StdString(sem.tryAcquire(0.0))))
	loop := New_sys__thread__EventLoop()
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("loop.wait_empty="), hxrt.StdString(loop.wait(0.0))))
	deque := New_sys__thread__Deque()
	deque.add(hxrt.StringFromLiteral("tail"))
	deque.push(hxrt.StringFromLiteral("head"))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("deque.pop1="), func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(deque.pop(false))))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("deque.pop2="), func(hx_value_3 any) *string {
		if hx_value_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_value_3.(*string)
	}(deque.pop(false))))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("deque.pop3="), hxrt.StdString(func(hx_value_5 any) *string {
		if hx_value_5 == nil {
			var hx_zero_6 *string
			return hx_zero_6
		}
		return hx_value_5.(*string)
	}(deque.pop(false)))))
	tls := New_sys__thread__Tls()
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("tls.initial="), hxrt.StdString(func(hx_value_7 any) *string {
		if hx_value_7 == nil {
			var hx_zero_8 *string
			return hx_zero_8
		}
		return hx_value_7.(*string)
	}(tls.get_value()))))
	func(hx_value_9 any) *string {
		if hx_value_9 == nil {
			var hx_zero_10 *string
			return hx_zero_10
		}
		return hx_value_9.(*string)
	}(tls.set_value(hxrt.StringFromLiteral("worker")))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("tls.after_set="), hxrt.StdString(func(hx_value_11 any) *string {
		if hx_value_11 == nil {
			var hx_zero_12 *string
			return hx_zero_12
		}
		return hx_value_11.(*string)
	}(tls.get_value()))))
	func(hx_value_13 any) *string {
		if hx_value_13 == nil {
			var hx_zero_14 *string
			return hx_zero_14
		}
		return hx_value_13.(*string)
	}(tls.set_value(nil))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("tls.after_clear="), hxrt.StdString(func(hx_value_15 any) *string {
		if hx_value_15 == nil {
			var hx_zero_16 *string
			return hx_zero_16
		}
		return hx_value_15.(*string)
	}(tls.get_value()))))
	noLoop := New_sys__thread__NoEventLoopException(hxrt.StringFromLiteral("Event loop is not available. Refer to sys.thread.Thread.runWithEventLoop."), nil)
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("noLoop.msg="), hxrt.ExceptionMessage(noLoop)))
	pool := New__Main__DummyPool()
	pool.run(func() {
		hxrt.Println(hxrt.StringFromLiteral("pool.task=ran"))
	})
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("pool.runs="), pool.runCount()))
	pool.shutdown()
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("pool.shutdown="), hxrt.StdString(pool.get_isShutdown())))
	hxrt.TryCatch(func() {
		pool.run(func() {
		})
	}, func(hx_caught_17 any) {
		switch hx_typed_18 := hx_caught_17.(type) {
		case *sys__thread__ThreadPoolException:
			err := hx_typed_18
			hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("pool.error="), hxrt.ExceptionMessage(err)))
		default:
			hxrt.Throw(hx_caught_17)
		}
	})
}

type I__Main__DummyPool interface {
	get_threadsCount() int
	get_isShutdown() bool
	run(task func())
	shutdown()
	runCount() int
}

type _Main__DummyPool struct {
	__hx_this    I__Main__DummyPool
	threadsCount int
	isShutdown   bool
	_isShutdown  bool
	runs         int
}

func New__Main__DummyPool() *_Main__DummyPool {
	self := &_Main__DummyPool{}
	self.__hx_this = self
	self.runs = 0
	self._isShutdown = false
	return self
}

func (self *_Main__DummyPool) get_threadsCount() int {
	return 0
}

func (self *_Main__DummyPool) get_isShutdown() bool {
	return self._isShutdown
}

func (self *_Main__DummyPool) run(task func()) {
	if self._isShutdown {
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("shutdown"), nil, nil))
	}
	self.runs = int(int32((self.runs + 1)))
	task()
}

func (self *_Main__DummyPool) shutdown() {
	self._isShutdown = true
}

func (self *_Main__DummyPool) runCount() int {
	return self.runs
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

func (self *haxe__ds__IntMap) set(key any, value any) {
	resolvedKey := hxrt.IntFromNullableAny(key)
	self.h[resolvedKey] = value
}

func (self *haxe__ds__IntMap) get(key any) any {
	resolvedKey := hxrt.IntFromNullableAny(key)
	value := self.h[resolvedKey]
	return value
}

func (self *haxe__ds__IntMap) exists(key any) bool {
	resolvedKey := hxrt.IntFromNullableAny(key)
	_, ok := self.h[resolvedKey]
	return ok
}

func (self *haxe__ds__IntMap) remove(key any) bool {
	resolvedKey := hxrt.IntFromNullableAny(key)
	_, ok := self.h[resolvedKey]
	delete(self.h, resolvedKey)
	return ok
}

func (self *haxe__ds__IntMap) keys() map[string]any {
	keys := make([]int, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() int { key := keys[index]; index++; return key }
	return iter
}

func (self *haxe__ds__IntMap) iterator() map[string]any {
	keys := make([]int, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return self.h[key] }
	return iter
}

func (self *haxe__ds__IntMap) keyValueIterator() map[string]any {
	keys := make([]int, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() map[string]any {
		key := keys[index]
		index++
		return map[string]any{"key": key, "value": self.h[key]}
	}
	return iter
}

func (self *haxe__ds__IntMap) copyIMap() haxe__IMap {
	copied := New_haxe__ds__IntMap()
	for key, value := range self.h {
		copied.h[key] = value
	}
	return copied
}

func (self *haxe__ds__IntMap) toString() *string {
	return hxrt.StringFromLiteral("{}")
}

func (self *haxe__ds__IntMap) clear() {
	self.h = map[int]any{}
}

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	return &haxe__ds__StringMap{h: map[string]any{}}
}

func (self *haxe__ds__StringMap) set(key any, value any) {
	self.h[*hxrt.StdString(key)] = value
}

func (self *haxe__ds__StringMap) get(key any) any {
	value := self.h[*hxrt.StdString(key)]
	return value
}

func (self *haxe__ds__StringMap) exists(key any) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	return ok
}

func (self *haxe__ds__StringMap) remove(key any) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	delete(self.h, *hxrt.StdString(key))
	return ok
}

func (self *haxe__ds__StringMap) keys() map[string]any {
	keys := make([]string, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() *string { key := keys[index]; index++; return hxrt.StringFromLiteral(key) }
	return iter
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	keys := make([]string, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return self.h[key] }
	return iter
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	keys := make([]string, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() map[string]any {
		key := keys[index]
		index++
		return map[string]any{"key": hxrt.StringFromLiteral(key), "value": self.h[key]}
	}
	return iter
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	copied := New_haxe__ds__StringMap()
	for key, value := range self.h {
		copied.h[key] = value
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	return hxrt.StringFromLiteral("{}")
}

func (self *haxe__ds__StringMap) clear() {
	self.h = map[string]any{}
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

func (self *haxe__ds__ObjectMap) keys() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return key }
	return iter
}

func (self *haxe__ds__ObjectMap) iterator() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return self.h[key] }
	return iter
}

func (self *haxe__ds__ObjectMap) keyValueIterator() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() map[string]any {
		key := keys[index]
		index++
		return map[string]any{"key": key, "value": self.h[key]}
	}
	return iter
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	copied := New_haxe__ds__ObjectMap()
	for key, value := range self.h {
		copied.h[key] = value
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	return hxrt.StringFromLiteral("{}")
}

func (self *haxe__ds__ObjectMap) clear() {
	self.h = map[any]any{}
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

func (self *haxe__ds__EnumValueMap) keys() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return key }
	return iter
}

func (self *haxe__ds__EnumValueMap) iterator() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return self.h[key] }
	return iter
}

func (self *haxe__ds__EnumValueMap) keyValueIterator() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() map[string]any {
		key := keys[index]
		index++
		return map[string]any{"key": key, "value": self.h[key]}
	}
	return iter
}

func (self *haxe__ds__EnumValueMap) copyIMap() haxe__IMap {
	copied := New_haxe__ds__EnumValueMap()
	for key, value := range self.h {
		copied.h[key] = value
	}
	return copied
}

func (self *haxe__ds__EnumValueMap) toString() *string {
	return hxrt.StringFromLiteral("{}")
}

func (self *haxe__ds__EnumValueMap) clear() {
	self.h = map[any]any{}
}

func New_haxe__ds__List() *haxe__ds__List {
	return &haxe__ds__List{items: []any{}, length: 0}
}

func (self *haxe__ds__List) add(item any) {
	self.items = append(self.items, item)
	self.length = len(self.items)
}

func (self *haxe__ds__List) push(item any) {
	self.items = append([]any{item}, self.items...)
	self.length = len(self.items)
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
