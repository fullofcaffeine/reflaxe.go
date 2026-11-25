package main

import (
	"math"
	"reflect"
	"snapshot/hxrt"
)

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_Holder interface {
}

type Holder struct {
	__hx_this I_Holder
	n         any
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder(n any) *Holder {
	return hx_init_Holder(hx_alloc_Holder(nil), n)
}

func hx_alloc_Holder(owner any) *Holder {
	self := &Holder{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Holder)
	}
	return self
}

func hx_init_Holder(self *Holder, n any) *Holder {
	self.n = n
	return self
}

func (self *Holder) HxClassName() string {
	return "Holder"
}

var Holder_shared any = any(1)

func arithmetic() {
	var i any = dyn(7)
	var f any = dyn(2.5)
	var s any = dyn(hxrt.StringFromLiteral("x"))
	var max_local any = dyn(2147483647)
	var min_local any = dyn(-2147483648)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("add "), hxrt.StdString(hxrt.DynamicAdd(i, 2))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicAdd(i, 0.5))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicAdd(f, 2))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicAdd(i, i))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicAdd(i, f))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicAdd(s, i))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicAdd(i, s))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("add types "), hxrt.StdString(Type_typeof(hxrt.DynamicAdd(i, 2)))), hxrt.StringFromLiteral(" ")), hxrt.StdString(Type_typeof(hxrt.DynamicAdd(i, 0.5)))), hxrt.StringFromLiteral(" ")), hxrt.StdString(Type_typeof(hxrt.DynamicAdd(i, i)))))
	hxrt.Println(v_1)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("sub "), hxrt.DynamicToFloat(hxrt.DynamicSub(i, 2))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicSub(2, i))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicSub(f, 0.5))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicSub(i, f)))))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("mul "), hxrt.DynamicToFloat(hxrt.DynamicMul(i, 2))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicMul(f, 2))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicMul(i, f)))))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("div "), hxrt.DynamicDiv(i, 2)), hxrt.StringFromLiteral(" ")), hxrt.DynamicDiv(f, 0.5)), hxrt.StringFromLiteral(" ")), hxrt.DynamicDiv(i, i))))
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("mod "), hxrt.DynamicToFloat(hxrt.DynamicMod(i, 4))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicMod(dyn(-7), 2))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicMod(f, 1))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicMod(dyn(-7.5), 2))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("zero "), hxrt.StdString(func() bool {
		f_1 := hxrt.DynamicDiv(dyn(0), 0)
		return math.IsNaN(f_1)
	}())), hxrt.StringFromLiteral(" ")), hxrt.StdString((hxrt.DynamicDiv(dyn(1), 0) == Math_POSITIVE_INFINITY))), hxrt.StringFromLiteral(" ")), hxrt.StdString(func() bool {
		f_2 := hxrt.DynamicToFloat(hxrt.DynamicMod(dyn(5), 0))
		return math.IsNaN(f_2)
	}())))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("overflow "), hxrt.StdString(hxrt.DynamicAdd(max_local, 1))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicMul(dyn(65536), dyn(65536)))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicMul(max_local, 2))), hxrt.StringFromLiteral(" ")), hxrt.StdString((hxrt.DynamicToFloat(hxrt.DynamicSub(min_local, 1)) == 2147483647))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.HaxeEqual(hxrt.DynamicAdd(max_local, 0.5), 2147483647.5))))
	hxrt.Println(v_4)
	var none any = dyn(nil)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null text "), hxrt.StringConcatStringPtr(hxrt.StdString(none), hxrt.StringFromLiteral("t"))), hxrt.StringFromLiteral(" ")), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("t"), hxrt.StdString(none))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicAdd(none, s))))
	hxrt.Println(v_5)
}

func attempt(name *string, run *hxrt.Function[func() any]) {
	var hx_try_1 *string
	hxrt.TryCatch(func() {
		hx_try_1 = hxrt.StdString(func() any {
			var hx_callable_2 *hxrt.Function[func() any] = run
			return hxrt.FunctionCall(hx_callable_2)()
		}())
	}, func(hx_caught_3 any) {
		hx_exception_4 := haxe__Exception_caught(hx_caught_3)
		hx_payload_5 := haxe__Exception_catchValue(hx_caught_3)
		_ = hx_exception_4
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_5) {
			var e *string = func(hx_value_6 any) *string {
				if hx_value_6 == nil {
					var hx_zero_7 *string
					return hx_zero_7
				}
				return hx_value_6.(*string)
			}(hx_payload_5)
			hx_try_1 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error "), e)
		} else {
			hxrt.Throw(hx_caught_3)
		}
	})
	text := hx_try_1
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(name, hxrt.StringFromLiteral(" = ")), text)))
}

func bitwise() {
	var b any = dyn(12)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("bits "), hxrt.DynamicAnd(b, 10)), hxrt.StringFromLiteral(" ")), hxrt.DynamicOr(b, 3)), hxrt.StringFromLiteral(" ")), hxrt.DynamicXor(b, 5)), hxrt.StringFromLiteral(" ")), hxrt.DynamicShl(b, 2)), hxrt.StringFromLiteral(" ")), hxrt.DynamicShr(b, 1)), hxrt.StringFromLiteral(" ")), hxrt.DynamicAnd(6, b))))
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("shifts "), hxrt.DynamicShr(dyn(-8), 1)), hxrt.StringFromLiteral(" ")), hxrt.DynamicUShr(dyn(-8), 1)), hxrt.StringFromLiteral(" ")), hxrt.DynamicShl(dyn(1), 33)), hxrt.StringFromLiteral(" ")), hxrt.DynamicShl(dyn(2147483647), 1)))
	hxrt.Println(v)
}

