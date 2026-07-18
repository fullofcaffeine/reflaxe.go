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

var Demo___rtti *string = hxrt.StringFromLiteral("<class path=\"Demo\" params=\"\" module=\"Main\">\n\t<field public=\"1\" expr=\"&quot;value&quot;\" line=\"6\" static=\"1\">\n\t\t<c path=\"String\"/>\n\t\t<meta><m n=\":value\"><e>\"value\"</e></m></meta>\n\t</field>\n\t<meta>\n\t\t<m n=\":directlyUsed\"/>\n\t\t<m n=\":rtti\"/>\n\t</meta>\n</class>")

var Demo_field *string = hxrt.StringFromLiteral("value")

func main() {
	info := haxe__rtti__Rtti_getRtti(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Demo")})
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("path="), func(hx_obj_1 map[string]any) *string {
		hx_field_2 := hx_obj_1["path"]
		if hx_field_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_field_2.(*string)
	}(info)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("staticType="), haxe__rtti__CTypeTools_toString(func(hx_obj_9 map[string]any) *haxe__rtti__CType {
		hx_field_10 := hx_obj_9["type"]
		if hx_field_10 == nil {
			var hx_zero_11 *haxe__rtti__CType
			return hx_zero_11
		}
		return hx_field_10.(*haxe__rtti__CType)
	}(func(hx_value_7 any) map[string]any {
		if hx_value_7 == nil {
			var hx_zero_8 map[string]any
			return hx_zero_8
		}
		return hx_value_7.(map[string]any)
	}(func(hx_obj_4 map[string]any) *hxrt.Array {
		hx_field_5 := hx_obj_4["statics"]
		if hx_field_5 == nil {
			var hx_zero_6 *hxrt.Array
			return hx_zero_6
		}
		return hx_field_5.(*hxrt.Array)
	}(info).Get(0))))))
	hxrt.Println(v_1)
	var rawRtti any = Reflect_field(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Demo")}, hxrt.StringFromLiteral("__rtti"))
	parsed := New_haxe__rtti__XmlParser().__hx_this.processElement(Xml_parse(hxrt.StdString(rawRtti)).__hx_this.firstElement())
	if parsed.tag == 1 {
		_g := parsed.params[0].(map[string]any)
		c := _g
		var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("parsedPath="), func(hx_obj_12 map[string]any) *string {
			hx_field_13 := hx_obj_12["path"]
			if hx_field_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_field_13.(*string)
		}(c)))
		hxrt.Println(v_2)
	} else {
		hxrt.Println(any(hxrt.StringFromLiteral("parsed=unexpected")))
	}
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := obj.(type) {
	case *Xml:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__rtti__XmlParser:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__xml__Printer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__xml__XmlParserException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Xml:
		return hxrt__generated_method_field__Xml(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_method_field__haxe__ds__StringMap(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_method_field__haxe__io__Bytes(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_method_field__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value, key)
	case *haxe__rtti__XmlParser:
		return hxrt__generated_method_field__haxe__rtti__XmlParser(value, key)
	case *haxe__xml__Printer:
		return hxrt__generated_method_field__haxe__xml__Printer(value, key)
	case *haxe__xml__XmlParserException:
		return hxrt__generated_method_field__haxe__xml__XmlParserException(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Xml(value *Xml, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "addChild":
		return value.addChild
	case "attributes":
		return value.attributes
	case "elements":
		return value.elements
	case "elementsNamed":
		return value.elementsNamed
	case "ensureElementType":
		return value.ensureElementType
	case "exists":
		return value.exists
	case "firstChild":
		return value.firstChild
	case "firstElement":
		return value.firstElement
	case "get":
		return value.get
	case "get_nodeName":
		return value.get_nodeName
	case "get_nodeValue":
		return value.get_nodeValue
	case "insertChild":
		return value.insertChild
	case "iterator":
		return value.iterator
	case "remove":
		return value.remove
	case "removeChild":
		return value.removeChild
	case "set":
		return value.set
	case "set_nodeName":
		return value.set_nodeName
	case "set_nodeValue":
		return value.set_nodeValue
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "clear":
		return value.clear
	case "copy":
		return value.copy
	case "copyIMap":
		return value.copyIMap
	case "exists":
		return value.exists
	case "existsIMap":
		return value.existsIMap
	case "get":
		return value.get
	case "getIMap":
		return value.getIMap
	case "iterator":
		return value.iterator
	case "keyValueIterator":
		return value.keyValueIterator
	case "keys":
		return value.keys
	case "remove":
		return value.remove
	case "removeIMap":
		return value.removeIMap
	case "set":
		return value.set
	case "setIMap":
		return value.setIMap
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_nativeView":
		return value.__hx_nativeView
	case "blit":
		return value.blit
	case "compare":
		return value.compare
	case "fill":
		return value.fill
	case "get":
		return value.get
	case "getData":
		return value.getData
	case "getDouble":
		return value.getDouble
	case "getFloat":
		return value.getFloat
	case "getInt32":
		return value.getInt32
	case "getInt64":
		return value.getInt64
	case "getString":
		return value.getString
	case "getUInt16":
		return value.getUInt16
	case "readString":
		return value.readString
	case "set":
		return value.set
	case "setDouble":
		return value.setDouble
	case "setFloat":
		return value.setFloat
	case "setInt32":
		return value.setInt32
	case "setInt64":
		return value.setInt64
	case "setUInt16":
		return value.setUInt16
	case "sub":
		return value.sub
	case "toHex":
		return value.toHex
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
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

func hxrt__generated_method_field__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
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

func hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
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

func hxrt__generated_method_field__haxe__rtti__XmlParser(value *haxe__rtti__XmlParser, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "defplat":
		return value.defplat
	case "elementName":
		return value.elementName
	case "findSeparator":
		return value.findSeparator
	case "hasNamedElement":
		return value.hasNamedElement
	case "innerData":
		return value.innerData
	case "innerHTML":
		return value.innerHTML
	case "joinStringArray":
		return value.joinStringArray
	case "merge":
		return value.merge
	case "mergeAbstracts":
		return value.mergeAbstracts
	case "mergeClasses":
		return value.mergeClasses
	case "mergeDoc":
		return value.mergeDoc
	case "mergeEnums":
		return value.mergeEnums
	case "mergeFields":
		return value.mergeFields
	case "mergeRights":
		return value.mergeRights
	case "mergeTypedefs":
		return value.mergeTypedefs
	case "mkPath":
		return value.mkPath
	case "mkRights":
		return value.mkRights
	case "mkTypeParams":
		return value.mkTypeParams
	case "nodeDisplayName":
		return value.nodeDisplayName
	case "parseIntString":
		return value.parseIntString
	case "process":
		return value.process
	case "processElement":
		return value.processElement
	case "requireAttr":
		return value.requireAttr
	case "requireFirstElement":
		return value.requireFirstElement
	case "requireNamedElement":
		return value.requireNamedElement
	case "sort":
		return value.sort
	case "sortFields":
		return value.sortFields
	case "splitString":
		return value.splitString
	case "xabstract":
		return value.xabstract
	case "xclass":
		return value.xclass
	case "xclassfield":
		return value.xclassfield
	case "xenum":
		return value.xenum
	case "xenumfield":
		return value.xenumfield
	case "xmeta":
		return value.xmeta
	case "xoverloads":
		return value.xoverloads
	case "xpath":
		return value.xpath
	case "xroot":
		return value.xroot
	case "xtype":
		return value.xtype
	case "xtypedef":
		return value.xtypedef
	case "xtypeparams":
		return value.xtypeparams
	}
	return nil
}

func hxrt__generated_method_field__haxe__xml__Printer(value *haxe__xml__Printer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasChildren":
		return value.hasChildren
	case "newline":
		return value.newline
	case "write":
		return value.write
	case "writeNode":
		return value.writeNode
	}
	return nil
}

func hxrt__generated_method_field__haxe__xml__XmlParserException(value *haxe__xml__XmlParserException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "toString":
		return value.toString
	}
	return nil
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
	generatedMethod := hxrt__generated_method_field(obj, key)
	if generatedMethod != nil {
		return generatedMethod
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
	generatedMethod := hxrt__generated_method_field(obj, key)
	if generatedMethod != nil {
		return true
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
	case "Demo":
		return nil, false
	case "Main":
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
	case "haxe.ds.ArraySort":
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
	case "haxe.rtti.CTypeTools":
		return nil, false
	case "haxe.rtti.Rtti":
		return nil, false
	case "haxe.rtti.TypeApi":
		return nil, false
	case "haxe.rtti.XmlParser":
		return hxrt_typeCallAny(New_haxe__rtti__XmlParser, args)
	case "haxe.xml.Parser":
		return nil, false
	case "haxe.xml.Printer":
		return hxrt_typeCallAny(New_haxe__xml__Printer, args)
	case "haxe.xml.XmlParserException":
		return hxrt_typeCallAny(New_haxe__xml__XmlParserException, args)
	case "haxe.xml._Parser.S_Impl_":
		return nil, false
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
	case "haxe.rtti.XmlParser":
		return &haxe__rtti__XmlParser{}, true
	case "haxe.xml.Printer":
		return &haxe__xml__Printer{}, true
	case "haxe.xml.XmlParserException":
		return &haxe__xml__XmlParserException{}, true
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
	case "haxe.rtti.CType":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__rtti__CType_CUnknown, true
			case 1:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CEnum, args)
			case 2:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CClass, args)
			case 3:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CTypedef, args)
			case 4:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CFunction, args)
			case 5:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CAnonymous, args)
			case 6:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CDynamic, args)
			case 7:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CAbstract, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "CUnknown":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__rtti__CType_CUnknown, true
		case "CEnum":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CEnum, args)
		case "CClass":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CClass, args)
		case "CTypedef":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CTypedef, args)
		case "CFunction":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CFunction, args)
		case "CAnonymous":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CAnonymous, args)
		case "CDynamic":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CDynamic, args)
		case "CAbstract":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CAbstract, args)
		default:
			return nil, false
		}
	case "haxe.rtti.Rights":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__rtti__Rights_RNormal, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__rtti__Rights_RNo, true
			case 2:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__Rights_RCall, args)
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__rtti__Rights_RMethod, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__rtti__Rights_RDynamic, true
			case 5:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__rtti__Rights_RInline, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "RNormal":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__rtti__Rights_RNormal, true
		case "RNo":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__rtti__Rights_RNo, true
		case "RCall":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__Rights_RCall, args)
		case "RMethod":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__rtti__Rights_RMethod, true
		case "RDynamic":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__rtti__Rights_RDynamic, true
		case "RInline":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__rtti__Rights_RInline, true
		default:
			return nil, false
		}
	case "haxe.rtti.TypeTree":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 3 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__TypeTree_TPackage, args)
			case 1:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__TypeTree_TClassdecl, args)
			case 2:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__TypeTree_TEnumdecl, args)
			case 3:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__TypeTree_TTypedecl, args)
			case 4:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__TypeTree_TAbstractdecl, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "TPackage":
			if len(args) != 3 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__TypeTree_TPackage, args)
		case "TClassdecl":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__TypeTree_TClassdecl, args)
		case "TEnumdecl":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__TypeTree_TEnumdecl, args)
		case "TTypedecl":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__TypeTree_TTypedecl, args)
		case "TAbstractdecl":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__TypeTree_TAbstractdecl, args)
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
	case *haxe__rtti__XmlParser:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.rtti.XmlParser")}
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
	case *haxe__rtti__CType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.rtti.CType")}
	case *haxe__rtti__Rights:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.rtti.Rights")}
	case *haxe__rtti__TypeTree:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.rtti.TypeTree")}
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
	case "Demo":
		return nil
	case "Main":
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
	case "haxe.ds.ArraySort":
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
	case "haxe.rtti.CTypeTools":
		return nil
	case "haxe.rtti.Rtti":
		return nil
	case "haxe.rtti.TypeApi":
		return nil
	case "haxe.rtti.XmlParser":
		return nil
	case "haxe.xml.Parser":
		return nil
	case "haxe.xml.Printer":
		return nil
	case "haxe.xml.XmlParserException":
		return nil
	case "haxe.xml._Parser.S_Impl_":
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
	case "Demo":
		return hxrt.NewArray(hxrt.StringFromLiteral("__rtti"), hxrt.StringFromLiteral("field"))
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("main"))
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
	case "haxe.ds.ArraySort":
		return hxrt.NewArray(hxrt.StringFromLiteral("doMerge"), hxrt.StringFromLiteral("gcd"), hxrt.StringFromLiteral("lower"), hxrt.StringFromLiteral("rec"), hxrt.StringFromLiteral("rotate"), hxrt.StringFromLiteral("sort"), hxrt.StringFromLiteral("swap"), hxrt.StringFromLiteral("upper"))
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
	case "haxe.rtti.CTypeTools":
		return hxrt.NewArray(hxrt.StringFromLiteral("classField"), hxrt.StringFromLiteral("functionArgumentName"), hxrt.StringFromLiteral("joinClassFields"), hxrt.StringFromLiteral("joinFunctionArguments"), hxrt.StringFromLiteral("joinStringArray"), hxrt.StringFromLiteral("nameWithParams"), hxrt.StringFromLiteral("toString"))
	case "haxe.rtti.Rtti":
		return hxrt.NewArray(hxrt.StringFromLiteral("getRtti"), hxrt.StringFromLiteral("hasRtti"))
	case "haxe.rtti.TypeApi":
		return hxrt.NewArray(hxrt.StringFromLiteral("constructorEq"), hxrt.StringFromLiteral("fieldEq"), hxrt.StringFromLiteral("isVar"), hxrt.StringFromLiteral("rightsEq"), hxrt.StringFromLiteral("sameClassFields"), hxrt.StringFromLiteral("sameConstructorArguments"), hxrt.StringFromLiteral("sameFunctionArguments"), hxrt.StringFromLiteral("sameTypeParamNames"), hxrt.StringFromLiteral("sameTypes"), hxrt.StringFromLiteral("typeEq"), hxrt.StringFromLiteral("typeInfos"))
	case "haxe.rtti.XmlParser":
		return hxrt.NewArray()
	case "haxe.xml.Parser":
		return hxrt.NewArray(hxrt.StringFromLiteral("doParse"), hxrt.StringFromLiteral("escapes"), hxrt.StringFromLiteral("parse"))
	case "haxe.xml.Printer":
		return hxrt.NewArray(hxrt.StringFromLiteral("print"))
	case "haxe.xml.XmlParserException":
		return hxrt.NewArray()
	case "haxe.xml._Parser.S_Impl_":
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
	case "Demo":
		return hxrt.NewArray()
	case "Main":
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
	case "haxe.ds.ArraySort":
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
	case "haxe.rtti.CTypeTools":
		return hxrt.NewArray()
	case "haxe.rtti.Rtti":
		return hxrt.NewArray()
	case "haxe.rtti.TypeApi":
		return hxrt.NewArray()
	case "haxe.rtti.XmlParser":
		return hxrt.NewArray(hxrt.StringFromLiteral("curplatform"), hxrt.StringFromLiteral("defplat"), hxrt.StringFromLiteral("elementName"), hxrt.StringFromLiteral("findSeparator"), hxrt.StringFromLiteral("hasNamedElement"), hxrt.StringFromLiteral("innerData"), hxrt.StringFromLiteral("innerHTML"), hxrt.StringFromLiteral("joinStringArray"), hxrt.StringFromLiteral("merge"), hxrt.StringFromLiteral("mergeAbstracts"), hxrt.StringFromLiteral("mergeClasses"), hxrt.StringFromLiteral("mergeDoc"), hxrt.StringFromLiteral("mergeEnums"), hxrt.StringFromLiteral("mergeFields"), hxrt.StringFromLiteral("mergeRights"), hxrt.StringFromLiteral("mergeTypedefs"), hxrt.StringFromLiteral("mkPath"), hxrt.StringFromLiteral("mkRights"), hxrt.StringFromLiteral("mkTypeParams"), hxrt.StringFromLiteral("newField"), hxrt.StringFromLiteral("nodeDisplayName"), hxrt.StringFromLiteral("parseIntString"), hxrt.StringFromLiteral("process"), hxrt.StringFromLiteral("processElement"), hxrt.StringFromLiteral("requireAttr"), hxrt.StringFromLiteral("requireFirstElement"), hxrt.StringFromLiteral("requireNamedElement"), hxrt.StringFromLiteral("root"), hxrt.StringFromLiteral("sort"), hxrt.StringFromLiteral("sortFields"), hxrt.StringFromLiteral("splitString"), hxrt.StringFromLiteral("xabstract"), hxrt.StringFromLiteral("xclass"), hxrt.StringFromLiteral("xclassfield"), hxrt.StringFromLiteral("xenum"), hxrt.StringFromLiteral("xenumfield"), hxrt.StringFromLiteral("xmeta"), hxrt.StringFromLiteral("xoverloads"), hxrt.StringFromLiteral("xpath"), hxrt.StringFromLiteral("xroot"), hxrt.StringFromLiteral("xtype"), hxrt.StringFromLiteral("xtypedef"), hxrt.StringFromLiteral("xtypeparams"))
	case "haxe.xml.Parser":
		return hxrt.NewArray()
	case "haxe.xml.Printer":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasChildren"), hxrt.StringFromLiteral("newline"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("pretty"), hxrt.StringFromLiteral("write"), hxrt.StringFromLiteral("writeNode"))
	case "haxe.xml.XmlParserException":
		return hxrt.NewArray(hxrt.StringFromLiteral("lineNumber"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("position"), hxrt.StringFromLiteral("positionAtLine"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("xml"))
	case "haxe.xml._Parser.S_Impl_":
		return hxrt.NewArray()
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
	case "Demo":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Main":
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
	case "haxe.ds.ArraySort":
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
	case "haxe.rtti.CTypeTools":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.rtti.Rtti":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.rtti.TypeApi":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.rtti.XmlParser":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.xml.Parser":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.xml.Printer":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.xml.XmlParserException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.xml._Parser.S_Impl_":
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
	case "haxe.rtti.CType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.rtti.Rights":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.rtti.TypeTree":
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
	case *haxe__rtti__CType:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("CUnknown")
		case 1:
			return hxrt.StringFromLiteral("CEnum")
		case 2:
			return hxrt.StringFromLiteral("CClass")
		case 3:
			return hxrt.StringFromLiteral("CTypedef")
		case 4:
			return hxrt.StringFromLiteral("CFunction")
		case 5:
			return hxrt.StringFromLiteral("CAnonymous")
		case 6:
			return hxrt.StringFromLiteral("CDynamic")
		case 7:
			return hxrt.StringFromLiteral("CAbstract")
		default:
			return nil
		}
	case *haxe__rtti__Rights:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("RNormal")
		case 1:
			return hxrt.StringFromLiteral("RNo")
		case 2:
			return hxrt.StringFromLiteral("RCall")
		case 3:
			return hxrt.StringFromLiteral("RMethod")
		case 4:
			return hxrt.StringFromLiteral("RDynamic")
		case 5:
			return hxrt.StringFromLiteral("RInline")
		default:
			return nil
		}
	case *haxe__rtti__TypeTree:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("TPackage")
		case 1:
			return hxrt.StringFromLiteral("TClassdecl")
		case 2:
			return hxrt.StringFromLiteral("TEnumdecl")
		case 3:
			return hxrt.StringFromLiteral("TTypedecl")
		case 4:
			return hxrt.StringFromLiteral("TAbstractdecl")
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
	case *haxe__rtti__CType:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__rtti__Rights:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__rtti__TypeTree:
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
	case "haxe.rtti.CType":
		return hxrt.NewArray(hxrt.StringFromLiteral("CUnknown"), hxrt.StringFromLiteral("CEnum"), hxrt.StringFromLiteral("CClass"), hxrt.StringFromLiteral("CTypedef"), hxrt.StringFromLiteral("CFunction"), hxrt.StringFromLiteral("CAnonymous"), hxrt.StringFromLiteral("CDynamic"), hxrt.StringFromLiteral("CAbstract"))
	case "haxe.rtti.Rights":
		return hxrt.NewArray(hxrt.StringFromLiteral("RNormal"), hxrt.StringFromLiteral("RNo"), hxrt.StringFromLiteral("RCall"), hxrt.StringFromLiteral("RMethod"), hxrt.StringFromLiteral("RDynamic"), hxrt.StringFromLiteral("RInline"))
	case "haxe.rtti.TypeTree":
		return hxrt.NewArray(hxrt.StringFromLiteral("TPackage"), hxrt.StringFromLiteral("TClassdecl"), hxrt.StringFromLiteral("TEnumdecl"), hxrt.StringFromLiteral("TTypedecl"), hxrt.StringFromLiteral("TAbstractdecl"))
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
	case *haxe__rtti__CType:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__rtti__Rights:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__rtti__TypeTree:
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
	case "haxe.rtti.CType":
		return hxrt.NewArray(haxe__rtti__CType_CUnknown)
	case "haxe.rtti.Rights":
		return hxrt.NewArray(haxe__rtti__Rights_RNormal, haxe__rtti__Rights_RNo, haxe__rtti__Rights_RMethod, haxe__rtti__Rights_RDynamic, haxe__rtti__Rights_RInline)
	case "haxe.rtti.TypeTree":
		return hxrt.NewArray()
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
	case "Demo":
		switch key {
		case "__rtti":
			return Demo___rtti, true
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}
