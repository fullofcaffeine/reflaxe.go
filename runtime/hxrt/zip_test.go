package hxrt

import "testing"

func TestZipRoundTripsLevelsAndBufferSizes(t *testing.T) {
	values := []int{'r', 'e', 'p', 'e', 'a', 't', '-', 'r', 'e', 'p', 'e', 'a', 't', 0, 127, 128, 255}
	for _, level := range []int{-1, 0, 1, 6, 9} {
		compressed := ZipCompress(values, level)
		if len(compressed) == 0 {
			t.Fatalf("level %d produced no zlib stream", level)
		}
		for _, bufferSize := range []int{1, 7, 65536} {
			if got := ZipUncompress(compressed, false, bufferSize); !equalZipValues(got, values) {
				t.Fatalf("level %d buffer %d round trip = %#v", level, bufferSize, got)
			}
		}
	}
}

func TestZipRawDeflateRoundTrip(t *testing.T) {
	values := []int{'z', 'i', 'p', '-', 'e', 'n', 't', 'r', 'y'}
	zlibStream := ZipCompress(values, 6)
	if len(zlibStream) < 6 {
		t.Fatalf("zlib stream is too short: %d", len(zlibStream))
	}
	rawDeflate := append([]int(nil), zlibStream[2:len(zlibStream)-4]...)
	if got := ZipUncompress(rawDeflate, true, 2); !equalZipValues(got, values) {
		t.Fatalf("raw DEFLATE round trip = %#v", got)
	}
}

func TestZipInvalidInputsCrossExceptionCarrier(t *testing.T) {
	assertZipPanics(t, "compression level", func() {
		ZipCompress([]int{'x'}, 10)
	})
	assertZipPanics(t, "zlib stream", func() {
		ZipUncompress([]int{'n', 'o', 't', '-', 'z', 'l', 'i', 'b'}, false, 8)
	})
}

func TestZipDeflateHandleStreamsAcrossPartialBuffers(t *testing.T) {
	handle := ZipDeflateCreate(6)
	parts := [][]int{
		{'s', 't', 'r', 'e', 'a', 'm', '-'},
		{'s', 'a', 'f', 'e', '-'},
		{'z', 'i', 'p', '-', 'z', 'i', 'p', '-', 'z', 'i', 'p'},
	}
	var compressed []int
	calls := 0
	for index, part := range parts {
		flushMode := ZipFlushNo
		if index == 1 {
			flushMode = ZipFlushSync
		} else if index == len(parts)-1 {
			flushMode = ZipFlushFinish
		}
		position := 0
		done := false
		for position < len(part) || (index == len(parts)-1 && !done) {
			step := ZipDeflateExecute(handle, part[position:], 5, flushMode)
			calls++
			assertZipStep(t, "deflate", step, len(part)-position, 5)
			position += step.Read
			compressed = append(compressed, step.Values...)
			done = step.Done
			if !done && step.Read == 0 && len(step.Values) == 0 {
				t.Fatal("deflate made no progress")
			}
		}
	}
	if calls < 2 {
		t.Fatalf("deflate used %d calls; want repeated execution", calls)
	}
	if got := ZipUncompress(compressed, false, 3); !equalZipValues(got, flattenZipParts(parts)) {
		t.Fatalf("progressive deflate round trip = %#v", got)
	}
	ZipDeflateClose(handle)
	ZipDeflateClose(handle)
}

