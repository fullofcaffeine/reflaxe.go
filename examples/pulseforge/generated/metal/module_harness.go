package main

import "examples_pulseforge_metal/hxrt"

func Harness_assertContract(runtime app__runtime__PulseRuntime) *string {
	report := Harness_runReport(runtime, Harness_baselineFrames())
	if !hxrt.StringEqualStringPtr(report.profile, runtime.profileId()) {
		hxrt.Throw(hxrt.StringFromLiteral("profile drift"))
	}
	if !hxrt.StringEqualStringPtr(report.variant, runtime.variantId()) {
		hxrt.Throw(hxrt.StringFromLiteral("variant drift"))
	}
	if !hxrt.StringEqualStringPtr(report.capability, runtime.capabilityId()) {
		hxrt.Throw(hxrt.StringFromLiteral("capability drift"))
	}
	if report.ingestReceived != 8 {
		hxrt.Throw(hxrt.StringFromLiteral("ingest.received drift"))
	}
	if report.ingestAccepted != 8 {
		hxrt.Throw(hxrt.StringFromLiteral("ingest.accepted drift"))
	}
	if report.backpressureEvents != 5 {
		hxrt.Throw(hxrt.StringFromLiteral("ingest.backpressure drift"))
	}
	if report.alertCount != 3 {
		hxrt.Throw(hxrt.StringFromLiteral("alert.count drift"))
	}
	if !hxrt.StringEqualStringPtr(report.alertDigest, hxrt.StringFromLiteral("3,5,7")) {
		hxrt.Throw(hxrt.StringFromLiteral("alert.events drift"))
	}
	var hx_if_1 int
	if hxrt.StringEqualStringPtr(runtime.variantId(), hxrt.StringFromLiteral("go_native")) {
		hx_if_1 = 146
	} else {
		hx_if_1 = 129
	}
	expectedScore := hx_if_1
	if report.runtimeScore != expectedScore {
		hxrt.Throw(hxrt.StringFromLiteral("runtime.score drift"))
	}
	return report.render()
}

func Harness_baselineFrames() *hxrt.Array {
	return hxrt.NewArray(New_app__core__PulseIngressFrame(1, hxrt.StringFromLiteral("edge"), 3, hxrt.StringFromLiteral("iad")), New_app__core__PulseIngressFrame(2, hxrt.StringFromLiteral("api"), 7, hxrt.StringFromLiteral("sfo")), New_app__core__PulseIngressFrame(3, hxrt.StringFromLiteral("db"), 11, hxrt.StringFromLiteral("fra")), New_app__core__PulseIngressFrame(4, hxrt.StringFromLiteral("edge"), 2, hxrt.StringFromLiteral("iad")), New_app__core__PulseIngressFrame(5, hxrt.StringFromLiteral("auth"), 13, hxrt.StringFromLiteral("gru")), New_app__core__PulseIngressFrame(6, hxrt.StringFromLiteral("worker"), 5, hxrt.StringFromLiteral("sfo")), New_app__core__PulseIngressFrame(7, hxrt.StringFromLiteral("api"), 9, hxrt.StringFromLiteral("fra")), New_app__core__PulseIngressFrame(8, hxrt.StringFromLiteral("db"), 4, hxrt.StringFromLiteral("iad")))
}

func Harness_cloneFrames(frames *hxrt.Array) *hxrt.Array {
	out := hxrt.NewArray()
	_g := 0
	for _g < frames.Len() {
		frame := func(hx_value_2 any) *app__core__PulseIngressFrame {
			if hx_value_2 == nil {
				var hx_zero_3 *app__core__PulseIngressFrame
				return hx_zero_3
			}
			return hx_value_2.(*app__core__PulseIngressFrame)
		}(frames.Get(_g))
		_g = int(int32((_g + 1)))
		out.Push(New_app__core__PulseIngressFrame(frame.sequence, frame.source, frame.value, frame.region))
	}
	return out
}

func Harness_run(runtime app__runtime__PulseRuntime) *string {
	return Harness_runReport(runtime, Harness_baselineFrames()).render()
}

func Harness_runReport(runtime app__runtime__PulseRuntime, frames *hxrt.Array) *app__core__PulseReport {
	pipeline := New_app__core__PulsePipeline(runtime)
	return pipeline.run(Harness_cloneFrames(frames))
}

func Harness_runWithFrames(runtime app__runtime__PulseRuntime, frames *hxrt.Array) *string {
	return Harness_runReport(runtime, frames).render()
}
