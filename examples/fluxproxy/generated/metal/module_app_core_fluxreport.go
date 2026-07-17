package main

import "examples_fluxproxy_metal/hxrt"

type I_app__core__FluxReport interface {
	lines() *hxrt.Array
	profileId() *string
	variantId() *string
	capabilityId() *string
	receivedCount() int
	acceptedCount() int
	backpressureCount() int
	retriesCount() int
	rateLimited() int
	breakerOpen() int
	errors() int
	score() int
	render() *string
}

type app__core__FluxReport struct {
	__hx_this           I_app__core__FluxReport
	profile             *string
	variant             *string
	capability          *string
	ingressReceived     int
	ingressAccepted     int
	ingressBackpressure int
	proxyResponses      int
	proxyRetries        int
	rateLimitedCount    int
	breakerOpenCount    int
	routesCount         int
	routesSummary       *string
	errorsCount         int
	runtimeScore        int
}

func New_app__core__FluxReport(profile *string, variant *string, capability *string, ingressReceived int, ingressAccepted int, ingressBackpressure int, proxyResponses int, proxyRetries int, rateLimitedCount int, breakerOpenCount int, routesCount int, routesSummary *string, errorsCount int, runtimeScore int) *app__core__FluxReport {
	self := &app__core__FluxReport{}
	self.__hx_this = self
	self.profile = profile
	self.variant = variant
	self.capability = capability
	self.ingressReceived = ingressReceived
	self.ingressAccepted = ingressAccepted
	self.ingressBackpressure = ingressBackpressure
	self.proxyResponses = proxyResponses
	self.proxyRetries = proxyRetries
	self.rateLimitedCount = rateLimitedCount
	self.breakerOpenCount = breakerOpenCount
	self.routesCount = routesCount
	self.routesSummary = routesSummary
	self.errorsCount = errorsCount
	self.runtimeScore = runtimeScore
	return self
}

func (self *app__core__FluxReport) lines() *hxrt.Array {
	return hxrt.NewArray(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fluxproxy.profile="), self.profile), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fluxproxy.variant="), self.variant), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("runtime.capability="), self.capability), hxrt.StringConcatAny(hxrt.StringFromLiteral("ingress.received="), self.ingressReceived), hxrt.StringConcatAny(hxrt.StringFromLiteral("ingress.accepted="), self.ingressAccepted), hxrt.StringConcatAny(hxrt.StringFromLiteral("ingress.backpressure="), self.ingressBackpressure), hxrt.StringConcatAny(hxrt.StringFromLiteral("proxy.responses="), self.proxyResponses), hxrt.StringConcatAny(hxrt.StringFromLiteral("proxy.retries="), self.proxyRetries), hxrt.StringConcatAny(hxrt.StringFromLiteral("policy.rate_limited="), self.rateLimitedCount), hxrt.StringConcatAny(hxrt.StringFromLiteral("policy.breaker_open="), self.breakerOpenCount), hxrt.StringConcatAny(hxrt.StringFromLiteral("routes.count="), self.routesCount), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("routes.summary="), self.routesSummary), hxrt.StringConcatAny(hxrt.StringFromLiteral("errors.count="), self.errorsCount), hxrt.StringConcatAny(hxrt.StringFromLiteral("runtime.score="), self.runtimeScore))
}

func (self *app__core__FluxReport) profileId() *string {
	return self.profile
}

func (self *app__core__FluxReport) variantId() *string {
	return self.variant
}

func (self *app__core__FluxReport) capabilityId() *string {
	return self.capability
}

func (self *app__core__FluxReport) receivedCount() int {
	return self.ingressReceived
}

func (self *app__core__FluxReport) acceptedCount() int {
	return self.ingressAccepted
}

func (self *app__core__FluxReport) backpressureCount() int {
	return self.ingressBackpressure
}

func (self *app__core__FluxReport) retriesCount() int {
	return self.proxyRetries
}

func (self *app__core__FluxReport) rateLimited() int {
	return self.rateLimitedCount
}

func (self *app__core__FluxReport) breakerOpen() int {
	return self.breakerOpenCount
}

func (self *app__core__FluxReport) errors() int {
	return self.errorsCount
}

func (self *app__core__FluxReport) score() int {
	return self.runtimeScore
}

func (self *app__core__FluxReport) render() *string {
	out := hxrt.StringFromLiteral("")
	values := self.lines()
	i := 0
	for i < values.Len() {
		if i > 0 {
			out = hxrt.StringConcatStringPtr(out, hxrt.StringFromLiteral("\n"))
		}
		out = hxrt.StringConcatAny(out, values.Get(i))
		i = int(int32((i + 1)))
	}
	return out
}
