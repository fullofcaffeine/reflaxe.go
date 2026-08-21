package main

import "examples_portable_beta/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type _Main__DeliveryState struct {
	tag    int
	params []any
}

var _Main__DeliveryState_Queued *_Main__DeliveryState = &_Main__DeliveryState{tag: 0}

var _Main__DeliveryState_Delivered *_Main__DeliveryState = &_Main__DeliveryState{tag: 1}

var SCRATCH_FILE *string = hxrt.StringFromLiteral(".portable_beta_contract.txt")

func checkExceptionContract() {
	observed := hxrt.StringFromLiteral("")
	hxrt.TryCatch(func() {
		hxrt.Throw(hxrt.StringFromLiteral("portable-beta-error"))
	}, func(hx_caught_1 any) {
		error := hxrt.ExceptionCaught(hx_caught_1)
		observed = hxrt.ExceptionMessage(error)
	})
	require(hxrt.StringEqualStringPtr(observed, hxrt.StringFromLiteral("portable-beta-error")), hxrt.StringFromLiteral("typed exception message contract"))
}

func checkFileLifecycle() {
	if sys__FileSystem_exists(hxrt.StringFromLiteral(".portable_beta_contract.txt")) {
		sys__FileSystem_deleteFile(hxrt.StringFromLiteral(".portable_beta_contract.txt"))
	}
	hxrt.TryCatch(func() {
		sys__io__File_saveContent(hxrt.StringFromLiteral(".portable_beta_contract.txt"), hxrt.StringFromLiteral("portable-beta-ok"))
		require(sys__FileSystem_exists(hxrt.StringFromLiteral(".portable_beta_contract.txt")), hxrt.StringFromLiteral("file existence contract"))
		require(hxrt.StringEqualStringPtr(sys__io__File_getContent(hxrt.StringFromLiteral(".portable_beta_contract.txt")), hxrt.StringFromLiteral("portable-beta-ok")), hxrt.StringFromLiteral("file content contract"))
		sys__FileSystem_deleteFile(hxrt.StringFromLiteral(".portable_beta_contract.txt"))
	}, func(hx_caught_3 any) {
		error := hxrt.ExceptionCaught(hx_caught_3)
		if sys__FileSystem_exists(hxrt.StringFromLiteral(".portable_beta_contract.txt")) {
			sys__FileSystem_deleteFile(hxrt.StringFromLiteral(".portable_beta_contract.txt"))
		}
		hxrt.Throw(error)
	})
	require(!sys__FileSystem_exists(hxrt.StringFromLiteral(".portable_beta_contract.txt")), hxrt.StringFromLiteral("file cleanup contract"))
}

func checkLanguageAndCollections() {
	deliveries := hxrt.NewArray()
	deliveries.Push(New__Main__Delivery(hxrt.StringFromLiteral("alpha"), _Main__DeliveryState_Delivered))
	deliveries.Push(New__Main__Delivery(hxrt.StringFromLiteral("beta"), _Main__DeliveryState_Queued))
	deliveries.Push(New__Main__Delivery(hxrt.StringFromLiteral("gamma"), _Main__DeliveryState_Delivered))
	delivered := 0
	_g := 0
	_g1 := deliveries.Len()
	for _g < _g1 {
		hx_post_8 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_8
		_g_1 := func(hx_value_9 any) *_Main__Delivery {
			if hx_value_9 == nil {
				var hx_zero_10 *_Main__Delivery
				return hx_zero_10
			}
			return hx_value_9.(*_Main__Delivery)
		}(deliveries.Get(index)).state
		switch _g_1.tag {
		case 0:
		case 1:
			delivered = int(int32((delivered + 1)))
		}
	}
	require((delivered == 2), hxrt.StringFromLiteral("enum and numeric-for contract"))
	removed := func() *_Main__Delivery {
		return func(hx_value_12 any) *_Main__Delivery {
			if hx_value_12 == nil {
				var hx_zero_13 *_Main__Delivery
				return hx_zero_13
			}
			return hx_value_12.(*_Main__Delivery)
		}(deliveries.Pop())
	}()
	require((removed != nil), hxrt.StringFromLiteral("array pop contract"))
	if removed != nil {
		require(hxrt.StringEqualStringPtr(removed.name, hxrt.StringFromLiteral("gamma")), hxrt.StringFromLiteral("array order contract"))
	}
	require((deliveries.Len() == 2), hxrt.StringFromLiteral("array length contract"))
	decorate := func(value *string) *string {
		return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), value), hxrt.StringFromLiteral("]"))
	}
	require(hxrt.StringEqualStringPtr(decorate(hxrt.StringFromLiteral("portable")), hxrt.StringFromLiteral("[portable]")), hxrt.StringFromLiteral("closure contract"))
}

