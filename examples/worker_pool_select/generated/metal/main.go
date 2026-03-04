package main

import (
	"errors"
	"examples_worker_pool_select_metal/hxrt"
	"reflect"
)

var EMPTY_TOKEN *string = hxrt.StringFromLiteral("__empty__")

var STOP_TOKEN *string = hxrt.StringFromLiteral("__stop__")

func main() {
	workerCount := 3
	tasks := []*string{hxrt.StringFromLiteral("alpha"), hxrt.StringFromLiteral("beta"), hxrt.StringFromLiteral("gamma"), hxrt.StringFromLiteral("delta")}
	jobs := go__concurrency_newChan___string_f613ccd0(int(int32((hxrt.Int32Wrap(len(tasks)) + hxrt.Int32Wrap(workerCount)))))
	results := go__concurrency_newChan___string_f613ccd0(len(tasks))
	_g := 0
	for _g < len(tasks) {
		task := tasks[_g]
		_g = int(int32((_g + 1)))
		go__concurrency_send___string_f613ccd0(jobs.__hx_native, task)
	}
	_g_1 := 0
	_g1 := workerCount
	for _g_1 < _g1 {
		hx_post_1 := _g_1
		_g_1 = int(int32((_g_1 + 1)))
		hx_tmp := hx_post_1
		_ = hx_tmp
		go__concurrency_send___string_f613ccd0(jobs.__hx_native, hxrt.StringFromLiteral("__stop__"))
	}
	_g_2 := 0
	_g1_1 := workerCount
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
	for received < len(tasks) {
		value := go__concurrency_recvOr___string_f613ccd0(results.__hx_native, hxrt.StringFromLiteral("__empty__"))
		if hxrt.StringEqualStringPtr(value, hxrt.StringFromLiteral("__empty__")) {
			continue
		}
		received = int(int32((received + 1)))
	}
	selectGate := go__concurrency_newChan__int_95e97e5e(1)
	_g_3 := go___Select_send_Int(selectGate, 5)
	var hx_switch_3 bool
	switch _g_3.tag {
	case 0:
		hx_switch_3 = true
	case 1:
		hx_switch_3 = false
	}
	firstTry := hx_switch_3
	_g_4 := go___Select_send_Int(selectGate, 6)
	var hx_switch_4 bool
	switch _g_4.tag {
	case 0:
		hx_switch_4 = true
	case 1:
		hx_switch_4 = false
	}
	secondTry := hx_switch_4
	_g_5 := go___Select_recv_Int(selectGate)
	var hx_switch_5 int
	switch _g_5.tag {
	case 0:
		_g_6 := _g_5.params[0].(int)
		value_1 := _g_6
		hx_switch_5 = value_1
	case 1:
		hx_switch_5 = -1
	}
	firstRecv := hx_switch_5
	_g_7 := go___Select_recv_Int(selectGate)
	var hx_switch_6 int
	switch _g_7.tag {
	case 0:
		_g_8 := _g_7.params[0].(int)
		value_2 := _g_8
		hx_switch_6 = value_2
	case 1:
		hx_switch_6 = 99
	}
	secondRecv := hx_switch_6
	left := go__concurrency_newChan___string_f613ccd0(1)
	right := go__concurrency_newChan___string_f613ccd0(1)
	go__concurrency_send___string_f613ccd0(right.__hx_native, hxrt.StringFromLiteral("right"))
	_g_9 := go___Select_recv2_String_String(left, right)
	var hx_switch_7 *string
	switch _g_9.tag {
	case 0:
		_g_10 := _g_9.params[0].(*string)
		value_3 := _g_10
		hx_switch_7 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("left:"), value_3)
	case 1:
		_g_11 := _g_9.params[0].(*string)
		value_4 := _g_11
		hx_switch_7 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("right:"), value_4)
	case 2:
		hx_switch_7 = hxrt.StringFromLiteral("none")
	}
	recv2 := hx_switch_7
	send2a := go__concurrency_newChan__int_95e97e5e(1)
	send2b := go__concurrency_newChan__int_95e97e5e(1)
	_g_12 := go___Select_send2_Int_Int(send2a, 11, send2b, 22)
	var hx_switch_8 *string
	switch _g_12.tag {
	case 0:
		hx_switch_8 = hxrt.StringFromLiteral("a")
	case 1:
		hx_switch_8 = hxrt.StringFromLiteral("b")
	case 2:
		hx_switch_8 = hxrt.StringFromLiteral("none")
	}
	send2 := hx_switch_8
	send2Values := hxrt.StringConcatAny(hxrt.StringConcatAny(go__concurrency_recvOr__int_95e97e5e(send2a.__hx_native, -1), hxrt.StringFromLiteral(",")), go__concurrency_recvOr__int_95e97e5e(send2b.__hx_native, -1))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("worker.count="), received))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("select.trySend="), hxrt.StdString(firstTry)), hxrt.StringFromLiteral(",")), hxrt.StdString(secondTry)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("select.recvOr="), firstRecv), hxrt.StringFromLiteral(",")), secondRecv))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("select.recv2="), recv2))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("select.send2="), send2), hxrt.StringFromLiteral(" values=")), send2Values))
}

