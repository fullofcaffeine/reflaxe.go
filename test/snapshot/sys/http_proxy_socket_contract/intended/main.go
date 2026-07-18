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
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy0="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v)
	sys__Http_PROXY = func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["host"] = hxrt.StringFromLiteral("proxy.local")
		hx_obj_1["port"] = 3128
		hx_obj_2 := map[string]any{}
		hx_obj_2["user"] = hxrt.StringFromLiteral("scott")
		hx_obj_2["pass"] = hxrt.StringFromLiteral("tiger")
		hx_obj_1["auth"] = hx_obj_2
		return hx_obj_1
	}()
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy1="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_1)
	sys__Http_PROXY = func() map[string]any {
		hx_obj_3 := map[string]any{}
		hx_obj_3["host"] = hxrt.StringFromLiteral("proxy.local")
		hx_obj_3["port"] = 80
		hx_obj_4 := map[string]any{}
		hx_obj_4["user"] = hxrt.StringFromLiteral("scott")
		hx_obj_4["pass"] = nil
		hx_obj_3["auth"] = hx_obj_4
		return hx_obj_3
	}()
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy2="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_2)
	sys__Http_PROXY = func() map[string]any {
		hx_obj_5 := map[string]any{}
		hx_obj_5["host"] = hxrt.StringFromLiteral("proxy.local:9000")
		hx_obj_5["port"] = 3128
		hx_obj_5["auth"] = nil
		return hx_obj_5
	}()
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy3="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_3)
	sys__Http_PROXY = func() map[string]any {
		hx_obj_6 := map[string]any{}
		hx_obj_6["host"] = nil
		hx_obj_6["port"] = 3128
		hx_obj_6["auth"] = nil
		return hx_obj_6
	}()
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy4="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_4)
	sys__Http_PROXY = nil
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,body"))
	sink := New_haxe__io__BytesOutput()
	http.__hx_this.customRequest(false, sink.haxe__io__Output, New_sys__net__Socket(), hxrt.StringFromLiteral("PATCH"))
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("methodSock="), sink.__hx_this.getBytes().__hx_this.toString()))
	hxrt.Println(v_5)
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
	case "StringBuf":
		return nil, false
	case "StringTools":
		return nil, false
	case "haxe.Int64Helper":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe.ds.StringMap":
		return hxrt_typeCallAny(New_haxe__ds__StringMap, args)
	case "haxe.exceptions.NotImplementedException":
		return hxrt_typeCallAny(New_haxe__exceptions__NotImplementedException, args)
	case "haxe.exceptions.PosException":
		return hxrt_typeCallAny(New_haxe__exceptions__PosException, args)
	case "haxe.http.HttpBase":
		return hxrt_typeCallAny(New_haxe__http__HttpBase, args)
	case "haxe.io.Bytes":
		return hxrt_typeCallAny(New_haxe__io__Bytes, args)
	case "haxe.io.BytesBuffer":
		return hxrt_typeCallAny(New_haxe__io__BytesBuffer, args)
	case "haxe.io.BytesOutput":
		return hxrt_typeCallAny(New_haxe__io__BytesOutput, args)
	case "haxe.io.Eof":
		return hxrt_typeCallAny(New_haxe__io__Eof, args)
	case "haxe.io.FPHelper":
		return nil, false
	case "haxe.io.Input":
		return hxrt_typeCallAny(New_haxe__io__Input, args)
	case "haxe.io.Output":
		return hxrt_typeCallAny(New_haxe__io__Output, args)
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__MapKeyValueIterator, args)
	case "haxe.iterators.StringIterator":
		return hxrt_typeCallAny(New_haxe__iterators__StringIterator, args)
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__StringKeyValueIterator, args)
	case "sys.Http":
		return hxrt_typeCallAny(New_sys__Http, args)
	case "sys.net.Host":
		return hxrt_typeCallAny(New_sys__net__Host, args)
	case "sys.net.Socket":
		return hxrt_typeCallAny(New_sys__net__Socket, args)
	case "sys.net.SocketInput":
		return hxrt_typeCallAny(New_sys__net__SocketInput, args)
	case "sys.net.SocketOutput":
		return hxrt_typeCallAny(New_sys__net__SocketOutput, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "haxe._Int64.___Int64":
		return &haxe___Int64_____Int64{}, true
	case "haxe.ds.StringMap":
		return &haxe__ds__StringMap{}, true
	case "haxe.exceptions.NotImplementedException":
		return &haxe__exceptions__NotImplementedException{}, true
	case "haxe.exceptions.PosException":
		return &haxe__exceptions__PosException{}, true
	case "haxe.http.HttpBase":
		return &haxe__http__HttpBase{}, true
	case "haxe.io.Bytes":
		return &haxe__io__Bytes{}, true
	case "haxe.io.BytesBuffer":
		return &haxe__io__BytesBuffer{}, true
	case "haxe.io.BytesOutput":
		return &haxe__io__BytesOutput{}, true
	case "haxe.io.Eof":
		return &haxe__io__Eof{}, true
	case "haxe.io.Input":
		return &haxe__io__Input{}, true
	case "haxe.io.Output":
		return &haxe__io__Output{}, true
	case "haxe.iterators.MapKeyValueIterator":
		return &haxe__iterators__MapKeyValueIterator{}, true
	case "haxe.iterators.StringIterator":
		return &haxe__iterators__StringIterator{}, true
	case "haxe.iterators.StringKeyValueIterator":
		return &haxe__iterators__StringKeyValueIterator{}, true
	case "sys.Http":
		return &sys__Http{}, true
	case "sys.net.Host":
		return &sys__net__Host{}, true
	case "sys.net.Socket":
		return &sys__net__Socket{}, true
	case "sys.net.SocketInput":
		return &sys__net__SocketInput{}, true
	case "sys.net.SocketOutput":
		return &sys__net__SocketOutput{}, true
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
	case "haxe.io.Encoding":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Encoding_UTF8, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Encoding_RawNative, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "UTF8":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Encoding_UTF8, true
		case "RawNative":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Encoding_RawNative, true
		default:
			return nil, false
		}
	case "haxe.io.Error":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_Blocked, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_Overflow, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_OutsideBounds, true
			case 3:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__io__Error_Custom, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "Blocked":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_Blocked, true
		case "Overflow":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_Overflow, true
		case "OutsideBounds":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_OutsideBounds, true
		case "Custom":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__io__Error_Custom, args)
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
	case *haxe__ds__StringMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.StringMap")}
	case *haxe__exceptions__NotImplementedException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.NotImplementedException")}
	case *haxe__exceptions__PosException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.PosException")}
	case *haxe__http__HttpBase:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.http.HttpBase")}
	case *haxe__io__Bytes:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Bytes")}
	case *haxe__io__BytesBuffer:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.BytesBuffer")}
	case *haxe__io__BytesOutput:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.BytesOutput")}
	case *haxe__io__Eof:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Eof")}
	case *haxe__io__Input:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Input")}
	case *haxe__io__Output:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Output")}
	case *haxe__iterators__MapKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.MapKeyValueIterator")}
	case *haxe__iterators__StringIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.StringIterator")}
	case *haxe__iterators__StringKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.StringKeyValueIterator")}
	case *sys__Http:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.Http")}
	case *sys__net__Host:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.Host")}
	case *sys__net__Socket:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.Socket")}
	case *sys__net__SocketInput:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.SocketInput")}
	case *sys__net__SocketOutput:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.SocketOutput")}
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
	case *haxe__io__Encoding:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.io.Encoding")}
	case *haxe__io__Error:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.io.Error")}
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
	case "StringBuf":
		return nil
	case "StringTools":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe.ds.StringMap":
		return nil
	case "haxe.exceptions.NotImplementedException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.PosException")}
	case "haxe.exceptions.PosException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "haxe.http.HttpBase":
		return nil
	case "haxe.io.Bytes":
		return nil
	case "haxe.io.BytesBuffer":
		return nil
	case "haxe.io.BytesOutput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Output")}
	case "haxe.io.Eof":
		return nil
	case "haxe.io.FPHelper":
		return nil
	case "haxe.io.Input":
		return nil
	case "haxe.io.Output":
		return nil
	case "haxe.iterators.MapKeyValueIterator":
		return nil
	case "haxe.iterators.StringIterator":
		return nil
	case "haxe.iterators.StringKeyValueIterator":
		return nil
	case "sys.Http":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.http.HttpBase")}
	case "sys.net.Host":
		return nil
	case "sys.net.Socket":
		return nil
	case "sys.net.SocketInput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Input")}
	case "sys.net.SocketOutput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Output")}
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
		return hxrt.NewArray(hxrt.StringFromLiteral("main"))
	case "StringBuf":
		return hxrt.NewArray()
	case "StringTools":
		return hxrt.NewArray(hxrt.StringFromLiteral("MAX_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("contains"), hxrt.StringFromLiteral("containsImpl"), hxrt.StringFromLiteral("endsWith"), hxrt.StringFromLiteral("endsWithImpl"), hxrt.StringFromLiteral("fastCodeAt"), hxrt.StringFromLiteral("hex"), hxrt.StringFromLiteral("hexDigitValue"), hxrt.StringFromLiteral("htmlEscape"), hxrt.StringFromLiteral("htmlUnescape"), hxrt.StringFromLiteral("isEof"), hxrt.StringFromLiteral("isSpace"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("lpad"), hxrt.StringFromLiteral("ltrim"), hxrt.StringFromLiteral("replace"), hxrt.StringFromLiteral("rpad"), hxrt.StringFromLiteral("rtrim"), hxrt.StringFromLiteral("startsWith"), hxrt.StringFromLiteral("startsWithImpl"), hxrt.StringFromLiteral("trim"), hxrt.StringFromLiteral("unsafeCodeAt"), hxrt.StringFromLiteral("urlDecode"), hxrt.StringFromLiteral("urlEncode"), hxrt.StringFromLiteral("utf16CodePointAt"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe.ds.StringMap":
		return hxrt.NewArray()
	case "haxe.exceptions.NotImplementedException":
		return hxrt.NewArray()
	case "haxe.exceptions.PosException":
		return hxrt.NewArray()
	case "haxe.http.HttpBase":
		return hxrt.NewArray()
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_fromNativeView"), hxrt.StringFromLiteral("alloc"), hxrt.StringFromLiteral("fastGet"), hxrt.StringFromLiteral("ofData"), hxrt.StringFromLiteral("ofHex"), hxrt.StringFromLiteral("ofString"), hxrt.StringFromLiteral("rawNativeUsesUtf16LE"))
	case "haxe.io.BytesBuffer":
		return hxrt.NewArray()
	case "haxe.io.BytesOutput":
		return hxrt.NewArray()
	case "haxe.io.Eof":
		return hxrt.NewArray()
	case "haxe.io.FPHelper":
		return hxrt.NewArray(hxrt.StringFromLiteral("doubleToI64"), hxrt.StringFromLiteral("floatToI32"), hxrt.StringFromLiteral("i32ToFloat"), hxrt.StringFromLiteral("i64ToDouble"))
	case "haxe.io.Input":
		return hxrt.NewArray()
	case "haxe.io.Output":
		return hxrt.NewArray()
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray()
	case "haxe.iterators.StringIterator":
		return hxrt.NewArray()
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt.NewArray()
	case "sys.Http":
		return hxrt.NewArray(hxrt.StringFromLiteral("PROXY"), hxrt.StringFromLiteral("capture"), hxrt.StringFromLiteral("firstComma"), hxrt.StringFromLiteral("hxrt_proxyDescriptor"), hxrt.StringFromLiteral("normalizedMethod"), hxrt.StringFromLiteral("requestUrl"))
	case "sys.net.Host":
		return hxrt.NewArray(hxrt.StringFromLiteral("fromIPv4"), hxrt.StringFromLiteral("localhost"))
	case "sys.net.Socket":
		return hxrt.NewArray(hxrt.StringFromLiteral("pick"), hxrt.StringFromLiteral("publicAddress"), hxrt.StringFromLiteral("select"))
	case "sys.net.SocketInput":
		return hxrt.NewArray(hxrt.StringFromLiteral("translateReadStatus"))
	case "sys.net.SocketOutput":
		return hxrt.NewArray(hxrt.StringFromLiteral("translateWriteStatus"))
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
	case "StringBuf":
		return hxrt.NewArray()
	case "StringTools":
		return hxrt.NewArray()
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe.ds.StringMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.exceptions.NotImplementedException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("posInfos"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.exceptions.PosException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("posInfos"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.http.HttpBase":
		return hxrt.NewArray(hxrt.StringFromLiteral("addHeader"), hxrt.StringFromLiteral("addParameter"), hxrt.StringFromLiteral("emptyOnData"), hxrt.StringFromLiteral("get_responseData"), hxrt.StringFromLiteral("hasOnData"), hxrt.StringFromLiteral("headers"), hxrt.StringFromLiteral("onBytes"), hxrt.StringFromLiteral("onData"), hxrt.StringFromLiteral("onError"), hxrt.StringFromLiteral("onStatus"), hxrt.StringFromLiteral("params"), hxrt.StringFromLiteral("postBytes"), hxrt.StringFromLiteral("postData"), hxrt.StringFromLiteral("request"), hxrt.StringFromLiteral("responseAsString"), hxrt.StringFromLiteral("responseBytes"), hxrt.StringFromLiteral("responseData"), hxrt.StringFromLiteral("setHeader"), hxrt.StringFromLiteral("setParameter"), hxrt.StringFromLiteral("setPostBytes"), hxrt.StringFromLiteral("setPostData"), hxrt.StringFromLiteral("success"), hxrt.StringFromLiteral("url"))
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_dataExposed"), hxrt.StringFromLiteral("__hx_nativeView"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("blit"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("fill"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getData"), hxrt.StringFromLiteral("getDouble"), hxrt.StringFromLiteral("getFloat"), hxrt.StringFromLiteral("getInt32"), hxrt.StringFromLiteral("getInt64"), hxrt.StringFromLiteral("getString"), hxrt.StringFromLiteral("getUInt16"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setDouble"), hxrt.StringFromLiteral("setFloat"), hxrt.StringFromLiteral("setInt32"), hxrt.StringFromLiteral("setInt64"), hxrt.StringFromLiteral("setUInt16"), hxrt.StringFromLiteral("sub"), hxrt.StringFromLiteral("toHex"), hxrt.StringFromLiteral("toString"))
	case "haxe.io.BytesBuffer":
		return hxrt.NewArray(hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("addByte"), hxrt.StringFromLiteral("addBytes"), hxrt.StringFromLiteral("addDouble"), hxrt.StringFromLiteral("addFloat"), hxrt.StringFromLiteral("addInt32"), hxrt.StringFromLiteral("addInt64"), hxrt.StringFromLiteral("addString"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("getBytes"), hxrt.StringFromLiteral("get_length"), hxrt.StringFromLiteral("length"))
	case "haxe.io.BytesOutput":
		return hxrt.NewArray(hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("flush"), hxrt.StringFromLiteral("getBytes"), hxrt.StringFromLiteral("get_length"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("prepare"), hxrt.StringFromLiteral("set_bigEndian"), hxrt.StringFromLiteral("write"), hxrt.StringFromLiteral("writeByte"), hxrt.StringFromLiteral("writeBytes"), hxrt.StringFromLiteral("writeDouble"), hxrt.StringFromLiteral("writeFloat"), hxrt.StringFromLiteral("writeFullBytes"), hxrt.StringFromLiteral("writeInput"), hxrt.StringFromLiteral("writeInt16"), hxrt.StringFromLiteral("writeInt24"), hxrt.StringFromLiteral("writeInt32"), hxrt.StringFromLiteral("writeInt8"), hxrt.StringFromLiteral("writeString"), hxrt.StringFromLiteral("writeUInt16"), hxrt.StringFromLiteral("writeUInt24"))
	case "haxe.io.Eof":
		return hxrt.NewArray(hxrt.StringFromLiteral("toString"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray()
	case "haxe.io.Input":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("readAll"), hxrt.StringFromLiteral("readByte"), hxrt.StringFromLiteral("readBytes"), hxrt.StringFromLiteral("readDouble"), hxrt.StringFromLiteral("readFloat"), hxrt.StringFromLiteral("readFullBytes"), hxrt.StringFromLiteral("readInt16"), hxrt.StringFromLiteral("readInt24"), hxrt.StringFromLiteral("readInt32"), hxrt.StringFromLiteral("readInt8"), hxrt.StringFromLiteral("readLine"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("readUInt16"), hxrt.StringFromLiteral("readUInt24"), hxrt.StringFromLiteral("readUntil"), hxrt.StringFromLiteral("set_bigEndian"))
	case "haxe.io.Output":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("flush"), hxrt.StringFromLiteral("prepare"), hxrt.StringFromLiteral("set_bigEndian"), hxrt.StringFromLiteral("write"), hxrt.StringFromLiteral("writeByte"), hxrt.StringFromLiteral("writeBytes"), hxrt.StringFromLiteral("writeDouble"), hxrt.StringFromLiteral("writeFloat"), hxrt.StringFromLiteral("writeFullBytes"), hxrt.StringFromLiteral("writeInput"), hxrt.StringFromLiteral("writeInt16"), hxrt.StringFromLiteral("writeInt24"), hxrt.StringFromLiteral("writeInt32"), hxrt.StringFromLiteral("writeInt8"), hxrt.StringFromLiteral("writeString"), hxrt.StringFromLiteral("writeUInt16"), hxrt.StringFromLiteral("writeUInt24"))
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("next"))
	case "haxe.iterators.StringIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case "sys.Http":
		return hxrt.NewArray(hxrt.StringFromLiteral("addHeader"), hxrt.StringFromLiteral("addParameter"), hxrt.StringFromLiteral("buildMultipartBody"), hxrt.StringFromLiteral("cnxTimeout"), hxrt.StringFromLiteral("customRequest"), hxrt.StringFromLiteral("emptyOnData"), hxrt.StringFromLiteral("encodedParameters"), hxrt.StringFromLiteral("file"), hxrt.StringFromLiteral("fileTransfer"), hxrt.StringFromLiteral("fileTransfert"), hxrt.StringFromLiteral("getResponseHeaderValues"), hxrt.StringFromLiteral("get_responseData"), hxrt.StringFromLiteral("handleDataRequest"), hxrt.StringFromLiteral("hasHeader"), hxrt.StringFromLiteral("hasOnData"), hxrt.StringFromLiteral("headers"), hxrt.StringFromLiteral("noShutdown"), hxrt.StringFromLiteral("onBytes"), hxrt.StringFromLiteral("onData"), hxrt.StringFromLiteral("onError"), hxrt.StringFromLiteral("onStatus"), hxrt.StringFromLiteral("params"), hxrt.StringFromLiteral("postBytes"), hxrt.StringFromLiteral("postData"), hxrt.StringFromLiteral("recordResponseHeaders"), hxrt.StringFromLiteral("request"), hxrt.StringFromLiteral("requestWith"), hxrt.StringFromLiteral("resetResponseHeaders"), hxrt.StringFromLiteral("responseAsString"), hxrt.StringFromLiteral("responseBytes"), hxrt.StringFromLiteral("responseData"), hxrt.StringFromLiteral("responseHeaders"), hxrt.StringFromLiteral("responseHeadersSameKey"), hxrt.StringFromLiteral("setHeader"), hxrt.StringFromLiteral("setParameter"), hxrt.StringFromLiteral("setPostBytes"), hxrt.StringFromLiteral("setPostData"), hxrt.StringFromLiteral("success"), hxrt.StringFromLiteral("url"))
	case "sys.net.Host":
		return hxrt.NewArray(hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("ip"), hxrt.StringFromLiteral("reverse"), hxrt.StringFromLiteral("toString"))
	case "sys.net.Socket":
		return hxrt.NewArray(hxrt.StringFromLiteral("accept"), hxrt.StringFromLiteral("bind"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("connect"), hxrt.StringFromLiteral("custom"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("input"), hxrt.StringFromLiteral("listen"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("peer"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("replaceHandle"), hxrt.StringFromLiteral("setBlocking"), hxrt.StringFromLiteral("setFastSend"), hxrt.StringFromLiteral("setTimeout"), hxrt.StringFromLiteral("shutdown"), hxrt.StringFromLiteral("waitForRead"), hxrt.StringFromLiteral("write"))
	case "sys.net.SocketInput":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("readAll"), hxrt.StringFromLiteral("readByte"), hxrt.StringFromLiteral("readBytes"), hxrt.StringFromLiteral("readDouble"), hxrt.StringFromLiteral("readFloat"), hxrt.StringFromLiteral("readFullBytes"), hxrt.StringFromLiteral("readInt16"), hxrt.StringFromLiteral("readInt24"), hxrt.StringFromLiteral("readInt32"), hxrt.StringFromLiteral("readInt8"), hxrt.StringFromLiteral("readLine"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("readUInt16"), hxrt.StringFromLiteral("readUInt24"), hxrt.StringFromLiteral("readUntil"), hxrt.StringFromLiteral("set_bigEndian"))
	case "sys.net.SocketOutput":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("flush"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("prepare"), hxrt.StringFromLiteral("set_bigEndian"), hxrt.StringFromLiteral("write"), hxrt.StringFromLiteral("writeByte"), hxrt.StringFromLiteral("writeBytes"), hxrt.StringFromLiteral("writeDouble"), hxrt.StringFromLiteral("writeFloat"), hxrt.StringFromLiteral("writeFullBytes"), hxrt.StringFromLiteral("writeInput"), hxrt.StringFromLiteral("writeInt16"), hxrt.StringFromLiteral("writeInt24"), hxrt.StringFromLiteral("writeInt32"), hxrt.StringFromLiteral("writeInt8"), hxrt.StringFromLiteral("writeString"), hxrt.StringFromLiteral("writeUInt16"), hxrt.StringFromLiteral("writeUInt24"))
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
	case "StringBuf":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringTools":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.StringMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.exceptions.NotImplementedException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.exceptions.PosException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.http.HttpBase":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Bytes":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.BytesBuffer":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.BytesOutput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Eof":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.FPHelper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Input":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Output":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.MapKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.StringIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.StringKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.Http":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.net.Host":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.net.Socket":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.net.SocketInput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.net.SocketOutput":
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
	case "haxe.io.Encoding":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Error":
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
	case *haxe__io__Encoding:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("UTF8")
		case 1:
			return hxrt.StringFromLiteral("RawNative")
		default:
			return nil
		}
	case *haxe__io__Error:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("Blocked")
		case 1:
			return hxrt.StringFromLiteral("Overflow")
		case 2:
			return hxrt.StringFromLiteral("OutsideBounds")
		case 3:
			return hxrt.StringFromLiteral("Custom")
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
	case *haxe__io__Encoding:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__io__Error:
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
	case "haxe.io.Encoding":
		return hxrt.NewArray(hxrt.StringFromLiteral("UTF8"), hxrt.StringFromLiteral("RawNative"))
	case "haxe.io.Error":
		return hxrt.NewArray(hxrt.StringFromLiteral("Blocked"), hxrt.StringFromLiteral("Overflow"), hxrt.StringFromLiteral("OutsideBounds"), hxrt.StringFromLiteral("Custom"))
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
	case *haxe__io__Encoding:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__io__Error:
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
	case "haxe.io.Encoding":
		return hxrt.NewArray(haxe__io__Encoding_UTF8, haxe__io__Encoding_RawNative)
	case "haxe.io.Error":
		return hxrt.NewArray(haxe__io__Error_Blocked, haxe__io__Error_Overflow, haxe__io__Error_OutsideBounds)
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
