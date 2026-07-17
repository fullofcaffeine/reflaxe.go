package main

import "examples_pulseforge_portable/hxrt"

func app__core__PulseCodec_enrich(event *app__core__PulseEvent) *app__core__PulseEnrichedEvent {
	severity := app__core__PulseCodec_severityFor(event.value)
	weightedValue := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(event.value) * hxrt.Int32Wrap(severity))))) + hxrt.Int32Wrap(app__core__PulseCodec_regionBoost(event.region)))))
	return New_app__core__PulseEnrichedEvent(event, severity, weightedValue)
}

func app__core__PulseCodec_normalizeToken(value *string, fallback *string) *string {
	trimmed := StringTools_trim(value)
	if hxrt.StringEqualStringPtr(trimmed, hxrt.StringFromLiteral("")) {
		return fallback
	}
	return trimmed
}

func app__core__PulseCodec_parse(frame *app__core__PulseIngressFrame) *app__core__PulseEvent {
	source := app__core__PulseCodec_normalizeToken(frame.source, hxrt.StringFromLiteral("unknown"))
	region := app__core__PulseCodec_normalizeToken(frame.region, hxrt.StringFromLiteral("global"))
	value := frame.value
	if value < 0 {
		value = 0
	}
	return New_app__core__PulseEvent(frame.sequence, source, region, value)
}

func app__core__PulseCodec_regionBoost(region *string) int {
	var hx_switch_27 int
	switch *hxrt.StdString(region) {
	case *hxrt.StdString(hxrt.StringFromLiteral("fra")):
		hx_switch_27 = 3
	case *hxrt.StdString(hxrt.StringFromLiteral("gru")), *hxrt.StdString(hxrt.StringFromLiteral("iad")):
		hx_switch_27 = 2
	case *hxrt.StdString(hxrt.StringFromLiteral("sfo")):
		hx_switch_27 = 1
	default:
		hx_switch_27 = 0
	}
	return hx_switch_27
}

func app__core__PulseCodec_severityFor(value int) int {
	if value >= 12 {
		return 3
	}
	if value >= 8 {
		return 2
	}
	return 1
}
