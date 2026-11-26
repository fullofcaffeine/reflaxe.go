package main

import "examples_fluxproxy_portable/hxrt"

func hasArg(flag *string) bool {
	_g := 0
	_g1 := Sys_args()
	for _g < _g1.Len() {
		arg := func(hx_value_1 any) *string {
			if hx_value_1 == nil {
				var hx_zero_2 *string
				return hx_zero_2
			}
			return hx_value_1.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(arg, flag) {
			return true
		}
	}
	return false
}

func main() {
	var runtime app__runtime__FluxRuntime = app__runtime__RuntimeFactory_create()
	if hasArg(hxrt.StringFromLiteral("--scripted")) {
		var v any = any(Harness_run(runtime))
		hxrt.Println(v)
	} else {
		InteractiveCli_run(runtime)
	}
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *app__core__FluxIngestResult:
		return func(hx_structural_source_1 *app__core__FluxIngestResult) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "receivedCount":
					return hxrt.AnonymousField{Value: hx_structural_source_1.receivedCount, Found: true}
				case "acceptedRequests":
					return hxrt.AnonymousField{Value: hx_structural_source_1.acceptedRequests, Found: true}
				case "backpressureEvents":
					return hxrt.AnonymousField{Value: hx_structural_source_1.backpressureEvents, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *app__core__FluxPipeline:
		return func(hx_structural_source_2 *app__core__FluxPipeline) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_2, func(name string) hxrt.AnonymousField {
				switch name {
				case "runtime":
					return hxrt.AnonymousField{Value: hx_structural_source_2.runtime, Found: true}
				case "run":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array) *app__core__FluxReport] {
						hx_receiver_3 := hx_structural_source_2
						return hxrt.BoundFunction(func(requests *hxrt.Array) *app__core__FluxReport {
							return hx_receiver_3.__hx_this.run(requests)
						}, func() any {
							return hx_receiver_3.__hx_this
						}, "run")
					}(), Found: true}
				case "ingest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, int) *app__core__FluxIngestResult] {
						hx_receiver_4 := hx_structural_source_2
						return hxrt.BoundFunction(func(requests *hxrt.Array, capacity int) *app__core__FluxIngestResult {
							return hx_receiver_4.__hx_this.ingest(requests, capacity)
						}, func() any {
							return hx_receiver_4.__hx_this
						}, "ingest")
					}(), Found: true}
				case "aggregate":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array) hxrt.Anonymous] {
						hx_receiver_5 := hx_structural_source_2
						return hxrt.BoundFunction(func(responses *hxrt.Array) hxrt.Anonymous {
							return hx_receiver_5.__hx_this.aggregate(responses)
						}, func() any {
							return hx_receiver_5.__hx_this
						}, "aggregate")
					}(), Found: true}
				case "retries":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array) int] {
						hx_receiver_6 := hx_structural_source_2
						return hxrt.BoundFunction(func(responses *hxrt.Array) int {
							return hx_receiver_6.__hx_this.retries(responses)
						}, func() any {
							return hx_receiver_6.__hx_this
						}, "retries")
					}(), Found: true}
				case "errors":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array) int] {
						hx_receiver_7 := hx_structural_source_2
						return hxrt.BoundFunction(func(responses *hxrt.Array) int {
							return hx_receiver_7.__hx_this.errors(responses)
						}, func() any {
							return hx_receiver_7.__hx_this
						}, "errors")
					}(), Found: true}
				case "applyRoutePolicies":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, int, int, int) hxrt.Anonymous] {
						hx_receiver_8 := hx_structural_source_2
						return hxrt.BoundFunction(func(requests *hxrt.Array, perRouteLimit int, breakerFailureThreshold int, timeoutMs int) hxrt.Anonymous {
							return hx_receiver_8.__hx_this.applyRoutePolicies(requests, perRouteLimit, breakerFailureThreshold, timeoutMs)
						}, func() any {
							return hx_receiver_8.__hx_this
						}, "applyRoutePolicies")
					}(), Found: true}
				case "orderedResponses":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, *hxrt.Array, *hxrt.Array) *hxrt.Array] {
						hx_receiver_9 := hx_structural_source_2
						return hxrt.BoundFunction(func(synthetic *hxrt.Array, dispatched *hxrt.Array, acceptedRequests *hxrt.Array) *hxrt.Array {
							return hx_receiver_9.__hx_this.orderedResponses(synthetic, dispatched, acceptedRequests)
						}, func() any {
							return hx_receiver_9.__hx_this
						}, "orderedResponses")
					}(), Found: true}
				case "findRouteAggregate":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, *string) *app__core__FluxRouteAggregate] {
						hx_receiver_10 := hx_structural_source_2
						return hxrt.BoundFunction(func(routes *hxrt.Array, route *string) *app__core__FluxRouteAggregate {
							return hx_receiver_10.__hx_this.findRouteAggregate(routes, route)
						}, func() any {
							return hx_receiver_10.__hx_this
						}, "findRouteAggregate")
					}(), Found: true}
				case "getStringIntStateValue":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, *string) int] {
						hx_receiver_11 := hx_structural_source_2
						return hxrt.BoundFunction(func(states *hxrt.Array, key *string) int {
							return hx_receiver_11.__hx_this.getStringIntStateValue(states, key)
						}, func() any {
							return hx_receiver_11.__hx_this
						}, "getStringIntStateValue")
					}(), Found: true}
				case "setStringIntStateValue":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, *string, int) *hxrt.Array] {
						hx_receiver_12 := hx_structural_source_2
						return hxrt.BoundFunction(func(states *hxrt.Array, key *string, value int) *hxrt.Array {
							return hx_receiver_12.__hx_this.setStringIntStateValue(states, key, value)
						}, func() any {
							return hx_receiver_12.__hx_this
						}, "setStringIntStateValue")
					}(), Found: true}
				case "getResponseById":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, int) *app__core__FluxProxyResponse] {
						hx_receiver_13 := hx_structural_source_2
						return hxrt.BoundFunction(func(states *hxrt.Array, requestId int) *app__core__FluxProxyResponse {
							return hx_receiver_13.__hx_this.getResponseById(states, requestId)
						}, func() any {
							return hx_receiver_13.__hx_this
						}, "getResponseById")
					}(), Found: true}
				case "setResponseById":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, int, *app__core__FluxProxyResponse) *hxrt.Array] {
						hx_receiver_14 := hx_structural_source_2
						return hxrt.BoundFunction(func(states *hxrt.Array, requestId int, response *app__core__FluxProxyResponse) *hxrt.Array {
							return hx_receiver_14.__hx_this.setResponseById(states, requestId, response)
						}, func() any {
							return hx_receiver_14.__hx_this
						}, "setResponseById")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *app__core__FluxProxyResponse:
		return func(hx_structural_source_15 *app__core__FluxProxyResponse) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_15, func(name string) hxrt.AnonymousField {
				switch name {
				case "requestId":
					return hxrt.AnonymousField{Value: hx_structural_source_15.requestId, Found: true}
				case "route":
					return hxrt.AnonymousField{Value: hx_structural_source_15.route, Found: true}
				case "upstream":
					return hxrt.AnonymousField{Value: hx_structural_source_15.upstream, Found: true}
				case "status":
					return hxrt.AnonymousField{Value: hx_structural_source_15.status, Found: true}
				case "latencyMs":
					return hxrt.AnonymousField{Value: hx_structural_source_15.latencyMs, Found: true}
				case "attempts":
					return hxrt.AnonymousField{Value: hx_structural_source_15.attempts, Found: true}
				case "success":
					return hxrt.AnonymousField{Value: hx_structural_source_15.success, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *app__core__FluxReport:
		return func(hx_structural_source_16 *app__core__FluxReport) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_16, func(name string) hxrt.AnonymousField {
				switch name {
				case "profile":
					return hxrt.AnonymousField{Value: hx_structural_source_16.profile, Found: true}
				case "variant":
					return hxrt.AnonymousField{Value: hx_structural_source_16.variant, Found: true}
				case "capability":
					return hxrt.AnonymousField{Value: hx_structural_source_16.capability, Found: true}
				case "ingressReceived":
					return hxrt.AnonymousField{Value: hx_structural_source_16.ingressReceived, Found: true}
				case "ingressAccepted":
					return hxrt.AnonymousField{Value: hx_structural_source_16.ingressAccepted, Found: true}
				case "ingressBackpressure":
					return hxrt.AnonymousField{Value: hx_structural_source_16.ingressBackpressure, Found: true}
				case "proxyResponses":
					return hxrt.AnonymousField{Value: hx_structural_source_16.proxyResponses, Found: true}
				case "proxyRetries":
					return hxrt.AnonymousField{Value: hx_structural_source_16.proxyRetries, Found: true}
				case "rateLimitedCount":
					return hxrt.AnonymousField{Value: hx_structural_source_16.rateLimitedCount, Found: true}
				case "breakerOpenCount":
					return hxrt.AnonymousField{Value: hx_structural_source_16.breakerOpenCount, Found: true}
				case "routesCount":
					return hxrt.AnonymousField{Value: hx_structural_source_16.routesCount, Found: true}
				case "routesSummary":
					return hxrt.AnonymousField{Value: hx_structural_source_16.routesSummary, Found: true}
				case "errorsCount":
					return hxrt.AnonymousField{Value: hx_structural_source_16.errorsCount, Found: true}
				case "runtimeScore":
					return hxrt.AnonymousField{Value: hx_structural_source_16.runtimeScore, Found: true}
				case "lines":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.Array] {
						hx_receiver_17 := hx_structural_source_16
						return hxrt.BoundFunction(func() *hxrt.Array {
							return hx_receiver_17.__hx_this.lines()
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "lines")
					}(), Found: true}
				case "profileId":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_18 := hx_structural_source_16
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_18.__hx_this.profileId()
						}, func() any {
							return hx_receiver_18.__hx_this
						}, "profileId")
					}(), Found: true}
				case "variantId":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_19 := hx_structural_source_16
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_19.__hx_this.variantId()
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "variantId")
					}(), Found: true}
				case "capabilityId":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_20 := hx_structural_source_16
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_20.__hx_this.capabilityId()
						}, func() any {
							return hx_receiver_20.__hx_this
						}, "capabilityId")
					}(), Found: true}
				case "receivedCount":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_21 := hx_structural_source_16
						return hxrt.BoundFunction(func() int {
							return hx_receiver_21.__hx_this.receivedCount()
						}, func() any {
							return hx_receiver_21.__hx_this
						}, "receivedCount")
					}(), Found: true}
				case "acceptedCount":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_22 := hx_structural_source_16
						return hxrt.BoundFunction(func() int {
							return hx_receiver_22.__hx_this.acceptedCount()
						}, func() any {
							return hx_receiver_22.__hx_this
						}, "acceptedCount")
					}(), Found: true}
				case "backpressureCount":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_23 := hx_structural_source_16
						return hxrt.BoundFunction(func() int {
							return hx_receiver_23.__hx_this.backpressureCount()
						}, func() any {
							return hx_receiver_23.__hx_this
						}, "backpressureCount")
					}(), Found: true}
				case "retriesCount":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_24 := hx_structural_source_16
						return hxrt.BoundFunction(func() int {
							return hx_receiver_24.__hx_this.retriesCount()
						}, func() any {
							return hx_receiver_24.__hx_this
						}, "retriesCount")
					}(), Found: true}
				case "rateLimited":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_25 := hx_structural_source_16
						return hxrt.BoundFunction(func() int {
							return hx_receiver_25.__hx_this.rateLimited()
						}, func() any {
							return hx_receiver_25.__hx_this
						}, "rateLimited")
					}(), Found: true}
				case "breakerOpen":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_26 := hx_structural_source_16
						return hxrt.BoundFunction(func() int {
							return hx_receiver_26.__hx_this.breakerOpen()
						}, func() any {
							return hx_receiver_26.__hx_this
						}, "breakerOpen")
					}(), Found: true}
				case "errors":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_27 := hx_structural_source_16
						return hxrt.BoundFunction(func() int {
							return hx_receiver_27.__hx_this.errors()
						}, func() any {
							return hx_receiver_27.__hx_this
						}, "errors")
					}(), Found: true}
				case "score":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_28 := hx_structural_source_16
						return hxrt.BoundFunction(func() int {
							return hx_receiver_28.__hx_this.score()
						}, func() any {
							return hx_receiver_28.__hx_this
						}, "score")
					}(), Found: true}
				case "render":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_29 := hx_structural_source_16
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_29.__hx_this.render()
						}, func() any {
							return hx_receiver_29.__hx_this
						}, "render")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *app__core__FluxRequest:
		return func(hx_structural_source_30 *app__core__FluxRequest) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_30, func(name string) hxrt.AnonymousField {
				switch name {
				case "id":
					return hxrt.AnonymousField{Value: hx_structural_source_30.id, Found: true}
				case "route":
					return hxrt.AnonymousField{Value: hx_structural_source_30.route, Found: true}
				case "latencyMs":
					return hxrt.AnonymousField{Value: hx_structural_source_30.latencyMs, Found: true}
				case "status":
					return hxrt.AnonymousField{Value: hx_structural_source_30.status, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *app__core__FluxRouteAggregate:
		return func(hx_structural_source_31 *app__core__FluxRouteAggregate) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_31, func(name string) hxrt.AnonymousField {
				switch name {
				case "route":
					return hxrt.AnonymousField{Value: hx_structural_source_31.route, Found: true}
				case "count":
					return hxrt.AnonymousField{Value: hx_structural_source_31.count, Found: true}
				case "successCount":
					return hxrt.AnonymousField{Value: hx_structural_source_31.successCount, Found: true}
				case "errorCount":
					return hxrt.AnonymousField{Value: hx_structural_source_31.errorCount, Found: true}
				case "totalLatencyMs":
					return hxrt.AnonymousField{Value: hx_structural_source_31.totalLatencyMs, Found: true}
				case "record":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*app__core__FluxProxyResponse)] {
						hx_receiver_32 := hx_structural_source_31
						return hxrt.BoundFunction(func(response *app__core__FluxProxyResponse) {
							hx_receiver_32.__hx_this.record(response)
						}, func() any {
							return hx_receiver_32.__hx_this
						}, "record")
					}(), Found: true}
				case "averageLatencyMs":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_33 := hx_structural_source_31
						return hxrt.BoundFunction(func() int {
							return hx_receiver_33.__hx_this.averageLatencyMs()
						}, func() any {
							return hx_receiver_33.__hx_this
						}, "averageLatencyMs")
					}(), Found: true}
				case "summaryToken":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_34 := hx_structural_source_31
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_34.__hx_this.summaryToken()
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "summaryToken")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *app__runtime__CoreRuntime:
		return func(hx_structural_source_35 *app__runtime__CoreRuntime) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_35, func(name string) hxrt.AnonymousField {
				switch name {
				case "profileId":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_36 := hx_structural_source_35
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_36.__hx_this.profileId()
						}, func() any {
							return hx_receiver_36.__hx_this
						}, "profileId")
					}(), Found: true}
				case "variantId":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_37 := hx_structural_source_35
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_37.__hx_this.variantId()
						}, func() any {
							return hx_receiver_37.__hx_this
						}, "variantId")
					}(), Found: true}
				case "capabilityId":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_38 := hx_structural_source_35
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_38.__hx_this.capabilityId()
						}, func() any {
							return hx_receiver_38.__hx_this
						}, "capabilityId")
					}(), Found: true}
				case "dispatch":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, int) *hxrt.Array] {
						hx_receiver_39 := hx_structural_source_35
						return hxrt.BoundFunction(func(requests *hxrt.Array, workerCount int) *hxrt.Array {
							return hx_receiver_39.__hx_this.dispatch(requests, workerCount)
						}, func() any {
							return hx_receiver_39.__hx_this
						}, "dispatch")
					}(), Found: true}
				case "stageScore":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, int, int) int] {
						hx_receiver_40 := hx_structural_source_35
						return hxrt.BoundFunction(func(responses *hxrt.Array, retryCount int, backpressureEvents int) int {
							return hx_receiver_40.__hx_this.stageScore(responses, retryCount, backpressureEvents)
						}, func() any {
							return hx_receiver_40.__hx_this
						}, "stageScore")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_41 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_41, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_41.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_41.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_41.high = func(hx_value_42 any) int {
						if hx_value_42 == nil {
							var hx_zero_43 int
							return hx_zero_43
						}
						return hx_value_42.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_41.low = func(hx_value_44 any) int {
						if hx_value_44 == nil {
							var hx_zero_45 int
							return hx_zero_45
						}
						return hx_value_44.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_46 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_46, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_46.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_46.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_46.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_46.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_49 := hx_structural_source_46
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_49.__hx_this.get_message()
						}, func() any {
							return hx_receiver_49.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_50 := hx_structural_source_46
						return hxrt.BoundFunction(func() any {
							return hx_receiver_50.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_50.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_51 := hx_structural_source_46
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_51.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_51.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_52 := hx_structural_source_46
						return hxrt.BoundFunction(func() any {
							return hx_receiver_52.__hx_this.get_native()
						}, func() any {
							return hx_receiver_52.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_53 := hx_structural_source_46
						return hxrt.BoundFunction(func() any {
							return hx_receiver_53.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_53.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_54 := hx_structural_source_46
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_54.__hx_this.toString()
						}, func() any {
							return hx_receiver_54.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_55 := hx_structural_source_46
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_55.__hx_this.details()
						}, func() any {
							return hx_receiver_55.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_46.savedMessage = func(hx_value_47 any) *string {
						if hx_value_47 == nil {
							var hx_zero_48 *string
							return hx_zero_48
						}
						return hx_value_47.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_46.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_46.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_46.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_56 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_56, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_56.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_56.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_61 := hx_structural_source_56
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_61.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_61.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_62 := hx_structural_source_56
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_62.__hx_this.next()
						}, func() any {
							return hx_receiver_62.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_56.offset = func(hx_value_57 any) int {
						if hx_value_57 == nil {
							var hx_zero_58 int
							return hx_zero_58
						}
						return hx_value_57.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_56.s = func(hx_value_59 any) *string {
						if hx_value_59 == nil {
							var hx_zero_60 *string
							return hx_zero_60
						}
						return hx_value_59.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_63 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_63, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_63.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_63.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_68 := hx_structural_source_63
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_68.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_68.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_69 := hx_structural_source_63
						return hxrt.BoundFunction(func() int {
							return hx_receiver_69.__hx_this.next()
						}, func() any {
							return hx_receiver_69.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_63.offset = func(hx_value_64 any) int {
						if hx_value_64 == nil {
							var hx_zero_65 int
							return hx_zero_65
						}
						return hx_value_64.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_63.s = func(hx_value_66 any) *string {
						if hx_value_66 == nil {
							var hx_zero_67 *string
							return hx_zero_67
						}
						return hx_value_66.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_70 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_70, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_70.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_70.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_70.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_70.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_70.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_79 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_79.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_79.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_80 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_80.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_80.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_81 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len_local int) {
							hx_receiver_81.__hx_this.blit(pos, src, srcpos, len_local)
						}, func() any {
							return hx_receiver_81.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_82 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int, len_local int, value int) {
							hx_receiver_82.__hx_this.fill(pos, len_local, value)
						}, func() any {
							return hx_receiver_82.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_83 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int, len_local int) *haxe__io__Bytes {
							return hx_receiver_83.__hx_this.sub(pos, len_local)
						}, func() any {
							return hx_receiver_83.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_84 := hx_structural_source_70
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_84.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_84.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_85 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_85.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_85.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_86 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_86.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_86.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_87 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_87.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_87.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_88 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_88.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_88.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_89 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_89.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_89.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_90 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_90.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_90.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_91 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_91.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_91.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_92 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_92.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_92.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_93 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_93.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_93.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_94 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_94.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_94.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_95 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int, len_local int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_95.__hx_this.getString(pos, len_local, encoding)
						}, func() any {
							return hx_receiver_95.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_96 := hx_structural_source_70
						return hxrt.BoundFunction(func(pos int, len_local int) *string {
							return hx_receiver_96.__hx_this.readString(pos, len_local)
						}, func() any {
							return hx_receiver_96.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_97 := hx_structural_source_70
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_97.__hx_this.toString()
						}, func() any {
							return hx_receiver_97.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_98 := hx_structural_source_70
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_98.__hx_this.toHex()
						}, func() any {
							return hx_receiver_98.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_99 := hx_structural_source_70
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_99.__hx_this.getData()
						}, func() any {
							return hx_receiver_99.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_100 := hx_structural_source_70
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_100.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_100.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_70.b = func(hx_value_71 any) []int {
						if hx_value_71 == nil {
							var hx_zero_72 []int
							return hx_zero_72
						}
						return hx_value_71.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_70.__hx_raw = func(hx_value_73 any) *hxrt.ByteView {
						if hx_value_73 == nil {
							var hx_zero_74 *hxrt.ByteView
							return hx_zero_74
						}
						return hx_value_73.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_70.__hx_rawValid = func(hx_value_75 any) bool {
						if hx_value_75 == nil {
							var hx_zero_76 bool
							return hx_zero_76
						}
						return hx_value_75.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_70.__hx_dataExposed = func(hx_value_77 any) bool {
						if hx_value_77 == nil {
							var hx_zero_78 bool
							return hx_zero_78
						}
						return hx_value_77.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_101 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_101, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_101.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_102 := hx_structural_source_101
						return hxrt.BoundFunction(func() any {
							return hx_receiver_102.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_102.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_101.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_101.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_101.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_101.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_105 := hx_structural_source_101
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_105.__hx_this.get_message()
						}, func() any {
							return hx_receiver_105.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_106 := hx_structural_source_101
						return hxrt.BoundFunction(func() any {
							return hx_receiver_106.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_106.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_107 := hx_structural_source_101
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_107.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_107.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_108 := hx_structural_source_101
						return hxrt.BoundFunction(func() any {
							return hx_receiver_108.__hx_this.get_native()
						}, func() any {
							return hx_receiver_108.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_109 := hx_structural_source_101
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_109.__hx_this.toString()
						}, func() any {
							return hx_receiver_109.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_110 := hx_structural_source_101
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_110.__hx_this.details()
						}, func() any {
							return hx_receiver_110.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_101.savedMessage = func(hx_value_103 any) *string {
						if hx_value_103 == nil {
							var hx_zero_104 *string
							return hx_zero_104
						}
						return hx_value_103.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_101.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_101.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_101.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_class_app__core__FluxProxyResponse(value any) *app__core__FluxProxyResponse {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__core__FluxProxyResponse:
		return source
	default:
		return value.(*app__core__FluxProxyResponse)
	}
}

func hx_recover_class_app__core__FluxRequest(value any) *app__core__FluxRequest {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__core__FluxRequest:
		return source
	default:
		return value.(*app__core__FluxRequest)
	}
}

func hx_recover_class_app__core__FluxRouteAggregate(value any) *app__core__FluxRouteAggregate {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__core__FluxRouteAggregate:
		return source
	default:
		return value.(*app__core__FluxRouteAggregate)
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