func worker(jobs *go___Chan, results *go___Chan) {
	for true {
		job := go__concurrency_recvOr___string_f613ccd0(jobs.__hx_native, hxrt.StringFromLiteral("__stop__"))
		if hxrt.StringEqualStringPtr(job, hxrt.StringFromLiteral("__stop__")) {
			return
		}
		go__concurrency_send___string_f613ccd0(results.__hx_native, job)
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
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return
	}
	sendValue := reflect.ValueOf(value)
	if !sendValue.IsValid() {
		sendValue = reflect.Zero(chanValue.Type().Elem())
	} else if !sendValue.Type().AssignableTo(chanValue.Type().Elem()) {
		if sendValue.Type().ConvertibleTo(chanValue.Type().Elem()) {
			sendValue = sendValue.Convert(chanValue.Type().Elem())
		} else {
			return
		}
	}
	chanValue.Send(sendValue)
}

func go__concurrency_trySend(channel any, value any) bool {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return false
	}
	sendValue := reflect.ValueOf(value)
	if !sendValue.IsValid() {
		sendValue = reflect.Zero(chanValue.Type().Elem())
	} else if !sendValue.Type().AssignableTo(chanValue.Type().Elem()) {
		if sendValue.Type().ConvertibleTo(chanValue.Type().Elem()) {
			sendValue = sendValue.Convert(chanValue.Type().Elem())
		} else {
			return false
		}
	}
	cases := []reflect.SelectCase{
		{Dir: reflect.SelectSend, Chan: chanValue, Send: sendValue},
		{Dir: reflect.SelectDefault},
	}
	chosen, _, _ := reflect.Select(cases)
	return chosen == 0
}

func go__concurrency_recv(channel any) any {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return nil
	}
	recvValue, _ := chanValue.Recv()
	if !recvValue.IsValid() {
		return nil
	}
	return recvValue.Interface()
}

func go__concurrency_recvOr(channel any, defaultValue any) any {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return defaultValue
	}
	cases := []reflect.SelectCase{
		{Dir: reflect.SelectRecv, Chan: chanValue},
		{Dir: reflect.SelectDefault},
	}
	chosen, recvValue, _ := reflect.Select(cases)
	if chosen == 0 {
		if !recvValue.IsValid() {
			return defaultValue
		}
		return recvValue.Interface()
	}
	return defaultValue
}

func go__concurrency_tryRecv(channel any) *go___Result {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("empty")))
	}
	cases := []reflect.SelectCase{
		{Dir: reflect.SelectRecv, Chan: chanValue},
		{Dir: reflect.SelectDefault},
	}
	chosen, recvValue, _ := reflect.Select(cases)
	if chosen == 0 {
		if !recvValue.IsValid() {
			return New_go___Result(nil, nil)
		}
		return New_go___Result(recvValue.Interface(), nil)
	}
	return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("empty")))
}

func go__concurrency_close(channel any) {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return
	}
	chanValue.Close()
}

func go__concurrency_spawn(fn func()) {
	go fn()
}

func go__concurrency_makeChan___string_f613ccd0(buffer int) any {
	if buffer > 0 {
		return make(chan *string, buffer)
	}
	return make(chan *string)
}

func go__concurrency_setBuffer___string_f613ccd0(channel *go___Chan, buffer int) {
	if channel == nil {
		return
	}
	channel.__hx_native = go__concurrency_makeChan___string_f613ccd0(buffer)
}

func go__concurrency_newChan___string_f613ccd0(buffer int) *go___Chan {
	channel := New_go___Chan()
	go__concurrency_setBuffer___string_f613ccd0(channel, buffer)
	return channel
}

func go__concurrency_send___string_f613ccd0(channel any, value *string) {
	channel.(chan *string) <- value
}

func go__concurrency_trySend___string_f613ccd0(channel any, value *string) bool {
	select {
	case channel.(chan *string) <- value:
		return true
	default:
		return false
	}
}

func go__concurrency_recv___string_f613ccd0(channel any) *string {
	return <-channel.(chan *string)
}

func go__concurrency_recvOr___string_f613ccd0(channel any, defaultValue *string) *string {
	select {
	case value := <-channel.(chan *string):
		return value
	default:
		return defaultValue
	}
}

func go__concurrency_tryRecv___string_f613ccd0(channel any) *go___Result {
	select {
	case value := <-channel.(chan *string):
		return New_go___Result(value, nil)
	default:
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("empty")))
	}
}

