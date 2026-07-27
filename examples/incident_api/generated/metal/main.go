package main

import "examples_incident_api_metal/hxrt"

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

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := obj.(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__Incident:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__IncidentApi:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__IncidentStore:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__http__TinyHttpServer:
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
	case *Date:
		return hxrt__generated_method_field__Date(value, key)
	case *app__core__Incident:
		return hxrt__generated_method_field__app__core__Incident(value, key)
	case *app__core__IncidentApi:
		return hxrt__generated_method_field__app__core__IncidentApi(value, key)
	case *app__core__IncidentStore:
		return hxrt__generated_method_field__app__core__IncidentStore(value, key)
	case *app__http__TinyHttpServer:
		return hxrt__generated_method_field__app__http__TinyHttpServer(value, key)
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
	case *haxe__iterators__StringIterator:
		return hxrt__generated_method_field__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value, key)
	case *sys__io__FileInput:
		return hxrt__generated_method_field__sys__io__FileInput(value, key)
	case *sys__io__FileOutput:
		return hxrt__generated_method_field__sys__io__FileOutput(value, key)
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

func hxrt__generated_method_field__app__core__Incident(value *app__core__Incident, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "toJson":
		return value.toJson
	}
	return nil
}

func hxrt__generated_method_field__app__core__IncidentApi(value *app__core__IncidentApi, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "createIncident":
		return value.createIncident
	case "handle":
		return value.handle
	case "updateIncident":
		return value.updateIncident
	}
	return nil
}

func hxrt__generated_method_field__app__core__IncidentStore(value *app__core__IncidentStore, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "acknowledge":
		return value.acknowledge
	case "create":
		return value.create
	case "find":
		return value.find
	case "listJson":
		return value.listJson
	case "load":
		return value.load
	case "metricsJson":
		return value.metricsJson
	case "resolve":
		return value.resolve
	case "save":
		return value.save
	}
	return nil
}

