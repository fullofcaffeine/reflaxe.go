package main

import "examples_fluxproxy_metal/hxrt"

type I_app__core__FluxPipeline interface {
	run(requests *hxrt.Array) *app__core__FluxReport
	ingest(requests *hxrt.Array, capacity int) *app__core__FluxIngestResult
	aggregate(responses *hxrt.Array) hxrt.Anonymous
	retries(responses *hxrt.Array) int
	errors(responses *hxrt.Array) int
	applyRoutePolicies(requests *hxrt.Array, perRouteLimit int, breakerFailureThreshold int, timeoutMs int) hxrt.Anonymous
	orderedResponses(synthetic *hxrt.Array, dispatched *hxrt.Array, acceptedRequests *hxrt.Array) *hxrt.Array
	findRouteAggregate(routes *hxrt.Array, route *string) *app__core__FluxRouteAggregate
	getStringIntStateValue(states *hxrt.Array, key *string) int
	setStringIntStateValue(states *hxrt.Array, key *string, value int) *hxrt.Array
	getResponseById(states *hxrt.Array, requestId int) *app__core__FluxProxyResponse
	setResponseById(states *hxrt.Array, requestId int, response *app__core__FluxProxyResponse) *hxrt.Array
}

type app__core__FluxPipeline struct {
	__hx_this I_app__core__FluxPipeline
	runtime   app__runtime__FluxRuntime
}

func (self *app__core__FluxPipeline) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__core__FluxPipeline(runtime app__runtime__FluxRuntime) *app__core__FluxPipeline {
	return hx_init_app__core__FluxPipeline(hx_alloc_app__core__FluxPipeline(nil), runtime)
}

func hx_alloc_app__core__FluxPipeline(owner any) *app__core__FluxPipeline {
	self := &app__core__FluxPipeline{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__core__FluxPipeline)
	}
	return self
}

func hx_init_app__core__FluxPipeline(self *app__core__FluxPipeline, runtime app__runtime__FluxRuntime) *app__core__FluxPipeline {
	self.runtime = runtime
	return self
}

func (self *app__core__FluxPipeline) run(requests *hxrt.Array) *app__core__FluxReport {
	ingest := self.ingest(requests, 3)
	planned := self.applyRoutePolicies(ingest.acceptedRequests, 2, 2, 50)
	dispatched := self.runtime.dispatch(func(hx_value_1 any) *hxrt.Array {
		if hx_value_1 == nil {
			var hx_zero_2 *hxrt.Array
			return hx_zero_2
		}
		return hx_value_1.(*hxrt.Array)
	}(hxrt.AnonymousGet(planned, "dispatchable")), 1)
	responses := self.orderedResponses(func(hx_value_3 any) *hxrt.Array {
		if hx_value_3 == nil {
			var hx_zero_4 *hxrt.Array
			return hx_zero_4
		}
		return hx_value_3.(*hxrt.Array)
	}(hxrt.AnonymousGet(planned, "synthetic")), dispatched, ingest.acceptedRequests)
	aggregates := self.aggregate(responses)
	retryCount := self.retries(responses)
	errorCount := self.errors(responses)
	score := self.runtime.stageScore(responses, retryCount, ingest.backpressureEvents)
	return New_app__core__FluxReport(self.runtime.profileId(), self.runtime.variantId(), self.runtime.capabilityId(), ingest.receivedCount, ingest.acceptedRequests.Len(), ingest.backpressureEvents, responses.Len(), retryCount, func(hx_value_5 any) int {
		if hx_value_5 == nil {
			var hx_zero_6 int
			return hx_zero_6
		}
		return hx_value_5.(int)
	}(hxrt.AnonymousGet(planned, "rateLimited")), func(hx_value_7 any) int {
		if hx_value_7 == nil {
			var hx_zero_8 int
			return hx_zero_8
		}
		return hx_value_7.(int)
	}(hxrt.AnonymousGet(planned, "breakerOpen")), func(hx_value_9 any) *hxrt.Array {
		if hx_value_9 == nil {
			var hx_zero_10 *hxrt.Array
			return hx_zero_10
		}
		return hx_value_9.(*hxrt.Array)
	}(hxrt.AnonymousGet(aggregates, "routes")).Len(), func(hx_value_11 any) *string {
		if hx_value_11 == nil {
			var hx_zero_12 *string
			return hx_zero_12
		}
		return hx_value_11.(*string)
	}(hxrt.AnonymousGet(aggregates, "summary")), errorCount, score)
}

