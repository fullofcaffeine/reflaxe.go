package main

import (
	"reflect"
	"snapshot/hxrt"
)

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

type Std struct {
}

type Type struct {
}

type Reflect struct {
}

func Reflect_compare(a any, b any) int {
	toFloat := func(value any) (float64, bool) {
		switch v := value.(type) {
		case int:
			return float64(v), true
		case int8:
			return float64(v), true
		case int16:
			return float64(v), true
		case int32:
			return float64(v), true
		case int64:
			return float64(v), true
		case uint:
			return float64(v), true
		case uint8:
			return float64(v), true
		case uint16:
			return float64(v), true
		case uint32:
			return float64(v), true
		case uint64:
			return float64(v), true
		case float32:
			return float64(v), true
		case float64:
			return v, true
		default:
			return 0, false
		}
	}
	if af, ok := toFloat(a); ok {
		if bf, okB := toFloat(b); okB {
			if af < bf {
				return -1
			}
			if af > bf {
				return 1
			}
			return 0
		}
	}
	aStr := *hxrt.StdString(a)
	bStr := *hxrt.StdString(b)
	if aStr < bStr {
		return -1
	}
	if aStr > bStr {
		return 1
	}
	return 0
}

func Reflect_compareMethods(a any, b any) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	if !av.IsValid() || !bv.IsValid() {
		return !av.IsValid() && !bv.IsValid()
	}
	if av.Kind() == reflect.Func && bv.Kind() == reflect.Func {
		if av.IsNil() || bv.IsNil() {
			return av.IsNil() && bv.IsNil()
		}
		return av.Pointer() == bv.Pointer()
	}
	return reflect.DeepEqual(a, b)
}

func Reflect_field(obj any, field *string) any {
	if obj == nil {
		return nil
	}
	key := *hxrt.StdString(field)
	if metadataValue, ok := hxrt_typeClassMetadataField(obj, key); ok {
		return metadataValue
	}
	switch value := obj.(type) {
	case map[string]any:
		return value[key]
	case map[any]any:
		return value[key]
	case *map[string]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	case *map[any]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return nil
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if fieldValue := rv.FieldByName(key); fieldValue.IsValid() && fieldValue.CanInterface() {
			return fieldValue.Interface()
		}
	}
	method := reflect.ValueOf(obj).MethodByName(key)
	if method.IsValid() {
		return method.Interface()
	}
	return nil
}

func Reflect_hasField(obj any, field *string) bool {
	if obj == nil {
		return false
	}
	key := *hxrt.StdString(field)
	if _, ok := hxrt_typeClassMetadataField(obj, key); ok {
		return true
	}
	switch value := obj.(type) {
	case map[string]any:
		_, ok := value[key]
		return ok
	case map[any]any:
		_, ok := value[key]
		return ok
	case *map[string]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	case *map[any]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return false
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if rv.FieldByName(key).IsValid() {
			return true
		}
	}
	return reflect.ValueOf(obj).MethodByName(key).IsValid()
}