func checkTextAndData() {
	cleaned := StringTools_trim(hxrt.StringFromLiteral("  alpha-go  "))
	require((hxrt.StringLengthStringPtr(cleaned) == 8), hxrt.StringFromLiteral("string length contract"))
	require(hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(cleaned, 5), hxrt.StringFromLiteral("-")), hxrt.StringFromLiteral("string charAt contract"))
	require(hxrt.StringEqualStringPtr(hxrt.StringSubstringStringPtr(cleaned, 0, 5), hxrt.StringFromLiteral("alpha")), hxrt.StringFromLiteral("string substring contract"))
	require(StringTools_startsWith(cleaned, hxrt.StringFromLiteral("alpha")), hxrt.StringFromLiteral("startsWith contract"))
	require(StringTools_endsWith(cleaned, hxrt.StringFromLiteral("go")), hxrt.StringFromLiteral("endsWith contract"))
	require(StringTools_contains(cleaned, hxrt.StringFromLiteral("ha-g")), hxrt.StringFromLiteral("contains contract"))
	require(hxrt.StringEqualStringPtr(StringTools_replace(cleaned, hxrt.StringFromLiteral("-"), hxrt.StringFromLiteral(" ")), hxrt.StringFromLiteral("alpha go")), hxrt.StringFromLiteral("replace contract"))
	bytes := haxe__io__Bytes_ofString(cleaned, nil)
	require(hxrt.StringEqualStringPtr(bytes.__hx_this.toString(), cleaned), hxrt.StringFromLiteral("bytes round-trip contract"))
	names := New_haxe__ds__StringMap()
	names.__hx_this.set(hxrt.StringFromLiteral("primary"), hxrt.StringFromLiteral("alpha"))
	require(func(hx_value_14 any) bool {
		if hx_value_14 == nil {
			var hx_zero_15 bool
			return hx_zero_15
		}
		return hx_value_14.(bool)
	}(names.__hx_this.exists(hxrt.StringFromLiteral("primary"))), hxrt.StringFromLiteral("StringMap exists contract"))
	require(hxrt.StringEqualStringPtr(func(hx_value_16 any) *string {
		if hx_value_16 == nil {
			var hx_zero_17 *string
			return hx_zero_17
		}
		return hx_value_16.(*string)
	}(names.__hx_this.get(hxrt.StringFromLiteral("primary"))), hxrt.StringFromLiteral("alpha")), hxrt.StringFromLiteral("StringMap get contract"))
	require(func(hx_value_18 any) bool {
		if hx_value_18 == nil {
			var hx_zero_19 bool
			return hx_zero_19
		}
		return hx_value_18.(bool)
	}(names.__hx_this.remove(hxrt.StringFromLiteral("primary"))), hxrt.StringFromLiteral("StringMap remove contract"))
	require(!func(hx_value_20 any) bool {
		if hx_value_20 == nil {
			var hx_zero_21 bool
			return hx_zero_21
		}
		return hx_value_20.(bool)
	}(names.__hx_this.exists(hxrt.StringFromLiteral("primary"))), hxrt.StringFromLiteral("StringMap removal contract"))
	var parsed any = hxrt.JsonParse(hxrt.StringFromLiteral("{\"name\":\"alpha\",\"ready\":true}"))
	require(Reflect_hasField(parsed, hxrt.StringFromLiteral("name")), hxrt.StringFromLiteral("JSON field contract"))
	parsedName := func(hx_value_22 any) *string {
		if hx_value_22 == nil {
			var hx_zero_23 *string
			return hx_zero_23
		}
		return hx_value_22.(*string)
	}(Reflect_field(parsed, hxrt.StringFromLiteral("name")))
	require(hxrt.StringEqualStringPtr(parsedName, hxrt.StringFromLiteral("alpha")), hxrt.StringFromLiteral("JSON value contract"))
	var space *string = nil
	encoded := hxrt.StdString(hxrt.JsonStringify(any(func() map[string]any {
		hx_obj_24 := map[string]any{}
		hx_obj_24["name"] = hxrt.StringFromLiteral("alpha")
		return hx_obj_24
	}()), space))
	require(StringTools_contains(encoded, hxrt.StringFromLiteral("\"name\":\"alpha\"")), hxrt.StringFromLiteral("JSON stringify contract"))
	path := New_haxe__io__Path(haxe__io__Path_join(hxrt.NewArray(hxrt.StringFromLiteral("reports"), hxrt.StringFromLiteral("result.json"))))
	require(hxrt.StringEqualStringPtr(path.dir, hxrt.StringFromLiteral("reports")), hxrt.StringFromLiteral("Path directory contract"))
	require(hxrt.StringEqualStringPtr(path.file, hxrt.StringFromLiteral("result")), hxrt.StringFromLiteral("Path file contract"))
	require(hxrt.StringEqualStringPtr(path.ext, hxrt.StringFromLiteral("json")), hxrt.StringFromLiteral("Path extension contract"))
}

