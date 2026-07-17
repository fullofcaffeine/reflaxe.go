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
	dispatched := self.runtime.dispatch(func(hx_obj_29 map[string]any) *hxrt.Array {
		hx_field_30 := hx_obj_29["dispatchable"]
		if hx_field_30 == nil {
			var hx_zero_31 *hxrt.Array
			return hx_zero_31
		}
		return hx_field_30.(*hxrt.Array)
	}(planned), 1)
	responses := self.orderedResponses(func(hx_obj_32 map[string]any) *hxrt.Array {
		hx_field_33 := hx_obj_32["synthetic"]
		if hx_field_33 == nil {
			var hx_zero_34 *hxrt.Array
			return hx_zero_34
		}
		return hx_field_33.(*hxrt.Array)
	}(planned), dispatched, ingest.acceptedRequests)
	aggregates := self.aggregate(responses)
	retryCount := self.retries(responses)
	errorCount := self.errors(responses)
	score := self.runtime.stageScore(responses, retryCount, ingest.backpressureEvents)
	return New_app__core__FluxReport(self.runtime.profileId(), self.runtime.variantId(), self.runtime.capabilityId(), ingest.receivedCount, ingest.acceptedRequests.Len(), ingest.backpressureEvents, responses.Len(), retryCount, func(hx_obj_35 map[string]any) int {
		hx_field_36 := hx_obj_35["rateLimited"]
		if hx_field_36 == nil {
			var hx_zero_37 int
			return hx_zero_37
		}
		return hx_field_36.(int)
	}(planned), func(hx_obj_38 map[string]any) int {
		hx_field_39 := hx_obj_38["breakerOpen"]
		if hx_field_39 == nil {
			var hx_zero_40 int
			return hx_zero_40
		}
		return hx_field_39.(int)
	}(planned), func(hx_obj_41 map[string]any) *hxrt.Array {
		hx_field_42 := hx_obj_41["routes"]
		if hx_field_42 == nil {
			var hx_zero_43 *hxrt.Array
			return hx_zero_43
		}
		return hx_field_42.(*hxrt.Array)
	}(aggregates).Len(), func(hx_obj_44 map[string]any) *string {
		hx_field_45 := hx_obj_44["summary"]
		if hx_field_45 == nil {
			var hx_zero_46 *string
			return hx_zero_46
		}
		return hx_field_45.(*string)
	}(aggregates), errorCount, score)
}

func (self *app__core__FluxPipeline) ingest(requests *hxrt.Array, capacity int) *app__core__FluxIngestResult {
	var hx_if_47 int
	if capacity <= 0 {
		hx_if_47 = 1
	} else {
		hx_if_47 = capacity
	}
	boundedCapacity := hx_if_47
	queue := hxrt.NewArray()
	queueHead := 0
	accepted := hxrt.NewArray()
	backpressureEvents := 0
	_g := 0
	for _g < requests.Len() {
		request := func(hx_value_48 any) *app__core__FluxRequest {
			if hx_value_48 == nil {
				var hx_zero_49 *app__core__FluxRequest
				return hx_zero_49
			}
			return hx_value_48.(*app__core__FluxRequest)
		}(requests.Get(_g))
		_g = int(int32((_g + 1)))
		if int(int32((hxrt.Int32Wrap(queue.Len()) - hxrt.Int32Wrap(queueHead)))) >= boundedCapacity {
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
		response := func(hx_value_53 any) *app__core__FluxProxyResponse {
			if hx_value_53 == nil {
				var hx_zero_54 *app__core__FluxProxyResponse
				return hx_zero_54
			}
			return hx_value_53.(*app__core__FluxProxyResponse)
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
		item := func(hx_value_56 any) *app__core__FluxRouteAggregate {
			if hx_value_56 == nil {
				var hx_zero_57 *app__core__FluxRouteAggregate
				return hx_zero_57
			}
			return hx_value_56.(*app__core__FluxRouteAggregate)
		}(routes.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		if !hxrt.StringEqualStringPtr(digest, hxrt.StringFromLiteral("")) {
			digest = hxrt.StringConcatStringPtr(digest, hxrt.StringFromLiteral(","))
		}
		digest = hxrt.StringConcatStringPtr(digest, item.summaryToken())
	}
	hx_obj_58 := map[string]any{}
	hx_obj_58["routes"] = routes
	hx_obj_58["summary"] = digest
	return hx_obj_58
}

func (self *app__core__FluxPipeline) retries(responses *hxrt.Array) int {
	total := 0
	_g := 0
	for _g < responses.Len() {
		response := func(hx_value_59 any) *app__core__FluxProxyResponse {
			if hx_value_59 == nil {
				var hx_zero_60 *app__core__FluxProxyResponse
				return hx_zero_60
			}
			return hx_value_59.(*app__core__FluxProxyResponse)
		}(responses.Get(_g))
		_g = int(int32((_g + 1)))
		total = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(response.attempts) - hxrt.Int32Wrap(1))))))))
	}
	return total
}

