package main

import "examples_fluxproxy_portable/hxrt"

type I_app__core__FluxPipeline interface {
	run(requests []*app__core__FluxRequest) *app__core__FluxReport
	ingest(requests []*app__core__FluxRequest, capacity int) *app__core__FluxIngestResult
	aggregate(responses []*app__core__FluxProxyResponse) map[string]any
	retries(responses []*app__core__FluxProxyResponse) int
	errors(responses []*app__core__FluxProxyResponse) int
	applyRoutePolicies(requests []*app__core__FluxRequest, perRouteLimit int, breakerFailureThreshold int, timeoutMs int) map[string]any
	orderedResponses(synthetic []*app__core__FluxProxyResponse, dispatched []*app__core__FluxProxyResponse, acceptedRequests []*app__core__FluxRequest) []*app__core__FluxProxyResponse
	findRouteAggregate(routes []*app__core__FluxRouteAggregate, route *string) *app__core__FluxRouteAggregate
	getStringIntStateValue(states []map[string]any, key *string) int
	setStringIntStateValue(states []map[string]any, key *string, value int) []map[string]any
	getResponseById(states []map[string]any, requestId int) *app__core__FluxProxyResponse
	setResponseById(states []map[string]any, requestId int, response *app__core__FluxProxyResponse) []map[string]any
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
	dispatched := self.runtime.dispatch(func(hx_obj_5 map[string]any) []*app__core__FluxRequest {
		hx_field_6 := hx_obj_5["dispatchable"]
		if hx_field_6 == nil {
			var hx_zero_7 []*app__core__FluxRequest
			return hx_zero_7
		}
		return hx_field_6.([]*app__core__FluxRequest)
	}(planned), 1)
	responses := self.orderedResponses(func(hx_obj_8 map[string]any) []*app__core__FluxProxyResponse {
		hx_field_9 := hx_obj_8["synthetic"]
		if hx_field_9 == nil {
			var hx_zero_10 []*app__core__FluxProxyResponse
			return hx_zero_10
		}
		return hx_field_9.([]*app__core__FluxProxyResponse)
	}(planned), dispatched, ingest.acceptedRequests)
	aggregates := self.aggregate(responses)
	retryCount := self.retries(responses)
	errorCount := self.errors(responses)
	score := self.runtime.stageScore(responses, retryCount, ingest.backpressureEvents)
	return New_app__core__FluxReport(self.runtime.profileId(), self.runtime.variantId(), self.runtime.capabilityId(), ingest.receivedCount, len(ingest.acceptedRequests), ingest.backpressureEvents, len(responses), retryCount, func(hx_obj_11 map[string]any) int {
		hx_field_12 := hx_obj_11["rateLimited"]
		if hx_field_12 == nil {
			var hx_zero_13 int
			return hx_zero_13
		}
		return hx_field_12.(int)
	}(planned), func(hx_obj_14 map[string]any) int {
		hx_field_15 := hx_obj_14["breakerOpen"]
		if hx_field_15 == nil {
			var hx_zero_16 int
			return hx_zero_16
		}
		return hx_field_15.(int)
	}(planned), len(func(hx_obj_17 map[string]any) []*app__core__FluxRouteAggregate {
		hx_field_18 := hx_obj_17["routes"]
		if hx_field_18 == nil {
			var hx_zero_19 []*app__core__FluxRouteAggregate
			return hx_zero_19
		}
		return hx_field_18.([]*app__core__FluxRouteAggregate)
	}(aggregates)), func(hx_obj_20 map[string]any) *string {
		hx_field_21 := hx_obj_20["summary"]
		if hx_field_21 == nil {
			var hx_zero_22 *string
			return hx_zero_22
		}
		return hx_field_21.(*string)
	}(aggregates), errorCount, score)
}