func comparisons() {
	var i any = dyn(7)
	var none any = dyn(nil)
	hx_obj_8 := map[string]any{}
	hx_obj_8["a"] = 1
	var o any = hxrt.AnonymousFromMap(hx_obj_8)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("numbers "), hxrt.StdString(hxrt.DynamicLess(i, 8))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicGreater(i, 8))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicLessEqual(i, 7))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicGreaterEqual(i, 7.5))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicLess(dyn(2.5), i))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicLess(3, i))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("strings "), hxrt.StdString(hxrt.DynamicLess(dyn(hxrt.StringFromLiteral("a")), dyn(hxrt.StringFromLiteral("b"))))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicLess(dyn(hxrt.StringFromLiteral("10")), hxrt.StringFromLiteral("9")))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicGreaterEqual(dyn(hxrt.StringFromLiteral("b")), hxrt.StringFromLiteral("a")))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bools "), hxrt.StdString(hxrt.DynamicLess(dyn(false), dyn(true)))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicLessEqual(dyn(true), dyn(true)))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mixed "), hxrt.StdString(hxrt.DynamicLess(dyn(hxrt.StringFromLiteral("x")), 1))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicGreater(dyn(hxrt.StringFromLiteral("x")), 1))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicLess(i, dyn(hxrt.StringFromLiteral("x"))))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null "), hxrt.StdString(hxrt.DynamicLess(none, 1))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicGreaterEqual(none, 0))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicLessEqual(none, none))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicLess(none, none))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicGreaterEqual(i, none))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("objects "), hxrt.StdString(hxrt.DynamicLessEqual(o, o))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicLess(o, o))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicGreaterEqual(o, dyn(func() hxrt.Anonymous {
		hx_obj_9 := map[string]any{}
		hx_obj_9["a"] = 1
		return hxrt.AnonymousFromMap(hx_obj_9)
	}())))))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("equality "), hxrt.StdString(hxrt.HaxeEqual(i, 7))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.HaxeEqual(i, 7.0))), hxrt.StringFromLiteral(" ")), hxrt.StdString(!hxrt.HaxeEqual(i, dyn(8)))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.AnyEqualsNull(none))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.StringEqualAny(dyn(hxrt.StringFromLiteral("x")), hxrt.StringFromLiteral("x")))))
	hxrt.Println(v_6)
}

func compound() {
	var d any = dyn(10)
	d = hxrt.DynamicAdd(d, 5)
	d = hxrt.DynamicSub(d, 3)
	d = hxrt.DynamicMul(d, 2)
	d = hxrt.DynamicMod(d, 7)
	d = hxrt.DynamicShl(d, 2)
	d = hxrt.DynamicOr(d, 1)
	d = hxrt.DynamicAnd(d, 7)
	d = hxrt.DynamicXor(d, 3)
	d = hxrt.DynamicShr(d, 1)
	d = hxrt.DynamicUShr(d, 1)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("local "), hxrt.StdString(d)))
	hxrt.Println(v)
	d = hxrt.DynamicDiv(d, 4)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("divided "), hxrt.StdString(d)))
	hxrt.Println(v_1)
	result := hxrt.DynamicToString(func() any {
		d = hxrt.DynamicAdd(d, 1)
		return d
	}())
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("value "), result), hxrt.StringFromLiteral(" ")), hxrt.StdString(d)))
	hxrt.Println(v_2)
	var text any = dyn(hxrt.StringFromLiteral("a"))
	text = hxrt.DynamicAdd(text, 1)
	text = hxrt.StringConcatStringPtr(hxrt.StdString(text), hxrt.StringFromLiteral("b"))
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("text "), hxrt.StdString(text)))
	hxrt.Println(v_3)
	var wrap any = dyn(2147483647)
	wrap = hxrt.DynamicAdd(wrap, 1)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("wrap "), hxrt.StdString(wrap)))
	hxrt.Println(v_4)
	hx_obj_10 := map[string]any{}
	hx_obj_10["n"] = 1
	hx_obj_10["s"] = hxrt.StringFromLiteral("s")
	var o any = hxrt.AnonymousFromMap(hx_obj_10)
	func() any {
		hx_dynamic_target_11 := o
		hx_dynamic_old_12 := Reflect_field(hx_dynamic_target_11, hxrt.StringFromLiteral("n"))
		hx_dynamic_value_13 := hxrt.DynamicAdd(hx_dynamic_old_12, 1)
		Reflect_setField(hx_dynamic_target_11, hxrt.StringFromLiteral("n"), hx_dynamic_value_13)
		return hx_dynamic_value_13
	}()
	func() any {
		hx_dynamic_target_14 := o
		hx_dynamic_old_15 := Reflect_field(hx_dynamic_target_14, hxrt.StringFromLiteral("n"))
		hx_dynamic_value_16 := hxrt.DynamicMul(hx_dynamic_old_15, 10)
		Reflect_setField(hx_dynamic_target_14, hxrt.StringFromLiteral("n"), hx_dynamic_value_16)
		return hx_dynamic_value_16
	}()
	func() any {
		hx_dynamic_target_17 := o
		hx_dynamic_old_18 := Reflect_field(hx_dynamic_target_17, hxrt.StringFromLiteral("s"))
		hx_dynamic_value_19 := hxrt.DynamicAdd(hx_dynamic_old_18, 2)
		Reflect_setField(hx_dynamic_target_17, hxrt.StringFromLiteral("s"), hx_dynamic_value_19)
		return hx_dynamic_value_19
	}()
	fieldResult := hxrt.DynamicToString(func() any {
		hx_dynamic_target_20 := o
		hx_dynamic_old_21 := Reflect_field(hx_dynamic_target_20, hxrt.StringFromLiteral("n"))
		hx_dynamic_value_22 := hxrt.DynamicSub(hx_dynamic_old_21, 5)
		Reflect_setField(hx_dynamic_target_20, hxrt.StringFromLiteral("n"), hx_dynamic_value_22)
		return hx_dynamic_value_22
	}())
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field "), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("n")))), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("s")))), hxrt.StringFromLiteral(" ")), fieldResult))
	hxrt.Println(v_5)
	h := New_Holder(2)
	h.n = hxrt.DynamicAdd(h.n, 0.5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("class field "), hxrt.StdString(h.n)))
	hxrt.Println(v_6)
	var values_0 any
	values_0 = 1
	values_0 = hxrt.DynamicAdd(values_0, 2)
	values_0 = hxrt.DynamicMul(values_0, 3)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array "), hxrt.StdString(values_0)))
	hxrt.Println(v_7)
	typedInt := 1
	typedInt = hxrt.IntFromNullableAny(hxrt.DynamicAdd(typedInt, dyn(2)))
	typedFloat := 1.5
	typedFloat = hxrt.DynamicToFloat(hxrt.DynamicAdd(typedFloat, dyn(2)))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("typed targets "), typedInt), hxrt.StringFromLiteral(" ")), typedFloat)))
}