func Reflect_setField(obj any, field *string, value any) {
	if obj == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	key := *hxrt.StdString(field)
	switch target := obj.(type) {
	case map[string]any:
		target[key] = value
		return
	case map[any]any:
		target[key] = value
		return
	case *map[string]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	case *map[any]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() || rv.Kind() != reflect.Pointer {
		return
	}
	if rv.IsNil() {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return
	}
	fieldValue := rv.FieldByName(key)
	if !fieldValue.IsValid() || !fieldValue.CanSet() {
		return
	}
	if value == nil {
		fieldValue.Set(reflect.Zero(fieldValue.Type()))
		return
	}
	incoming := reflect.ValueOf(value)
	if incoming.Type().AssignableTo(fieldValue.Type()) {
		fieldValue.Set(incoming)
		return
	}
	if incoming.Type().ConvertibleTo(fieldValue.Type()) {
		fieldValue.Set(incoming.Convert(fieldValue.Type()))
		return
	}
	if fieldValue.Kind() == reflect.Interface {
		fieldValue.Set(incoming)
	}
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
}

type ValueType struct {
	tag    int
	params []any
}

var ValueType_TNull *ValueType = &ValueType{tag: 0, params: []any{}}

var ValueType_TInt *ValueType = &ValueType{tag: 1, params: []any{}}

var ValueType_TFloat *ValueType = &ValueType{tag: 2, params: []any{}}

var ValueType_TBool *ValueType = &ValueType{tag: 3, params: []any{}}

var ValueType_TObject *ValueType = &ValueType{tag: 4, params: []any{}}

var ValueType_TFunction *ValueType = &ValueType{tag: 5, params: []any{}}

var ValueType_TUnknown *ValueType = &ValueType{tag: 8, params: []any{}}

func ValueType_TClass(c any) *ValueType {
	return &ValueType{tag: 6, params: []any{c}}
}

func ValueType_TEnum(e any) *ValueType {
	return &ValueType{tag: 7, params: []any{e}}
}

func hxrt_typeCallAny(callable any, args []any) (any, bool) {
	result := any(nil)
	ok := false
	defer func() {
		if recover() != nil {
			result = nil
			ok = false
		}
	}()
	if callable == nil {
		return nil, false
	}
	fn := reflect.ValueOf(callable)
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil, false
	}
	fnType := fn.Type()
	if fnType.NumIn() != len(args) {
		return nil, false
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		paramType := fnType.In(i)
		arg := args[i]
		if arg == nil {
			in[i] = reflect.Zero(paramType)
			continue
		}
		v := reflect.ValueOf(arg)
		if v.IsValid() && v.Type().AssignableTo(paramType) {
			in[i] = v
			continue
		}
		if v.IsValid() && v.Type().ConvertibleTo(paramType) {
			in[i] = v.Convert(paramType)
			continue
		}
		if paramType.Kind() == reflect.Interface && v.IsValid() {
			in[i] = v
			continue
		}
		return nil, false
	}
	out := fn.Call(in)
	if len(out) == 0 {
		return nil, true
	}
	first := out[0]
	if !first.IsValid() {
		return nil, true
	}
	result = first.Interface()
	ok = true
	return result, ok
}

func hxrt_typeArrayValues(value *hxrt.Array) []any {
	if value == nil {
		return []any{}
	}
	return value.Values()
}

func hxrt_typeResolvedClassName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeClassValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeClassValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeResolvedEnumName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeEnumValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeEnumValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeCreateClassInstance(className string, args []any) (any, bool) {
	switch className {
	case "Main":
		return nil, false
	case "haxe.Int64Helper":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe.ds.ArraySort":
		return nil, false
	case "haxe.ds.List":
		return hxrt_typeCallAny(New_haxe__ds__List, args)
	case "haxe.ds._List.GoListIterator":
		return hxrt_typeCallAny(New_haxe__ds___List__GoListIterator, args)
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__ds___List__GoListKeyValueIterator, args)
	case "sys.thread.Deque":
		return hxrt_typeCallAny(New_sys__thread__Deque, args)
	case "sys.thread.ElasticThreadPool":
		return hxrt_typeCallAny(New_sys__thread__ElasticThreadPool, args)
	case "sys.thread.ElasticThreadPoolWorker":
		return hxrt_typeCallAny(New_sys__thread__ElasticThreadPoolWorker, args)
	case "sys.thread.EventLoop":
		return hxrt_typeCallAny(New_sys__thread__EventLoop, args)
	case "sys.thread.FixedThreadPool":
		return hxrt_typeCallAny(New_sys__thread__FixedThreadPool, args)
	case "sys.thread.FixedThreadPoolShutdownException":
		return hxrt_typeCallAny(New_sys__thread__FixedThreadPoolShutdownException, args)
	case "sys.thread.FixedThreadPoolWorker":
		return hxrt_typeCallAny(New_sys__thread__FixedThreadPoolWorker, args)
	case "sys.thread.Lock":
		return hxrt_typeCallAny(New_sys__thread__Lock, args)
	case "sys.thread.Mutex":
		return hxrt_typeCallAny(New_sys__thread__Mutex, args)
	case "sys.thread.NoEventLoopException":
		return hxrt_typeCallAny(New_sys__thread__NoEventLoopException, args)
	case "sys.thread.Thread":
		return hxrt_typeCallAny(New_sys__thread__Thread, args)
	case "sys.thread.ThreadPoolException":
		return hxrt_typeCallAny(New_sys__thread__ThreadPoolException, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "haxe._Int64.___Int64":
		return &haxe___Int64_____Int64{}, true
	case "haxe.ds.List":
		return &haxe__ds__List{}, true
	case "haxe.ds._List.GoListIterator":
		return &haxe__ds___List__GoListIterator{}, true
	case "haxe.ds._List.GoListKeyValueIterator":
		return &haxe__ds___List__GoListKeyValueIterator{}, true
	case "sys.thread.Deque":
		return &sys__thread__Deque{}, true
	case "sys.thread.ElasticThreadPool":
		return &sys__thread__ElasticThreadPool{}, true
	case "sys.thread.ElasticThreadPoolWorker":
		return &sys__thread__ElasticThreadPoolWorker{}, true
	case "sys.thread.EventLoop":
		return &sys__thread__EventLoop{}, true
	case "sys.thread.FixedThreadPool":
		return &sys__thread__FixedThreadPool{}, true
	case "sys.thread.FixedThreadPoolShutdownException":
		return &sys__thread__FixedThreadPoolShutdownException{}, true
	case "sys.thread.FixedThreadPoolWorker":
		return &sys__thread__FixedThreadPoolWorker{}, true
	case "sys.thread.Lock":
		return &sys__thread__Lock{}, true
	case "sys.thread.Mutex":
		return &sys__thread__Mutex{}, true
	case "sys.thread.NoEventLoopException":
		return &sys__thread__NoEventLoopException{}, true
	case "sys.thread.Thread":
		return &sys__thread__Thread{}, true
	case "sys.thread.ThreadPoolException":
		return &sys__thread__ThreadPoolException{}, true
	default:
		return nil, false
	}
}

