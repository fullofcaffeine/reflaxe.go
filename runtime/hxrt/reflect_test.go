package hxrt

import "testing"

type reflectRuntimeFixture struct {
	Name  string
	Count int
}

func (fixture *reflectRuntimeFixture) Add(value int) int {
	return fixture.Count + value
}

func reflectTestName(value string) *string {
	return StringFromLiteral(value)
}

func TestReflectRuntimeFieldAndMethodLookup(t *testing.T) {
	fixture := &reflectRuntimeFixture{Name: "Ada", Count: 4}
	field := ReflectLookupField(fixture, reflectTestName("Name"))
	if field == nil || !field.Found || field.Value != "Ada" {
		t.Fatalf("unexpected field lookup: %#v", field)
	}
	missing := ReflectLookupField(fixture, reflectTestName("Missing"))
	if missing == nil || missing.Found {
		t.Fatalf("unexpected missing lookup: %#v", missing)
	}
	method := ReflectLookupMethod(fixture, reflectTestName("Add"))
	if method == nil || !method.Found {
		t.Fatalf("unexpected method lookup: %#v", method)
	}
	if got := ReflectCallMethod(method.Value, []any{3}); got != 7 {
		t.Fatalf("unexpected method result: %#v", got)
	}
}

func TestReflectRuntimeHidesStdStringProtocolMethods(t *testing.T) {
	if lookup := ReflectLookupMethod(&stdStringTestClass{}, reflectTestName("HxClassName")); lookup == nil || lookup.Found {
		t.Fatalf("HxClassName leaked as a Haxe member: %#v", lookup)
	}
	if lookup := ReflectLookupMethod(&stdStringTestEnum{}, reflectTestName("HxEnumValue")); lookup == nil || lookup.Found {
		t.Fatalf("HxEnumValue leaked as a Haxe member: %#v", lookup)
	}
}

func TestReflectRuntimeAnonymousObjectOperations(t *testing.T) {
	object := map[string]any{"name": "Ada", "count": 2}
	ReflectSetField(object, reflectTestName("name"), "Bea")
	if got := object["name"]; got != "Bea" {
		t.Fatalf("set field = %#v", got)
	}
	if !ReflectDeleteField(object, reflectTestName("count")) {
		t.Fatal("expected count deletion")
	}
	if _, found := object["count"]; found {
		t.Fatal("count remained after deletion")
	}
	copied, ok := ReflectCopy(object).(map[string]any)
	if !ok {
		t.Fatalf("copy type = %T", ReflectCopy(object))
	}
	copied["name"] = "Cleo"
	if object["name"] != "Bea" {
		t.Fatalf("copy mutated source: %#v", object)
	}
	fields := ReflectFields(object)
	if len(fields) != 1 || *fields[0] != "name" {
		t.Fatalf("fields = %#v", fields)
	}
}

func TestReflectRuntimeClassificationComparisonAndVarArgs(t *testing.T) {
	if ReflectCompare(2, 3) >= 0 || ReflectCompare(reflectTestName("z"), reflectTestName("a")) <= 0 {
		t.Fatal("unexpected comparison ordering")
	}
	callback := func(arguments *Array) any {
		total := 0
		for _, argument := range arguments.Values() {
			total += argument.(int)
		}
		return total
	}
	variadic, ok := ReflectMakeVarArgs(callback).(*Function[func(...any) any])
	if !ok {
		t.Fatalf("variadic type = %T", ReflectMakeVarArgs(callback))
	}
	if got := FunctionCall(variadic)(1, 2, 3, 4); got != 10 {
		t.Fatalf("variadic result = %#v", got)
	}
	if !ReflectIsFunction(callback) || !ReflectIsObject(map[string]any{}) || !ReflectIsObject(reflectTestName("text")) || ReflectIsObject("text") {
		t.Fatal("unexpected runtime classification")
	}
}

// reflectBuiltinCall reads one built-in member and calls it as Reflect.callMethod
// does, failing the test when the member is missing.
func reflectBuiltinCall(t *testing.T, receiver any, name string, arguments ...any) any {
	t.Helper()
	lookup := ReflectLookupBuiltin(receiver, reflectTestName(name))
	if lookup == nil || !lookup.Found || !ReflectIsFunction(lookup.Value) {
		t.Fatalf("built-in %s missing: %#v", name, lookup)
	}
	return ReflectCallMethod(lookup.Value, arguments)
}

func reflectArrayText(value any) string {
	return *StdString(value)
}

