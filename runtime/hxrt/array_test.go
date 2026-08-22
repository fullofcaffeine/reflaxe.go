package hxrt

import (
	"strings"
	"testing"
)

func TestArraySharesLengthChangingMutations(t *testing.T) {
	array := NewArray(1, 2)
	alias := array

	if got := alias.Push(3); got != 3 {
		t.Fatalf("Push() length = %d, want 3", got)
	}
	alias.Insert(1, 9)
	alias.RemoveAt(2)
	if got := alias.Pop(); got != 3 {
		t.Fatalf("Pop() = %#v, want 3", got)
	}

	if array.Len() != 2 || array.Get(0) != 1 || array.Get(1) != 9 {
		t.Fatalf("shared array = %#v, want [1 9]", array.Values())
	}
}

func TestArraySparseSetPreservesNullHoles(t *testing.T) {
	array := NewArray()
	if got := array.Set(2, 7); got != 7 {
		t.Fatalf("Set() = %#v, want 7", got)
	}

	if array.Len() != 3 {
		t.Fatalf("Len() = %d, want 3", array.Len())
	}
	if array.Get(0) != nil || array.Get(1) != nil || array.Get(2) != 7 {
		t.Fatalf("sparse array = %#v, want [nil nil 7]", array.Values())
	}
	if array.Get(-1) != nil || array.Get(3) != nil {
		t.Fatalf("out-of-range reads must return nil")
	}
}

func TestArrayPositionAndLengthRules(t *testing.T) {
	array := NewArray(2, 3)
	array.Insert(-99, 1)
	array.Insert(99, 4)
	array.SetLength(6)

	want := []any{1, 2, 3, 4, nil, nil}
	got := array.Values()
	if len(got) != len(want) {
		t.Fatalf("Values() length = %d, want %d", len(got), len(want))
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("Values()[%d] = %#v, want %#v", index, got[index], want[index])
		}
	}

	array.SetLength(2)
	if array.Len() != 2 || array.Get(0) != 1 || array.Get(1) != 2 {
		t.Fatalf("truncated array = %#v, want [1 2]", array.Values())
	}
}

func TestArrayCopyHasDistinctIdentity(t *testing.T) {
	array := NewArray("a")
	copy := array.Copy()
	if copy == array {
		t.Fatalf("Copy() retained source identity")
	}
	copy.Push("b")
	if array.Len() != 1 || copy.Len() != 2 {
		t.Fatalf("copy mutation leaked: source=%#v copy=%#v", array.Values(), copy.Values())
	}
}

func TestArraySortMutatesSharedIdentity(t *testing.T) {
	array := NewArray("beta", "alpha", "gamma")
	alias := array
	ArraySort(array, func(left, right any) int {
		return strings.Compare(left.(string), right.(string))
	})
	if got, want := alias.Values(), []any{"alpha", "beta", "gamma"}; !equalArrayValues(got, want) {
		t.Fatalf("Sort() through alias = %#v, want %#v", got, want)
	}
}

func TestArraySliceUsesHaxeBoundsAndCopies(t *testing.T) {
	array := NewArray(0, 1, 2, 3)
	tests := []struct {
		name  string
		start int
		end   int
		open  bool
		want  []any
	}{
		{name: "start only", start: 1, open: true, want: []any{1, 2, 3}},
		{name: "bounded", start: 1, end: 3, want: []any{1, 2}},
		{name: "negative start", start: -2, open: true, want: []any{2, 3}},
		{name: "negative end", start: 0, end: -1, want: []any{0, 1, 2}},
		{name: "past end", start: 7, open: true, want: []any{}},
		{name: "reversed", start: 3, end: 1, want: []any{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got *Array
			if test.open {
				got = array.SliceFrom(test.start)
			} else {
				got = array.Slice(test.start, test.end)
			}
			if !equalArrayValues(got.Values(), test.want) {
				t.Fatalf("Slice(%d, %v) = %#v, want %#v", test.start, test.end, got.Values(), test.want)
			}
		})
	}

	copy := array.Slice(0, 2)
	copy.Set(0, 9)
	if got := array.Get(0); got != 0 {
		t.Fatalf("mutating slice copy changed source to %v", got)
	}

	native := []int{0, 1, 2, 3}
	nativeCopy := SliceValues(native, 1, 3)
	nativeCopy[0] = 9
	if native[1] != 1 {
		t.Fatalf("mutating native slice copy changed source to %v", native)
	}
	if got := SliceValuesFrom(native, 2); len(got) != 2 || got[0] != 2 || got[1] != 3 {
		t.Fatalf("SliceValuesFrom() = %#v, want [2 3]", got)
	}
	if got := array.SliceOptional(1, nil); !equalArrayValues(got.Values(), []any{1, 2, 3}) {
		t.Fatalf("SliceOptional(1, nil) = %#v, want [1 2 3]", got.Values())
	}
	if got := array.SliceOptional(1, 3); !equalArrayValues(got.Values(), []any{1, 2}) {
		t.Fatalf("SliceOptional(1, 3) = %#v, want [1 2]", got.Values())
	}
	if got := SliceValuesOptional(native, 1, nil); len(got) != 3 || got[0] != 1 || got[2] != 3 {
		t.Fatalf("SliceValuesOptional(1, nil) = %#v, want [1 2 3]", got)
	}
}

func equalArrayValues(left, right []any) bool {
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

func TestArrayValuesCopyDetachesNativeStorage(t *testing.T) {
	array := NewArray("source")
	values := array.ValuesCopy()
	values[0] = "native"
	if got := array.Get(0); got != "source" {
		t.Fatalf("ValuesCopy mutation changed portable Array to %#v", got)
	}
}