func conditions() {
	if hxrt.DynamicIsTrue(dyn(true)) {
		hxrt.Println(any(hxrt.StringFromLiteral("if call true")))
	}
	if hxrt.DynamicIsTrue(dyn(1)) {
		hxrt.Println(any(hxrt.StringFromLiteral("if one true")))
	} else {
		hxrt.Println(any(hxrt.StringFromLiteral("if one false")))
	}
	var none any = dyn(nil)
	if hxrt.DynamicIsTrue(none) {
		hxrt.Println(any(hxrt.StringFromLiteral("if null true")))
	} else {
		hxrt.Println(any(hxrt.StringFromLiteral("if null false")))
	}
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ternary "), func() *string {
		var hx_if_23 *string
		if hxrt.DynamicIsTrue(dyn(true)) {
			hx_if_23 = hxrt.StringFromLiteral("yes")
		} else {
			hx_if_23 = hxrt.StringFromLiteral("no")
		}
		return hx_if_23
	}()), hxrt.StringFromLiteral(" ")), func() *string {
		var hx_if_24 *string
		if hxrt.DynamicIsTrue(dyn(hxrt.StringFromLiteral("x"))) {
			hx_if_24 = hxrt.StringFromLiteral("yes")
		} else {
			hx_if_24 = hxrt.StringFromLiteral("no")
		}
		return hx_if_24
	}()), hxrt.StringFromLiteral(" ")), func() *string {
		var hx_if_25 *string
		if hxrt.DynamicIsTrue(none) {
			hx_if_25 = hxrt.StringFromLiteral("yes")
		} else {
			hx_if_25 = hxrt.StringFromLiteral("no")
		}
		return hx_if_25
	}()))
	hxrt.Println(v)
	counter = 0
	loops := 0
	for hxrt.DynamicIsTrue(more()) {
		loops = int(int32((loops + 1)))
	}
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("while "), loops)))
	counter = 0
	passes := 0
	hx_do_first_26 := true
	for hx_do_first_26 || hxrt.DynamicIsTrue(more()) {
		hx_do_first_26 = false
		passes = int(int32((passes + 1)))
	}
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("do-while "), passes)))
	var running any = dyn(true)
	steps := 0
	for hxrt.DynamicIsTrue(running) {
		steps = int(int32((steps + 1)))
		if steps == 3 {
			running = false
		}
	}
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("while local "), steps)))
	var high any = dyn(4)
	sum := 0
	_g := 0
	_g1 := hxrt.IntFromNullableAny(high)
	for _g < _g1 {
		hx_post_27 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_27
		sum = int((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(i)))
	}
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("for range "), sum)))
	var limit any = dyn(3)
	k := 0
	for hxrt.DynamicLess(k, limit) {
		k = int(int32((k + 1)))
	}
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("compare loop "), k)))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("logic "), hxrt.StdString((hxrt.DynamicIsTrue(dyn(true)) && hxrt.DynamicIsTrue(dyn(true))))), hxrt.StringFromLiteral(" ")), hxrt.StdString((hxrt.DynamicIsTrue(dyn(false)) || hxrt.DynamicIsTrue(dyn(true))))), hxrt.StringFromLiteral(" ")), hxrt.StdString((hxrt.DynamicIsTrue(dyn(true)) && false))), hxrt.StringFromLiteral(" ")), hxrt.StdString(dyn(false))))
	hxrt.Println(v_1)
	if hxrt.DynamicIsTrue(dyn(1)) {
		hxrt.Println(any(hxrt.StringFromLiteral("and one true")))
	} else {
		hxrt.Println(any(hxrt.StringFromLiteral("and one false")))
	}
	if hxrt.DynamicIsTrue(dyn(false)) || hxrt.DynamicIsTrue(dyn(true)) {
		hxrt.Println(any(hxrt.StringFromLiteral("or true")))
	}
}

var counter int = 0

func dyn(value any) any {
	return value
}

