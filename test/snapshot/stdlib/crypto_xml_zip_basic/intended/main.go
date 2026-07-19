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

func invalidZipThrows() bool {
	hx_try_return_1 := false
	var hx_try_value_2 bool
	hxrt.TryCatch(func() {
		haxe__zip__Uncompress_run(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("not-zlib"), nil), nil)
	}, func(hx_caught_3 any) {
		hx_tmp := hx_caught_3
		_ = hx_tmp
		hx_try_value_2 = true
		hx_try_return_1 = true
		return
	})
	if hx_try_return_1 {
		return hx_try_value_2
	}
	return false
}

func main() {
	payload := hxrt.StringFromLiteral("ab")
	bytes := haxe__io__Bytes_ofString(payload, nil)
	var v any = any(haxe__crypto__Base64_encode(bytes, true))
	hxrt.Println(v)
	var v_1 any = any(haxe__crypto__Base64_encode(bytes, false))
	hxrt.Println(v_1)
	var v_2 any = any(haxe__crypto__Base64_decode(hxrt.StringFromLiteral("YWI="), true).__hx_this.toString())
	hxrt.Println(v_2)
	var v_3 any = any(haxe__crypto__Base64_urlEncode(bytes, true))
	hxrt.Println(v_3)
	var v_4 any = any(haxe__crypto__Base64_urlDecode(hxrt.StringFromLiteral("YWI"), false).__hx_this.toString())
	hxrt.Println(v_4)
	var v_5 any = any(haxe__crypto__Md5_encode(payload))
	hxrt.Println(v_5)
	var v_6 any = any(haxe__crypto__Sha1_encode(payload))
	hxrt.Println(v_6)
	var v_7 any = any(haxe__crypto__Sha224_encode(payload))
	hxrt.Println(v_7)
	var v_8 any = any(haxe__crypto__Sha256_encode(payload))
	hxrt.Println(v_8)
	doc := haxe__xml__Parser_parse(hxrt.StringFromLiteral("<root><item n=\"1\">x</item></root>"), false)
	var v_9 any = any(haxe__xml__Printer_print(doc, false))
	hxrt.Println(v_9)
	compressed := haxe__zip__Compress_run(bytes, 9)
	roundtrip := haxe__zip__Uncompress_run(compressed, nil)
	var v_10 any = any(roundtrip.__hx_this.toString())
	hxrt.Println(v_10)
	var v_11 any = any((compressed.length > 0))
	hxrt.Println(v_11)
	var v_12 any = any(invalidZipThrows())
	hxrt.Println(v_12)
}

