package main

import "snapshot/hxrt"

type I_Bounds interface {
}

type Bounds struct {
	__hx_this I_Bounds
	maybeEnd  any
}

func (self *Bounds) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Bounds(maybeEnd any) *Bounds {
	return hx_init_Bounds(hx_alloc_Bounds(nil), maybeEnd)
}

func hx_alloc_Bounds(owner any) *Bounds {
	self := &Bounds{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Bounds)
	}
	return self
}

func hx_init_Bounds(self *Bounds, maybeEnd any) *Bounds {
	self.maybeEnd = maybeEnd
	return self
}

func (self *Bounds) HxClassName() string {
	return "Bounds"
}

var Bounds_shared any = nil

func four(s *string, value any) *string {
	return hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringSubstrOptionalStringPtr(s, 1, value), hxrt.StringSubstringOptionalStringPtr(s, 0, value), hxrt.StdString(hxrt.StringIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), value)), hxrt.StdString(hxrt.StringLastIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), value)), hxrt.StdString(hxrt.StringIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral(""), value))).Values(), hxrt.StringFromLiteral("|"))
}

func main() {
	s := hxrt.StringFromLiteral("abcabc")
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("omitted="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringSubstrStringPtr(s, 1, 0, false), hxrt.StringSubstringStringPtr(s, 1, hxrt.StringLengthStringPtr(s)), hxrt.StringSubstringStringPtr(s, 4, hxrt.StringLengthStringPtr(s)), hxrt.StdString(hxrt.StringIndexOfStringPtr(s, hxrt.StringFromLiteral("b"), 0, false)), hxrt.StdString(hxrt.StringLastIndexOfStringPtr(s, hxrt.StringFromLiteral("b"), 0, false))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("literal.null="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringSubstrStringPtr(s, 1, 0, false), hxrt.StringSubstringStringPtr(s, 1, hxrt.StringLengthStringPtr(s)), hxrt.StdString(hxrt.StringIndexOfStringPtr(s, hxrt.StringFromLiteral("b"), 0, false)), hxrt.StdString(hxrt.StringLastIndexOfStringPtr(s, hxrt.StringFromLiteral("b"), 0, false))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("constant="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringSubstrStringPtr(s, 1, 2, true), hxrt.StringSubstringStringPtr(s, 1, 2), hxrt.StdString(hxrt.StringIndexOfStringPtr(s, hxrt.StringFromLiteral("b"), 2, true)), hxrt.StdString(hxrt.StringLastIndexOfStringPtr(s, hxrt.StringFromLiteral("b"), 2, true))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("constant.negative="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringSubstrStringPtr(s, 1, -2, true), hxrt.StringSubstringStringPtr(s, 0, -2), hxrt.StdString(hxrt.StringIndexOfStringPtr(s, hxrt.StringFromLiteral("b"), -2, true)), hxrt.StdString(hxrt.StringLastIndexOfStringPtr(s, hxrt.StringFromLiteral("b"), -2, true))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_3)
	var value any = nil
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param."), value), hxrt.StringFromLiteral("=")), four(s, value)))
	hxrt.Println(v_4)
	var value_1 any = 2
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param."), value_1), hxrt.StringFromLiteral("=")), four(s, value_1)))
	hxrt.Println(v_5)
	var value_2 any = -2
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param."), value_2), hxrt.StringFromLiteral("=")), four(s, value_2)))
	hxrt.Println(v_6)
	var value_3 any = 99
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param."), value_3), hxrt.StringFromLiteral("=")), four(s, value_3)))
	hxrt.Println(v_7)
	var value_4 any = -99
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param."), value_4), hxrt.StringFromLiteral("=")), four(s, value_4)))
	hxrt.Println(v_8)
	var value_5 any = 0
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param."), value_5), hxrt.StringFromLiteral("=")), four(s, value_5)))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("optional.omitted="), optional(s, nil)))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("optional.null="), optional(s, nil)))
	hxrt.Println(v_11)
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("optional.value="), optional(s, 3)))
	hxrt.Println(v_12)
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("narrowed="), narrowed(s, 3)), hxrt.StringFromLiteral(",")), narrowed(s, nil)))
	hxrt.Println(v_13)
	var none any = noInt()
	var some any = someInt(2)
	var v_14 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("local="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringSubstrOptionalStringPtr(s, 1, none), hxrt.StringSubstringOptionalStringPtr(s, 1, none), hxrt.StdString(hxrt.StringIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), none)), hxrt.StdString(hxrt.StringLastIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), none)), hxrt.StringSubstrOptionalStringPtr(s, 1, some), hxrt.StdString(hxrt.StringLastIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), some))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_14)
	bounds := New_Bounds(nil)
	var window_maybeEnd any
	window_maybeEnd = nil
	var v_15 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null.field="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringSubstrOptionalStringPtr(s, 1, bounds.maybeEnd), hxrt.StringSubstringOptionalStringPtr(s, 1, Bounds_shared), hxrt.StdString(hxrt.StringIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), window_maybeEnd)), hxrt.StdString(hxrt.StringLastIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), bounds.maybeEnd))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_15)
	bounds.maybeEnd = 3
	Bounds_shared = 2
	window_maybeEnd = 2
	var v_16 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("value.field="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringSubstrOptionalStringPtr(s, 1, bounds.maybeEnd), hxrt.StringSubstringOptionalStringPtr(s, 1, Bounds_shared), hxrt.StdString(hxrt.StringIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), window_maybeEnd)), hxrt.StdString(hxrt.StringLastIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), bounds.maybeEnd))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_16)
	ends := hxrt.NewArray(nil, 3)
	i := 0
	var v_17 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringSubstrOptionalStringPtr(s, 1, ends.Get(0)), hxrt.StringSubstringOptionalStringPtr(s, 1, ends.Get(i)), hxrt.StdString(hxrt.StringIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), ends.Get(i))), hxrt.StdString(hxrt.StringLastIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), ends.Get(0))), hxrt.StringSubstrOptionalStringPtr(s, 1, ends.Get(int((hxrt.Int32Wrap(i)+hxrt.Int32Wrap(1))))), hxrt.StdString(hxrt.StringIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), ends.Get(1)))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_17)
	flag := (hxrt.StringLengthStringPtr(s) > 3)
	var dynamicNull any = noInt()
	var dynamicValue any = someInt(2)
	var v_18 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("call="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringSubstrOptionalStringPtr(s, 1, noInt()), hxrt.StringSubstringOptionalStringPtr(s, 1, noInt()), hxrt.StdString(hxrt.StringIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), noInt())), hxrt.StdString(hxrt.StringLastIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), noInt())), hxrt.StringSubstringOptionalStringPtr(s, 1, someInt(3)), hxrt.StringSubstrOptionalStringPtr(s, 1, func() any {
		var hx_if_1 any
		if flag {
			hx_if_1 = nil
		} else {
			hx_if_1 = 1
		}
		return hx_if_1
	}()), hxrt.StringSubstrOptionalStringPtr(s, 1, dynamicNull), hxrt.StdString(hxrt.StringLastIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), dynamicValue))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_18)
	accented := hxrt.StringFromLiteral("héllo")
	var v_19 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("unicode="), hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringSubstrOptionalStringPtr(accented, 1, none), hxrt.StringSubstringOptionalStringPtr(accented, 1, none), hxrt.StdString(hxrt.StringIndexOfOptionalStringPtr(accented, hxrt.StringFromLiteral("l"), none)), hxrt.StdString(hxrt.StringLastIndexOfOptionalStringPtr(accented, hxrt.StringFromLiteral("l"), none))).Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_19)
	order := hxrt.NewArray(hxrt.StringSubstrOptionalStringPtr(tickText(hxrt.StringFromLiteral("a"), s), tickInt(hxrt.StringFromLiteral("b"), 1), tick(hxrt.StringFromLiteral("c"), nil)), hxrt.StringSubstringOptionalStringPtr(tickText(hxrt.StringFromLiteral("d"), s), tickInt(hxrt.StringFromLiteral("e"), 1), tick(hxrt.StringFromLiteral("f"), nil)), hxrt.StdString(hxrt.StringIndexOfOptionalStringPtr(tickText(hxrt.StringFromLiteral("g"), s), tickText(hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("b")), tick(hxrt.StringFromLiteral("i"), nil))), hxrt.StdString(hxrt.StringLastIndexOfOptionalStringPtr(tickText(hxrt.StringFromLiteral("j"), s), tickText(hxrt.StringFromLiteral("k"), hxrt.StringFromLiteral("b")), tick(hxrt.StringFromLiteral("l"), nil))), hxrt.StringSubstrOptionalStringPtr(tickText(hxrt.StringFromLiteral("m"), s), tickInt(hxrt.StringFromLiteral("n"), 1), tick(hxrt.StringFromLiteral("o"), 2)), hxrt.StringSubstringOptionalStringPtr(tickText(hxrt.StringFromLiteral("p"), s), tickInt(hxrt.StringFromLiteral("q"), 1), nil))
	var v_20 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order="), hxrt.StringJoinAny(order.Values(), hxrt.StringFromLiteral("|"))), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(trail.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_20)
}

