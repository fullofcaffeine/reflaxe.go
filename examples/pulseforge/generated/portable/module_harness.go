package main

import "examples_pulseforge_portable/hxrt"

func Harness_assertContract(runtime app__runtime__PulseRuntime) *string {
	report := Harness_runReport(runtime, Harness_baselineFrames())
	if !hxrt.StringEqualStringPtr(report.profile, runtime.profileId()) {
		hxrt.Throw(hxrt.StringFromLiteral("profile drift"))
		var hx_throw_zero_1 *string
		return hx_throw_zero_1
	}
	if !hxrt.StringEqualStringPtr(report.variant, runtime.variantId()) {
		hxrt.Throw(hxrt.StringFromLiteral("variant drift"))
		var hx_throw_zero_2 *string
		return hx_throw_zero_2
	}
	if !hxrt.StringEqualStringPtr(report.capability, runtime.capabilityId()) {
		hxrt.Throw(hxrt.StringFromLiteral("capability drift"))
		var hx_throw_zero_3 *string
		return hx_throw_zero_3
	}
	if report.ingestReceived != 8 {
		hxrt.Throw(hxrt.StringFromLiteral("ingest.received drift"))
		var hx_throw_zero_4 *string
		return hx_throw_zero_4
	}
	if report.ingestAccepted != 8 {
		hxrt.Throw(hxrt.StringFromLiteral("ingest.accepted drift"))
		var hx_throw_zero_5 *string
		return hx_throw_zero_5
	}
	if report.backpressureEvents != 5 {
		hxrt.Throw(hxrt.StringFromLiteral("ingest.backpressure drift"))
		var hx_throw_zero_6 *string
		return hx_throw_zero_6
	}
	if report.alertCount != 2 {
		hxrt.Throw(hxrt.StringFromLiteral("alert.count drift"))
		var hx_throw_zero_7 *string
		return hx_throw_zero_7
	}
	if !hxrt.StringEqualStringPtr(report.alertDigest, hxrt.StringFromLiteral("3,5")) {
		hxrt.Throw(hxrt.StringFromLiteral("alert.events drift"))
		var hx_throw_zero_8 *string
		return hx_throw_zero_8
	}
	var hx_if_9 int
	if hxrt.StringEqualStringPtr(runtime.variantId(), hxrt.StringFromLiteral("go_native")) {
		hx_if_9 = 123
	} else {
		hx_if_9 = 108
	}
	expectedScore := hx_if_9
	if report.runtimeScore != expectedScore {
		hxrt.Throw(hxrt.StringFromLiteral("runtime.score drift"))
		var hx_throw_zero_10 *string
		return hx_throw_zero_10
	}
	return report.render()
}

func Harness_baselineFrames() []*app__core__PulseIngressFrame {
	return []*app__core__PulseIngressFrame{New_app__core__PulseIngressFrame(1, hxrt.StringFromLiteral("edge"), 3, hxrt.StringFromLiteral("iad")), New_app__core__PulseIngressFrame(2, hxrt.StringFromLiteral("api"), 7, hxrt.StringFromLiteral("sfo")), New_app__core__PulseIngressFrame(3, hxrt.StringFromLiteral("db"), 11, hxrt.StringFromLiteral("fra")), New_app__core__PulseIngressFrame(4, hxrt.StringFromLiteral("edge"), 2, hxrt.StringFromLiteral("iad")), New_app__core__PulseIngressFrame(5, hxrt.StringFromLiteral("auth"), 13, hxrt.StringFromLiteral("gru")), New_app__core__PulseIngressFrame(6, hxrt.StringFromLiteral("worker"), 5, hxrt.StringFromLiteral("sfo")), New_app__core__PulseIngressFrame(7, hxrt.StringFromLiteral("api"), 9, hxrt.StringFromLiteral("fra")), New_app__core__PulseIngressFrame(8, hxrt.StringFromLiteral("db"), 4, hxrt.StringFromLiteral("iad"))}
}

func Harness_cloneFrames(frames []*app__core__PulseIngressFrame) []*app__core__PulseIngressFrame {
	out := []*app__core__PulseIngressFrame{}
	_g := 0
	for _g < len(frames) {
		frame := frames[_g]
		_g = int(int32((_g + 1)))
		out = append(out, New_app__core__PulseIngressFrame(frame.sequence, frame.source, frame.value, frame.region))
	}
	return out
}

func Harness_run(runtime app__runtime__PulseRuntime) *string {
	return Harness_runReport(runtime, Harness_baselineFrames()).render()
}

func Harness_runReport(runtime app__runtime__PulseRuntime, frames []*app__core__PulseIngressFrame) *app__core__PulseReport {
	pipeline := New_app__core__PulsePipeline(runtime)
	return pipeline.run(Harness_cloneFrames(frames))
}

func Harness_runWithFrames(runtime app__runtime__PulseRuntime, frames []*app__core__PulseIngressFrame) *string {
	return Harness_runReport(runtime, frames).render()
}
