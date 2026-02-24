package main

import (
	"strings"
	"testing"
)

const expectedCoreOutput = `pulseforge.profile=pure_go
pulseforge.variant=core
runtime.capability=core_loop
ingest.received=8
ingest.accepted=8
ingest.backpressure=5
parse.events=8
enrich.events=8
aggregate.sources=5
aggregate.total=54
aggregate.weighted_total=100
aggregate.summary=edge:2/5/5/sev1,api:2/16/25/sev2,db:2/15/26/sev2,auth:1/13/39/sev3,worker:1/5/5/sev1
alert.count=2
alert.events=3,5
runtime.score=108`

const expectedGoNativeOutput = `pulseforge.profile=pure_go
pulseforge.variant=go_native
runtime.capability=chan_fanout_select
ingest.received=8
ingest.accepted=8
ingest.backpressure=5
parse.events=8
enrich.events=8
aggregate.sources=5
aggregate.total=54
aggregate.weighted_total=100
aggregate.summary=edge:2/5/5/sev1,api:2/16/25/sev2,db:2/15/26/sev2,auth:1/13/39/sev3,worker:1/5/5/sev1
alert.count=2
alert.events=3,5
runtime.score=123`

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

func TestSharedDomainSemanticsAcrossVariants(t *testing.T) {
	coreRuntime, err := newRuntime("pure_go", "core")
	if err != nil {
		t.Fatalf("newRuntime(core): %v", err)
	}
	nativeRuntime, err := newRuntime("pure_go", "go_native")
	if err != nil {
		t.Fatalf("newRuntime(go_native): %v", err)
	}

	coreReport := runReport(coreRuntime, cloneFrames(baselineFrames()))
	nativeReport := runReport(nativeRuntime, cloneFrames(baselineFrames()))

	assertFieldEq := func(name string, a string, b string) {
		t.Helper()
		if a != b {
			t.Fatalf("%s mismatch: %q != %q", name, a, b)
		}
	}

	assertFieldEq("ingest.received", keyValue(coreReport.render(), "ingest.received"), keyValue(nativeReport.render(), "ingest.received"))
	assertFieldEq("ingest.accepted", keyValue(coreReport.render(), "ingest.accepted"), keyValue(nativeReport.render(), "ingest.accepted"))
	assertFieldEq("ingest.backpressure", keyValue(coreReport.render(), "ingest.backpressure"), keyValue(nativeReport.render(), "ingest.backpressure"))
	assertFieldEq("parse.events", keyValue(coreReport.render(), "parse.events"), keyValue(nativeReport.render(), "parse.events"))
	assertFieldEq("enrich.events", keyValue(coreReport.render(), "enrich.events"), keyValue(nativeReport.render(), "enrich.events"))
	assertFieldEq("aggregate.summary", keyValue(coreReport.render(), "aggregate.summary"), keyValue(nativeReport.render(), "aggregate.summary"))
	assertFieldEq("alert.events", keyValue(coreReport.render(), "alert.events"), keyValue(nativeReport.render(), "alert.events"))
}

func keyValue(rendered string, key string) string {
	lines := strings.Split(rendered, "\n")
	prefix := key + "="
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			return strings.TrimPrefix(line, prefix)
		}
	}
	return ""
}

func BenchmarkPulseForgeCore(b *testing.B) {
	runtime, err := newRuntime("pure_go", "core")
	if err != nil {
		b.Fatalf("newRuntime(core): %v", err)
	}
	frames := baselineFrames()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		report := runReport(runtime, cloneFrames(frames))
		if report.RuntimeScore == 0 {
			b.Fatalf("unexpected score")
		}
	}
}

func BenchmarkPulseForgeGoNative(b *testing.B) {
	runtime, err := newRuntime("pure_go", "go_native")
	if err != nil {
		b.Fatalf("newRuntime(go_native): %v", err)
	}
	frames := baselineFrames()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		report := runReport(runtime, cloneFrames(frames))
		if report.RuntimeScore == 0 {
			b.Fatalf("unexpected score")
		}
	}
}