func (self *app__core__FluxPipeline) ingest(requests []*app__core__FluxRequest, capacity int) *app__core__FluxIngestResult {
	var hx_if_23 int
	if capacity <= 0 {
		hx_if_23 = 1
	} else {
		hx_if_23 = capacity
	}
	boundedCapacity := hx_if_23
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
	routes := []*app__core__FluxRouteAggregate{}
	_g := 0
	for _g < len(responses) {
		response := responses[_g]
		_g = int(int32((_g + 1)))
		route := response.route
		bucket := self.findRouteAggregate(routes, route)
		if bucket == nil {
			bucket = New_app__core__FluxRouteAggregate(route)
			routes = append(routes, bucket)
		}
		bucket.record(response)
	}
	digest := hxrt.StringFromLiteral("")
	_g_1 := 0
	for _g_1 < len(routes) {
		item := routes[_g_1]
		_g_1 = int(int32((_g_1 + 1)))
		if !hxrt.StringEqualStringPtr(digest, hxrt.StringFromLiteral("")) {
			digest = hxrt.StringConcatStringPtr(digest, hxrt.StringFromLiteral(","))
		}
		digest = hxrt.StringConcatStringPtr(digest, item.summaryToken())
	}
	hx_obj_28 := map[string]any{}
	hx_obj_28["routes"] = routes
	hx_obj_28["summary"] = digest
	return hx_obj_28
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
	var hx_if_29 int
	if perRouteLimit <= 0 {
		hx_if_29 = 1
	} else {
		hx_if_29 = perRouteLimit
	}
	normalizedLimit := hx_if_29
	var hx_if_30 int
	if breakerFailureThreshold <= 0 {
		hx_if_30 = 1
	} else {
		hx_if_30 = breakerFailureThreshold
	}
	normalizedBreaker := hx_if_30
	routeCounts := []map[string]any{}
	failureStreak := []map[string]any{}
	dispatchable := []*app__core__FluxRequest{}
	synthetic := []*app__core__FluxProxyResponse{}
	rateLimited := 0
	breakerOpen := 0
	_g := 0
	for _g < len(requests) {
		request := requests[_g]
		_g = int(int32((_g + 1)))
		route := app__core__FluxCodec_normalizedRoute(request.route)
		streak := self.getStringIntStateValue(failureStreak, route)
		if streak >= normalizedBreaker {
			synthetic = append(synthetic, app__core__FluxCodec_breakerOpen(request))
			breakerOpen = int(int32((breakerOpen + 1)))
			continue
		}
		routeCount := self.getStringIntStateValue(routeCounts, route)
		if routeCount >= normalizedLimit {
			synthetic = append(synthetic, app__core__FluxCodec_rateLimited(request))
			rateLimited = int(int32((rateLimited + 1)))
			continue
		}
		routeCounts = self.setStringIntStateValue(routeCounts, route, int(int32((hxrt.Int32Wrap(routeCount) + hxrt.Int32Wrap(1)))))
		dispatchable = append(dispatchable, request)
		predictsFailure := ((request.status >= 500) || (request.latencyMs > timeoutMs))
		if predictsFailure {
			failureStreak = self.setStringIntStateValue(failureStreak, route, int(int32((hxrt.Int32Wrap(streak) + hxrt.Int32Wrap(1)))))
		} else {
			failureStreak = self.setStringIntStateValue(failureStreak, route, 0)
		}
	}
	hx_obj_34 := map[string]any{}
	hx_obj_34["dispatchable"] = dispatchable
	hx_obj_34["synthetic"] = synthetic
	hx_obj_34["rateLimited"] = rateLimited
	hx_obj_34["breakerOpen"] = breakerOpen
	return hx_obj_34
}