func Main_errors() {
	attempt(hxrt.StringFromLiteral("null + 1"), hxrt.NewFunction(func() any {
		return hxrt.DynamicAdd(dyn(nil), 1)
	}))
	attempt(hxrt.StringFromLiteral("true + 1"), hxrt.NewFunction(func() any {
		return hxrt.DynamicAdd(dyn(true), 1)
	}))
	attempt(hxrt.StringFromLiteral("x - 1"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() float64 {
		return hxrt.DynamicToFloat(hxrt.DynamicSub(dyn(hxrt.StringFromLiteral("x")), 1))
	})))
	attempt(hxrt.StringFromLiteral("null * 2"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() float64 {
		return hxrt.DynamicToFloat(hxrt.DynamicMul(dyn(nil), 2))
	})))
	attempt(hxrt.StringFromLiteral("2.5 & 1"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() int {
		return hxrt.DynamicAnd(dyn(2.5), 1)
	})))
	attempt(hxrt.StringFromLiteral("2.0 | 1"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() int {
		return hxrt.DynamicOr(dyn(2.0), 1)
	})))
	attempt(hxrt.StringFromLiteral("-null"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() float64 {
		return hxrt.DynamicToFloat(hxrt.DynamicNeg(dyn(nil)))
	})))
	attempt(hxrt.StringFromLiteral("~x"), hx_recover_callable_c5309b43(hxrt.NewFunction(func() int {
		return hxrt.IntFromNullableAny(hxrt.DynamicComplement(dyn(hxrt.StringFromLiteral("x"))))
	})))
	attempt(hxrt.StringFromLiteral("null++"), hxrt.NewFunction(func() any {
		var d any = dyn(nil)
		d = hxrt.DynamicAdd(d, 1)
		return d
	}))
	attempt(hxrt.StringFromLiteral("ok"), hxrt.NewFunction(func() any {
		return hxrt.DynamicAdd(dyn(1), 1)
	}))
}

func evaluationOrder() {
	order = hxrt.StringFromLiteral("")
	sum := hxrt.DynamicToString(hxrt.DynamicAdd(step(hxrt.StringFromLiteral("a"), 1), step(hxrt.StringFromLiteral("b"), 2)))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("operands "), order), hxrt.StringFromLiteral(" ")), sum))
	hxrt.Println(v)
	order = hxrt.StringFromLiteral("")
	less := hxrt.DynamicLess(step(hxrt.StringFromLiteral("l"), 1), step(hxrt.StringFromLiteral("r"), 2))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("compare "), order), hxrt.StringFromLiteral(" ")), hxrt.StdString(less)))
	hxrt.Println(v_1)
	hx_obj_28 := map[string]any{}
	hx_obj_28["n"] = 1
	var target any = hxrt.AnonymousFromMap(hx_obj_28)
	order = hxrt.StringFromLiteral("")
	var fh any = step(hxrt.StringFromLiteral("t"), target)
	func() any {
		hx_dynamic_target_29 := fh
		hx_dynamic_old_30 := Reflect_field(hx_dynamic_target_29, hxrt.StringFromLiteral("n"))
		hx_dynamic_value_31 := hxrt.DynamicAdd(hx_dynamic_old_30, step(hxrt.StringFromLiteral("v"), 2))
		Reflect_setField(hx_dynamic_target_29, hxrt.StringFromLiteral("n"), hx_dynamic_value_31)
		return hx_dynamic_value_31
	}()
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field compound "), order), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(target, hxrt.StringFromLiteral("n")))))
	hxrt.Println(v_2)
	order = hxrt.StringFromLiteral("")
	func() any {
		hx_dynamic_target_32 := step(hxrt.StringFromLiteral("t"), target)
		hx_dynamic_old_33 := Reflect_field(hx_dynamic_target_32, hxrt.StringFromLiteral("n"))
		hx_dynamic_value_34 := hxrt.DynamicAdd(hx_dynamic_old_33, 1)
		Reflect_setField(hx_dynamic_target_32, hxrt.StringFromLiteral("n"), hx_dynamic_value_34)
		return hx_dynamic_old_33
	}()
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field increment "), order), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(target, hxrt.StringFromLiteral("n")))))
	hxrt.Println(v_3)
	var y any = dyn(10)
	z := func() float64 {
		hx_dynamic_left_35 := y
		return hxrt.DynamicToFloat(hxrt.DynamicSub(hx_dynamic_left_35, func() any {
			y = 3
			return y
		}()))
	}()
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("left first "), z), hxrt.StringFromLiteral(" ")), hxrt.StdString(y)))
	hxrt.Println(v_4)
	var w any = dyn(5)
	bump := hxrt.NewFunction(func() int {
		w = 100
		return 1
	})
	captured := hxrt.DynamicToString(func() any {
		hx_dynamic_left_37 := w
		return hxrt.DynamicAdd(hx_dynamic_left_37, func() int {
			var hx_callable_36 *hxrt.Function[func() int] = bump
			return hxrt.FunctionCall(hx_callable_36)()
		}())
	}())
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("captured left first "), captured), hxrt.StringFromLiteral(" ")), hxrt.StdString(w)))
	hxrt.Println(v_5)
}