func go__concurrency_close___string_f613ccd0(channel any) {
	close(channel.(chan *string))
}

func go__concurrency_makeChan__int_95e97e5e(buffer int) any {
	if buffer > 0 {
		return make(chan int, buffer)
	}
	return make(chan int)
}

func go__concurrency_setBuffer__int_95e97e5e(channel *go___Chan, buffer int) {
	if channel == nil {
		return
	}
	channel.__hx_native = go__concurrency_makeChan__int_95e97e5e(buffer)
}

func go__concurrency_newChan__int_95e97e5e(buffer int) *go___Chan {
	channel := New_go___Chan()
	go__concurrency_setBuffer__int_95e97e5e(channel, buffer)
	return channel
}

func go__concurrency_send__int_95e97e5e(channel any, value int) {
	channel.(chan int) <- value
}

func go__concurrency_trySend__int_95e97e5e(channel any, value int) bool {
	select {
	case channel.(chan int) <- value:
		return true
	default:
		return false
	}
}

func go__concurrency_recv__int_95e97e5e(channel any) int {
	return <-channel.(chan int)
}

func go__concurrency_recvOr__int_95e97e5e(channel any, defaultValue int) int {
	select {
	case value := <-channel.(chan int):
		return value
	default:
		return defaultValue
	}
}

func go__concurrency_tryRecv__int_95e97e5e(channel any) *go___Result {
	select {
	case value := <-channel.(chan int):
		return New_go___Result(value, nil)
	default:
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("empty")))
	}
}

func go__concurrency_close__int_95e97e5e(channel any) {
	close(channel.(chan int))
}

func go__result_fromValueError(value any, err error) *go___Result {
	if err != nil {
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral(err.Error())))
	}
	return New_go___Result(value, nil)
}

func go__result_ok___string_f613ccd0(value *string) *go___Result {
	return New_go___Result(value, nil)
}

func go__result_failure___string_f613ccd0(message *string) *go___Result {
	return New_go___Result(nil, New_go___Error(message))
}

func go__result_valueError___string_f613ccd0(result *go___Result) (*string, error) {
	var zero *string
	if result == nil {
		return zero, errors.New("nil go.Result")
	}
	if result.errorValue != nil {
		return zero, errors.New(*hxrt.StdString(result.errorValue.message))
	}
	if result.value == nil {
		return zero, nil
	}
	return result.value.(*string), nil
}

func go__result_isOk___string_f613ccd0(result *go___Result) bool {
	_, err := go__result_valueError___string_f613ccd0(result)
	return (err == nil)
}

func go__result_isErr___string_f613ccd0(result *go___Result) bool {
	_, err := go__result_valueError___string_f613ccd0(result)
	return (err != nil)
}

func go__result_unwrap___string_f613ccd0(result *go___Result) *string {
	value, err := go__result_valueError___string_f613ccd0(result)
	if err != nil {
		hxrt.Throw(hxrt.StringFromLiteral(err.Error()))
		var zero *string
		return zero
	}
	return value
}

func go__result_error___string_f613ccd0(result *go___Result) *string {
	_, err := go__result_valueError___string_f613ccd0(result)
	if err == nil {
		return nil
	}
	return hxrt.StringFromLiteral(err.Error())
}

func go__result_ok__int_95e97e5e(value int) *go___Result {
	return New_go___Result(value, nil)
}

func go__result_failure__int_95e97e5e(message *string) *go___Result {
	return New_go___Result(nil, New_go___Error(message))
}

func go__result_valueError__int_95e97e5e(result *go___Result) (int, error) {
	var zero int
	if result == nil {
		return zero, errors.New("nil go.Result")
	}
	if result.errorValue != nil {
		return zero, errors.New(*hxrt.StdString(result.errorValue.message))
	}
	if result.value == nil {
		return zero, nil
	}
	return result.value.(int), nil
}

func go__result_isOk__int_95e97e5e(result *go___Result) bool {
	_, err := go__result_valueError__int_95e97e5e(result)
	return (err == nil)
}

func go__result_isErr__int_95e97e5e(result *go___Result) bool {
	_, err := go__result_valueError__int_95e97e5e(result)
	return (err != nil)
}

func go__result_unwrap__int_95e97e5e(result *go___Result) int {
	value, err := go__result_valueError__int_95e97e5e(result)
	if err != nil {
		hxrt.Throw(hxrt.StringFromLiteral(err.Error()))
		var zero int
		return zero
	}
	return value
}

func go__result_error__int_95e97e5e(result *go___Result) *string {
	_, err := go__result_valueError__int_95e97e5e(result)
	if err == nil {
		return nil
	}
	return hxrt.StringFromLiteral(err.Error())
}