func hxrt_typeCreateEnumInstance(enumName string, constructorName string, constructorIndex int, useIndex bool, args []any) (any, bool) {
	switch enumName {
	case "ValueType":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TNull, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TInt, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFloat, true
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TBool, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TObject, true
			case 5:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFunction, true
			case 6:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TClass, args)
			case 7:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TEnum, args)
			case 8:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TUnknown, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "TNull":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TNull, true
		case "TInt":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TInt, true
		case "TFloat":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFloat, true
		case "TBool":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TBool, true
		case "TObject":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TObject, true
		case "TFunction":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFunction, true
		case "TClass":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TClass, args)
		case "TEnum":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TEnum, args)
		case "TUnknown":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TUnknown, true
		default:
			return nil, false
		}
	case "sys.thread.NextEventTime":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return sys__thread__NextEventTime_Now, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return sys__thread__NextEventTime_Never, true
			case 2:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(sys__thread__NextEventTime_AnyTime, args)
			case 3:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(sys__thread__NextEventTime_At, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "Now":
			if len(args) != 0 {
				return nil, false
			}
			return sys__thread__NextEventTime_Now, true
		case "Never":
			if len(args) != 0 {
				return nil, false
			}
			return sys__thread__NextEventTime_Never, true
		case "AnyTime":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(sys__thread__NextEventTime_AnyTime, args)
		case "At":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(sys__thread__NextEventTime_At, args)
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

func Type_getClass(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeClassValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeClassValue:
		copyValue := value
		return &copyValue
	case *hxrt.Array:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	case *haxe__ds__List:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.List")}
	case *haxe__ds___List__GoListIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds._List.GoListIterator")}
	case *haxe__ds___List__GoListKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds._List.GoListKeyValueIterator")}
	case *sys__thread__Deque:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.Deque")}
	case *sys__thread__ElasticThreadPool:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.ElasticThreadPool")}
	case *sys__thread__ElasticThreadPoolWorker:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.ElasticThreadPoolWorker")}
	case *sys__thread__EventLoop:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.EventLoop")}
	case *sys__thread__FixedThreadPool:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.FixedThreadPool")}
	case *sys__thread__FixedThreadPoolShutdownException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.FixedThreadPoolShutdownException")}
	case *sys__thread__FixedThreadPoolWorker:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.FixedThreadPoolWorker")}
	case *sys__thread__Lock:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.Lock")}
	case *sys__thread__Mutex:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.Mutex")}
	case *sys__thread__NoEventLoopException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.NoEventLoopException")}
	case *sys__thread__Thread:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.Thread")}
	case *sys__thread__ThreadPoolException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.ThreadPoolException")}
	default:
		return nil
	}
}

