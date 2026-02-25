package main

import "examples_fluxproxy_metal/hxrt"

type I_app__core__FluxPipeline interface {
	run(requests []*app__core__FluxRequest) *app__core__FluxReport
	ingest(requests []*app__core__FluxRequest, capacity int) *app__core__FluxIngestResult
	aggregate(responses []*app__core__FluxProxyResponse) map[string]any
	retries(responses []*app__core__FluxProxyResponse) int
	errors(responses []*app__core__FluxProxyResponse) int
	applyRoutePolicies(requests []*app__core__FluxRequest, perRouteLimit int, breakerFailureThreshold int, timeoutMs int) map[string]any
	orderedResponses(synthetic []*app__core__FluxProxyResponse, dispatched []*app__core__FluxProxyResponse, acceptedRequests []*app__core__FluxRequest) []*app__core__FluxProxyResponse
}

type app__core__FluxPipeline struct {
	__hx_this I_app__core__FluxPipeline
	runtime   app__runtime__FluxRuntime
}

func New_app__core__FluxPipeline(runtime app__runtime__FluxRuntime) *app__core__FluxPipeline {
	self := &app__core__FluxPipeline{}
	self.__hx_this = self
	self.runtime = runtime
	return self
}

func (self *app__core__FluxPipeline) run(requests []*app__core__FluxRequest) *app__core__FluxReport {
	ingest := self.ingest(requests, 3)
	planned := self.applyRoutePolicies(ingest.acceptedRequests, 2, 2, 50)
	dispatched := self.runtime.dispatch(func(hx_obj_14 map[string]any) []*app__core__FluxRequest {
		hx_field_15 := hx_obj_14["dispatchable"]
		if hx_field_15 == nil {
			var hx_zero_16 []*app__core__FluxRequest
			return hx_zero_16
		}
		return hx_field_15.([]*app__core__FluxRequest)
	}(planned), 1)
	responses := self.orderedResponses(func(hx_obj_17 map[string]any) []*app__core__FluxProxyResponse {
		hx_field_18 := hx_obj_17["synthetic"]
		if hx_field_18 == nil {
			var hx_zero_19 []*app__core__FluxProxyResponse
			return hx_zero_19
		}
		return hx_field_18.([]*app__core__FluxProxyResponse)
	}(planned), dispatched, ingest.acceptedRequests)
	aggregates := self.aggregate(responses)
	retryCount := self.retries(responses)
	errorCount := self.errors(responses)
	score := self.runtime.stageScore(responses, retryCount, ingest.backpressureEvents)
	return New_app__core__FluxReport(self.runtime.profileId(), self.runtime.variantId(), self.runtime.capabilityId(), ingest.receivedCount, len(ingest.acceptedRequests), ingest.backpressureEvents, len(responses), retryCount, func(hx_obj_20 map[string]any) int {
		hx_field_21 := hx_obj_20["rateLimited"]
		if hx_field_21 == nil {
			var hx_zero_22 int
			return hx_zero_22
		}
		return hx_field_21.(int)
	}(planned), func(hx_obj_23 map[string]any) int {
		hx_field_24 := hx_obj_23["breakerOpen"]
		if hx_field_24 == nil {
			var hx_zero_25 int
			return hx_zero_25
		}
		return hx_field_24.(int)
	}(planned), len(func(hx_obj_26 map[string]any) []*app__core__FluxRouteAggregate {
		hx_field_27 := hx_obj_26["routes"]
		if hx_field_27 == nil {
			var hx_zero_28 []*app__core__FluxRouteAggregate
			return hx_zero_28
		}
		return hx_field_27.([]*app__core__FluxRouteAggregate)
	}(aggregates)), func(hx_obj_29 map[string]any) *string {
		hx_field_30 := hx_obj_29["summary"]
		if hx_field_30 == nil {
			var hx_zero_31 *string
			return hx_zero_31
		}
		return hx_field_30.(*string)
	}(aggregates), errorCount, score)
}

