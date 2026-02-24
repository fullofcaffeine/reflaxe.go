package main

import (
	"fmt"
	"strings"
)

const (
	ingestQueueCapacity     = 3
	perRouteLimit           = 2
	breakerFailureThreshold = 2
	timeoutMs               = 50
)

type fluxRequest struct {
	id        int
	route     string
	latencyMs int
	status    int
}

type fluxProxyResponse struct {
	requestID int
	route     string
	upstream  string
	status    int
	latencyMs int
	attempts  int
	success   bool
}

type fluxRouteAggregate struct {
	route        string
	count        int
	successCount int
	errorCount   int
	totalLatency int
}

func newFluxRouteAggregate(route string) *fluxRouteAggregate {
	return &fluxRouteAggregate{route: route}
}

func (a *fluxRouteAggregate) record(response fluxProxyResponse) {
	a.count++
	a.totalLatency += response.latencyMs
	if response.success {
		a.successCount++
	} else {
		a.errorCount++
	}
}

func (a *fluxRouteAggregate) averageLatencyMs() int {
	if a.count == 0 {
		return 0
	}
	remaining := a.totalLatency
	quotient := 0
	for remaining >= a.count {
		remaining -= a.count
		quotient++
	}
	return quotient
}

func (a *fluxRouteAggregate) summaryToken() string {
	return fmt.Sprintf("%s:%d/%d/%d/%d", a.route, a.count, a.successCount, a.errorCount, a.averageLatencyMs())
}

type fluxIngestResult struct {
	receivedCount      int
	acceptedRequests   []fluxRequest
	backpressureEvents int
}

type fluxReport struct {
	profile             string
	variant             string
	capability          string
	ingressReceived     int
	ingressAccepted     int
	ingressBackpressure int
	proxyResponses      int
	proxyRetries        int
	rateLimitedCount    int
	breakerOpenCount    int
	routesCount         int
	routesSummary       string
	errorsCount         int
	runtimeScore        int
}

func (r fluxReport) lines() []string {
	return []string{
		"fluxproxy.profile=" + r.profile,
		"fluxproxy.variant=" + r.variant,
		"runtime.capability=" + r.capability,
		"ingress.received=" + fmt.Sprintf("%d", r.ingressReceived),
		"ingress.accepted=" + fmt.Sprintf("%d", r.ingressAccepted),
		"ingress.backpressure=" + fmt.Sprintf("%d", r.ingressBackpressure),
		"proxy.responses=" + fmt.Sprintf("%d", r.proxyResponses),
		"proxy.retries=" + fmt.Sprintf("%d", r.proxyRetries),
		"policy.rate_limited=" + fmt.Sprintf("%d", r.rateLimitedCount),
		"policy.breaker_open=" + fmt.Sprintf("%d", r.breakerOpenCount),
		"routes.count=" + fmt.Sprintf("%d", r.routesCount),
		"routes.summary=" + r.routesSummary,
		"errors.count=" + fmt.Sprintf("%d", r.errorsCount),
		"runtime.score=" + fmt.Sprintf("%d", r.runtimeScore),
	}
}

func (r fluxReport) render() string {
	return strings.Join(r.lines(), "\n")
}

type fluxRuntime interface {
	profileID() string
	variantID() string
	capabilityID() string
	dispatch(requests []fluxRequest, workerCount int) []fluxProxyResponse
	stageScore(responses []fluxProxyResponse, retryCount int, backpressureEvents int) int
	dispatchWorkers() int
}

type coreRuntime struct {
	profile string
}

func (r coreRuntime) profileID() string    { return r.profile }
func (r coreRuntime) variantID() string    { return "core" }
func (r coreRuntime) capabilityID() string { return "loop_dispatch" }
func (r coreRuntime) dispatchWorkers() int { return 1 }

func (r coreRuntime) dispatch(requests []fluxRequest, _ int) []fluxProxyResponse {
	responses := make([]fluxProxyResponse, 0, len(requests))
	for _, request := range requests {
		responses = append(responses, proxy(request, timeoutMs))
	}
	return responses
}

