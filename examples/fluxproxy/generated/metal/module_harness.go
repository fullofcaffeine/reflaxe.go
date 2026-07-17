package main

import "examples_fluxproxy_metal/hxrt"

func Harness_assertBreakerScenario(runtime app__runtime__FluxRuntime) {
	breakerRequests := hxrt.NewArray(New_app__core__FluxRequest(1, hxrt.StringFromLiteral("/breaker/api"), 60, 503), New_app__core__FluxRequest(2, hxrt.StringFromLiteral("/breaker/api"), 65, 502), New_app__core__FluxRequest(3, hxrt.StringFromLiteral("/breaker/api"), 20, 200), New_app__core__FluxRequest(4, hxrt.StringFromLiteral("/breaker/api"), 18, 200))
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
	}
	if !hxrt.StringEqualStringPtr(report.variant, runtime.variantId()) {
		hxrt.Throw(hxrt.StringFromLiteral("variant drift"))
	}
	if !hxrt.StringEqualStringPtr(report.capability, runtime.capabilityId()) {
		hxrt.Throw(hxrt.StringFromLiteral("capability drift"))
	}
	if report.ingressReceived != 8 {
		hxrt.Throw(hxrt.StringFromLiteral("ingress.received drift"))
	}
	if report.ingressAccepted != 8 {
		hxrt.Throw(hxrt.StringFromLiteral("ingress.accepted drift"))
	}
	if report.ingressBackpressure != 5 {
		hxrt.Throw(hxrt.StringFromLiteral("ingress.backpressure drift"))
	}
	if report.proxyRetries != 2 {
		hxrt.Throw(hxrt.StringFromLiteral("proxy.retries drift"))
	}
	if report.rateLimitedCount != 1 {
		hxrt.Throw(hxrt.StringFromLiteral("policy.rate_limited drift"))
	}
	if report.breakerOpenCount != 0 {
		hxrt.Throw(hxrt.StringFromLiteral("policy.breaker_open drift"))
	}
	if report.errorsCount != 3 {
		hxrt.Throw(hxrt.StringFromLiteral("errors.count drift"))
	}
	var hx_if_1 int
	if hxrt.StringEqualStringPtr(runtime.variantId(), hxrt.StringFromLiteral("go_native")) {
		hx_if_1 = 35
	} else {
		hx_if_1 = 26
	}
	expectedScore := hx_if_1
	if report.runtimeScore != expectedScore {
		hxrt.Throw(hxrt.StringFromLiteral("runtime.score drift"))
	}
	Harness_assertBreakerScenario(runtime)
	return report.render()
}

func Harness_baselineRequests() *hxrt.Array {
	return hxrt.NewArray(New_app__core__FluxRequest(1, hxrt.StringFromLiteral("/v1/items"), 30, 200), New_app__core__FluxRequest(2, hxrt.StringFromLiteral("/v1/items"), 70, 503), New_app__core__FluxRequest(3, hxrt.StringFromLiteral("/assets/logo.png"), 12, 200), New_app__core__FluxRequest(4, hxrt.StringFromLiteral("/health"), 4, 200), New_app__core__FluxRequest(5, hxrt.StringFromLiteral("/v1/auth"), 40, 502), New_app__core__FluxRequest(6, hxrt.StringFromLiteral("/v1/items"), 18, 200), New_app__core__FluxRequest(7, hxrt.StringFromLiteral("/assets/main.css"), 9, 200), New_app__core__FluxRequest(8, hxrt.StringFromLiteral("/v1/auth"), 28, 200))
}

func Harness_cloneRequests(requests *hxrt.Array) *hxrt.Array {
	out := hxrt.NewArray()
	_g := 0
	for _g < requests.Len() {
		request := func(hx_value_2 any) *app__core__FluxRequest {
			if hx_value_2 == nil {
				var hx_zero_3 *app__core__FluxRequest
				return hx_zero_3
			}
			return hx_value_2.(*app__core__FluxRequest)
		}(requests.Get(_g))
		_g = int(int32((_g + 1)))
		out.Push(New_app__core__FluxRequest(request.id, request.route, request.latencyMs, request.status))
	}
	return out
}

func Harness_run(runtime app__runtime__FluxRuntime) *string {
	return Harness_runReport(runtime, Harness_baselineRequests()).render()
}

func Harness_runReport(runtime app__runtime__FluxRuntime, requests *hxrt.Array) *app__core__FluxReport {
	pipeline := New_app__core__FluxPipeline(runtime)
	return pipeline.run(Harness_cloneRequests(requests))
}

func Harness_runWithRequests(runtime app__runtime__FluxRuntime, requests *hxrt.Array) *string {
	return Harness_runReport(runtime, requests).render()
}
