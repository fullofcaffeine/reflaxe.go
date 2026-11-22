package main

import "snapshot/hxrt"

type I_Counter interface {
}

type Counter struct {
	__hx_this  I_Counter
	maybeCount any
	maybeRatio any
}

func (self *Counter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Counter() *Counter {
	return hx_init_Counter(hx_alloc_Counter(nil))
}

func hx_alloc_Counter(owner any) *Counter {
	self := &Counter{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Counter)
	}
	return self
}

func hx_init_Counter(self *Counter) *Counter {
	self.maybeRatio = nil
	self.maybeCount = nil
	return self
}

func (self *Counter) HxClassName() string {
	return "Counter"
}

var Counter_spare any = nil

func flag(value bool) *string {
	var hx_if_1 *string
	if value {
		hx_if_1 = hxrt.StringFromLiteral("t")
	} else {
		hx_if_1 = hxrt.StringFromLiteral("f")
	}
	return hx_if_1
}

func intInt(a any, b int) *string {
	return hxrt.StringJoinAny(func() *hxrt.Array {
		f := hxrt.NamedFunction(flag, "Main.flag")
		_g := hxrt.NewArray()
		_g1 := 0
		_g2 := hxrt.NewArray(hxrt.NullableIntLess(a, b), hxrt.NullableIntLessEqual(a, b), hxrt.NullableIntGreater(a, b), hxrt.NullableIntGreaterEqual(a, b), hxrt.IntLessNullable(b, a), hxrt.IntLessEqualNullable(b, a), hxrt.IntGreaterNullable(b, a), hxrt.IntGreaterEqualNullable(b, a))
		for _g1 < _g2.Len() {
			v := func(hx_value_2 any) bool {
				if hx_value_2 == nil {
					var hx_zero_3 bool
					return hx_zero_3
				}
				return hx_value_2.(bool)
			}(_g2.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func() *string {
				var hx_callable_5 *hxrt.Function[func(bool) *string] = f
				var hx_argument_6 bool = v
				return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
			}())
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(""))
}

func intReal(a any, b float64) *string {
	return hxrt.StringJoinAny(func() *hxrt.Array {
		f := hxrt.NamedFunction(flag, "Main.flag")
		_g := hxrt.NewArray()
		_g1 := 0
		_g2 := hxrt.NewArray((func(hx_value_7 any) float64 {
			if hx_value_7 == nil {
				var hx_zero_8 float64
				return (0 / hx_zero_8)
			}
			return float64(hxrt.IntFromNullableAny(hx_value_7))
		}(a) < b), (func(hx_value_9 any) float64 {
			if hx_value_9 == nil {
				var hx_zero_10 float64
				return (0 / hx_zero_10)
			}
			return float64(hxrt.IntFromNullableAny(hx_value_9))
		}(a) <= b), (func(hx_value_11 any) float64 {
			if hx_value_11 == nil {
				var hx_zero_12 float64
				return (0 / hx_zero_12)
			}
			return float64(hxrt.IntFromNullableAny(hx_value_11))
		}(a) > b), (func(hx_value_13 any) float64 {
			if hx_value_13 == nil {
				var hx_zero_14 float64
				return (0 / hx_zero_14)
			}
			return float64(hxrt.IntFromNullableAny(hx_value_13))
		}(a) >= b), (b < func(hx_value_15 any) float64 {
			if hx_value_15 == nil {
				var hx_zero_16 float64
				return (0 / hx_zero_16)
			}
			return float64(hxrt.IntFromNullableAny(hx_value_15))
		}(a)), (b <= func(hx_value_17 any) float64 {
			if hx_value_17 == nil {
				var hx_zero_18 float64
				return (0 / hx_zero_18)
			}
			return float64(hxrt.IntFromNullableAny(hx_value_17))
		}(a)), (b > func(hx_value_19 any) float64 {
			if hx_value_19 == nil {
				var hx_zero_20 float64
				return (0 / hx_zero_20)
			}
			return float64(hxrt.IntFromNullableAny(hx_value_19))
		}(a)), (b >= func(hx_value_21 any) float64 {
			if hx_value_21 == nil {
				var hx_zero_22 float64
				return (0 / hx_zero_22)
			}
			return float64(hxrt.IntFromNullableAny(hx_value_21))
		}(a)))
		for _g1 < _g2.Len() {
			v := func(hx_value_23 any) bool {
				if hx_value_23 == nil {
					var hx_zero_24 bool
					return hx_zero_24
				}
				return hx_value_23.(bool)
			}(_g2.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func() *string {
				var hx_callable_26 *hxrt.Function[func(bool) *string] = f
				var hx_argument_27 bool = v
				return hxrt.FunctionCall(hx_callable_26)(hx_argument_27)
			}())
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(""))
}

func main() {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("int.null="), intInt(nil, 1)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("int.null.zero="), intInt(nil, 0)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("int.null.negative="), intInt(nil, -1)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("int.value="), intInt(2, 1)))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("int.equal="), intInt(1, 1)))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ints.both.null="), nullInts(nil, nil)))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ints.left.null="), nullInts(nil, 0)))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ints.right.null="), nullInts(0, nil)))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ints.value="), nullInts(1, 2)))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ints.equal="), nullInts(2, 2)))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("reals.both.null="), nullReals(nil, nil)))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("reals.left.null="), nullReals(nil, 1.5)))
	hxrt.Println(v_11)
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("reals.value="), nullReals(1.5, 2.5)))
	hxrt.Println(v_12)
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("reals.nan="), nullReals(Math_NaN, Math_NaN)))
	hxrt.Println(v_13)
	var v_14 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("reals.nan.null="), nullReals(Math_NaN, nil)))
	hxrt.Println(v_14)
	var v_15 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mixed.both.null="), nullMixed(nil, nil)))
	hxrt.Println(v_15)
	var v_16 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mixed.left.null="), nullMixed(nil, 2.5)))
	hxrt.Println(v_16)
	var v_17 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mixed.right.null="), nullMixed(2, nil)))
	hxrt.Println(v_17)
	var v_18 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mixed.value="), nullMixed(2, 2.5)))
	hxrt.Println(v_18)
	var v_19 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mixed.equal="), nullMixed(2, 2.0)))
	hxrt.Println(v_19)
	var v_20 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("int.real.null="), intReal(nil, 1.5)))
	hxrt.Println(v_20)
	var v_21 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("real.int.null="), realInt(nil, 1)))
	hxrt.Println(v_21)
	var v_22 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("real.int.value="), realInt(1.5, 1)))
	hxrt.Println(v_22)
	var v_23 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("optional.null="), optional(nil)))
	hxrt.Println(v_23)
	var v_24 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("optional.value="), optional(0)))
	hxrt.Println(v_24)
	var none any = noWhole()
	var other any = noWhole()
	var one any = someWhole(1)
	x := 1
	row(hxrt.StringFromLiteral("local"), hxrt.NewArray(hxrt.NullableIntLess(none, x), hxrt.IntGreaterNullable(x, none), hxrt.NullableIntLessEqual(none, 1), hxrt.IntGreaterEqualNullable(1, none), hxrt.NullableIntGreaterEqual(none, -1), hxrt.IntLessNullable(0, none), hxrt.NullableIntsLess(none, other), hxrt.NullableIntsLessEqual(none, other), hxrt.NullableIntsGreater(none, other), hxrt.NullableIntsGreaterEqual(none, other), hxrt.NullableIntsLess(one, none), hxrt.NullableIntsLessEqual(none, one), hxrt.NullableIntLessEqual(one, x), hxrt.NullableIntGreater(one, 0)))
	counter := New_Counter()
	var sample_maybeWhole any
	var sample_maybeReal any
	sample_maybeWhole = nil
	sample_maybeReal = nil
	row(hxrt.StringFromLiteral("null.field"), hxrt.NewArray(hxrt.NullableIntLess(counter.maybeCount, 1), hxrt.NullableIntGreaterEqual(Counter_spare, x), hxrt.NullableIntLessEqual(sample_maybeWhole, 0), hxrt.NullableIntsLessEqual(counter.maybeCount, Counter_spare), hxrt.NullableIntsGreaterEqual(Counter_spare, sample_maybeWhole), hxrt.NullableIntsLess(counter.maybeCount, Counter_spare), hxrt.NullableFloatsLessEqual(counter.maybeRatio, sample_maybeReal), hxrt.NullableFloatsGreaterEqual(sample_maybeReal, counter.maybeCount), hxrt.NullableFloatsGreater(sample_maybeReal, counter.maybeRatio)))
	counter.maybeCount = 2
	Counter_spare = 3
	sample_maybeWhole = 2
	sample_maybeReal = 2.5
	row(hxrt.StringFromLiteral("value.field"), hxrt.NewArray(hxrt.NullableIntLess(counter.maybeCount, 3), hxrt.NullableIntGreaterEqual(Counter_spare, x), hxrt.NullableIntLessEqual(sample_maybeWhole, 2), hxrt.NullableIntsLessEqual(counter.maybeCount, Counter_spare), hxrt.NullableIntsGreaterEqual(Counter_spare, sample_maybeWhole), hxrt.NullableIntsLess(counter.maybeCount, sample_maybeWhole), hxrt.NullableFloatsLessEqual(counter.maybeRatio, sample_maybeReal), hxrt.NullableFloatsGreaterEqual(sample_maybeReal, counter.maybeCount), hxrt.NullableFloatsGreater(sample_maybeReal, counter.maybeRatio)))
	maybeInts := hxrt.NewArray(nil, 2, nil)
	var maybeReals_1 any
	var maybeReals_0 any
	maybeReals_0 = nil
	maybeReals_1 = 2.5
	i := 1
	row(hxrt.StringFromLiteral("array"), hxrt.NewArray(hxrt.NullableIntLess(maybeInts.Get(0), 1), hxrt.IntGreaterNullable(1, maybeInts.Get(0)), hxrt.NullableIntGreaterEqual(maybeInts.Get(1), 2), hxrt.NullableIntsLessEqual(maybeInts.Get(0), maybeInts.Get(2)), hxrt.NullableIntsLess(maybeInts.Get(0), maybeInts.Get(2)), hxrt.NullableIntsGreater(maybeInts.Get(1), maybeInts.Get(0)), hxrt.NullableIntsGreaterEqual(maybeInts.Get(i), maybeInts.Get(int((hxrt.Int32Wrap(i)+hxrt.Int32Wrap(1))))), hxrt.NullableFloatsGreaterEqual(maybeInts.Get(0), maybeReals_0), hxrt.NullableFloatsLessEqual(maybeReals_0, maybeInts.Get(2)), hxrt.NullableFloatsLess(maybeInts.Get(1), maybeReals_1), hxrt.NullableFloatsGreaterEqual(maybeReals_0, maybeReals_0)))
	row(hxrt.StringFromLiteral("call"), hxrt.NewArray(hxrt.NullableIntLess(noWhole(), 1), hxrt.IntGreaterEqualNullable(0, noWhole()), hxrt.NullableIntsLessEqual(noWhole(), noWhole()), hxrt.NullableIntsGreater(noWhole(), noWhole()), hxrt.NullableIntsGreaterEqual(someWhole(2), noWhole()), hxrt.NullableFloatsGreaterEqual(noWhole(), noReal()), hxrt.NullableFloatsLessEqual(noReal(), noReal()), hxrt.NullableIntsLess(someWhole(1), someWhole(2)), hxrt.IntGreaterEqualNullable(2, someWhole(2))))
	var v_25 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("narrowed="), narrowed(2, nil)), hxrt.StringFromLiteral(",")), narrowed(2, 3)), hxrt.StringFromLiteral(",")), narrowed(nil, 1)))
	hxrt.Println(v_25)
	narrow := 5
	row(hxrt.StringFromLiteral("int32"), hxrt.NewArray(hxrt.IntLessNullable(narrow, none), hxrt.NullableIntGreaterEqual(none, narrow), hxrt.IntGreaterNullable(narrow, one), hxrt.NullableIntLessEqual(one, narrow)))
	frames := hxrt.NewArray(haxe__StackItem_FilePos(nil, hxrt.StringFromLiteral("Main.hx"), 3, nil), haxe__StackItem_FilePos(haxe__StackItem_Method(hxrt.StringFromLiteral("Main"), hxrt.StringFromLiteral("run")), hxrt.StringFromLiteral("Main.hx"), 7, 2))
	var v_26 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("stack="), StringTools_replace(haxe___CallStack__CallStack_Impl__toString(frames), hxrt.StringFromLiteral("\n"), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_26)
	taken := hxrt.NewArray()
	if hxrt.NullableIntLess(none, 3) {
		taken.Push(hxrt.StringFromLiteral("never"))
	} else {
		taken.Push(hxrt.StringFromLiteral("else"))
	}
	var hx_if_31 *string
	if hxrt.NullableIntsLessEqual(none, other) {
		hx_if_31 = hxrt.StringFromLiteral("both")
	} else {
		hx_if_31 = hxrt.StringFromLiteral("never")
	}
	taken.Push(hx_if_31)
	var hx_if_33 *string
	if hxrt.NullableIntsLess(none, other) {
		hx_if_33 = hxrt.StringFromLiteral("never")
	} else {
		hx_if_33 = hxrt.StringFromLiteral("strict")
	}
	taken.Push(hx_if_33)
	var hx_if_35 *string
	if hxrt.NullableIntsGreaterEqual(one, none) || hxrt.NullableIntsGreaterEqual(none, other) {
		hx_if_35 = hxrt.StringFromLiteral("or")
	} else {
		hx_if_35 = hxrt.StringFromLiteral("never")
	}
	taken.Push(hx_if_35)
	steps := 0
	var limit any = nil
	for hxrt.IntLessNullable(steps, limit) {
		steps = int(int32((steps + 1)))
	}
	var counted any = 0
	for hxrt.NullableIntLess(counted, 3) {
		counted = int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(counted))) + hxrt.Int32Wrap(1)))
	}
	stored := hxrt.NullableIntsGreaterEqual(none, other)
	var v_27 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("flow="), hxrt.StringJoinAny(taken.Values(), hxrt.StringFromLiteral(","))), hxrt.StringFromLiteral(":")), steps), hxrt.StringFromLiteral(":")), counted), hxrt.StringFromLiteral(":")), flag(stored)))
	hxrt.Println(v_27)
	order := hxrt.NewArray(hxrt.NullableIntLess(tick(hxrt.StringFromLiteral("a"), nil), tickWhole(hxrt.StringFromLiteral("b"), 1)), hxrt.IntGreaterEqualNullable(tickWhole(hxrt.StringFromLiteral("c"), 1), tick(hxrt.StringFromLiteral("d"), nil)), hxrt.NullableIntsLessEqual(tick(hxrt.StringFromLiteral("e"), nil), tick(hxrt.StringFromLiteral("f"), nil)), hxrt.NullableIntsGreater(tick(hxrt.StringFromLiteral("g"), nil), tick(hxrt.StringFromLiteral("h"), nil)), hxrt.NullableFloatsGreaterEqual(tickReal(hxrt.StringFromLiteral("i"), nil), tickReal(hxrt.StringFromLiteral("j"), nil)), hxrt.NullableFloatsLessEqual(tick(hxrt.StringFromLiteral("k"), nil), tickReal(hxrt.StringFromLiteral("l"), nil)), hxrt.NullableIntsLess(tick(hxrt.StringFromLiteral("m"), 1), tick(hxrt.StringFromLiteral("n"), 2)), hxrt.IntLessEqualNullable(tickWhole(hxrt.StringFromLiteral("o"), 2), tick(hxrt.StringFromLiteral("p"), 2)))
	var v_28 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order="), hxrt.StringJoinAny(func() *hxrt.Array {
		f := hxrt.NamedFunction(flag, "Main.flag")
		_g := hxrt.NewArray()
		_g1 := 0
		_g2 := order
		for _g1 < _g2.Len() {
			v_29 := func(hx_value_36 any) bool {
				if hx_value_36 == nil {
					var hx_zero_37 bool
					return hx_zero_37
				}
				return hx_value_36.(bool)
			}(_g2.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func() *string {
				var hx_callable_39 *hxrt.Function[func(bool) *string] = f
				var hx_argument_40 bool = v_29
				return hxrt.FunctionCall(hx_callable_39)(hx_argument_40)
			}())
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(""))), hxrt.StringFromLiteral(":")), hxrt.StringJoinAny(trail.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_28)
}