func (r coreRuntime) stageScore(responses []fluxProxyResponse, retryCount int, backpressureEvents int) int {
	successCount := 0
	errorCount := 0
	for _, response := range responses {
		if response.success {
			successCount++
		} else {
			errorCount++
		}
	}

	score := 0
	score += successCount * 10
	score -= errorCount * 6
	score -= backpressureEvents * 2
	score -= retryCount * 2
	score += len(responses)
	return score
}

type goNativeRuntime struct {
	profile string
}

func (r goNativeRuntime) profileID() string    { return r.profile }
func (r goNativeRuntime) variantID() string    { return "go_native" }
func (r goNativeRuntime) capabilityID() string { return "worker_chan_fanout" }
func (r goNativeRuntime) dispatchWorkers() int { return 3 }

func (r goNativeRuntime) dispatch(requests []fluxRequest, workerCount int) []fluxProxyResponse {
	if len(requests) == 0 {
		return []fluxProxyResponse{}
	}

	workers := normalizeWorkers(workerCount)
	inbox := make(chan fluxRequest, len(requests))
	out := make(chan fluxProxyResponse, len(requests))
	done := make(chan int, workers)

	for _, request := range requests {
		inbox <- request
	}
	close(inbox)

	for i := 0; i < workers; i++ {
		go func() {
			processed := 0
			for request := range inbox {
				out <- proxy(request, timeoutMs)
				processed++
			}
			done <- processed
		}()
	}

	waitForWorkers(done, workers)
	return orderResponsesForRequests(drainResponses(out, len(requests)), requests)
}

func (r goNativeRuntime) stageScore(responses []fluxProxyResponse, retryCount int, backpressureEvents int) int {
	latency := make(chan int, len(responses))
	score := 0

	for _, response := range responses {
		step := -8
		if response.success {
			step = 5
		}
		select {
		case latency <- response.latencyMs:
			score += step
		default:
			score -= 20
		}
	}

	remaining := len(responses)
	for remaining > 0 {
		select {
		case value := <-latency:
			score += divFloorPositive(value, 20)
		default:
			score += 0
		}
		remaining--
	}

	close(latency)
	score += len(responses) * 4
	score -= backpressureEvents * 3
	score += retryCount * 5
	return score
}

func divFloorPositive(numerator int, denominator int) int {
	if denominator <= 0 {
		return 0
	}
	remaining := numerator
	quotient := 0
	for remaining >= denominator {
		remaining -= denominator
		quotient++
	}
	return quotient
}

func normalizeWorkers(workerCount int) int {
	if workerCount <= 0 {
		return 1
	}
	return workerCount
}

func waitForWorkers(done chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
}

func drainResponses(out chan fluxProxyResponse, expected int) []fluxProxyResponse {
	responses := make([]fluxProxyResponse, 0, expected)
	for i := 0; i < expected; i++ {
		responses = append(responses, <-out)
	}
	close(out)
	return responses
}

func normalizedRoute(route string) string {
	trimmed := strings.TrimSpace(route)
	if trimmed == "" {
		return "/unknown"
	}
	return trimmed
}

func upstreamForRoute(route string) string {
	if strings.HasPrefix(route, "/assets") {
		return "cdn"
	}
	if route == "/health" {
		return "healthz"
	}
	return "core-api"
}

func proxy(request fluxRequest, timeout int) fluxProxyResponse {
	route := normalizedRoute(request.route)
	latency := request.latencyMs
	if latency < 0 {
		latency = 0
	}
	status := request.status
	attempts := 1
	if status >= 500 {
		attempts = 2
	}
	if latency > timeout {
		status = 504
		attempts = 2
	}
	success := status < 500
	return fluxProxyResponse{
		requestID: request.id,
		route:     route,
		upstream:  upstreamForRoute(route),
		status:    status,
		latencyMs: latency,
		attempts:  attempts,
		success:   success,
	}
}

func rateLimited(request fluxRequest) fluxProxyResponse {
	route := normalizedRoute(request.route)
	latency := request.latencyMs
	if latency < 0 {
		latency = 0
	}
	return fluxProxyResponse{
		requestID: request.id,
		route:     route,
		upstream:  "rate-limit",
		status:    429,
		latencyMs: latency,
		attempts:  1,
		success:   false,
	}
}

