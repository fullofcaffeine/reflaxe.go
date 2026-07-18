package hxrt

import "testing"

func TestTemplateArrayValues(t *testing.T) {
	want := []any{1, 2, 3}
	if got := TemplateArrayValues([]int{1, 2, 3}); len(got) != len(want) || got[0] != want[0] || got[1] != want[1] || got[2] != want[2] {
		t.Fatalf("TemplateArrayValues(slice) = %#v, want %#v", got, want)
	}
	if got := TemplateArrayValues(NewArray("portable", 7)); len(got) != 2 || got[0] != "portable" || got[1] != 7 {
		t.Fatalf("TemplateArrayValues(portable Array) = %#v, want [portable 7]", got)
	}

	array := [2]string{"a", "b"}
	if got := TemplateArrayValues(&array); len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("TemplateArrayValues(pointer-to-array) = %#v, want [a b]", got)
	}

	if got := TemplateArrayValues("not-an-array"); got != nil {
		t.Fatalf("TemplateArrayValues(non-array) = %#v, want nil", got)
	}
	if got := TemplateArrayValues(nil); got != nil {
		t.Fatalf("TemplateArrayValues(nil) = %#v, want nil", got)
	}
}

func TestTemplateIsObject(t *testing.T) {
	type record struct{ Value int }
	value := record{Value: 1}

	for _, test := range []struct {
		name  string
		value any
		want  bool
	}{
		{name: "map", value: map[string]any{"value": 1}, want: true},
		{name: "struct", value: value, want: true},
		{name: "pointer", value: &value, want: true},
		{name: "slice", value: []int{1}, want: false},
		{name: "number", value: 1, want: false},
		{name: "nil", value: nil, want: false},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := TemplateIsObject(test.value); got != test.want {
				t.Fatalf("TemplateIsObject(%s) = %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestTemplateCall(t *testing.T) {
	join := func(prefix string, value int) string {
		return prefix + string(rune('0'+value))
	}
	if got := TemplateCall(join, []any{"v", 3}); got != "v3" {
		t.Fatalf("TemplateCall(join) = %#v, want v3", got)
	}

	called := false
	if got := TemplateCall(func() { called = true }, nil); got != nil {
		t.Fatalf("TemplateCall(void) = %#v, want nil", got)
	}
	if !called {
		t.Fatal("TemplateCall(void) did not invoke the function")
	}

	if got := TemplateCall(nil, nil); got != nil {
		t.Fatalf("TemplateCall(nil) = %#v, want nil", got)
	}
	if got := TemplateCall("not-a-function", nil); got != nil {
		t.Fatalf("TemplateCall(non-function) = %#v, want nil", got)
	}
}

type templateCallCounter struct {
	value int
}

func (counter *templateCallCounter) add(amount int) int {
	counter.value += amount
	return counter.value
}

func TestTemplateCallInvokesBoundMethodValue(t *testing.T) {
	counter := &templateCallCounter{value: 4}

	got := TemplateCall(counter.add, []any{3})

	if got != 7 {
		t.Fatalf("expected bound method result 7, got %v", got)
	}
	if counter.value != 7 {
		t.Fatalf("expected bound receiver mutation to persist, got %d", counter.value)
	}
}
