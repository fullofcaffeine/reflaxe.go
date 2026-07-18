package main

import (
	"examples_incident_api_metal/hxrt"
	"reflect"
)

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func argValue(name *string, fallback *string) *string {
	args := hxrt.ArrayFromValues(func(hx_sort_src_36 []*string) []any {
		hx_sort_out_38 := make([]any, 0, len(hx_sort_src_36))
		for _, hx_sort_item_37 := range hx_sort_src_36 {
			hx_sort_out_38 = append(hx_sort_out_38, hx_sort_item_37)
		}
		return hx_sort_out_38
	}(hxrt.SysArgs()))
	i := 0
	for i < int(int32((hxrt.Int32Wrap(args.Len()) - hxrt.Int32Wrap(1)))) {
		if hxrt.StringEqualAny(args.Get(i), name) {
			return func(hx_value_39 any) *string {
				if hx_value_39 == nil {
					var hx_zero_40 *string
					return hx_zero_40
				}
				return hx_value_39.(*string)
			}(args.Get(int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1))))))
		}
		i = int(int32((i + 1)))
	}
	return fallback
}

func hasArg(name *string) bool {
	_g := 0
	_g1 := hxrt.ArrayFromValues(func(hx_sort_src_41 []*string) []any {
		hx_sort_out_43 := make([]any, 0, len(hx_sort_src_41))
		for _, hx_sort_item_42 := range hx_sort_src_41 {
			hx_sort_out_43 = append(hx_sort_out_43, hx_sort_item_42)
		}
		return hx_sort_out_43
	}(hxrt.SysArgs()))
	for _g < _g1.Len() {
		arg := func(hx_value_44 any) *string {
			if hx_value_44 == nil {
				var hx_zero_45 *string
				return hx_zero_45
			}
			return hx_value_44.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(arg, name) {
			return true
		}
	}
	return false
}

func main() {
	if hasArg(hxrt.StringFromLiteral("--scripted")) {
		var v any = any(Harness_run())
		hxrt.Println(v)
		return
	}
	if hasArg(hxrt.StringFromLiteral("init-config")) {
		configPath := argValue(hxrt.StringFromLiteral("--config"), hxrt.StringFromLiteral("config.json"))
		app__core__IncidentConfig_saveExample(configPath)
		hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("wrote "), configPath)))
		return
	}
	if hasArg(hxrt.StringFromLiteral("serve")) {
		serve(argValue(hxrt.StringFromLiteral("--config"), hxrt.StringFromLiteral("config.json")))
		return
	}
	printHelp()
}

func printHelp() {
	hxrt.Println(any(hxrt.StringFromLiteral("incident_api commands:")))
	hxrt.Println(any(hxrt.StringFromLiteral("  --scripted")))
	hxrt.Println(any(hxrt.StringFromLiteral("  init-config --config <path>")))
	hxrt.Println(any(hxrt.StringFromLiteral("  serve --config <path>")))
	hxrt.Println(any(hxrt.StringFromLiteral("curl examples:")))
	hxrt.Println(any(hxrt.StringFromLiteral("  curl http://127.0.0.1:8080/health")))
	hxrt.Println(any(hxrt.StringFromLiteral("  curl -X POST -d '{\"title\":\"Database lag\",\"severity\":\"high\"}' http://127.0.0.1:8080/incidents")))
}

