package hxrt

import "testing"

func TestTryCatchCatchesOnlyHaxeExceptions(t *testing.T) {
	marker := &struct{ name string }{name: "haxe"}
	var caught any
	TryCatch(func() {
		Throw(marker)
	}, func(value any) {
		caught = value
	})
	if caught != marker {
		t.Fatalf("caught value = %v, want explicit Haxe throw carrier", caught)
	}
}

func TestTryCatchRepanicsForeignPanics(t *testing.T) {
	marker := &struct{ name string }{name: "native"}
	catchCalled := false
	var recovered any
	func() {
		defer func() {
			recovered = recover()
		}()
		TryCatch(func() {
			panic(marker)
		}, func(any) {
			catchCalled = true
		})
	}()

	if catchCalled {
		t.Fatal("Haxe catch handled a foreign Go panic")
	}
	if recovered != marker {
		t.Fatalf("recovered panic = %v, want original foreign panic", recovered)
	}
}

func TestPortableRuntimeFailuresUseHaxeExceptionCarrier(t *testing.T) {
	for _, probe := range []struct {
		name      string
		value     any
		wantError string
	}{
		{name: "null", value: nil, wantError: "Invalid operation: null"},
		{name: "non-int", value: struct{}{}, wantError: "Invalid operation: expected Int-compatible value"},
	} {
		t.Run(probe.name, func(t *testing.T) {
			var caught any
			TryCatch(func() {
				IntFromNullableAny(probe.value)
			}, func(value any) {
				caught = value
			})
			if caught == nil {
				t.Fatal("portable runtime failure escaped the Haxe exception boundary")
			}
			if got := *StdString(caught); got != probe.wantError {
				t.Fatalf("caught runtime failure = %q, want %q", got, probe.wantError)
			}
		})
	}
}