func (self *app__core__FluxPipeline) errors(responses *hxrt.Array) int {
	total := 0
	_g := 0
	for _g < responses.Len() {
		response := func(hx_value_61 any) *app__core__FluxProxyResponse {
			if hx_value_61 == nil {
				var hx_zero_62 *app__core__FluxProxyResponse
				return hx_zero_62
			}
			return hx_value_61.(*app__core__FluxProxyResponse)
		}(responses.Get(_g))
		_g = int(int32((_g + 1)))
		if !response.success {
			total = int(int32((total + 1)))
		}
	}
	return total
}

func (self *app__core__FluxPipeline) applyRoutePolicies(requests *hxrt.Array, perRouteLimit int, breakerFailureThreshold int, timeoutMs int) map[string]any {
	var hx_if_63 int
	if perRouteLimit <= 0 {
		hx_if_63 = 1
	} else {
		hx_if_63 = perRouteLimit
	}
	normalizedLimit := hx_if_63
	var hx_if_64 int
	if breakerFailureThreshold <= 0 {
		hx_if_64 = 1
	} else {
		hx_if_64 = breakerFailureThreshold
	}
	normalizedBreaker := hx_if_64
	routeCounts := hxrt.NewArray()
	failureStreak := hxrt.NewArray()
	dispatchable := hxrt.NewArray()
	synthetic := hxrt.NewArray()
	rateLimited := 0
	breakerOpen := 0
	_g := 0
	for _g < requests.Len() {
		request := func(hx_value_65 any) *app__core__FluxRequest {
			if hx_value_65 == nil {
				var hx_zero_66 *app__core__FluxRequest
				return hx_zero_66
			}
			return hx_value_65.(*app__core__FluxRequest)
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
		routeCounts = self.setStringIntStateValue(routeCounts, route, int(int32((hxrt.Int32Wrap(routeCount) + hxrt.Int32Wrap(1)))))
		dispatchable.Push(request)
		predictsFailure := ((request.status >= 500) || (request.latencyMs > timeoutMs))
		if predictsFailure {
			failureStreak = self.setStringIntStateValue(failureStreak, route, int(int32((hxrt.Int32Wrap(streak) + hxrt.Int32Wrap(1)))))
		} else {
			failureStreak = self.setStringIntStateValue(failureStreak, route, 0)
		}
	}
	hx_obj_70 := map[string]any{}
	hx_obj_70["dispatchable"] = dispatchable
	hx_obj_70["synthetic"] = synthetic
	hx_obj_70["rateLimited"] = rateLimited
	hx_obj_70["breakerOpen"] = breakerOpen
	return hx_obj_70
}

func (self *app__core__FluxPipeline) orderedResponses(synthetic *hxrt.Array, dispatched *hxrt.Array, acceptedRequests *hxrt.Array) *hxrt.Array {
	byId := hxrt.NewArray()
	_g := 0
	for _g < synthetic.Len() {
		response := func(hx_value_71 any) *app__core__FluxProxyResponse {
			if hx_value_71 == nil {
				var hx_zero_72 *app__core__FluxProxyResponse
				return hx_zero_72
			}
			return hx_value_71.(*app__core__FluxProxyResponse)
		}(synthetic.Get(_g))
		_g = int(int32((_g + 1)))
		byId = self.setResponseById(byId, response.requestId, response)
	}
	_g_1 := 0
	for _g_1 < dispatched.Len() {
		response_1 := func(hx_value_73 any) *app__core__FluxProxyResponse {
			if hx_value_73 == nil {
				var hx_zero_74 *app__core__FluxProxyResponse
				return hx_zero_74
			}
			return hx_value_73.(*app__core__FluxProxyResponse)
		}(dispatched.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		byId = self.setResponseById(byId, response_1.requestId, response_1)
	}
	ordered := hxrt.NewArray()
	_g_2 := 0
	for _g_2 < acceptedRequests.Len() {
		request := func(hx_value_75 any) *app__core__FluxRequest {
			if hx_value_75 == nil {
				var hx_zero_76 *app__core__FluxRequest
				return hx_zero_76
			}
			return hx_value_75.(*app__core__FluxRequest)
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
		item := func(hx_value_78 any) *app__core__FluxRouteAggregate {
			if hx_value_78 == nil {
				var hx_zero_79 *app__core__FluxRouteAggregate
				return hx_zero_79
			}
			return hx_value_78.(*app__core__FluxRouteAggregate)
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
		state := func(hx_value_80 any) map[string]any {
			if hx_value_80 == nil {
				var hx_zero_81 map[string]any
				return hx_zero_81
			}
			return hx_value_80.(map[string]any)
		}(states.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_obj_85 map[string]any) *string {
			hx_field_86 := hx_obj_85["key"]
			if hx_field_86 == nil {
				var hx_zero_87 *string
				return hx_zero_87
			}
			return hx_field_86.(*string)
		}(state), key) {
			return func(hx_obj_82 map[string]any) int {
				hx_field_83 := hx_obj_82["value"]
				if hx_field_83 == nil {
					var hx_zero_84 int
					return hx_zero_84
				}
				return hx_field_83.(int)
			}(state)
		}
	}
	return 0
}

func (self *app__core__FluxPipeline) setStringIntStateValue(states *hxrt.Array, key *string, value int) *hxrt.Array {
	_g := 0
	for _g < states.Len() {
		state := func(hx_value_88 any) map[string]any {
			if hx_value_88 == nil {
				var hx_zero_89 map[string]any
				return hx_zero_89
			}
			return hx_value_88.(map[string]any)
		}(states.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(func(hx_obj_90 map[string]any) *string {
			hx_field_91 := hx_obj_90["key"]
			if hx_field_91 == nil {
				var hx_zero_92 *string
				return hx_zero_92
			}
			return hx_field_91.(*string)
		}(state), key) {
			state["value"] = value
			return states
		}
	}
	hx_obj_94 := map[string]any{}
	hx_obj_94["key"] = key
	hx_obj_94["value"] = value
	states.Push(hx_obj_94)
	return states
}

func (self *app__core__FluxPipeline) getResponseById(states *hxrt.Array, requestId int) *app__core__FluxProxyResponse {
	_g := 0
	for _g < states.Len() {
		state := func(hx_value_95 any) map[string]any {
			if hx_value_95 == nil {
				var hx_zero_96 map[string]any
				return hx_zero_96
			}
			return hx_value_95.(map[string]any)
		}(states.Get(_g))
		_g = int(int32((_g + 1)))
		if func(hx_obj_100 map[string]any) int {
			hx_field_101 := hx_obj_100["requestId"]
			if hx_field_101 == nil {
				var hx_zero_102 int
				return hx_zero_102
			}
			return hx_field_101.(int)
		}(state) == requestId {
			return func(hx_obj_97 map[string]any) *app__core__FluxProxyResponse {
				hx_field_98 := hx_obj_97["response"]
				if hx_field_98 == nil {
					var hx_zero_99 *app__core__FluxProxyResponse
					return hx_zero_99
				}
				return hx_field_98.(*app__core__FluxProxyResponse)
			}(state)
		}
	}
	return nil
}

func (self *app__core__FluxPipeline) setResponseById(states *hxrt.Array, requestId int, response *app__core__FluxProxyResponse) *hxrt.Array {
	_g := 0
	for _g < states.Len() {
		state := func(hx_value_103 any) map[string]any {
			if hx_value_103 == nil {
				var hx_zero_104 map[string]any
				return hx_zero_104
			}
			return hx_value_103.(map[string]any)
		}(states.Get(_g))
		_g = int(int32((_g + 1)))
		if func(hx_obj_105 map[string]any) int {
			hx_field_106 := hx_obj_105["requestId"]
			if hx_field_106 == nil {
				var hx_zero_107 int
				return hx_zero_107
			}
			return hx_field_106.(int)
		}(state) == requestId {
			state["response"] = response
			return states
		}
	}
	hx_obj_109 := map[string]any{}
	hx_obj_109["requestId"] = requestId
	hx_obj_109["response"] = response
	states.Push(hx_obj_109)
	return states
}
