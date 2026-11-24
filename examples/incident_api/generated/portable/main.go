package main

import "examples_incident_api_portable/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func argValue(name *string, fallback *string) *string {
	args := hxrt.ArrayFromValues(func(hx_sort_src_1 []*string) []any {
		hx_sort_out_3 := make([]any, 0, len(hx_sort_src_1))
		for _, hx_sort_item_2 := range hx_sort_src_1 {
			hx_sort_out_3 = append(hx_sort_out_3, hx_sort_item_2)
		}
		return hx_sort_out_3
	}(hxrt.SysArgs()))
	i := 0
	for i < int((hxrt.Int32Wrap(args.Len()) - hxrt.Int32Wrap(1))) {
		if hxrt.StringEqualAny(args.Get(i), name) {
			return func(hx_value_4 any) *string {
				if hx_value_4 == nil {
					var hx_zero_5 *string
					return hx_zero_5
				}
				return hx_value_4.(*string)
			}(args.Get(int((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1)))))
		}
		i = int(int32((i + 1)))
	}
	return fallback
}

func hasArg(name *string) bool {
	_g := 0
	_g1 := hxrt.ArrayFromValues(func(hx_sort_src_6 []*string) []any {
		hx_sort_out_8 := make([]any, 0, len(hx_sort_src_6))
		for _, hx_sort_item_7 := range hx_sort_src_6 {
			hx_sort_out_8 = append(hx_sort_out_8, hx_sort_item_7)
		}
		return hx_sort_out_8
	}(hxrt.SysArgs()))
	for _g < _g1.Len() {
		arg := func(hx_value_9 any) *string {
			if hx_value_9 == nil {
				var hx_zero_10 *string
				return hx_zero_10
			}
			return hx_value_9.(*string)
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
	switch value := hxrt.ObjectSource(obj).(type) {
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
	case *app__http__TinyHttpServer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
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
	case *app__core__Incident:
		return hxrt__generated_method_field__app__core__Incident(value, key)
	case *app__core__IncidentApi:
		return hxrt__generated_method_field__app__core__IncidentApi(value, key)
	case *app__core__IncidentRequestException:
		return hxrt__generated_method_field__app__core__IncidentRequestException(value, key)
	case *app__core__IncidentStore:
		return hxrt__generated_method_field__app__core__IncidentStore(value, key)
	case *app__http__TinyHttpServer:
		return hxrt__generated_method_field__app__http__TinyHttpServer(value, key)
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
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

func hxrt__generated_method_field__app__core__Incident(value *app__core__Incident, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "toJson":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toJson()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toJson")
		}()
	}
	return nil
}

func hxrt__generated_method_field__app__core__IncidentApi(value *app__core__IncidentApi, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "createIncident":
		return func() *hxrt.Function[func(*string) *app__http__HttpResponse] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(body *string) *app__http__HttpResponse {
				return hx_receiver_1.__hx_this.createIncident(body)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "createIncident")
		}()
	case "handle":
		return func() *hxrt.Function[func(*app__http__HttpRequest) *app__http__HttpResponse] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(request *app__http__HttpRequest) *app__http__HttpResponse {
				return hx_receiver_1.__hx_this.handle(request)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "handle")
		}()
	case "updateIncident":
		return func() *hxrt.Function[func(*string, *string) *app__http__HttpResponse] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(path *string, action *string) *app__http__HttpResponse {
				return hx_receiver_1.__hx_this.updateIncident(path, action)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "updateIncident")
		}()
	}
	return nil
}

func hxrt__generated_method_field__app__core__IncidentRequestException(value *app__core__IncidentRequestException, key string) any {
	if value == nil {
		return nil
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__app__core__IncidentStore(value *app__core__IncidentStore, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "acknowledge":
		return func() *hxrt.Function[func(int) *app__core__Incident] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(id int) *app__core__Incident {
				return hx_receiver_1.__hx_this.acknowledge(id)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "acknowledge")
		}()
	case "create":
		return func() *hxrt.Function[func(*string, *string) *app__core__Incident] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(title *string, severity *string) *app__core__Incident {
				return hx_receiver_1.__hx_this.create(title, severity)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "create")
		}()
	case "find":
		return func() *hxrt.Function[func(int) *app__core__Incident] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(id int) *app__core__Incident {
				return hx_receiver_1.__hx_this.find(id)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "find")
		}()
	case "listJson":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.listJson()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "listJson")
		}()
	case "load":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.load()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "load")
		}()
	case "metricsJson":
		return func() *hxrt.Function[func(*string, int) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(serviceName *string, requests int) *string {
				return hx_receiver_1.__hx_this.metricsJson(serviceName, requests)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "metricsJson")
		}()
	case "resolve":
		return func() *hxrt.Function[func(int) *app__core__Incident] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(id int) *app__core__Incident {
				return hx_receiver_1.__hx_this.resolve(id)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "resolve")
		}()
	case "save":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.save()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "save")
		}()
	}
	return nil
}

