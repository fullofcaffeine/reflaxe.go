package main

import "examples_pulseforge_portable/hxrt"

func InteractiveCli_decodeToken(raw *string) *string {
	return StringTools_replace(raw, hxrt.StringFromLiteral("_"), hxrt.StringFromLiteral(" "))
}

func InteractiveCli_failUsage(message *string) {
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error: "), message)))
	hxrt.Println(any(hxrt.StringFromLiteral("run `help` for command syntax")))
}

func InteractiveCli_liveLine(report *app__core__PulseReport) *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("live ingest.received="), report.ingestReceived), hxrt.StringFromLiteral(",ingest.backpressure=")), report.backpressureEvents), hxrt.StringFromLiteral(",alert.count=")), report.alertCount), hxrt.StringFromLiteral(",runtime.score=")), report.runtimeScore)
}

func InteractiveCli_nextSequence(frames *hxrt.Array) int {
	next := 1
	_g := 0
	for _g < frames.Len() {
		frame := func(hx_value_1 any) *app__core__PulseIngressFrame {
			if hx_value_1 == nil {
				var hx_zero_2 *app__core__PulseIngressFrame
				return hx_zero_2
			}
			return hx_value_1.(*app__core__PulseIngressFrame)
		}(frames.Get(_g))
		_g = int(int32((_g + 1)))
		if frame.sequence >= next {
			next = int(int32((hxrt.Int32Wrap(frame.sequence) + hxrt.Int32Wrap(1))))
		}
	}
	return next
}

func InteractiveCli_parsePositiveInt(raw *string) int {
	if hxrt.StringEqualStringPtr(raw, hxrt.StringFromLiteral("")) {
		return -1
	}
	bytes := haxe__io__Bytes_ofString(raw, nil)
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

func InteractiveCli_printHelp(runtime app__runtime__PulseRuntime) {
	hxrt.Println(any(hxrt.StringFromLiteral("commands:")))
	hxrt.Println(any(hxrt.StringFromLiteral("  help")))
	hxrt.Println(any(hxrt.StringFromLiteral("  profile")))
	hxrt.Println(any(hxrt.StringFromLiteral("  reset")))
	hxrt.Println(any(hxrt.StringFromLiteral("  status")))
	hxrt.Println(any(hxrt.StringFromLiteral("  scripted")))
	hxrt.Println(any(hxrt.StringFromLiteral("  ingest <source_token> <value> <region_token>")))
	hxrt.Println(any(hxrt.StringFromLiteral("token note: use '_' for spaces")))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("runtime="), runtime.profileId()), hxrt.StringFromLiteral("/")), runtime.variantId()), hxrt.StringFromLiteral("/")), runtime.capabilityId()))
	hxrt.Println(v)
}

func InteractiveCli_printUsage(runtime app__runtime__PulseRuntime) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("pulseforge interactive command session ("), runtime.profileId()), hxrt.StringFromLiteral(")")))
	hxrt.Println(v)
	hxrt.Println(any(hxrt.StringFromLiteral("run scripted contract mode with: --scripted")))
	hxrt.Println(any(hxrt.StringFromLiteral("commands:")))
	hxrt.Println(any(hxrt.StringFromLiteral("  pulseforge help")))
	hxrt.Println(any(hxrt.StringFromLiteral("  pulseforge profile")))
	hxrt.Println(any(hxrt.StringFromLiteral("  pulseforge status")))
	hxrt.Println(any(hxrt.StringFromLiteral("  pulseforge ingest edge 8 iad status")))
	hxrt.Println(any(hxrt.StringFromLiteral("generated-source invocation:")))
	hxrt.Println(any(hxrt.StringFromLiteral("  go run . --scripted")))
	hxrt.Println(any(hxrt.StringFromLiteral("  go run . status")))
}

func InteractiveCli_run(runtime app__runtime__PulseRuntime) {
	frames := Harness_baselineFrames()
	args := hxrt.ArrayFromValues(func(hx_sort_src_3 []*string) []any {
		hx_sort_out_5 := make([]any, 0, len(hx_sort_src_3))
		for _, hx_sort_item_4 := range hx_sort_src_3 {
			hx_sort_out_5 = append(hx_sort_out_5, hx_sort_item_4)
		}
		return hx_sort_out_5
	}(hxrt.SysArgs()))
	if args.Len() == 0 {
		InteractiveCli_printUsage(runtime)
		return
	}
	i := 0
	for i < args.Len() {
		cmd := func(hx_value_6 any) *string {
			if hx_value_6 == nil {
				var hx_zero_7 *string
				return hx_zero_7
			}
			return hx_value_6.(*string)
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
			frames = Harness_baselineFrames()
			resetReport := InteractiveCli_runReport(runtime, frames)
			hxrt.Println(any(hxrt.StringFromLiteral("ok reset")))
			var v_1 any = any(InteractiveCli_liveLine(resetReport))
			hxrt.Println(v_1)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("status")) {
			statusReport := InteractiveCli_runReport(runtime, frames)
			var v_2 any = any(statusReport.render())
			hxrt.Println(v_2)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("scripted")) {
			var v_3 any = any(Harness_runWithFrames(runtime, frames))
			hxrt.Println(v_3)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("ingest")) {
			if int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3)))) >= args.Len() {
				InteractiveCli_failUsage(hxrt.StringFromLiteral("ingest requires <source_token> <value> <region_token>"))
				return
			}
			source := InteractiveCli_decodeToken(func(hx_value_8 any) *string {
				if hx_value_8 == nil {
					var hx_zero_9 *string
					return hx_zero_9
				}
				return hx_value_8.(*string)
			}(args.Get(int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1)))))))
			value := InteractiveCli_parsePositiveInt(func(hx_value_10 any) *string {
				if hx_value_10 == nil {
					var hx_zero_11 *string
					return hx_zero_11
				}
				return hx_value_10.(*string)
			}(args.Get(int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2)))))))
			if value < 0 {
				InteractiveCli_failUsage(hxrt.StringConcatAny(hxrt.StringFromLiteral("invalid value: "), args.Get(int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2)))))))
				return
			}
			region := InteractiveCli_decodeToken(func(hx_value_14 any) *string {
				if hx_value_14 == nil {
					var hx_zero_15 *string
					return hx_zero_15
				}
				return hx_value_14.(*string)
			}(args.Get(int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3)))))))
			sequence := InteractiveCli_nextSequence(frames)
			frames.Push(New_app__core__PulseIngressFrame(sequence, source, value, region))
			ingestReport := InteractiveCli_runReport(runtime, frames)
			hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("ok ingest seq="), sequence)))
			var v_4 any = any(InteractiveCli_liveLine(ingestReport))
			hxrt.Println(v_4)
			i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(4))))
			continue
		}
		InteractiveCli_failUsage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("unknown command: "), cmd))
		return
	}
}

func InteractiveCli_runReport(runtime app__runtime__PulseRuntime, frames *hxrt.Array) *app__core__PulseReport {
	pipeline := New_app__core__PulsePipeline(runtime)
	return pipeline.run(frames)
}
