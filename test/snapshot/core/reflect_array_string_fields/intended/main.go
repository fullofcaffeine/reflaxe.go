package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

var arrayMembers *hxrt.Array = hxrt.NewArray(hxrt.StringFromLiteral("concat"), hxrt.StringFromLiteral("contains"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("filter"), hxrt.StringFromLiteral("indexOf"), hxrt.StringFromLiteral("insert"), hxrt.StringFromLiteral("join"), hxrt.StringFromLiteral("lastIndexOf"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("pop"), hxrt.StringFromLiteral("push"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("resize"), hxrt.StringFromLiteral("reverse"), hxrt.StringFromLiteral("shift"), hxrt.StringFromLiteral("slice"), hxrt.StringFromLiteral("sort"), hxrt.StringFromLiteral("splice"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unshift"))

func boundArrayMethods() {
	a := hxrt.NewArray(3, 1, 2, 1)
	row(hxrt.StringFromLiteral("indexOf"), call(a, hxrt.StringFromLiteral("indexOf"), hxrt.NewArray(1)))
	row(hxrt.StringFromLiteral("indexOf from"), call(a, hxrt.StringFromLiteral("indexOf"), hxrt.NewArray(1, 2)))
	row(hxrt.StringFromLiteral("indexOf negative from"), call(a, hxrt.StringFromLiteral("indexOf"), hxrt.NewArray(1, -1)))
	row(hxrt.StringFromLiteral("indexOf null from"), call(a, hxrt.StringFromLiteral("indexOf"), hxrt.NewArray(1, nil)))
	row(hxrt.StringFromLiteral("lastIndexOf"), call(a, hxrt.StringFromLiteral("lastIndexOf"), hxrt.NewArray(1)))
	row(hxrt.StringFromLiteral("lastIndexOf from"), call(a, hxrt.StringFromLiteral("lastIndexOf"), hxrt.NewArray(1, 2)))
	row(hxrt.StringFromLiteral("contains"), hxrt.NewArray(call(a, hxrt.StringFromLiteral("contains"), hxrt.NewArray(2)), call(a, hxrt.StringFromLiteral("contains"), hxrt.NewArray(9))))
	row(hxrt.StringFromLiteral("join"), call(a, hxrt.StringFromLiteral("join"), hxrt.NewArray(hxrt.StringFromLiteral("|"))))
	row(hxrt.StringFromLiteral("slice"), hxrt.NewArray(call(a, hxrt.StringFromLiteral("slice"), hxrt.NewArray(1)), call(a, hxrt.StringFromLiteral("slice"), hxrt.NewArray(1, 2)), call(a, hxrt.StringFromLiteral("slice"), hxrt.NewArray(1, nil))))
	row(hxrt.StringFromLiteral("copy"), call(a, hxrt.StringFromLiteral("copy"), hxrt.NewArray()))
	row(hxrt.StringFromLiteral("concat"), call(a, hxrt.StringFromLiteral("concat"), hxrt.NewArray(hxrt.NewArray(9))))
	row(hxrt.StringFromLiteral("toString"), call(a, hxrt.StringFromLiteral("toString"), hxrt.NewArray()))
	row(hxrt.StringFromLiteral("map"), call(a, hxrt.StringFromLiteral("map"), hxrt.NewArray(hxrt.NewFunction(func(x int) int {
		return int((hxrt.Int32Wrap(x) * hxrt.Int32Wrap(10)))
	}))))
	row(hxrt.StringFromLiteral("filter"), call(a, hxrt.StringFromLiteral("filter"), hxrt.NewArray(hxrt.NewFunction(func(x int) bool {
		return (x > 1)
	}))))
	row(hxrt.StringFromLiteral("remove"), call(a, hxrt.StringFromLiteral("remove"), hxrt.NewArray(1)))
	row(hxrt.StringFromLiteral("after remove"), a)
	row(hxrt.StringFromLiteral("shift"), call(a, hxrt.StringFromLiteral("shift"), hxrt.NewArray()))
	row(hxrt.StringFromLiteral("unshift"), call(a, hxrt.StringFromLiteral("unshift"), hxrt.NewArray(0)))
	row(hxrt.StringFromLiteral("push"), call(a, hxrt.StringFromLiteral("push"), hxrt.NewArray(7)))
	row(hxrt.StringFromLiteral("after push"), a)
	row(hxrt.StringFromLiteral("sort"), call(a, hxrt.StringFromLiteral("sort"), hxrt.NewArray(hxrt.NewFunction(func(x int, y int) int {
		return int((hxrt.Int32Wrap(x) - hxrt.Int32Wrap(y)))
	}))))
	row(hxrt.StringFromLiteral("after sort"), a)
	row(hxrt.StringFromLiteral("insert"), call(a, hxrt.StringFromLiteral("insert"), hxrt.NewArray(1, 5)))
	row(hxrt.StringFromLiteral("reverse"), call(a, hxrt.StringFromLiteral("reverse"), hxrt.NewArray()))
	row(hxrt.StringFromLiteral("after reverse"), a)
	row(hxrt.StringFromLiteral("splice"), call(a, hxrt.StringFromLiteral("splice"), hxrt.NewArray(1, 2)))
	row(hxrt.StringFromLiteral("pop"), call(a, hxrt.StringFromLiteral("pop"), hxrt.NewArray()))
	row(hxrt.StringFromLiteral("resize"), call(a, hxrt.StringFromLiteral("resize"), hxrt.NewArray(4)))
	row(hxrt.StringFromLiteral("after resize"), a)
	empty := hxrt.NewArray()
	row(hxrt.StringFromLiteral("empty pop and shift"), hxrt.NewArray(call(empty, hxrt.StringFromLiteral("pop"), hxrt.NewArray()), call(empty, hxrt.StringFromLiteral("shift"), hxrt.NewArray())))
	words := hxrt.NewArray(hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("c"))
	row(hxrt.StringFromLiteral("String elements"), hxrt.NewArray(call(words, hxrt.StringFromLiteral("indexOf"), hxrt.NewArray(hxrt.StringFromLiteral("c"))), call(words, hxrt.StringFromLiteral("join"), hxrt.NewArray(hxrt.StringFromLiteral("+")))))
}

func boundIdentity() {
	a := hxrt.NewArray(1)
	alias := a
	var push any = Reflect_field(a, hxrt.StringFromLiteral("push"))
	a.Push(2)
	reflaxe__go___internal__CompilerReflect_callMethod(push, []any{3})
	Reflect_callMethod(nil, push, hxrt.NewArray(4))
	row(hxrt.StringFromLiteral("bound push"), alias)
	var length any = Reflect_field(a, hxrt.StringFromLiteral("length"))
	a.Push(5)
	row(hxrt.StringFromLiteral("length read is a value"), length)
	var again any = Reflect_field(a, hxrt.StringFromLiteral("push"))
	row(hxrt.StringFromLiteral("compareMethods same"), hxrt.ReflectCompareMethods(push, again))
	row(hxrt.StringFromLiteral("compareMethods other Array"), func() bool {
		var f2 any = Reflect_field(hxrt.NewArray(1), hxrt.StringFromLiteral("push"))
		return hxrt.ReflectCompareMethods(push, f2)
	}())
	row(hxrt.StringFromLiteral("compareMethods other method"), func() bool {
		var f2_1 any = Reflect_field(a, hxrt.StringFromLiteral("pop"))
		return hxrt.ReflectCompareMethods(push, f2_1)
	}())
	row(hxrt.StringFromLiteral("fresh value"), hxrt.HaxeEqual(push, again))
	s := hxrt.StringFromLiteral("ab")
	row(hxrt.StringFromLiteral("String compareMethods"), func() bool {
		var f1 any = Reflect_field(s, hxrt.StringFromLiteral("charAt"))
		var f2_2 any = Reflect_field(s, hxrt.StringFromLiteral("charAt"))
		return hxrt.ReflectCompareMethods(f1, f2_2)
	}())
}

func boundStringMethods() {
	s := hxrt.StringFromLiteral("x,y,z")
	row(hxrt.StringFromLiteral("charAt"), hxrt.NewArray(call(s, hxrt.StringFromLiteral("charAt"), hxrt.NewArray(2)), call(s, hxrt.StringFromLiteral("charAt"), hxrt.NewArray(10))))
	row(hxrt.StringFromLiteral("charCodeAt"), hxrt.NewArray(call(s, hxrt.StringFromLiteral("charCodeAt"), hxrt.NewArray(0)), call(s, hxrt.StringFromLiteral("charCodeAt"), hxrt.NewArray(10))))
	row(hxrt.StringFromLiteral("indexOf"), hxrt.NewArray(call(s, hxrt.StringFromLiteral("indexOf"), hxrt.NewArray(hxrt.StringFromLiteral(","))), call(s, hxrt.StringFromLiteral("indexOf"), hxrt.NewArray(hxrt.StringFromLiteral(","), 2)), call(s, hxrt.StringFromLiteral("indexOf"), hxrt.NewArray(hxrt.StringFromLiteral(","), nil))))
	row(hxrt.StringFromLiteral("lastIndexOf"), hxrt.NewArray(call(s, hxrt.StringFromLiteral("lastIndexOf"), hxrt.NewArray(hxrt.StringFromLiteral(","))), call(s, hxrt.StringFromLiteral("lastIndexOf"), hxrt.NewArray(hxrt.StringFromLiteral(","), 2))))
	row(hxrt.StringFromLiteral("split"), call(s, hxrt.StringFromLiteral("split"), hxrt.NewArray(hxrt.StringFromLiteral(","))))
	row(hxrt.StringFromLiteral("substr"), hxrt.NewArray(call(s, hxrt.StringFromLiteral("substr"), hxrt.NewArray(2)), call(s, hxrt.StringFromLiteral("substr"), hxrt.NewArray(2, 1)), call(s, hxrt.StringFromLiteral("substr"), hxrt.NewArray(2, nil))))
	row(hxrt.StringFromLiteral("substring"), hxrt.NewArray(call(s, hxrt.StringFromLiteral("substring"), hxrt.NewArray(2)), call(s, hxrt.StringFromLiteral("substring"), hxrt.NewArray(2, 3)), call(s, hxrt.StringFromLiteral("substring"), hxrt.NewArray(2, nil))))
	row(hxrt.StringFromLiteral("case"), hxrt.NewArray(call(s, hxrt.StringFromLiteral("toUpperCase"), hxrt.NewArray()), call(hxrt.StringFromLiteral("ABC"), hxrt.StringFromLiteral("toLowerCase"), hxrt.NewArray())))
	row(hxrt.StringFromLiteral("toString"), call(s, hxrt.StringFromLiteral("toString"), hxrt.NewArray()))
}

func call(receiver any, name *string, arguments *hxrt.Array) any {
	return Reflect_callMethod(receiver, Reflect_field(receiver, name), arguments)
}

func dynamicReads() {
	var d any = hxrt.NewArray(1, 2, 3)
	row(hxrt.StringFromLiteral("d.length"), Reflect_field(d, hxrt.StringFromLiteral("length")))
	pushed := hxrt.IntFromNullableAny(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(d, hxrt.StringFromLiteral("push")), []any{4}))
	row(hxrt.StringFromLiteral("d.push"), pushed)
	row(hxrt.StringFromLiteral("d.indexOf"), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(d, hxrt.StringFromLiteral("indexOf")), []any{2}))
	row(hxrt.StringFromLiteral("d.join"), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(d, hxrt.StringFromLiteral("join")), []any{hxrt.StringFromLiteral(",")}))
	row(hxrt.StringFromLiteral("d.pop"), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(d, hxrt.StringFromLiteral("pop")), []any{}))
	row(hxrt.StringFromLiteral("d.length + 1"), hxrt.DynamicAdd(Reflect_field(d, hxrt.StringFromLiteral("length")), 1))
	n := hxrt.IntFromNullableAny(Reflect_field(d, hxrt.StringFromLiteral("length")))
	row(hxrt.StringFromLiteral("typed Int"), int((hxrt.Int32Wrap(n) * hxrt.Int32Wrap(2))))
	row(hxrt.StringFromLiteral("interpolation"), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("size "), hxrt.StdString(Reflect_field(d, hxrt.StringFromLiteral("length")))))
	row(hxrt.StringFromLiteral("condition"), func() any {
		var hx_if_3 any
		if hxrt.DynamicGreater(Reflect_field(d, hxrt.StringFromLiteral("length")), 2) {
			hx_if_3 = hxrt.StringFromLiteral("big")
		} else {
			hx_if_3 = hxrt.StringFromLiteral("small")
		}
		return hx_if_3
	}())
	row(hxrt.StringFromLiteral("comparison"), hxrt.HaxeEqual(Reflect_field(d, hxrt.StringFromLiteral("length")), 3))
	var s any = hxrt.StringFromLiteral("x,y,z")
	row(hxrt.StringFromLiteral("s.length"), Reflect_field(s, hxrt.StringFromLiteral("length")))
	row(hxrt.StringFromLiteral("s.charAt"), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(s, hxrt.StringFromLiteral("charAt")), []any{0}))
	row(hxrt.StringFromLiteral("s.split"), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(s, hxrt.StringFromLiteral("split")), []any{hxrt.StringFromLiteral(",")}))
	row(hxrt.StringFromLiteral("s.split.length"), Reflect_field(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(s, hxrt.StringFromLiteral("split")), []any{hxrt.StringFromLiteral(",")}), hxrt.StringFromLiteral("length")))
	row(hxrt.StringFromLiteral("s.indexOf"), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(s, hxrt.StringFromLiteral("indexOf")), []any{hxrt.StringFromLiteral("y")}))
	row(hxrt.StringFromLiteral("s.toUpperCase"), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(s, hxrt.StringFromLiteral("toUpperCase")), []any{}))
	row(hxrt.StringFromLiteral("s.substr"), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(s, hxrt.StringFromLiteral("substr")), []any{2}))
	row(hxrt.StringFromLiteral("s.charCodeAt"), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(s, hxrt.StringFromLiteral("charCodeAt")), []any{0}))
	c := hxrt.DynamicToString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(s, hxrt.StringFromLiteral("charAt")), []any{4}))
	row(hxrt.StringFromLiteral("typed String"), hxrt.StringConcatStringPtr(c, hxrt.StringFromLiteral("!")))
	hx_obj_4 := map[string]any{}
	hx_obj_4["items"] = hxrt.NewArray(1, 2)
	hx_obj_4["name"] = hxrt.StringFromLiteral("ab")
	var o any = hxrt.AnonymousFromMap(hx_obj_4)
	row(hxrt.StringFromLiteral("nested"), hxrt.DynamicAdd(Reflect_field(Reflect_field(o, hxrt.StringFromLiteral("items")), hxrt.StringFromLiteral("length")), Reflect_field(Reflect_field(o, hxrt.StringFromLiteral("name")), hxrt.StringFromLiteral("length"))))
	row(hxrt.StringFromLiteral("nested call"), hxrt.DynamicAdd(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(Reflect_field(o, hxrt.StringFromLiteral("name")), hxrt.StringFromLiteral("charAt")), []any{1}), reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(Reflect_field(o, hxrt.StringFromLiteral("items")), hxrt.StringFromLiteral("indexOf")), []any{2})))
	var parsed any = hxrt.JsonParse(hxrt.StringFromLiteral("{\"tags\":[\"a\",\"b\",\"c\"],\"title\":\"Hello\"}"))
	row(hxrt.StringFromLiteral("JSON"), hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(Reflect_field(Reflect_field(parsed, hxrt.StringFromLiteral("tags")), hxrt.StringFromLiteral("length"))), hxrt.StringFromLiteral(" ")), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(Reflect_field(parsed, hxrt.StringFromLiteral("tags")), hxrt.StringFromLiteral("join")), []any{hxrt.StringFromLiteral("/")}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(Reflect_field(parsed, hxrt.StringFromLiteral("title")), hxrt.StringFromLiteral("toLowerCase")), []any{}))))
}