func (self *app__core__FluxPipeline) ingest(requests *hxrt.Array, capacity int) *app__core__FluxIngestResult {
	var hx_if_13 int
	if capacity <= 0 {
		hx_if_13 = 1
	} else {
		hx_if_13 = capacity
	}
	boundedCapacity := hx_if_13
	queue := hxrt.NewArray()
	queueHead := 0
	accepted := hxrt.NewArray()
	backpressureEvents := 0
	_g := 0
	for _g < requests.Len() {
		request := hx_recover_class_0(requests.Get(_g))
		_g = int(int32((_g + 1)))
		if int((hxrt.Int32Wrap(queue.Len()) - hxrt.Int32Wrap(queueHead))) >= boundedCapacity {
			backpressureEvents = int(int32((backpressureEvents + 1)))
			accepted.Push(queue.Get(queueHead))
			queueHead = int(int32((queueHead + 1)))
		}
		queue.Push(request)
	}
	for queueHead < queue.Len() {
		accepted.Push(queue.Get(queueHead))
		queueHead = int(int32((queueHead + 1)))
	}
	return New_app__core__FluxIngestResult(requests.Len(), accepted, backpressureEvents)
}

func (self *app__core__FluxPipeline) aggregate(responses *hxrt.Array) hxrt.Anonymous {
	routes := hxrt.NewArray()
	_g := 0
	for _g < responses.Len() {
		response := hx_recover_class_1(responses.Get(_g))
		_g = int(int32((_g + 1)))
		route := response.route
		bucket := self.findRouteAggregate(routes, route)
		if bucket == nil {
			bucket = New_app__core__FluxRouteAggregate(route)
			routes.Push(bucket)
		}
		bucket.record(response)
	}
	digest := hxrt.StringFromLiteral("")
	_g_1 := 0
	for _g_1 < routes.Len() {
		item := hx_recover_class_2(routes.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		if !hxrt.StringEqualStringPtr(digest, hxrt.StringFromLiteral("")) {
			digest = hxrt.StringConcatStringPtr(digest, hxrt.StringFromLiteral(","))
		}
		digest = hxrt.StringConcatStringPtr(digest, item.summaryToken())
	}
	hx_obj_18 := map[string]any{}
	hx_obj_18["routes"] = routes
	hx_obj_18["summary"] = digest
	return hxrt.AnonymousFromMap(hx_obj_18)
}

func (self *app__core__FluxPipeline) retries(responses *hxrt.Array) int {
	total := 0
	_g := 0
	for _g < responses.Len() {
		response := hx_recover_class_1(responses.Get(_g))
		_g = int(int32((_g + 1)))
		total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(int((hxrt.Int32Wrap(response.attempts) - hxrt.Int32Wrap(1))))))
	}
	return total
}

func (self *app__core__FluxPipeline) errors(responses *hxrt.Array) int {
	total := 0
	_g := 0
	for _g < responses.Len() {
		response := hx_recover_class_1(responses.Get(_g))
		_g = int(int32((_g + 1)))
		if !response.success {
			total = int(int32((total + 1)))
		}
	}
	return total
}

func (self *app__core__FluxPipeline) applyRoutePolicies(requests *hxrt.Array, perRouteLimit int, breakerFailureThreshold int, timeoutMs int) hxrt.Anonymous {
	var hx_if_19 int
	if perRouteLimit <= 0 {
		hx_if_19 = 1
	} else {
		hx_if_19 = perRouteLimit
	}
	normalizedLimit := hx_if_19
	var hx_if_20 int
	if breakerFailureThreshold <= 0 {
		hx_if_20 = 1
	} else {
		hx_if_20 = breakerFailureThreshold
	}
	normalizedBreaker := hx_if_20
	routeCounts := hxrt.NewArray()
	failureStreak := hxrt.NewArray()
	dispatchable := hxrt.NewArray()
	synthetic := hxrt.NewArray()
	rateLimited := 0
	breakerOpen := 0
	_g := 0
	for _g < requests.Len() {
		request := hx_recover_class_0(requests.Get(_g))
		_g = int(int32((_g + 1)))
		route := app__core__FluxCodec_normalizedRoute(request.route)
		streak := self.getStringIntStateValue(failureStreak, route)
		if streak >= normalizedBreaker {
			synthetic.Push(app__core__FluxCodec_breakerOpen(request))
			breakerOpen = int(int32((breakerOpen + 1)))
			continue
		}
		routeCount := self.getStringIntStateValue(routeCounts, route)
		if routeCount >= normalizedLimit {
			synthetic.Push(app__core__FluxCodec_rateLimited(request))
			rateLimited = int(int32((rateLimited + 1)))
			continue
		}
		routeCounts = self.setStringIntStateValue(routeCounts, route, int((hxrt.Int32Wrap(routeCount) + hxrt.Int32Wrap(1))))
		dispatchable.Push(request)
		predictsFailure := ((request.status >= 500) || (request.latencyMs > timeoutMs))
		if predictsFailure {
			failureStreak = self.setStringIntStateValue(failureStreak, route, int((hxrt.Int32Wrap(streak) + hxrt.Int32Wrap(1))))
		} else {
			failureStreak = self.setStringIntStateValue(failureStreak, route, 0)
		}
	}
	hx_obj_24 := map[string]any{}
	hx_obj_24["dispatchable"] = dispatchable
	hx_obj_24["synthetic"] = synthetic
	hx_obj_24["rateLimited"] = rateLimited
	hx_obj_24["breakerOpen"] = breakerOpen
	return hxrt.AnonymousFromMap(hx_obj_24)
}