func increments() {
	var d any = dyn(7)
	d = hxrt.DynamicAdd(d, 1)
	d = hxrt.DynamicAdd(d, 1)
	d = hxrt.DynamicSub(d, 1)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("local "), hxrt.StdString(d)))
	hxrt.Println(v)
	hx_post_38 := d
	d = hxrt.DynamicAdd(d, 1)
	post := hxrt.DynamicToFloat(hx_post_38)
	pre := hxrt.DynamicToFloat(func() any {
		d = hxrt.DynamicAdd(d, 1)
		return d
	}())
	hx_post_39 := d
	d = hxrt.DynamicSub(d, 1)
	postDown := hxrt.DynamicToFloat(hx_post_39)
	preDown := hxrt.DynamicToFloat(func() any {
		d = hxrt.DynamicSub(d, 1)
		return d
	}())
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("values "), post), hxrt.StringFromLiteral(" ")), pre), hxrt.StringFromLiteral(" ")), postDown), hxrt.StringFromLiteral(" ")), preDown), hxrt.StringFromLiteral(" ")), hxrt.StdString(d)))
	hxrt.Println(v_1)
	var f any = dyn(2.5)
	f = hxrt.DynamicAdd(f, 1)
	hx_post_40 := f
	f = hxrt.DynamicSub(f, 1)
	floatPost := hxrt.DynamicToFloat(hx_post_40)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("float "), hxrt.StdString(f)), hxrt.StringFromLiteral(" ")), floatPost))
	hxrt.Println(v_2)
	var max_local any = dyn(2147483647)
	max_local = hxrt.DynamicAdd(max_local, 1)
	var text any = dyn(hxrt.StringFromLiteral("s"))
	text = hxrt.DynamicAdd(text, 1)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("wrap "), hxrt.StdString(max_local)), hxrt.StringFromLiteral(" text ")), hxrt.StdString(text)))
	hxrt.Println(v_3)
	hx_obj_41 := map[string]any{}
	hx_obj_41["n"] = 1
	var o any = hxrt.AnonymousFromMap(hx_obj_41)
	func() any {
		hx_dynamic_target_42 := o
		hx_dynamic_old_43 := Reflect_field(hx_dynamic_target_42, hxrt.StringFromLiteral("n"))
		hx_dynamic_value_44 := hxrt.DynamicAdd(hx_dynamic_old_43, 1)
		Reflect_setField(hx_dynamic_target_42, hxrt.StringFromLiteral("n"), hx_dynamic_value_44)
		return hx_dynamic_old_43
	}()
	func() any {
		hx_dynamic_target_45 := o
		hx_dynamic_old_46 := Reflect_field(hx_dynamic_target_45, hxrt.StringFromLiteral("n"))
		hx_dynamic_value_47 := hxrt.DynamicAdd(hx_dynamic_old_46, 1)
		Reflect_setField(hx_dynamic_target_45, hxrt.StringFromLiteral("n"), hx_dynamic_value_47)
		return hx_dynamic_value_47
	}()
	fieldPost := hxrt.DynamicToFloat(func() any {
		hx_dynamic_target_48 := o
		hx_dynamic_old_49 := Reflect_field(hx_dynamic_target_48, hxrt.StringFromLiteral("n"))
		hx_dynamic_value_50 := hxrt.DynamicAdd(hx_dynamic_old_49, 1)
		Reflect_setField(hx_dynamic_target_48, hxrt.StringFromLiteral("n"), hx_dynamic_value_50)
		return hx_dynamic_old_49
	}())
	fieldPre := hxrt.DynamicToFloat(func() any {
		hx_dynamic_target_51 := o
		hx_dynamic_old_52 := Reflect_field(hx_dynamic_target_51, hxrt.StringFromLiteral("n"))
		hx_dynamic_value_53 := hxrt.DynamicSub(hx_dynamic_old_52, 1)
		Reflect_setField(hx_dynamic_target_51, hxrt.StringFromLiteral("n"), hx_dynamic_value_53)
		return hx_dynamic_value_53
	}())
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("anon field "), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("n")))), hxrt.StringFromLiteral(" ")), fieldPost), hxrt.StringFromLiteral(" ")), fieldPre))
	hxrt.Println(v_4)
	h := New_Holder(5)
	h.n = hxrt.DynamicAdd(h.n, 1)
	hx_post_54 := h.n
	h.n = hxrt.DynamicSub(h.n, 1)
	held := hxrt.DynamicToFloat(hx_post_54)
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("class field "), hxrt.StdString(h.n)), hxrt.StringFromLiteral(" ")), held))
	hxrt.Println(v_5)
	var viaDynamic any = New_Holder(1.5)
	func() any {
		hx_dynamic_target_55 := viaDynamic
		hx_dynamic_old_56 := Reflect_field(hx_dynamic_target_55, hxrt.StringFromLiteral("n"))
		hx_dynamic_value_57 := hxrt.DynamicAdd(hx_dynamic_old_56, 1)
		Reflect_setField(hx_dynamic_target_55, hxrt.StringFromLiteral("n"), hx_dynamic_value_57)
		return hx_dynamic_old_56
	}()
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("class via dynamic "), hxrt.StdString(Reflect_field(viaDynamic, hxrt.StringFromLiteral("n")))))
	hxrt.Println(v_6)
	Holder_shared = hxrt.DynamicAdd(Holder_shared, 1)
	sharedPre := hxrt.DynamicToFloat(func() any {
		Holder_shared = hxrt.DynamicAdd(Holder_shared, 1)
		return Holder_shared
	}())
	var v_7 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("static "), hxrt.StdString(Holder_shared)), hxrt.StringFromLiteral(" ")), sharedPre))
	hxrt.Println(v_7)
	values := hxrt.NewArray(1, 2.5)
	func() any {
		hx_array_target_58 := values
		hx_array_index_59 := 0
		var hx_array_current_60 any = hx_array_target_58.Get(hx_array_index_59)
		var hx_array_next_61 any = hxrt.DynamicAdd(hx_array_current_60, 1)
		hx_array_target_58.Set(hx_array_index_59, hx_array_next_61)
		return hx_array_current_60
	}()
	elementPost := hxrt.DynamicToFloat(func() any {
		hx_array_target_62 := values
		hx_array_index_63 := 1
		var hx_array_current_64 any = hx_array_target_62.Get(hx_array_index_63)
		var hx_array_next_65 any = hxrt.DynamicAdd(hx_array_current_64, 1)
		hx_array_target_62.Set(hx_array_index_63, hx_array_next_65)
		return hx_array_current_64
	}())
	var v_8 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array "), hxrt.StdString(values)), hxrt.StringFromLiteral(" ")), elementPost))
	hxrt.Println(v_8)
}

func main() {
	arithmetic()
	typedOperands()
	bitwise()
	Main_errors()
	increments()
	compound()
	unary()
	conditions()
	comparisons()
	evaluationOrder()
}

func more() any {
	counter = int(int32((counter + 1)))
	return (counter < 3)
}

var order *string = hxrt.StringFromLiteral("")

func step(name *string, value any) any {
	order = hxrt.StringConcatStringPtr(order, name)
	return value
}