func lengthWrites() {
	w := hxrt.NewArray(5, 6, 7)
	Reflect_setField(w, hxrt.StringFromLiteral("length"), 1)
	row(hxrt.StringFromLiteral("setField length 1"), w)
	grown := hxrt.NewArray(8)
	Reflect_setField(grown, hxrt.StringFromLiteral("length"), 3)
	row(hxrt.StringFromLiteral("setField length 3"), grown)
	Reflect_setProperty(grown, hxrt.StringFromLiteral("length"), 2)
	row(hxrt.StringFromLiteral("setProperty length 2"), grown)
	var d any = hxrt.NewArray(5, 6, 7)
	alias := func(hx_value_5 any) *hxrt.Array {
		if hx_value_5 == nil {
			var hx_zero_6 *hxrt.Array
			return hx_zero_6
		}
		return hx_value_5.(*hxrt.Array)
	}(d)
	row(hxrt.StringFromLiteral("d.length = 0"), func() any {
		hx_dynamic_target_7 := d
		hx_dynamic_value_9 := 0
		Reflect_setField(hx_dynamic_target_7, hxrt.StringFromLiteral("length"), hx_dynamic_value_9)
		return hx_dynamic_value_9
	}())
	row(hxrt.StringFromLiteral("after d.length = 0"), hxrt.NewArray(d, alias))
}

