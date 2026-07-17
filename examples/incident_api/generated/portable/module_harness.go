package main

import "examples_incident_api_portable/hxrt"

var Harness_CONFIG_FILE *string = hxrt.StringFromLiteral(".incident_api_scripted_config.json")

var Harness_STATE_FILE *string = hxrt.StringFromLiteral(".incident_api_scripted_state.json")

func Harness_cleanup() {
	path := hxrt.StringFromLiteral(".incident_api_scripted_config.json")
	hxrt.TryCatch(func() {
		if sys__FileSystem_exists(path) {
			sys__FileSystem_deleteFile(path)
		}
	}, func(hx_caught_1 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_1)
		_ = hx_tmp
	})
	path_1 := hxrt.StringFromLiteral(".incident_api_scripted_state.json")
	hxrt.TryCatch(func() {
		if sys__FileSystem_exists(path_1) {
			sys__FileSystem_deleteFile(path_1)
		}
	}, func(hx_caught_3 any) {
		hx_tmp_1 := hxrt.ExceptionCaught(hx_caught_3)
		_ = hx_tmp_1
	})
}

func Harness_request(server *app__http__TinyHttpServer, method *string, path *string, body *string) *string {
	client := New_sys__net__Socket()
	result := hxrt.StringFromLiteral("")
	hxrt.TryCatch(func() {
		client.connect(New_sys__net__Host(server.host), server.port)
		contentLength := haxe__io__Bytes_ofString(body).length
		client.output.writeString(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(method, hxrt.StringFromLiteral(" ")), path), hxrt.StringFromLiteral(" HTTP/1.1\r\n")))
		client.output.writeString(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Host: "), server.host), hxrt.StringFromLiteral(":")), server.port), hxrt.StringFromLiteral("\r\n")))
		client.output.writeString(hxrt.StringFromLiteral("Content-Type: application/json\r\n"))
		client.output.writeString(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("Content-Length: "), contentLength), hxrt.StringFromLiteral("\r\n")))
		client.output.writeString(hxrt.StringFromLiteral("Connection: close\r\n\r\n"))
		client.output.writeString(body)
		client.output.flush()
		server.serveOnce()
		result = Harness_summarize(client.input.readAll().toString())
	}, func(hx_caught_5 any) {
		error := hxrt.ExceptionCaught(hx_caught_5)
		_ = error
		result = hxrt.StringFromLiteral("HTTP/1.1 000 Client Error body={\"error\":\"client_error\"}")
	})
	hxrt.TryCatch(func() {
		client.close()
	}, func(hx_caught_7 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_7)
		_ = hx_tmp
	})
	return result
}

