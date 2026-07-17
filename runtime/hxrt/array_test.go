package hxrt

import "testing"

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

func TestArrayValuesCopyDetachesNativeStorage(t *testing.T) {
	array := NewArray("source")
	values := array.ValuesCopy()
	values[0] = "native"
	if got := array.Get(0); got != "source" {
		t.Fatalf("ValuesCopy mutation changed portable Array to %#v", got)
	}
}