func hxrt__generated_method_field__app__http__TinyHttpServer(value *app__http__TinyHttpServer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.close()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "close")
		}()
	case "readBody":
		return func() *hxrt.Function[func(*sys__net__Socket, int) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(peer *sys__net__Socket, length int) *string {
				return hx_receiver_1.__hx_this.readBody(peer, length)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readBody")
		}()
	case "readRequest":
		return func() *hxrt.Function[func(*sys__net__Socket) *app__http__HttpRequest] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(peer *sys__net__Socket) *app__http__HttpRequest {
				return hx_receiver_1.__hx_this.readRequest(peer)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readRequest")
		}()
	case "serveOnce":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.serveOnce()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "serveOnce")
		}()
	case "writeResponse":
		return func() *hxrt.Function[func(*sys__net__Socket, *app__http__HttpResponse)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(peer *sys__net__Socket, response *app__http__HttpResponse) {
				hx_receiver_1.__hx_this.writeResponse(peer, response)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeResponse")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.details()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.get_message()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_native()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_1.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
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
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
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
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_nativeView":
		return func() *hxrt.Function[func() *hxrt.ByteView] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *hxrt.ByteView {
				return hx_receiver_1.__hx_this.__hx_nativeView()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "__hx_nativeView")
		}()
	case "blit":
		return func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
				hx_receiver_1.__hx_this.blit(pos, src, srcpos, len)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "blit")
		}()
	case "compare":
		return func() *hxrt.Function[func(*haxe__io__Bytes) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
				return hx_receiver_1.__hx_this.compare(other)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "compare")
		}()
	case "fill":
		return func() *hxrt.Function[func(int, int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len int, value int) {
				hx_receiver_1.__hx_this.fill(pos, len, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "fill")
		}()
	case "get":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_1.__hx_this.get(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get")
		}()
	case "getData":
		return func() *hxrt.Function[func() []int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() []int {
				return hx_receiver_1.__hx_this.getData()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getData")
		}()
	case "getDouble":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_1.__hx_this.getDouble(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getDouble")
		}()
	case "getFloat":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_1.__hx_this.getFloat(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getFloat")
		}()
	case "getInt32":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_1.__hx_this.getInt32(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getInt32")
		}()
	case "getInt64":
		return func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
				return hx_receiver_1.__hx_this.getInt64(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getInt64")
		}()
	case "getString":
		return func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
				return hx_receiver_1.__hx_this.getString(pos, len, encoding)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getString")
		}()
	case "getUInt16":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_1.__hx_this.getUInt16(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getUInt16")
		}()
	case "readString":
		return func() *hxrt.Function[func(int, int) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len int) *string {
				return hx_receiver_1.__hx_this.readString(pos, len)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readString")
		}()
	case "set":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_1.__hx_this.set(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "set")
		}()
	case "setDouble":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_1.__hx_this.setDouble(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setDouble")
		}()
	case "setFloat":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_1.__hx_this.setFloat(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setFloat")
		}()
	case "setInt32":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_1.__hx_this.setInt32(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setInt32")
		}()
	case "setInt64":
		return func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
				hx_receiver_1.__hx_this.setInt64(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setInt64")
		}()
	case "setUInt16":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_1.__hx_this.setUInt16(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setUInt16")
		}()
	case "sub":
		return func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
				return hx_receiver_1.__hx_this.sub(pos, len)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "sub")
		}()
	case "toHex":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toHex()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toHex")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__BytesBuffer(value *haxe__io__BytesBuffer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "add":
		return func() *hxrt.Function[func(*haxe__io__Bytes)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(source *haxe__io__Bytes) {
				hx_receiver_1.__hx_this.add(source)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "add")
		}()
	case "addByte":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_1.__hx_this.addByte(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "addByte")
		}()
	case "addBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(source *haxe__io__Bytes, pos int, len int) {
				hx_receiver_1.__hx_this.addBytes(source, pos, len)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "addBytes")
		}()
	case "addDouble":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value float64) {
				hx_receiver_1.__hx_this.addDouble(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "addDouble")
		}()
	case "addFloat":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value float64) {
				hx_receiver_1.__hx_this.addFloat(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "addFloat")
		}()
	case "addInt32":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_1.__hx_this.addInt32(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "addInt32")
		}()
	case "addInt64":
		return func() *hxrt.Function[func(*haxe___Int64_____Int64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value *haxe___Int64_____Int64) {
				hx_receiver_1.__hx_this.addInt64(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "addInt64")
		}()
	case "addString":
		return func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
				hx_receiver_1.__hx_this.addString(value, encoding)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "addString")
		}()
	case "getBytes":
		return func() *hxrt.Function[func() *haxe__io__Bytes] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *haxe__io__Bytes {
				return hx_receiver_1.__hx_this.getBytes()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getBytes")
		}()
	case "get_length":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.get_length()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_length")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Eof(value *haxe__io__Eof, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Input(value *haxe__io__Input, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.close()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "close")
		}()
	case "read":
		return func() *hxrt.Function[func(int) *haxe__io__Bytes] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
				return hx_receiver_1.__hx_this.read(nbytes)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "read")
		}()
	case "readAll":
		return func() *hxrt.Function[func(any) *haxe__io__Bytes] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
				return hx_receiver_1.__hx_this.readAll(bufsize)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readAll")
		}()
	case "readByte":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.readByte()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readByte")
		}()
	case "readBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
				return hx_receiver_1.__hx_this.readBytes(bytes, pos, len)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readBytes")
		}()
	case "readDouble":
		return func() *hxrt.Function[func() float64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() float64 {
				return hx_receiver_1.__hx_this.readDouble()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readDouble")
		}()
	case "readFloat":
		return func() *hxrt.Function[func() float64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() float64 {
				return hx_receiver_1.__hx_this.readFloat()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readFloat")
		}()
	case "readFullBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
				hx_receiver_1.__hx_this.readFullBytes(bytes, pos, len)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readFullBytes")
		}()
	case "readInt16":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.readInt16()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readInt16")
		}()
	case "readInt24":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.readInt24()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readInt24")
		}()
	case "readInt32":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.readInt32()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readInt32")
		}()
	case "readInt8":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.readInt8()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readInt8")
		}()
	case "readLine":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.readLine()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readLine")
		}()
	case "readString":
		return func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
				return hx_receiver_1.__hx_this.readString(len, encoding)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readString")
		}()
	case "readUInt16":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.readUInt16()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readUInt16")
		}()
	case "readUInt24":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.readUInt24()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readUInt24")
		}()
	case "readUntil":
		return func() *hxrt.Function[func(int) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(end int) *string {
				return hx_receiver_1.__hx_this.readUntil(end)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readUntil")
		}()
	case "set_bigEndian":
		return func() *hxrt.Function[func(bool) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value bool) bool {
				return hx_receiver_1.__hx_this.set_bigEndian(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "set_bigEndian")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Output(value *haxe__io__Output, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.close()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "close")
		}()
	case "flush":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.flush()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "flush")
		}()
	case "prepare":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(nbytes int) {
				hx_receiver_1.__hx_this.prepare(nbytes)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "prepare")
		}()
	case "set_bigEndian":
		return func() *hxrt.Function[func(bool) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value bool) bool {
				return hx_receiver_1.__hx_this.set_bigEndian(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "set_bigEndian")
		}()
	case "write":
		return func() *hxrt.Function[func(*haxe__io__Bytes)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
				hx_receiver_1.__hx_this.write(bytes)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "write")
		}()
	case "writeByte":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_1.__hx_this.writeByte(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeByte")
		}()
	case "writeBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
				return hx_receiver_1.__hx_this.writeBytes(bytes, pos, len)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeBytes")
		}()
	case "writeDouble":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value float64) {
				hx_receiver_1.__hx_this.writeDouble(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeDouble")
		}()
	case "writeFloat":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value float64) {
				hx_receiver_1.__hx_this.writeFloat(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeFloat")
		}()
	case "writeFullBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
				hx_receiver_1.__hx_this.writeFullBytes(bytes, pos, len)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeFullBytes")
		}()
	case "writeInput":
		return func() *hxrt.Function[func(*haxe__io__Input, any)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
				hx_receiver_1.__hx_this.writeInput(input, bufsize)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeInput")
		}()
	case "writeInt16":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_1.__hx_this.writeInt16(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeInt16")
		}()
	case "writeInt24":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_1.__hx_this.writeInt24(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeInt24")
		}()
	case "writeInt32":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_1.__hx_this.writeInt32(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeInt32")
		}()
	case "writeInt8":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_1.__hx_this.writeInt8(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeInt8")
		}()
	case "writeString":
		return func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
				hx_receiver_1.__hx_this.writeString(value, encoding)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeString")
		}()
	case "writeUInt16":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_1.__hx_this.writeUInt16(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeUInt16")
		}()
	case "writeUInt24":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_1.__hx_this.writeUInt24(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeUInt24")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__io__FileInput(value *sys__io__FileInput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.close()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "close")
		}()
	case "eof":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.eof()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "eof")
		}()
	case "readByte":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.readByte()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readByte")
		}()
	case "readBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
				return hx_receiver_1.__hx_this.readBytes(bytes, pos, length)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readBytes")
		}()
	case "seek":
		return func() *hxrt.Function[func(int, *sys__io__FileSeek)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(p int, pos *sys__io__FileSeek) {
				hx_receiver_1.__hx_this.seek(p, pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "seek")
		}()
	case "tell":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.tell()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "tell")
		}()
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
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.close()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "close")
		}()
	case "flush":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.flush()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "flush")
		}()
	case "seek":
		return func() *hxrt.Function[func(int, *sys__io__FileSeek)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(p int, pos *sys__io__FileSeek) {
				hx_receiver_1.__hx_this.seek(p, pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "seek")
		}()
	case "tell":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.tell()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "tell")
		}()
	case "writeByte":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_1.__hx_this.writeByte(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeByte")
		}()
	case "writeBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
				return hx_receiver_1.__hx_this.writeBytes(bytes, pos, length)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeBytes")
		}()
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
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.reverse()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "reverse")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__net__Socket(value *sys__net__Socket, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "accept":
		return func() *hxrt.Function[func() *sys__net__Socket] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *sys__net__Socket {
				return hx_receiver_1.__hx_this.accept()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "accept")
		}()
	case "bind":
		return func() *hxrt.Function[func(*sys__net__Host, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
				hx_receiver_1.__hx_this.bind(host, port)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "bind")
		}()
	case "close":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.close()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "close")
		}()
	case "connect":
		return func() *hxrt.Function[func(*sys__net__Host, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
				hx_receiver_1.__hx_this.connect(host, port)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "connect")
		}()
	case "host":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.host()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "host")
		}()
	case "listen":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(connections int) {
				hx_receiver_1.__hx_this.listen(connections)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "listen")
		}()
	case "peer":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.peer()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "peer")
		}()
	case "read":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.read()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "read")
		}()
	case "replaceHandle":
		return func() *hxrt.Function[func(*hxrt.SocketHandle)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(next *hxrt.SocketHandle) {
				hx_receiver_1.__hx_this.replaceHandle(next)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "replaceHandle")
		}()
	case "setBlocking":
		return func() *hxrt.Function[func(bool)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(blocking bool) {
				hx_receiver_1.__hx_this.setBlocking(blocking)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setBlocking")
		}()
	case "setFastSend":
		return func() *hxrt.Function[func(bool)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(fastSend bool) {
				hx_receiver_1.__hx_this.setFastSend(fastSend)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setFastSend")
		}()
	case "setTimeout":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(timeout float64) {
				hx_receiver_1.__hx_this.setTimeout(timeout)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setTimeout")
		}()
	case "shutdown":
		return func() *hxrt.Function[func(bool, bool)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(read bool, write bool) {
				hx_receiver_1.__hx_this.shutdown(read, write)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "shutdown")
		}()
	case "waitForRead":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.waitForRead()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "waitForRead")
		}()
	case "write":
		return func() *hxrt.Function[func(*string)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(content *string) {
				hx_receiver_1.__hx_this.write(content)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "write")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__net__SocketInput(value *sys__net__SocketInput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.close()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "close")
		}()
	case "readByte":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.readByte()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readByte")
		}()
	case "readBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
				return hx_receiver_1.__hx_this.readBytes(bytes, pos, length)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readBytes")
		}()
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
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.close()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "close")
		}()
	case "flush":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.flush()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "flush")
		}()
	case "writeByte":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_1.__hx_this.writeByte(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeByte")
		}()
	case "writeBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
				return hx_receiver_1.__hx_this.writeBytes(bytes, pos, length)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeBytes")
		}()
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
	switch value := hxrt.ObjectSource(object).(type) {
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
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
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
	switch value := hxrt.ObjectSource(object).(type) {
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
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
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
	switch value := hxrt.ObjectSource(object).(type) {
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
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
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
	switch value := hxrt.ObjectSource(object).(type) {
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
	case *app__core__Incident:
		return hxrt.NewArray(hxrt.StringFromLiteral("id"), hxrt.StringFromLiteral("title"), hxrt.StringFromLiteral("severity"), hxrt.StringFromLiteral("acknowledged"), hxrt.StringFromLiteral("resolved"), hxrt.StringFromLiteral("createdAt"))
	case *app__core__IncidentApi:
		return hxrt.NewArray(hxrt.StringFromLiteral("config"), hxrt.StringFromLiteral("store"), hxrt.StringFromLiteral("requests"))
	case *app__core__IncidentConfig:
		return hxrt.NewArray(hxrt.StringFromLiteral("serviceName"), hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("port"), hxrt.StringFromLiteral("statePath"))
	case *app__core__IncidentRequestException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("code"))
	case *app__core__IncidentStore:
		return hxrt.NewArray(hxrt.StringFromLiteral("statePath"), hxrt.StringFromLiteral("incidents"), hxrt.StringFromLiteral("nextId"))
	case *app__http__HttpRequest:
		return hxrt.NewArray(hxrt.StringFromLiteral("method"), hxrt.StringFromLiteral("path"), hxrt.StringFromLiteral("body"))
	case *app__http__HttpResponse:
		return hxrt.NewArray(hxrt.StringFromLiteral("status"), hxrt.StringFromLiteral("body"))
	case *app__http__TinyHttpServer:
		return hxrt.NewArray(hxrt.StringFromLiteral("api"), hxrt.StringFromLiteral("server"), hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("port"))
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__exceptions__NotImplementedException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("posInfos"))
	case *haxe__exceptions__PosException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("posInfos"))
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

func hxrt__generated_field_lookup__app__core__Incident(value *app__core__Incident, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "acknowledged":
		return hxrt.ReadStoredField(value.acknowledged)
	case "createdAt":
		return hxrt.ReadStoredField(value.createdAt)
	case "id":
		return hxrt.ReadStoredField(value.id)
	case "resolved":
		return hxrt.ReadStoredField(value.resolved)
	case "severity":
		return hxrt.ReadStoredField(value.severity)
	case "title":
		return hxrt.ReadStoredField(value.title)
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
		return hxrt.ReadStoredField(value.config)
	case "requests":
		return hxrt.ReadStoredField(value.requests)
	case "store":
		return hxrt.ReadStoredField(value.store)
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
		return hxrt.ReadStoredField(value.host)
	case "port":
		return hxrt.ReadStoredField(value.port)
	case "serviceName":
		return hxrt.ReadStoredField(value.serviceName)
	case "statePath":
		return hxrt.ReadStoredField(value.statePath)
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
		return hxrt.ReadStoredField(value.code)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__app__core__IncidentRequestException(value *app__core__IncidentRequestException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "code":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
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
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__app__core__IncidentStore(value *app__core__IncidentStore, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "incidents":
		return hxrt.ReadStoredField(value.incidents)
	case "nextId":
		return hxrt.ReadStoredField(value.nextId)
	case "statePath":
		return hxrt.ReadStoredField(value.statePath)
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
		return hxrt.ReadStoredField(value.body)
	case "method":
		return hxrt.ReadStoredField(value.method)
	case "path":
		return hxrt.ReadStoredField(value.path)
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
		return hxrt.ReadStoredField(value.body)
	case "status":
		return hxrt.ReadStoredField(value.status)
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
		return hxrt.ReadStoredField(value.api)
	case "host":
		return hxrt.ReadStoredField(value.host)
	case "port":
		return hxrt.ReadStoredField(value.port)
	case "server":
		return hxrt.ReadStoredField(value.server)
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
		return hxrt.ReadStoredField(value.posInfos)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__exceptions__PosException(value *haxe__exceptions__PosException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "posInfos":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__exceptions__PosException(value *haxe__exceptions__PosException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "posInfos":
		if incoming == nil {
			var zero hxrt.Anonymous
			value.posInfos = zero
			return true
		}
		switch typed := incoming.(type) {
		case hxrt.Anonymous:
			value.posInfos = typed
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

func hxrt__generated_field_lookup__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_dataExposed":
		return hxrt.ReadStoredField(value.__hx_dataExposed)
	case "__hx_raw":
		return hxrt.ReadStoredField(value.__hx_raw)
	case "__hx_rawValid":
		return hxrt.ReadStoredField(value.__hx_rawValid)
	case "b":
		return hxrt.ReadStoredField(value.b)
	case "length":
		return hxrt.ReadStoredField(value.length)
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
		return hxrt.ReadStoredField(value.b)
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
		return hxrt.ReadStoredField(value.bigEndian)
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
		return hxrt.ReadStoredField(value.bigEndian)
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
		return hxrt.ReadStoredField(value.offset)
	case "s":
		return hxrt.ReadStoredField(value.s)
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
		return hxrt.ReadStoredField(value.offset)
	case "s":
		return hxrt.ReadStoredField(value.s)
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
		return hxrt.ReadStoredField(value.handle)
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
		return hxrt.ReadStoredField(value.handle)
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
		return hxrt.ReadStoredField(value.host)
	case "ip":
		return hxrt.ReadStoredField(value.ip)
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
		return hxrt.ReadStoredField(value.custom)
	case "handle":
		return hxrt.ReadStoredField(value.handle)
	case "input":
		return hxrt.ReadStoredField(value.input)
	case "output":
		return hxrt.ReadStoredField(value.output)
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
		return hxrt.ReadStoredField(value.handle)
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
		return hxrt.ReadStoredField(value.handle)
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
	case *haxe__StackItem:
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

func hx_is_of_type_token(hx_value any, hx_type any) bool {
	switch hx_type_marker := hx_type.(type) {
	case *hxrt__TypeClassValue:
		if hx_type_marker == nil {
			return false
		}
		if hx_type_marker.name == nil {
			return false
		}
		switch *hx_type_marker.name {
		case "Array":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *hxrt.Array:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Bool":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case bool:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Class":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *hxrt__TypeClassValue:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Dynamic":
			return (hx_value != nil)
		case "Enum":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *hxrt__TypeEnumValue:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Float":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case int:
					return true
				case int8:
					return true
				case int16:
					return true
				case int32:
					return true
				case int64:
					return true
				case uint:
					return true
				case uint8:
					return true
				case uint16:
					return true
				case uint32:
					return true
				case uint64:
					return true
				case uintptr:
					return true
				case float32:
					return true
				case float64:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Int":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case int:
					return true
				case int8:
					return true
				case int16:
					return true
				case int32:
					return true
				case int64:
					return true
				case uint:
					return true
				case uint8:
					return true
				case uint16:
					return true
				case uint32:
					return true
				case uint64:
					return true
				case uintptr:
					return true
				default:
					return false
				}
			}(hx_value)
		case "String":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *string:
					return true
				case string:
					return true
				default:
					return false
				}
			}(hx_value)
		case "app.core.Incident":
			return hx_is_nominal_app__core__Incident(hx_value)
		case "app.core.IncidentApi":
			return hx_is_nominal_app__core__IncidentApi(hx_value)
		case "app.core.IncidentConfig":
			return hx_is_nominal_app__core__IncidentConfig(hx_value)
		case "app.core.IncidentRequestException":
			return hx_is_nominal_app__core__IncidentRequestException(hx_value)
		case "app.core.IncidentStore":
			return hx_is_nominal_app__core__IncidentStore(hx_value)
		case "app.http.HttpRequest":
			return hx_is_nominal_app__http__HttpRequest(hx_value)
		case "app.http.HttpResponse":
			return hx_is_nominal_app__http__HttpResponse(hx_value)
		case "app.http.TinyHttpServer":
			return hx_is_nominal_app__http__TinyHttpServer(hx_value)
		case "haxe.Exception":
			return hx_is_nominal_haxe__Exception(hx_value)
		case "haxe.ValueException":
			return hx_is_nominal_haxe__ValueException(hx_value)
		case "haxe._Int64.___Int64":
			return hx_is_nominal_haxe___Int64_____Int64(hx_value)
		case "haxe.exceptions.NotImplementedException":
			return hx_is_nominal_haxe__exceptions__NotImplementedException(hx_value)
		case "haxe.exceptions.PosException":
			return hx_is_nominal_haxe__exceptions__PosException(hx_value)
		case "haxe.io.Bytes":
			return hx_is_nominal_haxe__io__Bytes(hx_value)
		case "haxe.io.BytesBuffer":
			return hx_is_nominal_haxe__io__BytesBuffer(hx_value)
		case "haxe.io.Eof":
			return hx_is_nominal_haxe__io__Eof(hx_value)
		case "haxe.io.Input":
			return hx_is_nominal_haxe__io__Input(hx_value)
		case "haxe.io.Output":
			return hx_is_nominal_haxe__io__Output(hx_value)
		case "haxe.iterators.StringIterator":
			return hx_is_nominal_haxe__iterators__StringIterator(hx_value)
		case "haxe.iterators.StringKeyValueIterator":
			return hx_is_nominal_haxe__iterators__StringKeyValueIterator(hx_value)
		case "sys.io.FileInput":
			return hx_is_nominal_sys__io__FileInput(hx_value)
		case "sys.io.FileOutput":
			return hx_is_nominal_sys__io__FileOutput(hx_value)
		case "sys.net.Host":
			return hx_is_nominal_sys__net__Host(hx_value)
		case "sys.net.Socket":
			return hx_is_nominal_sys__net__Socket(hx_value)
		case "sys.net.SocketInput":
			return hx_is_nominal_sys__net__SocketInput(hx_value)
		case "sys.net.SocketOutput":
			return hx_is_nominal_sys__net__SocketOutput(hx_value)
		default:
			return false
		}
	case *hxrt__TypeEnumValue:
		if hx_type_marker == nil {
			return false
		}
		if hx_type_marker.name == nil {
			return false
		}
		switch *hx_type_marker.name {
		case "haxe.StackItem":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__StackItem:
					return true
				default:
					return false
				}
			}(hx_value)
		case "haxe.io.Encoding":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__io__Encoding:
					return true
				default:
					return false
				}
			}(hx_value)
		case "haxe.io.Error":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__io__Error:
					return true
				default:
					return false
				}
			}(hx_value)
		case "sys.io.FileSeek":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *sys__io__FileSeek:
					return true
				default:
					return false
				}
			}(hx_value)
		default:
			return false
		}
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
	case *hxrt.Function[func(*string) *app__http__HttpResponse]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_1 any) *string {
			if hx_value_1 == nil {
				var hx_zero_2 *string
				return hx_zero_2
			}
			return hx_value_1.(*string)
		}(argument_0))
	case *hxrt.Function[func(*app__http__HttpRequest) *app__http__HttpResponse]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_app__http__HttpRequest(argument_0))
	case *hxrt.Function[func(*string, *string) *app__http__HttpResponse]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_3 any) *string {
			if hx_value_3 == nil {
				var hx_zero_4 *string
				return hx_zero_4
			}
			return hx_value_3.(*string)
		}(argument_0), func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(argument_1))
	case *hxrt.Function[func(int) *app__core__Incident]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_7 any) int {
			if hx_value_7 == nil {
				var hx_zero_8 int
				return hx_zero_8
			}
			return hx_value_7.(int)
		}(argument_0))
	case *hxrt.Function[func(*string, *string) *app__core__Incident]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_9 any) *string {
			if hx_value_9 == nil {
				var hx_zero_10 *string
				return hx_zero_10
			}
			return hx_value_9.(*string)
		}(argument_0), func(hx_value_11 any) *string {
			if hx_value_11 == nil {
				var hx_zero_12 *string
				return hx_zero_12
			}
			return hx_value_11.(*string)
		}(argument_1))
	case *hxrt.Function[func()]:
		hxrt.FunctionCall(source)()
		return nil
	case *hxrt.Function[func(*string, int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_13 any) *string {
			if hx_value_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_value_13.(*string)
		}(argument_0), func(hx_value_15 any) int {
			if hx_value_15 == nil {
				var hx_zero_16 int
				return hx_zero_16
			}
			return hx_value_15.(int)
		}(argument_1))
	case *hxrt.Function[func(*sys__net__Socket, int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_sys__net__Socket(argument_0), func(hx_value_17 any) int {
			if hx_value_17 == nil {
				var hx_zero_18 int
				return hx_zero_18
			}
			return hx_value_17.(int)
		}(argument_1))
	case *hxrt.Function[func(*sys__net__Socket) *app__http__HttpRequest]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_sys__net__Socket(argument_0))
	case *hxrt.Function[func(*sys__net__Socket, *app__http__HttpResponse)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(hx_recover_class_sys__net__Socket(argument_0), hx_recover_class_app__http__HttpResponse(argument_1))
		return nil
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int, *haxe__io__Bytes, int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		var argument_3 any = nil
		if 3 < len(arguments) {
			argument_3 = arguments[3]
		}
		hxrt.FunctionCall(source)(func(hx_value_19 any) int {
			if hx_value_19 == nil {
				var hx_zero_20 int
				return hx_zero_20
			}
			return hx_value_19.(int)
		}(argument_0), hx_recover_class_haxe__io__Bytes(argument_1), func(hx_value_21 any) int {
			if hx_value_21 == nil {
				var hx_zero_22 int
				return hx_zero_22
			}
			return hx_value_21.(int)
		}(argument_2), func(hx_value_23 any) int {
			if hx_value_23 == nil {
				var hx_zero_24 int
				return hx_zero_24
			}
			return hx_value_23.(int)
		}(argument_3))
		return nil
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(argument_0))
	case *hxrt.Function[func(int, int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		hxrt.FunctionCall(source)(func(hx_value_25 any) int {
			if hx_value_25 == nil {
				var hx_zero_26 int
				return hx_zero_26
			}
			return hx_value_25.(int)
		}(argument_0), func(hx_value_27 any) int {
			if hx_value_27 == nil {
				var hx_zero_28 int
				return hx_zero_28
			}
			return hx_value_27.(int)
		}(argument_1), func(hx_value_29 any) int {
			if hx_value_29 == nil {
				var hx_zero_30 int
				return hx_zero_30
			}
			return hx_value_29.(int)
		}(argument_2))
		return nil
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_31 any) int {
			if hx_value_31 == nil {
				var hx_zero_32 int
				return hx_zero_32
			}
			return hx_value_31.(int)
		}(argument_0))
	case *hxrt.Function[func() []int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) float64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_33 any) int {
			if hx_value_33 == nil {
				var hx_zero_34 int
				return hx_zero_34
			}
			return hx_value_33.(int)
		}(argument_0))
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_35 any) int {
			if hx_value_35 == nil {
				var hx_zero_36 int
				return hx_zero_36
			}
			return hx_value_35.(int)
		}(argument_0))
	case *hxrt.Function[func(int, int, *haxe__io__Encoding) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(func(hx_value_37 any) int {
			if hx_value_37 == nil {
				var hx_zero_38 int
				return hx_zero_38
			}
			return hx_value_37.(int)
		}(argument_0), func(hx_value_39 any) int {
			if hx_value_39 == nil {
				var hx_zero_40 int
				return hx_zero_40
			}
			return hx_value_39.(int)
		}(argument_1), func(hx_value_41 any) *haxe__io__Encoding {
			if hx_value_41 == nil {
				var hx_zero_42 *haxe__io__Encoding
				return hx_zero_42
			}
			return hx_value_41.(*haxe__io__Encoding)
		}(argument_2))
	case *hxrt.Function[func(int, int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_43 any) int {
			if hx_value_43 == nil {
				var hx_zero_44 int
				return hx_zero_44
			}
			return hx_value_43.(int)
		}(argument_0), func(hx_value_45 any) int {
			if hx_value_45 == nil {
				var hx_zero_46 int
				return hx_zero_46
			}
			return hx_value_45.(int)
		}(argument_1))
	case *hxrt.Function[func(int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_47 any) int {
			if hx_value_47 == nil {
				var hx_zero_48 int
				return hx_zero_48
			}
			return hx_value_47.(int)
		}(argument_0), func(hx_value_49 any) int {
			if hx_value_49 == nil {
				var hx_zero_50 int
				return hx_zero_50
			}
			return hx_value_49.(int)
		}(argument_1))
		return nil
	case *hxrt.Function[func(int, float64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_51 any) int {
			if hx_value_51 == nil {
				var hx_zero_52 int
				return hx_zero_52
			}
			return hx_value_51.(int)
		}(argument_0), func(hx_value_53 any) float64 {
			if hx_value_53 == nil {
				var hx_zero_54 float64
				return hx_zero_54
			}
			return hx_value_53.(float64)
		}(argument_1))
		return nil
	case *hxrt.Function[func(int, *haxe___Int64_____Int64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_55 any) int {
			if hx_value_55 == nil {
				var hx_zero_56 int
				return hx_zero_56
			}
			return hx_value_55.(int)
		}(argument_0), hx_recover_class_haxe___Int64_____Int64(argument_1))
		return nil
	case *hxrt.Function[func(int, int) *haxe__io__Bytes]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_57 any) int {
			if hx_value_57 == nil {
				var hx_zero_58 int
				return hx_zero_58
			}
			return hx_value_57.(int)
		}(argument_0), func(hx_value_59 any) int {
			if hx_value_59 == nil {
				var hx_zero_60 int
				return hx_zero_60
			}
			return hx_value_59.(int)
		}(argument_1))
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(argument_0))
		return nil
	case *hxrt.Function[func(int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_61 any) int {
			if hx_value_61 == nil {
				var hx_zero_62 int
				return hx_zero_62
			}
			return hx_value_61.(int)
		}(argument_0))
		return nil
	case *hxrt.Function[func(*haxe__io__Bytes, int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(argument_0), func(hx_value_63 any) int {
			if hx_value_63 == nil {
				var hx_zero_64 int
				return hx_zero_64
			}
			return hx_value_63.(int)
		}(argument_1), func(hx_value_65 any) int {
			if hx_value_65 == nil {
				var hx_zero_66 int
				return hx_zero_66
			}
			return hx_value_65.(int)
		}(argument_2))
		return nil
	case *hxrt.Function[func(float64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_67 any) float64 {
			if hx_value_67 == nil {
				var hx_zero_68 float64
				return hx_zero_68
			}
			return hx_value_67.(float64)
		}(argument_0))
		return nil
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_haxe___Int64_____Int64(argument_0))
		return nil
	case *hxrt.Function[func(*string, *haxe__io__Encoding)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_69 any) *string {
			if hx_value_69 == nil {
				var hx_zero_70 *string
				return hx_zero_70
			}
			return hx_value_69.(*string)
		}(argument_0), func(hx_value_71 any) *haxe__io__Encoding {
			if hx_value_71 == nil {
				var hx_zero_72 *haxe__io__Encoding
				return hx_zero_72
			}
			return hx_value_71.(*haxe__io__Encoding)
		}(argument_1))
		return nil
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_73 any) int {
			if hx_value_73 == nil {
				var hx_zero_74 int
				return hx_zero_74
			}
			return hx_value_73.(int)
		}(argument_0))
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(*haxe__io__Bytes, int, int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(argument_0), func(hx_value_75 any) int {
			if hx_value_75 == nil {
				var hx_zero_76 int
				return hx_zero_76
			}
			return hx_value_75.(int)
		}(argument_1), func(hx_value_77 any) int {
			if hx_value_77 == nil {
				var hx_zero_78 int
				return hx_zero_78
			}
			return hx_value_77.(int)
		}(argument_2))
	case *hxrt.Function[func() float64]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int, *haxe__io__Encoding) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_79 any) int {
			if hx_value_79 == nil {
				var hx_zero_80 int
				return hx_zero_80
			}
			return hx_value_79.(int)
		}(argument_0), func(hx_value_81 any) *haxe__io__Encoding {
			if hx_value_81 == nil {
				var hx_zero_82 *haxe__io__Encoding
				return hx_zero_82
			}
			return hx_value_81.(*haxe__io__Encoding)
		}(argument_1))
	case *hxrt.Function[func(int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_83 any) int {
			if hx_value_83 == nil {
				var hx_zero_84 int
				return hx_zero_84
			}
			return hx_value_83.(int)
		}(argument_0))
	case *hxrt.Function[func(bool) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_85 any) bool {
			if hx_value_85 == nil {
				var hx_zero_86 bool
				return hx_zero_86
			}
			return hx_value_85.(bool)
		}(argument_0))
	case *hxrt.Function[func(*haxe__io__Input, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Input(argument_0), argument_1)
		return nil
	case *hxrt.Function[func() bool]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int, *sys__io__FileSeek)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_87 any) int {
			if hx_value_87 == nil {
				var hx_zero_88 int
				return hx_zero_88
			}
			return hx_value_87.(int)
		}(argument_0), func(hx_value_89 any) *sys__io__FileSeek {
			if hx_value_89 == nil {
				var hx_zero_90 *sys__io__FileSeek
				return hx_zero_90
			}
			return hx_value_89.(*sys__io__FileSeek)
		}(argument_1))
		return nil
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*sys__net__Host, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(hx_recover_class_sys__net__Host(argument_0), func(hx_value_91 any) int {
			if hx_value_91 == nil {
				var hx_zero_92 int
				return hx_zero_92
			}
			return hx_value_91.(int)
		}(argument_1))
		return nil
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_93 any) *hxrt.SocketHandle {
			if hx_value_93 == nil {
				var hx_zero_94 *hxrt.SocketHandle
				return hx_zero_94
			}
			return hx_value_93.(*hxrt.SocketHandle)
		}(argument_0))
		return nil
	case *hxrt.Function[func(bool)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_95 any) bool {
			if hx_value_95 == nil {
				var hx_zero_96 bool
				return hx_zero_96
			}
			return hx_value_95.(bool)
		}(argument_0))
		return nil
	case *hxrt.Function[func(bool, bool)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_97 any) bool {
			if hx_value_97 == nil {
				var hx_zero_98 bool
				return hx_zero_98
			}
			return hx_value_97.(bool)
		}(argument_0), func(hx_value_99 any) bool {
			if hx_value_99 == nil {
				var hx_zero_100 bool
				return hx_zero_100
			}
			return hx_value_99.(bool)
		}(argument_1))
		return nil
	case *hxrt.Function[func(*string)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_101 any) *string {
			if hx_value_101 == nil {
				var hx_zero_102 *string
				return hx_zero_102
			}
			return hx_value_101.(*string)
		}(argument_0))
		return nil
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_class_app__core__Incident(value any) *app__core__Incident {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__core__Incident:
		return source
	default:
		return value.(*app__core__Incident)
	}
}