func narrowed(s *string, value any) *string {
	if value != nil {
		return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringSubstrStringPtr(s, 1, value.(int), true), hxrt.StringFromLiteral("|")), hxrt.StringSubstringStringPtr(s, 1, value.(int))), hxrt.StringFromLiteral("|")), hxrt.StringIndexOfStringPtr(s, hxrt.StringFromLiteral("b"), value.(int), true)), hxrt.StringFromLiteral("|")), hxrt.StringLastIndexOfStringPtr(s, hxrt.StringFromLiteral("b"), value.(int), true))
	}
	return hxrt.StringFromLiteral("missing")
}

func noInt() any {
	return nil
}

func optional(s *string, value any) *string {
	return hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringSubstrOptionalStringPtr(s, 1, value), hxrt.StringSubstringOptionalStringPtr(s, 1, value), hxrt.StdString(hxrt.StringIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), value)), hxrt.StdString(hxrt.StringLastIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral("b"), value)), hxrt.StdString(hxrt.StringLastIndexOfOptionalStringPtr(s, hxrt.StringFromLiteral(""), value))).Values(), hxrt.StringFromLiteral("|"))
}

func someInt(value int) any {
	return value
}

func tick(label *string, value any) any {
	hx_arr_2 := trail
	hx_arr_2.Push(label)
	return value
}

func tickInt(label *string, value int) int {
	hx_arr_3 := trail
	hx_arr_3.Push(label)
	return value
}

func tickText(label *string, value *string) *string {
	hx_arr_4 := trail
	hx_arr_4.Push(label)
	return value
}

var trail *hxrt.Array = hxrt.NewArray()