func typedOperands() {
	n := 3
	x := 0.5
	t := hxrt.StringFromLiteral("t")
	var d any = dyn(4)
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typed "), hxrt.StdString(hxrt.DynamicAdd(d, n))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicAdd(n, d))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicSub(d, x))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicMul(x, d))), hxrt.StringFromLiteral(" ")), hxrt.StringConcatStringPtr(hxrt.StdString(d), t)), hxrt.StringFromLiteral(" ")), hxrt.StringConcatStringPtr(t, hxrt.StdString(d))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicSub(n, d))), hxrt.StringFromLiteral(" ")), hxrt.DynamicDiv(n, d)))
	hxrt.Println(v)
	plain := int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(n) * hxrt.Int32Wrap(2)))) + hxrt.Int32Wrap(1)))
	real_local := ((x * float64(4)) - float64(1))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("plain "), plain), hxrt.StringFromLiteral(" ")), real_local)))
}

func unary() {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("neg "), hxrt.DynamicToFloat(hxrt.DynamicNeg(dyn(7)))), hxrt.StringFromLiteral(" ")), hxrt.DynamicToFloat(hxrt.DynamicNeg(dyn(2.5)))), hxrt.StringFromLiteral(" ")), hxrt.StdString((hxrt.DynamicToFloat(hxrt.DynamicNeg(dyn(-2147483648))) == -2147483648))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("not "), hxrt.StdString(hxrt.DynamicNot(dyn(true)))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicNot(dyn(false)))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicNot(dyn(nil)))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicNot(dyn(0)))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicNot(dyn(hxrt.StringFromLiteral(""))))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.DynamicNot(dyn(1)))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("complement "), hxrt.IntFromNullableAny(hxrt.DynamicComplement(dyn(7)))), hxrt.StringFromLiteral(" ")), hxrt.IntFromNullableAny(hxrt.DynamicComplement(dyn(0)))), hxrt.StringFromLiteral(" ")), hxrt.IntFromNullableAny(hxrt.DynamicComplement(dyn(-2147483648)))))
	hxrt.Println(v_2)
	var flag any = dyn(false)
	if hxrt.DynamicNot(flag) {
		hxrt.Println(any(hxrt.StringFromLiteral("not local")))
	}
}

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

type Type struct {
}

func hxrt_typeCreateEmpty__Holder() *Holder {
	instance := &Holder{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__Exception() *haxe__Exception {
	instance := &haxe__Exception{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ValueException() *haxe__ValueException {
	instance := &haxe__ValueException{}
	instance.haxe__Exception = &haxe__Exception{}
	instance.haxe__Exception.__hx_this = instance
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe___Int64_____Int64() *haxe___Int64_____Int64 {
	instance := &haxe___Int64_____Int64{}
	instance.__hx_this = instance
	return instance
}

type ValueType struct {
	tag    int
	params []any
}

func (self *ValueType) HxEnumValue() (string, []any) {
	return [9]string{"TNull", "TInt", "TFloat", "TBool", "TObject", "TFunction", "TClass", "TEnum", "TUnknown"}[self.tag], self.params
}

var ValueType_TNull *ValueType = &ValueType{tag: 0, params: []any{}}

var ValueType_TInt *ValueType = &ValueType{tag: 1, params: []any{}}

var ValueType_TFloat *ValueType = &ValueType{tag: 2, params: []any{}}

var ValueType_TBool *ValueType = &ValueType{tag: 3, params: []any{}}

var ValueType_TObject *ValueType = &ValueType{tag: 4, params: []any{}}

var ValueType_TFunction *ValueType = &ValueType{tag: 5, params: []any{}}

var ValueType_TUnknown *ValueType = &ValueType{tag: 8, params: []any{}}

func ValueType_TClass(c any) *ValueType {
	return &ValueType{tag: 6, params: []any{c}}
}

func ValueType_TEnum(e any) *ValueType {
	return &ValueType{tag: 7, params: []any{e}}
}

func hxrt_typeCallAny(callable any, args []any) (any, bool) {
	result := any(nil)
	ok := false
	defer func() {
		if recover() != nil {
			result = nil
			ok = false
		}
	}()
	if callable == nil {
		return nil, false
	}
	fn := reflect.ValueOf(callable)
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil, false
	}
	fnType := fn.Type()
	if fnType.NumIn() != len(args) {
		return nil, false
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		paramType := fnType.In(i)
		arg := args[i]
		if arg == nil {
			in[i] = reflect.Zero(paramType)
			continue
		}
		v := reflect.ValueOf(arg)
		if v.IsValid() && v.Type().AssignableTo(paramType) {
			in[i] = v
			continue
		}
		if v.IsValid() && v.Type().ConvertibleTo(paramType) {
			in[i] = v.Convert(paramType)
			continue
		}
		if paramType.Kind() == reflect.Interface && v.IsValid() {
			in[i] = v
			continue
		}
		return nil, false
	}
	out := fn.Call(in)
	if len(out) == 0 {
		return nil, true
	}
	first := out[0]
	if !first.IsValid() {
		return nil, true
	}
	result = first.Interface()
	ok = true
	return result, ok
}

func hxrt_typeArrayValues(value *hxrt.Array) []any {
	if value == nil {
		return []any{}
	}
	return value.Values()
}

func hxrt_typeResolvedClassName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeClassValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeClassValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeResolvedEnumName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeEnumValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeEnumValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeCreateClassInstance(className string, args []any) (any, bool) {
	switch className {
	case "Holder":
		return hxrt_typeCallAny(New_Holder, args)
	case "Main":
		return nil, false
	case "Math":
		return nil, false
	case "Reflect":
		return nil, false
	case "haxe.Exception":
		return hxrt_typeCallAny(New_haxe__Exception, args)
	case "haxe.Int64Helper":
		return nil, false
	case "haxe.ValueException":
		return hxrt_typeCallAny(New_haxe__ValueException, args)
	case "haxe._CallStack.CallStack_Impl_":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "Holder":
		return hxrt_typeCreateEmpty__Holder(), true
	case "haxe.Exception":
		return hxrt_typeCreateEmpty__haxe__Exception(), true
	case "haxe.ValueException":
		return hxrt_typeCreateEmpty__haxe__ValueException(), true
	case "haxe._Int64.___Int64":
		return hxrt_typeCreateEmpty__haxe___Int64_____Int64(), true
	default:
		return nil, false
	}
}

func hxrt_typeCreateEnumInstance(enumName string, constructorName string, constructorIndex int, useIndex bool, args []any) (any, bool) {
	switch enumName {
	case "ValueType":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TNull, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TInt, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFloat, true
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TBool, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TObject, true
			case 5:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFunction, true
			case 6:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TClass, args)
			case 7:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TEnum, args)
			case 8:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TUnknown, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "TNull":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TNull, true
		case "TInt":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TInt, true
		case "TFloat":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFloat, true
		case "TBool":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TBool, true
		case "TObject":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TObject, true
		case "TFunction":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFunction, true
		case "TClass":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TClass, args)
		case "TEnum":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TEnum, args)
		case "TUnknown":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TUnknown, true
		default:
			return nil, false
		}
	case "haxe.StackItem":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__StackItem_CFunction, true
			case 1:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_Module, args)
			case 2:
				if len(args) != 4 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_FilePos, args)
			case 3:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_Method, args)
			case 4:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_LocalFunction, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "CFunction":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__StackItem_CFunction, true
		case "Module":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_Module, args)
		case "FilePos":
			if len(args) != 4 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_FilePos, args)
		case "Method":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_Method, args)
		case "LocalFunction":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_LocalFunction, args)
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

