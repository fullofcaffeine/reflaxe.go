package main

import "examples_fluxproxy_portable/hxrt"

func InteractiveCli_decodeToken(raw *string) *string {
	return StringTools_replace(raw, hxrt.StringFromLiteral("_"), hxrt.StringFromLiteral(" "))
}

func InteractiveCli_failUsage(message *string) {
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error: "), message))
	hxrt.Println(hxrt.StringFromLiteral("run `help` for command syntax"))
}

func InteractiveCli_liveLine(report *app__core__FluxReport) *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("live ingress.received="), report.ingressReceived), hxrt.StringFromLiteral(",ingress.backpressure=")), report.ingressBackpressure), hxrt.StringFromLiteral(",proxy.retries=")), report.proxyRetries), hxrt.StringFromLiteral(",errors.count=")), report.errorsCount), hxrt.StringFromLiteral(",runtime.score=")), report.runtimeScore)
}

func InteractiveCli_nextId(requests []*app__core__FluxRequest) int {
	next := 1
	_g := 0
	for _g < len(requests) {
		request := requests[_g]
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
	hxrt.Println(hxrt.StringFromLiteral("commands:"))
	hxrt.Println(hxrt.StringFromLiteral("  help"))
	hxrt.Println(hxrt.StringFromLiteral("  profile"))
	hxrt.Println(hxrt.StringFromLiteral("  reset"))
	hxrt.Println(hxrt.StringFromLiteral("  status"))
	hxrt.Println(hxrt.StringFromLiteral("  scripted"))
	hxrt.Println(hxrt.StringFromLiteral("  ingest <route_token> <latency_ms> <status_code>"))
	hxrt.Println(hxrt.StringFromLiteral("token note: use '_' for spaces"))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("runtime="), runtime.profileId()), hxrt.StringFromLiteral("/")), runtime.variantId()), hxrt.StringFromLiteral("/")), runtime.capabilityId()))
}

func InteractiveCli_printUsage(runtime app__runtime__FluxRuntime) {
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fluxproxy interactive command session ("), runtime.profileId()), hxrt.StringFromLiteral(")")))
	hxrt.Println(hxrt.StringFromLiteral("run scripted contract mode with: --scripted"))
	hxrt.Println(hxrt.StringFromLiteral("commands:"))
	hxrt.Println(hxrt.StringFromLiteral("  fluxproxy help"))
	hxrt.Println(hxrt.StringFromLiteral("  fluxproxy profile"))
	hxrt.Println(hxrt.StringFromLiteral("  fluxproxy status"))
	hxrt.Println(hxrt.StringFromLiteral("  fluxproxy ingest /v1/items 45 200 status"))
	hxrt.Println(hxrt.StringFromLiteral("generated-source invocation:"))
	hxrt.Println(hxrt.StringFromLiteral("  go run . --scripted"))
	hxrt.Println(hxrt.StringFromLiteral("  go run . status"))
}

func InteractiveCli_run(runtime app__runtime__FluxRuntime) {
	requests := Harness_baselineRequests()
	args := Sys_args()
	if len(args) == 0 {
		InteractiveCli_printUsage(runtime)
		return
	}
	i := 0
	for i < len(args) {
		cmd := args[i]
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("help")) {
			InteractiveCli_printHelp(runtime)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("profile")) {
			hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("profile="), runtime.profileId()), hxrt.StringFromLiteral(",variant=")), runtime.variantId()), hxrt.StringFromLiteral(",capability=")), runtime.capabilityId()))
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("reset")) {
			requests = Harness_baselineRequests()
			resetReport := InteractiveCli_runReport(runtime, requests)
			hxrt.Println(hxrt.StringFromLiteral("ok reset"))
			hxrt.Println(InteractiveCli_liveLine(resetReport))
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("status")) {
			statusReport := InteractiveCli_runReport(runtime, requests)
			hxrt.Println(statusReport.render())
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("scripted")) {
			hxrt.Println(Harness_runWithRequests(runtime, requests))
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("ingest")) {
			if int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3)))) >= len(args) {
				InteractiveCli_failUsage(hxrt.StringFromLiteral("ingest requires <route_token> <latency_ms> <status_code>"))
				return
			}
			route := InteractiveCli_decodeToken(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1))))])
			latency := InteractiveCli_parsePositiveInt(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2))))])
			if latency < 0 {
				InteractiveCli_failUsage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("invalid latency_ms: "), args[int(int32((hxrt.Int32Wrap(i)+hxrt.Int32Wrap(2))))]))
				return
			}
			status := InteractiveCli_parsePositiveInt(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3))))])
			if (status < 100) || (status > 599) {
				InteractiveCli_failUsage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("invalid status_code: "), args[int(int32((hxrt.Int32Wrap(i)+hxrt.Int32Wrap(3))))]))
				return
			}
			requestId := InteractiveCli_nextId(requests)
			requests = append(requests, New_app__core__FluxRequest(requestId, route, latency, status))
			ingestReport := InteractiveCli_runReport(runtime, requests)
			hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("ok ingest id="), requestId))
			hxrt.Println(InteractiveCli_liveLine(ingestReport))
			i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(4))))
			continue
		}
		InteractiveCli_failUsage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("unknown command: "), cmd))
		return
	}
}

func InteractiveCli_runReport(runtime app__runtime__FluxRuntime, requests []*app__core__FluxRequest) *app__core__FluxReport {
	pipeline := New_app__core__FluxPipeline(runtime)
	return pipeline.run(requests)
}