func breakerOpen(request fluxRequest) fluxProxyResponse {
	route := normalizedRoute(request.route)
	return fluxProxyResponse{
		requestID: request.id,
		route:     route,
		upstream:  "breaker-open",
		status:    503,
		latencyMs: 0,
		attempts:  1,
		success:   false,
	}
}

func baselineRequests() []fluxRequest {
	return []fluxRequest{
		{id: 1, route: "/v1/items", latencyMs: 30, status: 200},
		{id: 2, route: "/v1/items", latencyMs: 70, status: 503},
		{id: 3, route: "/assets/logo.png", latencyMs: 12, status: 200},
		{id: 4, route: "/health", latencyMs: 4, status: 200},
		{id: 5, route: "/v1/auth", latencyMs: 40, status: 502},
		{id: 6, route: "/v1/items", latencyMs: 18, status: 200},
		{id: 7, route: "/assets/main.css", latencyMs: 9, status: 200},
		{id: 8, route: "/v1/auth", latencyMs: 28, status: 200},
	}
}

func cloneRequests(requests []fluxRequest) []fluxRequest {
	out := make([]fluxRequest, len(requests))
	copy(out, requests)
	return out
}

func ingest(requests []fluxRequest, capacity int) fluxIngestResult {
	boundedCapacity := capacity
	if boundedCapacity <= 0 {
		boundedCapacity = 1
	}
	queue := make([]fluxRequest, 0, len(requests))
	queueHead := 0
	accepted := make([]fluxRequest, 0, len(requests))
	backpressureEvents := 0

	for _, request := range requests {
		if len(queue)-queueHead >= boundedCapacity {
			backpressureEvents++
			accepted = append(accepted, queue[queueHead])
			queueHead++
		}
		queue = append(queue, request)
	}

	for queueHead < len(queue) {
		accepted = append(accepted, queue[queueHead])
		queueHead++
	}

	return fluxIngestResult{
		receivedCount:      len(requests),
		acceptedRequests:   accepted,
		backpressureEvents: backpressureEvents,
	}
}

type routePolicyPlan struct {
	dispatchable []fluxRequest
	synthetic    []fluxProxyResponse
	rateLimited  int
	breakerOpen  int
}

func applyRoutePolicies(requests []fluxRequest, limit int, breakerThreshold int, timeout int) routePolicyPlan {
	normalizedLimit := limit
	if normalizedLimit <= 0 {
		normalizedLimit = 1
	}
	normalizedBreaker := breakerThreshold
	if normalizedBreaker <= 0 {
		normalizedBreaker = 1
	}

	routeCounts := map[string]int{}
	failureStreak := map[string]int{}
	dispatchable := make([]fluxRequest, 0, len(requests))
	synthetic := make([]fluxProxyResponse, 0)
	rateLimitedCount := 0
	breakerOpenCount := 0

	for _, request := range requests {
		route := normalizedRoute(request.route)
		streak := failureStreak[route]
		if streak >= normalizedBreaker {
			synthetic = append(synthetic, breakerOpen(request))
			breakerOpenCount++
			continue
		}

		routeCount := routeCounts[route]
		if routeCount >= normalizedLimit {
			synthetic = append(synthetic, rateLimited(request))
			rateLimitedCount++
			continue
		}
		routeCounts[route] = routeCount + 1

		dispatchable = append(dispatchable, request)
		predictsFailure := request.status >= 500 || request.latencyMs > timeout
		if predictsFailure {
			failureStreak[route] = streak + 1
		} else {
			failureStreak[route] = 0
		}
	}

	return routePolicyPlan{
		dispatchable: dispatchable,
		synthetic:    synthetic,
		rateLimited:  rateLimitedCount,
		breakerOpen:  breakerOpenCount,
	}
}