func main() {
	reflectReads()
	boundArrayMethods()
	boundStringMethods()
	boundIdentity()
	dynamicReads()
	lengthWrites()
}

func reflectReads() {
	numbers := hxrt.NewArray(1, 2, 3)
	text := hxrt.StringFromLiteral("abcd")
	row(hxrt.StringFromLiteral("field Array length"), Reflect_field(numbers, hxrt.StringFromLiteral("length")))
	row(hxrt.StringFromLiteral("field String length"), Reflect_field(text, hxrt.StringFromLiteral("length")))
	row(hxrt.StringFromLiteral("field empty length"), hxrt.DynamicAdd(Reflect_field(hxrt.NewArray(), hxrt.StringFromLiteral("length")), Reflect_field(hxrt.StringFromLiteral(""), hxrt.StringFromLiteral("length"))))
	_g := 0
	_g1 := arrayMembers
	for _g < _g1.Len() {
		name := func(hx_value_10 any) *string {
			if hx_value_10 == nil {
				var hx_zero_11 *string
				return hx_zero_11
			}
			return hx_value_10.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		row(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Array "), name), hxrt.StringFromLiteral(" is a function")), func() bool {
			var f any = Reflect_field(numbers, name)
			return hxrt.ReflectIsFunction(f)
		}())
	}
	_g_1 := 0
	_g1_1 := stringMembers
	for _g_1 < _g1_1.Len() {
		name_1 := func(hx_value_12 any) *string {
			if hx_value_12 == nil {
				var hx_zero_13 *string
				return hx_zero_13
			}
			return hx_value_12.(*string)
		}(_g1_1.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		row(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("String "), name_1), hxrt.StringFromLiteral(" is a function")), func() bool {
			var f_1 any = Reflect_field(text, name_1)
			return hxrt.ReflectIsFunction(f_1)
		}())
	}
	name_2 := hxrt.StringFromLiteral("missing")
	row(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Array "), name_2), Reflect_field(numbers, name_2))
	name_3 := hxrt.StringFromLiteral("Len")
	row(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Array "), name_3), Reflect_field(numbers, name_3))
	name_4 := hxrt.StringFromLiteral("Push")
	row(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Array "), name_4), Reflect_field(numbers, name_4))
	name_5 := hxrt.StringFromLiteral("Values")
	row(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Array "), name_5), Reflect_field(numbers, name_5))
	row(hxrt.StringFromLiteral("String missing"), Reflect_field(text, hxrt.StringFromLiteral("missing")))
	row(hxrt.StringFromLiteral("hasField Array"), hxrt.NewArray(Reflect_hasField(numbers, hxrt.StringFromLiteral("length")), Reflect_hasField(numbers, hxrt.StringFromLiteral("push")), Reflect_hasField(numbers, hxrt.StringFromLiteral("Len"))))
	row(hxrt.StringFromLiteral("hasField String"), hxrt.NewArray(Reflect_hasField(text, hxrt.StringFromLiteral("length")), Reflect_hasField(text, hxrt.StringFromLiteral("charAt"))))
	row(hxrt.StringFromLiteral("fields Array"), Reflect_fields(numbers))
	row(hxrt.StringFromLiteral("fields String"), Reflect_fields(text))
	row(hxrt.StringFromLiteral("getProperty"), hxrt.NewArray(Reflect_getProperty(numbers, hxrt.StringFromLiteral("length")), Reflect_getProperty(text, hxrt.StringFromLiteral("length"))))
	length := hxrt.IntFromNullableAny(Reflect_field(text, hxrt.StringFromLiteral("length")))
	row(hxrt.StringFromLiteral("typed length"), int((hxrt.Int32Wrap(length) * hxrt.Int32Wrap(2))))
}

func row(label *string, value any) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(": ")), hxrt.StdString(value)))
	hxrt.Println(v)
}