func serve(configPath *string) {
	config := app__core__IncidentConfig_load(configPath)
	api := New_app__core__IncidentApi(config, New_app__core__IncidentStore(config.statePath))
	server := New_app__http__TinyHttpServer(api, config.host, config.port)
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("incident_api listening on http://"), server.host), hxrt.StringFromLiteral(":")), server.port))
	hxrt.Println(v)
	for true {
		server.serveOnce()
	}
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
	case "Date":
		return hxrt_typeCallAny(New_Date, args)
	case "Harness":
		return nil, false
	case "Main":
		return nil, false
	case "StringBuf":
		return nil, false
	case "StringTools":
		return nil, false
	case "app.core.Incident":
		return hxrt_typeCallAny(New_app__core__Incident, args)
	case "app.core.IncidentApi":
		return hxrt_typeCallAny(New_app__core__IncidentApi, args)
	case "app.core.IncidentConfig":
		return hxrt_typeCallAny(New_app__core__IncidentConfig, args)
	case "app.core.IncidentRequestException":
		return hxrt_typeCallAny(New_app__core__IncidentRequestException, args)
	case "app.core.IncidentStore":
		return hxrt_typeCallAny(New_app__core__IncidentStore, args)
	case "app.http.HttpRequest":
		return hxrt_typeCallAny(New_app__http__HttpRequest, args)
	case "app.http.HttpResponse":
		return hxrt_typeCallAny(New_app__http__HttpResponse, args)
	case "app.http.TinyHttpServer":
		return hxrt_typeCallAny(New_app__http__TinyHttpServer, args)
	case "haxe.Int64Helper":
		return nil, false
	case "haxe.Json":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe.exceptions.NotImplementedException":
		return hxrt_typeCallAny(New_haxe__exceptions__NotImplementedException, args)
	case "haxe.exceptions.PosException":
		return hxrt_typeCallAny(New_haxe__exceptions__PosException, args)
	case "haxe.io.Bytes":
		return hxrt_typeCallAny(New_haxe__io__Bytes, args)
	case "haxe.io.BytesBuffer":
		return hxrt_typeCallAny(New_haxe__io__BytesBuffer, args)
	case "haxe.io.Eof":
		return hxrt_typeCallAny(New_haxe__io__Eof, args)
	case "haxe.io.FPHelper":
		return nil, false
	case "haxe.io.Input":
		return hxrt_typeCallAny(New_haxe__io__Input, args)
	case "haxe.io.Output":
		return hxrt_typeCallAny(New_haxe__io__Output, args)
	case "haxe.iterators.StringIterator":
		return hxrt_typeCallAny(New_haxe__iterators__StringIterator, args)
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__StringKeyValueIterator, args)
	case "sys.FileSystem":
		return nil, false
	case "sys.io.File":
		return nil, false
	case "sys.io.FileInput":
		return hxrt_typeCallAny(New_sys__io__FileInput, args)
	case "sys.io.FileOutput":
		return hxrt_typeCallAny(New_sys__io__FileOutput, args)
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
	case "Date":
		return &Date{}, true
	case "app.core.Incident":
		return &app__core__Incident{}, true
	case "app.core.IncidentApi":
		return &app__core__IncidentApi{}, true
	case "app.core.IncidentConfig":
		return &app__core__IncidentConfig{}, true
	case "app.core.IncidentRequestException":
		return &app__core__IncidentRequestException{}, true
	case "app.core.IncidentStore":
		return &app__core__IncidentStore{}, true
	case "app.http.HttpRequest":
		return &app__http__HttpRequest{}, true
	case "app.http.HttpResponse":
		return &app__http__HttpResponse{}, true
	case "app.http.TinyHttpServer":
		return &app__http__TinyHttpServer{}, true
	case "haxe._Int64.___Int64":
		return &haxe___Int64_____Int64{}, true
	case "haxe.exceptions.NotImplementedException":
		return &haxe__exceptions__NotImplementedException{}, true
	case "haxe.exceptions.PosException":
		return &haxe__exceptions__PosException{}, true
	case "haxe.io.Bytes":
		return &haxe__io__Bytes{}, true
	case "haxe.io.BytesBuffer":
		return &haxe__io__BytesBuffer{}, true
	case "haxe.io.Eof":
		return &haxe__io__Eof{}, true
	case "haxe.io.Input":
		return &haxe__io__Input{}, true
	case "haxe.io.Output":
		return &haxe__io__Output{}, true
	case "haxe.iterators.StringIterator":
		return &haxe__iterators__StringIterator{}, true
	case "haxe.iterators.StringKeyValueIterator":
		return &haxe__iterators__StringKeyValueIterator{}, true
	case "sys.io.FileInput":
		return &sys__io__FileInput{}, true
	case "sys.io.FileOutput":
		return &sys__io__FileOutput{}, true
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
	case "sys.io.FileSeek":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return sys__io__FileSeek_SeekBegin, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return sys__io__FileSeek_SeekCur, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return sys__io__FileSeek_SeekEnd, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "SeekBegin":
			if len(args) != 0 {
				return nil, false
			}
			return sys__io__FileSeek_SeekBegin, true
		case "SeekCur":
			if len(args) != 0 {
				return nil, false
			}
			return sys__io__FileSeek_SeekCur, true
		case "SeekEnd":
			if len(args) != 0 {
				return nil, false
			}
			return sys__io__FileSeek_SeekEnd, true
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
	case *Date:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Date")}
	case *app__core__Incident:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("app.core.Incident")}
	case *app__core__IncidentApi:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("app.core.IncidentApi")}
	case *app__core__IncidentConfig:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("app.core.IncidentConfig")}
	case *app__core__IncidentRequestException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("app.core.IncidentRequestException")}
	case *app__core__IncidentStore:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("app.core.IncidentStore")}
	case *app__http__HttpRequest:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("app.http.HttpRequest")}
	case *app__http__HttpResponse:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("app.http.HttpResponse")}
	case *app__http__TinyHttpServer:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("app.http.TinyHttpServer")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
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
	case *sys__io__FileInput:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.io.FileInput")}
	case *sys__io__FileOutput:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.io.FileOutput")}
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
	case *sys__io__FileSeek:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("sys.io.FileSeek")}
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
	case "Date":
		return nil
	case "Harness":
		return nil
	case "Main":
		return nil
	case "StringBuf":
		return nil
	case "StringTools":
		return nil
	case "app.core.Incident":
		return nil
	case "app.core.IncidentApi":
		return nil
	case "app.core.IncidentConfig":
		return nil
	case "app.core.IncidentRequestException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "app.core.IncidentStore":
		return nil
	case "app.http.HttpRequest":
		return nil
	case "app.http.HttpResponse":
		return nil
	case "app.http.TinyHttpServer":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe.Json":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe.exceptions.NotImplementedException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.PosException")}
	case "haxe.exceptions.PosException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "haxe.io.Bytes":
		return nil
	case "haxe.io.BytesBuffer":
		return nil
	case "haxe.io.Eof":
		return nil
	case "haxe.io.FPHelper":
		return nil
	case "haxe.io.Input":
		return nil
	case "haxe.io.Output":
		return nil
	case "haxe.iterators.StringIterator":
		return nil
	case "haxe.iterators.StringKeyValueIterator":
		return nil
	case "sys.FileSystem":
		return nil
	case "sys.io.File":
		return nil
	case "sys.io.FileInput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Input")}
	case "sys.io.FileOutput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Output")}
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
	case "Date":
		return hxrt.NewArray(hxrt.StringFromLiteral("fromMilliseconds"), hxrt.StringFromLiteral("fromString"), hxrt.StringFromLiteral("fromTime"), hxrt.StringFromLiteral("now"))
	case "Harness":
		return hxrt.NewArray(hxrt.StringFromLiteral("CONFIG_FILE"), hxrt.StringFromLiteral("STATE_FILE"), hxrt.StringFromLiteral("cleanup"), hxrt.StringFromLiteral("request"), hxrt.StringFromLiteral("run"), hxrt.StringFromLiteral("summarize"))
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("argValue"), hxrt.StringFromLiteral("hasArg"), hxrt.StringFromLiteral("main"), hxrt.StringFromLiteral("printHelp"), hxrt.StringFromLiteral("serve"))
	case "StringBuf":
		return hxrt.NewArray()
	case "StringTools":
		return hxrt.NewArray(hxrt.StringFromLiteral("MAX_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("contains"), hxrt.StringFromLiteral("containsImpl"), hxrt.StringFromLiteral("endsWith"), hxrt.StringFromLiteral("endsWithImpl"), hxrt.StringFromLiteral("fastCodeAt"), hxrt.StringFromLiteral("hex"), hxrt.StringFromLiteral("hexDigitValue"), hxrt.StringFromLiteral("htmlEscape"), hxrt.StringFromLiteral("htmlUnescape"), hxrt.StringFromLiteral("isEof"), hxrt.StringFromLiteral("isSpace"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("lpad"), hxrt.StringFromLiteral("ltrim"), hxrt.StringFromLiteral("replace"), hxrt.StringFromLiteral("rpad"), hxrt.StringFromLiteral("rtrim"), hxrt.StringFromLiteral("startsWith"), hxrt.StringFromLiteral("startsWithImpl"), hxrt.StringFromLiteral("trim"), hxrt.StringFromLiteral("unsafeCodeAt"), hxrt.StringFromLiteral("urlDecode"), hxrt.StringFromLiteral("urlEncode"), hxrt.StringFromLiteral("utf16CodePointAt"))
	case "app.core.Incident":
		return hxrt.NewArray(hxrt.StringFromLiteral("boolJson"), hxrt.StringFromLiteral("jsonEscape"))
	case "app.core.IncidentApi":
		return hxrt.NewArray(hxrt.StringFromLiteral("fieldString"), hxrt.StringFromLiteral("parseJsonBody"))
	case "app.core.IncidentConfig":
		return hxrt.NewArray(hxrt.StringFromLiteral("defaults"), hxrt.StringFromLiteral("intField"), hxrt.StringFromLiteral("load"), hxrt.StringFromLiteral("saveExample"), hxrt.StringFromLiteral("stringField"))
	case "app.core.IncidentRequestException":
		return hxrt.NewArray()
	case "app.core.IncidentStore":
		return hxrt.NewArray(hxrt.StringFromLiteral("boolField"), hxrt.StringFromLiteral("intField"), hxrt.StringFromLiteral("normalizeSeverity"), hxrt.StringFromLiteral("stringField"))
	case "app.http.HttpRequest":
		return hxrt.NewArray()
	case "app.http.HttpResponse":
		return hxrt.NewArray(hxrt.StringFromLiteral("json"))
	case "app.http.TinyHttpServer":
		return hxrt.NewArray(hxrt.StringFromLiteral("closePeer"), hxrt.StringFromLiteral("reason"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Json":
		return hxrt.NewArray(hxrt.StringFromLiteral("parse"), hxrt.StringFromLiteral("stringify"))
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe.exceptions.NotImplementedException":
		return hxrt.NewArray()
	case "haxe.exceptions.PosException":
		return hxrt.NewArray()
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_fromNativeView"), hxrt.StringFromLiteral("alloc"), hxrt.StringFromLiteral("fastGet"), hxrt.StringFromLiteral("ofData"), hxrt.StringFromLiteral("ofHex"), hxrt.StringFromLiteral("ofString"), hxrt.StringFromLiteral("rawNativeUsesUtf16LE"))
	case "haxe.io.BytesBuffer":
		return hxrt.NewArray()
	case "haxe.io.Eof":
		return hxrt.NewArray()
	case "haxe.io.FPHelper":
		return hxrt.NewArray(hxrt.StringFromLiteral("doubleToI64"), hxrt.StringFromLiteral("floatToI32"), hxrt.StringFromLiteral("i32ToFloat"), hxrt.StringFromLiteral("i64ToDouble"))
	case "haxe.io.Input":
		return hxrt.NewArray()
	case "haxe.io.Output":
		return hxrt.NewArray()
	case "haxe.iterators.StringIterator":
		return hxrt.NewArray()
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt.NewArray()
	case "sys.FileSystem":
		return hxrt.NewArray(hxrt.StringFromLiteral("absolutePath"), hxrt.StringFromLiteral("createDirectory"), hxrt.StringFromLiteral("deleteDirectory"), hxrt.StringFromLiteral("deleteFile"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("fullPath"), hxrt.StringFromLiteral("isDirectory"), hxrt.StringFromLiteral("readDirectory"), hxrt.StringFromLiteral("rename"), hxrt.StringFromLiteral("stat"))
	case "sys.io.File":
		return hxrt.NewArray(hxrt.StringFromLiteral("append"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("getBytes"), hxrt.StringFromLiteral("getContent"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("saveBytes"), hxrt.StringFromLiteral("saveContent"), hxrt.StringFromLiteral("update"), hxrt.StringFromLiteral("write"))
	case "sys.io.FileInput":
		return hxrt.NewArray()
	case "sys.io.FileOutput":
		return hxrt.NewArray()
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
	case "Date":
		return hxrt.NewArray(hxrt.StringFromLiteral("getDate"), hxrt.StringFromLiteral("getDay"), hxrt.StringFromLiteral("getFullYear"), hxrt.StringFromLiteral("getHours"), hxrt.StringFromLiteral("getMinutes"), hxrt.StringFromLiteral("getMonth"), hxrt.StringFromLiteral("getSeconds"), hxrt.StringFromLiteral("getTime"), hxrt.StringFromLiteral("getTimezoneOffset"), hxrt.StringFromLiteral("getUTCDate"), hxrt.StringFromLiteral("getUTCDay"), hxrt.StringFromLiteral("getUTCFullYear"), hxrt.StringFromLiteral("getUTCHours"), hxrt.StringFromLiteral("getUTCMinutes"), hxrt.StringFromLiteral("getUTCMonth"), hxrt.StringFromLiteral("getUTCSeconds"), hxrt.StringFromLiteral("localParts"), hxrt.StringFromLiteral("ms"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("utcParts"))
	case "Harness":
		return hxrt.NewArray()
	case "Main":
		return hxrt.NewArray()
	case "StringBuf":
		return hxrt.NewArray()
	case "StringTools":
		return hxrt.NewArray()
	case "app.core.Incident":
		return hxrt.NewArray(hxrt.StringFromLiteral("acknowledged"), hxrt.StringFromLiteral("createdAt"), hxrt.StringFromLiteral("id"), hxrt.StringFromLiteral("resolved"), hxrt.StringFromLiteral("severity"), hxrt.StringFromLiteral("title"), hxrt.StringFromLiteral("toJson"))
	case "app.core.IncidentApi":
		return hxrt.NewArray(hxrt.StringFromLiteral("config"), hxrt.StringFromLiteral("createIncident"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("requests"), hxrt.StringFromLiteral("store"), hxrt.StringFromLiteral("updateIncident"))
	case "app.core.IncidentConfig":
		return hxrt.NewArray(hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("port"), hxrt.StringFromLiteral("serviceName"), hxrt.StringFromLiteral("statePath"))
	case "app.core.IncidentRequestException":
		return hxrt.NewArray(hxrt.StringFromLiteral("code"), hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "app.core.IncidentStore":
		return hxrt.NewArray(hxrt.StringFromLiteral("acknowledge"), hxrt.StringFromLiteral("create"), hxrt.StringFromLiteral("find"), hxrt.StringFromLiteral("incidents"), hxrt.StringFromLiteral("listJson"), hxrt.StringFromLiteral("load"), hxrt.StringFromLiteral("metricsJson"), hxrt.StringFromLiteral("nextId"), hxrt.StringFromLiteral("resolve"), hxrt.StringFromLiteral("save"), hxrt.StringFromLiteral("statePath"))
	case "app.http.HttpRequest":
		return hxrt.NewArray(hxrt.StringFromLiteral("body"), hxrt.StringFromLiteral("method"), hxrt.StringFromLiteral("path"))
	case "app.http.HttpResponse":
		return hxrt.NewArray(hxrt.StringFromLiteral("body"), hxrt.StringFromLiteral("status"))
	case "app.http.TinyHttpServer":
		return hxrt.NewArray(hxrt.StringFromLiteral("api"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("port"), hxrt.StringFromLiteral("readBody"), hxrt.StringFromLiteral("readRequest"), hxrt.StringFromLiteral("serveOnce"), hxrt.StringFromLiteral("server"), hxrt.StringFromLiteral("writeResponse"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Json":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe.exceptions.NotImplementedException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("posInfos"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.exceptions.PosException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("posInfos"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_dataExposed"), hxrt.StringFromLiteral("__hx_nativeView"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("blit"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("fill"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getData"), hxrt.StringFromLiteral("getDouble"), hxrt.StringFromLiteral("getFloat"), hxrt.StringFromLiteral("getInt32"), hxrt.StringFromLiteral("getInt64"), hxrt.StringFromLiteral("getString"), hxrt.StringFromLiteral("getUInt16"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setDouble"), hxrt.StringFromLiteral("setFloat"), hxrt.StringFromLiteral("setInt32"), hxrt.StringFromLiteral("setInt64"), hxrt.StringFromLiteral("setUInt16"), hxrt.StringFromLiteral("sub"), hxrt.StringFromLiteral("toHex"), hxrt.StringFromLiteral("toString"))
	case "haxe.io.BytesBuffer":
		return hxrt.NewArray(hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("addByte"), hxrt.StringFromLiteral("addBytes"), hxrt.StringFromLiteral("addDouble"), hxrt.StringFromLiteral("addFloat"), hxrt.StringFromLiteral("addInt32"), hxrt.StringFromLiteral("addInt64"), hxrt.StringFromLiteral("addString"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("getBytes"), hxrt.StringFromLiteral("get_length"), hxrt.StringFromLiteral("length"))
	case "haxe.io.Eof":
		return hxrt.NewArray(hxrt.StringFromLiteral("toString"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray()
	case "haxe.io.Input":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("readAll"), hxrt.StringFromLiteral("readByte"), hxrt.StringFromLiteral("readBytes"), hxrt.StringFromLiteral("readDouble"), hxrt.StringFromLiteral("readFloat"), hxrt.StringFromLiteral("readFullBytes"), hxrt.StringFromLiteral("readInt16"), hxrt.StringFromLiteral("readInt24"), hxrt.StringFromLiteral("readInt32"), hxrt.StringFromLiteral("readInt8"), hxrt.StringFromLiteral("readLine"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("readUInt16"), hxrt.StringFromLiteral("readUInt24"), hxrt.StringFromLiteral("readUntil"), hxrt.StringFromLiteral("set_bigEndian"))
	case "haxe.io.Output":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("flush"), hxrt.StringFromLiteral("prepare"), hxrt.StringFromLiteral("set_bigEndian"), hxrt.StringFromLiteral("write"), hxrt.StringFromLiteral("writeByte"), hxrt.StringFromLiteral("writeBytes"), hxrt.StringFromLiteral("writeDouble"), hxrt.StringFromLiteral("writeFloat"), hxrt.StringFromLiteral("writeFullBytes"), hxrt.StringFromLiteral("writeInput"), hxrt.StringFromLiteral("writeInt16"), hxrt.StringFromLiteral("writeInt24"), hxrt.StringFromLiteral("writeInt32"), hxrt.StringFromLiteral("writeInt8"), hxrt.StringFromLiteral("writeString"), hxrt.StringFromLiteral("writeUInt16"), hxrt.StringFromLiteral("writeUInt24"))
	case "haxe.iterators.StringIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case "sys.FileSystem":
		return hxrt.NewArray()
	case "sys.io.File":
		return hxrt.NewArray()
	case "sys.io.FileInput":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("eof"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("readAll"), hxrt.StringFromLiteral("readByte"), hxrt.StringFromLiteral("readBytes"), hxrt.StringFromLiteral("readDouble"), hxrt.StringFromLiteral("readFloat"), hxrt.StringFromLiteral("readFullBytes"), hxrt.StringFromLiteral("readInt16"), hxrt.StringFromLiteral("readInt24"), hxrt.StringFromLiteral("readInt32"), hxrt.StringFromLiteral("readInt8"), hxrt.StringFromLiteral("readLine"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("readUInt16"), hxrt.StringFromLiteral("readUInt24"), hxrt.StringFromLiteral("readUntil"), hxrt.StringFromLiteral("seek"), hxrt.StringFromLiteral("set_bigEndian"), hxrt.StringFromLiteral("tell"))
	case "sys.io.FileOutput":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("flush"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("prepare"), hxrt.StringFromLiteral("seek"), hxrt.StringFromLiteral("set_bigEndian"), hxrt.StringFromLiteral("tell"), hxrt.StringFromLiteral("write"), hxrt.StringFromLiteral("writeByte"), hxrt.StringFromLiteral("writeBytes"), hxrt.StringFromLiteral("writeDouble"), hxrt.StringFromLiteral("writeFloat"), hxrt.StringFromLiteral("writeFullBytes"), hxrt.StringFromLiteral("writeInput"), hxrt.StringFromLiteral("writeInt16"), hxrt.StringFromLiteral("writeInt24"), hxrt.StringFromLiteral("writeInt32"), hxrt.StringFromLiteral("writeInt8"), hxrt.StringFromLiteral("writeString"), hxrt.StringFromLiteral("writeUInt16"), hxrt.StringFromLiteral("writeUInt24"))
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
	case "Date":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Harness":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringBuf":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringTools":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "app.core.Incident":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "app.core.IncidentApi":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "app.core.IncidentConfig":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "app.core.IncidentRequestException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "app.core.IncidentStore":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "app.http.HttpRequest":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "app.http.HttpResponse":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "app.http.TinyHttpServer":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Json":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.exceptions.NotImplementedException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.exceptions.PosException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Bytes":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.BytesBuffer":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Eof":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.FPHelper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Input":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Output":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.StringIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.StringKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.FileSystem":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.io.File":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.io.FileInput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.io.FileOutput":
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
	case "sys.io.FileSeek":
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
	case *sys__io__FileSeek:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("SeekBegin")
		case 1:
			return hxrt.StringFromLiteral("SeekCur")
		case 2:
			return hxrt.StringFromLiteral("SeekEnd")
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
	case *sys__io__FileSeek:
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
	case "sys.io.FileSeek":
		return hxrt.NewArray(hxrt.StringFromLiteral("SeekBegin"), hxrt.StringFromLiteral("SeekCur"), hxrt.StringFromLiteral("SeekEnd"))
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
	case *sys__io__FileSeek:
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
	case "sys.io.FileSeek":
		return hxrt.NewArray(sys__io__FileSeek_SeekBegin, sys__io__FileSeek_SeekCur, sys__io__FileSeek_SeekEnd)
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
