package main

import "examples_fluxproxy_metal/hxrt"

func InteractiveCli_decodeToken(raw *string) *string {
	return StringTools_replace(raw, hxrt.StringFromLiteral("_"), hxrt.StringFromLiteral(" "))
}

func InteractiveCli_failUsage(message *string) {
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error: "), message)))
	hxrt.Println(any(hxrt.StringFromLiteral("run `help` for command syntax")))
}

func InteractiveCli_liveLine(report *app__core__FluxReport) *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("live ingress.received="), report.ingressReceived), hxrt.StringFromLiteral(",ingress.backpressure=")), report.ingressBackpressure), hxrt.StringFromLiteral(",proxy.retries=")), report.proxyRetries), hxrt.StringFromLiteral(",errors.count=")), report.errorsCount), hxrt.StringFromLiteral(",runtime.score=")), report.runtimeScore)
}

func InteractiveCli_nextId(requests *hxrt.Array) int {
	next := 1
	_g := 0
	for _g < requests.Len() {
		request := func(hx_value_5 any) *app__core__FluxRequest {
			if hx_value_5 == nil {
				var hx_zero_6 *app__core__FluxRequest
				return hx_zero_6
			}
			return hx_value_5.(*app__core__FluxRequest)
		}(requests.Get(_g))
		_g = int(int32((_g + 1)))
		if request.id >= next {
			next = int(int32((hxrt.Int32Wrap(request.id) + hxrt.Int32Wrap(1))))
		}
	}
	return next
}

func InteractiveCli_parsePositiveInt(raw *string) int {
	if hxrt.StringEqualStringPtr(raw, hxrt.StringFromLiteral("")) {
		return -1
	}
	bytes := haxe__io__Bytes_ofString(raw)
	value := 0
	i := 0
	for i < bytes.length {
		code := bytes.b[i]
		if (code < 48) || (code > 57) {
			return -1
		}
		value = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(10))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(48))))))))
		i = int(int32((i + 1)))
	}
	return value
}

func InteractiveCli_printHelp(runtime app__runtime__FluxRuntime) {
	hxrt.Println(any(hxrt.StringFromLiteral("commands:")))
	hxrt.Println(any(hxrt.StringFromLiteral("  help")))
	hxrt.Println(any(hxrt.StringFromLiteral("  profile")))
	hxrt.Println(any(hxrt.StringFromLiteral("  reset")))
	hxrt.Println(any(hxrt.StringFromLiteral("  status")))
	hxrt.Println(any(hxrt.StringFromLiteral("  scripted")))
	hxrt.Println(any(hxrt.StringFromLiteral("  ingest <route_token> <latency_ms> <status_code>")))
	hxrt.Println(any(hxrt.StringFromLiteral("token note: use '_' for spaces")))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("runtime="), runtime.profileId()), hxrt.StringFromLiteral("/")), runtime.variantId()), hxrt.StringFromLiteral("/")), runtime.capabilityId()))
	hxrt.Println(v)
}

func InteractiveCli_printUsage(runtime app__runtime__FluxRuntime) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fluxproxy interactive command session ("), runtime.profileId()), hxrt.StringFromLiteral(")")))
	hxrt.Println(v)
	hxrt.Println(any(hxrt.StringFromLiteral("run scripted contract mode with: --scripted")))
	hxrt.Println(any(hxrt.StringFromLiteral("commands:")))
	hxrt.Println(any(hxrt.StringFromLiteral("  fluxproxy help")))
	hxrt.Println(any(hxrt.StringFromLiteral("  fluxproxy profile")))
	hxrt.Println(any(hxrt.StringFromLiteral("  fluxproxy status")))
	hxrt.Println(any(hxrt.StringFromLiteral("  fluxproxy ingest /v1/items 45 200 status")))
	hxrt.Println(any(hxrt.StringFromLiteral("generated-source invocation:")))
	hxrt.Println(any(hxrt.StringFromLiteral("  go run . --scripted")))
	hxrt.Println(any(hxrt.StringFromLiteral("  go run . status")))
}