func (self *app__core__FluxPipeline) ingest(requests []*app__core__FluxRequest, capacity int) *app__core__FluxIngestResult {
	var hx_if_32 int
	if capacity <= 0 {
		hx_if_32 = 1
	} else {
		hx_if_32 = capacity
	}
	boundedCapacity := hx_if_32
	queue := []*app__core__FluxRequest{}
	queueHead := 0
	accepted := []*app__core__FluxRequest{}
	backpressureEvents := 0
	_g := 0
	for _g < len(requests) {
		request := requests[_g]
		_g = int(int32((_g + 1)))
		if int(int32((hxrt.Int32Wrap(len(queue)) - hxrt.Int32Wrap(queueHead)))) >= boundedCapacity {
			backpressureEvents = int(int32((backpressureEvents + 1)))
			accepted = append(accepted, queue[queueHead])
			queueHead = int(int32((queueHead + 1)))
		}
		queue = append(queue, request)
	}
	for queueHead < len(queue) {
		accepted = append(accepted, queue[queueHead])
		queueHead = int(int32((queueHead + 1)))
	}
	return New_app__core__FluxIngestResult(len(requests), accepted, backpressureEvents)
}

func (self *app__core__FluxPipeline) aggregate(responses []*app__core__FluxProxyResponse) map[string]any {
	byRoute := New_haxe__ds__StringMap()
	routeKeys := []*string{}
	_g := 0
	for _g < len(responses) {
		response := responses[_g]
		_g = int(int32((_g + 1)))
		route := response.route
		bucket := func(hx_value_33 any) *app__core__FluxRouteAggregate {
			if hx_value_33 == nil {
				var hx_zero_34 *app__core__FluxRouteAggregate
				return hx_zero_34
			}
			return hx_value_33.(*app__core__FluxRouteAggregate)
		}(byRoute.get(route))
		if bucket == nil {
			bucket = New_app__core__FluxRouteAggregate(route)
			byRoute.set(route, bucket)
			routeKeys = append(routeKeys, route)
		}
		bucket.record(response)
	}
	routes := []*app__core__FluxRouteAggregate{}
	digest := hxrt.StringFromLiteral("")
	i := 0
	for i < len(routeKeys) {
		route_1 := routeKeys[i]
		item := func(hx_value_35 any) *app__core__FluxRouteAggregate {
			if hx_value_35 == nil {
				var hx_zero_36 *app__core__FluxRouteAggregate
				return hx_zero_36
			}
			return hx_value_35.(*app__core__FluxRouteAggregate)
		}(byRoute.get(route_1))
		if item != nil {
			routes = append(routes, item)
			if !hxrt.StringEqualStringPtr(digest, hxrt.StringFromLiteral("")) {
				digest = hxrt.StringConcatStringPtr(digest, hxrt.StringFromLiteral(","))
			}
			digest = hxrt.StringConcatStringPtr(digest, item.summaryToken())
		}
		i = int(int32((i + 1)))
	}
	hx_obj_37 := map[string]any{}
	hx_obj_37["routes"] = routes
	hx_obj_37["summary"] = digest
	return hx_obj_37
}

func (self *app__core__FluxPipeline) retries(responses []*app__core__FluxProxyResponse) int {
	total := 0
	_g := 0
	for _g < len(responses) {
		response := responses[_g]
		_g = int(int32((_g + 1)))
		total = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(response.attempts) - hxrt.Int32Wrap(1))))))))
	}
	return total
}

func (self *app__core__FluxPipeline) errors(responses []*app__core__FluxProxyResponse) int {
	total := 0
	_g := 0
	for _g < len(responses) {
		response := responses[_g]
		_g = int(int32((_g + 1)))
		if !response.success {
			total = int(int32((total + 1)))
		}
	}
	return total
}

