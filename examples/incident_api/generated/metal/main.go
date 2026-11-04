package main

import "examples_incident_api_metal/hxrt"

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
	case *Date:
		return hxrt__generated_method_field__Date(value, key)
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

func hxrt__generated_method_field__Date(value *Date, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "getDate":
		return func() *hxrt.Function[func() int] {
			hx_receiver_257 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_257.__hx_this.getDate()
			}, func() any {
				return hx_receiver_257.__hx_this
			}, "getDate")
		}()
	case "getDay":
		return func() *hxrt.Function[func() int] {
			hx_receiver_258 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_258.__hx_this.getDay()
			}, func() any {
				return hx_receiver_258.__hx_this
			}, "getDay")
		}()
	case "getFullYear":
		return func() *hxrt.Function[func() int] {
			hx_receiver_259 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_259.__hx_this.getFullYear()
			}, func() any {
				return hx_receiver_259.__hx_this
			}, "getFullYear")
		}()
	case "getHours":
		return func() *hxrt.Function[func() int] {
			hx_receiver_260 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_260.__hx_this.getHours()
			}, func() any {
				return hx_receiver_260.__hx_this
			}, "getHours")
		}()
	case "getMinutes":
		return func() *hxrt.Function[func() int] {
			hx_receiver_261 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_261.__hx_this.getMinutes()
			}, func() any {
				return hx_receiver_261.__hx_this
			}, "getMinutes")
		}()
	case "getMonth":
		return func() *hxrt.Function[func() int] {
			hx_receiver_262 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_262.__hx_this.getMonth()
			}, func() any {
				return hx_receiver_262.__hx_this
			}, "getMonth")
		}()
	case "getSeconds":
		return func() *hxrt.Function[func() int] {
			hx_receiver_263 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_263.__hx_this.getSeconds()
			}, func() any {
				return hx_receiver_263.__hx_this
			}, "getSeconds")
		}()
	case "getTime":
		return func() *hxrt.Function[func() float64] {
			hx_receiver_264 := value
			return hxrt.BoundFunction(func() float64 {
				return hx_receiver_264.__hx_this.getTime()
			}, func() any {
				return hx_receiver_264.__hx_this
			}, "getTime")
		}()
	case "getTimezoneOffset":
		return func() *hxrt.Function[func() int] {
			hx_receiver_265 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_265.__hx_this.getTimezoneOffset()
			}, func() any {
				return hx_receiver_265.__hx_this
			}, "getTimezoneOffset")
		}()
	case "getUTCDate":
		return func() *hxrt.Function[func() int] {
			hx_receiver_266 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_266.__hx_this.getUTCDate()
			}, func() any {
				return hx_receiver_266.__hx_this
			}, "getUTCDate")
		}()
	case "getUTCDay":
		return func() *hxrt.Function[func() int] {
			hx_receiver_267 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_267.__hx_this.getUTCDay()
			}, func() any {
				return hx_receiver_267.__hx_this
			}, "getUTCDay")
		}()
	case "getUTCFullYear":
		return func() *hxrt.Function[func() int] {
			hx_receiver_268 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_268.__hx_this.getUTCFullYear()
			}, func() any {
				return hx_receiver_268.__hx_this
			}, "getUTCFullYear")
		}()
	case "getUTCHours":
		return func() *hxrt.Function[func() int] {
			hx_receiver_269 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_269.__hx_this.getUTCHours()
			}, func() any {
				return hx_receiver_269.__hx_this
			}, "getUTCHours")
		}()
	case "getUTCMinutes":
		return func() *hxrt.Function[func() int] {
			hx_receiver_270 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_270.__hx_this.getUTCMinutes()
			}, func() any {
				return hx_receiver_270.__hx_this
			}, "getUTCMinutes")
		}()
	case "getUTCMonth":
		return func() *hxrt.Function[func() int] {
			hx_receiver_271 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_271.__hx_this.getUTCMonth()
			}, func() any {
				return hx_receiver_271.__hx_this
			}, "getUTCMonth")
		}()
	case "getUTCSeconds":
		return func() *hxrt.Function[func() int] {
			hx_receiver_272 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_272.__hx_this.getUTCSeconds()
			}, func() any {
				return hx_receiver_272.__hx_this
			}, "getUTCSeconds")
		}()
	case "localParts":
		return func() *hxrt.Function[func() *hxrt.DateParts] {
			hx_receiver_273 := value
			return hxrt.BoundFunction(func() *hxrt.DateParts {
				return hx_receiver_273.__hx_this.localParts()
			}, func() any {
				return hx_receiver_273.__hx_this
			}, "localParts")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_274 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_274.__hx_this.toString()
			}, func() any {
				return hx_receiver_274.__hx_this
			}, "toString")
		}()
	case "utcParts":
		return func() *hxrt.Function[func() *hxrt.DateParts] {
			hx_receiver_275 := value
			return hxrt.BoundFunction(func() *hxrt.DateParts {
				return hx_receiver_275.__hx_this.utcParts()
			}, func() any {
				return hx_receiver_275.__hx_this
			}, "utcParts")
		}()
	}
	return nil
}