func InteractiveCli_run(runtime app__runtime__FluxRuntime) {
	requests := Harness_baselineRequests()
	args := hxrt.ArrayFromValues(func(hx_sort_src_7 []*string) []any {
		hx_sort_out_9 := make([]any, 0, len(hx_sort_src_7))
		for _, hx_sort_item_8 := range hx_sort_src_7 {
			hx_sort_out_9 = append(hx_sort_out_9, hx_sort_item_8)
		}
		return hx_sort_out_9
	}(hxrt.SysArgs()))
	if args.Len() == 0 {
		InteractiveCli_printUsage(runtime)
		return
	}
	i := 0
	for i < args.Len() {
		cmd := func(hx_value_10 any) *string {
			if hx_value_10 == nil {
				var hx_zero_11 *string
				return hx_zero_11
			}
			return hx_value_10.(*string)
		}(args.Get(i))
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("help")) {
			InteractiveCli_printHelp(runtime)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("profile")) {
			var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("profile="), runtime.profileId()), hxrt.StringFromLiteral(",variant=")), runtime.variantId()), hxrt.StringFromLiteral(",capability=")), runtime.capabilityId()))
			hxrt.Println(v)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("reset")) {
			requests = Harness_baselineRequests()
			resetReport := InteractiveCli_runReport(runtime, requests)
			hxrt.Println(any(hxrt.StringFromLiteral("ok reset")))
			var v_1 any = any(InteractiveCli_liveLine(resetReport))
			hxrt.Println(v_1)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("status")) {
			statusReport := InteractiveCli_runReport(runtime, requests)
			var v_2 any = any(statusReport.render())
			hxrt.Println(v_2)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("scripted")) {
			var v_3 any = any(Harness_runWithRequests(runtime, requests))
			hxrt.Println(v_3)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("ingest")) {
			if int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3)))) >= args.Len() {
				InteractiveCli_failUsage(hxrt.StringFromLiteral("ingest requires <route_token> <latency_ms> <status_code>"))
				return
			}
			route := InteractiveCli_decodeToken(func(hx_value_12 any) *string {
				if hx_value_12 == nil {
					var hx_zero_13 *string
					return hx_zero_13
				}
				return hx_value_12.(*string)
			}(args.Get(int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1)))))))
			latency := InteractiveCli_parsePositiveInt(func(hx_value_14 any) *string {
				if hx_value_14 == nil {
					var hx_zero_15 *string
					return hx_zero_15
				}
				return hx_value_14.(*string)
			}(args.Get(int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2)))))))
			if latency < 0 {
				InteractiveCli_failUsage(hxrt.StringConcatAny(hxrt.StringFromLiteral("invalid latency_ms: "), args.Get(int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2)))))))
				return
			}
			status := InteractiveCli_parsePositiveInt(func(hx_value_18 any) *string {
				if hx_value_18 == nil {
					var hx_zero_19 *string
					return hx_zero_19
				}
				return hx_value_18.(*string)
			}(args.Get(int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3)))))))
			if (status < 100) || (status > 599) {
				InteractiveCli_failUsage(hxrt.StringConcatAny(hxrt.StringFromLiteral("invalid status_code: "), args.Get(int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3)))))))
				return
			}
			requestId := InteractiveCli_nextId(requests)
			requests.Push(New_app__core__FluxRequest(requestId, route, latency, status))
			ingestReport := InteractiveCli_runReport(runtime, requests)
			hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("ok ingest id="), requestId)))
			var v_4 any = any(InteractiveCli_liveLine(ingestReport))
			hxrt.Println(v_4)
			i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(4))))
			continue
		}
		InteractiveCli_failUsage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("unknown command: "), cmd))
		return
	}
}

func InteractiveCli_runReport(runtime app__runtime__FluxRuntime, requests *hxrt.Array) *app__core__FluxReport {
	pipeline := New_app__core__FluxPipeline(runtime)
	return pipeline.run(requests)
}