func Type_getClass(o any) any {
	o = hxrt.ObjectSource(o)
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeClassValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeClassValue:
		copyValue := value
		return &copyValue
	case *hxrt.Array:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")}
	case *Holder:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Holder")}
	case *haxe__Exception:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case *haxe__ValueException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ValueException")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	default:
		return nil
	}
}

func Type_getEnum(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeEnumValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeEnumValue:
		copyValue := value
		return &copyValue
	case *ValueType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("ValueType")}
	case *haxe__StackItem:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.StackItem")}
	default:
		return nil
	}
}

func Type_getSuperClass(c any) any {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	switch className {
	case "Holder":
		return nil
	case "Main":
		return nil
	case "Math":
		return nil
	case "Reflect":
		return nil
	case "haxe.Exception":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe.ValueException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "haxe._CallStack.CallStack_Impl_":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	default:
		return nil
	}
}

func Type_getClassName(c any) *string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(className)
}

func Type_getClassFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Holder":
		return hxrt.NewArray(hxrt.StringFromLiteral("shared"))
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("arithmetic"), hxrt.StringFromLiteral("attempt"), hxrt.StringFromLiteral("bitwise"), hxrt.StringFromLiteral("comparisons"), hxrt.StringFromLiteral("compound"), hxrt.StringFromLiteral("conditions"), hxrt.StringFromLiteral("counter"), hxrt.StringFromLiteral("dyn"), hxrt.StringFromLiteral("errors"), hxrt.StringFromLiteral("evaluationOrder"), hxrt.StringFromLiteral("increments"), hxrt.StringFromLiteral("main"), hxrt.StringFromLiteral("more"), hxrt.StringFromLiteral("order"), hxrt.StringFromLiteral("step"), hxrt.StringFromLiteral("typedOperands"), hxrt.StringFromLiteral("unary"))
	case "Math":
		return hxrt.NewArray(hxrt.StringFromLiteral("NEGATIVE_INFINITY"), hxrt.StringFromLiteral("NaN"), hxrt.StringFromLiteral("PI"), hxrt.StringFromLiteral("POSITIVE_INFINITY"), hxrt.StringFromLiteral("abs"), hxrt.StringFromLiteral("acos"), hxrt.StringFromLiteral("asin"), hxrt.StringFromLiteral("atan"), hxrt.StringFromLiteral("atan2"), hxrt.StringFromLiteral("ceil"), hxrt.StringFromLiteral("cos"), hxrt.StringFromLiteral("exp"), hxrt.StringFromLiteral("fceil"), hxrt.StringFromLiteral("ffloor"), hxrt.StringFromLiteral("floor"), hxrt.StringFromLiteral("fround"), hxrt.StringFromLiteral("isFinite"), hxrt.StringFromLiteral("isNaN"), hxrt.StringFromLiteral("log"), hxrt.StringFromLiteral("max"), hxrt.StringFromLiteral("min"), hxrt.StringFromLiteral("pow"), hxrt.StringFromLiteral("random"), hxrt.StringFromLiteral("round"), hxrt.StringFromLiteral("sin"), hxrt.StringFromLiteral("sqrt"), hxrt.StringFromLiteral("tan"))
	case "Reflect":
		return hxrt.NewArray(hxrt.StringFromLiteral("callMethod"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("compareMethods"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("deleteField"), hxrt.StringFromLiteral("field"), hxrt.StringFromLiteral("fields"), hxrt.StringFromLiteral("getProperty"), hxrt.StringFromLiteral("hasField"), hxrt.StringFromLiteral("isEnumValue"), hxrt.StringFromLiteral("isFunction"), hxrt.StringFromLiteral("isObject"), hxrt.StringFromLiteral("makeVarArgs"), hxrt.StringFromLiteral("setField"), hxrt.StringFromLiteral("setProperty"))
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("catchValue"), hxrt.StringFromLiteral("caught"), hxrt.StringFromLiteral("thrown"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.ValueException":
		return hxrt.NewArray()
	case "haxe._CallStack.CallStack_Impl_":
		return hxrt.NewArray(hxrt.StringFromLiteral("asArray"), hxrt.StringFromLiteral("callStack"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("equalItems"), hxrt.StringFromLiteral("exceptionStack"), hxrt.StringFromLiteral("exceptionToString"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("get_length"), hxrt.StringFromLiteral("itemToString"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("subtract"), hxrt.StringFromLiteral("toString"))
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	default:
		return hxrt.NewArray()
	}
}

func Type_getInstanceFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Holder":
		return hxrt.NewArray(hxrt.StringFromLiteral("n"))
	case "Main":
		return hxrt.NewArray()
	case "Math":
		return hxrt.NewArray()
	case "Reflect":
		return hxrt.NewArray()
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.ValueException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"), hxrt.StringFromLiteral("value"))
	case "haxe._CallStack.CallStack_Impl_":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	default:
		return hxrt.NewArray()
	}
}

func Type_getEnumName(e any) *string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(enumName)
}

func Type_resolveClass(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "Holder":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Math":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Reflect":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Exception":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ValueException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._CallStack.CallStack_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_resolveEnum(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "ValueType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.StackItem":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_createInstance(cl any, args *hxrt.Array) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassInstance(className, hxrt_typeArrayValues(args))
	if !ok {
		return nil
	}
	return instance
}

func Type_createEmptyInstance(cl any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassEmptyInstance(className)
	if !ok {
		return nil
	}
	return instance
}

func Type_createEnum(e any, constr *string, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	constructorName := ""
	if constr != nil {
		constructorName = *hxrt.StdString(constr)
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_createEnumIndex(e any, index int, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, "", index, true, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_enumConstructor(e any) *string {
	if hxrt.AnyEqualsNull(e) {
		return nil
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("TNull")
		case 1:
			return hxrt.StringFromLiteral("TInt")
		case 2:
			return hxrt.StringFromLiteral("TFloat")
		case 3:
			return hxrt.StringFromLiteral("TBool")
		case 4:
			return hxrt.StringFromLiteral("TObject")
		case 5:
			return hxrt.StringFromLiteral("TFunction")
		case 6:
			return hxrt.StringFromLiteral("TClass")
		case 7:
			return hxrt.StringFromLiteral("TEnum")
		case 8:
			return hxrt.StringFromLiteral("TUnknown")
		default:
			return nil
		}
	case *haxe__StackItem:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("CFunction")
		case 1:
			return hxrt.StringFromLiteral("Module")
		case 2:
			return hxrt.StringFromLiteral("FilePos")
		case 3:
			return hxrt.StringFromLiteral("Method")
		case 4:
			return hxrt.StringFromLiteral("LocalFunction")
		default:
			return nil
		}
	default:
		return nil
	}
}

func Type_enumIndex(e any) int {
	if hxrt.AnyEqualsNull(e) {
		return -1
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__StackItem:
		if value == nil {
			return -1
		}
		return value.tag
	default:
		return -1
	}
}

func Type_getEnumConstructs(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown"))
	case "haxe.StackItem":
		return hxrt.NewArray(hxrt.StringFromLiteral("CFunction"), hxrt.StringFromLiteral("Module"), hxrt.StringFromLiteral("FilePos"), hxrt.StringFromLiteral("Method"), hxrt.StringFromLiteral("LocalFunction"))
	default:
		return hxrt.NewArray()
	}
}

func Type_enumParameters(e any) *hxrt.Array {
	if hxrt.AnyEqualsNull(e) {
		return hxrt.NewArray()
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__StackItem:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	default:
		return hxrt.NewArray()
	}
}

func Type_allEnums(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown)
	case "haxe.StackItem":
		return hxrt.NewArray(haxe__StackItem_CFunction)
	default:
		return hxrt.NewArray()
	}
}

