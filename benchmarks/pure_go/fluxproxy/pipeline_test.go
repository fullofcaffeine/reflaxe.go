package main

import "testing"

const expectedCoreOutput = `fluxproxy.profile=pure_go
fluxproxy.variant=core
runtime.capability=loop_dispatch
ingress.received=8
ingress.accepted=8
ingress.backpressure=5
proxy.responses=8
proxy.retries=2
policy.rate_limited=1
policy.breaker_open=0
routes.count=5
routes.summary=/v1/items:3/1/2/39,/assets/logo.png:1/1/0/12,/health:1/1/0/4,/v1/auth:2/1/1/34,/assets/main.css:1/1/0/9
errors.count=3
runtime.score=26`

const expectedGoNativeOutput = `fluxproxy.profile=pure_go
fluxproxy.variant=go_native
runtime.capability=worker_chan_fanout
ingress.received=8
ingress.accepted=8
ingress.backpressure=5
proxy.responses=8
proxy.retries=2
policy.rate_limited=1
policy.breaker_open=0
routes.count=5
routes.summary=/v1/items:3/1/2/39,/assets/logo.png:1/1/0/12,/health:1/1/0/4,/v1/auth:2/1/1/34,/assets/main.css:1/1/0/9
errors.count=3
runtime.score=35`

func TestScriptedCoreOutput(t *testing.T) {
	runtime, err := newRuntime("pure_go", "core")
	if err != nil {
		t.Fatalf("newRuntime(core): %v", err)
	}
	got := runScripted(runtime)
	if got != expectedCoreOutput {
		t.Fatalf("core output mismatch\nexpected:\n%s\n\ngot:\n%s", expectedCoreOutput, got)
	}
}

func TestScriptedGoNativeOutput(t *testing.T) {
	runtime, err := newRuntime("pure_go", "go_native")
	if err != nil {
		t.Fatalf("newRuntime(go_native): %v", err)
	}
	got := runScripted(runtime)
	if got != expectedGoNativeOutput {
		t.Fatalf("go_native output mismatch\nexpected:\n%s\n\ngot:\n%s", expectedGoNativeOutput, got)
	}
}

func TestBreakerScenario(t *testing.T) {
	runtime, err := newRuntime("pure_go", "core")
	if err != nil {
		t.Fatalf("newRuntime(core): %v", err)
	}
	report := runBreakerScenario(runtime)
	if report.breakerOpenCount != 2 {
		t.Fatalf("expected breakerOpen=2, got %d", report.breakerOpenCount)
	}
	if report.rateLimitedCount != 0 {
		t.Fatalf("expected rateLimited=0, got %d", report.rateLimitedCount)
	}
	if report.proxyRetries != 2 {
		t.Fatalf("expected retries=2, got %d", report.proxyRetries)
	}
}

func BenchmarkFluxProxyCore(b *testing.B) {
	runtime, err := newRuntime("pure_go", "core")
	if err != nil {
		b.Fatalf("newRuntime(core): %v", err)
	}
	requests := baselineRequests()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		report := runReport(runtime, cloneRequests(requests))
		if report.runtimeScore == 0 {
			b.Fatalf("unexpected score")
		}
	}
}

func BenchmarkFluxProxyGoNative(b *testing.B) {
	runtime, err := newRuntime("pure_go", "go_native")
	if err != nil {
		b.Fatalf("newRuntime(go_native): %v", err)
	}
	requests := baselineRequests()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		report := runReport(runtime, cloneRequests(requests))
		if report.runtimeScore == 0 {
			b.Fatalf("unexpected score")
		}
	}
}
