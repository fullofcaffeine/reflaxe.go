package main

import "examples_worker_pool_select_portable/hxrt"

var EMPTY_TOKEN *string = hxrt.StringFromLiteral("__empty__")

var STOP_TOKEN *string = hxrt.StringFromLiteral("__stop__")

func main() {
	workerCount := 3
	_ = workerCount
	tasks := []*string{hxrt.StringFromLiteral("alpha"), hxrt.StringFromLiteral("beta"), hxrt.StringFromLiteral("gamma"), hxrt.StringFromLiteral("delta")}
	_ = tasks
	jobs := go___Go_newChan(int(int32((hxrt.Int32Wrap(len(tasks)) + hxrt.Int32Wrap(workerCount)))))
	_ = jobs
	results := go___Go_newChan(len(tasks))
	_ = results
	_g := 0
	_ = _g
	for _g < len(tasks) {
		task := tasks[_g]
		_ = task
		_g = int(int32((_g + 1)))
		jobs.__hx_this.send(task)
	}
	_g_1 := 0
	_ = _g_1
	_g1 := workerCount
	_ = _g1
	for _g_1 < _g1 {
		hx_post_1 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		hx_tmp := hx_post_1
		_ = hx_tmp
		jobs.__hx_this.send(hxrt.StringFromLiteral("__stop__"))
	}
	_g_2 := 0
	_ = _g_2
	_g1_1 := workerCount
	_ = _g1_1
	for _g_2 < _g1_1 {
		hx_post_2 := _g_2
		_g_2 = int(int32((_g_2 + 1)))
		hx_tmp_1 := hx_post_2
		_ = hx_tmp_1
		go___Go_spawn(func() {
			worker(jobs, results)
		})
	}
	received := 0
	_ = received
	for received < len(tasks) {
		value := func(hx_value_3 any) *string {
			if hx_value_3 == nil {
				var hx_zero_4 *string
				return hx_zero_4
			}
			return hx_value_3.(*string)
		}(results.__hx_this.recvOr(hxrt.StringFromLiteral("__empty__")))
		_ = value
		if hxrt.StringEqualAny(value, hxrt.StringFromLiteral("__empty__")) {
			continue
		}
		received = int(int32((received + 1)))
	}
	selectGate := go___Go_newChan(1)
	_ = selectGate
	firstTry := selectGate.__hx_this.trySend(5)
	_ = firstTry
	secondTry := selectGate.__hx_this.trySend(6)
	_ = secondTry
	firstRecv := func(hx_value_5 any) int {
		if hx_value_5 == nil {
			var hx_zero_6 int
			return hx_zero_6
		}
		return hx_value_5.(int)
	}(selectGate.__hx_this.recvOr(-1))
	_ = firstRecv
	secondRecv := func(hx_value_7 any) int {
		if hx_value_7 == nil {
			var hx_zero_8 int
			return hx_zero_8
		}
		return hx_value_7.(int)
	}(selectGate.__hx_this.recvOr(99))
	_ = secondRecv
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("worker.count="), received))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("select.trySend="), hxrt.StdString(firstTry)), hxrt.StringFromLiteral(",")), hxrt.StdString(secondTry)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("select.recvOr="), firstRecv), hxrt.StringFromLiteral(",")), secondRecv))
}

func worker(jobs *go___Chan, results *go___Chan) {
	for true {
		job := func(hx_value_9 any) *string {
			if hx_value_9 == nil {
				var hx_zero_10 *string
				return hx_zero_10
			}
			return hx_value_9.(*string)
		}(jobs.__hx_this.recvOr(hxrt.StringFromLiteral("__stop__")))
		_ = job
		if hxrt.StringEqualAny(job, hxrt.StringFromLiteral("__stop__")) {
			return
		}
		results.__hx_this.send(job)
	}
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

func go__concurrency_makeChan(buffer int) any {
	if buffer > 0 {
		return make(chan any, buffer)
	}
	return make(chan any)
}

func go__concurrency_send(channel any, value any) {
	channel.(chan any) <- value
}

func go__concurrency_trySend(channel any, value any) bool {
	select {
	case channel.(chan any) <- value:
		return true
	default:
		return false
	}
}

func go__concurrency_recv(channel any) any {
	return <-channel.(chan any)
}

func go__concurrency_recvOr(channel any, defaultValue any) any {
	select {
	case value := <-channel.(chan any):
		return value
	default:
		return defaultValue
	}
}

func go__concurrency_tryRecv(channel any) *go___Result {
	select {
	case value := <-channel.(chan any):
		return New_go___Result(value, nil)
	default:
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("empty")))
	}
}

func go__concurrency_close(channel any) {
	close(channel.(chan any))
}

func go__concurrency_spawn(fn func()) {
	go fn()
}