func hxrt__generated_method_field__app__http__TinyHttpServer(value *app__http__TinyHttpServer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return value.close
	case "readBody":
		return value.readBody
	case "readRequest":
		return value.readRequest
	case "serveOnce":
		return value.serveOnce
	case "writeResponse":
		return value.writeResponse
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
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__Incident:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__IncidentApi:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__IncidentConfig:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__IncidentRequestException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__IncidentStore:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__http__HttpRequest:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__http__HttpResponse:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__http__TinyHttpServer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
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
	case *Date:
		return hxrt__generated_field_lookup__Date(value, key)
	case *app__core__Incident:
		return hxrt__generated_field_lookup__app__core__Incident(value, key)
	case *app__core__IncidentApi:
		return hxrt__generated_field_lookup__app__core__IncidentApi(value, key)
	case *app__core__IncidentConfig:
		return hxrt__generated_field_lookup__app__core__IncidentConfig(value, key)
	case *app__core__IncidentRequestException:
		return hxrt__generated_field_lookup__app__core__IncidentRequestException(value, key)
	case *app__core__IncidentStore:
		return hxrt__generated_field_lookup__app__core__IncidentStore(value, key)
	case *app__http__HttpRequest:
		return hxrt__generated_field_lookup__app__http__HttpRequest(value, key)
	case *app__http__HttpResponse:
		return hxrt__generated_field_lookup__app__http__HttpResponse(value, key)
	case *app__http__TinyHttpServer:
		return hxrt__generated_field_lookup__app__http__TinyHttpServer(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
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
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringKeyValueIterator(value, key)
	case *sys__io__FileInput:
		return hxrt__generated_field_lookup__sys__io__FileInput(value, key)
	case *sys__io__FileOutput:
		return hxrt__generated_field_lookup__sys__io__FileOutput(value, key)
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
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__core__Incident:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__core__IncidentApi:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__core__IncidentConfig:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__core__IncidentRequestException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__core__IncidentStore:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__http__HttpRequest:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__http__HttpResponse:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__http__TinyHttpServer:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
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
	case *Date:
		return hxrt__generated_field_has__Date(value, key)
	case *app__core__Incident:
		return hxrt__generated_field_has__app__core__Incident(value, key)
	case *app__core__IncidentApi:
		return hxrt__generated_field_has__app__core__IncidentApi(value, key)
	case *app__core__IncidentConfig:
		return hxrt__generated_field_has__app__core__IncidentConfig(value, key)
	case *app__core__IncidentRequestException:
		return hxrt__generated_field_has__app__core__IncidentRequestException(value, key)
	case *app__core__IncidentStore:
		return hxrt__generated_field_has__app__core__IncidentStore(value, key)
	case *app__http__HttpRequest:
		return hxrt__generated_field_has__app__http__HttpRequest(value, key)
	case *app__http__HttpResponse:
		return hxrt__generated_field_has__app__http__HttpResponse(value, key)
	case *app__http__TinyHttpServer:
		return hxrt__generated_field_has__app__http__TinyHttpServer(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
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
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_has__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__StringKeyValueIterator(value, key)
	case *sys__io__FileInput:
		return hxrt__generated_field_has__sys__io__FileInput(value, key)
	case *sys__io__FileOutput:
		return hxrt__generated_field_has__sys__io__FileOutput(value, key)
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
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__core__Incident:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__core__IncidentApi:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__core__IncidentConfig:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__core__IncidentRequestException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__core__IncidentStore:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__http__HttpRequest:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__http__HttpResponse:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *app__http__TinyHttpServer:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
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
	case *Date:
		return hxrt__generated_field_set__Date(value, key, incoming)
	case *app__core__Incident:
		return hxrt__generated_field_set__app__core__Incident(value, key, incoming)
	case *app__core__IncidentApi:
		return hxrt__generated_field_set__app__core__IncidentApi(value, key, incoming)
	case *app__core__IncidentConfig:
		return hxrt__generated_field_set__app__core__IncidentConfig(value, key, incoming)
	case *app__core__IncidentRequestException:
		return hxrt__generated_field_set__app__core__IncidentRequestException(value, key, incoming)
	case *app__core__IncidentStore:
		return hxrt__generated_field_set__app__core__IncidentStore(value, key, incoming)
	case *app__http__HttpRequest:
		return hxrt__generated_field_set__app__http__HttpRequest(value, key, incoming)
	case *app__http__HttpResponse:
		return hxrt__generated_field_set__app__http__HttpResponse(value, key, incoming)
	case *app__http__TinyHttpServer:
		return hxrt__generated_field_set__app__http__TinyHttpServer(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
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
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_set__haxe__iterators__StringIterator(value, key, incoming)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__StringKeyValueIterator(value, key, incoming)
	case *sys__io__FileInput:
		return hxrt__generated_field_set__sys__io__FileInput(value, key, incoming)
	case *sys__io__FileOutput:
		return hxrt__generated_field_set__sys__io__FileOutput(value, key, incoming)
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
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__Incident:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__IncidentApi:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__IncidentConfig:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__IncidentRequestException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__core__IncidentStore:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__http__HttpRequest:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__http__HttpResponse:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *app__http__TinyHttpServer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
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
	case *Date:
		return hxrt.NewArray(hxrt.StringFromLiteral("ms"))
	case *app__core__Incident:
		return hxrt.NewArray(hxrt.StringFromLiteral("id"), hxrt.StringFromLiteral("title"), hxrt.StringFromLiteral("severity"), hxrt.StringFromLiteral("acknowledged"), hxrt.StringFromLiteral("resolved"), hxrt.StringFromLiteral("createdAt"))
	case *app__core__IncidentApi:
		return hxrt.NewArray(hxrt.StringFromLiteral("config"), hxrt.StringFromLiteral("store"), hxrt.StringFromLiteral("requests"))
	case *app__core__IncidentConfig:
		return hxrt.NewArray(hxrt.StringFromLiteral("serviceName"), hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("port"), hxrt.StringFromLiteral("statePath"))
	case *app__core__IncidentRequestException:
		return hxrt.NewArray(hxrt.StringFromLiteral("code"))
	case *app__core__IncidentStore:
		return hxrt.NewArray(hxrt.StringFromLiteral("statePath"), hxrt.StringFromLiteral("incidents"), hxrt.StringFromLiteral("nextId"))
	case *app__http__HttpRequest:
		return hxrt.NewArray(hxrt.StringFromLiteral("method"), hxrt.StringFromLiteral("path"), hxrt.StringFromLiteral("body"))
	case *app__http__HttpResponse:
		return hxrt.NewArray(hxrt.StringFromLiteral("status"), hxrt.StringFromLiteral("body"))
	case *app__http__TinyHttpServer:
		return hxrt.NewArray(hxrt.StringFromLiteral("api"), hxrt.StringFromLiteral("server"), hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("port"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
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
	case *haxe__iterators__StringIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *sys__io__FileInput:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("handle"))
	case *sys__io__FileOutput:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("handle"))
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

func hxrt__generated_field_lookup__app__core__Incident(value *app__core__Incident, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "acknowledged":
		return value.acknowledged
	case "createdAt":
		return value.createdAt
	case "id":
		return value.id
	case "resolved":
		return value.resolved
	case "severity":
		return value.severity
	case "title":
		return value.title
	}
	return nil
}

func hxrt__generated_field_has__app__core__Incident(value *app__core__Incident, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "acknowledged":
		return true
	case "createdAt":
		return true
	case "id":
		return true
	case "resolved":
		return true
	case "severity":
		return true
	case "title":
		return true
	}
	return false
}

func hxrt__generated_field_set__app__core__Incident(value *app__core__Incident, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "acknowledged":
		if incoming == nil {
			var zero bool
			value.acknowledged = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.acknowledged = typed
			return true
		default:
			return false
		}
	case "createdAt":
		if incoming == nil {
			var zero *string
			value.createdAt = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.createdAt = typed
			return true
		default:
			return false
		}
	case "id":
		if incoming == nil {
			var zero int
			value.id = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.id = typed
			return true
		default:
			return false
		}
	case "resolved":
		if incoming == nil {
			var zero bool
			value.resolved = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.resolved = typed
			return true
		default:
			return false
		}
	case "severity":
		if incoming == nil {
			var zero *string
			value.severity = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.severity = typed
			return true
		default:
			return false
		}
	case "title":
		if incoming == nil {
			var zero *string
			value.title = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.title = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__app__core__IncidentApi(value *app__core__IncidentApi, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "config":
		return value.config
	case "requests":
		return value.requests
	case "store":
		return value.store
	}
	return nil
}

func hxrt__generated_field_has__app__core__IncidentApi(value *app__core__IncidentApi, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "config":
		return true
	case "requests":
		return true
	case "store":
		return true
	}
	return false
}

func hxrt__generated_field_set__app__core__IncidentApi(value *app__core__IncidentApi, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "config":
		if incoming == nil {
			var zero *app__core__IncidentConfig
			value.config = zero
			return true
		}
		switch typed := incoming.(type) {
		case *app__core__IncidentConfig:
			value.config = typed
			return true
		default:
			return false
		}
	case "requests":
		if incoming == nil {
			var zero int
			value.requests = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.requests = typed
			return true
		default:
			return false
		}
	case "store":
		if incoming == nil {
			var zero *app__core__IncidentStore
			value.store = zero
			return true
		}
		switch typed := incoming.(type) {
		case *app__core__IncidentStore:
			value.store = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__app__core__IncidentConfig(value *app__core__IncidentConfig, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "host":
		return value.host
	case "port":
		return value.port
	case "serviceName":
		return value.serviceName
	case "statePath":
		return value.statePath
	}
	return nil
}

func hxrt__generated_field_has__app__core__IncidentConfig(value *app__core__IncidentConfig, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "host":
		return true
	case "port":
		return true
	case "serviceName":
		return true
	case "statePath":
		return true
	}
	return false
}

func hxrt__generated_field_set__app__core__IncidentConfig(value *app__core__IncidentConfig, key string, incoming any) bool {
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
	case "port":
		if incoming == nil {
			var zero int
			value.port = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.port = typed
			return true
		default:
			return false
		}
	case "serviceName":
		if incoming == nil {
			var zero *string
			value.serviceName = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.serviceName = typed
			return true
		default:
			return false
		}
	case "statePath":
		if incoming == nil {
			var zero *string
			value.statePath = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.statePath = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__app__core__IncidentRequestException(value *app__core__IncidentRequestException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "code":
		return value.code
	}
	return nil
}

func hxrt__generated_field_has__app__core__IncidentRequestException(value *app__core__IncidentRequestException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "code":
		return true
	}
	return false
}

func hxrt__generated_field_set__app__core__IncidentRequestException(value *app__core__IncidentRequestException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "code":
		if incoming == nil {
			var zero *string
			value.code = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.code = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__app__core__IncidentStore(value *app__core__IncidentStore, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "incidents":
		return value.incidents
	case "nextId":
		return value.nextId
	case "statePath":
		return value.statePath
	}
	return nil
}

func hxrt__generated_field_has__app__core__IncidentStore(value *app__core__IncidentStore, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "incidents":
		return true
	case "nextId":
		return true
	case "statePath":
		return true
	}
	return false
}

func hxrt__generated_field_set__app__core__IncidentStore(value *app__core__IncidentStore, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "incidents":
		if incoming == nil {
			var zero *hxrt.Array
			value.incidents = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.incidents = typed
			return true
		default:
			return false
		}
	case "nextId":
		if incoming == nil {
			var zero int
			value.nextId = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.nextId = typed
			return true
		default:
			return false
		}
	case "statePath":
		if incoming == nil {
			var zero *string
			value.statePath = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.statePath = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__app__http__HttpRequest(value *app__http__HttpRequest, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "body":
		return value.body
	case "method":
		return value.method
	case "path":
		return value.path
	}
	return nil
}

func hxrt__generated_field_has__app__http__HttpRequest(value *app__http__HttpRequest, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "body":
		return true
	case "method":
		return true
	case "path":
		return true
	}
	return false
}

func hxrt__generated_field_set__app__http__HttpRequest(value *app__http__HttpRequest, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "body":
		if incoming == nil {
			var zero *string
			value.body = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.body = typed
			return true
		default:
			return false
		}
	case "method":
		if incoming == nil {
			var zero *string
			value.method = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.method = typed
			return true
		default:
			return false
		}
	case "path":
		if incoming == nil {
			var zero *string
			value.path = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.path = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__app__http__HttpResponse(value *app__http__HttpResponse, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "body":
		return value.body
	case "status":
		return value.status
	}
	return nil
}

func hxrt__generated_field_has__app__http__HttpResponse(value *app__http__HttpResponse, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "body":
		return true
	case "status":
		return true
	}
	return false
}

func hxrt__generated_field_set__app__http__HttpResponse(value *app__http__HttpResponse, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "body":
		if incoming == nil {
			var zero *string
			value.body = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.body = typed
			return true
		default:
			return false
		}
	case "status":
		if incoming == nil {
			var zero int
			value.status = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.status = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__app__http__TinyHttpServer(value *app__http__TinyHttpServer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "api":
		return value.api
	case "host":
		return value.host
	case "port":
		return value.port
	case "server":
		return value.server
	}
	return nil
}

func hxrt__generated_field_has__app__http__TinyHttpServer(value *app__http__TinyHttpServer, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "api":
		return true
	case "host":
		return true
	case "port":
		return true
	case "server":
		return true
	}
	return false
}

func hxrt__generated_field_set__app__http__TinyHttpServer(value *app__http__TinyHttpServer, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "api":
		if incoming == nil {
			var zero *app__core__IncidentApi
			value.api = zero
			return true
		}
		switch typed := incoming.(type) {
		case *app__core__IncidentApi:
			value.api = typed
			return true
		default:
			return false
		}
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
	case "port":
		if incoming == nil {
			var zero int
			value.port = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.port = typed
			return true
		default:
			return false
		}
	case "server":
		if incoming == nil {
			var zero *sys__net__Socket
			value.server = zero
			return true
		}
		switch typed := incoming.(type) {
		case *sys__net__Socket:
			value.server = typed
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
	case *sys__io__FileSeek:
		return (enumValue != nil)
	default:
		return false
	}
}
