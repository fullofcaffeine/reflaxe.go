package hxrt

import "testing"

var benchmarkFunctionSink *Function[func(int) int]
var benchmarkIntSink int
var benchmarkBoolSink bool

// These benchmarks isolate carrier costs. They do not establish application
// parity: generated-program and pinned-Beads workloads remain separate gates.
func BenchmarkFunctionInvocation(b *testing.B) {
	call := func(value int) int { return value + 1 }
	carrier := NewFunction(call)
	b.Run("native", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			benchmarkIntSink = call(7)
		}
	})
	b.Run("carrier", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			benchmarkIntSink = FunctionCall(carrier)(7)
		}
	})
}

func BenchmarkFunctionEscapingCreation(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value := 7
		benchmarkFunctionSink = NewFunction(func(delta int) int { value += delta; return value })
	}
}

func BenchmarkFunctionAlias(b *testing.B) {
	value := NewFunction(func(delta int) int { return delta + 7 })
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkFunctionSink = value
	}
}

func BenchmarkFunctionBoundExtraction(b *testing.B) {
	receiver := &struct{ value int }{7}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkFunctionSink = BoundFunction(func(delta int) int { return receiver.value + delta }, func() any { return receiver }, "step")
	}
}

func BenchmarkFunctionAdaptation(b *testing.B) {
	source := NewFunction(func(value int64) int64 { return value + 1 })
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkFunctionSink = AdaptFunction(source, func(value int) int { return int(FunctionCall(source)(int64(value))) })
	}
}

// A depth series makes accidental wrapper-chain growth visible. Construction is
// outside the timed loop; reads and identity observation measure traversal only.
func BenchmarkAnonymousViewDepth(b *testing.B) {
	for _, size := range []struct {
		name  string
		depth int
	}{{"1", 1}, {"8", 8}, {"64", 64}} {
		raw := map[string]any{"value": 7}
		var view Anonymous = AnonymousFromMap(raw)
		for i := 0; i < size.depth; i++ {
			previous := view
			view = NewAnonymousView(previous, func(name string) AnonymousField { return AnonymousLookup(previous, name) }, nil)
		}
		if AnonymousGet(view, "value") != 7 || !HaxeEqual(view, raw) {
			b.Fatal("view contract failed")
		}
		b.Run(size.name+"/read", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				benchmarkIntSink = AnonymousGet(view, "value").(int)
			}
		})
		b.Run(size.name+"/identity", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				benchmarkBoolSink = HaxeEqual(view, raw)
			}
		})
	}
}