func Harness_run() *string {
	Harness_cleanup()
	app__core__IncidentConfig_saveExample(hxrt.StringFromLiteral(".incident_api_scripted_config.json"))
	rawConfig := sys__io__File_getContent(hxrt.StringFromLiteral(".incident_api_scripted_config.json"))
	config := app__core__IncidentConfig_load(hxrt.StringFromLiteral(".incident_api_scripted_config.json"))
	config.statePath = hxrt.StringFromLiteral(".incident_api_scripted_state.json")
	store := New_app__core__IncidentStore(config.statePath)
	api := New_app__core__IncidentApi(config, store)
	server := New_app__http__TinyHttpServer(api, config.host, config.port)
	out := hxrt.NewArray()
	hxrt.TryCatch(func() {
		out.Push(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("config="), StringTools_trim(rawConfig)))
		out.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("listen="), server.host), hxrt.StringFromLiteral(":<ephemeral>")))
		out.Push(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("health="), Harness_request(server, hxrt.StringFromLiteral("GET"), hxrt.StringFromLiteral("/health"), hxrt.StringFromLiteral(""))))
		out.Push(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("create="), Harness_request(server, hxrt.StringFromLiteral("POST"), hxrt.StringFromLiteral("/incidents"), hxrt.StringFromLiteral("{\"title\":\"Database lag\",\"severity\":\"high\"}"))))
		out.Push(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("list="), Harness_request(server, hxrt.StringFromLiteral("GET"), hxrt.StringFromLiteral("/incidents"), hxrt.StringFromLiteral(""))))
		out.Push(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ack="), Harness_request(server, hxrt.StringFromLiteral("POST"), hxrt.StringFromLiteral("/incidents/1/ack"), hxrt.StringFromLiteral(""))))
		out.Push(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("resolve="), Harness_request(server, hxrt.StringFromLiteral("POST"), hxrt.StringFromLiteral("/incidents/1/resolve"), hxrt.StringFromLiteral(""))))
		out.Push(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("metrics="), Harness_request(server, hxrt.StringFromLiteral("GET"), hxrt.StringFromLiteral("/metrics"), hxrt.StringFromLiteral(""))))
		out.Push(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("state="), StringTools_trim(sys__io__File_getContent(hxrt.StringFromLiteral(".incident_api_scripted_state.json")))))
	}, func(hx_caught_9 any) {
		error := hxrt.ExceptionCaught(hx_caught_9)
		out.Push(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error="), hxrt.StdString(error)))
	})
	server.close()
	Harness_cleanup()
	return hxrt.StringJoinAny(out.Values(), hxrt.StringFromLiteral("\n"))
}

func Harness_summarize(raw *string) *string {
	normalized := StringTools_replace(raw, hxrt.StringFromLiteral("\r\n"), hxrt.StringFromLiteral("\n"))
	sections := hxrt.ArrayFromValues(func(hx_sort_src_21 []*string) []any {
		hx_sort_out_23 := make([]any, 0, len(hx_sort_src_21))
		for _, hx_sort_item_22 := range hx_sort_src_21 {
			hx_sort_out_23 = append(hx_sort_out_23, hx_sort_item_22)
		}
		return hx_sort_out_23
	}(hxrt.StringSplitStringPtr(normalized, hxrt.StringFromLiteral("\n\n"))))
	var hx_if_29 *hxrt.Array
	if sections.Len() > 0 {
		hx_if_29 = hxrt.ArrayFromValues(func(hx_sort_src_26 []*string) []any {
			hx_sort_out_28 := make([]any, 0, len(hx_sort_src_26))
			for _, hx_sort_item_27 := range hx_sort_src_26 {
				hx_sort_out_28 = append(hx_sort_out_28, hx_sort_item_27)
			}
			return hx_sort_out_28
		}(hxrt.StringSplitStringPtr(func(hx_value_24 any) *string {
			if hx_value_24 == nil {
				var hx_zero_25 *string
				return hx_zero_25
			}
			return hx_value_24.(*string)
		}(sections.Get(0)), hxrt.StringFromLiteral("\n"))))
	} else {
		hx_if_29 = hxrt.NewArray()
	}
	headerLines := hx_if_29
	var hx_if_32 *string
	if headerLines.Len() > 0 {
		hx_if_32 = hxrt.StdString(func(hx_value_30 any) *string {
			if hx_value_30 == nil {
				var hx_zero_31 *string
				return hx_zero_31
			}
			return hx_value_30.(*string)
		}(headerLines.Get(0)))
	} else {
		hx_if_32 = hxrt.StringFromLiteral("HTTP/1.1 000 Missing")
	}
	status := hx_if_32
	var hx_if_35 *string
	if sections.Len() > 1 {
		hx_if_35 = hxrt.StdString(func(hx_value_33 any) *string {
			if hx_value_33 == nil {
				var hx_zero_34 *string
				return hx_zero_34
			}
			return hx_value_33.(*string)
		}(sections.Get(1)))
	} else {
		hx_if_35 = hxrt.StringFromLiteral("")
	}
	body := hx_if_35
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(status, hxrt.StringFromLiteral(" body=")), body)
}