func TestZipInflateHandleStreamsAcrossPartialBuffers(t *testing.T) {
	want := []int{'s', 't', 'r', 'e', 'a', 'm', '-', 's', 'a', 'f', 'e', '-', 'z', 'i', 'p'}
	compressed := ZipCompress(want, 6)
	handle := ZipInflateCreate(false)
	var got []int
	compressedPosition := 0
	calls := 0
	done := false
	for !done {
		remaining := len(compressed) - compressedPosition
		partLength := 2
		if remaining < partLength {
			partLength = remaining
		}
		step := ZipInflateExecute(handle, compressed[compressedPosition:compressedPosition+partLength], 3, ZipFlushSync)
		calls++
		assertZipStep(t, "inflate", step, partLength, 3)
		compressedPosition += step.Read
		got = append(got, step.Values...)
		done = step.Done
		if !done && step.Read == 0 && len(step.Values) == 0 {
			t.Fatal("inflate made no progress")
		}
		if !done && compressedPosition == len(compressed) && len(step.Values) == 0 {
			t.Fatal("inflate exhausted input without completing")
		}
	}
	if calls < 2 {
		t.Fatalf("inflate used %d calls; want repeated execution", calls)
	}
	if compressedPosition != len(compressed) {
		t.Fatalf("inflate read %d of %d compressed bytes", compressedPosition, len(compressed))
	}
	if !equalZipValues(got, want) {
		t.Fatalf("progressive inflate = %#v", got)
	}
	ZipInflateClose(handle)
	ZipInflateClose(handle)
}

func TestZipInflateHandleSupportsRawDeflate(t *testing.T) {
	want := []int{'r', 'a', 'w', '-', 'd', 'e', 'f', 'l', 'a', 't', 'e'}
	zlibStream := ZipCompress(want, 6)
	rawDeflate := zlibStream[2 : len(zlibStream)-4]
	handle := ZipInflateCreate(true)
	var got []int
	position := 0
	done := false
	for !done {
		step := ZipInflateExecute(handle, rawDeflate[position:], 2, ZipFlushSync)
		assertZipStep(t, "raw inflate", step, len(rawDeflate)-position, 2)
		position += step.Read
		got = append(got, step.Values...)
		done = step.Done
		if !done && step.Read == 0 && len(step.Values) == 0 {
			t.Fatal("raw inflate made no progress")
		}
	}
	if !equalZipValues(got, want) {
		t.Fatalf("progressive raw inflate = %#v", got)
	}
	ZipInflateClose(handle)
}

func TestZipInflateHandleDoesNotConsumeTrailingBytes(t *testing.T) {
	want := []int{'o', 'n', 'e', '-', 's', 't', 'r', 'e', 'a', 'm'}
	compressed := ZipCompress(want, 6)
	input := append(append([]int(nil), compressed...), 'x', 'y')
	handle := ZipInflateCreate(false)
	step := ZipInflateExecute(handle, input, 256, ZipFlushSync)
	if !step.Done {
		t.Fatalf("inflate with trailing bytes was not done: %#v", step)
	}
	if step.Read != len(compressed) {
		t.Fatalf("inflate consumed %d bytes; want stream length %d", step.Read, len(compressed))
	}
	if !equalZipValues(step.Values, want) {
		t.Fatalf("inflate with trailing bytes = %#v", step.Values)
	}
	ZipInflateClose(handle)
}

func TestZipInflateHandlePreservesStateAcrossVariedFragments(t *testing.T) {
	want := make([]int, 8193)
	for index := range want {
		want[index] = (index*73 + index/7) & 0xff
	}
	compressed := ZipCompress(want, 6)
	handle := ZipInflateCreate(false)
	var got []int
	position := 0
	done := false
	calls := 0
	for !done {
		remaining := len(compressed) - position
		inputLength := calls%11 + 1
		if remaining < inputLength {
			inputLength = remaining
		}
		outputLimit := calls%19 + 1
		step := ZipInflateExecute(handle, compressed[position:position+inputLength], outputLimit, ZipFlushSync)
		assertZipStep(t, "varied inflate", step, inputLength, outputLimit)
		position += step.Read
		got = append(got, step.Values...)
		done = step.Done
		calls++
		if calls > len(compressed)+len(want)+100 {
			t.Fatal("varied inflate did not converge")
		}
		if !done && step.Read == 0 && len(step.Values) == 0 {
			t.Fatal("varied inflate made no progress")
		}
	}
	if position != len(compressed) {
		t.Fatalf("varied inflate read %d of %d bytes", position, len(compressed))
	}
	if !equalZipValues(got, want) {
		t.Fatalf("varied progressive inflate length = %d; want %d", len(got), len(want))
	}
	ZipInflateClose(handle)
}

