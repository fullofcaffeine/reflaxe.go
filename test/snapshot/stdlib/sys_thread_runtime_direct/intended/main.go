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
	worker := sys__thread__Thread_create(func() {
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("worker-ready"))
		var payload any = sys__thread__Thread_readMessage(true)
		mainThread.__hx_this.sendMessage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("worker-echo="), hxrt.StdString(payload)))
	})
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("thread.msg1="), hxrt.StdString(sys__thread__Thread_readMessage(true))))
	hxrt.Println(v)
	worker.__hx_this.sendMessage(hxrt.StringFromLiteral("ping"))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("thread.msg2="), hxrt.StdString(sys__thread__Thread_readMessage(true))))
	hxrt.Println(v_1)
	hxrt.TryCatch(func() {
		worker.__hx_this.get_events().__hx_this.progress()
		hxrt.Println(any(hxrt.StringFromLiteral("thread.worker_events=available")))
	}, func(hx_caught_1 any) {
		switch hx_typed_2 := hx_caught_1.(type) {
		case *sys__thread__NoEventLoopException:
			err := hx_typed_2
			var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("thread.worker_events="), hxrt.ExceptionMessage(err)))
			hxrt.Println(v_2)
		default:
			hxrt.Throw(hx_caught_1)
		}
	})
	runLoop := New_sys__thread__EventLoop()
	runLoop.__hx_this.run(func() {
		hxrt.Println(any(hxrt.StringFromLiteral("loop.run=once")))
	})
	runLoop.__hx_this.loop()
	hxrt.Println(any(hxrt.StringFromLiteral("loop.loop_after=done")))
	promisedLoop := New_sys__thread__EventLoop()
	promisedLoop.__hx_this.promise()
	promisedLoop.__hx_this.runPromised(func() {
		hxrt.Println(any(hxrt.StringFromLiteral("loop.runPromised=ok")))
	})
	promisedLoop.__hx_this.loop()
	hxrt.Println(any(hxrt.StringFromLiteral("loop.promised_after=done")))
	loop := New_sys__thread__EventLoop()
	repeats := 0
	var handler any = any(0)
	handler = loop.__hx_this.repeat(func() {
		repeats = int(int32((repeats + 1)))
		hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("loop.repeat="), repeats)))
		loop.__hx_this.cancel(handler)
	}, 10)
	loop.__hx_this.loop()
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("loop.repeat_after="), repeats)))
	sys__thread__Thread_create(func() {
		sys__thread__Thread_runWithEventLoop(func() {
			sys__thread__Thread_current().__hx_this.get_events().__hx_this.promise()
			sys__thread__Thread_current().__hx_this.get_events().__hx_this.runPromised(func() {
				mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("runWithEventLoop=ok"))
			})
		})
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("runWithEventLoop=after"))
	})
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("thread.runWithEventLoop="), sortedMessages(2)))
	hxrt.Println(v_3)
	sys__thread__Thread_createWithEventLoop(func() {
		sys__thread__Thread_current().__hx_this.get_events().__hx_this.promise()
		sys__thread__Thread_current().__hx_this.get_events().__hx_this.runPromised(func() {
			mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("createWithEventLoop=ok"))
		})
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("createWithEventLoop=after-job"))
	})
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("thread.createWithEventLoop="), sortedMessages(2)))
	hxrt.Println(v_4)
	fixed := New_sys__thread__FixedThreadPool(2)
	fixed.__hx_this.run(func() {
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("fixed:b"))
	})
	fixed.__hx_this.run(func() {
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("fixed:a"))
	})
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fixed.msgs="), sortedMessages(2)))
	hxrt.Println(v_5)
	fixed.__hx_this.shutdown()
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fixed.shutdown="), hxrt.StdString(fixed.__hx_this.get_isShutdown())))
	hxrt.Println(v_6)
	hxrt.TryCatch(func() {
		fixed.__hx_this.run(func() {
		})
	}, func(hx_caught_3 any) {
		switch hx_typed_4 := hx_caught_3.(type) {
		case *sys__thread__ThreadPoolException:
			err_1 := hx_typed_4
			var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fixed.error="), hxrt.ExceptionMessage(err_1)))
			hxrt.Println(v_7)
		default:
			hxrt.Throw(hx_caught_3)
		}
	})
	elastic := New_sys__thread__ElasticThreadPool(2, 0.05)
	elastic.__hx_this.run(func() {
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("elastic:a"))
	})
	elastic.__hx_this.run(func() {
		mainThread.__hx_this.sendMessage(hxrt.StringFromLiteral("elastic:b"))
	})
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("elastic.msgs="), sortedMessages(2)))
	hxrt.Println(v_8)
	elastic.__hx_this.shutdown()
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("elastic.shutdown="), hxrt.StdString(elastic.__hx_this.get_isShutdown())))
	hxrt.Println(v_9)
	hxrt.TryCatch(func() {
		elastic.__hx_this.run(func() {
		})
	}, func(hx_caught_5 any) {
		switch hx_typed_6 := hx_caught_5.(type) {
		case *sys__thread__ThreadPoolException:
			err_2 := hx_typed_6
			var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("elastic.error="), hxrt.ExceptionMessage(err_2)))
			hxrt.Println(v_10)
		default:
			hxrt.Throw(hx_caught_5)
		}
	})
}

