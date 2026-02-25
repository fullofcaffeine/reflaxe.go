package main

import "examples_fluxproxy_metal/hxrt"

func Harness_assertBreakerScenario(runtime app__runtime__FluxRuntime) {
	breakerRequests := []*app__core__FluxRequest{New_app__core__FluxRequest(1, hxrt.StringFromLiteral("/breaker/api"), 60, 503), New_app__core__FluxRequest(2, hxrt.StringFromLiteral("/breaker/api"), 65, 502), New_app__core__FluxRequest(3, hxrt.StringFromLiteral("/breaker/api"), 20, 200), New_app__core__FluxRequest(4, hxrt.StringFromLiteral("/breaker/api"), 18, 200)}
	report := Harness_runReport(runtime, breakerRequests)
	if report.breakerOpenCount != 2 {
		hxrt.Throw(hxrt.StringFromLiteral("policy.breaker_open scenario drift"))
	}
	if report.rateLimitedCount != 0 {
		hxrt.Throw(hxrt.StringFromLiteral("policy.rate_limited scenario drift"))
	}
	if report.proxyRetries != 2 {
		hxrt.Throw(hxrt.StringFromLiteral("proxy.retries scenario drift"))
	}
}

func Harness_assertContract(runtime app__runtime__FluxRuntime) *string {
	report := Harness_runReport(runtime, Harness_baselineRequests())
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
	if report.ingressReceived != 8 {
		hxrt.Throw(hxrt.StringFromLiteral("ingress.received drift"))
		var hx_throw_zero_4 *string
		return hx_throw_zero_4
	}
	if report.ingressAccepted != 8 {
		hxrt.Throw(hxrt.StringFromLiteral("ingress.accepted drift"))
		var hx_throw_zero_5 *string
		return hx_throw_zero_5
	}
	if report.ingressBackpressure != 5 {
		hxrt.Throw(hxrt.StringFromLiteral("ingress.backpressure drift"))
		var hx_throw_zero_6 *string
		return hx_throw_zero_6
	}
	if report.proxyRetries != 2 {
		hxrt.Throw(hxrt.StringFromLiteral("proxy.retries drift"))
		var hx_throw_zero_7 *string
		return hx_throw_zero_7
	}
	if report.rateLimitedCount != 1 {
		hxrt.Throw(hxrt.StringFromLiteral("policy.rate_limited drift"))
		var hx_throw_zero_8 *string
		return hx_throw_zero_8
	}
	if report.breakerOpenCount != 0 {
		hxrt.Throw(hxrt.StringFromLiteral("policy.breaker_open drift"))
		var hx_throw_zero_9 *string
		return hx_throw_zero_9
	}
	if report.errorsCount != 3 {
		hxrt.Throw(hxrt.StringFromLiteral("errors.count drift"))
		var hx_throw_zero_10 *string
		return hx_throw_zero_10
	}
	var hx_if_11 int
	if hxrt.StringEqualStringPtr(runtime.variantId(), hxrt.StringFromLiteral("go_native")) {
		hx_if_11 = 35
	} else {
		hx_if_11 = 26
	}
	expectedScore := hx_if_11
	if report.runtimeScore != expectedScore {
		hxrt.Throw(hxrt.StringFromLiteral("runtime.score drift"))
		var hx_throw_zero_12 *string
		return hx_throw_zero_12
	}
	Harness_assertBreakerScenario(runtime)
	return report.render()
}

func Harness_baselineRequests() []*app__core__FluxRequest {
	return []*app__core__FluxRequest{New_app__core__FluxRequest(1, hxrt.StringFromLiteral("/v1/items"), 30, 200), New_app__core__FluxRequest(2, hxrt.StringFromLiteral("/v1/items"), 70, 503), New_app__core__FluxRequest(3, hxrt.StringFromLiteral("/assets/logo.png"), 12, 200), New_app__core__FluxRequest(4, hxrt.StringFromLiteral("/health"), 4, 200), New_app__core__FluxRequest(5, hxrt.StringFromLiteral("/v1/auth"), 40, 502), New_app__core__FluxRequest(6, hxrt.StringFromLiteral("/v1/items"), 18, 200), New_app__core__FluxRequest(7, hxrt.StringFromLiteral("/assets/main.css"), 9, 200), New_app__core__FluxRequest(8, hxrt.StringFromLiteral("/v1/auth"), 28, 200)}
}

func Harness_cloneRequests(requests []*app__core__FluxRequest) []*app__core__FluxRequest {
	out := []*app__core__FluxRequest{}
	_g := 0
	for _g < len(requests) {
		request := requests[_g]
		_g = int(int32((_g + 1)))
		out = append(out, New_app__core__FluxRequest(request.id, request.route, request.latencyMs, request.status))
	}
	return out
}

func Harness_run(runtime app__runtime__FluxRuntime) *string {
	return Harness_runReport(runtime, Harness_baselineRequests()).render()
}

func Harness_runReport(runtime app__runtime__FluxRuntime, requests []*app__core__FluxRequest) *app__core__FluxReport {
	pipeline := New_app__core__FluxPipeline(runtime)
	return pipeline.run(Harness_cloneRequests(requests))
}

func Harness_runWithRequests(runtime app__runtime__FluxRuntime, requests []*app__core__FluxRequest) *string {
	return Harness_runReport(runtime, requests).render()
}
