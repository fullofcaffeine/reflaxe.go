package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func main() {
	defer hxrt.ThreadWaitForAll()
	mainThread := sys__thread__Thread_current()
	worker := sys__thread__Thread_create(hxrt.NewFunction(func() {
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("worker-ready"))
		var payload any = sys__thread__Thread_readMessage(true)
		mainThread.__hx_this.sendMessage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("worker-echo="), hxrt.StdString(payload)))
	}))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("thread.msg1="), hxrt.StdString(sys__thread__Thread_readMessage(true))))
	hxrt.Println(v)
	worker.__hx_this.sendMessage(hxrt.StringFromLiteral("ping"))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("thread.msg2="), hxrt.StdString(sys__thread__Thread_readMessage(true))))
	hxrt.Println(v_1)
	hxrt.TryCatch(func() {
		worker.__hx_this.get_events().__hx_this.progress()
		hxrt.Println(any(hxrt.StringFromLiteral("thread.worker_events=available")))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_payload_3
		if hx_is_nominal_0(hx_exception_2) {
			var err *sys__thread__NoEventLoopException = hx_recover_class_0(hx_exception_2)
			var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("thread.worker_events="), err.__hx_this.get_message()))
			hxrt.Println(v_2)
		} else {
			hxrt.Throw(hx_caught_1)
		}
	})
	runLoop := New_sys__thread__EventLoop()
	runLoop.__hx_this.run(hxrt.NewFunction(func() {
		hxrt.Println(any(hxrt.StringFromLiteral("loop.run=once")))
	}))
	runLoop.__hx_this.loop()
	hxrt.Println(any(hxrt.StringFromLiteral("loop.loop_after=done")))
	promisedLoop := New_sys__thread__EventLoop()
	promisedLoop.__hx_this.promise()
	promisedLoop.__hx_this.runPromised(hxrt.NewFunction(func() {
		hxrt.Println(any(hxrt.StringFromLiteral("loop.runPromised=ok")))
	}))
	promisedLoop.__hx_this.loop()
	hxrt.Println(any(hxrt.StringFromLiteral("loop.promised_after=done")))
	loop := New_sys__thread__EventLoop()
	repeats := 0
	var handler any = any(0)
	handler = loop.__hx_this.repeat(hxrt.NewFunction(func() {
		repeats = int(int32((repeats + 1)))
		hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("loop.repeat="), repeats)))
		loop.__hx_this.cancel(handler)
	}), 10)
	loop.__hx_this.loop()
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("loop.repeat_after="), repeats)))
	sys__thread__Thread_create(hxrt.NewFunction(func() {
		sys__thread__Thread_runWithEventLoop(hxrt.NewFunction(func() {
			sys__thread__Thread_current().__hx_this.get_events().__hx_this.promise()
			sys__thread__Thread_current().__hx_this.get_events().__hx_this.runPromised(hxrt.NewFunction(func() {
				mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("runWithEventLoop=ok"))
			}))
		}))
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("runWithEventLoop=after"))
	}))
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("thread.runWithEventLoop="), sortedMessages(2)))
	hxrt.Println(v_3)
	sys__thread__Thread_createWithEventLoop(hxrt.NewFunction(func() {
		sys__thread__Thread_current().__hx_this.get_events().__hx_this.promise()
		sys__thread__Thread_current().__hx_this.get_events().__hx_this.runPromised(hxrt.NewFunction(func() {
			mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("createWithEventLoop=ok"))
		}))
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("createWithEventLoop=after-job"))
	}))
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("thread.createWithEventLoop="), sortedMessages(2)))
	hxrt.Println(v_4)
	fixed := New_sys__thread__FixedThreadPool(2)
	fixed.__hx_this.run(hxrt.NewFunction(func() {
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("fixed:b"))
	}))
	fixed.__hx_this.run(hxrt.NewFunction(func() {
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("fixed:a"))
	}))
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fixed.msgs="), sortedMessages(2)))
	hxrt.Println(v_5)
	fixed.__hx_this.shutdown()
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fixed.shutdown="), hxrt.StdString(fixed.__hx_this.get_isShutdown())))
	hxrt.Println(v_6)
	hxrt.TryCatch(func() {
		fixed.__hx_this.run(hxrt.NewFunction(func() {
		}))
	}, func(hx_caught_4 any) {
		hx_exception_5 := haxe__Exception_caught(hx_caught_4)
		hx_payload_6 := haxe__Exception_catchValue(hx_caught_4)
		_ = hx_payload_6
		if hx_is_nominal_1(hx_exception_5) {
			var err_1 *sys__thread__ThreadPoolException = hx_recover_class_1(hx_exception_5)
			var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fixed.error="), err_1.__hx_this.get_message()))
			hxrt.Println(v_7)
		} else {
			hxrt.Throw(hx_caught_4)
		}
	})
	elastic := New_sys__thread__ElasticThreadPool(2, 0.05)
	elastic.__hx_this.run(hxrt.NewFunction(func() {
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("elastic:a"))
	}))
	elastic.__hx_this.run(hxrt.NewFunction(func() {
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("elastic:b"))
	}))
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("elastic.msgs="), sortedMessages(2)))
	hxrt.Println(v_8)
	elastic.__hx_this.shutdown()
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("elastic.shutdown="), hxrt.StdString(elastic.__hx_this.get_isShutdown())))
	hxrt.Println(v_9)
	hxrt.TryCatch(func() {
		elastic.__hx_this.run(hxrt.NewFunction(func() {
		}))
	}, func(hx_caught_7 any) {
		hx_exception_8 := haxe__Exception_caught(hx_caught_7)
		hx_payload_9 := haxe__Exception_catchValue(hx_caught_7)
		_ = hx_payload_9
		if hx_is_nominal_1(hx_exception_8) {
			var err_2 *sys__thread__ThreadPoolException = hx_recover_class_1(hx_exception_8)
			var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("elastic.error="), err_2.__hx_this.get_message()))
			hxrt.Println(v_10)
		} else {
			hxrt.Throw(hx_caught_7)
		}
	})
}