func narrowed(value any, other any) *string {
	if value != nil {
		plain := value.(int)
		return hxrt.StringJoinAny(func() *hxrt.Array {
			f := hxrt.NamedFunction(flag, "Main.flag")
			_g := hxrt.NewArray()
			_g1 := 0
			_g2 := hxrt.NewArray((plain > 1), hxrt.IntLessEqualNullable(plain, other), hxrt.NullableIntGreaterEqual(other, plain), (value.(int) < 3), hxrt.IntGreaterEqualNullable(value.(int), other), hxrt.NullableIntGreater(other, value.(int)))
			for _g1 < _g2.Len() {
				v := func(hx_value_41 any) bool {
					if hx_value_41 == nil {
						var hx_zero_42 bool
						return hx_zero_42
					}
					return hx_value_41.(bool)
				}(_g2.Get(_g1))
				_g1 = int(int32((_g1 + 1)))
				_g.Push(func() *string {
					var hx_callable_44 *hxrt.Function[func(bool) *string] = f
					var hx_argument_45 bool = v
					return hxrt.FunctionCall(hx_callable_44)(hx_argument_45)
				}())
			}
			return _g
		}().Values(), hxrt.StringFromLiteral(""))
	}
	return hxrt.StringFromLiteral("missing")
}

func noReal() any {
	return nil
}

