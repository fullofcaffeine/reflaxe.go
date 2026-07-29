package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func main() {
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,hello%20from%20haxe.go"))
	sink := New_haxe__io__BytesOutput()
	http.__hx_this.customRequest(false, sink.haxe__io__Output, nil, nil)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("custom="), sink.__hx_this.getBytes().__hx_this.toString()))
	hxrt.Println(v)
	values := http.__hx_this.getResponseHeaderValues(hxrt.StringFromLiteral("Content-Type"))
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("headers="), func() int {
		var hx_if_1 int
		if values == nil {
			hx_if_1 = -1
		} else {
			hx_if_1 = values.Len()
		}
		return hx_if_1
	}()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("header0="), func() *string {
		var hx_if_4 *string
		if (values != nil) && (values.Len() > 0) {
			hx_if_4 = hxrt.StdString(func(hx_value_2 any) *string {
				if hx_value_2 == nil {
					var hx_zero_3 *string
					return hx_zero_3
				}
				return hx_value_2.(*string)
			}(values.Get(0)))
		} else {
			hx_if_4 = hxrt.StringFromLiteral("none")
		}
		return hx_if_4
	}()))
	hxrt.Println(v_2)
	putSink := New_haxe__io__BytesOutput()
	http.__hx_this.customRequest(false, putSink.haxe__io__Output, nil, hxrt.StringFromLiteral("PUT"))
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("method="), putSink.__hx_this.getBytes().__hx_this.toString()))
	hxrt.Println(v_3)
	upload := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,ignored"))
	upload.__hx_this.setParameter(hxrt.StringFromLiteral("token"), hxrt.StringFromLiteral("42"))
	upload.__hx_this.fileTransfer(hxrt.StringFromLiteral("asset"), hxrt.StringFromLiteral("demo.txt"), func(hx_value_5 any) *haxe__io__Input {
		if hx_value_5 == nil {
			var hx_zero_6 *haxe__io__Input
			return hx_zero_6
		}
		return hx_value_5.(*haxe__io__Input)
	}(nil), 4, hxrt.StringFromLiteral("text/plain"))
	uploadSink := New_haxe__io__BytesOutput()
	upload.__hx_this.customRequest(true, uploadSink.haxe__io__Output, nil, nil)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("upload="), uploadSink.__hx_this.getBytes().__hx_this.toString()))
	hxrt.Println(v_4)
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := obj.(type) {
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
	case *haxe__http__HttpBase:
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
	case *haxe__io__BytesOutput:
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
	case *sys__Http:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__net__Host:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__net__Socket:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__net__SocketInput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__net__SocketOutput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe__ds__StringMap:
		return hxrt__generated_method_field__haxe__ds__StringMap(value, key)
	case *haxe__exceptions__NotImplementedException:
		return hxrt__generated_method_field__haxe__exceptions__NotImplementedException(value, key)
	case *haxe__exceptions__PosException:
		return hxrt__generated_method_field__haxe__exceptions__PosException(value, key)
	case *haxe__http__HttpBase:
		return hxrt__generated_method_field__haxe__http__HttpBase(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_method_field__haxe__io__Bytes(value, key)
	case *haxe__io__BytesBuffer:
		return hxrt__generated_method_field__haxe__io__BytesBuffer(value, key)
	case *haxe__io__BytesOutput:
		return hxrt__generated_method_field__haxe__io__BytesOutput(value, key)
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
	case *sys__Http:
		return hxrt__generated_method_field__sys__Http(value, key)
	case *sys__net__Host:
		return hxrt__generated_method_field__sys__net__Host(value, key)
	case *sys__net__Socket:
		return hxrt__generated_method_field__sys__net__Socket(value, key)
	case *sys__net__SocketInput:
		return hxrt__generated_method_field__sys__net__SocketInput(value, key)
	case *sys__net__SocketOutput:
		return hxrt__generated_method_field__sys__net__SocketOutput(value, key)
	default:
		return nil
	}
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

func hxrt__generated_method_field__haxe__http__HttpBase(value *haxe__http__HttpBase, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "addHeader":
		return value.addHeader
	case "addParameter":
		return value.addParameter
	case "get_responseData":
		return value.get_responseData
	case "hasOnData":
		return value.hasOnData
	case "request":
		return value.request
	case "setHeader":
		return value.setHeader
	case "setParameter":
		return value.setParameter
	case "setPostBytes":
		return value.setPostBytes
	case "setPostData":
		return value.setPostData
	case "success":
		return value.success
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

func hxrt__generated_method_field__haxe__io__BytesOutput(value *haxe__io__BytesOutput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "getBytes":
		return value.getBytes
	case "get_length":
		return value.get_length
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

func hxrt__generated_method_field__sys__Http(value *sys__Http, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "customRequest":
		return value.customRequest
	case "encodedParameters":
		return value.encodedParameters
	case "fileTransfer":
		return value.fileTransfer
	case "fileTransfert":
		return value.fileTransfert
	case "getResponseHeaderValues":
		return value.getResponseHeaderValues
	case "handleDataRequest":
		return value.handleDataRequest
	case "hasHeader":
		return value.hasHeader
	case "recordResponseHeaders":
		return value.recordResponseHeaders
	case "request":
		return value.request
	case "requestWith":
		return value.requestWith
	case "resetResponseHeaders":
		return value.resetResponseHeaders
	}
	if value.haxe__http__HttpBase == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__http__HttpBase(value.haxe__http__HttpBase, key)
}

func hxrt__generated_method_field__sys__net__Host(value *sys__net__Host, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "reverse":
		return value.reverse
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__sys__net__Socket(value *sys__net__Socket, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "accept":
		return value.accept
	case "bind":
		return value.bind
	case "close":
		return value.close
	case "connect":
		return value.connect
	case "host":
		return value.host
	case "listen":
		return value.listen
	case "peer":
		return value.peer
	case "read":
		return value.read
	case "replaceHandle":
		return value.replaceHandle
	case "setBlocking":
		return value.setBlocking
	case "setFastSend":
		return value.setFastSend
	case "setTimeout":
		return value.setTimeout
	case "shutdown":
		return value.shutdown
	case "waitForRead":
		return value.waitForRead
	case "write":
		return value.write
	}
	return nil
}

func hxrt__generated_method_field__sys__net__SocketInput(value *sys__net__SocketInput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return value.close
	case "readByte":
		return value.readByte
	case "readBytes":
		return value.readBytes
	}
	if value.haxe__io__Input == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__io__Input(value.haxe__io__Input, key)
}

func hxrt__generated_method_field__sys__net__SocketOutput(value *sys__net__SocketOutput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return value.close
	case "flush":
		return value.flush
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
	case *haxe__http__HttpBase:
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
	case *haxe__io__BytesOutput:
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
	case *sys__Http:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__net__Host:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__net__Socket:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__net__SocketInput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__net__SocketOutput:
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
	case *haxe__ds__StringMap:
		return hxrt__generated_field_lookup__haxe__ds__StringMap(value, key)
	case *haxe__exceptions__NotImplementedException:
		return hxrt__generated_field_lookup__haxe__exceptions__NotImplementedException(value, key)
	case *haxe__exceptions__PosException:
		return hxrt__generated_field_lookup__haxe__exceptions__PosException(value, key)
	case *haxe__http__HttpBase:
		return hxrt__generated_field_lookup__haxe__http__HttpBase(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_lookup__haxe__io__Bytes(value, key)
	case *haxe__io__BytesBuffer:
		return hxrt__generated_field_lookup__haxe__io__BytesBuffer(value, key)
	case *haxe__io__BytesOutput:
		return hxrt__generated_field_lookup__haxe__io__BytesOutput(value, key)
	case *haxe__io__Input:
		return hxrt__generated_field_lookup__haxe__io__Input(value, key)
	case *haxe__io__Output:
		return hxrt__generated_field_lookup__haxe__io__Output(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringKeyValueIterator(value, key)
	case *sys__Http:
		return hxrt__generated_field_lookup__sys__Http(value, key)
	case *sys__net__Host:
		return hxrt__generated_field_lookup__sys__net__Host(value, key)
	case *sys__net__Socket:
		return hxrt__generated_field_lookup__sys__net__Socket(value, key)
	case *sys__net__SocketInput:
		return hxrt__generated_field_lookup__sys__net__SocketInput(value, key)
	case *sys__net__SocketOutput:
		return hxrt__generated_field_lookup__sys__net__SocketOutput(value, key)
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
	case *haxe__http__HttpBase:
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
	case *haxe__io__BytesOutput:
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
	case *sys__Http:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__net__Host:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__net__Socket:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__net__SocketInput:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__net__SocketOutput:
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
	case *haxe__ds__StringMap:
		return hxrt__generated_field_has__haxe__ds__StringMap(value, key)
	case *haxe__exceptions__NotImplementedException:
		return hxrt__generated_field_has__haxe__exceptions__NotImplementedException(value, key)
	case *haxe__exceptions__PosException:
		return hxrt__generated_field_has__haxe__exceptions__PosException(value, key)
	case *haxe__http__HttpBase:
		return hxrt__generated_field_has__haxe__http__HttpBase(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_has__haxe__io__Bytes(value, key)
	case *haxe__io__BytesBuffer:
		return hxrt__generated_field_has__haxe__io__BytesBuffer(value, key)
	case *haxe__io__BytesOutput:
		return hxrt__generated_field_has__haxe__io__BytesOutput(value, key)
	case *haxe__io__Input:
		return hxrt__generated_field_has__haxe__io__Input(value, key)
	case *haxe__io__Output:
		return hxrt__generated_field_has__haxe__io__Output(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_has__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__StringKeyValueIterator(value, key)
	case *sys__Http:
		return hxrt__generated_field_has__sys__Http(value, key)
	case *sys__net__Host:
		return hxrt__generated_field_has__sys__net__Host(value, key)
	case *sys__net__Socket:
		return hxrt__generated_field_has__sys__net__Socket(value, key)
	case *sys__net__SocketInput:
		return hxrt__generated_field_has__sys__net__SocketInput(value, key)
	case *sys__net__SocketOutput:
		return hxrt__generated_field_has__sys__net__SocketOutput(value, key)
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
	case *haxe__http__HttpBase:
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
	case *haxe__io__BytesOutput:
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
	case *sys__Http:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__net__Host:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__net__Socket:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__net__SocketInput:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__net__SocketOutput:
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
	case *haxe__ds__StringMap:
		return hxrt__generated_field_set__haxe__ds__StringMap(value, key, incoming)
	case *haxe__exceptions__NotImplementedException:
		return hxrt__generated_field_set__haxe__exceptions__NotImplementedException(value, key, incoming)
	case *haxe__exceptions__PosException:
		return hxrt__generated_field_set__haxe__exceptions__PosException(value, key, incoming)
	case *haxe__http__HttpBase:
		return hxrt__generated_field_set__haxe__http__HttpBase(value, key, incoming)
	case *haxe__io__Bytes:
		return hxrt__generated_field_set__haxe__io__Bytes(value, key, incoming)
	case *haxe__io__BytesBuffer:
		return hxrt__generated_field_set__haxe__io__BytesBuffer(value, key, incoming)
	case *haxe__io__BytesOutput:
		return hxrt__generated_field_set__haxe__io__BytesOutput(value, key, incoming)
	case *haxe__io__Input:
		return hxrt__generated_field_set__haxe__io__Input(value, key, incoming)
	case *haxe__io__Output:
		return hxrt__generated_field_set__haxe__io__Output(value, key, incoming)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value, key, incoming)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_set__haxe__iterators__StringIterator(value, key, incoming)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__StringKeyValueIterator(value, key, incoming)
	case *sys__Http:
		return hxrt__generated_field_set__sys__Http(value, key, incoming)
	case *sys__net__Host:
		return hxrt__generated_field_set__sys__net__Host(value, key, incoming)
	case *sys__net__Socket:
		return hxrt__generated_field_set__sys__net__Socket(value, key, incoming)
	case *sys__net__SocketInput:
		return hxrt__generated_field_set__sys__net__SocketInput(value, key, incoming)
	case *sys__net__SocketOutput:
		return hxrt__generated_field_set__sys__net__SocketOutput(value, key, incoming)
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
	case *haxe__http__HttpBase:
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
	case *haxe__io__BytesOutput:
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
	case *sys__Http:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__net__Host:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__net__Socket:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__net__SocketInput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__net__SocketOutput:
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
	case *haxe__ds__StringMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("h"))
	case *haxe__exceptions__NotImplementedException:
		return hxrt.NewArray(hxrt.StringFromLiteral("posInfos"))
	case *haxe__exceptions__PosException:
		return hxrt.NewArray(hxrt.StringFromLiteral("posInfos"))
	case *haxe__http__HttpBase:
		return hxrt.NewArray(hxrt.StringFromLiteral("url"), hxrt.StringFromLiteral("responseBytes"), hxrt.StringFromLiteral("responseAsString"), hxrt.StringFromLiteral("postData"), hxrt.StringFromLiteral("postBytes"), hxrt.StringFromLiteral("headers"), hxrt.StringFromLiteral("params"), hxrt.StringFromLiteral("emptyOnData"), hxrt.StringFromLiteral("onData"), hxrt.StringFromLiteral("onBytes"), hxrt.StringFromLiteral("onError"), hxrt.StringFromLiteral("onStatus"))
	case *haxe__io__Bytes:
		return hxrt.NewArray(hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("__hx_dataExposed"))
	case *haxe__io__BytesBuffer:
		return hxrt.NewArray(hxrt.StringFromLiteral("b"))
	case *haxe__io__BytesOutput:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("b"))
	case *haxe__io__Eof:
		return hxrt.NewArray()
	case *haxe__io__Input:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"))
	case *haxe__io__Output:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"))
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("keys"))
	case *haxe__iterators__StringIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *sys__Http:
		return hxrt.NewArray(hxrt.StringFromLiteral("url"), hxrt.StringFromLiteral("responseBytes"), hxrt.StringFromLiteral("responseAsString"), hxrt.StringFromLiteral("postData"), hxrt.StringFromLiteral("postBytes"), hxrt.StringFromLiteral("headers"), hxrt.StringFromLiteral("params"), hxrt.StringFromLiteral("emptyOnData"), hxrt.StringFromLiteral("onData"), hxrt.StringFromLiteral("onBytes"), hxrt.StringFromLiteral("onError"), hxrt.StringFromLiteral("onStatus"), hxrt.StringFromLiteral("noShutdown"), hxrt.StringFromLiteral("cnxTimeout"), hxrt.StringFromLiteral("responseHeaders"), hxrt.StringFromLiteral("responseHeadersSameKey"), hxrt.StringFromLiteral("file"))
	case *sys__net__Host:
		return hxrt.NewArray(hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("ip"))
	case *sys__net__Socket:
		return hxrt.NewArray(hxrt.StringFromLiteral("input"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("custom"), hxrt.StringFromLiteral("handle"))
	case *sys__net__SocketInput:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("handle"))
	case *sys__net__SocketOutput:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("handle"))
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

func hxrt__generated_field_lookup__haxe__http__HttpBase(value *haxe__http__HttpBase, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "emptyOnData":
		return value.emptyOnData
	case "headers":
		return value.headers
	case "onBytes":
		return value.onBytes
	case "onData":
		return value.onData
	case "onError":
		return value.onError
	case "onStatus":
		return value.onStatus
	case "params":
		return value.params
	case "postBytes":
		return value.postBytes
	case "postData":
		return value.postData
	case "responseAsString":
		return value.responseAsString
	case "responseBytes":
		return value.responseBytes
	case "url":
		return value.url
	}
	return nil
}

