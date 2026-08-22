package main

import "examples_fluxproxy_portable/hxrt"

func app__core__FluxCodec_breakerOpen(request *app__core__FluxRequest) *app__core__FluxProxyResponse {
	route := app__core__FluxCodec_normalizedRoute(request.route)
	return New_app__core__FluxProxyResponse(request.id, route, hxrt.StringFromLiteral("breaker-open"), 503, 0, 1, false)
}

func app__core__FluxCodec_normalizedRoute(route *string) *string {
	trimmed := StringTools_trim(route)
	if hxrt.StringEqualStringPtr(trimmed, hxrt.StringFromLiteral("")) {
		return hxrt.StringFromLiteral("/unknown")
	}
	return trimmed
}

func app__core__FluxCodec_proxy(request *app__core__FluxRequest, timeoutMs int) *app__core__FluxProxyResponse {
	route := app__core__FluxCodec_normalizedRoute(request.route)
	latency := request.latencyMs
	if latency < 0 {
		latency = 0
	}
	status := request.status
	var hx_if_1 int
	if status >= 500 {
		hx_if_1 = 2
	} else {
		hx_if_1 = 1
	}
	attempts := hx_if_1
	if latency > timeoutMs {
		status = 504
		attempts = 2
	}
	success := (status < 500)
	return New_app__core__FluxProxyResponse(request.id, route, app__core__FluxCodec_upstreamForRoute(route), status, latency, attempts, success)
}

func app__core__FluxCodec_rateLimited(request *app__core__FluxRequest) *app__core__FluxProxyResponse {
	route := app__core__FluxCodec_normalizedRoute(request.route)
	latency := request.latencyMs
	if latency < 0 {
		latency = 0
	}
	return New_app__core__FluxProxyResponse(request.id, route, hxrt.StringFromLiteral("rate-limit"), 429, latency, 1, false)
}

func app__core__FluxCodec_upstreamForRoute(route *string) *string {
	if StringTools_startsWith(route, hxrt.StringFromLiteral("/assets")) {
		return hxrt.StringFromLiteral("cdn")
	}
	if hxrt.StringEqualStringPtr(route, hxrt.StringFromLiteral("/health")) {
		return hxrt.StringFromLiteral("healthz")
	}
	return hxrt.StringFromLiteral("core-api")
}