func orderResponsesForRequests(items []fluxProxyResponse, requests []fluxRequest) []fluxProxyResponse {
	byID := make(map[int]fluxProxyResponse, len(items))
	for _, item := range items {
		byID[item.requestID] = item
	}
	ordered := make([]fluxProxyResponse, 0, len(requests))
	for _, request := range requests {
		if response, ok := byID[request.id]; ok {
			ordered = append(ordered, response)
		}
	}
	return ordered
}

func orderedResponses(synthetic []fluxProxyResponse, dispatched []fluxProxyResponse, accepted []fluxRequest) []fluxProxyResponse {
	combined := make([]fluxProxyResponse, 0, len(synthetic)+len(dispatched))
	combined = append(combined, synthetic...)
	combined = append(combined, dispatched...)
	return orderResponsesForRequests(combined, accepted)
}

func aggregate(responses []fluxProxyResponse) (routes []*fluxRouteAggregate, summary string) {
	byRoute := map[string]*fluxRouteAggregate{}
	routeKeys := make([]string, 0, len(responses))

	for _, response := range responses {
		route := response.route
		bucket, ok := byRoute[route]
		if !ok {
			bucket = newFluxRouteAggregate(route)
			byRoute[route] = bucket
			routeKeys = append(routeKeys, route)
		}
		bucket.record(response)
	}

	routes = make([]*fluxRouteAggregate, 0, len(routeKeys))
	tokens := make([]string, 0, len(routeKeys))
	for _, route := range routeKeys {
		item := byRoute[route]
		if item == nil {
			continue
		}
		routes = append(routes, item)
		tokens = append(tokens, item.summaryToken())
	}
	summary = strings.Join(tokens, ",")
	return routes, summary
}

func retries(responses []fluxProxyResponse) int {
	total := 0
	for _, response := range responses {
		total += response.attempts - 1
	}
	return total
}

func errors(responses []fluxProxyResponse) int {
	total := 0
	for _, response := range responses {
		if !response.success {
			total++
		}
	}
	return total
}

func runReport(runtime fluxRuntime, requests []fluxRequest) fluxReport {
	ingestResult := ingest(requests, ingestQueueCapacity)
	plan := applyRoutePolicies(ingestResult.acceptedRequests, perRouteLimit, breakerFailureThreshold, timeoutMs)
	dispatched := runtime.dispatch(plan.dispatchable, runtime.dispatchWorkers())
	responses := orderedResponses(plan.synthetic, dispatched, ingestResult.acceptedRequests)
	routeAggs, routeSummary := aggregate(responses)
	retryCount := retries(responses)
	errorCount := errors(responses)
	runtimeScore := runtime.stageScore(responses, retryCount, ingestResult.backpressureEvents)

	return fluxReport{
		profile:             runtime.profileID(),
		variant:             runtime.variantID(),
		capability:          runtime.capabilityID(),
		ingressReceived:     ingestResult.receivedCount,
		ingressAccepted:     len(ingestResult.acceptedRequests),
		ingressBackpressure: ingestResult.backpressureEvents,
		proxyResponses:      len(responses),
		proxyRetries:        retryCount,
		rateLimitedCount:    plan.rateLimited,
		breakerOpenCount:    plan.breakerOpen,
		routesCount:         len(routeAggs),
		routesSummary:       routeSummary,
		errorsCount:         errorCount,
		runtimeScore:        runtimeScore,
	}
}

func runScripted(runtime fluxRuntime) string {
	return runReport(runtime, cloneRequests(baselineRequests())).render()
}

func newRuntime(profile, variant string) (fluxRuntime, error) {
	switch variant {
	case "core":
		return coreRuntime{profile: profile}, nil
	case "go_native":
		return goNativeRuntime{profile: profile}, nil
	default:
		return nil, fmt.Errorf("unsupported variant: %s", variant)
	}
}

func runBreakerScenario(runtime fluxRuntime) fluxReport {
	requests := []fluxRequest{
		{id: 1, route: "/breaker/api", latencyMs: 60, status: 503},
		{id: 2, route: "/breaker/api", latencyMs: 65, status: 502},
		{id: 3, route: "/breaker/api", latencyMs: 20, status: 200},
		{id: 4, route: "/breaker/api", latencyMs: 18, status: 200},
	}
	return runReport(runtime, requests)
}