func hxrt__generated_field_has__haxe__http__HttpBase(value *haxe__http__HttpBase, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "emptyOnData":
		return true
	case "headers":
		return true
	case "onBytes":
		return true
	case "onData":
		return true
	case "onError":
		return true
	case "onStatus":
		return true
	case "params":
		return true
	case "postBytes":
		return true
	case "postData":
		return true
	case "responseAsString":
		return true
	case "responseBytes":
		return true
	case "url":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__http__HttpBase(value *haxe__http__HttpBase, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "emptyOnData":
		if incoming == nil {
			var zero func(*string)
			value.emptyOnData = zero
			return true
		}
		switch typed := incoming.(type) {
		case func(*string):
			value.emptyOnData = typed
			return true
		default:
			return false
		}
	case "headers":
		if incoming == nil {
			var zero *hxrt.Array
			value.headers = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.headers = typed
			return true
		default:
			return false
		}
	case "onBytes":
		if incoming == nil {
			var zero func(*haxe__io__Bytes)
			value.onBytes = zero
			return true
		}
		switch typed := incoming.(type) {
		case func(*haxe__io__Bytes):
			value.onBytes = typed
			return true
		default:
			return false
		}
	case "onData":
		if incoming == nil {
			var zero func(*string)
			value.onData = zero
			return true
		}
		switch typed := incoming.(type) {
		case func(*string):
			value.onData = typed
			return true
		default:
			return false
		}
	case "onError":
		if incoming == nil {
			var zero func(*string)
			value.onError = zero
			return true
		}
		switch typed := incoming.(type) {
		case func(*string):
			value.onError = typed
			return true
		default:
			return false
		}
	case "onStatus":
		if incoming == nil {
			var zero func(int)
			value.onStatus = zero
			return true
		}
		switch typed := incoming.(type) {
		case func(int):
			value.onStatus = typed
			return true
		default:
			return false
		}
	case "params":
		if incoming == nil {
			var zero *hxrt.Array
			value.params = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.params = typed
			return true
		default:
			return false
		}
	case "postBytes":
		if incoming == nil {
			var zero *haxe__io__Bytes
			value.postBytes = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__io__Bytes:
			value.postBytes = typed
			return true
		default:
			return false
		}
	case "postData":
		if incoming == nil {
			var zero *string
			value.postData = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.postData = typed
			return true
		default:
			return false
		}
	case "responseAsString":
		if incoming == nil {
			var zero *string
			value.responseAsString = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.responseAsString = typed
			return true
		default:
			return false
		}
	case "responseBytes":
		if incoming == nil {
			var zero *haxe__io__Bytes
			value.responseBytes = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__io__Bytes:
			value.responseBytes = typed
			return true
		default:
			return false
		}
	case "url":
		if incoming == nil {
			var zero *string
			value.url = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.url = typed
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

func hxrt__generated_field_lookup__haxe__io__BytesOutput(value *haxe__io__BytesOutput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "b":
		return value.b
	}
	if value.haxe__io__Output == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__io__Output(value.haxe__io__Output, key)
}

func hxrt__generated_field_has__haxe__io__BytesOutput(value *haxe__io__BytesOutput, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "b":
		return true
	}
	if value.haxe__io__Output == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__io__Output(value.haxe__io__Output, key)
}

func hxrt__generated_field_set__haxe__io__BytesOutput(value *haxe__io__BytesOutput, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "b":
		if incoming == nil {
			var zero *haxe__io__BytesBuffer
			value.b = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__io__BytesBuffer:
			value.b = typed
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

func hxrt__generated_field_lookup__sys__Http(value *sys__Http, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "cnxTimeout":
		return value.cnxTimeout
	case "file":
		return value.file
	case "noShutdown":
		return value.noShutdown
	case "responseHeaders":
		return value.responseHeaders
	case "responseHeadersSameKey":
		return value.responseHeadersSameKey
	}
	if value.haxe__http__HttpBase == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__http__HttpBase(value.haxe__http__HttpBase, key)
}

func hxrt__generated_field_has__sys__Http(value *sys__Http, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "cnxTimeout":
		return true
	case "file":
		return true
	case "noShutdown":
		return true
	case "responseHeaders":
		return true
	case "responseHeadersSameKey":
		return true
	}
	if value.haxe__http__HttpBase == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__http__HttpBase(value.haxe__http__HttpBase, key)
}

func hxrt__generated_field_set__sys__Http(value *sys__Http, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "cnxTimeout":
		if incoming == nil {
			var zero float64
			value.cnxTimeout = zero
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.cnxTimeout = typed
			return true
		case int:
			value.cnxTimeout = float64(typed)
			return true
		default:
			return false
		}
	case "file":
		if incoming == nil {
			var zero map[string]any
			value.file = zero
			return true
		}
		switch typed := incoming.(type) {
		case map[string]any:
			value.file = typed
			return true
		default:
			return false
		}
	case "noShutdown":
		if incoming == nil {
			var zero bool
			value.noShutdown = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.noShutdown = typed
			return true
		default:
			return false
		}
	case "responseHeaders":
		if incoming == nil {
			var zero *haxe__ds__StringMap
			value.responseHeaders = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds__StringMap:
			value.responseHeaders = typed
			return true
		default:
			return false
		}
	case "responseHeadersSameKey":
		if incoming == nil {
			var zero *haxe__ds__StringMap
			value.responseHeadersSameKey = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds__StringMap:
			value.responseHeadersSameKey = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__http__HttpBase == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__http__HttpBase(value.haxe__http__HttpBase, key, incoming)
}

func hxrt__generated_field_lookup__sys__net__Host(value *sys__net__Host, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "host":
		return value.host
	case "ip":
		return value.ip
	}
	return nil
}

func hxrt__generated_field_has__sys__net__Host(value *sys__net__Host, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "host":
		return true
	case "ip":
		return true
	}
	return false
}