func TestReflectBuiltinArrayMembers(t *testing.T) {
	array := NewArray(3, 1, 2, 1)
	if lookup := ReflectLookupBuiltin(array, reflectTestName("length")); lookup == nil || !lookup.Found || lookup.Value != 4 {
		t.Fatalf("length = %#v", lookup)
	}
	cases := []struct {
		name      string
		arguments []any
		want      string
	}{
		{"indexOf", []any{1}, "1"},
		{"indexOf", []any{1, 2}, "3"},
		{"indexOf", []any{1, -1}, "3"},
		{"indexOf", []any{1, nil}, "1"},
		{"lastIndexOf", []any{1}, "3"},
		{"lastIndexOf", []any{1, 2}, "1"},
		{"lastIndexOf", []any{1, nil}, "3"},
		{"contains", []any{2}, "true"},
		{"contains", []any{9}, "false"},
		{"join", []any{reflectTestName("|")}, "3|1|2|1"},
		{"slice", []any{1}, "[1,2,1]"},
		{"slice", []any{1, 2}, "[1]"},
		{"slice", []any{1, nil}, "[1,2,1]"},
		{"copy", nil, "[3,1,2,1]"},
		{"concat", []any{NewArray(9)}, "[3,1,2,1,9]"},
		{"toString", nil, "[3,1,2,1]"},
	}
	for _, test := range cases {
		if got := reflectArrayText(reflectBuiltinCall(t, array, test.name, test.arguments...)); got != test.want {
			t.Fatalf("%s%v = %s, want %s", test.name, test.arguments, got, test.want)
		}
	}
	double := NewFunction(func(value int) int { return value * 10 })
	if got := reflectArrayText(reflectBuiltinCall(t, array, "map", double)); got != "[30,10,20,10]" {
		t.Fatalf("map = %s", got)
	}
	keep := NewFunction(func(value int) bool { return value > 1 })
	if got := reflectArrayText(reflectBuiltinCall(t, array, "filter", keep)); got != "[3,2]" {
		t.Fatalf("filter = %s", got)
	}
	if got := reflectBuiltinCall(t, array, "remove", 1); got != true || reflectArrayText(array) != "[3,2,1]" {
		t.Fatalf("remove = %#v, array %s", got, reflectArrayText(array))
	}
	if got := reflectBuiltinCall(t, array, "remove", 9); got != false {
		t.Fatalf("remove missing = %#v", got)
	}
	if got := reflectBuiltinCall(t, array, "shift"); got != 3 {
		t.Fatalf("shift = %#v", got)
	}
	reflectBuiltinCall(t, array, "unshift", 0)
	if got := reflectBuiltinCall(t, array, "push", 7); got != 4 {
		t.Fatalf("push = %#v", got)
	}
	ascending := NewFunction(func(left int, right int) int { return left - right })
	reflectBuiltinCall(t, array, "sort", ascending)
	if got := reflectArrayText(array); got != "[0,1,2,7]" {
		t.Fatalf("sort = %s", got)
	}
	reflectBuiltinCall(t, array, "insert", 1, 5)
	reflectBuiltinCall(t, array, "reverse")
	if got := reflectArrayText(array); got != "[7,2,1,5,0]" {
		t.Fatalf("insert and reverse = %s", got)
	}
	if got := reflectArrayText(reflectBuiltinCall(t, array, "splice", 1, 2)); got != "[2,1]" {
		t.Fatalf("splice = %s", got)
	}
	if got := reflectBuiltinCall(t, array, "pop"); got != 0 {
		t.Fatalf("pop = %#v", got)
	}
	reflectBuiltinCall(t, array, "resize", 4)
	if got := reflectArrayText(array); got != "[7,5,null,null]" {
		t.Fatalf("resize = %s", got)
	}
	empty := NewArray()
	if got := reflectBuiltinCall(t, empty, "pop"); got != nil {
		t.Fatalf("empty pop = %#v", got)
	}
	if got := reflectBuiltinCall(t, empty, "shift"); got != nil {
		t.Fatalf("empty shift = %#v", got)
	}
}