func hx_recover_class_app__core__IncidentRequestException(value any) *app__core__IncidentRequestException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__core__IncidentRequestException:
		return source
	default:
		return value.(*app__core__IncidentRequestException)
	}
}

func hx_recover_class_app__http__HttpRequest(value any) *app__http__HttpRequest {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__http__HttpRequest:
		return source
	default:
		return value.(*app__http__HttpRequest)
	}
}

func hx_recover_class_app__http__HttpResponse(value any) *app__http__HttpResponse {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__http__HttpResponse:
		return source
	default:
		return value.(*app__http__HttpResponse)
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__core__IncidentRequestException:
		return source.haxe__Exception
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
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

func hx_recover_class_haxe___Int64_____Int64(value any) *haxe___Int64_____Int64 {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return source
	default:
		return value.(*haxe___Int64_____Int64)
	}
}

func hx_recover_class_haxe__io__Bytes(value any) *haxe__io__Bytes {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return source
	default:
		return value.(*haxe__io__Bytes)
	}
}

func hx_recover_class_haxe__io__Eof(value any) *haxe__io__Eof {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Eof:
		return source
	default:
		return value.(*haxe__io__Eof)
	}
}

func hx_recover_class_haxe__io__Input(value any) *haxe__io__Input {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__net__SocketInput:
		return source.haxe__io__Input
	case *sys__io__FileInput:
		return source.haxe__io__Input
	case *haxe__io__Input:
		return source
	default:
		return value.(*haxe__io__Input)
	}
}