func sortedMessages(count int) *string {
	values := hxrt.NewArray()
	_g := 0
	_g1 := count
	for _g < _g1 {
		hx_post_7 := _g
		_g = int(int32((_g + 1)))
		hx_tmp := hx_post_7
		_ = hx_tmp
		values.Push(hxrt.StdString(sys__thread__Thread_readMessage(true)))
	}
	haxe__ds__ArraySort_sort(values, func(hx_cmp_left_9 any, hx_cmp_right_10 any) int {
		return Reflect_compare(func(hx_value_11 any) *string {
			if hx_value_11 == nil {
				var hx_zero_12 *string
				return hx_zero_12
			}
			return hx_value_11.(*string)
		}(hx_cmp_left_9), func(hx_value_13 any) *string {
			if hx_value_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_value_13.(*string)
		}(hx_cmp_right_10))
	})
	var buf_b *string
	buf_b = hxrt.StringFromLiteral("")
	var _g_current int
	var _g_array *hxrt.Array
	_g_current = 0
	_g_array = values
	for _g_current < _g_array.Len() {
		var _g_value *string
		var _g_key int
		_g_value = func(hx_value_15 any) *string {
			if hx_value_15 == nil {
				var hx_zero_16 *string
				return hx_zero_16
			}
			return hx_value_15.(*string)
		}(_g_array.Get(_g_current))
		hx_post_17 := _g_current
		_g_current = int(int32((_g_current + 1)))
		_g_key = hx_post_17
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
	switch value := obj.(type) {
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
	case *sys__thread__Thread:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
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
	case *sys__thread__FixedThreadPoolWorker:
		return hxrt__generated_method_field__sys__thread__FixedThreadPoolWorker(value, key)
	case *sys__thread__Lock:
		return hxrt__generated_method_field__sys__thread__Lock(value, key)
	case *sys__thread__Mutex:
		return hxrt__generated_method_field__sys__thread__Mutex(value, key)
	case *sys__thread__Thread:
		return hxrt__generated_method_field__sys__thread__Thread(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__haxe__ds__List(value *haxe__ds__List, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "add":
		return value.add
	case "clear":
		return value.clear
	case "filter":
		return value.filter
	case "first":
		return value.first
	case "isEmpty":
		return value.isEmpty
	case "iterator":
		return value.iterator
	case "join":
		return value.join
	case "keyValueIterator":
		return value.keyValueIterator
	case "last":
		return value.last
	case "map":
		return value.map_
	case "pop":
		return value.pop
	case "push":
		return value.push
	case "remove":
		return value.remove
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds___List__GoListIterator(value *haxe__ds___List__GoListIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return value.hasNext
	case "next":
		return value.next
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds___List__GoListKeyValueIterator(value *haxe__ds___List__GoListKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return value.hasNext
	case "next":
		return value.next
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__Deque(value *sys__thread__Deque, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "add":
		return value.add
	case "pop":
		return value.pop
	case "push":
		return value.push
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__ElasticThreadPool(value *sys__thread__ElasticThreadPool, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "get_isShutdown":
		return value.get_isShutdown
	case "get_threadsCount":
		return value.get_threadsCount
	case "retireWorkerLocked":
		return value.retireWorkerLocked
	case "run":
		return value.run
	case "shutdown":
		return value.shutdown
	case "startWorkerLocked":
		return value.startWorkerLocked
	case "workerResolveWait":
		return value.workerResolveWait
	case "workerTaskFailed":
		return value.workerTaskFailed
	case "workerTaskFinished":
		return value.workerTaskFinished
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__ElasticThreadPoolWorker(value *sys__thread__ElasticThreadPoolWorker, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "loop":
		return value.loop
	case "start":
		return value.start
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__EventLoop(value *sys__thread__EventLoop, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "cancel":
		return value.cancel
	case "loop":
		return value.loop
	case "progress":
		return value.progress
	case "promise":
		return value.promise
	case "repeat":
		return value.repeat
	case "run":
		return value.run
	case "runPromised":
		return value.runPromised
	case "wait":
		return value.wait
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__FixedThreadPool(value *sys__thread__FixedThreadPool, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "get_isShutdown":
		return value.get_isShutdown
	case "get_threadsCount":
		return value.get_threadsCount
	case "run":
		return value.run
	case "shutdown":
		return value.shutdown
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__FixedThreadPoolWorker(value *sys__thread__FixedThreadPoolWorker, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "loop":
		return value.loop
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__Lock(value *sys__thread__Lock, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "release":
		return value.release
	case "wait":
		return value.wait
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__Mutex(value *sys__thread__Mutex, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "acquire":
		return value.acquire
	case "release":
		return value.release
	case "tryAcquire":
		return value.tryAcquire
	}
	return nil
}

func hxrt__generated_method_field__sys__thread__Thread(value *sys__thread__Thread, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "get_events":
		return value.get_events
	case "sendMessage":
		return value.sendMessage
	}
	return nil
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
	switch value := object.(type) {
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
	case *sys__thread__Thread:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
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
	case *sys__thread__FixedThreadPoolWorker:
		return hxrt__generated_field_lookup__sys__thread__FixedThreadPoolWorker(value, key)
	case *sys__thread__Lock:
		return hxrt__generated_field_lookup__sys__thread__Lock(value, key)
	case *sys__thread__Mutex:
		return hxrt__generated_field_lookup__sys__thread__Mutex(value, key)
	case *sys__thread__Thread:
		return hxrt__generated_field_lookup__sys__thread__Thread(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
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
	case *sys__thread__Thread:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
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
	case *sys__thread__FixedThreadPoolWorker:
		return hxrt__generated_field_has__sys__thread__FixedThreadPoolWorker(value, key)
	case *sys__thread__Lock:
		return hxrt__generated_field_has__sys__thread__Lock(value, key)
	case *sys__thread__Mutex:
		return hxrt__generated_field_has__sys__thread__Mutex(value, key)
	case *sys__thread__Thread:
		return hxrt__generated_field_has__sys__thread__Thread(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
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
	case *sys__thread__Thread:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
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
	case *sys__thread__FixedThreadPoolWorker:
		return hxrt__generated_field_set__sys__thread__FixedThreadPoolWorker(value, key, incoming)
	case *sys__thread__Lock:
		return hxrt__generated_field_set__sys__thread__Lock(value, key, incoming)
	case *sys__thread__Mutex:
		return hxrt__generated_field_set__sys__thread__Mutex(value, key, incoming)
	case *sys__thread__Thread:
		return hxrt__generated_field_set__sys__thread__Thread(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := object.(type) {
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
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__ds__List:
		return hxrt.NewArray(hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("length"))
	case *haxe__ds___List__GoListIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"))
	case *haxe__ds___List__GoListKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"))
	case *sys__thread__Deque:
		return hxrt.NewArray(hxrt.StringFromLiteral("__available"), hxrt.StringFromLiteral("__items"), hxrt.StringFromLiteral("__mutex"))
	case *sys__thread__ElasticThreadPool:
		return hxrt.NewArray(hxrt.StringFromLiteral("_isShutdown"), hxrt.StringFromLiteral("available"), hxrt.StringFromLiteral("liveWorkers"), hxrt.StringFromLiteral("maxThreadsCount"), hxrt.StringFromLiteral("mutex"), hxrt.StringFromLiteral("pendingTasks"), hxrt.StringFromLiteral("pool"), hxrt.StringFromLiteral("queue"), hxrt.StringFromLiteral("threadTimeout"))
	case *sys__thread__ElasticThreadPoolWorker:
		return hxrt.NewArray(hxrt.StringFromLiteral("available"), hxrt.StringFromLiteral("dead"), hxrt.StringFromLiteral("owner"), hxrt.StringFromLiteral("task"), hxrt.StringFromLiteral("timeout"))
	case *sys__thread__EventLoop:
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"))
	case *sys__thread__FixedThreadPool:
		return hxrt.NewArray(hxrt.StringFromLiteral("_isShutdown"), hxrt.StringFromLiteral("mutex"), hxrt.StringFromLiteral("pool"), hxrt.StringFromLiteral("queue"))
	case *sys__thread__FixedThreadPoolShutdownException:
		return hxrt.NewArray()
	case *sys__thread__FixedThreadPoolWorker:
		return hxrt.NewArray(hxrt.StringFromLiteral("queue"))
	case *sys__thread__Lock:
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"))
	case *sys__thread__Mutex:
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"))
	case *sys__thread__NoEventLoopException:
		return hxrt.NewArray()
	case *sys__thread__Thread:
		return hxrt.NewArray(hxrt.StringFromLiteral("__id"))
	case *sys__thread__ThreadPoolException:
		return hxrt.NewArray()
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return value.high
	case "low":
		return value.low
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
		return value.items
	case "length":
		return value.length
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
		return value.index
	case "items":
		return value.items
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
		return value.index
	case "items":
		return value.items
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
		return value.__available
	case "__items":
		return value.__items
	case "__mutex":
		return value.__mutex
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
		return value._isShutdown
	case "available":
		return value.available
	case "liveWorkers":
		return value.liveWorkers
	case "maxThreadsCount":
		return value.maxThreadsCount
	case "mutex":
		return value.mutex
	case "pendingTasks":
		return value.pendingTasks
	case "pool":
		return value.pool
	case "queue":
		return value.queue
	case "threadTimeout":
		return value.threadTimeout
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
		return value.available
	case "dead":
		return value.dead
	case "owner":
		return value.owner
	case "task":
		return value.task
	case "timeout":
		return value.timeout
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
		if incoming == nil {
			var zero func()
			value.task = zero
			return true
		}
		switch typed := incoming.(type) {
		case func():
			value.task = typed
			return true
		default:
			return false
		}
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
		return value.__h
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
		return value._isShutdown
	case "mutex":
		return value.mutex
	case "pool":
		return value.pool
	case "queue":
		return value.queue
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

func hxrt__generated_field_lookup__sys__thread__FixedThreadPoolWorker(value *sys__thread__FixedThreadPoolWorker, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "queue":
		return value.queue
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
		return value.__h
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
		return value.__h
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

func hxrt__generated_field_lookup__sys__thread__Thread(value *sys__thread__Thread, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__id":
		return value.__id
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
	case *sys__thread__NextEventTime:
		return (enumValue != nil)
	default:
		return false
	}
}