func TestReflectBuiltinStringMembers(t *testing.T) {
	text := reflectTestName("x,y,z")
	if lookup := ReflectLookupBuiltin(text, reflectTestName("length")); lookup == nil || !lookup.Found || lookup.Value != 5 {
		t.Fatalf("length = %#v", lookup)
	}
	cases := []struct {
		name      string
		arguments []any
		want      string
	}{
		{"charAt", []any{2}, "y"},
		{"charAt", []any{10}, ""},
		{"charCodeAt", []any{0}, "120"},
		{"charCodeAt", []any{10}, "null"},
		{"indexOf", []any{reflectTestName(",")}, "1"},
		{"indexOf", []any{reflectTestName(","), 2}, "3"},
		{"indexOf", []any{reflectTestName(","), nil}, "1"},
		{"lastIndexOf", []any{reflectTestName(",")}, "3"},
		{"lastIndexOf", []any{reflectTestName(","), 2}, "1"},
		{"split", []any{reflectTestName(",")}, "[x,y,z]"},
		{"substr", []any{2}, "y,z"},
		{"substr", []any{2, 1}, "y"},
		{"substr", []any{2, nil}, "y,z"},
		{"substring", []any{2}, "y,z"},
		{"substring", []any{2, 3}, "y"},
		{"substring", []any{2, nil}, "y,z"},
		{"toUpperCase", nil, "X,Y,Z"},
		{"toLowerCase", nil, "x,y,z"},
		{"toString", nil, "x,y,z"},
	}
	for _, test := range cases {
		if got := reflectArrayText(reflectBuiltinCall(t, text, test.name, test.arguments...)); got != test.want {
			t.Fatalf("%s%v = %s, want %s", test.name, test.arguments, got, test.want)
		}
	}
	if parts, ok := reflectBuiltinCall(t, text, "split", reflectTestName(",")).(*Array); !ok || parts.Len() != 3 {
		t.Fatalf("split must return a portable Array: %#v", parts)
	}
}

func TestReflectBuiltinMembersStayBoundAndHidden(t *testing.T) {
	array := NewArray(1)
	alias := array
	push := ReflectLookupBuiltin(array, reflectTestName("push")).Value
	ReflectCallMethod(push, []any{2})
	if got := reflectArrayText(alias); got != "[1,2]" {
		t.Fatalf("bound push = %s", got)
	}
	again := ReflectLookupBuiltin(array, reflectTestName("push")).Value
	other := ReflectLookupBuiltin(NewArray(1), reflectTestName("push")).Value
	pop := ReflectLookupBuiltin(array, reflectTestName("pop")).Value
	if !ReflectCompareMethods(push, again) || ReflectCompareMethods(push, other) || ReflectCompareMethods(push, pop) {
		t.Fatal("bound built-in methods must compare by receiver and member")
	}
	if HaxeEqual(push, again) {
		t.Fatal("each read of a bound built-in method is a fresh function value")
	}
	for _, name := range []string{"Len", "Push", "Values", "iterator", "missing"} {
		if lookup := ReflectLookupBuiltin(array, reflectTestName(name)); lookup == nil || lookup.Found {
			t.Fatalf("Array member %s must be absent: %#v", name, lookup)
		}
		if lookup := ReflectLookupMethod(array, reflectTestName(name)); lookup == nil || lookup.Found {
			t.Fatalf("Array carrier method %s leaked as a Haxe member: %#v", name, lookup)
		}
	}
	// Staged Reflect.hasField reads only these two probes; --interp reports
	// no field for the members of an Array or a String.
	for _, receiver := range []any{array, reflectTestName("text")} {
		for _, name := range []string{"length", "push", "charAt"} {
			if ReflectLookupField(receiver, reflectTestName(name)).Found || ReflectLookupMethod(receiver, reflectTestName(name)).Found {
				t.Fatalf("%T member %s must not read as an object field", receiver, name)
			}
		}
	}
	if lookup := ReflectLookupBuiltin(map[string]any{"length": 1}, reflectTestName("length")); lookup == nil || lookup.Found {
		t.Fatalf("anonymous objects have no built-in members: %#v", lookup)
	}
	var missing *Array
	if lookup := ReflectLookupBuiltin(missing, reflectTestName("length")); lookup == nil || lookup.Found {
		t.Fatalf("a null Array has no members: %#v", lookup)
	}
}

func TestReflectBuiltinLengthWriteAndFields(t *testing.T) {
	array := NewArray(5, 6, 7)
	if !ReflectSetField(array, reflectTestName("length"), 1) || reflectArrayText(array) != "[5]" {
		t.Fatalf("length write must truncate: %s", reflectArrayText(array))
	}
	if !ReflectSetField(array, reflectTestName("length"), 3) || reflectArrayText(array) != "[5,null,null]" {
		t.Fatalf("length write must extend with null: %s", reflectArrayText(array))
	}
	if ReflectSetField(array, reflectTestName("other"), 1) || ReflectSetField(reflectTestName("abc"), reflectTestName("length"), 1) {
		t.Fatal("only an Array length is writable")
	}
	for _, receiver := range []any{array, reflectTestName("abc")} {
		fields := ReflectFields(receiver)
		if len(fields) != 1 || *fields[0] != "length" {
			t.Fatalf("fields of %T = %#v", receiver, fields)
		}
	}
}