func Type_getEnum(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeEnumValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeEnumValue:
		copyValue := value
		return &copyValue
	case *ValueType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("ValueType")}
	case *sys__thread__NextEventTime:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("sys.thread.NextEventTime")}
	default:
		return nil
	}
}

func Type_getSuperClass(c any) any {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	switch className {
	case "Main":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe.ds.ArraySort":
		return nil
	case "haxe.ds.List":
		return nil
	case "haxe.ds._List.GoListIterator":
		return nil
	case "haxe.ds._List.GoListKeyValueIterator":
		return nil
	case "sys.thread.Deque":
		return nil
	case "sys.thread.ElasticThreadPool":
		return nil
	case "sys.thread.ElasticThreadPoolWorker":
		return nil
	case "sys.thread.EventLoop":
		return nil
	case "sys.thread.FixedThreadPool":
		return nil
	case "sys.thread.FixedThreadPoolShutdownException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "sys.thread.FixedThreadPoolWorker":
		return nil
	case "sys.thread.Lock":
		return nil
	case "sys.thread.Mutex":
		return nil
	case "sys.thread.NoEventLoopException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "sys.thread.Thread":
		return nil
	case "sys.thread.ThreadPoolException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	default:
		return nil
	}
}

func Type_getClassName(c any) *string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(className)
}

func Type_getClassFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("main"), hxrt.StringFromLiteral("sortedMessages"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe.ds.ArraySort":
		return hxrt.NewArray(hxrt.StringFromLiteral("doMerge"), hxrt.StringFromLiteral("gcd"), hxrt.StringFromLiteral("lower"), hxrt.StringFromLiteral("rec"), hxrt.StringFromLiteral("rotate"), hxrt.StringFromLiteral("sort"), hxrt.StringFromLiteral("swap"), hxrt.StringFromLiteral("upper"))
	case "haxe.ds.List":
		return hxrt.NewArray(hxrt.StringFromLiteral("sameValue"))
	case "haxe.ds._List.GoListIterator":
		return hxrt.NewArray()
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt.NewArray()
	case "sys.thread.Deque":
		return hxrt.NewArray()
	case "sys.thread.ElasticThreadPool":
		return hxrt.NewArray()
	case "sys.thread.ElasticThreadPoolWorker":
		return hxrt.NewArray()
	case "sys.thread.EventLoop":
		return hxrt.NewArray(hxrt.StringFromLiteral("__fromHandle"))
	case "sys.thread.FixedThreadPool":
		return hxrt.NewArray(hxrt.StringFromLiteral("shutdownTask"))
	case "sys.thread.FixedThreadPoolShutdownException":
		return hxrt.NewArray()
	case "sys.thread.FixedThreadPoolWorker":
		return hxrt.NewArray()
	case "sys.thread.Lock":
		return hxrt.NewArray()
	case "sys.thread.Mutex":
		return hxrt.NewArray()
	case "sys.thread.NoEventLoopException":
		return hxrt.NewArray()
	case "sys.thread.Thread":
		return hxrt.NewArray(hxrt.StringFromLiteral("create"), hxrt.StringFromLiteral("createWithEventLoop"), hxrt.StringFromLiteral("current"), hxrt.StringFromLiteral("processEvents"), hxrt.StringFromLiteral("readMessage"), hxrt.StringFromLiteral("runWithEventLoop"))
	case "sys.thread.ThreadPoolException":
		return hxrt.NewArray()
	default:
		return hxrt.NewArray()
	}
}