func sortedMessages(count int) *string {
	values := hxrt.NewArray()
	_g := 0
	_g1 := count
	for _g < _g1 {
		hx_post_10 := _g
		_g = int(int32((_g + 1)))
		hx_tmp := hx_post_10
		_ = hx_tmp
		values.Push(hxrt.StdString(sys__thread__Thread_readMessage(true)))
	}
	haxe__ds__ArraySort_sort(values, func(hx_comparator_14 *hxrt.Function[func(*string, *string) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_14, func(hx_cmp_left_12 any, hx_cmp_right_13 any) int {
			return hxrt.FunctionCall(hx_comparator_14)(func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(hx_cmp_left_12), func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(hx_cmp_right_13))
		})
	}(hx_recover_callable_0(hxrt.NamedFunction(Reflect_compare, "Reflect.compare"))))
	var buf_b *string
	buf_b = hxrt.StringFromLiteral("")
	var _g_current int
	var _g_array *hxrt.Array
	_g_current = 0
	_g_array = values
	for _g_current < _g_array.Len() {
		var _g_value *string
		var _g_key int
		_g_value = func(hx_value_19 any) *string {
			if hx_value_19 == nil {
				var hx_zero_20 *string
				return hx_zero_20
			}
			return hx_value_19.(*string)
		}(_g_array.Get(_g_current))
		hx_post_21 := _g_current
		_g_current = int(int32((_g_current + 1)))
		_g_key = hx_post_21
		index := _g_key
		value := _g_value
		if index > 0 {
			buf_b = hxrt.StringConcatStringPtr(buf_b, hxrt.StringFromLiteral(","))
		}
		buf_b = hxrt.StringConcatStringPtr(buf_b, hxrt.StdString(value))
	}
	return buf_b
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__List:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__Deque:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__ElasticThreadPool:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__ElasticThreadPoolWorker:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__EventLoop:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPool:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPoolShutdownException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPoolWorker:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__Lock:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__Mutex:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__NoEventLoopException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__Thread:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__ThreadPoolException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	case *haxe__ds__List:
		return hxrt__generated_method_field__haxe__ds__List(value, key)
	case *haxe__ds___List__GoListIterator:
		return hxrt__generated_method_field__haxe__ds___List__GoListIterator(value, key)
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt__generated_method_field__haxe__ds___List__GoListKeyValueIterator(value, key)
	case *sys__thread__Deque:
		return hxrt__generated_method_field__sys__thread__Deque(value, key)
	case *sys__thread__ElasticThreadPool:
		return hxrt__generated_method_field__sys__thread__ElasticThreadPool(value, key)
	case *sys__thread__ElasticThreadPoolWorker:
		return hxrt__generated_method_field__sys__thread__ElasticThreadPoolWorker(value, key)
	case *sys__thread__EventLoop:
		return hxrt__generated_method_field__sys__thread__EventLoop(value, key)
	case *sys__thread__FixedThreadPool:
		return hxrt__generated_method_field__sys__thread__FixedThreadPool(value, key)
	case *sys__thread__FixedThreadPoolShutdownException:
		return hxrt__generated_method_field__sys__thread__FixedThreadPoolShutdownException(value, key)
	case *sys__thread__FixedThreadPoolWorker:
		return hxrt__generated_method_field__sys__thread__FixedThreadPoolWorker(value, key)
	case *sys__thread__Lock:
		return hxrt__generated_method_field__sys__thread__Lock(value, key)
	case *sys__thread__Mutex:
		return hxrt__generated_method_field__sys__thread__Mutex(value, key)
	case *sys__thread__NoEventLoopException:
		return hxrt__generated_method_field__sys__thread__NoEventLoopException(value, key)
	case *sys__thread__Thread:
		return hxrt__generated_method_field__sys__thread__Thread(value, key)
	case *sys__thread__ThreadPoolException:
		return hxrt__generated_method_field__sys__thread__ThreadPoolException(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_103 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_103.__hx_this.details()
			}, func() any {
				return hx_receiver_103.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_104 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_104.__hx_this.get_message()
			}, func() any {
				return hx_receiver_104.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_105 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_105.__hx_this.get_native()
			}, func() any {
				return hx_receiver_105.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_106 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_106.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_106.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_107 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_107.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_107.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_108 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_108.__hx_this.toString()
			}, func() any {
				return hx_receiver_108.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_109 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_109.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_109.__hx_this
			}, "unwrap")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_110 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_110.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_110.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__haxe__ds__List(value *haxe__ds__List, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "add":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_111 := value
			return hxrt.BoundFunction(func(item any) {
				hx_receiver_111.__hx_this.add(item)
			}, func() any {
				return hx_receiver_111.__hx_this
			}, "add")
		}()
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_112 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_112.__hx_this.clear()
			}, func() any {
				return hx_receiver_112.__hx_this
			}, "clear")
		}()
	case "filter":
		return func() *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List] {
			hx_receiver_113 := value
			return hxrt.BoundFunction(func(predicate *hxrt.Function[func(any) bool]) *haxe__ds__List {
				return hx_receiver_113.__hx_this.filter(predicate)
			}, func() any {
				return hx_receiver_113.__hx_this
			}, "filter")
		}()
	case "first":
		return func() *hxrt.Function[func() any] {
			hx_receiver_114 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_114.__hx_this.first()
			}, func() any {
				return hx_receiver_114.__hx_this
			}, "first")
		}()
	case "isEmpty":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_115 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_115.__hx_this.isEmpty()
			}, func() any {
				return hx_receiver_115.__hx_this
			}, "isEmpty")
		}()
	case "iterator":
		return func() *hxrt.Function[func() *haxe__ds___List__GoListIterator] {
			hx_receiver_116 := value
			return hxrt.BoundFunction(func() *haxe__ds___List__GoListIterator {
				return hx_receiver_116.__hx_this.iterator()
			}, func() any {
				return hx_receiver_116.__hx_this
			}, "iterator")
		}()
	case "join":
		return func() *hxrt.Function[func(*string) *string] {
			hx_receiver_117 := value
			return hxrt.BoundFunction(func(separator *string) *string {
				return hx_receiver_117.__hx_this.join(separator)
			}, func() any {
				return hx_receiver_117.__hx_this
			}, "join")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator] {
			hx_receiver_118 := value
			return hxrt.BoundFunction(func() *haxe__ds___List__GoListKeyValueIterator {
				return hx_receiver_118.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_118.__hx_this
			}, "keyValueIterator")
		}()
	case "last":
		return func() *hxrt.Function[func() any] {
			hx_receiver_119 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_119.__hx_this.last()
			}, func() any {
				return hx_receiver_119.__hx_this
			}, "last")
		}()
	case "map":
		return func() *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List] {
			hx_receiver_120 := value
			return hxrt.BoundFunction(func(transform *hxrt.Function[func(any) any]) *haxe__ds__List {
				return hx_receiver_120.__hx_this.map_(transform)
			}, func() any {
				return hx_receiver_120.__hx_this
			}, "map")
		}()
	case "pop":
		return func() *hxrt.Function[func() any] {
			hx_receiver_121 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_121.__hx_this.pop()
			}, func() any {
				return hx_receiver_121.__hx_this
			}, "pop")
		}()
	case "push":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_122 := value
			return hxrt.BoundFunction(func(item any) {
				hx_receiver_122.__hx_this.push(item)
			}, func() any {
				return hx_receiver_122.__hx_this
			}, "push")
		}()
	case "remove":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_123 := value
			return hxrt.BoundFunction(func(value any) bool {
				return hx_receiver_123.__hx_this.remove(value)
			}, func() any {
				return hx_receiver_123.__hx_this
			}, "remove")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_124 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_124.__hx_this.toString()
			}, func() any {
				return hx_receiver_124.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds___List__GoListIterator(value *haxe__ds___List__GoListIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_125 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_125.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_125.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() any] {
			hx_receiver_126 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_126.__hx_this.next()
			}, func() any {
				return hx_receiver_126.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds___List__GoListKeyValueIterator(value *haxe__ds___List__GoListKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_127 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_127.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_127.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_128 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_128.__hx_this.next()
			}, func() any {
				return hx_receiver_128.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__Deque(value *sys__thread__Deque, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "add":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_129 := value
			return hxrt.BoundFunction(func(i any) {
				hx_receiver_129.__hx_this.add(i)
			}, func() any {
				return hx_receiver_129.__hx_this
			}, "add")
		}()
	case "pop":
		return func() *hxrt.Function[func(bool) any] {
			hx_receiver_130 := value
			return hxrt.BoundFunction(func(block bool) any {
				return hx_receiver_130.__hx_this.pop(block)
			}, func() any {
				return hx_receiver_130.__hx_this
			}, "pop")
		}()
	case "push":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_131 := value
			return hxrt.BoundFunction(func(i any) {
				hx_receiver_131.__hx_this.push(i)
			}, func() any {
				return hx_receiver_131.__hx_this
			}, "push")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__ElasticThreadPool(value *sys__thread__ElasticThreadPool, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "get_isShutdown":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_132 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_132.__hx_this.get_isShutdown()
			}, func() any {
				return hx_receiver_132.__hx_this
			}, "get_isShutdown")
		}()
	case "get_threadsCount":
		return func() *hxrt.Function[func() int] {
			hx_receiver_133 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_133.__hx_this.get_threadsCount()
			}, func() any {
				return hx_receiver_133.__hx_this
			}, "get_threadsCount")
		}()
	case "retireWorkerLocked":
		return func() *hxrt.Function[func(*sys__thread__ElasticThreadPoolWorker)] {
			hx_receiver_134 := value
			return hxrt.BoundFunction(func(worker *sys__thread__ElasticThreadPoolWorker) {
				hx_receiver_134.__hx_this.retireWorkerLocked(worker)
			}, func() any {
				return hx_receiver_134.__hx_this
			}, "retireWorkerLocked")
		}()
	case "run":
		return func() *hxrt.Function[func(*hxrt.Function[func()])] {
			hx_receiver_135 := value
			return hxrt.BoundFunction(func(task *hxrt.Function[func()]) {
				hx_receiver_135.__hx_this.run(task)
			}, func() any {
				return hx_receiver_135.__hx_this
			}, "run")
		}()
	case "shutdown":
		return func() *hxrt.Function[func()] {
			hx_receiver_136 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_136.__hx_this.shutdown()
			}, func() any {
				return hx_receiver_136.__hx_this
			}, "shutdown")
		}()
	case "startWorkerLocked":
		return func() *hxrt.Function[func()] {
			hx_receiver_137 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_137.__hx_this.startWorkerLocked()
			}, func() any {
				return hx_receiver_137.__hx_this
			}, "startWorkerLocked")
		}()
	case "workerResolveWait":
		return func() *hxrt.Function[func(*sys__thread__ElasticThreadPoolWorker, bool) bool] {
			hx_receiver_138 := value
			return hxrt.BoundFunction(func(worker *sys__thread__ElasticThreadPoolWorker, woke bool) bool {
				return hx_receiver_138.__hx_this.workerResolveWait(worker, woke)
			}, func() any {
				return hx_receiver_138.__hx_this
			}, "workerResolveWait")
		}()
	case "workerTaskFailed":
		return func() *hxrt.Function[func(*sys__thread__ElasticThreadPoolWorker)] {
			hx_receiver_139 := value
			return hxrt.BoundFunction(func(worker *sys__thread__ElasticThreadPoolWorker) {
				hx_receiver_139.__hx_this.workerTaskFailed(worker)
			}, func() any {
				return hx_receiver_139.__hx_this
			}, "workerTaskFailed")
		}()
	case "workerTaskFinished":
		return func() *hxrt.Function[func(*sys__thread__ElasticThreadPoolWorker)] {
			hx_receiver_140 := value
			return hxrt.BoundFunction(func(worker *sys__thread__ElasticThreadPoolWorker) {
				hx_receiver_140.__hx_this.workerTaskFinished(worker)
			}, func() any {
				return hx_receiver_140.__hx_this
			}, "workerTaskFinished")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__ElasticThreadPoolWorker(value *sys__thread__ElasticThreadPoolWorker, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "loop":
		return func() *hxrt.Function[func()] {
			hx_receiver_141 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_141.__hx_this.loop()
			}, func() any {
				return hx_receiver_141.__hx_this
			}, "loop")
		}()
	case "start":
		return func() *hxrt.Function[func()] {
			hx_receiver_142 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_142.__hx_this.start()
			}, func() any {
				return hx_receiver_142.__hx_this
			}, "start")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__EventLoop(value *sys__thread__EventLoop, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "cancel":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_143 := value
			return hxrt.BoundFunction(func(eventHandler any) {
				hx_receiver_143.__hx_this.cancel(eventHandler)
			}, func() any {
				return hx_receiver_143.__hx_this
			}, "cancel")
		}()
	case "loop":
		return func() *hxrt.Function[func()] {
			hx_receiver_144 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_144.__hx_this.loop()
			}, func() any {
				return hx_receiver_144.__hx_this
			}, "loop")
		}()
	case "progress":
		return func() *hxrt.Function[func() *sys__thread__NextEventTime] {
			hx_receiver_145 := value
			return hxrt.BoundFunction(func() *sys__thread__NextEventTime {
				return hx_receiver_145.__hx_this.progress()
			}, func() any {
				return hx_receiver_145.__hx_this
			}, "progress")
		}()
	case "promise":
		return func() *hxrt.Function[func()] {
			hx_receiver_146 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_146.__hx_this.promise()
			}, func() any {
				return hx_receiver_146.__hx_this
			}, "promise")
		}()
	case "repeat":
		return func() *hxrt.Function[func(*hxrt.Function[func()], int) any] {
			hx_receiver_147 := value
			return hxrt.BoundFunction(func(event *hxrt.Function[func()], intervalMs int) any {
				return hx_receiver_147.__hx_this.repeat(event, intervalMs)
			}, func() any {
				return hx_receiver_147.__hx_this
			}, "repeat")
		}()
	case "run":
		return func() *hxrt.Function[func(*hxrt.Function[func()])] {
			hx_receiver_148 := value
			return hxrt.BoundFunction(func(event *hxrt.Function[func()]) {
				hx_receiver_148.__hx_this.run(event)
			}, func() any {
				return hx_receiver_148.__hx_this
			}, "run")
		}()
	case "runPromised":
		return func() *hxrt.Function[func(*hxrt.Function[func()])] {
			hx_receiver_149 := value
			return hxrt.BoundFunction(func(event *hxrt.Function[func()]) {
				hx_receiver_149.__hx_this.runPromised(event)
			}, func() any {
				return hx_receiver_149.__hx_this
			}, "runPromised")
		}()
	case "wait":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_150 := value
			return hxrt.BoundFunction(func(timeout any) bool {
				return hx_receiver_150.__hx_this.wait(timeout)
			}, func() any {
				return hx_receiver_150.__hx_this
			}, "wait")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__FixedThreadPool(value *sys__thread__FixedThreadPool, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "get_isShutdown":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_151 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_151.__hx_this.get_isShutdown()
			}, func() any {
				return hx_receiver_151.__hx_this
			}, "get_isShutdown")
		}()
	case "get_threadsCount":
		return func() *hxrt.Function[func() int] {
			hx_receiver_152 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_152.__hx_this.get_threadsCount()
			}, func() any {
				return hx_receiver_152.__hx_this
			}, "get_threadsCount")
		}()
	case "run":
		return func() *hxrt.Function[func(*hxrt.Function[func()])] {
			hx_receiver_153 := value
			return hxrt.BoundFunction(func(task *hxrt.Function[func()]) {
				hx_receiver_153.__hx_this.run(task)
			}, func() any {
				return hx_receiver_153.__hx_this
			}, "run")
		}()
	case "shutdown":
		return func() *hxrt.Function[func()] {
			hx_receiver_154 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_154.__hx_this.shutdown()
			}, func() any {
				return hx_receiver_154.__hx_this
			}, "shutdown")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__FixedThreadPoolShutdownException(value *sys__thread__FixedThreadPoolShutdownException, key string) any {
	if value == nil {
		return nil
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__sys__thread__FixedThreadPoolWorker(value *sys__thread__FixedThreadPoolWorker, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "loop":
		return func() *hxrt.Function[func()] {
			hx_receiver_155 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_155.__hx_this.loop()
			}, func() any {
				return hx_receiver_155.__hx_this
			}, "loop")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__Lock(value *sys__thread__Lock, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "release":
		return func() *hxrt.Function[func()] {
			hx_receiver_156 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_156.__hx_this.release()
			}, func() any {
				return hx_receiver_156.__hx_this
			}, "release")
		}()
	case "wait":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_157 := value
			return hxrt.BoundFunction(func(timeout any) bool {
				return hx_receiver_157.__hx_this.wait(timeout)
			}, func() any {
				return hx_receiver_157.__hx_this
			}, "wait")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__Mutex(value *sys__thread__Mutex, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "acquire":
		return func() *hxrt.Function[func()] {
			hx_receiver_158 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_158.__hx_this.acquire()
			}, func() any {
				return hx_receiver_158.__hx_this
			}, "acquire")
		}()
	case "release":
		return func() *hxrt.Function[func()] {
			hx_receiver_159 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_159.__hx_this.release()
			}, func() any {
				return hx_receiver_159.__hx_this
			}, "release")
		}()
	case "tryAcquire":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_160 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_160.__hx_this.tryAcquire()
			}, func() any {
				return hx_receiver_160.__hx_this
			}, "tryAcquire")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__NoEventLoopException(value *sys__thread__NoEventLoopException, key string) any {
	if value == nil {
		return nil
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__sys__thread__Thread(value *sys__thread__Thread, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "get_events":
		return func() *hxrt.Function[func() *sys__thread__EventLoop] {
			hx_receiver_161 := value
			return hxrt.BoundFunction(func() *sys__thread__EventLoop {
				return hx_receiver_161.__hx_this.get_events()
			}, func() any {
				return hx_receiver_161.__hx_this
			}, "get_events")
		}()
	case "sendMessage":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_162 := value
			return hxrt.BoundFunction(func(msg any) {
				hx_receiver_162.__hx_this.sendMessage(msg)
			}, func() any {
				return hx_receiver_162.__hx_this
			}, "sendMessage")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__ThreadPoolException(value *sys__thread__ThreadPoolException, key string) any {
	if value == nil {
		return nil
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	classValue, ok := value.(*hxrt__TypeClassValue)
	if !ok || classValue == nil {
		return nil, false
	}
	className := *hxrt.StdString(classValue.name)
	switch className {
	default:
		return nil, false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedField(object any, field *string) any {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__List:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__Deque:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__ElasticThreadPool:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__ElasticThreadPoolWorker:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__EventLoop:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPool:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPoolShutdownException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPoolWorker:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__Lock:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__Mutex:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__NoEventLoopException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__Thread:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__ThreadPoolException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	case *haxe__ds__List:
		return hxrt__generated_field_lookup__haxe__ds__List(value, key)
	case *haxe__ds___List__GoListIterator:
		return hxrt__generated_field_lookup__haxe__ds___List__GoListIterator(value, key)
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__ds___List__GoListKeyValueIterator(value, key)
	case *sys__thread__Deque:
		return hxrt__generated_field_lookup__sys__thread__Deque(value, key)
	case *sys__thread__ElasticThreadPool:
		return hxrt__generated_field_lookup__sys__thread__ElasticThreadPool(value, key)
	case *sys__thread__ElasticThreadPoolWorker:
		return hxrt__generated_field_lookup__sys__thread__ElasticThreadPoolWorker(value, key)
	case *sys__thread__EventLoop:
		return hxrt__generated_field_lookup__sys__thread__EventLoop(value, key)
	case *sys__thread__FixedThreadPool:
		return hxrt__generated_field_lookup__sys__thread__FixedThreadPool(value, key)
	case *sys__thread__FixedThreadPoolShutdownException:
		return hxrt__generated_field_lookup__sys__thread__FixedThreadPoolShutdownException(value, key)
	case *sys__thread__FixedThreadPoolWorker:
		return hxrt__generated_field_lookup__sys__thread__FixedThreadPoolWorker(value, key)
	case *sys__thread__Lock:
		return hxrt__generated_field_lookup__sys__thread__Lock(value, key)
	case *sys__thread__Mutex:
		return hxrt__generated_field_lookup__sys__thread__Mutex(value, key)
	case *sys__thread__NoEventLoopException:
		return hxrt__generated_field_lookup__sys__thread__NoEventLoopException(value, key)
	case *sys__thread__Thread:
		return hxrt__generated_field_lookup__sys__thread__Thread(value, key)
	case *sys__thread__ThreadPoolException:
		return hxrt__generated_field_lookup__sys__thread__ThreadPoolException(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__List:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__Deque:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__ElasticThreadPool:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__ElasticThreadPoolWorker:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__EventLoop:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPool:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPoolShutdownException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPoolWorker:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__Lock:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__Mutex:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__NoEventLoopException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__Thread:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__ThreadPoolException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	case *haxe__ds__List:
		return hxrt__generated_field_has__haxe__ds__List(value, key)
	case *haxe__ds___List__GoListIterator:
		return hxrt__generated_field_has__haxe__ds___List__GoListIterator(value, key)
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt__generated_field_has__haxe__ds___List__GoListKeyValueIterator(value, key)
	case *sys__thread__Deque:
		return hxrt__generated_field_has__sys__thread__Deque(value, key)
	case *sys__thread__ElasticThreadPool:
		return hxrt__generated_field_has__sys__thread__ElasticThreadPool(value, key)
	case *sys__thread__ElasticThreadPoolWorker:
		return hxrt__generated_field_has__sys__thread__ElasticThreadPoolWorker(value, key)
	case *sys__thread__EventLoop:
		return hxrt__generated_field_has__sys__thread__EventLoop(value, key)
	case *sys__thread__FixedThreadPool:
		return hxrt__generated_field_has__sys__thread__FixedThreadPool(value, key)
	case *sys__thread__FixedThreadPoolShutdownException:
		return hxrt__generated_field_has__sys__thread__FixedThreadPoolShutdownException(value, key)
	case *sys__thread__FixedThreadPoolWorker:
		return hxrt__generated_field_has__sys__thread__FixedThreadPoolWorker(value, key)
	case *sys__thread__Lock:
		return hxrt__generated_field_has__sys__thread__Lock(value, key)
	case *sys__thread__Mutex:
		return hxrt__generated_field_has__sys__thread__Mutex(value, key)
	case *sys__thread__NoEventLoopException:
		return hxrt__generated_field_has__sys__thread__NoEventLoopException(value, key)
	case *sys__thread__Thread:
		return hxrt__generated_field_has__sys__thread__Thread(value, key)
	case *sys__thread__ThreadPoolException:
		return hxrt__generated_field_has__sys__thread__ThreadPoolException(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__List:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__Deque:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__ElasticThreadPool:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__ElasticThreadPoolWorker:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__EventLoop:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPool:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPoolShutdownException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPoolWorker:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__Lock:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__Mutex:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__NoEventLoopException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__Thread:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__thread__ThreadPoolException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	case *haxe__ds__List:
		return hxrt__generated_field_set__haxe__ds__List(value, key, incoming)
	case *haxe__ds___List__GoListIterator:
		return hxrt__generated_field_set__haxe__ds___List__GoListIterator(value, key, incoming)
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt__generated_field_set__haxe__ds___List__GoListKeyValueIterator(value, key, incoming)
	case *sys__thread__Deque:
		return hxrt__generated_field_set__sys__thread__Deque(value, key, incoming)
	case *sys__thread__ElasticThreadPool:
		return hxrt__generated_field_set__sys__thread__ElasticThreadPool(value, key, incoming)
	case *sys__thread__ElasticThreadPoolWorker:
		return hxrt__generated_field_set__sys__thread__ElasticThreadPoolWorker(value, key, incoming)
	case *sys__thread__EventLoop:
		return hxrt__generated_field_set__sys__thread__EventLoop(value, key, incoming)
	case *sys__thread__FixedThreadPool:
		return hxrt__generated_field_set__sys__thread__FixedThreadPool(value, key, incoming)
	case *sys__thread__FixedThreadPoolShutdownException:
		return hxrt__generated_field_set__sys__thread__FixedThreadPoolShutdownException(value, key, incoming)
	case *sys__thread__FixedThreadPoolWorker:
		return hxrt__generated_field_set__sys__thread__FixedThreadPoolWorker(value, key, incoming)
	case *sys__thread__Lock:
		return hxrt__generated_field_set__sys__thread__Lock(value, key, incoming)
	case *sys__thread__Mutex:
		return hxrt__generated_field_set__sys__thread__Mutex(value, key, incoming)
	case *sys__thread__NoEventLoopException:
		return hxrt__generated_field_set__sys__thread__NoEventLoopException(value, key, incoming)
	case *sys__thread__Thread:
		return hxrt__generated_field_set__sys__thread__Thread(value, key, incoming)
	case *sys__thread__ThreadPoolException:
		return hxrt__generated_field_set__sys__thread__ThreadPoolException(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__List:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds___List__GoListKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__Deque:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__ElasticThreadPool:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__ElasticThreadPoolWorker:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__EventLoop:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPool:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPoolShutdownException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__FixedThreadPoolWorker:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__Lock:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__Mutex:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__NoEventLoopException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__Thread:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__thread__ThreadPoolException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__ds__List:
		return hxrt.NewArray(hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("length"))
	case *haxe__ds___List__GoListIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("index"))
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("index"))
	case *sys__thread__Deque:
		return hxrt.NewArray(hxrt.StringFromLiteral("__mutex"), hxrt.StringFromLiteral("__available"), hxrt.StringFromLiteral("__items"))
	case *sys__thread__ElasticThreadPool:
		return hxrt.NewArray(hxrt.StringFromLiteral("maxThreadsCount"), hxrt.StringFromLiteral("_isShutdown"), hxrt.StringFromLiteral("liveWorkers"), hxrt.StringFromLiteral("pendingTasks"), hxrt.StringFromLiteral("pool"), hxrt.StringFromLiteral("queue"), hxrt.StringFromLiteral("available"), hxrt.StringFromLiteral("mutex"), hxrt.StringFromLiteral("threadTimeout"))
	case *sys__thread__ElasticThreadPoolWorker:
		return hxrt.NewArray(hxrt.StringFromLiteral("task"), hxrt.StringFromLiteral("dead"), hxrt.StringFromLiteral("owner"), hxrt.StringFromLiteral("available"), hxrt.StringFromLiteral("timeout"))
	case *sys__thread__EventLoop:
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"))
	case *sys__thread__FixedThreadPool:
		return hxrt.NewArray(hxrt.StringFromLiteral("_isShutdown"), hxrt.StringFromLiteral("pool"), hxrt.StringFromLiteral("queue"), hxrt.StringFromLiteral("mutex"))
	case *sys__thread__FixedThreadPoolShutdownException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *sys__thread__FixedThreadPoolWorker:
		return hxrt.NewArray(hxrt.StringFromLiteral("queue"))
	case *sys__thread__Lock:
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"))
	case *sys__thread__Mutex:
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"))
	case *sys__thread__NoEventLoopException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *sys__thread__Thread:
		return hxrt.NewArray(hxrt.StringFromLiteral("__id"))
	case *sys__thread__ThreadPoolException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "savedMessage":
		return hxrt.ReadStoredField(value.savedMessage)
	case "savedNative":
		return hxrt.ReadStoredField(value.savedNative)
	case "savedPrevious":
		return hxrt.ReadStoredField(value.savedPrevious)
	case "savedStack":
		return hxrt.ReadStoredField(value.savedStack)
	}
	return nil
}

func hxrt__generated_field_has__haxe__Exception(value *haxe__Exception, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		return true
	case "savedNative":
		return true
	case "savedPrevious":
		return true
	case "savedStack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Exception(value *haxe__Exception, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		if incoming == nil {
			var zero *string
			value.savedMessage = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.savedMessage = typed
			return true
		default:
			return false
		}
	case "savedNative":
		if incoming == nil {
			var zero any
			value.savedNative = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedNative = typed
			return true
		default:
			return false
		}
	case "savedPrevious":
		if incoming == nil {
			var zero *haxe__Exception
			value.savedPrevious = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__Exception:
			value.savedPrevious = typed
			return true
		default:
			return false
		}
	case "savedStack":
		if incoming == nil {
			var zero any
			value.savedStack = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedStack = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "value":
		return hxrt.ReadStoredField(value.value)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__ValueException(value *haxe__ValueException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__ValueException(value *haxe__ValueException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		if incoming == nil {
			var zero any
			value.value = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.value = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return hxrt.ReadStoredField(value.high)
	case "low":
		return hxrt.ReadStoredField(value.low)
	}
	return nil
}

func hxrt__generated_field_has__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		return true
	case "low":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		if incoming == nil {
			var zero int
			value.high = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.high = typed
			return true
		default:
			return false
		}
	case "low":
		if incoming == nil {
			var zero int
			value.low = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.low = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds__List(value *haxe__ds__List, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "items":
		return hxrt.ReadStoredField(value.items)
	case "length":
		return hxrt.ReadStoredField(value.length)
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__List(value *haxe__ds__List, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "items":
		return true
	case "length":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__List(value *haxe__ds__List, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "items":
		if incoming == nil {
			var zero *hxrt.Array
			value.items = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.items = typed
			return true
		default:
			return false
		}
	case "length":
		if incoming == nil {
			var zero int
			value.length = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.length = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds___List__GoListIterator(value *haxe__ds___List__GoListIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "index":
		return hxrt.ReadStoredField(value.index)
	case "items":
		return hxrt.ReadStoredField(value.items)
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds___List__GoListIterator(value *haxe__ds___List__GoListIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		return true
	case "items":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds___List__GoListIterator(value *haxe__ds___List__GoListIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		if incoming == nil {
			var zero int
			value.index = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.index = typed
			return true
		default:
			return false
		}
	case "items":
		if incoming == nil {
			var zero *hxrt.Array
			value.items = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.items = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds___List__GoListKeyValueIterator(value *haxe__ds___List__GoListKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "index":
		return hxrt.ReadStoredField(value.index)
	case "items":
		return hxrt.ReadStoredField(value.items)
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds___List__GoListKeyValueIterator(value *haxe__ds___List__GoListKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		return true
	case "items":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds___List__GoListKeyValueIterator(value *haxe__ds___List__GoListKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "index":
		if incoming == nil {
			var zero int
			value.index = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.index = typed
			return true
		default:
			return false
		}
	case "items":
		if incoming == nil {
			var zero *hxrt.Array
			value.items = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.items = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__thread__Deque(value *sys__thread__Deque, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__available":
		return hxrt.ReadStoredField(value.__available)
	case "__items":
		return hxrt.ReadStoredField(value.__items)
	case "__mutex":
		return hxrt.ReadStoredField(value.__mutex)
	}
	return nil
}

func hxrt__generated_field_has__sys__thread__Deque(value *sys__thread__Deque, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__available":
		return true
	case "__items":
		return true
	case "__mutex":
		return true
	}
	return false
}

func hxrt__generated_field_set__sys__thread__Deque(value *sys__thread__Deque, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__available":
		if incoming == nil {
			var zero *sys__thread__Lock
			value.__available = zero
			return true
		}
		switch typed := incoming.(type) {
		case *sys__thread__Lock:
			value.__available = typed
			return true
		default:
			return false
		}
	case "__items":
		if incoming == nil {
			var zero *haxe__ds__List
			value.__items = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds__List:
			value.__items = typed
			return true
		default:
			return false
		}
	case "__mutex":
		if incoming == nil {
			var zero *sys__thread__Mutex
			value.__mutex = zero
			return true
		}
		switch typed := incoming.(type) {
		case *sys__thread__Mutex:
			value.__mutex = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__thread__ElasticThreadPool(value *sys__thread__ElasticThreadPool, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "_isShutdown":
		return hxrt.ReadStoredField(value._isShutdown)
	case "available":
		return hxrt.ReadStoredField(value.available)
	case "liveWorkers":
		return hxrt.ReadStoredField(value.liveWorkers)
	case "maxThreadsCount":
		return hxrt.ReadStoredField(value.maxThreadsCount)
	case "mutex":
		return hxrt.ReadStoredField(value.mutex)
	case "pendingTasks":
		return hxrt.ReadStoredField(value.pendingTasks)
	case "pool":
		return hxrt.ReadStoredField(value.pool)
	case "queue":
		return hxrt.ReadStoredField(value.queue)
	case "threadTimeout":
		return hxrt.ReadStoredField(value.threadTimeout)
	}
	return nil
}

func hxrt__generated_field_has__sys__thread__ElasticThreadPool(value *sys__thread__ElasticThreadPool, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "_isShutdown":
		return true
	case "available":
		return true
	case "liveWorkers":
		return true
	case "maxThreadsCount":
		return true
	case "mutex":
		return true
	case "pendingTasks":
		return true
	case "pool":
		return true
	case "queue":
		return true
	case "threadTimeout":
		return true
	}
	return false
}

func hxrt__generated_field_set__sys__thread__ElasticThreadPool(value *sys__thread__ElasticThreadPool, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "_isShutdown":
		if incoming == nil {
			var zero bool
			value._isShutdown = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value._isShutdown = typed
			return true
		default:
			return false
		}
	case "available":
		if incoming == nil {
			var zero *sys__thread__Lock
			value.available = zero
			return true
		}
		switch typed := incoming.(type) {
		case *sys__thread__Lock:
			value.available = typed
			return true
		default:
			return false
		}
	case "liveWorkers":
		if incoming == nil {
			var zero int
			value.liveWorkers = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.liveWorkers = typed
			return true
		default:
			return false
		}
	case "maxThreadsCount":
		if incoming == nil {
			var zero int
			value.maxThreadsCount = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.maxThreadsCount = typed
			return true
		default:
			return false
		}
	case "mutex":
		if incoming == nil {
			var zero *sys__thread__Mutex
			value.mutex = zero
			return true
		}
		switch typed := incoming.(type) {
		case *sys__thread__Mutex:
			value.mutex = typed
			return true
		default:
			return false
		}
	case "pendingTasks":
		if incoming == nil {
			var zero int
			value.pendingTasks = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.pendingTasks = typed
			return true
		default:
			return false
		}
	case "pool":
		if incoming == nil {
			var zero *hxrt.Array
			value.pool = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.pool = typed
			return true
		default:
			return false
		}
	case "queue":
		if incoming == nil {
			var zero *sys__thread__Deque
			value.queue = zero
			return true
		}
		switch typed := incoming.(type) {
		case *sys__thread__Deque:
			value.queue = typed
			return true
		default:
			return false
		}
	case "threadTimeout":
		if incoming == nil {
			var zero float64
			value.threadTimeout = zero
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.threadTimeout = typed
			return true
		case int:
			value.threadTimeout = float64(typed)
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__thread__ElasticThreadPoolWorker(value *sys__thread__ElasticThreadPoolWorker, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "available":
		return hxrt.ReadStoredField(value.available)
	case "dead":
		return hxrt.ReadStoredField(value.dead)
	case "owner":
		return hxrt.ReadStoredField(value.owner)
	case "task":
		return hxrt.ReadStoredField(value.task)
	case "timeout":
		return hxrt.ReadStoredField(value.timeout)
	}
	return nil
}

func hxrt__generated_field_has__sys__thread__ElasticThreadPoolWorker(value *sys__thread__ElasticThreadPoolWorker, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "available":
		return true
	case "dead":
		return true
	case "owner":
		return true
	case "task":
		return true
	case "timeout":
		return true
	}
	return false
}

func hxrt__generated_field_set__sys__thread__ElasticThreadPoolWorker(value *sys__thread__ElasticThreadPoolWorker, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "available":
		if incoming == nil {
			var zero *sys__thread__Lock
			value.available = zero
			return true
		}
		switch typed := incoming.(type) {
		case *sys__thread__Lock:
			value.available = typed
			return true
		default:
			return false
		}
	case "dead":
		if incoming == nil {
			var zero bool
			value.dead = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.dead = typed
			return true
		default:
			return false
		}
	case "owner":
		if incoming == nil {
			var zero *sys__thread__ElasticThreadPool
			value.owner = zero
			return true
		}
		switch typed := incoming.(type) {
		case *sys__thread__ElasticThreadPool:
			value.owner = typed
			return true
		default:
			return false
		}
	case "task":
		if (hxrt.FunctionOrigin(incoming) == nil) || hxrt.ReflectIsFunction(incoming) {
			value.task = hx_recover_callable_1(incoming)
			return true
		}
		return false
	case "timeout":
		if incoming == nil {
			var zero float64
			value.timeout = zero
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.timeout = typed
			return true
		case int:
			value.timeout = float64(typed)
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__thread__EventLoop(value *sys__thread__EventLoop, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__h":
		return hxrt.ReadStoredField(value.__h)
	}
	return nil
}

func hxrt__generated_field_has__sys__thread__EventLoop(value *sys__thread__EventLoop, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__h":
		return true
	}
	return false
}

func hxrt__generated_field_set__sys__thread__EventLoop(value *sys__thread__EventLoop, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__h":
		if incoming == nil {
			var zero *hxrt.EventLoopHandle
			value.__h = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.EventLoopHandle:
			value.__h = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__thread__FixedThreadPool(value *sys__thread__FixedThreadPool, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "_isShutdown":
		return hxrt.ReadStoredField(value._isShutdown)
	case "mutex":
		return hxrt.ReadStoredField(value.mutex)
	case "pool":
		return hxrt.ReadStoredField(value.pool)
	case "queue":
		return hxrt.ReadStoredField(value.queue)
	}
	return nil
}

func hxrt__generated_field_has__sys__thread__FixedThreadPool(value *sys__thread__FixedThreadPool, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "_isShutdown":
		return true
	case "mutex":
		return true
	case "pool":
		return true
	case "queue":
		return true
	}
	return false
}

func hxrt__generated_field_set__sys__thread__FixedThreadPool(value *sys__thread__FixedThreadPool, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "_isShutdown":
		if incoming == nil {
			var zero bool
			value._isShutdown = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value._isShutdown = typed
			return true
		default:
			return false
		}
	case "mutex":
		if incoming == nil {
			var zero *sys__thread__Mutex
			value.mutex = zero
			return true
		}
		switch typed := incoming.(type) {
		case *sys__thread__Mutex:
			value.mutex = typed
			return true
		default:
			return false
		}
	case "pool":
		if incoming == nil {
			var zero *hxrt.Array
			value.pool = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.pool = typed
			return true
		default:
			return false
		}
	case "queue":
		if incoming == nil {
			var zero *sys__thread__Deque
			value.queue = zero
			return true
		}
		switch typed := incoming.(type) {
		case *sys__thread__Deque:
			value.queue = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__thread__FixedThreadPoolShutdownException(value *sys__thread__FixedThreadPoolShutdownException, key string) any {
	if value == nil {
		return nil
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__sys__thread__FixedThreadPoolShutdownException(value *sys__thread__FixedThreadPoolShutdownException, key string) bool {
	if value == nil {
		return false
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__sys__thread__FixedThreadPoolShutdownException(value *sys__thread__FixedThreadPoolShutdownException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__sys__thread__FixedThreadPoolWorker(value *sys__thread__FixedThreadPoolWorker, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "queue":
		return hxrt.ReadStoredField(value.queue)
	}
	return nil
}

func hxrt__generated_field_has__sys__thread__FixedThreadPoolWorker(value *sys__thread__FixedThreadPoolWorker, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "queue":
		return true
	}
	return false
}

func hxrt__generated_field_set__sys__thread__FixedThreadPoolWorker(value *sys__thread__FixedThreadPoolWorker, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "queue":
		if incoming == nil {
			var zero *sys__thread__Deque
			value.queue = zero
			return true
		}
		switch typed := incoming.(type) {
		case *sys__thread__Deque:
			value.queue = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__thread__Lock(value *sys__thread__Lock, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__h":
		return hxrt.ReadStoredField(value.__h)
	}
	return nil
}

func hxrt__generated_field_has__sys__thread__Lock(value *sys__thread__Lock, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__h":
		return true
	}
	return false
}

func hxrt__generated_field_set__sys__thread__Lock(value *sys__thread__Lock, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__h":
		if incoming == nil {
			var zero *hxrt.LockHandle
			value.__h = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.LockHandle:
			value.__h = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__thread__Mutex(value *sys__thread__Mutex, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__h":
		return hxrt.ReadStoredField(value.__h)
	}
	return nil
}

func hxrt__generated_field_has__sys__thread__Mutex(value *sys__thread__Mutex, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__h":
		return true
	}
	return false
}

func hxrt__generated_field_set__sys__thread__Mutex(value *sys__thread__Mutex, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__h":
		if incoming == nil {
			var zero *hxrt.MutexHandle
			value.__h = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.MutexHandle:
			value.__h = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__thread__NoEventLoopException(value *sys__thread__NoEventLoopException, key string) any {
	if value == nil {
		return nil
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__sys__thread__NoEventLoopException(value *sys__thread__NoEventLoopException, key string) bool {
	if value == nil {
		return false
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__sys__thread__NoEventLoopException(value *sys__thread__NoEventLoopException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__sys__thread__Thread(value *sys__thread__Thread, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__id":
		return hxrt.ReadStoredField(value.__id)
	}
	return nil
}

func hxrt__generated_field_has__sys__thread__Thread(value *sys__thread__Thread, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__id":
		return true
	}
	return false
}

func hxrt__generated_field_set__sys__thread__Thread(value *sys__thread__Thread, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__id":
		if incoming == nil {
			var zero int
			value.__id = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.__id = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__thread__ThreadPoolException(value *sys__thread__ThreadPoolException, key string) any {
	if value == nil {
		return nil
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__sys__thread__ThreadPoolException(value *sys__thread__ThreadPoolException, key string) bool {
	if value == nil {
		return false
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__sys__thread__ThreadPoolException(value *sys__thread__ThreadPoolException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func reflaxe__go___internal__CompilerReflect_typeField(object any, field *string) any {
	key := *hxrt.StdString(field)
	value, found := hxrt_typeClassMetadataField(object, key)
	if !found {
		return nil
	}
	return value
}

func reflaxe__go___internal__CompilerReflect_hasTypeField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	_, found := hxrt_typeClassMetadataField(object, key)
	return found
}

func reflaxe__go___internal__CompilerReflect_generatedMethod(object any, field *string) any {
	key := *hxrt.StdString(field)
	return hxrt__generated_method_field(object, key)
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	switch enumValue := value.(type) {
	case *haxe__StackItem:
		return (enumValue != nil)
	case *sys__thread__NextEventTime:
		return (enumValue != nil)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func()]:
		hxrt.FunctionCall(source)()
		return nil
	case *hxrt.Function[func(any, any) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(argument_0, argument_1)
	case *hxrt.Function[func(any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(argument_0)
		return nil
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_2(argument_0))
	case *hxrt.Function[func() bool]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_163 any) *string {
			if hx_value_163 == nil {
				var hx_zero_164 *string
				return hx_zero_164
			}
			return hx_value_163.(*string)
		}(any(argument_0)))
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_3(argument_0))
	case *hxrt.Function[func(any) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(bool) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_165 any) bool {
			if hx_value_165 == nil {
				var hx_zero_166 bool
				return hx_zero_166
			}
			return hx_value_165.(bool)
		}(any(argument_0)))
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*sys__thread__ElasticThreadPoolWorker)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_4(argument_0))
		return nil
	case *hxrt.Function[func(*hxrt.Function[func()])]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_callable_1(argument_0))
		return nil
	case *hxrt.Function[func(*sys__thread__ElasticThreadPoolWorker, bool) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_4(argument_0), func(hx_value_167 any) bool {
			if hx_value_167 == nil {
				var hx_zero_168 bool
				return hx_zero_168
			}
			return hx_value_167.(bool)
		}(any(argument_1)))
	case *hxrt.Function[func() *sys__thread__NextEventTime]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*hxrt.Function[func()], int) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_1(argument_0), func(hx_value_169 any) int {
			if hx_value_169 == nil {
				var hx_zero_170 int
				return hx_zero_170
			}
			return hx_value_169.(int)
		}(any(argument_1)))
	case *hxrt.Function[func() *sys__thread__EventLoop]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func(*string, *string) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*string, *string) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any, any) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string, arg_1 *string) int {
			return hxrt.FunctionCall(source)(arg_0, arg_1)
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *string, arg_1 *string) int {
			return func(hx_value_171 any) int {
				if hx_value_171 == nil {
					var hx_zero_172 int
					return hx_zero_172
				}
				return hx_value_171.(int)
			}(any(hxrt.FunctionCall(source)(arg_0, arg_1)))
		})
	case *hxrt.Function[func(*sys__thread__ElasticThreadPoolWorker, bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string, arg_1 *string) int {
			return func(hx_value_175 any) int {
				if hx_value_175 == nil {
					var hx_zero_176 int
					return hx_zero_176
				}
				return hx_value_175.(int)
			}(any(hxrt.FunctionCall(source)(hx_recover_class_4(arg_0), func(hx_value_173 any) bool {
				if hx_value_173 == nil {
					var hx_zero_174 bool
					return hx_zero_174
				}
				return hx_value_173.(bool)
			}(any(arg_1)))))
		})
	case *hxrt.Function[func(*hxrt.Function[func()], int) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *string, arg_1 *string) int {
			return func(hx_value_179 any) int {
				if hx_value_179 == nil {
					var hx_zero_180 int
					return hx_zero_180
				}
				return hx_value_179.(int)
			}(any(hxrt.FunctionCall(source)(hx_recover_callable_1(arg_0), func(hx_value_177 any) int {
				if hx_value_177 == nil {
					var hx_zero_178 int
					return hx_zero_178
				}
				return hx_value_177.(int)
			}(any(arg_1)))))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func()] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func()]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func()]:
		return source
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() {
			hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() {
			hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() {
			hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() {
			hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() {
			hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() {
			hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() {
			hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() {
			hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() {
			hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *sys__thread__NextEventTime]:
		return hxrt.ProjectFunction(source, func() {
			hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *sys__thread__EventLoop]:
		return hxrt.ProjectFunction(source, func() {
			hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func(any) bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any) bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			hxrt.FunctionCall(source)(arg_0)
			return func(hx_value_181 any) bool {
				if hx_value_181 == nil {
					var hx_zero_182 bool
					return hx_zero_182
				}
				return hx_value_181.(bool)
			}(any(nil))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_183 any) bool {
				if hx_value_183 == nil {
					var hx_zero_184 bool
					return hx_zero_184
				}
				return hx_value_183.(bool)
			}(any(hxrt.FunctionCall(source)(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_185 any) bool {
				if hx_value_185 == nil {
					var hx_zero_186 bool
					return hx_zero_186
				}
				return hx_value_185.(bool)
			}(any(hxrt.FunctionCall(source)(hx_recover_callable_2(arg_0))))
		})
	case *hxrt.Function[func(*string) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_189 any) bool {
				if hx_value_189 == nil {
					var hx_zero_190 bool
					return hx_zero_190
				}
				return hx_value_189.(bool)
			}(any(hxrt.FunctionCall(source)(func(hx_value_187 any) *string {
				if hx_value_187 == nil {
					var hx_zero_188 *string
					return hx_zero_188
				}
				return hx_value_187.(*string)
			}(any(arg_0)))))
		})
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_191 any) bool {
				if hx_value_191 == nil {
					var hx_zero_192 bool
					return hx_zero_192
				}
				return hx_value_191.(bool)
			}(any(hxrt.FunctionCall(source)(hx_recover_callable_3(arg_0))))
		})
	case *hxrt.Function[func(any) bool]:
		return source
	case *hxrt.Function[func(bool) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			return func(hx_value_195 any) bool {
				if hx_value_195 == nil {
					var hx_zero_196 bool
					return hx_zero_196
				}
				return hx_value_195.(bool)
			}(any(hxrt.FunctionCall(source)(func(hx_value_193 any) bool {
				if hx_value_193 == nil {
					var hx_zero_194 bool
					return hx_zero_194
				}
				return hx_value_193.(bool)
			}(any(arg_0)))))
		})
	case *hxrt.Function[func(*sys__thread__ElasticThreadPoolWorker)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			hxrt.FunctionCall(source)(hx_recover_class_4(arg_0))
			return func(hx_value_197 any) bool {
				if hx_value_197 == nil {
					var hx_zero_198 bool
					return hx_zero_198
				}
				return hx_value_197.(bool)
			}(any(nil))
		})
	case *hxrt.Function[func(*hxrt.Function[func()])]:
		return hxrt.ProjectFunction(source, func(arg_0 any) bool {
			hxrt.FunctionCall(source)(hx_recover_callable_1(arg_0))
			return func(hx_value_199 any) bool {
				if hx_value_199 == nil {
					var hx_zero_200 bool
					return hx_zero_200
				}
				return hx_value_199.(bool)
			}(any(nil))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_3(value any) *hxrt.Function[func(any) any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any) any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			hxrt.FunctionCall(source)(arg_0)
			return nil
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(hx_recover_callable_2(arg_0))
		})
	case *hxrt.Function[func(*string) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_201 any) *string {
				if hx_value_201 == nil {
					var hx_zero_202 *string
					return hx_zero_202
				}
				return hx_value_201.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(hx_recover_callable_3(arg_0))
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(bool) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_203 any) bool {
				if hx_value_203 == nil {
					var hx_zero_204 bool
					return hx_zero_204
				}
				return hx_value_203.(bool)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*sys__thread__ElasticThreadPoolWorker)]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			hxrt.FunctionCall(source)(hx_recover_class_4(arg_0))
			return nil
		})
	case *hxrt.Function[func(*hxrt.Function[func()])]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			hxrt.FunctionCall(source)(hx_recover_callable_1(arg_0))
			return nil
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *sys__thread__NoEventLoopException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__thread__NoEventLoopException:
		return source
	default:
		return value.(*sys__thread__NoEventLoopException)
	}
}

func hx_recover_class_1(value any) *sys__thread__ThreadPoolException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__thread__ThreadPoolException:
		return source
	default:
		return value.(*sys__thread__ThreadPoolException)
	}
}

func hx_recover_class_2(value any) *sys__thread__FixedThreadPoolShutdownException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__thread__FixedThreadPoolShutdownException:
		return source
	default:
		return value.(*sys__thread__FixedThreadPoolShutdownException)
	}
}

func hx_recover_class_3(value any) *sys__thread__FixedThreadPoolWorker {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__thread__FixedThreadPoolWorker:
		return source
	default:
		return value.(*sys__thread__FixedThreadPoolWorker)
	}
}

func hx_recover_class_4(value any) *sys__thread__ElasticThreadPoolWorker {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__thread__ElasticThreadPoolWorker:
		return source
	default:
		return value.(*sys__thread__ElasticThreadPoolWorker)
	}
}

func hx_recover_class_5(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_6(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__thread__ThreadPoolException:
		return source.haxe__Exception
	case *sys__thread__NoEventLoopException:
		return source.haxe__Exception
	case *sys__thread__FixedThreadPoolShutdownException:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_7(value any) *haxe__ds___List__GoListIterator {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds___List__GoListIterator:
		return source
	default:
		return value.(*haxe__ds___List__GoListIterator)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__thread__NoEventLoopException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__thread__ThreadPoolException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_2(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__thread__FixedThreadPoolShutdownException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_3(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_4(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__thread__ThreadPoolException:
		return (source != nil)
	case *sys__thread__NoEventLoopException:
		return (source != nil)
	case *sys__thread__FixedThreadPoolShutdownException:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}