func hxrt__generated_field_set__sys__net__Host(value *sys__net__Host, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "host":
		if incoming == nil {
			var zero *string
			value.host = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.host = typed
			return true
		default:
			return false
		}
	case "ip":
		if incoming == nil {
			var zero int
			value.ip = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.ip = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__net__Socket(value *sys__net__Socket, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "custom":
		return value.custom
	case "handle":
		return value.handle
	case "input":
		return value.input
	case "output":
		return value.output
	}
	return nil
}

func hxrt__generated_field_has__sys__net__Socket(value *sys__net__Socket, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "custom":
		return true
	case "handle":
		return true
	case "input":
		return true
	case "output":
		return true
	}
	return false
}

func hxrt__generated_field_set__sys__net__Socket(value *sys__net__Socket, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "custom":
		if incoming == nil {
			var zero any
			value.custom = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.custom = typed
			return true
		default:
			return false
		}
	case "handle":
		if incoming == nil {
			var zero *hxrt.SocketHandle
			value.handle = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.SocketHandle:
			value.handle = typed
			return true
		default:
			return false
		}
	case "input":
		if incoming == nil {
			var zero *haxe__io__Input
			value.input = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__io__Input:
			value.input = typed
			return true
		default:
			return false
		}
	case "output":
		if incoming == nil {
			var zero *haxe__io__Output
			value.output = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__io__Output:
			value.output = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__net__SocketInput(value *sys__net__SocketInput, key string) any {
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

func hxrt__generated_field_has__sys__net__SocketInput(value *sys__net__SocketInput, key string) bool {
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

func hxrt__generated_field_set__sys__net__SocketInput(value *sys__net__SocketInput, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "handle":
		if incoming == nil {
			var zero *hxrt.SocketHandle
			value.handle = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.SocketHandle:
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

func hxrt__generated_field_lookup__sys__net__SocketOutput(value *sys__net__SocketOutput, key string) any {
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

func hxrt__generated_field_has__sys__net__SocketOutput(value *sys__net__SocketOutput, key string) bool {
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

func hxrt__generated_field_set__sys__net__SocketOutput(value *sys__net__SocketOutput, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "handle":
		if incoming == nil {
			var zero *hxrt.SocketHandle
			value.handle = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.SocketHandle:
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
	case *haxe__io__Encoding:
		return (enumValue != nil)
	case *haxe__io__Error:
		return (enumValue != nil)
	default:
		return false
	}
}