func Type_getInstanceFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Main":
		return hxrt.NewArray()
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe.ds.ArraySort":
		return hxrt.NewArray()
	case "haxe.ds.List":
		return hxrt.NewArray(hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("filter"), hxrt.StringFromLiteral("first"), hxrt.StringFromLiteral("isEmpty"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("join"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("last"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("pop"), hxrt.StringFromLiteral("push"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds._List.GoListIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("next"))
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("next"))
	case "sys.thread.Deque":
		return hxrt.NewArray(hxrt.StringFromLiteral("__available"), hxrt.StringFromLiteral("__items"), hxrt.StringFromLiteral("__mutex"), hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("pop"), hxrt.StringFromLiteral("push"))
	case "sys.thread.ElasticThreadPool":
		return hxrt.NewArray(hxrt.StringFromLiteral("_isShutdown"), hxrt.StringFromLiteral("available"), hxrt.StringFromLiteral("get_isShutdown"), hxrt.StringFromLiteral("get_threadsCount"), hxrt.StringFromLiteral("isShutdown"), hxrt.StringFromLiteral("liveWorkers"), hxrt.StringFromLiteral("maxThreadsCount"), hxrt.StringFromLiteral("mutex"), hxrt.StringFromLiteral("pendingTasks"), hxrt.StringFromLiteral("pool"), hxrt.StringFromLiteral("queue"), hxrt.StringFromLiteral("retireWorkerLocked"), hxrt.StringFromLiteral("run"), hxrt.StringFromLiteral("shutdown"), hxrt.StringFromLiteral("startWorkerLocked"), hxrt.StringFromLiteral("threadTimeout"), hxrt.StringFromLiteral("threadsCount"), hxrt.StringFromLiteral("workerResolveWait"), hxrt.StringFromLiteral("workerTaskFailed"), hxrt.StringFromLiteral("workerTaskFinished"))
	case "sys.thread.ElasticThreadPoolWorker":
		return hxrt.NewArray(hxrt.StringFromLiteral("available"), hxrt.StringFromLiteral("dead"), hxrt.StringFromLiteral("loop"), hxrt.StringFromLiteral("owner"), hxrt.StringFromLiteral("start"), hxrt.StringFromLiteral("task"), hxrt.StringFromLiteral("timeout"))
	case "sys.thread.EventLoop":
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"), hxrt.StringFromLiteral("cancel"), hxrt.StringFromLiteral("loop"), hxrt.StringFromLiteral("progress"), hxrt.StringFromLiteral("promise"), hxrt.StringFromLiteral("repeat"), hxrt.StringFromLiteral("run"), hxrt.StringFromLiteral("runPromised"), hxrt.StringFromLiteral("wait"))
	case "sys.thread.FixedThreadPool":
		return hxrt.NewArray(hxrt.StringFromLiteral("_isShutdown"), hxrt.StringFromLiteral("get_isShutdown"), hxrt.StringFromLiteral("get_threadsCount"), hxrt.StringFromLiteral("isShutdown"), hxrt.StringFromLiteral("mutex"), hxrt.StringFromLiteral("pool"), hxrt.StringFromLiteral("queue"), hxrt.StringFromLiteral("run"), hxrt.StringFromLiteral("shutdown"), hxrt.StringFromLiteral("threadsCount"))
	case "sys.thread.FixedThreadPoolShutdownException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "sys.thread.FixedThreadPoolWorker":
		return hxrt.NewArray(hxrt.StringFromLiteral("loop"), hxrt.StringFromLiteral("queue"))
	case "sys.thread.Lock":
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"), hxrt.StringFromLiteral("release"), hxrt.StringFromLiteral("wait"))
	case "sys.thread.Mutex":
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"), hxrt.StringFromLiteral("acquire"), hxrt.StringFromLiteral("release"), hxrt.StringFromLiteral("tryAcquire"))
	case "sys.thread.NoEventLoopException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "sys.thread.Thread":
		return hxrt.NewArray(hxrt.StringFromLiteral("__id"), hxrt.StringFromLiteral("events"), hxrt.StringFromLiteral("get_events"), hxrt.StringFromLiteral("sendMessage"))
	case "sys.thread.ThreadPoolException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	default:
		return hxrt.NewArray()
	}
}

func Type_getEnumName(e any) *string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(enumName)
}

func Type_resolveClass(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.ArraySort":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.List":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._List.GoListIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._List.GoListKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.Deque":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.ElasticThreadPool":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.ElasticThreadPoolWorker":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.EventLoop":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.FixedThreadPool":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.FixedThreadPoolShutdownException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.FixedThreadPoolWorker":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.Lock":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.Mutex":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.NoEventLoopException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.Thread":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.ThreadPoolException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_resolveEnum(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "ValueType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.NextEventTime":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_createInstance(cl any, args *hxrt.Array) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassInstance(className, hxrt_typeArrayValues(args))
	if !ok {
		return nil
	}
	return instance
}

