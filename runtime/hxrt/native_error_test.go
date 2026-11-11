package hxrt

import (
	"errors"
	"fmt"
	"testing"
)

type observedNativeError struct{ calls *int }

func (e *observedNativeError) Error() string {
	*e.calls++
	return "observed"
}

type sliceNativeError []int

func (sliceNativeError) Error() string { return "slice" }

type nilSafeNativeError struct{}

func (*nilSafeNativeError) Error() string { return "nil-safe" }

type matchingNativeError struct{ target error }

func (e *matchingNativeError) Error() string        { panic("transport formatted error") }
func (e *matchingNativeError) Is(target error) bool { return target == e.target }
func (e *matchingNativeError) As(target any) bool {
	if out, ok := target.(**observedNativeError); ok {
		*out = e.target.(*observedNativeError)
		return true
	}
	return false
}

func TestNativeErrorTransportPreservesOriginal(t *testing.T) {
	if ImportNativeError(nil) != nil || ExportNativeError(nil) != nil {
		t.Fatal("nil interface did not survive")
	}
	calls := 0
	root := &observedNativeError{&calls}
	wrapped := fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", root))
	calls = 0 // Constructing fmt wrappers formats; transport must not.
	joined := errors.Join(wrapped, errors.New("other"))
	for _, original := range []error{root, wrapped, joined, &matchingNativeError{root}} {
		value := ImportNativeError(original)
		for i := 0; i < 3; i++ {
			if value == nil || ExportNativeError(value) != original {
				t.Fatal("transport changed the original root")
			}
			value = ImportNativeError(ExportNativeError(value))
		}
		var found *observedNativeError
		if !errors.Is(ExportNativeError(value), root) || !errors.As(ExportNativeError(value), &found) || found != root {
			t.Fatal("native matching changed")
		}
	}
	if calls != 0 {
		t.Fatal("transport called Error")
	}
	other := &observedNativeError{&calls}
	joinedOrder := errors.Join(root, other)
	var first *observedNativeError
	if !errors.As(ExportNativeError(ImportNativeError(joinedOrder)), &first) || first != root {
		t.Fatal("joined matching order changed")
	}
	left, right := errors.New("same"), errors.New("same")
	if errors.Is(ExportNativeError(ImportNativeError(left)), right) {
		t.Fatal("distinct equal-message errors collapsed")
	}
	nonComparable := sliceNativeError{1, 2}
	got := ExportNativeError(ImportNativeError(nonComparable)).(sliceNativeError)
	got[0] = 3
	if nonComparable[0] != 3 {
		t.Fatal("transport copied native value")
	}
}

func TestNativeErrorTypedNilAndTextObservation(t *testing.T) {
	var safe *nilSafeNativeError
	if value := ImportNativeError(safe); value == nil || *value.Text() != "nil-safe" || ExportNativeError(value) == nil {
		t.Fatal("nil-safe typed nil changed")
	}
	var pointer *observedNativeError
	var original error = pointer
	carrier := ImportNativeError(original)
	if carrier == nil || AnyEqualsNull(carrier) || !AnyEqualsNull(pointer) {
		t.Fatal("generic erasure lost the native interface distinction")
	}
	if exported := ExportNativeError(carrier); exported == nil || exported != original {
		t.Fatal("typed nil became a nil interface")
	}
	calls := 0
	value := ImportNativeError(&observedNativeError{&calls})
	if calls != 0 || *value.Text() != "observed" || calls != 1 {
		t.Fatal("text was not an explicit observation")
	}
	defer func() {
		if recover() == nil {
			t.Fatal("typed nil text should retain native panic")
		}
	}()
	carrier.Text()
}

func TestNativeErrorAliasPreservesFunctionSignature(t *testing.T) {
	var native func(error) error = func(value RawError) RawError { return value }
	root := errors.New("root")
	if native(root) != root {
		t.Fatal("raw alias changed native ABI")
	}
}