func hx_recover_class_sys__net__Host(value any) *sys__net__Host {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__net__Host:
		return source
	default:
		return value.(*sys__net__Host)
	}
}

func hx_recover_class_sys__net__Socket(value any) *sys__net__Socket {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__net__Socket:
		return source
	default:
		return value.(*sys__net__Socket)
	}
}

func hx_is_nominal_app__core__Incident(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__core__Incident:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_app__core__IncidentApi(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__core__IncidentApi:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_app__core__IncidentConfig(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__core__IncidentConfig:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_app__core__IncidentRequestException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__core__IncidentRequestException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_app__core__IncidentStore(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__core__IncidentStore:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_app__http__HttpRequest(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__http__HttpRequest:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_app__http__HttpResponse(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__http__HttpResponse:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_app__http__TinyHttpServer(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__http__TinyHttpServer:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__core__IncidentRequestException:
		return (source != nil)
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	case *haxe__exceptions__PosException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe___Int64_____Int64(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__exceptions__NotImplementedException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__exceptions__PosException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__exceptions__PosException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__Bytes(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__BytesBuffer(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__BytesBuffer:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__Eof(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__Input(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketInput:
		return (source != nil)
	case *sys__io__FileInput:
		return (source != nil)
	case *haxe__io__Input:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__Output(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketOutput:
		return (source != nil)
	case *sys__io__FileOutput:
		return (source != nil)
	case *haxe__io__Output:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__iterators__StringIterator(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__StringIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__iterators__StringKeyValueIterator(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__StringKeyValueIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__io__FileInput(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__io__FileInput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__io__FileOutput(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__io__FileOutput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__net__Host(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__Host:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__net__Socket(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__Socket:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__net__SocketInput(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketInput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__net__SocketOutput(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketOutput:
		return (source != nil)
	default:
		return false
	}
}
