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