func (self *app__core__FluxPipeline) orderedResponses(synthetic *hxrt.Array, dispatched *hxrt.Array, acceptedRequests *hxrt.Array) *hxrt.Array {
	byId := hxrt.NewArray()
	_g := 0
	for _g < synthetic.Len() {
		response := hx_recover_class_1(synthetic.Get(_g))
		_g = int(int32((_g + 1)))
		byId = self.setResponseById(byId, response.requestId, response)
	}
	_g_1 := 0
	for _g_1 < dispatched.Len() {
		response_1 := hx_recover_class_1(dispatched.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		byId = self.setResponseById(byId, response_1.requestId, response_1)
	}
	ordered := hxrt.NewArray()
	_g_2 := 0
	for _g_2 < acceptedRequests.Len() {
		request := hx_recover_class_0(acceptedRequests.Get(_g_2))
		_g_2 = int(int32((_g_2 + 1)))
		response_2 := self.getResponseById(byId, request.id)
		if response_2 != nil {
			ordered.Push(response_2)
		}
	}
	return ordered
}

func (self *app__core__FluxPipeline) findRouteAggregate(routes *hxrt.Array, route *string) *app__core__FluxRouteAggregate {
	_g := 0
	for _g < routes.Len() {
		item := hx_recover_class_2(routes.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(item.route, route) {
			return item
		}
	}
	return nil
}

func (self *app__core__FluxPipeline) getStringIntStateValue(states *hxrt.Array, key *string) int {
	_g := 0
	for _g < states.Len() {
		state := hx_structural_recover(states.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_value_28 any) *string {
			if hx_value_28 == nil {
				var hx_zero_29 *string
				return hx_zero_29
			}
			return hx_value_28.(*string)
		}(hxrt.AnonymousGet(state, "key")), key) {
			return func(hx_value_26 any) int {
				if hx_value_26 == nil {
					var hx_zero_27 int
					return hx_zero_27
				}
				return hx_value_26.(int)
			}(hxrt.AnonymousGet(state, "value"))
		}
	}
	return 0
}

func (self *app__core__FluxPipeline) setStringIntStateValue(states *hxrt.Array, key *string, value int) *hxrt.Array {
	_g := 0
	for _g < states.Len() {
		state := hx_structural_recover(states.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_value_33 any) *string {
			if hx_value_33 == nil {
				var hx_zero_34 *string
				return hx_zero_34
			}
			return hx_value_33.(*string)
		}(hxrt.AnonymousGet(state, "key")), key) {
			func() int {
				hx_structural_target_30 := state
				hx_structural_value_32 := value
				hxrt.AnonymousSet(hx_structural_target_30, "value", hx_structural_value_32)
				return hx_structural_value_32
			}()
			return states
		}
	}
	hx_obj_36 := map[string]any{}
	hx_obj_36["key"] = key
	hx_obj_36["value"] = value
	states.Push(hxrt.AnonymousFromMap(hx_obj_36))
	return states
}

func (self *app__core__FluxPipeline) getResponseById(states *hxrt.Array, requestId int) *app__core__FluxProxyResponse {
	_g := 0
	for _g < states.Len() {
		state := hx_structural_recover(states.Get(_g))
		_g = int(int32((_g + 1)))
		if func(hx_value_37 any) int {
			if hx_value_37 == nil {
				var hx_zero_38 int
				return hx_zero_38
			}
			return hx_value_37.(int)
		}(hxrt.AnonymousGet(state, "requestId")) == requestId {
			return hx_recover_class_1(hxrt.AnonymousGet(state, "response"))
		}
	}
	return nil
}

func (self *app__core__FluxPipeline) setResponseById(states *hxrt.Array, requestId int, response *app__core__FluxProxyResponse) *hxrt.Array {
	_g := 0
	for _g < states.Len() {
		state := hx_structural_recover(states.Get(_g))
		_g = int(int32((_g + 1)))
		if func(hx_value_42 any) int {
			if hx_value_42 == nil {
				var hx_zero_43 int
				return hx_zero_43
			}
			return hx_value_42.(int)
		}(hxrt.AnonymousGet(state, "requestId")) == requestId {
			func() *app__core__FluxProxyResponse {
				hx_structural_target_39 := state
				hx_structural_value_41 := response
				hxrt.AnonymousSet(hx_structural_target_39, "response", hx_structural_value_41)
				return hx_structural_value_41
			}()
			return states
		}
	}
	hx_obj_45 := map[string]any{}
	hx_obj_45["requestId"] = requestId
	hx_obj_45["response"] = response
	states.Push(hxrt.AnonymousFromMap(hx_obj_45))
	return states
}