func TestZipInflateHandleBoundsPendingExpansionByDestination(t *testing.T) {
	want := make([]int, 64*1024)
	compressed := ZipCompress(want, 9)
	handle := ZipInflateCreate(false)
	var got []int
	position := 0
	done := false
	const outputLimit = 257
	for !done {
		step := ZipInflateExecute(handle, compressed[position:], outputLimit, ZipFlushSync)
		position += step.Read
		got = append(got, step.Values...)
		handle.mu.Lock()
		buffered := handle.output.Len()
		retained := buffered + handle.outputBudget + handle.outputInFlight
		handle.mu.Unlock()
		if buffered > outputLimit {
			t.Fatalf("inflate retained %d decoded bytes behind capacity %d", buffered, outputLimit)
		}
		if retained > outputLimit {
			t.Fatalf("inflate retained %d bytes of output state behind capacity %d", retained, outputLimit)
		}
		done = step.Done
		if !done && step.Read == 0 && len(step.Values) == 0 {
			t.Fatal("bounded inflate made no progress")
		}
	}
	if position != len(compressed) {
		t.Fatalf("bounded inflate read %d of %d bytes", position, len(compressed))
	}
	if !equalZipValues(got, want) {
		t.Fatalf("bounded inflate length = %d; want %d", len(got), len(want))
	}
	ZipInflateClose(handle)
}

func TestZipCodecHandlesEnforceCapacityFlushAndLifecyclePolicy(t *testing.T) {
	deflate := ZipDeflateCreate(6)
	zero := ZipDeflateExecute(deflate, []int{'x'}, 0, ZipFlushNo)
	if zero.Read != 0 || len(zero.Values) != 0 || zero.Done {
		t.Fatalf("zero-capacity deflate = %#v", zero)
	}

	assertZipPanics(t, "FULL flush", func() {
		ZipDeflateExecute(deflate, nil, 8, ZipFlushFull)
	})
	assertZipPanics(t, "BLOCK flush", func() {
		ZipDeflateExecute(deflate, nil, 8, ZipFlushBlock)
	})
	ZipDeflateClose(deflate)
	ZipDeflateClose(deflate)
	assertZipPanics(t, "deflate use after close", func() {
		ZipDeflateExecute(deflate, []int{'x'}, 8, ZipFlushFinish)
	})

	inflate := ZipInflateCreate(false)
	zero = ZipInflateExecute(inflate, []int{1}, 0, ZipFlushSync)
	if zero.Read != 0 || len(zero.Values) != 0 || zero.Done {
		t.Fatalf("zero-capacity inflate = %#v", zero)
	}
	assertZipPanics(t, "inflate FULL flush", func() {
		ZipInflateExecute(inflate, nil, 8, ZipFlushFull)
	})
	assertZipPanics(t, "inflate BLOCK flush", func() {
		ZipInflateExecute(inflate, nil, 8, ZipFlushBlock)
	})
	ZipInflateClose(inflate)
	ZipInflateClose(inflate)
	assertZipPanics(t, "inflate use after close", func() {
		ZipInflateExecute(inflate, []int{1}, 8, ZipFlushSync)
	})
}

func assertZipStep(t *testing.T, name string, step *ZipCodecStep, availableInput int, outputLimit int) {
	t.Helper()
	if step == nil {
		t.Fatalf("%s returned a nil step", name)
	}
	if step.Read < 0 || step.Read > availableInput {
		t.Fatalf("%s read %d of %d available bytes", name, step.Read, availableInput)
	}
	if len(step.Values) > outputLimit {
		t.Fatalf("%s wrote %d bytes into capacity %d", name, len(step.Values), outputLimit)
	}
}

func flattenZipParts(parts [][]int) []int {
	var values []int
	for _, part := range parts {
		values = append(values, part...)
	}
	return values
}

func assertZipPanics(t *testing.T, name string, block func()) {
	t.Helper()
	deferred := false
	func() {
		defer func() {
			deferred = recover() != nil
		}()
		block()
	}()
	if !deferred {
		t.Fatalf("invalid %s did not cross the hxrt exception carrier", name)
	}
}

func equalZipValues(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
