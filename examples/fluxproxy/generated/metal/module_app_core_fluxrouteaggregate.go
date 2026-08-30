package main

import "examples_fluxproxy_metal/hxrt"

type I_app__core__FluxRouteAggregate interface {
	record(response *app__core__FluxProxyResponse)
	averageLatencyMs() int
	summaryToken() *string
}

type app__core__FluxRouteAggregate struct {
	__hx_this      I_app__core__FluxRouteAggregate
	route          *string
	count          int
	successCount   int
	errorCount     int
	totalLatencyMs int
}

func New_app__core__FluxRouteAggregate(route *string) *app__core__FluxRouteAggregate {
	self := &app__core__FluxRouteAggregate{}
	self.__hx_this = self
	self.route = route
	self.count = 0
	self.successCount = 0
	self.errorCount = 0
	self.totalLatencyMs = 0
	return self
}

func (self *app__core__FluxRouteAggregate) record(response *app__core__FluxProxyResponse) {
	self.count = int(int32((self.count + 1)))
	self.totalLatencyMs = int((hxrt.Int32Wrap(self.totalLatencyMs) + hxrt.Int32Wrap(response.latencyMs)))
	if response.success {
		self.successCount = int(int32((self.successCount + 1)))
	} else {
		self.errorCount = int(int32((self.errorCount + 1)))
	}
}

func (self *app__core__FluxRouteAggregate) averageLatencyMs() int {
	if self.count == 0 {
		return 0
	}
	remaining := self.totalLatencyMs
	quotient := 0
	for remaining >= self.count {
		remaining = int((hxrt.Int32Wrap(remaining) - hxrt.Int32Wrap(self.count)))
		quotient = int(int32((quotient + 1)))
	}
	return quotient
}

func (self *app__core__FluxRouteAggregate) summaryToken() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(self.route, hxrt.StringFromLiteral(":")), self.count), hxrt.StringFromLiteral("/")), self.successCount), hxrt.StringFromLiteral("/")), self.errorCount), hxrt.StringFromLiteral("/")), self.averageLatencyMs())
}