var stringMembers *hxrt.Array = hxrt.NewArray(hxrt.StringFromLiteral("charAt"), hxrt.StringFromLiteral("charCodeAt"), hxrt.StringFromLiteral("indexOf"), hxrt.StringFromLiteral("lastIndexOf"), hxrt.StringFromLiteral("split"), hxrt.StringFromLiteral("substr"), hxrt.StringFromLiteral("substring"), hxrt.StringFromLiteral("toLowerCase"), hxrt.StringFromLiteral("toUpperCase"), hxrt.StringFromLiteral("toString"))

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.details()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.get_message()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_native()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_1.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	classValue, ok := value.(*hxrt__TypeClassValue)
	if !ok || classValue == nil {
		return nil, false
	}
	className := *hxrt.StdString(classValue.name)
	switch className {
	default:
		return nil, false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedField(object any, field *string) any {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "savedMessage":
		return hxrt.ReadStoredField(value.savedMessage)
	case "savedNative":
		return hxrt.ReadStoredField(value.savedNative)
	case "savedPrevious":
		return hxrt.ReadStoredField(value.savedPrevious)
	case "savedStack":
		return hxrt.ReadStoredField(value.savedStack)
	}
	return nil
}

func hxrt__generated_field_has__haxe__Exception(value *haxe__Exception, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		return true
	case "savedNative":
		return true
	case "savedPrevious":
		return true
	case "savedStack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Exception(value *haxe__Exception, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		if incoming == nil {
			var zero *string
			value.savedMessage = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.savedMessage = typed
			return true
		default:
			return false
		}
	case "savedNative":
		if incoming == nil {
			var zero any
			value.savedNative = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedNative = typed
			return true
		default:
			return false
		}
	case "savedPrevious":
		if incoming == nil {
			var zero *haxe__Exception
			value.savedPrevious = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__Exception:
			value.savedPrevious = typed
			return true
		default:
			return false
		}
	case "savedStack":
		if incoming == nil {
			var zero any
			value.savedStack = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedStack = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "value":
		return hxrt.ReadStoredField(value.value)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__ValueException(value *haxe__ValueException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__ValueException(value *haxe__ValueException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		if incoming == nil {
			var zero any
			value.value = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.value = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return hxrt.ReadStoredField(value.high)
	case "low":
		return hxrt.ReadStoredField(value.low)
	}
	return nil
}

func hxrt__generated_field_has__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		return true
	case "low":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		if incoming == nil {
			var zero int
			value.high = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.high = typed
			return true
		default:
			return false
		}
	case "low":
		if incoming == nil {
			var zero int
			value.low = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.low = typed
			return true
		default:
			return false
		}
	}
	return false
}