func (self *app__core__FluxPipeline) applyRoutePolicies(requests []*app__core__FluxRequest, perRouteLimit int, breakerFailureThreshold int, timeoutMs int) map[string]any {
	var hx_if_38 int
	if perRouteLimit <= 0 {
		hx_if_38 = 1
	} else {
		hx_if_38 = perRouteLimit
	}
	normalizedLimit := hx_if_38
	var hx_if_39 int
	if breakerFailureThreshold <= 0 {
		hx_if_39 = 1
	} else {
		hx_if_39 = breakerFailureThreshold
	}
	normalizedBreaker := hx_if_39
	routeCounts := New_haxe__ds__StringMap()
	failureStreak := New_haxe__ds__StringMap()
	dispatchable := []*app__core__FluxRequest{}
	synthetic := []*app__core__FluxProxyResponse{}
	rateLimited := 0
	breakerOpen := 0
	_g := 0
	for _g < len(requests) {
		request := requests[_g]
		_g = int(int32((_g + 1)))
		route := app__core__FluxCodec_normalizedRoute(request.route)
		var hx_if_44 int
		if func(hx_value_40 any) bool {
			if hx_value_40 == nil {
				var hx_zero_41 bool
				return hx_zero_41
			}
			return hx_value_40.(bool)
		}(failureStreak.exists(route)) {
			hx_if_44 = func(hx_value_42 any) int {
				if hx_value_42 == nil {
					var hx_zero_43 int
					return hx_zero_43
				}
				return hx_value_42.(int)
			}(failureStreak.get(route))
		} else {
			hx_if_44 = 0
		}
		streak := hx_if_44
		if streak >= normalizedBreaker {
			synthetic = append(synthetic, app__core__FluxCodec_breakerOpen(request))
			breakerOpen = int(int32((breakerOpen + 1)))
			continue
		}
		var hx_if_49 int
		if func(hx_value_45 any) bool {
			if hx_value_45 == nil {
				var hx_zero_46 bool
				return hx_zero_46
			}
			return hx_value_45.(bool)
		}(routeCounts.exists(route)) {
			hx_if_49 = func(hx_value_47 any) int {
				if hx_value_47 == nil {
					var hx_zero_48 int
					return hx_zero_48
				}
				return hx_value_47.(int)
			}(routeCounts.get(route))
		} else {
			hx_if_49 = 0
		}
		routeCount := hx_if_49
		if routeCount >= normalizedLimit {
			synthetic = append(synthetic, app__core__FluxCodec_rateLimited(request))
			rateLimited = int(int32((rateLimited + 1)))
			continue
		}
		routeCounts.set(route, int(int32((hxrt.Int32Wrap(routeCount) + hxrt.Int32Wrap(1)))))
		dispatchable = append(dispatchable, request)
		predictsFailure := ((request.status >= 500) || (request.latencyMs > timeoutMs))
		if predictsFailure {
			failureStreak.set(route, int(int32((hxrt.Int32Wrap(streak) + hxrt.Int32Wrap(1)))))
		} else {
			failureStreak.set(route, 0)
		}
	}
	hx_obj_50 := map[string]any{}
	hx_obj_50["dispatchable"] = dispatchable
	hx_obj_50["synthetic"] = synthetic
	hx_obj_50["rateLimited"] = rateLimited
	hx_obj_50["breakerOpen"] = breakerOpen
	return hx_obj_50
}

func (self *app__core__FluxPipeline) orderedResponses(synthetic []*app__core__FluxProxyResponse, dispatched []*app__core__FluxProxyResponse, acceptedRequests []*app__core__FluxRequest) []*app__core__FluxProxyResponse {
	byId := New_haxe__ds__IntMap()
	_g := 0
	for _g < len(synthetic) {
		response := synthetic[_g]
		_g = int(int32((_g + 1)))
		byId.set(response.requestId, response)
	}
	_g_1 := 0
	for _g_1 < len(dispatched) {
		response_1 := dispatched[_g_1]
		_g_1 = int(int32((_g_1 + 1)))
		byId.set(response_1.requestId, response_1)
	}
	ordered := []*app__core__FluxProxyResponse{}
	_g_2 := 0
	for _g_2 < len(acceptedRequests) {
		request := acceptedRequests[_g_2]
		_g_2 = int(int32((_g_2 + 1)))
		response_2 := func(hx_value_51 any) *app__core__FluxProxyResponse {
			if hx_value_51 == nil {
				var hx_zero_52 *app__core__FluxProxyResponse
				return hx_zero_52
			}
			return hx_value_51.(*app__core__FluxProxyResponse)
		}(byId.get(request.id))
		if response_2 != nil {
			ordered = append(ordered, response_2)
		}
	}
	return ordered
}
