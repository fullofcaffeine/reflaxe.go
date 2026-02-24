package main

import "examples_pulseforge_metal/hxrt"

func InteractiveCli_decodeToken(raw *string) *string {
	return StringTools_replace(raw, hxrt.StringFromLiteral("_"), hxrt.StringFromLiteral(" "))
}

func InteractiveCli_failUsage(message *string) {
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error: "), message))
	hxrt.Println(hxrt.StringFromLiteral("run `help` for command syntax"))
}

func InteractiveCli_liveLine(report *app__core__PulseReport) *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("live ingest.received="), report.ingestReceived), hxrt.StringFromLiteral(",ingest.backpressure=")), report.backpressureEvents), hxrt.StringFromLiteral(",alert.count=")), report.alertCount), hxrt.StringFromLiteral(",runtime.score=")), report.runtimeScore)
}

func InteractiveCli_nextSequence(frames []*app__core__PulseIngressFrame) int {
	next := 1
	_ = next
	_g := 0
	for _g < len(frames) {
		frame := frames[_g]
		_ = frame
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
	bytes := haxe__io__Bytes_ofString(raw)
	_ = bytes
	value := 0
	_ = value
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
	hxrt.Println(hxrt.StringFromLiteral("commands:"))
	hxrt.Println(hxrt.StringFromLiteral("  help"))
	hxrt.Println(hxrt.StringFromLiteral("  profile"))
	hxrt.Println(hxrt.StringFromLiteral("  reset"))
	hxrt.Println(hxrt.StringFromLiteral("  status"))
	hxrt.Println(hxrt.StringFromLiteral("  scripted"))
	hxrt.Println(hxrt.StringFromLiteral("  ingest <source_token> <value> <region_token>"))
	hxrt.Println(hxrt.StringFromLiteral("token note: use '_' for spaces"))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("runtime="), runtime.profileId()), hxrt.StringFromLiteral("/")), runtime.variantId()), hxrt.StringFromLiteral("/")), runtime.capabilityId()))
}

func InteractiveCli_printUsage(runtime app__runtime__PulseRuntime) {
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("pulseforge interactive command session ("), runtime.profileId()), hxrt.StringFromLiteral(")")))
	hxrt.Println(hxrt.StringFromLiteral("run scripted contract mode with: --scripted"))
	hxrt.Println(hxrt.StringFromLiteral("commands:"))
	hxrt.Println(hxrt.StringFromLiteral("  pulseforge help"))
	hxrt.Println(hxrt.StringFromLiteral("  pulseforge profile"))
	hxrt.Println(hxrt.StringFromLiteral("  pulseforge status"))
	hxrt.Println(hxrt.StringFromLiteral("  pulseforge ingest edge 8 iad status"))
	hxrt.Println(hxrt.StringFromLiteral("generated-source invocation:"))
	hxrt.Println(hxrt.StringFromLiteral("  go run . --scripted"))
	hxrt.Println(hxrt.StringFromLiteral("  go run . status"))
}

func InteractiveCli_run(runtime app__runtime__PulseRuntime) {
	frames := Harness_baselineFrames()
	_ = frames
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
			frames = Harness_baselineFrames()
			resetReport := InteractiveCli_runReport(runtime, frames)
			_ = resetReport
			hxrt.Println(hxrt.StringFromLiteral("ok reset"))
			hxrt.Println(InteractiveCli_liveLine(resetReport))
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("status")) {
			statusReport := InteractiveCli_runReport(runtime, frames)
			hxrt.Println(statusReport.render())
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("scripted")) {
			hxrt.Println(Harness_runWithFrames(runtime, frames))
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("ingest")) {
			if int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3)))) >= len(args) {
				InteractiveCli_failUsage(hxrt.StringFromLiteral("ingest requires <source_token> <value> <region_token>"))
				return
			}
			source := InteractiveCli_decodeToken(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1))))])
			_ = source
			value := InteractiveCli_parsePositiveInt(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2))))])
			if value < 0 {
				InteractiveCli_failUsage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("invalid value: "), args[int(int32((hxrt.Int32Wrap(i)+hxrt.Int32Wrap(2))))]))
				return
			}
			region := InteractiveCli_decodeToken(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3))))])
			_ = region
			sequence := InteractiveCli_nextSequence(frames)
			frames = append(frames, New_app__core__PulseIngressFrame(sequence, source, value, region))
			ingestReport := InteractiveCli_runReport(runtime, frames)
			_ = ingestReport
			hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("ok ingest seq="), sequence))
			hxrt.Println(InteractiveCli_liveLine(ingestReport))
			i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(4))))
			continue
		}
		InteractiveCli_failUsage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("unknown command: "), cmd))
		return
	}
}

func InteractiveCli_runReport(runtime app__runtime__PulseRuntime, frames []*app__core__PulseIngressFrame) *app__core__PulseReport {
	pipeline := New_app__core__PulsePipeline(runtime)
	return pipeline.run(frames)
}