func reflaxe__go___internal__CompilerReflect_typeField(object any, field *string) any {
	key := *hxrt.StdString(field)
	value, found := hxrt_typeClassMetadataField(object, key)
	if !found {
		return nil
	}
	return value
}

func reflaxe__go___internal__CompilerReflect_hasTypeField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	_, found := hxrt_typeClassMetadataField(object, key)
	return found
}

func reflaxe__go___internal__CompilerReflect_generatedMethod(object any, field *string) any {
	key := *hxrt.StdString(field)
	return hxrt__generated_method_field(object, key)
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	switch enumValue := value.(type) {
	case *haxe__StackItem:
		return (enumValue != nil)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_1 any) int {
			if hx_value_1 == nil {
				var hx_zero_2 int
				return hx_zero_2
			}
			return hx_value_1.(int)
		}(argument_0))
	case *hxrt.Function[func(int) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_3 any) int {
			if hx_value_3 == nil {
				var hx_zero_4 int
				return hx_zero_4
			}
			return hx_value_3.(int)
		}(argument_0))
	case *hxrt.Function[func(int, int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_5 any) int {
			if hx_value_5 == nil {
				var hx_zero_6 int
				return hx_zero_6
			}
			return hx_value_5.(int)
		}(argument_0), func(hx_value_7 any) int {
			if hx_value_7 == nil {
				var hx_zero_8 int
				return hx_zero_8
			}
			return hx_value_7.(int)
		}(argument_1))
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