func noWhole() any {
	return nil
}

func nullInts(a any, b any) *string {
	return hxrt.StringJoinAny(func() *hxrt.Array {
		f := hxrt.NamedFunction(flag, "Main.flag")
		_g := hxrt.NewArray()
		_g1 := 0
		_g2 := hxrt.NewArray(hxrt.NullableIntsLess(a, b), hxrt.NullableIntsLessEqual(a, b), hxrt.NullableIntsGreater(a, b), hxrt.NullableIntsGreaterEqual(a, b), hxrt.NullableIntsLess(b, a), hxrt.NullableIntsLessEqual(b, a), hxrt.NullableIntsGreater(b, a), hxrt.NullableIntsGreaterEqual(b, a))
		for _g1 < _g2.Len() {
			v := func(hx_value_46 any) bool {
				if hx_value_46 == nil {
					var hx_zero_47 bool
					return hx_zero_47
				}
				return hx_value_46.(bool)
			}(_g2.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func() *string {
				var hx_callable_49 *hxrt.Function[func(bool) *string] = f
				var hx_argument_50 bool = v
				return hxrt.FunctionCall(hx_callable_49)(hx_argument_50)
			}())
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(""))
}

func nullMixed(a any, b any) *string {
	return hxrt.StringJoinAny(func() *hxrt.Array {
		f := hxrt.NamedFunction(flag, "Main.flag")
		_g := hxrt.NewArray()
		_g1 := 0
		_g2 := hxrt.NewArray(hxrt.NullableFloatsLess(a, b), hxrt.NullableFloatsLessEqual(a, b), hxrt.NullableFloatsGreater(a, b), hxrt.NullableFloatsGreaterEqual(a, b), hxrt.NullableFloatsLess(b, a), hxrt.NullableFloatsLessEqual(b, a), hxrt.NullableFloatsGreater(b, a), hxrt.NullableFloatsGreaterEqual(b, a))
		for _g1 < _g2.Len() {
			v := func(hx_value_51 any) bool {
				if hx_value_51 == nil {
					var hx_zero_52 bool
					return hx_zero_52
				}
				return hx_value_51.(bool)
			}(_g2.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func() *string {
				var hx_callable_54 *hxrt.Function[func(bool) *string] = f
				var hx_argument_55 bool = v
				return hxrt.FunctionCall(hx_callable_54)(hx_argument_55)
			}())
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(""))
}

func nullReals(a any, b any) *string {
	return hxrt.StringJoinAny(func() *hxrt.Array {
		f := hxrt.NamedFunction(flag, "Main.flag")
		_g := hxrt.NewArray()
		_g1 := 0
		_g2 := hxrt.NewArray(hxrt.NullableFloatsLess(a, b), hxrt.NullableFloatsLessEqual(a, b), hxrt.NullableFloatsGreater(a, b), hxrt.NullableFloatsGreaterEqual(a, b), hxrt.NullableFloatsLess(b, a), hxrt.NullableFloatsLessEqual(b, a), hxrt.NullableFloatsGreater(b, a), hxrt.NullableFloatsGreaterEqual(b, a))
		for _g1 < _g2.Len() {
			v := func(hx_value_56 any) bool {
				if hx_value_56 == nil {
					var hx_zero_57 bool
					return hx_zero_57
				}
				return hx_value_56.(bool)
			}(_g2.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func() *string {
				var hx_callable_59 *hxrt.Function[func(bool) *string] = f
				var hx_argument_60 bool = v
				return hxrt.FunctionCall(hx_callable_59)(hx_argument_60)
			}())
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(""))
}

func optional(limit any) *string {
	return hxrt.StringJoinAny(func() *hxrt.Array {
		f := hxrt.NamedFunction(flag, "Main.flag")
		_g := hxrt.NewArray()
		_g1 := 0
		_g2 := hxrt.NewArray(hxrt.NullableIntLess(limit, 1), hxrt.NullableIntLessEqual(limit, 0), hxrt.NullableIntGreater(limit, -1), hxrt.NullableIntGreaterEqual(limit, 0), hxrt.IntGreaterNullable(1, limit))
		for _g1 < _g2.Len() {
			v := func(hx_value_61 any) bool {
				if hx_value_61 == nil {
					var hx_zero_62 bool
					return hx_zero_62
				}
				return hx_value_61.(bool)
			}(_g2.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func() *string {
				var hx_callable_64 *hxrt.Function[func(bool) *string] = f
				var hx_argument_65 bool = v
				return hxrt.FunctionCall(hx_callable_64)(hx_argument_65)
			}())
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(""))
}

func realInt(a any, b int) *string {
	return hxrt.StringJoinAny(func() *hxrt.Array {
		f := hxrt.NamedFunction(flag, "Main.flag")
		_g := hxrt.NewArray()
		_g1 := 0
		_g2 := hxrt.NewArray((func(hx_value_66 any) float64 {
			if hx_value_66 == nil {
				var hx_zero_67 float64
				return (0 / hx_zero_67)
			}
			return hx_value_66.(float64)
		}(a) < float64(b)), (func(hx_value_68 any) float64 {
			if hx_value_68 == nil {
				var hx_zero_69 float64
				return (0 / hx_zero_69)
			}
			return hx_value_68.(float64)
		}(a) <= float64(b)), (func(hx_value_70 any) float64 {
			if hx_value_70 == nil {
				var hx_zero_71 float64
				return (0 / hx_zero_71)
			}
			return hx_value_70.(float64)
		}(a) > float64(b)), (func(hx_value_72 any) float64 {
			if hx_value_72 == nil {
				var hx_zero_73 float64
				return (0 / hx_zero_73)
			}
			return hx_value_72.(float64)
		}(a) >= float64(b)), (float64(b) < func(hx_value_74 any) float64 {
			if hx_value_74 == nil {
				var hx_zero_75 float64
				return (0 / hx_zero_75)
			}
			return hx_value_74.(float64)
		}(a)), (float64(b) <= func(hx_value_76 any) float64 {
			if hx_value_76 == nil {
				var hx_zero_77 float64
				return (0 / hx_zero_77)
			}
			return hx_value_76.(float64)
		}(a)), (float64(b) > func(hx_value_78 any) float64 {
			if hx_value_78 == nil {
				var hx_zero_79 float64
				return (0 / hx_zero_79)
			}
			return hx_value_78.(float64)
		}(a)), (float64(b) >= func(hx_value_80 any) float64 {
			if hx_value_80 == nil {
				var hx_zero_81 float64
				return (0 / hx_zero_81)
			}
			return hx_value_80.(float64)
		}(a)))
		for _g1 < _g2.Len() {
			v := func(hx_value_82 any) bool {
				if hx_value_82 == nil {
					var hx_zero_83 bool
					return hx_zero_83
				}
				return hx_value_82.(bool)
			}(_g2.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func() *string {
				var hx_callable_85 *hxrt.Function[func(bool) *string] = f
				var hx_argument_86 bool = v
				return hxrt.FunctionCall(hx_callable_85)(hx_argument_86)
			}())
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(""))
}

func row(label *string, values *hxrt.Array) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), hxrt.StringJoinAny(func() *hxrt.Array {
		f := hxrt.NamedFunction(flag, "Main.flag")
		_g := hxrt.NewArray()
		_g1 := 0
		_g2 := values
		for _g1 < _g2.Len() {
			v_1 := func(hx_value_87 any) bool {
				if hx_value_87 == nil {
					var hx_zero_88 bool
					return hx_zero_88
				}
				return hx_value_87.(bool)
			}(_g2.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(func() *string {
				var hx_callable_90 *hxrt.Function[func(bool) *string] = f
				var hx_argument_91 bool = v_1
				return hxrt.FunctionCall(hx_callable_90)(hx_argument_91)
			}())
		}
		return _g
	}().Values(), hxrt.StringFromLiteral(""))))
	hxrt.Println(v)
}

func someWhole(value int) any {
	return value
}

func tick(label *string, value any) any {
	hx_arr_92 := trail
	hx_arr_92.Push(label)
	return value
}

func tickReal(label *string, value any) any {
	hx_arr_93 := trail
	hx_arr_93.Push(label)
	return value
}

func tickWhole(label *string, value int) int {
	hx_arr_94 := trail
	hx_arr_94.Push(label)
	return value
}

var trail *hxrt.Array = hxrt.NewArray()

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
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
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
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