func main() {
	checkLanguageAndCollections()
	checkTextAndData()
	checkFileLifecycle()
	checkExceptionContract()
}

func require(condition bool, message *string) {
	if !condition {
		hxrt.Throw(message)
	}
}

type I__Main__Delivery interface {
}

type _Main__Delivery struct {
	__hx_this I__Main__Delivery
	name      *string
	state     *_Main__DeliveryState
}

func New__Main__Delivery(name *string, state *_Main__DeliveryState) *_Main__Delivery {
	self := &_Main__Delivery{}
	self.__hx_this = self
	self.name = name
	self.state = state
	return self
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := obj.(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__exceptions__NotImplementedException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__exceptions__PosException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__BytesBuffer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Eof:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Input:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Output:
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
	case *sys__io__FileInput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__io__FileOutput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Date:
		return hxrt__generated_method_field__Date(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_method_field__haxe__ds__StringMap(value, key)
	case *haxe__exceptions__NotImplementedException:
		return hxrt__generated_method_field__haxe__exceptions__NotImplementedException(value, key)
	case *haxe__exceptions__PosException:
		return hxrt__generated_method_field__haxe__exceptions__PosException(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_method_field__haxe__io__Bytes(value, key)
	case *haxe__io__BytesBuffer:
		return hxrt__generated_method_field__haxe__io__BytesBuffer(value, key)
	case *haxe__io__Eof:
		return hxrt__generated_method_field__haxe__io__Eof(value, key)
	case *haxe__io__Input:
		return hxrt__generated_method_field__haxe__io__Input(value, key)
	case *haxe__io__Output:
		return hxrt__generated_method_field__haxe__io__Output(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_method_field__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value, key)
	case *sys__io__FileInput:
		return hxrt__generated_method_field__sys__io__FileInput(value, key)
	case *sys__io__FileOutput:
		return hxrt__generated_method_field__sys__io__FileOutput(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Date(value *Date, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "getDate":
		return value.getDate
	case "getDay":
		return value.getDay
	case "getFullYear":
		return value.getFullYear
	case "getHours":
		return value.getHours
	case "getMinutes":
		return value.getMinutes
	case "getMonth":
		return value.getMonth
	case "getSeconds":
		return value.getSeconds
	case "getTime":
		return value.getTime
	case "getTimezoneOffset":
		return value.getTimezoneOffset
	case "getUTCDate":
		return value.getUTCDate
	case "getUTCDay":
		return value.getUTCDay
	case "getUTCFullYear":
		return value.getUTCFullYear
	case "getUTCHours":
		return value.getUTCHours
	case "getUTCMinutes":
		return value.getUTCMinutes
	case "getUTCMonth":
		return value.getUTCMonth
	case "getUTCSeconds":
		return value.getUTCSeconds
	case "localParts":
		return value.localParts
	case "toString":
		return value.toString
	case "utcParts":
		return value.utcParts
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

func hxrt__generated_method_field__haxe__exceptions__NotImplementedException(value *haxe__exceptions__NotImplementedException, key string) any {
	if value == nil {
		return nil
	}
	if value.haxe__exceptions__PosException == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__exceptions__PosException(value.haxe__exceptions__PosException, key)
}

func hxrt__generated_method_field__haxe__exceptions__PosException(value *haxe__exceptions__PosException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
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

func hxrt__generated_method_field__haxe__io__BytesBuffer(value *haxe__io__BytesBuffer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "add":
		return value.add
	case "addByte":
		return value.addByte
	case "addBytes":
		return value.addBytes
	case "addDouble":
		return value.addDouble
	case "addFloat":
		return value.addFloat
	case "addInt32":
		return value.addInt32
	case "addInt64":
		return value.addInt64
	case "addString":
		return value.addString
	case "getBytes":
		return value.getBytes
	case "get_length":
		return value.get_length
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Eof(value *haxe__io__Eof, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Input(value *haxe__io__Input, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return value.close
	case "read":
		return value.read
	case "readAll":
		return value.readAll
	case "readByte":
		return value.readByte
	case "readBytes":
		return value.readBytes
	case "readDouble":
		return value.readDouble
	case "readFloat":
		return value.readFloat
	case "readFullBytes":
		return value.readFullBytes
	case "readInt16":
		return value.readInt16
	case "readInt24":
		return value.readInt24
	case "readInt32":
		return value.readInt32
	case "readInt8":
		return value.readInt8
	case "readLine":
		return value.readLine
	case "readString":
		return value.readString
	case "readUInt16":
		return value.readUInt16
	case "readUInt24":
		return value.readUInt24
	case "readUntil":
		return value.readUntil
	case "set_bigEndian":
		return value.set_bigEndian
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Output(value *haxe__io__Output, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return value.close
	case "flush":
		return value.flush
	case "prepare":
		return value.prepare
	case "set_bigEndian":
		return value.set_bigEndian
	case "write":
		return value.write
	case "writeByte":
		return value.writeByte
	case "writeBytes":
		return value.writeBytes
	case "writeDouble":
		return value.writeDouble
	case "writeFloat":
		return value.writeFloat
	case "writeFullBytes":
		return value.writeFullBytes
	case "writeInput":
		return value.writeInput
	case "writeInt16":
		return value.writeInt16
	case "writeInt24":
		return value.writeInt24
	case "writeInt32":
		return value.writeInt32
	case "writeInt8":
		return value.writeInt8
	case "writeString":
		return value.writeString
	case "writeUInt16":
		return value.writeUInt16
	case "writeUInt24":
		return value.writeUInt24
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

func hxrt__generated_method_field__sys__io__FileInput(value *sys__io__FileInput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return value.close
	case "eof":
		return value.eof
	case "readByte":
		return value.readByte
	case "readBytes":
		return value.readBytes
	case "seek":
		return value.seek
	case "tell":
		return value.tell
	}
	if value.haxe__io__Input == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__io__Input(value.haxe__io__Input, key)
}

func hxrt__generated_method_field__sys__io__FileOutput(value *sys__io__FileOutput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return value.close
	case "flush":
		return value.flush
	case "seek":
		return value.seek
	case "tell":
		return value.tell
	case "writeByte":
		return value.writeByte
	case "writeBytes":
		return value.writeBytes
	}
	if value.haxe__io__Output == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__io__Output(value.haxe__io__Output, key)
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
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *_Main__Delivery:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__exceptions__NotImplementedException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__exceptions__PosException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__BytesBuffer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Input:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Output:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Path:
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
	case *sys__io__FileInput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__io__FileOutput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Date:
		return hxrt__generated_field_lookup__Date(value, key)
	case *_Main__Delivery:
		return hxrt__generated_field_lookup___Main__Delivery(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_lookup__haxe__ds__StringMap(value, key)
	case *haxe__exceptions__NotImplementedException:
		return hxrt__generated_field_lookup__haxe__exceptions__NotImplementedException(value, key)
	case *haxe__exceptions__PosException:
		return hxrt__generated_field_lookup__haxe__exceptions__PosException(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_lookup__haxe__io__Bytes(value, key)
	case *haxe__io__BytesBuffer:
		return hxrt__generated_field_lookup__haxe__io__BytesBuffer(value, key)
	case *haxe__io__Input:
		return hxrt__generated_field_lookup__haxe__io__Input(value, key)
	case *haxe__io__Output:
		return hxrt__generated_field_lookup__haxe__io__Output(value, key)
	case *haxe__io__Path:
		return hxrt__generated_field_lookup__haxe__io__Path(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringKeyValueIterator(value, key)
	case *sys__io__FileInput:
		return hxrt__generated_field_lookup__sys__io__FileInput(value, key)
	case *sys__io__FileOutput:
		return hxrt__generated_field_lookup__sys__io__FileOutput(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *_Main__Delivery:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__exceptions__NotImplementedException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__exceptions__PosException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__BytesBuffer:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Input:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Output:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Path:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__io__FileInput:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__io__FileOutput:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Date:
		return hxrt__generated_field_has__Date(value, key)
	case *_Main__Delivery:
		return hxrt__generated_field_has___Main__Delivery(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_has__haxe__ds__StringMap(value, key)
	case *haxe__exceptions__NotImplementedException:
		return hxrt__generated_field_has__haxe__exceptions__NotImplementedException(value, key)
	case *haxe__exceptions__PosException:
		return hxrt__generated_field_has__haxe__exceptions__PosException(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_has__haxe__io__Bytes(value, key)
	case *haxe__io__BytesBuffer:
		return hxrt__generated_field_has__haxe__io__BytesBuffer(value, key)
	case *haxe__io__Input:
		return hxrt__generated_field_has__haxe__io__Input(value, key)
	case *haxe__io__Output:
		return hxrt__generated_field_has__haxe__io__Output(value, key)
	case *haxe__io__Path:
		return hxrt__generated_field_has__haxe__io__Path(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_has__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__StringKeyValueIterator(value, key)
	case *sys__io__FileInput:
		return hxrt__generated_field_has__sys__io__FileInput(value, key)
	case *sys__io__FileOutput:
		return hxrt__generated_field_has__sys__io__FileOutput(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *_Main__Delivery:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__exceptions__NotImplementedException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__exceptions__PosException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__BytesBuffer:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Input:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Output:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Path:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__io__FileInput:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__io__FileOutput:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Date:
		return hxrt__generated_field_set__Date(value, key, incoming)
	case *_Main__Delivery:
		return hxrt__generated_field_set___Main__Delivery(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_set__haxe__ds__StringMap(value, key, incoming)
	case *haxe__exceptions__NotImplementedException:
		return hxrt__generated_field_set__haxe__exceptions__NotImplementedException(value, key, incoming)
	case *haxe__exceptions__PosException:
		return hxrt__generated_field_set__haxe__exceptions__PosException(value, key, incoming)
	case *haxe__io__Bytes:
		return hxrt__generated_field_set__haxe__io__Bytes(value, key, incoming)
	case *haxe__io__BytesBuffer:
		return hxrt__generated_field_set__haxe__io__BytesBuffer(value, key, incoming)
	case *haxe__io__Input:
		return hxrt__generated_field_set__haxe__io__Input(value, key, incoming)
	case *haxe__io__Output:
		return hxrt__generated_field_set__haxe__io__Output(value, key, incoming)
	case *haxe__io__Path:
		return hxrt__generated_field_set__haxe__io__Path(value, key, incoming)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value, key, incoming)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_set__haxe__iterators__StringIterator(value, key, incoming)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__StringKeyValueIterator(value, key, incoming)
	case *sys__io__FileInput:
		return hxrt__generated_field_set__sys__io__FileInput(value, key, incoming)
	case *sys__io__FileOutput:
		return hxrt__generated_field_set__sys__io__FileOutput(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := object.(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *_Main__Delivery:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__exceptions__NotImplementedException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__exceptions__PosException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__BytesBuffer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Eof:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Input:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Output:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Path:
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
	case *sys__io__FileInput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__io__FileOutput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *Date:
		return hxrt.NewArray(hxrt.StringFromLiteral("ms"))
	case *_Main__Delivery:
		return hxrt.NewArray(hxrt.StringFromLiteral("name"), hxrt.StringFromLiteral("state"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__ds__StringMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("h"))
	case *haxe__exceptions__NotImplementedException:
		return hxrt.NewArray(hxrt.StringFromLiteral("posInfos"))
	case *haxe__exceptions__PosException:
		return hxrt.NewArray(hxrt.StringFromLiteral("posInfos"))
	case *haxe__io__Bytes:
		return hxrt.NewArray(hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("__hx_dataExposed"))
	case *haxe__io__BytesBuffer:
		return hxrt.NewArray(hxrt.StringFromLiteral("b"))
	case *haxe__io__Eof:
		return hxrt.NewArray()
	case *haxe__io__Input:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"))
	case *haxe__io__Output:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"))
	case *haxe__io__Path:
		return hxrt.NewArray(hxrt.StringFromLiteral("dir"), hxrt.StringFromLiteral("file"), hxrt.StringFromLiteral("ext"), hxrt.StringFromLiteral("backslash"))
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("keys"))
	case *haxe__iterators__StringIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *sys__io__FileInput:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("handle"))
	case *sys__io__FileOutput:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("handle"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__Date(value *Date, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "ms":
		return value.ms
	}
	return nil
}

func hxrt__generated_field_has__Date(value *Date, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "ms":
		return true
	}
	return false
}

func hxrt__generated_field_set__Date(value *Date, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "ms":
		if incoming == nil {
			var zero float64
			value.ms = zero
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.ms = typed
			return true
		case int:
			value.ms = float64(typed)
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup___Main__Delivery(value *_Main__Delivery, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "name":
		return value.name
	case "state":
		return value.state
	}
	return nil
}

func hxrt__generated_field_has___Main__Delivery(value *_Main__Delivery, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "name":
		return true
	case "state":
		return true
	}
	return false
}

func hxrt__generated_field_set___Main__Delivery(value *_Main__Delivery, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "name":
		if incoming == nil {
			var zero *string
			value.name = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.name = typed
			return true
		default:
			return false
		}
	case "state":
		if incoming == nil {
			var zero *_Main__DeliveryState
			value.state = zero
			return true
		}
		switch typed := incoming.(type) {
		case *_Main__DeliveryState:
			value.state = typed
			return true
		default:
			return false
		}
	}
	return false
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

func hxrt__generated_field_lookup__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "h":
		return value.h
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__StringMap(value *haxe__ds__StringMap, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		if incoming == nil {
			var zero *hxrt.StringMapCell
			value.h = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.StringMapCell:
			value.h = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__exceptions__NotImplementedException(value *haxe__exceptions__NotImplementedException, key string) any {
	if value == nil {
		return nil
	}
	if value.haxe__exceptions__PosException == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__exceptions__PosException(value.haxe__exceptions__PosException, key)
}

func hxrt__generated_field_has__haxe__exceptions__NotImplementedException(value *haxe__exceptions__NotImplementedException, key string) bool {
	if value == nil {
		return false
	}
	if value.haxe__exceptions__PosException == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__exceptions__PosException(value.haxe__exceptions__PosException, key)
}

func hxrt__generated_field_set__haxe__exceptions__NotImplementedException(value *haxe__exceptions__NotImplementedException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	if value.haxe__exceptions__PosException == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__exceptions__PosException(value.haxe__exceptions__PosException, key, incoming)
}

func hxrt__generated_field_lookup__haxe__exceptions__PosException(value *haxe__exceptions__PosException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "posInfos":
		return value.posInfos
	}
	return nil
}

func hxrt__generated_field_has__haxe__exceptions__PosException(value *haxe__exceptions__PosException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "posInfos":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__exceptions__PosException(value *haxe__exceptions__PosException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "posInfos":
		if incoming == nil {
			var zero map[string]any
			value.posInfos = zero
			return true
		}
		switch typed := incoming.(type) {
		case map[string]any:
			value.posInfos = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_dataExposed":
		return value.__hx_dataExposed
	case "__hx_raw":
		return value.__hx_raw
	case "__hx_rawValid":
		return value.__hx_rawValid
	case "b":
		return value.b
	case "length":
		return value.length
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__Bytes(value *haxe__io__Bytes, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__hx_dataExposed":
		return true
	case "__hx_raw":
		return true
	case "__hx_rawValid":
		return true
	case "b":
		return true
	case "length":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__Bytes(value *haxe__io__Bytes, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__hx_dataExposed":
		if incoming == nil {
			var zero bool
			value.__hx_dataExposed = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.__hx_dataExposed = typed
			return true
		default:
			return false
		}
	case "__hx_raw":
		if incoming == nil {
			var zero *hxrt.ByteView
			value.__hx_raw = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.ByteView:
			value.__hx_raw = typed
			return true
		default:
			return false
		}
	case "__hx_rawValid":
		if incoming == nil {
			var zero bool
			value.__hx_rawValid = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.__hx_rawValid = typed
			return true
		default:
			return false
		}
	case "b":
		if incoming == nil {
			var zero []int
			value.b = zero
			return true
		}
		switch typed := incoming.(type) {
		case []int:
			value.b = typed
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

func hxrt__generated_field_lookup__haxe__io__BytesBuffer(value *haxe__io__BytesBuffer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "b":
		return value.b
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__BytesBuffer(value *haxe__io__BytesBuffer, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "b":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__BytesBuffer(value *haxe__io__BytesBuffer, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "b":
		if incoming == nil {
			var zero []int
			value.b = zero
			return true
		}
		switch typed := incoming.(type) {
		case []int:
			value.b = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__io__Input(value *haxe__io__Input, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "bigEndian":
		return value.bigEndian
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__Input(value *haxe__io__Input, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "bigEndian":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__Input(value *haxe__io__Input, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "bigEndian":
		if incoming == nil {
			var zero bool
			value.bigEndian = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.bigEndian = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__io__Output(value *haxe__io__Output, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "bigEndian":
		return value.bigEndian
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__Output(value *haxe__io__Output, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "bigEndian":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__Output(value *haxe__io__Output, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "bigEndian":
		if incoming == nil {
			var zero bool
			value.bigEndian = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.bigEndian = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__io__Path(value *haxe__io__Path, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "backslash":
		return value.backslash
	case "dir":
		return value.dir
	case "ext":
		return value.ext
	case "file":
		return value.file
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__Path(value *haxe__io__Path, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "backslash":
		return true
	case "dir":
		return true
	case "ext":
		return true
	case "file":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__Path(value *haxe__io__Path, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "backslash":
		if incoming == nil {
			var zero bool
			value.backslash = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.backslash = typed
			return true
		default:
			return false
		}
	case "dir":
		if incoming == nil {
			var zero *string
			value.dir = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.dir = typed
			return true
		default:
			return false
		}
	case "ext":
		if incoming == nil {
			var zero *string
			value.ext = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.ext = typed
			return true
		default:
			return false
		}
	case "file":
		if incoming == nil {
			var zero *string
			value.file = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.file = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "keys":
		return value.keys
	case "map":
		return value.map_
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		return true
	case "map":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		if incoming == nil {
			var zero map[string]any
			value.keys = zero
			return true
		}
		switch typed := incoming.(type) {
		case map[string]any:
			value.keys = typed
			return true
		default:
			return false
		}
	case "map":
		if incoming == nil {
			var zero haxe__IMap
			value.map_ = zero
			return true
		}
		switch typed := incoming.(type) {
		case haxe__IMap:
			value.map_ = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "offset":
		return value.offset
	case "s":
		return value.s
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		return true
	case "s":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		if incoming == nil {
			var zero int
			value.offset = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.offset = typed
			return true
		default:
			return false
		}
	case "s":
		if incoming == nil {
			var zero *string
			value.s = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.s = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "offset":
		return value.offset
	case "s":
		return value.s
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		return true
	case "s":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		if incoming == nil {
			var zero int
			value.offset = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.offset = typed
			return true
		default:
			return false
		}
	case "s":
		if incoming == nil {
			var zero *string
			value.s = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.s = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__io__FileInput(value *sys__io__FileInput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "handle":
		return value.handle
	}
	if value.haxe__io__Input == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__io__Input(value.haxe__io__Input, key)
}

func hxrt__generated_field_has__sys__io__FileInput(value *sys__io__FileInput, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "handle":
		return true
	}
	if value.haxe__io__Input == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__io__Input(value.haxe__io__Input, key)
}

func hxrt__generated_field_set__sys__io__FileInput(value *sys__io__FileInput, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "handle":
		if incoming == nil {
			var zero *hxrt.FileInput
			value.handle = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.FileInput:
			value.handle = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__io__Input == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__io__Input(value.haxe__io__Input, key, incoming)
}

func hxrt__generated_field_lookup__sys__io__FileOutput(value *sys__io__FileOutput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "handle":
		return value.handle
	}
	if value.haxe__io__Output == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__io__Output(value.haxe__io__Output, key)
}

func hxrt__generated_field_has__sys__io__FileOutput(value *sys__io__FileOutput, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "handle":
		return true
	}
	if value.haxe__io__Output == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__io__Output(value.haxe__io__Output, key)
}

func hxrt__generated_field_set__sys__io__FileOutput(value *sys__io__FileOutput, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "handle":
		if incoming == nil {
			var zero *hxrt.FileOutput
			value.handle = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.FileOutput:
			value.handle = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__io__Output == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__io__Output(value.haxe__io__Output, key, incoming)
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
	case *_Main__DeliveryState:
		return (enumValue != nil)
	case *haxe__io__Encoding:
		return (enumValue != nil)
	case *haxe__io__Error:
		return (enumValue != nil)
	case *sys__io__FileSeek:
		return (enumValue != nil)
	default:
		return false
	}
}
