package main

import "examples_fluxproxy_portable/hxrt"

func hasArg(flag *string) bool {
	_g := 0
	_g1 := hxrt.ArrayFromValues(func(hx_sort_src_1 []*string) []any {
		hx_sort_out_3 := make([]any, 0, len(hx_sort_src_1))
		for _, hx_sort_item_2 := range hx_sort_src_1 {
			hx_sort_out_3 = append(hx_sort_out_3, hx_sort_item_2)
		}
		return hx_sort_out_3
	}(hxrt.SysArgs()))
	for _g < _g1.Len() {
		arg := func(hx_value_4 any) *string {
			if hx_value_4 == nil {
				var hx_zero_5 *string
				return hx_zero_5
			}
			return hx_value_4.(*string)
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
		return func(hx_structural_source_111 *app__core__FluxIngestResult) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_111, func(name string) hxrt.AnonymousField {
				switch name {
				case "receivedCount":
					return hxrt.AnonymousField{Value: hx_structural_source_111.receivedCount, Found: true}
				case "acceptedRequests":
					return hxrt.AnonymousField{Value: hx_structural_source_111.acceptedRequests, Found: true}
				case "backpressureEvents":
					return hxrt.AnonymousField{Value: hx_structural_source_111.backpressureEvents, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *app__core__FluxPipeline:
		return func(hx_structural_source_112 *app__core__FluxPipeline) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_112, func(name string) hxrt.AnonymousField {
				switch name {
				case "runtime":
					return hxrt.AnonymousField{Value: hx_structural_source_112.runtime, Found: true}
				case "run":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array) *app__core__FluxReport] {
						hx_receiver_113 := hx_structural_source_112
						return hxrt.BoundFunction(func(requests *hxrt.Array) *app__core__FluxReport {
							return hx_receiver_113.__hx_this.run(requests)
						}, func() any {
							return hx_receiver_113.__hx_this
						}, "run")
					}(), Found: true}
				case "ingest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, int) *app__core__FluxIngestResult] {
						hx_receiver_114 := hx_structural_source_112
						return hxrt.BoundFunction(func(requests *hxrt.Array, capacity int) *app__core__FluxIngestResult {
							return hx_receiver_114.__hx_this.ingest(requests, capacity)
						}, func() any {
							return hx_receiver_114.__hx_this
						}, "ingest")
					}(), Found: true}
				case "aggregate":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array) hxrt.Anonymous] {
						hx_receiver_115 := hx_structural_source_112
						return hxrt.BoundFunction(func(responses *hxrt.Array) hxrt.Anonymous {
							return hx_receiver_115.__hx_this.aggregate(responses)
						}, func() any {
							return hx_receiver_115.__hx_this
						}, "aggregate")
					}(), Found: true}
				case "retries":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array) int] {
						hx_receiver_116 := hx_structural_source_112
						return hxrt.BoundFunction(func(responses *hxrt.Array) int {
							return hx_receiver_116.__hx_this.retries(responses)
						}, func() any {
							return hx_receiver_116.__hx_this
						}, "retries")
					}(), Found: true}
				case "errors":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array) int] {
						hx_receiver_117 := hx_structural_source_112
						return hxrt.BoundFunction(func(responses *hxrt.Array) int {
							return hx_receiver_117.__hx_this.errors(responses)
						}, func() any {
							return hx_receiver_117.__hx_this
						}, "errors")
					}(), Found: true}
				case "applyRoutePolicies":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, int, int, int) hxrt.Anonymous] {
						hx_receiver_118 := hx_structural_source_112
						return hxrt.BoundFunction(func(requests *hxrt.Array, perRouteLimit int, breakerFailureThreshold int, timeoutMs int) hxrt.Anonymous {
							return hx_receiver_118.__hx_this.applyRoutePolicies(requests, perRouteLimit, breakerFailureThreshold, timeoutMs)
						}, func() any {
							return hx_receiver_118.__hx_this
						}, "applyRoutePolicies")
					}(), Found: true}
				case "orderedResponses":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, *hxrt.Array, *hxrt.Array) *hxrt.Array] {
						hx_receiver_119 := hx_structural_source_112
						return hxrt.BoundFunction(func(synthetic *hxrt.Array, dispatched *hxrt.Array, acceptedRequests *hxrt.Array) *hxrt.Array {
							return hx_receiver_119.__hx_this.orderedResponses(synthetic, dispatched, acceptedRequests)
						}, func() any {
							return hx_receiver_119.__hx_this
						}, "orderedResponses")
					}(), Found: true}
				case "findRouteAggregate":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, *string) *app__core__FluxRouteAggregate] {
						hx_receiver_120 := hx_structural_source_112
						return hxrt.BoundFunction(func(routes *hxrt.Array, route *string) *app__core__FluxRouteAggregate {
							return hx_receiver_120.__hx_this.findRouteAggregate(routes, route)
						}, func() any {
							return hx_receiver_120.__hx_this
						}, "findRouteAggregate")
					}(), Found: true}
				case "getStringIntStateValue":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, *string) int] {
						hx_receiver_121 := hx_structural_source_112
						return hxrt.BoundFunction(func(states *hxrt.Array, key *string) int {
							return hx_receiver_121.__hx_this.getStringIntStateValue(states, key)
						}, func() any {
							return hx_receiver_121.__hx_this
						}, "getStringIntStateValue")
					}(), Found: true}
				case "setStringIntStateValue":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, *string, int) *hxrt.Array] {
						hx_receiver_122 := hx_structural_source_112
						return hxrt.BoundFunction(func(states *hxrt.Array, key *string, value int) *hxrt.Array {
							return hx_receiver_122.__hx_this.setStringIntStateValue(states, key, value)
						}, func() any {
							return hx_receiver_122.__hx_this
						}, "setStringIntStateValue")
					}(), Found: true}
				case "getResponseById":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, int) *app__core__FluxProxyResponse] {
						hx_receiver_123 := hx_structural_source_112
						return hxrt.BoundFunction(func(states *hxrt.Array, requestId int) *app__core__FluxProxyResponse {
							return hx_receiver_123.__hx_this.getResponseById(states, requestId)
						}, func() any {
							return hx_receiver_123.__hx_this
						}, "getResponseById")
					}(), Found: true}
				case "setResponseById":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, int, *app__core__FluxProxyResponse) *hxrt.Array] {
						hx_receiver_124 := hx_structural_source_112
						return hxrt.BoundFunction(func(states *hxrt.Array, requestId int, response *app__core__FluxProxyResponse) *hxrt.Array {
							return hx_receiver_124.__hx_this.setResponseById(states, requestId, response)
						}, func() any {
							return hx_receiver_124.__hx_this
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
		return func(hx_structural_source_125 *app__core__FluxProxyResponse) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_125, func(name string) hxrt.AnonymousField {
				switch name {
				case "requestId":
					return hxrt.AnonymousField{Value: hx_structural_source_125.requestId, Found: true}
				case "route":
					return hxrt.AnonymousField{Value: hx_structural_source_125.route, Found: true}
				case "upstream":
					return hxrt.AnonymousField{Value: hx_structural_source_125.upstream, Found: true}
				case "status":
					return hxrt.AnonymousField{Value: hx_structural_source_125.status, Found: true}
				case "latencyMs":
					return hxrt.AnonymousField{Value: hx_structural_source_125.latencyMs, Found: true}
				case "attempts":
					return hxrt.AnonymousField{Value: hx_structural_source_125.attempts, Found: true}
				case "success":
					return hxrt.AnonymousField{Value: hx_structural_source_125.success, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *app__core__FluxReport:
		return func(hx_structural_source_126 *app__core__FluxReport) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_126, func(name string) hxrt.AnonymousField {
				switch name {
				case "profile":
					return hxrt.AnonymousField{Value: hx_structural_source_126.profile, Found: true}
				case "variant":
					return hxrt.AnonymousField{Value: hx_structural_source_126.variant, Found: true}
				case "capability":
					return hxrt.AnonymousField{Value: hx_structural_source_126.capability, Found: true}
				case "ingressReceived":
					return hxrt.AnonymousField{Value: hx_structural_source_126.ingressReceived, Found: true}
				case "ingressAccepted":
					return hxrt.AnonymousField{Value: hx_structural_source_126.ingressAccepted, Found: true}
				case "ingressBackpressure":
					return hxrt.AnonymousField{Value: hx_structural_source_126.ingressBackpressure, Found: true}
				case "proxyResponses":
					return hxrt.AnonymousField{Value: hx_structural_source_126.proxyResponses, Found: true}
				case "proxyRetries":
					return hxrt.AnonymousField{Value: hx_structural_source_126.proxyRetries, Found: true}
				case "rateLimitedCount":
					return hxrt.AnonymousField{Value: hx_structural_source_126.rateLimitedCount, Found: true}
				case "breakerOpenCount":
					return hxrt.AnonymousField{Value: hx_structural_source_126.breakerOpenCount, Found: true}
				case "routesCount":
					return hxrt.AnonymousField{Value: hx_structural_source_126.routesCount, Found: true}
				case "routesSummary":
					return hxrt.AnonymousField{Value: hx_structural_source_126.routesSummary, Found: true}
				case "errorsCount":
					return hxrt.AnonymousField{Value: hx_structural_source_126.errorsCount, Found: true}
				case "runtimeScore":
					return hxrt.AnonymousField{Value: hx_structural_source_126.runtimeScore, Found: true}
				case "lines":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.Array] {
						hx_receiver_127 := hx_structural_source_126
						return hxrt.BoundFunction(func() *hxrt.Array {
							return hx_receiver_127.__hx_this.lines()
						}, func() any {
							return hx_receiver_127.__hx_this
						}, "lines")
					}(), Found: true}
				case "profileId":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_128 := hx_structural_source_126
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_128.__hx_this.profileId()
						}, func() any {
							return hx_receiver_128.__hx_this
						}, "profileId")
					}(), Found: true}
				case "variantId":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_129 := hx_structural_source_126
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_129.__hx_this.variantId()
						}, func() any {
							return hx_receiver_129.__hx_this
						}, "variantId")
					}(), Found: true}
				case "capabilityId":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_130 := hx_structural_source_126
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_130.__hx_this.capabilityId()
						}, func() any {
							return hx_receiver_130.__hx_this
						}, "capabilityId")
					}(), Found: true}
				case "receivedCount":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_131 := hx_structural_source_126
						return hxrt.BoundFunction(func() int {
							return hx_receiver_131.__hx_this.receivedCount()
						}, func() any {
							return hx_receiver_131.__hx_this
						}, "receivedCount")
					}(), Found: true}
				case "acceptedCount":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_132 := hx_structural_source_126
						return hxrt.BoundFunction(func() int {
							return hx_receiver_132.__hx_this.acceptedCount()
						}, func() any {
							return hx_receiver_132.__hx_this
						}, "acceptedCount")
					}(), Found: true}
				case "backpressureCount":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_133 := hx_structural_source_126
						return hxrt.BoundFunction(func() int {
							return hx_receiver_133.__hx_this.backpressureCount()
						}, func() any {
							return hx_receiver_133.__hx_this
						}, "backpressureCount")
					}(), Found: true}
				case "retriesCount":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_134 := hx_structural_source_126
						return hxrt.BoundFunction(func() int {
							return hx_receiver_134.__hx_this.retriesCount()
						}, func() any {
							return hx_receiver_134.__hx_this
						}, "retriesCount")
					}(), Found: true}
				case "rateLimited":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_135 := hx_structural_source_126
						return hxrt.BoundFunction(func() int {
							return hx_receiver_135.__hx_this.rateLimited()
						}, func() any {
							return hx_receiver_135.__hx_this
						}, "rateLimited")
					}(), Found: true}
				case "breakerOpen":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_136 := hx_structural_source_126
						return hxrt.BoundFunction(func() int {
							return hx_receiver_136.__hx_this.breakerOpen()
						}, func() any {
							return hx_receiver_136.__hx_this
						}, "breakerOpen")
					}(), Found: true}
				case "errors":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_137 := hx_structural_source_126
						return hxrt.BoundFunction(func() int {
							return hx_receiver_137.__hx_this.errors()
						}, func() any {
							return hx_receiver_137.__hx_this
						}, "errors")
					}(), Found: true}
				case "score":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_138 := hx_structural_source_126
						return hxrt.BoundFunction(func() int {
							return hx_receiver_138.__hx_this.score()
						}, func() any {
							return hx_receiver_138.__hx_this
						}, "score")
					}(), Found: true}
				case "render":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_139 := hx_structural_source_126
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_139.__hx_this.render()
						}, func() any {
							return hx_receiver_139.__hx_this
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
		return func(hx_structural_source_140 *app__core__FluxRequest) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_140, func(name string) hxrt.AnonymousField {
				switch name {
				case "id":
					return hxrt.AnonymousField{Value: hx_structural_source_140.id, Found: true}
				case "route":
					return hxrt.AnonymousField{Value: hx_structural_source_140.route, Found: true}
				case "latencyMs":
					return hxrt.AnonymousField{Value: hx_structural_source_140.latencyMs, Found: true}
				case "status":
					return hxrt.AnonymousField{Value: hx_structural_source_140.status, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *app__core__FluxRouteAggregate:
		return func(hx_structural_source_141 *app__core__FluxRouteAggregate) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_141, func(name string) hxrt.AnonymousField {
				switch name {
				case "route":
					return hxrt.AnonymousField{Value: hx_structural_source_141.route, Found: true}
				case "count":
					return hxrt.AnonymousField{Value: hx_structural_source_141.count, Found: true}
				case "successCount":
					return hxrt.AnonymousField{Value: hx_structural_source_141.successCount, Found: true}
				case "errorCount":
					return hxrt.AnonymousField{Value: hx_structural_source_141.errorCount, Found: true}
				case "totalLatencyMs":
					return hxrt.AnonymousField{Value: hx_structural_source_141.totalLatencyMs, Found: true}
				case "record":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*app__core__FluxProxyResponse)] {
						hx_receiver_142 := hx_structural_source_141
						return hxrt.BoundFunction(func(response *app__core__FluxProxyResponse) {
							hx_receiver_142.__hx_this.record(response)
						}, func() any {
							return hx_receiver_142.__hx_this
						}, "record")
					}(), Found: true}
				case "averageLatencyMs":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_143 := hx_structural_source_141
						return hxrt.BoundFunction(func() int {
							return hx_receiver_143.__hx_this.averageLatencyMs()
						}, func() any {
							return hx_receiver_143.__hx_this
						}, "averageLatencyMs")
					}(), Found: true}
				case "summaryToken":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_144 := hx_structural_source_141
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_144.__hx_this.summaryToken()
						}, func() any {
							return hx_receiver_144.__hx_this
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
		return func(hx_structural_source_145 *app__runtime__CoreRuntime) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_145, func(name string) hxrt.AnonymousField {
				switch name {
				case "profileId":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_146 := hx_structural_source_145
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_146.__hx_this.profileId()
						}, func() any {
							return hx_receiver_146.__hx_this
						}, "profileId")
					}(), Found: true}
				case "variantId":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_147 := hx_structural_source_145
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_147.__hx_this.variantId()
						}, func() any {
							return hx_receiver_147.__hx_this
						}, "variantId")
					}(), Found: true}
				case "capabilityId":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_148 := hx_structural_source_145
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_148.__hx_this.capabilityId()
						}, func() any {
							return hx_receiver_148.__hx_this
						}, "capabilityId")
					}(), Found: true}
				case "dispatch":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, int) *hxrt.Array] {
						hx_receiver_149 := hx_structural_source_145
						return hxrt.BoundFunction(func(requests *hxrt.Array, workerCount int) *hxrt.Array {
							return hx_receiver_149.__hx_this.dispatch(requests, workerCount)
						}, func() any {
							return hx_receiver_149.__hx_this
						}, "dispatch")
					}(), Found: true}
				case "stageScore":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, int, int) int] {
						hx_receiver_150 := hx_structural_source_145
						return hxrt.BoundFunction(func(responses *hxrt.Array, retryCount int, backpressureEvents int) int {
							return hx_receiver_150.__hx_this.stageScore(responses, retryCount, backpressureEvents)
						}, func() any {
							return hx_receiver_150.__hx_this
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
		return func(hx_structural_source_151 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_151, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_151.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_151.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_151.high = func(hx_value_152 any) int {
						if hx_value_152 == nil {
							var hx_zero_153 int
							return hx_zero_153
						}
						return hx_value_152.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_151.low = func(hx_value_154 any) int {
						if hx_value_154 == nil {
							var hx_zero_155 int
							return hx_zero_155
						}
						return hx_value_154.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_156 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_156, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_156.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_156.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_156.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_156.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_159 := hx_structural_source_156
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_159.__hx_this.get_message()
						}, func() any {
							return hx_receiver_159.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_160 := hx_structural_source_156
						return hxrt.BoundFunction(func() any {
							return hx_receiver_160.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_160.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_161 := hx_structural_source_156
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_161.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_161.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_162 := hx_structural_source_156
						return hxrt.BoundFunction(func() any {
							return hx_receiver_162.__hx_this.get_native()
						}, func() any {
							return hx_receiver_162.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_163 := hx_structural_source_156
						return hxrt.BoundFunction(func() any {
							return hx_receiver_163.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_163.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_164 := hx_structural_source_156
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_164.__hx_this.toString()
						}, func() any {
							return hx_receiver_164.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_165 := hx_structural_source_156
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_165.__hx_this.details()
						}, func() any {
							return hx_receiver_165.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_156.savedMessage = func(hx_value_157 any) *string {
						if hx_value_157 == nil {
							var hx_zero_158 *string
							return hx_zero_158
						}
						return hx_value_157.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_156.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_156.savedPrevious = hx_recover_class_4(incoming)
					return true
				case "savedNative":
					hx_structural_source_156.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_166 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_166, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_166.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_166.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_171 := hx_structural_source_166
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_171.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_171.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_172 := hx_structural_source_166
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_172.__hx_this.next()
						}, func() any {
							return hx_receiver_172.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_166.offset = func(hx_value_167 any) int {
						if hx_value_167 == nil {
							var hx_zero_168 int
							return hx_zero_168
						}
						return hx_value_167.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_166.s = func(hx_value_169 any) *string {
						if hx_value_169 == nil {
							var hx_zero_170 *string
							return hx_zero_170
						}
						return hx_value_169.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_173 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_173, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_173.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_173.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_178 := hx_structural_source_173
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_178.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_178.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_179 := hx_structural_source_173
						return hxrt.BoundFunction(func() int {
							return hx_receiver_179.__hx_this.next()
						}, func() any {
							return hx_receiver_179.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_173.offset = func(hx_value_174 any) int {
						if hx_value_174 == nil {
							var hx_zero_175 int
							return hx_zero_175
						}
						return hx_value_174.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_173.s = func(hx_value_176 any) *string {
						if hx_value_176 == nil {
							var hx_zero_177 *string
							return hx_zero_177
						}
						return hx_value_176.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_180 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_180, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_180.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_180.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_180.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_180.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_180.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_189 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_189.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_189.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_190 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_190.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_190.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_191 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
							hx_receiver_191.__hx_this.blit(pos, src, srcpos, len)
						}, func() any {
							return hx_receiver_191.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_192 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int, len int, value int) {
							hx_receiver_192.__hx_this.fill(pos, len, value)
						}, func() any {
							return hx_receiver_192.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_193 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
							return hx_receiver_193.__hx_this.sub(pos, len)
						}, func() any {
							return hx_receiver_193.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_194 := hx_structural_source_180
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_194.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_194.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_195 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_195.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_195.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_196 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_196.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_196.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_197 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_197.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_197.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_198 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_198.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_198.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_199 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_199.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_199.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_200 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_200.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_200.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_201 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_201.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_201.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_202 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_202.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_202.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_203 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_203.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_203.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_204 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_204.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_204.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_205 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_205.__hx_this.getString(pos, len, encoding)
						}, func() any {
							return hx_receiver_205.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_206 := hx_structural_source_180
						return hxrt.BoundFunction(func(pos int, len int) *string {
							return hx_receiver_206.__hx_this.readString(pos, len)
						}, func() any {
							return hx_receiver_206.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_207 := hx_structural_source_180
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_207.__hx_this.toString()
						}, func() any {
							return hx_receiver_207.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_208 := hx_structural_source_180
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_208.__hx_this.toHex()
						}, func() any {
							return hx_receiver_208.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_209 := hx_structural_source_180
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_209.__hx_this.getData()
						}, func() any {
							return hx_receiver_209.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_210 := hx_structural_source_180
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_210.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_210.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_180.b = func(hx_value_181 any) []int {
						if hx_value_181 == nil {
							var hx_zero_182 []int
							return hx_zero_182
						}
						return hx_value_181.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_180.__hx_raw = func(hx_value_183 any) *hxrt.ByteView {
						if hx_value_183 == nil {
							var hx_zero_184 *hxrt.ByteView
							return hx_zero_184
						}
						return hx_value_183.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_180.__hx_rawValid = func(hx_value_185 any) bool {
						if hx_value_185 == nil {
							var hx_zero_186 bool
							return hx_zero_186
						}
						return hx_value_185.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_180.__hx_dataExposed = func(hx_value_187 any) bool {
						if hx_value_187 == nil {
							var hx_zero_188 bool
							return hx_zero_188
						}
						return hx_value_187.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_211 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_211, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_211.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_212 := hx_structural_source_211
						return hxrt.BoundFunction(func() any {
							return hx_receiver_212.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_212.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_211.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_211.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_211.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_211.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_215 := hx_structural_source_211
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_215.__hx_this.get_message()
						}, func() any {
							return hx_receiver_215.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_216 := hx_structural_source_211
						return hxrt.BoundFunction(func() any {
							return hx_receiver_216.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_216.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_217 := hx_structural_source_211
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_217.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_217.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_218 := hx_structural_source_211
						return hxrt.BoundFunction(func() any {
							return hx_receiver_218.__hx_this.get_native()
						}, func() any {
							return hx_receiver_218.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_219 := hx_structural_source_211
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_219.__hx_this.toString()
						}, func() any {
							return hx_receiver_219.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_220 := hx_structural_source_211
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_220.__hx_this.details()
						}, func() any {
							return hx_receiver_220.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_211.savedMessage = func(hx_value_213 any) *string {
						if hx_value_213 == nil {
							var hx_zero_214 *string
							return hx_zero_214
						}
						return hx_value_213.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_211.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_211.savedPrevious = hx_recover_class_4(incoming)
					return true
				case "savedNative":
					hx_structural_source_211.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_class_0(value any) *app__core__FluxRequest {
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

func hx_recover_class_1(value any) *app__core__FluxProxyResponse {
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

func hx_recover_class_2(value any) *app__core__FluxRouteAggregate {
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

func hx_recover_class_3(value any) *haxe__ValueException {
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

func hx_recover_class_4(value any) *haxe__Exception {
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