func Type_typeof(v any) *ValueType {
	v = hxrt.ObjectSource(v)
	if hxrt.ReflectIsFunction(v) {
		return ValueType_TFunction
	}
	if hxrt.AnyEqualsNull(v) {
		return ValueType_TNull
	}
	switch v.(type) {
	case *hxrt__TypeClassValue, hxrt__TypeClassValue, *hxrt__TypeEnumValue, hxrt__TypeEnumValue:
		return ValueType_TObject
	}
	if enumValue := Type_getEnum(v); enumValue != nil {
		return ValueType_TEnum(enumValue)
	}
	if classValue := Type_getClass(v); classValue != nil {
		return ValueType_TClass(classValue)
	}
	switch v.(type) {
	case bool:
		return ValueType_TBool
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return ValueType_TInt
	case float32, float64:
		return ValueType_TFloat
	case string, *string:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("String")})
	case *hxrt.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	}
	ref := reflect.ValueOf(v)
	if !ref.IsValid() {
		return ValueType_TNull
	}
	switch ref.Kind() {
	case reflect.Slice, reflect.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	case reflect.Map, reflect.Struct, reflect.Interface, reflect.Pointer:
		return ValueType_TObject
	default:
		return ValueType_TUnknown
	}
}

func Type_enumEq(a any, b any) bool {
	return reflect.DeepEqual(a, b)
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
	case *Holder:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
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
	case *Holder:
		return hxrt__generated_field_lookup__Holder(value, key)
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
	case *Holder:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
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
	case *Holder:
		return hxrt__generated_field_has__Holder(value, key)
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
	case *Holder:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
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
	case *Holder:
		return hxrt__generated_field_set__Holder(value, key, incoming)
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
	case *Holder:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
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
	case *Holder:
		return hxrt.NewArray(hxrt.StringFromLiteral("n"))
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

func hxrt__generated_field_lookup__Holder(value *Holder, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "n":
		return hxrt.ReadStoredField(value.n)
	}
	return nil
}

func hxrt__generated_field_has__Holder(value *Holder, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "n":
		return true
	}
	return false
}

func hxrt__generated_field_set__Holder(value *Holder, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "n":
		if incoming == nil {
			var zero any
			value.n = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.n = typed
			return true
		default:
			return false
		}
	}
	return false
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
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() float64]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
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
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_c5309b43(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
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