func (self *app__core__FluxPipeline) orderedResponses(synthetic []*app__core__FluxProxyResponse, dispatched []*app__core__FluxProxyResponse, acceptedRequests []*app__core__FluxRequest) []*app__core__FluxProxyResponse {
	byId := []map[string]any{}
	_g := 0
	for _g < len(synthetic) {
		response := synthetic[_g]
		_g = int(int32((_g + 1)))
		byId = self.setResponseById(byId, response.requestId, response)
	}
	_g_1 := 0
	for _g_1 < len(dispatched) {
		response_1 := dispatched[_g_1]
		_g_1 = int(int32((_g_1 + 1)))
		byId = self.setResponseById(byId, response_1.requestId, response_1)
	}
	ordered := []*app__core__FluxProxyResponse{}
	_g_2 := 0
	for _g_2 < len(acceptedRequests) {
		request := acceptedRequests[_g_2]
		_g_2 = int(int32((_g_2 + 1)))
		response_2 := self.getResponseById(byId, request.id)
		if response_2 != nil {
			ordered = append(ordered, response_2)
		}
	}
	return ordered
}

func (self *app__core__FluxPipeline) findRouteAggregate(routes []*app__core__FluxRouteAggregate, route *string) *app__core__FluxRouteAggregate {
	_g := 0
	for _g < len(routes) {
		item := routes[_g]
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(item.route, route) {
			return item
		}
	}
	return nil
}

func (self *app__core__FluxPipeline) getStringIntStateValue(states []map[string]any, key *string) int {
	_g := 0
	for _g < len(states) {
		state := states[_g]
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_obj_39 map[string]any) *string {
			hx_field_40 := hx_obj_39["key"]
			if hx_field_40 == nil {
				var hx_zero_41 *string
				return hx_zero_41
			}
			return hx_field_40.(*string)
		}(state), key) {
			return func(hx_obj_36 map[string]any) int {
				hx_field_37 := hx_obj_36["value"]
				if hx_field_37 == nil {
					var hx_zero_38 int
					return hx_zero_38
				}
				return hx_field_37.(int)
			}(state)
		}
	}
	return 0
}

func (self *app__core__FluxPipeline) setStringIntStateValue(states []map[string]any, key *string, value int) []map[string]any {
	_g := 0
	for _g < len(states) {
		state := states[_g]
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_obj_42 map[string]any) *string {
			hx_field_43 := hx_obj_42["key"]
			if hx_field_43 == nil {
				var hx_zero_44 *string
				return hx_zero_44
			}
			return hx_field_43.(*string)
		}(state), key) {
			state["value"] = value
			return states
		}
	}
	states = append(states, func() map[string]any {
		hx_obj_46 := map[string]any{}
		hx_obj_46["key"] = key
		hx_obj_46["value"] = value
		return hx_obj_46
	}())
	return states
}

func (self *app__core__FluxPipeline) getResponseById(states []map[string]any, requestId int) *app__core__FluxProxyResponse {
	_g := 0
	for _g < len(states) {
		state := states[_g]
		_g = int(int32((_g + 1)))
		if func(hx_obj_50 map[string]any) int {
			hx_field_51 := hx_obj_50["requestId"]
			if hx_field_51 == nil {
				var hx_zero_52 int
				return hx_zero_52
			}
			return hx_field_51.(int)
		}(state) == requestId {
			return func(hx_obj_47 map[string]any) *app__core__FluxProxyResponse {
				hx_field_48 := hx_obj_47["response"]
				if hx_field_48 == nil {
					var hx_zero_49 *app__core__FluxProxyResponse
					return hx_zero_49
				}
				return hx_field_48.(*app__core__FluxProxyResponse)
			}(state)
		}
	}
	return nil
}

func (self *app__core__FluxPipeline) setResponseById(states []map[string]any, requestId int, response *app__core__FluxProxyResponse) []map[string]any {
	_g := 0
	for _g < len(states) {
		state := states[_g]
		_g = int(int32((_g + 1)))
		if func(hx_obj_53 map[string]any) int {
			hx_field_54 := hx_obj_53["requestId"]
			if hx_field_54 == nil {
				var hx_zero_55 int
				return hx_zero_55
			}
			return hx_field_54.(int)
		}(state) == requestId {
			state["response"] = response
			return states
		}
	}
	states = append(states, func() map[string]any {
		hx_obj_57 := map[string]any{}
		hx_obj_57["requestId"] = requestId
		hx_obj_57["response"] = response
		return hx_obj_57
	}())
	return states
}