func hxrt__generated_method_field__app__core__Incident(value *app__core__Incident, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "toJson":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_276 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_276.__hx_this.toJson()
			}, func() any {
				return hx_receiver_276.__hx_this
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
			hx_receiver_277 := value
			return hxrt.BoundFunction(func(body *string) *app__http__HttpResponse {
				return hx_receiver_277.__hx_this.createIncident(body)
			}, func() any {
				return hx_receiver_277.__hx_this
			}, "createIncident")
		}()
	case "handle":
		return func() *hxrt.Function[func(*app__http__HttpRequest) *app__http__HttpResponse] {
			hx_receiver_278 := value
			return hxrt.BoundFunction(func(request *app__http__HttpRequest) *app__http__HttpResponse {
				return hx_receiver_278.__hx_this.handle(request)
			}, func() any {
				return hx_receiver_278.__hx_this
			}, "handle")
		}()
	case "updateIncident":
		return func() *hxrt.Function[func(*string, *string) *app__http__HttpResponse] {
			hx_receiver_279 := value
			return hxrt.BoundFunction(func(path *string, action *string) *app__http__HttpResponse {
				return hx_receiver_279.__hx_this.updateIncident(path, action)
			}, func() any {
				return hx_receiver_279.__hx_this
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
			hx_receiver_280 := value
			return hxrt.BoundFunction(func(id int) *app__core__Incident {
				return hx_receiver_280.__hx_this.acknowledge(id)
			}, func() any {
				return hx_receiver_280.__hx_this
			}, "acknowledge")
		}()
	case "create":
		return func() *hxrt.Function[func(*string, *string) *app__core__Incident] {
			hx_receiver_281 := value
			return hxrt.BoundFunction(func(title *string, severity *string) *app__core__Incident {
				return hx_receiver_281.__hx_this.create(title, severity)
			}, func() any {
				return hx_receiver_281.__hx_this
			}, "create")
		}()
	case "find":
		return func() *hxrt.Function[func(int) *app__core__Incident] {
			hx_receiver_282 := value
			return hxrt.BoundFunction(func(id int) *app__core__Incident {
				return hx_receiver_282.__hx_this.find(id)
			}, func() any {
				return hx_receiver_282.__hx_this
			}, "find")
		}()
	case "listJson":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_283 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_283.__hx_this.listJson()
			}, func() any {
				return hx_receiver_283.__hx_this
			}, "listJson")
		}()
	case "load":
		return func() *hxrt.Function[func()] {
			hx_receiver_284 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_284.__hx_this.load()
			}, func() any {
				return hx_receiver_284.__hx_this
			}, "load")
		}()
	case "metricsJson":
		return func() *hxrt.Function[func(*string, int) *string] {
			hx_receiver_285 := value
			return hxrt.BoundFunction(func(serviceName *string, requests int) *string {
				return hx_receiver_285.__hx_this.metricsJson(serviceName, requests)
			}, func() any {
				return hx_receiver_285.__hx_this
			}, "metricsJson")
		}()
	case "resolve":
		return func() *hxrt.Function[func(int) *app__core__Incident] {
			hx_receiver_286 := value
			return hxrt.BoundFunction(func(id int) *app__core__Incident {
				return hx_receiver_286.__hx_this.resolve(id)
			}, func() any {
				return hx_receiver_286.__hx_this
			}, "resolve")
		}()
	case "save":
		return func() *hxrt.Function[func()] {
			hx_receiver_287 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_287.__hx_this.save()
			}, func() any {
				return hx_receiver_287.__hx_this
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
			hx_receiver_288 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_288.__hx_this.close()
			}, func() any {
				return hx_receiver_288.__hx_this
			}, "close")
		}()
	case "readBody":
		return func() *hxrt.Function[func(*sys__net__Socket, int) *string] {
			hx_receiver_289 := value
			return hxrt.BoundFunction(func(peer *sys__net__Socket, length int) *string {
				return hx_receiver_289.__hx_this.readBody(peer, length)
			}, func() any {
				return hx_receiver_289.__hx_this
			}, "readBody")
		}()
	case "readRequest":
		return func() *hxrt.Function[func(*sys__net__Socket) *app__http__HttpRequest] {
			hx_receiver_290 := value
			return hxrt.BoundFunction(func(peer *sys__net__Socket) *app__http__HttpRequest {
				return hx_receiver_290.__hx_this.readRequest(peer)
			}, func() any {
				return hx_receiver_290.__hx_this
			}, "readRequest")
		}()
	case "serveOnce":
		return func() *hxrt.Function[func()] {
			hx_receiver_291 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_291.__hx_this.serveOnce()
			}, func() any {
				return hx_receiver_291.__hx_this
			}, "serveOnce")
		}()
	case "writeResponse":
		return func() *hxrt.Function[func(*sys__net__Socket, *app__http__HttpResponse)] {
			hx_receiver_292 := value
			return hxrt.BoundFunction(func(peer *sys__net__Socket, response *app__http__HttpResponse) {
				hx_receiver_292.__hx_this.writeResponse(peer, response)
			}, func() any {
				return hx_receiver_292.__hx_this
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
			hx_receiver_293 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_293.__hx_this.details()
			}, func() any {
				return hx_receiver_293.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_294 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_294.__hx_this.get_message()
			}, func() any {
				return hx_receiver_294.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_295 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_295.__hx_this.get_native()
			}, func() any {
				return hx_receiver_295.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_296 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_296.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_296.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_297 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_297.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_297.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_298 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_298.__hx_this.toString()
			}, func() any {
				return hx_receiver_298.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_299 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_299.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_299.__hx_this
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
			hx_receiver_300 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_300.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_300.__hx_this
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
			hx_receiver_301 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_301.__hx_this.toString()
			}, func() any {
				return hx_receiver_301.__hx_this
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
			hx_receiver_302 := value
			return hxrt.BoundFunction(func() *hxrt.ByteView {
				return hx_receiver_302.__hx_this.__hx_nativeView()
			}, func() any {
				return hx_receiver_302.__hx_this
			}, "__hx_nativeView")
		}()
	case "blit":
		return func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
			hx_receiver_303 := value
			return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
				hx_receiver_303.__hx_this.blit(pos, src, srcpos, len)
			}, func() any {
				return hx_receiver_303.__hx_this
			}, "blit")
		}()
	case "compare":
		return func() *hxrt.Function[func(*haxe__io__Bytes) int] {
			hx_receiver_304 := value
			return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
				return hx_receiver_304.__hx_this.compare(other)
			}, func() any {
				return hx_receiver_304.__hx_this
			}, "compare")
		}()
	case "fill":
		return func() *hxrt.Function[func(int, int, int)] {
			hx_receiver_305 := value
			return hxrt.BoundFunction(func(pos int, len int, value int) {
				hx_receiver_305.__hx_this.fill(pos, len, value)
			}, func() any {
				return hx_receiver_305.__hx_this
			}, "fill")
		}()
	case "get":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_306 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_306.__hx_this.get(pos)
			}, func() any {
				return hx_receiver_306.__hx_this
			}, "get")
		}()
	case "getData":
		return func() *hxrt.Function[func() []int] {
			hx_receiver_307 := value
			return hxrt.BoundFunction(func() []int {
				return hx_receiver_307.__hx_this.getData()
			}, func() any {
				return hx_receiver_307.__hx_this
			}, "getData")
		}()
	case "getDouble":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_308 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_308.__hx_this.getDouble(pos)
			}, func() any {
				return hx_receiver_308.__hx_this
			}, "getDouble")
		}()
	case "getFloat":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_309 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_309.__hx_this.getFloat(pos)
			}, func() any {
				return hx_receiver_309.__hx_this
			}, "getFloat")
		}()
	case "getInt32":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_310 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_310.__hx_this.getInt32(pos)
			}, func() any {
				return hx_receiver_310.__hx_this
			}, "getInt32")
		}()
	case "getInt64":
		return func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
			hx_receiver_311 := value
			return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
				return hx_receiver_311.__hx_this.getInt64(pos)
			}, func() any {
				return hx_receiver_311.__hx_this
			}, "getInt64")
		}()
	case "getString":
		return func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
			hx_receiver_312 := value
			return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
				return hx_receiver_312.__hx_this.getString(pos, len, encoding)
			}, func() any {
				return hx_receiver_312.__hx_this
			}, "getString")
		}()
	case "getUInt16":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_313 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_313.__hx_this.getUInt16(pos)
			}, func() any {
				return hx_receiver_313.__hx_this
			}, "getUInt16")
		}()
	case "readString":
		return func() *hxrt.Function[func(int, int) *string] {
			hx_receiver_314 := value
			return hxrt.BoundFunction(func(pos int, len int) *string {
				return hx_receiver_314.__hx_this.readString(pos, len)
			}, func() any {
				return hx_receiver_314.__hx_this
			}, "readString")
		}()
	case "set":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_315 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_315.__hx_this.set(pos, value)
			}, func() any {
				return hx_receiver_315.__hx_this
			}, "set")
		}()
	case "setDouble":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_316 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_316.__hx_this.setDouble(pos, value)
			}, func() any {
				return hx_receiver_316.__hx_this
			}, "setDouble")
		}()
	case "setFloat":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_317 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_317.__hx_this.setFloat(pos, value)
			}, func() any {
				return hx_receiver_317.__hx_this
			}, "setFloat")
		}()
	case "setInt32":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_318 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_318.__hx_this.setInt32(pos, value)
			}, func() any {
				return hx_receiver_318.__hx_this
			}, "setInt32")
		}()
	case "setInt64":
		return func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
			hx_receiver_319 := value
			return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
				hx_receiver_319.__hx_this.setInt64(pos, value)
			}, func() any {
				return hx_receiver_319.__hx_this
			}, "setInt64")
		}()
	case "setUInt16":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_320 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_320.__hx_this.setUInt16(pos, value)
			}, func() any {
				return hx_receiver_320.__hx_this
			}, "setUInt16")
		}()
	case "sub":
		return func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
			hx_receiver_321 := value
			return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
				return hx_receiver_321.__hx_this.sub(pos, len)
			}, func() any {
				return hx_receiver_321.__hx_this
			}, "sub")
		}()
	case "toHex":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_322 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_322.__hx_this.toHex()
			}, func() any {
				return hx_receiver_322.__hx_this
			}, "toHex")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_323 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_323.__hx_this.toString()
			}, func() any {
				return hx_receiver_323.__hx_this
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
			hx_receiver_324 := value
			return hxrt.BoundFunction(func(source *haxe__io__Bytes) {
				hx_receiver_324.__hx_this.add(source)
			}, func() any {
				return hx_receiver_324.__hx_this
			}, "add")
		}()
	case "addByte":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_325 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_325.__hx_this.addByte(value)
			}, func() any {
				return hx_receiver_325.__hx_this
			}, "addByte")
		}()
	case "addBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
			hx_receiver_326 := value
			return hxrt.BoundFunction(func(source *haxe__io__Bytes, pos int, len int) {
				hx_receiver_326.__hx_this.addBytes(source, pos, len)
			}, func() any {
				return hx_receiver_326.__hx_this
			}, "addBytes")
		}()
	case "addDouble":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_327 := value
			return hxrt.BoundFunction(func(value float64) {
				hx_receiver_327.__hx_this.addDouble(value)
			}, func() any {
				return hx_receiver_327.__hx_this
			}, "addDouble")
		}()
	case "addFloat":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_328 := value
			return hxrt.BoundFunction(func(value float64) {
				hx_receiver_328.__hx_this.addFloat(value)
			}, func() any {
				return hx_receiver_328.__hx_this
			}, "addFloat")
		}()
	case "addInt32":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_329 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_329.__hx_this.addInt32(value)
			}, func() any {
				return hx_receiver_329.__hx_this
			}, "addInt32")
		}()
	case "addInt64":
		return func() *hxrt.Function[func(*haxe___Int64_____Int64)] {
			hx_receiver_330 := value
			return hxrt.BoundFunction(func(value *haxe___Int64_____Int64) {
				hx_receiver_330.__hx_this.addInt64(value)
			}, func() any {
				return hx_receiver_330.__hx_this
			}, "addInt64")
		}()
	case "addString":
		return func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
			hx_receiver_331 := value
			return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
				hx_receiver_331.__hx_this.addString(value, encoding)
			}, func() any {
				return hx_receiver_331.__hx_this
			}, "addString")
		}()
	case "getBytes":
		return func() *hxrt.Function[func() *haxe__io__Bytes] {
			hx_receiver_332 := value
			return hxrt.BoundFunction(func() *haxe__io__Bytes {
				return hx_receiver_332.__hx_this.getBytes()
			}, func() any {
				return hx_receiver_332.__hx_this
			}, "getBytes")
		}()
	case "get_length":
		return func() *hxrt.Function[func() int] {
			hx_receiver_333 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_333.__hx_this.get_length()
			}, func() any {
				return hx_receiver_333.__hx_this
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
			hx_receiver_334 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_334.__hx_this.toString()
			}, func() any {
				return hx_receiver_334.__hx_this
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
			hx_receiver_335 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_335.__hx_this.close()
			}, func() any {
				return hx_receiver_335.__hx_this
			}, "close")
		}()
	case "read":
		return func() *hxrt.Function[func(int) *haxe__io__Bytes] {
			hx_receiver_336 := value
			return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
				return hx_receiver_336.__hx_this.read(nbytes)
			}, func() any {
				return hx_receiver_336.__hx_this
			}, "read")
		}()
	case "readAll":
		return func() *hxrt.Function[func(any) *haxe__io__Bytes] {
			hx_receiver_337 := value
			return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
				return hx_receiver_337.__hx_this.readAll(bufsize)
			}, func() any {
				return hx_receiver_337.__hx_this
			}, "readAll")
		}()
	case "readByte":
		return func() *hxrt.Function[func() int] {
			hx_receiver_338 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_338.__hx_this.readByte()
			}, func() any {
				return hx_receiver_338.__hx_this
			}, "readByte")
		}()
	case "readBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_339 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
				return hx_receiver_339.__hx_this.readBytes(bytes, pos, len)
			}, func() any {
				return hx_receiver_339.__hx_this
			}, "readBytes")
		}()
	case "readDouble":
		return func() *hxrt.Function[func() float64] {
			hx_receiver_340 := value
			return hxrt.BoundFunction(func() float64 {
				return hx_receiver_340.__hx_this.readDouble()
			}, func() any {
				return hx_receiver_340.__hx_this
			}, "readDouble")
		}()
	case "readFloat":
		return func() *hxrt.Function[func() float64] {
			hx_receiver_341 := value
			return hxrt.BoundFunction(func() float64 {
				return hx_receiver_341.__hx_this.readFloat()
			}, func() any {
				return hx_receiver_341.__hx_this
			}, "readFloat")
		}()
	case "readFullBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
			hx_receiver_342 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
				hx_receiver_342.__hx_this.readFullBytes(bytes, pos, len)
			}, func() any {
				return hx_receiver_342.__hx_this
			}, "readFullBytes")
		}()
	case "readInt16":
		return func() *hxrt.Function[func() int] {
			hx_receiver_343 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_343.__hx_this.readInt16()
			}, func() any {
				return hx_receiver_343.__hx_this
			}, "readInt16")
		}()
	case "readInt24":
		return func() *hxrt.Function[func() int] {
			hx_receiver_344 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_344.__hx_this.readInt24()
			}, func() any {
				return hx_receiver_344.__hx_this
			}, "readInt24")
		}()
	case "readInt32":
		return func() *hxrt.Function[func() int] {
			hx_receiver_345 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_345.__hx_this.readInt32()
			}, func() any {
				return hx_receiver_345.__hx_this
			}, "readInt32")
		}()
	case "readInt8":
		return func() *hxrt.Function[func() int] {
			hx_receiver_346 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_346.__hx_this.readInt8()
			}, func() any {
				return hx_receiver_346.__hx_this
			}, "readInt8")
		}()
	case "readLine":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_347 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_347.__hx_this.readLine()
			}, func() any {
				return hx_receiver_347.__hx_this
			}, "readLine")
		}()
	case "readString":
		return func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
			hx_receiver_348 := value
			return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
				return hx_receiver_348.__hx_this.readString(len, encoding)
			}, func() any {
				return hx_receiver_348.__hx_this
			}, "readString")
		}()
	case "readUInt16":
		return func() *hxrt.Function[func() int] {
			hx_receiver_349 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_349.__hx_this.readUInt16()
			}, func() any {
				return hx_receiver_349.__hx_this
			}, "readUInt16")
		}()
	case "readUInt24":
		return func() *hxrt.Function[func() int] {
			hx_receiver_350 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_350.__hx_this.readUInt24()
			}, func() any {
				return hx_receiver_350.__hx_this
			}, "readUInt24")
		}()
	case "readUntil":
		return func() *hxrt.Function[func(int) *string] {
			hx_receiver_351 := value
			return hxrt.BoundFunction(func(end int) *string {
				return hx_receiver_351.__hx_this.readUntil(end)
			}, func() any {
				return hx_receiver_351.__hx_this
			}, "readUntil")
		}()
	case "set_bigEndian":
		return func() *hxrt.Function[func(bool) bool] {
			hx_receiver_352 := value
			return hxrt.BoundFunction(func(value bool) bool {
				return hx_receiver_352.__hx_this.set_bigEndian(value)
			}, func() any {
				return hx_receiver_352.__hx_this
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
			hx_receiver_353 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_353.__hx_this.close()
			}, func() any {
				return hx_receiver_353.__hx_this
			}, "close")
		}()
	case "flush":
		return func() *hxrt.Function[func()] {
			hx_receiver_354 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_354.__hx_this.flush()
			}, func() any {
				return hx_receiver_354.__hx_this
			}, "flush")
		}()
	case "prepare":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_355 := value
			return hxrt.BoundFunction(func(nbytes int) {
				hx_receiver_355.__hx_this.prepare(nbytes)
			}, func() any {
				return hx_receiver_355.__hx_this
			}, "prepare")
		}()
	case "set_bigEndian":
		return func() *hxrt.Function[func(bool) bool] {
			hx_receiver_356 := value
			return hxrt.BoundFunction(func(value bool) bool {
				return hx_receiver_356.__hx_this.set_bigEndian(value)
			}, func() any {
				return hx_receiver_356.__hx_this
			}, "set_bigEndian")
		}()
	case "write":
		return func() *hxrt.Function[func(*haxe__io__Bytes)] {
			hx_receiver_357 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
				hx_receiver_357.__hx_this.write(bytes)
			}, func() any {
				return hx_receiver_357.__hx_this
			}, "write")
		}()
	case "writeByte":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_358 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_358.__hx_this.writeByte(value)
			}, func() any {
				return hx_receiver_358.__hx_this
			}, "writeByte")
		}()
	case "writeBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_359 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
				return hx_receiver_359.__hx_this.writeBytes(bytes, pos, len)
			}, func() any {
				return hx_receiver_359.__hx_this
			}, "writeBytes")
		}()
	case "writeDouble":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_360 := value
			return hxrt.BoundFunction(func(value float64) {
				hx_receiver_360.__hx_this.writeDouble(value)
			}, func() any {
				return hx_receiver_360.__hx_this
			}, "writeDouble")
		}()
	case "writeFloat":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_361 := value
			return hxrt.BoundFunction(func(value float64) {
				hx_receiver_361.__hx_this.writeFloat(value)
			}, func() any {
				return hx_receiver_361.__hx_this
			}, "writeFloat")
		}()
	case "writeFullBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
			hx_receiver_362 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
				hx_receiver_362.__hx_this.writeFullBytes(bytes, pos, len)
			}, func() any {
				return hx_receiver_362.__hx_this
			}, "writeFullBytes")
		}()
	case "writeInput":
		return func() *hxrt.Function[func(*haxe__io__Input, any)] {
			hx_receiver_363 := value
			return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
				hx_receiver_363.__hx_this.writeInput(input, bufsize)
			}, func() any {
				return hx_receiver_363.__hx_this
			}, "writeInput")
		}()
	case "writeInt16":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_364 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_364.__hx_this.writeInt16(value)
			}, func() any {
				return hx_receiver_364.__hx_this
			}, "writeInt16")
		}()
	case "writeInt24":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_365 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_365.__hx_this.writeInt24(value)
			}, func() any {
				return hx_receiver_365.__hx_this
			}, "writeInt24")
		}()
	case "writeInt32":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_366 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_366.__hx_this.writeInt32(value)
			}, func() any {
				return hx_receiver_366.__hx_this
			}, "writeInt32")
		}()
	case "writeInt8":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_367 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_367.__hx_this.writeInt8(value)
			}, func() any {
				return hx_receiver_367.__hx_this
			}, "writeInt8")
		}()
	case "writeString":
		return func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
			hx_receiver_368 := value
			return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
				hx_receiver_368.__hx_this.writeString(value, encoding)
			}, func() any {
				return hx_receiver_368.__hx_this
			}, "writeString")
		}()
	case "writeUInt16":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_369 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_369.__hx_this.writeUInt16(value)
			}, func() any {
				return hx_receiver_369.__hx_this
			}, "writeUInt16")
		}()
	case "writeUInt24":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_370 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_370.__hx_this.writeUInt24(value)
			}, func() any {
				return hx_receiver_370.__hx_this
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
			hx_receiver_371 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_371.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_371.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() int] {
			hx_receiver_372 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_372.__hx_this.next()
			}, func() any {
				return hx_receiver_372.__hx_this
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
			hx_receiver_373 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_373.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_373.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_374 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_374.__hx_this.next()
			}, func() any {
				return hx_receiver_374.__hx_this
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
			hx_receiver_375 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_375.__hx_this.close()
			}, func() any {
				return hx_receiver_375.__hx_this
			}, "close")
		}()
	case "eof":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_376 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_376.__hx_this.eof()
			}, func() any {
				return hx_receiver_376.__hx_this
			}, "eof")
		}()
	case "readByte":
		return func() *hxrt.Function[func() int] {
			hx_receiver_377 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_377.__hx_this.readByte()
			}, func() any {
				return hx_receiver_377.__hx_this
			}, "readByte")
		}()
	case "readBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_378 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
				return hx_receiver_378.__hx_this.readBytes(bytes, pos, length)
			}, func() any {
				return hx_receiver_378.__hx_this
			}, "readBytes")
		}()
	case "seek":
		return func() *hxrt.Function[func(int, *sys__io__FileSeek)] {
			hx_receiver_379 := value
			return hxrt.BoundFunction(func(p int, pos *sys__io__FileSeek) {
				hx_receiver_379.__hx_this.seek(p, pos)
			}, func() any {
				return hx_receiver_379.__hx_this
			}, "seek")
		}()
	case "tell":
		return func() *hxrt.Function[func() int] {
			hx_receiver_380 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_380.__hx_this.tell()
			}, func() any {
				return hx_receiver_380.__hx_this
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
			hx_receiver_381 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_381.__hx_this.close()
			}, func() any {
				return hx_receiver_381.__hx_this
			}, "close")
		}()
	case "flush":
		return func() *hxrt.Function[func()] {
			hx_receiver_382 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_382.__hx_this.flush()
			}, func() any {
				return hx_receiver_382.__hx_this
			}, "flush")
		}()
	case "seek":
		return func() *hxrt.Function[func(int, *sys__io__FileSeek)] {
			hx_receiver_383 := value
			return hxrt.BoundFunction(func(p int, pos *sys__io__FileSeek) {
				hx_receiver_383.__hx_this.seek(p, pos)
			}, func() any {
				return hx_receiver_383.__hx_this
			}, "seek")
		}()
	case "tell":
		return func() *hxrt.Function[func() int] {
			hx_receiver_384 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_384.__hx_this.tell()
			}, func() any {
				return hx_receiver_384.__hx_this
			}, "tell")
		}()
	case "writeByte":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_385 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_385.__hx_this.writeByte(value)
			}, func() any {
				return hx_receiver_385.__hx_this
			}, "writeByte")
		}()
	case "writeBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_386 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
				return hx_receiver_386.__hx_this.writeBytes(bytes, pos, length)
			}, func() any {
				return hx_receiver_386.__hx_this
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
			hx_receiver_387 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_387.__hx_this.reverse()
			}, func() any {
				return hx_receiver_387.__hx_this
			}, "reverse")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_388 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_388.__hx_this.toString()
			}, func() any {
				return hx_receiver_388.__hx_this
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
			hx_receiver_389 := value
			return hxrt.BoundFunction(func() *sys__net__Socket {
				return hx_receiver_389.__hx_this.accept()
			}, func() any {
				return hx_receiver_389.__hx_this
			}, "accept")
		}()
	case "bind":
		return func() *hxrt.Function[func(*sys__net__Host, int)] {
			hx_receiver_390 := value
			return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
				hx_receiver_390.__hx_this.bind(host, port)
			}, func() any {
				return hx_receiver_390.__hx_this
			}, "bind")
		}()
	case "close":
		return func() *hxrt.Function[func()] {
			hx_receiver_391 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_391.__hx_this.close()
			}, func() any {
				return hx_receiver_391.__hx_this
			}, "close")
		}()
	case "connect":
		return func() *hxrt.Function[func(*sys__net__Host, int)] {
			hx_receiver_392 := value
			return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
				hx_receiver_392.__hx_this.connect(host, port)
			}, func() any {
				return hx_receiver_392.__hx_this
			}, "connect")
		}()
	case "host":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_393 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_393.__hx_this.host()
			}, func() any {
				return hx_receiver_393.__hx_this
			}, "host")
		}()
	case "listen":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_394 := value
			return hxrt.BoundFunction(func(connections int) {
				hx_receiver_394.__hx_this.listen(connections)
			}, func() any {
				return hx_receiver_394.__hx_this
			}, "listen")
		}()
	case "peer":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_395 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_395.__hx_this.peer()
			}, func() any {
				return hx_receiver_395.__hx_this
			}, "peer")
		}()
	case "read":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_396 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_396.__hx_this.read()
			}, func() any {
				return hx_receiver_396.__hx_this
			}, "read")
		}()
	case "replaceHandle":
		return func() *hxrt.Function[func(*hxrt.SocketHandle)] {
			hx_receiver_397 := value
			return hxrt.BoundFunction(func(next *hxrt.SocketHandle) {
				hx_receiver_397.__hx_this.replaceHandle(next)
			}, func() any {
				return hx_receiver_397.__hx_this
			}, "replaceHandle")
		}()
	case "setBlocking":
		return func() *hxrt.Function[func(bool)] {
			hx_receiver_398 := value
			return hxrt.BoundFunction(func(blocking bool) {
				hx_receiver_398.__hx_this.setBlocking(blocking)
			}, func() any {
				return hx_receiver_398.__hx_this
			}, "setBlocking")
		}()
	case "setFastSend":
		return func() *hxrt.Function[func(bool)] {
			hx_receiver_399 := value
			return hxrt.BoundFunction(func(fastSend bool) {
				hx_receiver_399.__hx_this.setFastSend(fastSend)
			}, func() any {
				return hx_receiver_399.__hx_this
			}, "setFastSend")
		}()
	case "setTimeout":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_400 := value
			return hxrt.BoundFunction(func(timeout float64) {
				hx_receiver_400.__hx_this.setTimeout(timeout)
			}, func() any {
				return hx_receiver_400.__hx_this
			}, "setTimeout")
		}()
	case "shutdown":
		return func() *hxrt.Function[func(bool, bool)] {
			hx_receiver_401 := value
			return hxrt.BoundFunction(func(read bool, write bool) {
				hx_receiver_401.__hx_this.shutdown(read, write)
			}, func() any {
				return hx_receiver_401.__hx_this
			}, "shutdown")
		}()
	case "waitForRead":
		return func() *hxrt.Function[func()] {
			hx_receiver_402 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_402.__hx_this.waitForRead()
			}, func() any {
				return hx_receiver_402.__hx_this
			}, "waitForRead")
		}()
	case "write":
		return func() *hxrt.Function[func(*string)] {
			hx_receiver_403 := value
			return hxrt.BoundFunction(func(content *string) {
				hx_receiver_403.__hx_this.write(content)
			}, func() any {
				return hx_receiver_403.__hx_this
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
			hx_receiver_404 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_404.__hx_this.close()
			}, func() any {
				return hx_receiver_404.__hx_this
			}, "close")
		}()
	case "readByte":
		return func() *hxrt.Function[func() int] {
			hx_receiver_405 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_405.__hx_this.readByte()
			}, func() any {
				return hx_receiver_405.__hx_this
			}, "readByte")
		}()
	case "readBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_406 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
				return hx_receiver_406.__hx_this.readBytes(bytes, pos, length)
			}, func() any {
				return hx_receiver_406.__hx_this
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
			hx_receiver_407 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_407.__hx_this.close()
			}, func() any {
				return hx_receiver_407.__hx_this
			}, "close")
		}()
	case "flush":
		return func() *hxrt.Function[func()] {
			hx_receiver_408 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_408.__hx_this.flush()
			}, func() any {
				return hx_receiver_408.__hx_this
			}, "flush")
		}()
	case "writeByte":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_409 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_409.__hx_this.writeByte(value)
			}, func() any {
				return hx_receiver_409.__hx_this
			}, "writeByte")
		}()
	case "writeBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_410 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
				return hx_receiver_410.__hx_this.writeBytes(bytes, pos, length)
			}, func() any {
				return hx_receiver_410.__hx_this
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
	case *Date:
		return hxrt.NewArray(hxrt.StringFromLiteral("ms"))
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

func hxrt__generated_field_lookup__Date(value *Date, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "ms":
		return hxrt.ReadStoredField(value.ms)
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
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() float64]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *hxrt.DateParts]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) *app__http__HttpResponse]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_411 any) *string {
			if hx_value_411 == nil {
				var hx_zero_412 *string
				return hx_zero_412
			}
			return hx_value_411.(*string)
		}(any(argument_0)))
	case *hxrt.Function[func(*app__http__HttpRequest) *app__http__HttpResponse]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_6(argument_0))
	case *hxrt.Function[func(*string, *string) *app__http__HttpResponse]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_413 any) *string {
			if hx_value_413 == nil {
				var hx_zero_414 *string
				return hx_zero_414
			}
			return hx_value_413.(*string)
		}(any(argument_0)), func(hx_value_415 any) *string {
			if hx_value_415 == nil {
				var hx_zero_416 *string
				return hx_zero_416
			}
			return hx_value_415.(*string)
		}(any(argument_1)))
	case *hxrt.Function[func(int) *app__core__Incident]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_417 any) int {
			if hx_value_417 == nil {
				var hx_zero_418 int
				return hx_zero_418
			}
			return hx_value_417.(int)
		}(any(argument_0)))
	case *hxrt.Function[func(*string, *string) *app__core__Incident]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_419 any) *string {
			if hx_value_419 == nil {
				var hx_zero_420 *string
				return hx_zero_420
			}
			return hx_value_419.(*string)
		}(any(argument_0)), func(hx_value_421 any) *string {
			if hx_value_421 == nil {
				var hx_zero_422 *string
				return hx_zero_422
			}
			return hx_value_421.(*string)
		}(any(argument_1)))
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
		return hxrt.FunctionCall(source)(func(hx_value_423 any) *string {
			if hx_value_423 == nil {
				var hx_zero_424 *string
				return hx_zero_424
			}
			return hx_value_423.(*string)
		}(any(argument_0)), func(hx_value_425 any) int {
			if hx_value_425 == nil {
				var hx_zero_426 int
				return hx_zero_426
			}
			return hx_value_425.(int)
		}(any(argument_1)))
	case *hxrt.Function[func(*sys__net__Socket, int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_3(argument_0), func(hx_value_427 any) int {
			if hx_value_427 == nil {
				var hx_zero_428 int
				return hx_zero_428
			}
			return hx_value_427.(int)
		}(any(argument_1)))
	case *hxrt.Function[func(*sys__net__Socket) *app__http__HttpRequest]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_3(argument_0))
	case *hxrt.Function[func(*sys__net__Socket, *app__http__HttpResponse)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(hx_recover_class_3(argument_0), hx_recover_class_7(argument_1))
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
		hxrt.FunctionCall(source)(func(hx_value_429 any) int {
			if hx_value_429 == nil {
				var hx_zero_430 int
				return hx_zero_430
			}
			return hx_value_429.(int)
		}(any(argument_0)), hx_recover_class_8(argument_1), func(hx_value_431 any) int {
			if hx_value_431 == nil {
				var hx_zero_432 int
				return hx_zero_432
			}
			return hx_value_431.(int)
		}(any(argument_2)), func(hx_value_433 any) int {
			if hx_value_433 == nil {
				var hx_zero_434 int
				return hx_zero_434
			}
			return hx_value_433.(int)
		}(any(argument_3)))
		return nil
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_8(argument_0))
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
		hxrt.FunctionCall(source)(func(hx_value_435 any) int {
			if hx_value_435 == nil {
				var hx_zero_436 int
				return hx_zero_436
			}
			return hx_value_435.(int)
		}(any(argument_0)), func(hx_value_437 any) int {
			if hx_value_437 == nil {
				var hx_zero_438 int
				return hx_zero_438
			}
			return hx_value_437.(int)
		}(any(argument_1)), func(hx_value_439 any) int {
			if hx_value_439 == nil {
				var hx_zero_440 int
				return hx_zero_440
			}
			return hx_value_439.(int)
		}(any(argument_2)))
		return nil
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_441 any) int {
			if hx_value_441 == nil {
				var hx_zero_442 int
				return hx_zero_442
			}
			return hx_value_441.(int)
		}(any(argument_0)))
	case *hxrt.Function[func() []int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) float64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_443 any) int {
			if hx_value_443 == nil {
				var hx_zero_444 int
				return hx_zero_444
			}
			return hx_value_443.(int)
		}(any(argument_0)))
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_445 any) int {
			if hx_value_445 == nil {
				var hx_zero_446 int
				return hx_zero_446
			}
			return hx_value_445.(int)
		}(any(argument_0)))
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
		return hxrt.FunctionCall(source)(func(hx_value_447 any) int {
			if hx_value_447 == nil {
				var hx_zero_448 int
				return hx_zero_448
			}
			return hx_value_447.(int)
		}(any(argument_0)), func(hx_value_449 any) int {
			if hx_value_449 == nil {
				var hx_zero_450 int
				return hx_zero_450
			}
			return hx_value_449.(int)
		}(any(argument_1)), func(hx_value_451 any) *haxe__io__Encoding {
			if hx_value_451 == nil {
				var hx_zero_452 *haxe__io__Encoding
				return hx_zero_452
			}
			return hx_value_451.(*haxe__io__Encoding)
		}(any(argument_2)))
	case *hxrt.Function[func(int, int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_453 any) int {
			if hx_value_453 == nil {
				var hx_zero_454 int
				return hx_zero_454
			}
			return hx_value_453.(int)
		}(any(argument_0)), func(hx_value_455 any) int {
			if hx_value_455 == nil {
				var hx_zero_456 int
				return hx_zero_456
			}
			return hx_value_455.(int)
		}(any(argument_1)))
	case *hxrt.Function[func(int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_457 any) int {
			if hx_value_457 == nil {
				var hx_zero_458 int
				return hx_zero_458
			}
			return hx_value_457.(int)
		}(any(argument_0)), func(hx_value_459 any) int {
			if hx_value_459 == nil {
				var hx_zero_460 int
				return hx_zero_460
			}
			return hx_value_459.(int)
		}(any(argument_1)))
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
		hxrt.FunctionCall(source)(func(hx_value_461 any) int {
			if hx_value_461 == nil {
				var hx_zero_462 int
				return hx_zero_462
			}
			return hx_value_461.(int)
		}(any(argument_0)), func(hx_value_463 any) float64 {
			if hx_value_463 == nil {
				var hx_zero_464 float64
				return hx_zero_464
			}
			return hx_value_463.(float64)
		}(any(argument_1)))
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
		hxrt.FunctionCall(source)(func(hx_value_465 any) int {
			if hx_value_465 == nil {
				var hx_zero_466 int
				return hx_zero_466
			}
			return hx_value_465.(int)
		}(any(argument_0)), hx_recover_class_9(argument_1))
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
		return hxrt.FunctionCall(source)(func(hx_value_467 any) int {
			if hx_value_467 == nil {
				var hx_zero_468 int
				return hx_zero_468
			}
			return hx_value_467.(int)
		}(any(argument_0)), func(hx_value_469 any) int {
			if hx_value_469 == nil {
				var hx_zero_470 int
				return hx_zero_470
			}
			return hx_value_469.(int)
		}(any(argument_1)))
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_8(argument_0))
		return nil
	case *hxrt.Function[func(int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_471 any) int {
			if hx_value_471 == nil {
				var hx_zero_472 int
				return hx_zero_472
			}
			return hx_value_471.(int)
		}(any(argument_0)))
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
		hxrt.FunctionCall(source)(hx_recover_class_8(argument_0), func(hx_value_473 any) int {
			if hx_value_473 == nil {
				var hx_zero_474 int
				return hx_zero_474
			}
			return hx_value_473.(int)
		}(any(argument_1)), func(hx_value_475 any) int {
			if hx_value_475 == nil {
				var hx_zero_476 int
				return hx_zero_476
			}
			return hx_value_475.(int)
		}(any(argument_2)))
		return nil
	case *hxrt.Function[func(float64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_477 any) float64 {
			if hx_value_477 == nil {
				var hx_zero_478 float64
				return hx_zero_478
			}
			return hx_value_477.(float64)
		}(any(argument_0)))
		return nil
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_9(argument_0))
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
		hxrt.FunctionCall(source)(func(hx_value_479 any) *string {
			if hx_value_479 == nil {
				var hx_zero_480 *string
				return hx_zero_480
			}
			return hx_value_479.(*string)
		}(any(argument_0)), func(hx_value_481 any) *haxe__io__Encoding {
			if hx_value_481 == nil {
				var hx_zero_482 *haxe__io__Encoding
				return hx_zero_482
			}
			return hx_value_481.(*haxe__io__Encoding)
		}(any(argument_1)))
		return nil
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_483 any) int {
			if hx_value_483 == nil {
				var hx_zero_484 int
				return hx_zero_484
			}
			return hx_value_483.(int)
		}(any(argument_0)))
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
		return hxrt.FunctionCall(source)(hx_recover_class_8(argument_0), func(hx_value_485 any) int {
			if hx_value_485 == nil {
				var hx_zero_486 int
				return hx_zero_486
			}
			return hx_value_485.(int)
		}(any(argument_1)), func(hx_value_487 any) int {
			if hx_value_487 == nil {
				var hx_zero_488 int
				return hx_zero_488
			}
			return hx_value_487.(int)
		}(any(argument_2)))
	case *hxrt.Function[func(int, *haxe__io__Encoding) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_489 any) int {
			if hx_value_489 == nil {
				var hx_zero_490 int
				return hx_zero_490
			}
			return hx_value_489.(int)
		}(any(argument_0)), func(hx_value_491 any) *haxe__io__Encoding {
			if hx_value_491 == nil {
				var hx_zero_492 *haxe__io__Encoding
				return hx_zero_492
			}
			return hx_value_491.(*haxe__io__Encoding)
		}(any(argument_1)))
	case *hxrt.Function[func(int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_493 any) int {
			if hx_value_493 == nil {
				var hx_zero_494 int
				return hx_zero_494
			}
			return hx_value_493.(int)
		}(any(argument_0)))
	case *hxrt.Function[func(bool) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_495 any) bool {
			if hx_value_495 == nil {
				var hx_zero_496 bool
				return hx_zero_496
			}
			return hx_value_495.(bool)
		}(any(argument_0)))
	case *hxrt.Function[func(*haxe__io__Input, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(hx_recover_class_10(argument_0), argument_1)
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
		hxrt.FunctionCall(source)(func(hx_value_497 any) int {
			if hx_value_497 == nil {
				var hx_zero_498 int
				return hx_zero_498
			}
			return hx_value_497.(int)
		}(any(argument_0)), func(hx_value_499 any) *sys__io__FileSeek {
			if hx_value_499 == nil {
				var hx_zero_500 *sys__io__FileSeek
				return hx_zero_500
			}
			return hx_value_499.(*sys__io__FileSeek)
		}(any(argument_1)))
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
		hxrt.FunctionCall(source)(hx_recover_class_11(argument_0), func(hx_value_501 any) int {
			if hx_value_501 == nil {
				var hx_zero_502 int
				return hx_zero_502
			}
			return hx_value_501.(int)
		}(any(argument_1)))
		return nil
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_503 any) *hxrt.SocketHandle {
			if hx_value_503 == nil {
				var hx_zero_504 *hxrt.SocketHandle
				return hx_zero_504
			}
			return hx_value_503.(*hxrt.SocketHandle)
		}(any(argument_0)))
		return nil
	case *hxrt.Function[func(bool)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_505 any) bool {
			if hx_value_505 == nil {
				var hx_zero_506 bool
				return hx_zero_506
			}
			return hx_value_505.(bool)
		}(any(argument_0)))
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
		hxrt.FunctionCall(source)(func(hx_value_507 any) bool {
			if hx_value_507 == nil {
				var hx_zero_508 bool
				return hx_zero_508
			}
			return hx_value_507.(bool)
		}(any(argument_0)), func(hx_value_509 any) bool {
			if hx_value_509 == nil {
				var hx_zero_510 bool
				return hx_zero_510
			}
			return hx_value_509.(bool)
		}(any(argument_1)))
		return nil
	case *hxrt.Function[func(*string)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_511 any) *string {
			if hx_value_511 == nil {
				var hx_zero_512 *string
				return hx_zero_512
			}
			return hx_value_511.(*string)
		}(any(argument_0)))
		return nil
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_class_0(value any) *haxe__Exception {
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

func hx_recover_class_1(value any) *app__core__IncidentRequestException {
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

func hx_recover_class_2(value any) *app__core__Incident {
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

func hx_recover_class_3(value any) *sys__net__Socket {
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

func hx_recover_class_4(value any) *haxe__io__Eof {
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

func hx_recover_class_5(value any) *haxe__ValueException {
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

func hx_recover_class_6(value any) *app__http__HttpRequest {
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

func hx_recover_class_7(value any) *app__http__HttpResponse {
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

func hx_recover_class_8(value any) *haxe__io__Bytes {
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

func hx_recover_class_9(value any) *haxe___Int64_____Int64 {
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

func hx_recover_class_10(value any) *haxe__io__Input {
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

func hx_recover_class_11(value any) *sys__net__Host {
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

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__core__IncidentRequestException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_2(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_3(value any) bool {
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

func hx_is_nominal_4(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *Date:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_5(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__core__Incident:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_6(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__core__IncidentApi:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_7(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__core__IncidentConfig:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_8(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__core__IncidentStore:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_9(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__http__HttpRequest:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_10(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__http__HttpResponse:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_11(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *app__http__TinyHttpServer:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_12(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_13(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_14(value any) bool {
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

func hx_is_nominal_15(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_16(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__BytesBuffer:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_17(value any) bool {
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

func hx_is_nominal_18(value any) bool {
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

func hx_is_nominal_19(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__StringIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_20(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__StringKeyValueIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_21(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__io__FileInput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_22(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__io__FileOutput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_23(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__Host:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_24(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__Socket:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_25(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketInput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_26(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketOutput:
		return (source != nil)
	default:
		return false
	}
}
