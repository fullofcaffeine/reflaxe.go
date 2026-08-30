package main

import "examples_fluxproxy_portable/hxrt"

type I_app__core__FluxPipeline interface {
	run(requests *hxrt.Array) *app__core__FluxReport
	ingest(requests *hxrt.Array, capacity int) *app__core__FluxIngestResult
	aggregate(responses *hxrt.Array) map[string]any
	retries(responses *hxrt.Array) int
	errors(responses *hxrt.Array) int
	applyRoutePolicies(requests *hxrt.Array, perRouteLimit int, breakerFailureThreshold int, timeoutMs int) map[string]any
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

func New_app__core__FluxPipeline(runtime app__runtime__FluxRuntime) *app__core__FluxPipeline {
	self := &app__core__FluxPipeline{}
	self.__hx_this = self
	self.runtime = runtime
	return self
}

func (self *app__core__FluxPipeline) run(requests *hxrt.Array) *app__core__FluxReport {
	ingest := self.ingest(requests, 3)
	planned := self.applyRoutePolicies(ingest.acceptedRequests, 2, 2, 50)
	dispatched := self.runtime.dispatch(func(hx_obj_1 map[string]any) *hxrt.Array {
		hx_field_2 := hx_obj_1["dispatchable"]
		if hx_field_2 == nil {
			var hx_zero_3 *hxrt.Array
			return hx_zero_3
		}
		return hx_field_2.(*hxrt.Array)
	}(planned), 1)
	responses := self.orderedResponses(func(hx_obj_4 map[string]any) *hxrt.Array {
		hx_field_5 := hx_obj_4["synthetic"]
		if hx_field_5 == nil {
			var hx_zero_6 *hxrt.Array
			return hx_zero_6
		}
		return hx_field_5.(*hxrt.Array)
	}(planned), dispatched, ingest.acceptedRequests)
	aggregates := self.aggregate(responses)
	retryCount := self.retries(responses)
	errorCount := self.errors(responses)
	score := self.runtime.stageScore(responses, retryCount, ingest.backpressureEvents)
	return New_app__core__FluxReport(self.runtime.profileId(), self.runtime.variantId(), self.runtime.capabilityId(), ingest.receivedCount, ingest.acceptedRequests.Len(), ingest.backpressureEvents, responses.Len(), retryCount, func(hx_obj_7 map[string]any) int {
		hx_field_8 := hx_obj_7["rateLimited"]
		if hx_field_8 == nil {
			var hx_zero_9 int
			return hx_zero_9
		}
		return hx_field_8.(int)
	}(planned), func(hx_obj_10 map[string]any) int {
		hx_field_11 := hx_obj_10["breakerOpen"]
		if hx_field_11 == nil {
			var hx_zero_12 int
			return hx_zero_12
		}
		return hx_field_11.(int)
	}(planned), func(hx_obj_13 map[string]any) *hxrt.Array {
		hx_field_14 := hx_obj_13["routes"]
		if hx_field_14 == nil {
			var hx_zero_15 *hxrt.Array
			return hx_zero_15
		}
		return hx_field_14.(*hxrt.Array)
	}(aggregates).Len(), func(hx_obj_16 map[string]any) *string {
		hx_field_17 := hx_obj_16["summary"]
		if hx_field_17 == nil {
			var hx_zero_18 *string
			return hx_zero_18
		}
		return hx_field_17.(*string)
	}(aggregates), errorCount, score)
}

func (self *app__core__FluxPipeline) ingest(requests *hxrt.Array, capacity int) *app__core__FluxIngestResult {
	var hx_if_19 int
	if capacity <= 0 {
		hx_if_19 = 1
	} else {
		hx_if_19 = capacity
	}
	boundedCapacity := hx_if_19
	queue := hxrt.NewArray()
	queueHead := 0
	accepted := hxrt.NewArray()
	backpressureEvents := 0
	_g := 0
	for _g < requests.Len() {
		request := func(hx_value_20 any) *app__core__FluxRequest {
			if hx_value_20 == nil {
				var hx_zero_21 *app__core__FluxRequest
				return hx_zero_21
			}
			return hx_value_20.(*app__core__FluxRequest)
		}(requests.Get(_g))
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

func (self *app__core__FluxPipeline) aggregate(responses *hxrt.Array) map[string]any {
	routes := hxrt.NewArray()
	_g := 0
	for _g < responses.Len() {
		response := func(hx_value_25 any) *app__core__FluxProxyResponse {
			if hx_value_25 == nil {
				var hx_zero_26 *app__core__FluxProxyResponse
				return hx_zero_26
			}
			return hx_value_25.(*app__core__FluxProxyResponse)
		}(responses.Get(_g))
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
		item := func(hx_value_28 any) *app__core__FluxRouteAggregate {
			if hx_value_28 == nil {
				var hx_zero_29 *app__core__FluxRouteAggregate
				return hx_zero_29
			}
			return hx_value_28.(*app__core__FluxRouteAggregate)
		}(routes.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		if !hxrt.StringEqualStringPtr(digest, hxrt.StringFromLiteral("")) {
			digest = hxrt.StringConcatStringPtr(digest, hxrt.StringFromLiteral(","))
		}
		digest = hxrt.StringConcatStringPtr(digest, item.summaryToken())
	}
	hx_obj_30 := map[string]any{}
	hx_obj_30["routes"] = routes
	hx_obj_30["summary"] = digest
	return hx_obj_30
}

func (self *app__core__FluxPipeline) retries(responses *hxrt.Array) int {
	total := 0
	_g := 0
	for _g < responses.Len() {
		response := func(hx_value_31 any) *app__core__FluxProxyResponse {
			if hx_value_31 == nil {
				var hx_zero_32 *app__core__FluxProxyResponse
				return hx_zero_32
			}
			return hx_value_31.(*app__core__FluxProxyResponse)
		}(responses.Get(_g))
		_g = int(int32((_g + 1)))
		total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(int((hxrt.Int32Wrap(response.attempts) - hxrt.Int32Wrap(1))))))
	}
	return total
}

func (self *app__core__FluxPipeline) errors(responses *hxrt.Array) int {
	total := 0
	_g := 0
	for _g < responses.Len() {
		response := func(hx_value_33 any) *app__core__FluxProxyResponse {
			if hx_value_33 == nil {
				var hx_zero_34 *app__core__FluxProxyResponse
				return hx_zero_34
			}
			return hx_value_33.(*app__core__FluxProxyResponse)
		}(responses.Get(_g))
		_g = int(int32((_g + 1)))
		if !response.success {
			total = int(int32((total + 1)))
		}
	}
	return total
}

func (self *app__core__FluxPipeline) applyRoutePolicies(requests *hxrt.Array, perRouteLimit int, breakerFailureThreshold int, timeoutMs int) map[string]any {
	var hx_if_35 int
	if perRouteLimit <= 0 {
		hx_if_35 = 1
	} else {
		hx_if_35 = perRouteLimit
	}
	normalizedLimit := hx_if_35
	var hx_if_36 int
	if breakerFailureThreshold <= 0 {
		hx_if_36 = 1
	} else {
		hx_if_36 = breakerFailureThreshold
	}
	normalizedBreaker := hx_if_36
	routeCounts := hxrt.NewArray()
	failureStreak := hxrt.NewArray()
	dispatchable := hxrt.NewArray()
	synthetic := hxrt.NewArray()
	rateLimited := 0
	breakerOpen := 0
	_g := 0
	for _g < requests.Len() {
		request := func(hx_value_37 any) *app__core__FluxRequest {
			if hx_value_37 == nil {
				var hx_zero_38 *app__core__FluxRequest
				return hx_zero_38
			}
			return hx_value_37.(*app__core__FluxRequest)
		}(requests.Get(_g))
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
	hx_obj_42 := map[string]any{}
	hx_obj_42["dispatchable"] = dispatchable
	hx_obj_42["synthetic"] = synthetic
	hx_obj_42["rateLimited"] = rateLimited
	hx_obj_42["breakerOpen"] = breakerOpen
	return hx_obj_42
}

func (self *app__core__FluxPipeline) orderedResponses(synthetic *hxrt.Array, dispatched *hxrt.Array, acceptedRequests *hxrt.Array) *hxrt.Array {
	byId := hxrt.NewArray()
	_g := 0
	for _g < synthetic.Len() {
		response := func(hx_value_43 any) *app__core__FluxProxyResponse {
			if hx_value_43 == nil {
				var hx_zero_44 *app__core__FluxProxyResponse
				return hx_zero_44
			}
			return hx_value_43.(*app__core__FluxProxyResponse)
		}(synthetic.Get(_g))
		_g = int(int32((_g + 1)))
		byId = self.setResponseById(byId, response.requestId, response)
	}
	_g_1 := 0
	for _g_1 < dispatched.Len() {
		response_1 := func(hx_value_45 any) *app__core__FluxProxyResponse {
			if hx_value_45 == nil {
				var hx_zero_46 *app__core__FluxProxyResponse
				return hx_zero_46
			}
			return hx_value_45.(*app__core__FluxProxyResponse)
		}(dispatched.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		byId = self.setResponseById(byId, response_1.requestId, response_1)
	}
	ordered := hxrt.NewArray()
	_g_2 := 0
	for _g_2 < acceptedRequests.Len() {
		request := func(hx_value_47 any) *app__core__FluxRequest {
			if hx_value_47 == nil {
				var hx_zero_48 *app__core__FluxRequest
				return hx_zero_48
			}
			return hx_value_47.(*app__core__FluxRequest)
		}(acceptedRequests.Get(_g_2))
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
		item := func(hx_value_50 any) *app__core__FluxRouteAggregate {
			if hx_value_50 == nil {
				var hx_zero_51 *app__core__FluxRouteAggregate
				return hx_zero_51
			}
			return hx_value_50.(*app__core__FluxRouteAggregate)
		}(routes.Get(_g))
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
		state := func(hx_value_52 any) map[string]any {
			if hx_value_52 == nil {
				var hx_zero_53 map[string]any
				return hx_zero_53
			}
			return hx_value_52.(map[string]any)
		}(states.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_obj_57 map[string]any) *string {
			hx_field_58 := hx_obj_57["key"]
			if hx_field_58 == nil {
				var hx_zero_59 *string
				return hx_zero_59
			}
			return hx_field_58.(*string)
		}(state), key) {
			return func(hx_obj_54 map[string]any) int {
				hx_field_55 := hx_obj_54["value"]
				if hx_field_55 == nil {
					var hx_zero_56 int
					return hx_zero_56
				}
				return hx_field_55.(int)
			}(state)
		}
	}
	return 0
}

func (self *app__core__FluxPipeline) setStringIntStateValue(states *hxrt.Array, key *string, value int) *hxrt.Array {
	_g := 0
	for _g < states.Len() {
		state := func(hx_value_60 any) map[string]any {
			if hx_value_60 == nil {
				var hx_zero_61 map[string]any
				return hx_zero_61
			}
			return hx_value_60.(map[string]any)
		}(states.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_obj_62 map[string]any) *string {
			hx_field_63 := hx_obj_62["key"]
			if hx_field_63 == nil {
				var hx_zero_64 *string
				return hx_zero_64
			}
			return hx_field_63.(*string)
		}(state), key) {
			state["value"] = value
			return states
		}
	}
	hx_obj_66 := map[string]any{}
	hx_obj_66["key"] = key
	hx_obj_66["value"] = value
	states.Push(hx_obj_66)
	return states
}

func (self *app__core__FluxPipeline) getResponseById(states *hxrt.Array, requestId int) *app__core__FluxProxyResponse {
	_g := 0
	for _g < states.Len() {
		state := func(hx_value_67 any) map[string]any {
			if hx_value_67 == nil {
				var hx_zero_68 map[string]any
				return hx_zero_68
			}
			return hx_value_67.(map[string]any)
		}(states.Get(_g))
		_g = int(int32((_g + 1)))
		if func(hx_obj_72 map[string]any) int {
			hx_field_73 := hx_obj_72["requestId"]
			if hx_field_73 == nil {
				var hx_zero_74 int
				return hx_zero_74
			}
			return hx_field_73.(int)
		}(state) == requestId {
			return func(hx_obj_69 map[string]any) *app__core__FluxProxyResponse {
				hx_field_70 := hx_obj_69["response"]
				if hx_field_70 == nil {
					var hx_zero_71 *app__core__FluxProxyResponse
					return hx_zero_71
				}
				return hx_field_70.(*app__core__FluxProxyResponse)
			}(state)
		}
	}
	return nil
}

func (self *app__core__FluxPipeline) setResponseById(states *hxrt.Array, requestId int, response *app__core__FluxProxyResponse) *hxrt.Array {
	_g := 0
	for _g < states.Len() {
		state := func(hx_value_75 any) map[string]any {
			if hx_value_75 == nil {
				var hx_zero_76 map[string]any
				return hx_zero_76
			}
			return hx_value_75.(map[string]any)
		}(states.Get(_g))
		_g = int(int32((_g + 1)))
		if func(hx_obj_77 map[string]any) int {
			hx_field_78 := hx_obj_77["requestId"]
			if hx_field_78 == nil {
				var hx_zero_79 int
				return hx_zero_79
			}
			return hx_field_78.(int)
		}(state) == requestId {
			state["response"] = response
			return states
		}
	}
	hx_obj_81 := map[string]any{}
	hx_obj_81["requestId"] = requestId
	hx_obj_81["response"] = response
	states.Push(hx_obj_81)
	return states
}