func Type_createEmptyInstance(cl any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassEmptyInstance(className)
	if !ok {
		return nil
	}
	return instance
}

func Type_createEnum(e any, constr *string, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	constructorName := ""
	if constr != nil {
		constructorName = *hxrt.StdString(constr)
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_createEnumIndex(e any, index int, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, "", index, true, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_enumConstructor(e any) *string {
	if hxrt.AnyEqualsNull(e) {
		return nil
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("TNull")
		case 1:
			return hxrt.StringFromLiteral("TInt")
		case 2:
			return hxrt.StringFromLiteral("TFloat")
		case 3:
			return hxrt.StringFromLiteral("TBool")
		case 4:
			return hxrt.StringFromLiteral("TObject")
		case 5:
			return hxrt.StringFromLiteral("TFunction")
		case 6:
			return hxrt.StringFromLiteral("TClass")
		case 7:
			return hxrt.StringFromLiteral("TEnum")
		case 8:
			return hxrt.StringFromLiteral("TUnknown")
		default:
			return nil
		}
	case *sys__thread__NextEventTime:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("Now")
		case 1:
			return hxrt.StringFromLiteral("Never")
		case 2:
			return hxrt.StringFromLiteral("AnyTime")
		case 3:
			return hxrt.StringFromLiteral("At")
		default:
			return nil
		}
	default:
		return nil
	}
}

func Type_enumIndex(e any) int {
	if hxrt.AnyEqualsNull(e) {
		return -1
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return -1
		}
		return value.tag
	case *sys__thread__NextEventTime:
		if value == nil {
			return -1
		}
		return value.tag
	default:
		return -1
	}
}

func Type_getEnumConstructs(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown"))
	case "sys.thread.NextEventTime":
		return hxrt.NewArray(hxrt.StringFromLiteral("Now"), hxrt.StringFromLiteral("Never"), hxrt.StringFromLiteral("AnyTime"), hxrt.StringFromLiteral("At"))
	default:
		return hxrt.NewArray()
	}
}

func Type_enumParameters(e any) *hxrt.Array {
	if hxrt.AnyEqualsNull(e) {
		return hxrt.NewArray()
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *sys__thread__NextEventTime:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	default:
		return hxrt.NewArray()
	}
}

func Type_allEnums(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown)
	case "sys.thread.NextEventTime":
		return hxrt.NewArray(sys__thread__NextEventTime_Now, sys__thread__NextEventTime_Never)
	default:
		return hxrt.NewArray()
	}
}

func Type_typeof(v any) *ValueType {
	if hxrt.AnyEqualsNull(v) {
		return ValueType_TNull
	}
	switch v.(type) {
	case *hxrt__TypeClassValue, hxrt__TypeClassValue, *hxrt__TypeEnumValue, hxrt__TypeEnumValue:
		return ValueType_TObject
	}
	if enumValue := Type_getEnum(v); enumValue != nil {
		return ValueType_TEnum(enumValue)
	}
	if classValue := Type_getClass(v); classValue != nil {
		return ValueType_TClass(classValue)
	}
	switch v.(type) {
	case bool:
		return ValueType_TBool
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return ValueType_TInt
	case float32, float64:
		return ValueType_TFloat
	case string, *string:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("String")})
	case *hxrt.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	}
	ref := reflect.ValueOf(v)
	if !ref.IsValid() {
		return ValueType_TNull
	}
	switch ref.Kind() {
	case reflect.Func:
		return ValueType_TFunction
	case reflect.Slice, reflect.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	case reflect.Map, reflect.Struct, reflect.Interface, reflect.Pointer:
		return ValueType_TObject
	default:
		return ValueType_TUnknown
	}
}

func Type_enumEq(a any, b any) bool {
	return reflect.DeepEqual(a, b)
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	className, ok := hxrt_typeResolvedClassName(value)
	if !ok {
		return nil, false
	}
	switch className {
	default:
		return nil, false
	}
}