type Type struct {
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
	case "Std":
		return nil, false
	case "StringBuf":
		return hxrt_typeCallAny(New_StringBuf, args)
	case "StringTools":
		return nil, false
	case "Xml":
		return hxrt_typeCallAny(New_Xml, args)
	case "_Xml.XmlType_Impl_":
		return nil, false
	case "haxe.Int64Helper":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe.crypto.Base64":
		return nil, false
	case "haxe.crypto.Md5":
		return nil, false
	case "haxe.crypto.Sha1":
		return nil, false
	case "haxe.crypto.Sha224":
		return nil, false
	case "haxe.crypto.Sha256":
		return nil, false
	case "haxe.ds.StringMap":
		return hxrt_typeCallAny(New_haxe__ds__StringMap, args)
	case "haxe.io.Bytes":
		return hxrt_typeCallAny(New_haxe__io__Bytes, args)
	case "haxe.io.FPHelper":
		return nil, false
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__MapKeyValueIterator, args)
	case "haxe.iterators.StringIterator":
		return hxrt_typeCallAny(New_haxe__iterators__StringIterator, args)
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__StringKeyValueIterator, args)
	case "haxe.xml.Parser":
		return nil, false
	case "haxe.xml.Printer":
		return hxrt_typeCallAny(New_haxe__xml__Printer, args)
	case "haxe.xml.XmlParserException":
		return hxrt_typeCallAny(New_haxe__xml__XmlParserException, args)
	case "haxe.xml._Parser.S_Impl_":
		return nil, false
	case "haxe.zip.Compress":
		return hxrt_typeCallAny(New_haxe__zip__Compress, args)
	case "haxe.zip.Uncompress":
		return hxrt_typeCallAny(New_haxe__zip__Uncompress, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "StringBuf":
		return &StringBuf{}, true
	case "Xml":
		return &Xml{}, true
	case "haxe._Int64.___Int64":
		return &haxe___Int64_____Int64{}, true
	case "haxe.ds.StringMap":
		return &haxe__ds__StringMap{}, true
	case "haxe.io.Bytes":
		return &haxe__io__Bytes{}, true
	case "haxe.iterators.MapKeyValueIterator":
		return &haxe__iterators__MapKeyValueIterator{}, true
	case "haxe.iterators.StringIterator":
		return &haxe__iterators__StringIterator{}, true
	case "haxe.iterators.StringKeyValueIterator":
		return &haxe__iterators__StringKeyValueIterator{}, true
	case "haxe.xml.Printer":
		return &haxe__xml__Printer{}, true
	case "haxe.xml.XmlParserException":
		return &haxe__xml__XmlParserException{}, true
	case "haxe.zip.Compress":
		return &haxe__zip__Compress{}, true
	case "haxe.zip.Uncompress":
		return &haxe__zip__Uncompress{}, true
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
	case "haxe.zip.FlushMode":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__zip__FlushMode_NO, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__zip__FlushMode_SYNC, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__zip__FlushMode_FULL, true
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__zip__FlushMode_FINISH, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__zip__FlushMode_BLOCK, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "NO":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__zip__FlushMode_NO, true
		case "SYNC":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__zip__FlushMode_SYNC, true
		case "FULL":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__zip__FlushMode_FULL, true
		case "FINISH":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__zip__FlushMode_FINISH, true
		case "BLOCK":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__zip__FlushMode_BLOCK, true
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
	case *StringBuf:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("StringBuf")}
	case *Xml:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Xml")}
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
	case *haxe__io__Bytes:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Bytes")}
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
	case *haxe__xml__Printer:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.xml.Printer")}
	case *haxe__xml__XmlParserException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.xml.XmlParserException")}
	case *haxe__zip__Compress:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.zip.Compress")}
	case *haxe__zip__Uncompress:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.zip.Uncompress")}
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
	case *haxe__zip__FlushMode:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.zip.FlushMode")}
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
	case "Std":
		return nil
	case "StringBuf":
		return nil
	case "StringTools":
		return nil
	case "Xml":
		return nil
	case "_Xml.XmlType_Impl_":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe.crypto.Base64":
		return nil
	case "haxe.crypto.Md5":
		return nil
	case "haxe.crypto.Sha1":
		return nil
	case "haxe.crypto.Sha224":
		return nil
	case "haxe.crypto.Sha256":
		return nil
	case "haxe.ds.StringMap":
		return nil
	case "haxe.io.Bytes":
		return nil
	case "haxe.io.FPHelper":
		return nil
	case "haxe.iterators.MapKeyValueIterator":
		return nil
	case "haxe.iterators.StringIterator":
		return nil
	case "haxe.iterators.StringKeyValueIterator":
		return nil
	case "haxe.xml.Parser":
		return nil
	case "haxe.xml.Printer":
		return nil
	case "haxe.xml.XmlParserException":
		return nil
	case "haxe.xml._Parser.S_Impl_":
		return nil
	case "haxe.zip.Compress":
		return nil
	case "haxe.zip.Uncompress":
		return nil
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
		return hxrt.NewArray(hxrt.StringFromLiteral("invalidZipThrows"), hxrt.StringFromLiteral("main"))
	case "Std":
		return hxrt.NewArray(hxrt.StringFromLiteral("digitValue"), hxrt.StringFromLiteral("downcast"), hxrt.StringFromLiteral("instance"), hxrt.StringFromLiteral("int"), hxrt.StringFromLiteral("invalidFloat"), hxrt.StringFromLiteral("is"), hxrt.StringFromLiteral("isDecimalDigit"), hxrt.StringFromLiteral("isSpaceCode"), hxrt.StringFromLiteral("parseFloat"), hxrt.StringFromLiteral("parseInt"), hxrt.StringFromLiteral("random"))
	case "StringBuf":
		return hxrt.NewArray()
	case "StringTools":
		return hxrt.NewArray(hxrt.StringFromLiteral("MAX_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("contains"), hxrt.StringFromLiteral("containsImpl"), hxrt.StringFromLiteral("endsWith"), hxrt.StringFromLiteral("endsWithImpl"), hxrt.StringFromLiteral("fastCodeAt"), hxrt.StringFromLiteral("hex"), hxrt.StringFromLiteral("hexDigitValue"), hxrt.StringFromLiteral("htmlEscape"), hxrt.StringFromLiteral("htmlUnescape"), hxrt.StringFromLiteral("isEof"), hxrt.StringFromLiteral("isSpace"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("lpad"), hxrt.StringFromLiteral("ltrim"), hxrt.StringFromLiteral("replace"), hxrt.StringFromLiteral("rpad"), hxrt.StringFromLiteral("rtrim"), hxrt.StringFromLiteral("startsWith"), hxrt.StringFromLiteral("startsWithImpl"), hxrt.StringFromLiteral("trim"), hxrt.StringFromLiteral("unsafeCodeAt"), hxrt.StringFromLiteral("urlDecode"), hxrt.StringFromLiteral("urlEncode"), hxrt.StringFromLiteral("utf16CodePointAt"))
	case "Xml":
		return hxrt.NewArray(hxrt.StringFromLiteral("CData"), hxrt.StringFromLiteral("Comment"), hxrt.StringFromLiteral("DocType"), hxrt.StringFromLiteral("Document"), hxrt.StringFromLiteral("Element"), hxrt.StringFromLiteral("PCData"), hxrt.StringFromLiteral("ProcessingInstruction"), hxrt.StringFromLiteral("createCData"), hxrt.StringFromLiteral("createComment"), hxrt.StringFromLiteral("createDocType"), hxrt.StringFromLiteral("createDocument"), hxrt.StringFromLiteral("createElement"), hxrt.StringFromLiteral("createPCData"), hxrt.StringFromLiteral("createProcessingInstruction"), hxrt.StringFromLiteral("parse"))
	case "_Xml.XmlType_Impl_":
		return hxrt.NewArray(hxrt.StringFromLiteral("CData"), hxrt.StringFromLiteral("Comment"), hxrt.StringFromLiteral("DocType"), hxrt.StringFromLiteral("Document"), hxrt.StringFromLiteral("Element"), hxrt.StringFromLiteral("PCData"), hxrt.StringFromLiteral("ProcessingInstruction"), hxrt.StringFromLiteral("toString"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe.crypto.Base64":
		return hxrt.NewArray(hxrt.StringFromLiteral("BYTES"), hxrt.StringFromLiteral("CHARS"), hxrt.StringFromLiteral("URL_BYTES"), hxrt.StringFromLiteral("URL_CHARS"), hxrt.StringFromLiteral("addPadding"), hxrt.StringFromLiteral("decode"), hxrt.StringFromLiteral("encode"), hxrt.StringFromLiteral("removePadding"), hxrt.StringFromLiteral("urlDecode"), hxrt.StringFromLiteral("urlEncode"))
	case "haxe.crypto.Md5":
		return hxrt.NewArray(hxrt.StringFromLiteral("encode"), hxrt.StringFromLiteral("make"))
	case "haxe.crypto.Sha1":
		return hxrt.NewArray(hxrt.StringFromLiteral("encode"), hxrt.StringFromLiteral("make"))
	case "haxe.crypto.Sha224":
		return hxrt.NewArray(hxrt.StringFromLiteral("encode"), hxrt.StringFromLiteral("make"))
	case "haxe.crypto.Sha256":
		return hxrt.NewArray(hxrt.StringFromLiteral("encode"), hxrt.StringFromLiteral("make"))
	case "haxe.ds.StringMap":
		return hxrt.NewArray()
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_fromNativeView"), hxrt.StringFromLiteral("alloc"), hxrt.StringFromLiteral("fastGet"), hxrt.StringFromLiteral("ofData"), hxrt.StringFromLiteral("ofHex"), hxrt.StringFromLiteral("ofString"), hxrt.StringFromLiteral("rawNativeUsesUtf16LE"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray(hxrt.StringFromLiteral("doubleToI64"), hxrt.StringFromLiteral("floatToI32"), hxrt.StringFromLiteral("i32ToFloat"), hxrt.StringFromLiteral("i64ToDouble"))
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray()
	case "haxe.iterators.StringIterator":
		return hxrt.NewArray()
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt.NewArray()
	case "haxe.xml.Parser":
		return hxrt.NewArray(hxrt.StringFromLiteral("doParse"), hxrt.StringFromLiteral("escapes"), hxrt.StringFromLiteral("parse"))
	case "haxe.xml.Printer":
		return hxrt.NewArray(hxrt.StringFromLiteral("print"))
	case "haxe.xml.XmlParserException":
		return hxrt.NewArray()
	case "haxe.xml._Parser.S_Impl_":
		return hxrt.NewArray()
	case "haxe.zip.Compress":
		return hxrt.NewArray(hxrt.StringFromLiteral("fromValues"), hxrt.StringFromLiteral("run"), hxrt.StringFromLiteral("toValues"), hxrt.StringFromLiteral("validateLevel"))
	case "haxe.zip.Uncompress":
		return hxrt.NewArray(hxrt.StringFromLiteral("fromValues"), hxrt.StringFromLiteral("run"), hxrt.StringFromLiteral("toValues"))
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
	case "Std":
		return hxrt.NewArray()
	case "StringBuf":
		return hxrt.NewArray(hxrt.StringFromLiteral("b"))
	case "StringTools":
		return hxrt.NewArray()
	case "Xml":
		return hxrt.NewArray(hxrt.StringFromLiteral("addChild"), hxrt.StringFromLiteral("attributeMap"), hxrt.StringFromLiteral("attributes"), hxrt.StringFromLiteral("children"), hxrt.StringFromLiteral("elements"), hxrt.StringFromLiteral("elementsNamed"), hxrt.StringFromLiteral("ensureElementType"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("firstChild"), hxrt.StringFromLiteral("firstElement"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("get_nodeName"), hxrt.StringFromLiteral("get_nodeValue"), hxrt.StringFromLiteral("insertChild"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("nodeName"), hxrt.StringFromLiteral("nodeType"), hxrt.StringFromLiteral("nodeValue"), hxrt.StringFromLiteral("parent"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeChild"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("set_nodeName"), hxrt.StringFromLiteral("set_nodeValue"), hxrt.StringFromLiteral("toString"))
	case "_Xml.XmlType_Impl_":
		return hxrt.NewArray()
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe.crypto.Base64":
		return hxrt.NewArray()
	case "haxe.crypto.Md5":
		return hxrt.NewArray()
	case "haxe.crypto.Sha1":
		return hxrt.NewArray()
	case "haxe.crypto.Sha224":
		return hxrt.NewArray()
	case "haxe.crypto.Sha256":
		return hxrt.NewArray()
	case "haxe.ds.StringMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_dataExposed"), hxrt.StringFromLiteral("__hx_nativeView"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("blit"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("fill"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getData"), hxrt.StringFromLiteral("getDouble"), hxrt.StringFromLiteral("getFloat"), hxrt.StringFromLiteral("getInt32"), hxrt.StringFromLiteral("getInt64"), hxrt.StringFromLiteral("getString"), hxrt.StringFromLiteral("getUInt16"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setDouble"), hxrt.StringFromLiteral("setFloat"), hxrt.StringFromLiteral("setInt32"), hxrt.StringFromLiteral("setInt64"), hxrt.StringFromLiteral("setUInt16"), hxrt.StringFromLiteral("sub"), hxrt.StringFromLiteral("toHex"), hxrt.StringFromLiteral("toString"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray()
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("next"))
	case "haxe.iterators.StringIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case "haxe.xml.Parser":
		return hxrt.NewArray()
	case "haxe.xml.Printer":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasChildren"), hxrt.StringFromLiteral("newline"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("pretty"), hxrt.StringFromLiteral("write"), hxrt.StringFromLiteral("writeNode"))
	case "haxe.xml.XmlParserException":
		return hxrt.NewArray(hxrt.StringFromLiteral("lineNumber"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("position"), hxrt.StringFromLiteral("positionAtLine"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("xml"))
	case "haxe.xml._Parser.S_Impl_":
		return hxrt.NewArray()
	case "haxe.zip.Compress":
		return hxrt.NewArray(hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("execute"), hxrt.StringFromLiteral("level"), hxrt.StringFromLiteral("setFlushMode"))
	case "haxe.zip.Uncompress":
		return hxrt.NewArray(hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("execute"), hxrt.StringFromLiteral("raw"), hxrt.StringFromLiteral("setFlushMode"))
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
	case "Std":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringBuf":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringTools":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Xml":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "_Xml.XmlType_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.crypto.Base64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.crypto.Md5":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.crypto.Sha1":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.crypto.Sha224":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.crypto.Sha256":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.StringMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Bytes":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.FPHelper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.MapKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.StringIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.StringKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.xml.Parser":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.xml.Printer":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.xml.XmlParserException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.xml._Parser.S_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.zip.Compress":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.zip.Uncompress":
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
	case "haxe.zip.FlushMode":
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
	case *haxe__zip__FlushMode:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("NO")
		case 1:
			return hxrt.StringFromLiteral("SYNC")
		case 2:
			return hxrt.StringFromLiteral("FULL")
		case 3:
			return hxrt.StringFromLiteral("FINISH")
		case 4:
			return hxrt.StringFromLiteral("BLOCK")
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
	case *haxe__zip__FlushMode:
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
	case "haxe.zip.FlushMode":
		return hxrt.NewArray(hxrt.StringFromLiteral("NO"), hxrt.StringFromLiteral("SYNC"), hxrt.StringFromLiteral("FULL"), hxrt.StringFromLiteral("FINISH"), hxrt.StringFromLiteral("BLOCK"))
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
	case *haxe__zip__FlushMode:
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
	case "haxe.zip.FlushMode":
		return hxrt.NewArray(haxe__zip__FlushMode_NO, haxe__zip__FlushMode_SYNC, haxe__zip__FlushMode_FULL, haxe__zip__FlushMode_FINISH, haxe__zip__FlushMode_BLOCK)
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
