package main

import "snapshot/hxrt"

type Wrap struct {
	tag    int
	params []any
}

func (self *Wrap) HxEnumValue() (string, []any) {
	return [1]string{"W"}[self.tag], self.params
}

func Wrap_W(n *string, c *string) *Wrap {
	enumValue := &Wrap{tag: 0}
	enumValue.params = []any{n, c}
	return enumValue
}

func aliasText(a *string) *string {
	n := a
	return n
}

func boxText(b *string) *string {
	return b
}

func codeText(c *string) *string {
	return c
}

func colorLetter(c *string) *string {
	var hx_switch_1 *string
	switch *hxrt.StdString(c) {
	case *hxrt.StdString(hxrt.StringFromLiteral("blue")):
		hx_switch_1 = hxrt.StringFromLiteral("B")
	case *hxrt.StdString(hxrt.StringFromLiteral("green")):
		hx_switch_1 = hxrt.StringFromLiteral("G")
	case *hxrt.StdString(hxrt.StringFromLiteral("red")):
		hx_switch_1 = hxrt.StringFromLiteral("R")
	}
	return hx_switch_1
}

func compareNames(x *string, y *string) int {
	a := x
	b := y
	var hx_if_3 int
	if hxrt.StringCompareStringPtr(a, b) < 0 {
		hx_if_3 = -1
	} else {
		var hx_if_2 int
		if hxrt.StringCompareStringPtr(a, b) > 0 {
			hx_if_2 = 1
		} else {
			hx_if_2 = 0
		}
		hx_if_3 = hx_if_2
	}
	return hx_if_3
}

func countOf(c any) int {
	return hxrt.IntFromNullableAny(c)
}

func defColor(c *string) *string {
	if c == nil {
		c = hxrt.StringFromLiteral("green")
	}
	return c
}

func defMode(m *string) *string {
	if m == nil {
		m = hxrt.StringFromLiteral("slow")
	}
	return m
}

func defName(n *string) *string {
	if n == nil {
		n = hxrt.StringFromLiteral("dflt")
	}
	return n
}

func firstOf(xs *hxrt.Array) any {
	return xs.Get(0)
}

func flagOf(f any) bool {
	return func(hx_value_4 any) bool {
		if hx_value_4 == nil {
			var hx_zero_5 bool
			return hx_zero_5
		}
		return hx_value_4.(bool)
	}(f)
}

func idName(n *string) *string {
	return n
}

func labelText(l *string) *string {
	return l
}

func levelOf(l any) int {
	return hxrt.IntFromNullableAny(l)
}

func looseText(l any) *string {
	return hxrt.StdString(l)
}

func main() {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("param.plain="), nameToString(hxrt.StringFromLiteral("alpha"))), hxrt.StringFromLiteral(",")), nameKind(hxrt.StringFromLiteral("alpha"))), hxrt.StringFromLiteral(",")), nameKind(hxrt.StringFromLiteral("beta"))), hxrt.StringFromLiteral(",")), nameKind(hxrt.StringFromLiteral("zeta"))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param.id="), idName(hxrt.StringFromLiteral("beta"))), hxrt.StringFromLiteral(",")), stringToName(hxrt.StringFromLiteral("gamma"))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("param.optional="), optName(nil)), hxrt.StringFromLiteral(",")), optName(nil)), hxrt.StringFromLiteral(",")), optName(hxrt.StringFromLiteral("set"))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("param.default="), defName(hxrt.StringFromLiteral("dflt"))), hxrt.StringFromLiteral(",")), defName(nil)), hxrt.StringFromLiteral(",")), defName(hxrt.StringFromLiteral("given"))), hxrt.StringFromLiteral(",")), defName(hxrt.StringFromLiteral(""))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("param.optionalDefault="), optDefName(hxrt.StringFromLiteral("odflt"))), hxrt.StringFromLiteral(",")), optDefName(nil)), hxrt.StringFromLiteral(",")), optDefName(hxrt.StringFromLiteral("given"))))
	hxrt.Println(v_4)
	var noName *string = nil
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("param.nullable="), nullName(noName)), hxrt.StringFromLiteral(",")), nullName(hxrt.StringFromLiteral("here"))), hxrt.StringFromLiteral(",")), hxrt.StdString((noName == nil))))
	hxrt.Println(v_5)
	local := hxrt.StringFromLiteral("local")
	copy := local
	back := copy
	reassigned := hxrt.StringFromLiteral("first")
	reassigned = hxrt.StringFromLiteral("second")
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("local="), local), hxrt.StringFromLiteral(",")), copy), hxrt.StringFromLiteral(",")), back), hxrt.StringFromLiteral(",")), reassigned), hxrt.StringFromLiteral(",")), hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral(""), local), hxrt.StringFromLiteral("!"))), hxrt.StringFromLiteral(",")), hxrt.StdString(StringTools_startsWith(local, hxrt.StringFromLiteral("lo")))))
	hxrt.Println(v_6)
	joined := hxrt.StringFromLiteral("ab1")
	literal := hxrt.StringFromLiteral("ab1")
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("equality="), hxrt.StdString(hxrt.StringEqualStringPtr(joined, literal))), hxrt.StringFromLiteral(",")), hxrt.StdString(!hxrt.StringEqualStringPtr(joined, literal))), hxrt.StringFromLiteral(",")), hxrt.StdString(hxrt.StringEqualStringPtr(local, hxrt.StringFromLiteral("local")))), hxrt.StringFromLiteral(",")), hxrt.StdString(hxrt.StringEqualStringPtr(hxrt.StringFromLiteral("local"), local))), hxrt.StringFromLiteral(",")), hxrt.StdString(hxrt.StringEqualStringPtr(local, copy))))
	hxrt.Println(v_7)
	p := New_Person(hxrt.StringFromLiteral("ann"), nil, hxrt.StringFromLiteral("red"))
	q := New_Person(hxrt.StringFromLiteral("bob"), hxrt.StringFromLiteral("bobby"), hxrt.StringFromLiteral("blue"))
	p.name = hxrt.StringFromLiteral("anna")
	Person_registry = hxrt.StringFromLiteral("people")
	fieldText := p.name
	var v_8 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field="), fieldText), hxrt.StringFromLiteral(",")), p.nick), hxrt.StringFromLiteral(",")), p.color), hxrt.StringFromLiteral(",")), q.name), hxrt.StringFromLiteral(",")), q.nick), hxrt.StringFromLiteral(",")), q.color), hxrt.StringFromLiteral(",")), Person_registry), hxrt.StringFromLiteral(",")), Person_lastColor))
	hxrt.Println(v_8)
	var namer Namer = q
	var v_9 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("interface="), namer.rename(hxrt.StringFromLiteral("x"), nil)), hxrt.StringFromLiteral(",")), namer.rename(hxrt.StringFromLiteral("x"), hxrt.StringFromLiteral("-y"))), hxrt.StringFromLiteral(",")), namer.rename(hxrt.StringFromLiteral("x"), nil)))
	hxrt.Println(v_9)
	names := hxrt.NewArray(hxrt.StringFromLiteral("c"), hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"))
	names.Push(hxrt.StringFromLiteral("d"))
	hxrt.ArraySort(names, func(hx_comparator_9 *hxrt.Function[func(*string, *string) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_9, func(hx_cmp_left_7 any, hx_cmp_right_8 any) int {
			return hxrt.FunctionCall(hx_comparator_9)(func(hx_value_10 any) *string {
				if hx_value_10 == nil {
					var hx_zero_11 *string
					return hx_zero_11
				}
				return hx_value_10.(*string)
			}(hx_cmp_left_7), func(hx_value_12 any) *string {
				if hx_value_12 == nil {
					var hx_zero_13 *string
					return hx_zero_13
				}
				return hx_value_12.(*string)
			}(hx_cmp_right_8))
		})
	}(hxrt.NamedFunction(compareNames, "Main.compareNames")))
	first := func(hx_value_14 any) *string {
		if hx_value_14 == nil {
			var hx_zero_15 *string
			return hx_zero_15
		}
		return hx_value_14.(*string)
	}(names.Get(0))
	firstText := func(hx_value_16 any) *string {
		if hx_value_16 == nil {
			var hx_zero_17 *string
			return hx_zero_17
		}
		return hx_value_16.(*string)
	}(names.Get(1))
	generic := func(hx_value_18 any) *string {
		if hx_value_18 == nil {
			var hx_zero_19 *string
			return hx_zero_19
		}
		return hx_value_18.(*string)
	}(firstOf(names))
	_g := hxrt.NewArray()
	_g1 := 0
	_g2 := names
	for _g1 < _g2.Len() {
		v_10 := func(hx_value_20 any) *string {
			if hx_value_20 == nil {
				var hx_zero_21 *string
				return hx_zero_21
			}
			return hx_value_20.(*string)
		}(_g2.Get(_g1))
		_g1 = int(int32((_g1 + 1)))
		_g.Push(hxrt.StringLengthStringPtr(v_10))
	}
	lengths := _g
	_g_1 := hxrt.NewArray()
	_g1_1 := 0
	_g2_1 := names
	for _g1_1 < _g2_1.Len() {
		v_11 := func(hx_value_23 any) *string {
			if hx_value_23 == nil {
				var hx_zero_24 *string
				return hx_zero_24
			}
			return hx_value_23.(*string)
		}(_g2_1.Get(_g1_1))
		_g1_1 = int(int32((_g1_1 + 1)))
		if !hxrt.StringEqualStringPtr(v_11, hxrt.StringFromLiteral("a")) {
			_g_1.Push(v_11)
		}
	}
	kept := _g_1
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), hxrt.StringJoinAny(names.Values(), hxrt.StringFromLiteral("|"))), hxrt.StringFromLiteral(",")), names.Len()), hxrt.StringFromLiteral(",")), first), hxrt.StringFromLiteral(",")), firstText), hxrt.StringFromLiteral(",")), generic), hxrt.StringFromLiteral(",")), func() int {
		hx_indexof_target_26 := names
		var hx_indexof_value_27 *string = hxrt.StringFromLiteral("b")
		hx_indexof_length_28 := hx_indexof_target_26.Len()
		hx_indexof_index_29 := 0
		for hx_indexof_index_29 < hx_indexof_length_28 {
			hx_indexof_element_30 := func(hx_value_31 any) *string {
				if hx_value_31 == nil {
					var hx_zero_32 *string
					return hx_zero_32
				}
				return hx_value_31.(*string)
			}(hx_indexof_target_26.Get(hx_indexof_index_29))
			if hxrt.StringEqualStringPtr(hx_indexof_element_30, hx_indexof_value_27) {
				return hx_indexof_index_29
			}
			hx_indexof_index_29 = (hx_indexof_index_29 + 1)
		}
		return -1
	}()), hxrt.StringFromLiteral(",")), hxrt.StdString(func() bool {
		hx_contains_target_33 := names
		var hx_contains_value_34 *string = hxrt.StringFromLiteral("zz")
		hx_contains_length_35 := hx_contains_target_33.Len()
		hx_contains_index_36 := 0
		for hx_contains_index_36 < hx_contains_length_35 {
			hx_contains_element_37 := func(hx_value_38 any) *string {
				if hx_value_38 == nil {
					var hx_zero_39 *string
					return hx_zero_39
				}
				return hx_value_38.(*string)
			}(hx_contains_target_33.Get(hx_contains_index_36))
			if hxrt.StringEqualStringPtr(hx_contains_element_37, hx_contains_value_34) {
				return true
			}
			hx_contains_index_36 = (hx_contains_index_36 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(",")), hxrt.StringJoinAny(lengths.Values(), hxrt.StringFromLiteral("|"))), hxrt.StringFromLiteral(",")), hxrt.StringJoinAny(kept.Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_12)
	hx_array_target_40 := names
	hx_array_index_41 := 2
	hx_array_target_40.Set(hx_array_index_41, hxrt.StringFromLiteral("bb"))
	loopOut := hxrt.NewArray()
	_g_2 := 0
	for _g_2 < names.Len() {
		n := func(hx_value_42 any) *string {
			if hx_value_42 == nil {
				var hx_zero_43 *string
				return hx_zero_43
			}
			return hx_value_42.(*string)
		}(names.Get(_g_2))
		_g_2 = int(int32((_g_2 + 1)))
		s := n
		loopOut.Push(hxrt.StringToUpperCaseStringPtr(s))
	}
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array.loop="), hxrt.StringJoinAny(loopOut.Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_13)
	byKey := New_haxe__ds__StringMap()
	byKey.__hx_this.set(hxrt.StringFromLiteral("k1"), hxrt.StringFromLiteral("v1"))
	byKey.__hx_this.set(hxrt.StringFromLiteral("k2"), hxrt.StringFromLiteral("v2"))
	hxrt.StringFromLiteral("v2")
	got := func(hx_value_45 any) *string {
		if hx_value_45 == nil {
			var hx_zero_46 *string
			return hx_zero_46
		}
		return hx_value_45.(*string)
	}(byKey.__hx_this.get(hxrt.StringFromLiteral("k1")))
	gotText := func(hx_value_47 any) *string {
		if hx_value_47 == nil {
			var hx_zero_48 *string
			return hx_zero_48
		}
		return hx_value_47.(*string)
	}(byKey.__hx_this.get(hxrt.StringFromLiteral("k2")))
	missing := func(hx_value_49 any) *string {
		if hx_value_49 == nil {
			var hx_zero_50 *string
			return hx_zero_50
		}
		return hx_value_49.(*string)
	}(byKey.__hx_this.get(hxrt.StringFromLiteral("none")))
	_g_3 := hxrt.NewArray()
	k := byKey.__hx_this.keys()
	for func() bool {
		var hx_callable_51 *hxrt.Function[func() bool] = hx_recover_callable_cf6928b1(hxrt.AnonymousGet(k, "hasNext"))
		return hxrt.FunctionCall(hx_callable_51)()
	}() {
		k_1 := func() *string {
			var hx_callable_52 *hxrt.Function[func() *string] = hx_recover_callable_b9d17ba0(hxrt.AnonymousGet(k, "next"))
			return hxrt.FunctionCall(hx_callable_52)()
		}()
		_g_3.Push(k_1)
	}
	keys := _g_3
	hxrt.ArraySort(keys, func(hx_comparator_56 *hxrt.Function[func(*string, *string) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_56, func(hx_cmp_left_54 any, hx_cmp_right_55 any) int {
			return hxrt.FunctionCall(hx_comparator_56)(func(hx_value_57 any) *string {
				if hx_value_57 == nil {
					var hx_zero_58 *string
					return hx_zero_58
				}
				return hx_value_57.(*string)
			}(hx_cmp_left_54), func(hx_value_59 any) *string {
				if hx_value_59 == nil {
					var hx_zero_60 *string
					return hx_zero_60
				}
				return hx_value_59.(*string)
			}(hx_cmp_right_55))
		})
	}(hxrt.NamedFunction(compareNames, "Main.compareNames")))
	_g_4 := hxrt.NewArray()
	v_14 := byKey.__hx_this.iterator()
	for func() bool {
		var hx_callable_61 *hxrt.Function[func() bool] = hx_recover_callable_cf6928b1(hxrt.AnonymousGet(v_14, "hasNext"))
		return hxrt.FunctionCall(hx_callable_61)()
	}() {
		v_15 := func() *string {
			var hx_callable_62 *hxrt.Function[func() *string] = hx_recover_callable_b9d17ba0(hxrt.AnonymousGet(v_14, "next"))
			return hxrt.FunctionCall(hx_callable_62)()
		}()
		_g_4.Push(v_15)
	}
	values := _g_4
	hxrt.ArraySort(values, func(hx_comparator_66 *hxrt.Function[func(*string, *string) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_66, func(hx_cmp_left_64 any, hx_cmp_right_65 any) int {
			return hxrt.FunctionCall(hx_comparator_66)(func(hx_value_67 any) *string {
				if hx_value_67 == nil {
					var hx_zero_68 *string
					return hx_zero_68
				}
				return hx_value_67.(*string)
			}(hx_cmp_left_64), func(hx_value_69 any) *string {
				if hx_value_69 == nil {
					var hx_zero_70 *string
					return hx_zero_70
				}
				return hx_value_69.(*string)
			}(hx_cmp_right_65))
		})
	}(hxrt.NamedFunction(compareNames, "Main.compareNames")))
	var v_16 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("map="), got), hxrt.StringFromLiteral(",")), gotText), hxrt.StringFromLiteral(",")), missing), hxrt.StringFromLiteral(",")), hxrt.StdString(func(hx_value_71 any) bool {
		if hx_value_71 == nil {
			var hx_zero_72 bool
			return hx_zero_72
		}
		return hx_value_71.(bool)
	}(byKey.__hx_this.exists(hxrt.StringFromLiteral("k1"))))), hxrt.StringFromLiteral(",")), hxrt.StringJoinAny(keys.Values(), hxrt.StringFromLiteral("|"))), hxrt.StringFromLiteral(",")), hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_16)
	byName := New_haxe__ds__StringMap()
	byName.__hx_this.set(hxrt.StringFromLiteral("n1"), 1)
	nameKey := hxrt.StringFromLiteral("n1")
	var v_17 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("map.key="), func(hx_value_73 any) any {
		if hx_value_73 == nil {
			return nil
		}
		return hx_value_73.(int)
	}(byName.__hx_this.get(nameKey))), hxrt.StringFromLiteral(",")), hxrt.StdString(func(hx_value_74 any) bool {
		if hx_value_74 == nil {
			var hx_zero_75 bool
			return hx_zero_75
		}
		return hx_value_74.(bool)
	}(byName.__hx_this.exists(hxrt.StringFromLiteral("n1"))))), hxrt.StringFromLiteral(",")), hxrt.StdString(func(hx_value_76 any) bool {
		if hx_value_76 == nil {
			var hx_zero_77 bool
			return hx_zero_77
		}
		return hx_value_76.(bool)
	}(byName.__hx_this.exists(hxrt.StringFromLiteral("n2"))))))
	hxrt.Println(v_17)
	c := hxrt.StringFromLiteral("green")
	colorText := c
	colors := hxrt.NewArray(hxrt.StringFromLiteral("red"), hxrt.StringFromLiteral("blue"))
	colors.Push(hxrt.StringFromLiteral("green"))
	colorMap := New_haxe__ds__StringMap()
	colorMap.__hx_this.set(hxrt.StringFromLiteral("sky"), hxrt.StringFromLiteral("blue"))
	var v_18 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("enumAbstract="), colorLetter(c)), hxrt.StringFromLiteral(",")), colorText), hxrt.StringFromLiteral(",")), hxrt.StringJoinAny(colors.Values(), hxrt.StringFromLiteral("|"))), hxrt.StringFromLiteral(",")), hxrt.StdString(hxrt.StringEqualStringPtr(c, hxrt.StringFromLiteral("green")))), hxrt.StringFromLiteral(",")), hxrt.StdString(hxrt.HaxeEqual(colors.Get(1), hxrt.StringFromLiteral("blue")))), hxrt.StringFromLiteral(",")), optColor(nil)), hxrt.StringFromLiteral(",")), optColor(hxrt.StringFromLiteral("blue"))), hxrt.StringFromLiteral(",")), defColor(hxrt.StringFromLiteral("green"))), hxrt.StringFromLiteral(",")), defColor(nil)), hxrt.StringFromLiteral(",")), defColor(hxrt.StringFromLiteral("red"))), hxrt.StringFromLiteral(",")), func(hx_value_79 any) *string {
		if hx_value_79 == nil {
			var hx_zero_80 *string
			return hx_zero_80
		}
		return hx_value_79.(*string)
	}(colorMap.__hx_this.get(hxrt.StringFromLiteral("sky")))))
	hxrt.Println(v_18)
	m := modeOf(hxrt.StringFromLiteral("fast"))
	var hx_switch_81 *string
	switch *hxrt.StdString(m) {
	case *hxrt.StdString(hxrt.StringFromLiteral("fast")):
		hx_switch_81 = hxrt.StringFromLiteral("F")
	case *hxrt.StdString(hxrt.StringFromLiteral("slow")):
		hx_switch_81 = hxrt.StringFromLiteral("S")
	}
	modeLetter := hx_switch_81
	var v_19 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("enumAbstract.from="), m), hxrt.StringFromLiteral(",")), hxrt.StdString(hxrt.StringEqualStringPtr(m, hxrt.StringFromLiteral("fast")))), hxrt.StringFromLiteral(",")), defMode(hxrt.StringFromLiteral("slow"))), hxrt.StringFromLiteral(",")), defMode(nil)), hxrt.StringFromLiteral(",")), modeLetter))
	hxrt.Println(v_19)
	fn := hxrt.NamedFunction(nameToString, "Main.nameToString")
	fromText := hxrt.NamedFunction(stringToName, "Main.stringToName")
	closure := hxrt.NewFunction(func(n_1 *string) *string {
		s_1 := n_1
		return hxrt.StringConcatStringPtr(s_1, hxrt.StringFromLiteral("!"))
	})
	pick := hxrt.NewFunction(func(n_1 *string, extra *string) *string {
		var hx_if_82 *string
		if extra == nil {
			hx_if_82 = n_1
		} else {
			hx_if_82 = extra
		}
		return hx_if_82
	})
	var v_20 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("function="), func() *string {
		var hx_callable_83 *hxrt.Function[func(*string) *string] = fn
		var hx_argument_84 *string = hxrt.StringFromLiteral("f1")
		return hxrt.FunctionCall(hx_callable_83)(hx_argument_84)
	}()), hxrt.StringFromLiteral(",")), func() *string {
		var hx_callable_85 *hxrt.Function[func(*string) *string] = fromText
		var hx_argument_86 *string = hxrt.StringFromLiteral("f2")
		return hxrt.FunctionCall(hx_callable_85)(hx_argument_86)
	}()), hxrt.StringFromLiteral(",")), func() *string {
		var hx_callable_87 *hxrt.Function[func(*string) *string] = closure
		var hx_argument_88 *string = hxrt.StringFromLiteral("f3")
		return hxrt.FunctionCall(hx_callable_87)(hx_argument_88)
	}()), hxrt.StringFromLiteral(",")), func() *string {
		var hx_callable_89 *hxrt.Function[func(*string, *string) *string] = pick
		var hx_argument_90 *string = hxrt.StringFromLiteral("f4")
		var hx_argument_91 *string = nil
		return hxrt.FunctionCall(hx_callable_89)(hx_argument_90, hx_argument_91)
	}()), hxrt.StringFromLiteral(",")), func() *string {
		var hx_callable_92 *hxrt.Function[func(*string, *string) *string] = pick
		var hx_argument_93 *string = hxrt.StringFromLiteral("f4")
		var hx_argument_94 *string = hxrt.StringFromLiteral("f5")
		return hxrt.FunctionCall(hx_callable_92)(hx_argument_93, hx_argument_94)
	}()))
	hxrt.Println(v_20)
	hx_obj_95 := map[string]any{}
	hx_obj_95["name"] = hxrt.StringFromLiteral("anon")
	hx_obj_95["color"] = hxrt.StringFromLiteral("red")
	entry := hxrt.AnonymousFromMap(hx_obj_95)
	hx_obj_96 := map[string]any{}
	hx_obj_96["name"] = hxrt.StringFromLiteral("anon2")
	entry2 := hxrt.AnonymousFromMap(hx_obj_96)
	var v_21 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("anon="), func(hx_value_97 any) *string {
		if hx_value_97 == nil {
			var hx_zero_98 *string
			return hx_zero_98
		}
		return hx_value_97.(*string)
	}(hxrt.AnonymousGet(entry, "name"))), hxrt.StringFromLiteral(",")), func(hx_value_99 any) *string {
		if hx_value_99 == nil {
			var hx_zero_100 *string
			return hx_zero_100
		}
		return hx_value_99.(*string)
	}(hxrt.AnonymousGet(entry, "color"))), hxrt.StringFromLiteral(",")), func(hx_value_101 any) *string {
		if hx_value_101 == nil {
			var hx_zero_102 *string
			return hx_zero_102
		}
		return hx_value_101.(*string)
	}(hxrt.AnonymousGet(entry2, "name"))), hxrt.StringFromLiteral(",")), func(hx_value_103 any) *string {
		if hx_value_103 == nil {
			var hx_zero_104 *string
			return hx_zero_104
		}
		return hx_value_103.(*string)
	}(hxrt.AnonymousGet(entry2, "color"))))
	hxrt.Println(v_21)
	wrapped := Wrap_W(hxrt.StringFromLiteral("wrapped"), hxrt.StringFromLiteral("blue"))
	_g_5 := func(hx_value_105 any) *string {
		if hx_value_105 == nil {
			var hx_zero_106 *string
			return hx_zero_106
		}
		return hx_value_105.(*string)
	}(wrapped.params[0])
	_g1_2 := func(hx_value_107 any) *string {
		if hx_value_107 == nil {
			var hx_zero_108 *string
			return hx_zero_108
		}
		return hx_value_107.(*string)
	}(wrapped.params[1])
	n_1 := _g_5
	wc := _g1_2
	var v_22 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("enum="), nameToString(n_1)), hxrt.StringFromLiteral(",")), wc))
	hxrt.Println(v_22)
	var d any = local
	fromDyn := func(hx_value_109 any) *string {
		if hx_value_109 == nil {
			var hx_zero_110 *string
			return hx_zero_110
		}
		return hx_value_109.(*string)
	}(d)
	fromDynText := fromDyn
	var v_23 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dynamic="), fromDynText), hxrt.StringFromLiteral(",")), hxrt.StdString(d)), hxrt.StringFromLiteral(",")), hxrt.StdString(func(hx_value any) bool {
		switch hx_value.(type) {
		case *string:
			return true
		case string:
			return true
		default:
			return false
		}
	}(d))))
	hxrt.Println(v_23)
	var v_24 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("shapes="), tagText(hxrt.StringFromLiteral("tag"))), hxrt.StringFromLiteral(",")), labelText(func() *string {
		var this1 *string
		this1 = hxrt.StringFromLiteral("label")
		return this1
	}())), hxrt.StringFromLiteral(",")), codeText(func() *string {
		var this1_1 *string
		this1_1 = hxrt.StringFromLiteral("code")
		return this1_1
	}())), hxrt.StringFromLiteral(",")), wordInfo(hxrt.StringFromLiteral("word"))), hxrt.StringFromLiteral(",")), textInfo(hxrt.StringFromLiteral("foo boo"))), hxrt.StringFromLiteral(",")), aliasText(hxrt.StringFromLiteral("alias"))), hxrt.StringFromLiteral(",")), plainText(hxrt.StringFromLiteral("plain"))), hxrt.StringFromLiteral(",")), boxText(hxrt.StringFromLiteral("box"))))
	hxrt.Println(v_24)
	counts := hxrt.NewArray(1, 2, 3)
	total := 0
	_g_6 := 0
	for _g_6 < counts.Len() {
		var k_2 any = counts.Get(_g_6)
		_g_6 = int(int32((_g_6 + 1)))
		kv := hxrt.IntFromNullableAny(k_2)
		total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(kv)))
	}
	_g_7 := New_haxe__ds__StringMap()
	_g_7.__hx_this.set(hxrt.StringFromLiteral("r"), any(0.5))
	ratios := _g_7
	ratio := func(hx_value_111 any) float64 {
		if hx_value_111 == nil {
			var hx_zero_112 float64
			return hx_zero_112
		}
		return hx_value_111.(float64)
	}(ratios.__hx_this.get(hxrt.StringFromLiteral("r")))
	var v_25 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("int="), countOf(5)), hxrt.StringFromLiteral(",")), total), hxrt.StringFromLiteral(",")), optCount(nil)), hxrt.StringFromLiteral(",")), optCount(4)), hxrt.StringFromLiteral(",")), levelOf(any(2))))
	hxrt.Println(v_25)
	var v_26 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("float.bool="), ratioOf(2.5)), hxrt.StringFromLiteral(",")), ratio), hxrt.StringFromLiteral(",")), hxrt.StdString(flagOf(true))), hxrt.StringFromLiteral(",")), hxrt.StdString(flagOf(false))))
	hxrt.Println(v_26)
	var intBox any = 9
	intBoxValue := hxrt.IntFromNullableAny(intBox)
	var v_27 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dynamicAbstract="), looseText(hxrt.StringFromLiteral("text"))), hxrt.StringFromLiteral(",")), looseText(42)), hxrt.StringFromLiteral(",")), intBoxValue))
	hxrt.Println(v_27)
}

func modeOf(s *string) *string {
	return s
}

func nameKind(n *string) *string {
	var hx_switch_113 *string
	switch *hxrt.StdString(n) {
	case *hxrt.StdString(hxrt.StringFromLiteral("alpha")):
		hx_switch_113 = hxrt.StringFromLiteral("first")
	case *hxrt.StdString(hxrt.StringFromLiteral("beta")):
		hx_switch_113 = hxrt.StringFromLiteral("second")
	default:
		hx_switch_113 = hxrt.StringFromLiteral("other")
	}
	return hx_switch_113
}

func nameToString(n *string) *string {
	return n
}

func nullName(n *string) *string {
	var hx_if_114 *string
	if n == nil {
		hx_if_114 = hxrt.StringFromLiteral("null")
	} else {
		hx_if_114 = n
	}
	return hx_if_114
}

func optColor(c *string) *string {
	var hx_if_115 *string
	if c == nil {
		hx_if_115 = hxrt.StringFromLiteral("none")
	} else {
		hx_if_115 = c
	}
	return hx_if_115
}

func optCount(c any) *string {
	var hx_if_116 *string
	if hxrt.AnyEqualsNull(c) {
		hx_if_116 = hxrt.StringFromLiteral("none")
	} else {
		hx_if_116 = hxrt.StdString(c)
	}
	return hx_if_116
}

func optDefName(n *string) *string {
	if n == nil {
		n = hxrt.StringFromLiteral("odflt")
	}
	return n
}

func optName(n *string) *string {
	var hx_if_117 *string
	if n == nil {
		hx_if_117 = hxrt.StringFromLiteral("none")
	} else {
		hx_if_117 = n
	}
	return hx_if_117
}

func plainText(p *string) *string {
	raw := p
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringToUpperCaseStringPtr(p), hxrt.StringFromLiteral("/")), raw)
}

func ratioOf(r any) float64 {
	return func(hx_value_118 any) float64 {
		if hx_value_118 == nil {
			var hx_zero_119 float64
			return hx_zero_119
		}
		return hx_value_118.(float64)
	}(r)
}

func stringToName(s *string) *string {
	return s
}

func tagText(t *string) *string {
	return t
}

func textInfo(t *string) *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringSubstrStringPtr(t, 1, 0, false), hxrt.StringFromLiteral(":")), hxrt.ArrayFromValues(func(hx_sort_src_120 []*string) []any {
		hx_sort_out_122 := make([]any, 0, len(hx_sort_src_120))
		for _, hx_sort_item_121 := range hx_sort_src_120 {
			hx_sort_out_122 = append(hx_sort_out_122, hx_sort_item_121)
		}
		return hx_sort_out_122
	}(hxrt.StringSplitStringPtr(t, hxrt.StringFromLiteral("o")))).Len())
}

func wordInfo(w *string) *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringLengthStringPtr(w), hxrt.StringFromLiteral(":")), hxrt.StringToUpperCaseStringPtr(w)), hxrt.StringFromLiteral(":")), hxrt.StringCharAtStringPtr(w, 0)), hxrt.StringFromLiteral(":")), hxrt.StringIndexOfStringPtr(w, hxrt.StringFromLiteral("r"), 0, false))
}

type Namer interface {
	rename(n *string, suffix *string) *string
}

type I_Person interface {
	rename(n *string, suffix *string) *string
}

type Person struct {
	__hx_this I_Person
	name      *string
	nick      *string
	color     *string
}

func (self *Person) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Person(name *string, nick *string, color *string) *Person {
	return hx_init_Person(hx_alloc_Person(nil), name, nick, color)
}

func hx_alloc_Person(owner any) *Person {
	self := &Person{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Person)
	}
	return self
}

func hx_init_Person(self *Person, name *string, nick *string, color *string) *Person {
	if color == nil {
		color = hxrt.StringFromLiteral("red")
	}
	self.name = name
	self.nick = nick
	self.color = color
	Person_lastColor = color
	return self
}

func (self *Person) rename(n *string, suffix *string) *string {
	if suffix == nil {
		return n
	}
	head := n
	tail := suffix
	return hxrt.StringConcatStringPtr(head, tail)
}

func (self *Person) HxClassName() string {
	return "Person"
}

var Person_lastColor *string = nil

var Person_registry *string = hxrt.StringFromLiteral("reg")

func _Main__Code_Impl___new(raw *string) *string {
	var this1 *string
	this1 = raw
	return this1
}

func _Main__Code_Impl__text(this1 *string) *string {
	return this1
}

func _Main__Label_Impl___new(raw *string) *string {
	var this1 *string
	this1 = raw
	return this1
}

func _Main__Tag_Impl__text(this1 *string) *string {
	return this1
}

func hx_recover_callable_b9d17ba0(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
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
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_cf6928b1(value any) *hxrt.Function[func() bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(hxrt.FunctionCall(source)())
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

func hx_recover_class_haxe__ds__StringMap(value any) *haxe__ds__StringMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__StringMap:
		return source
	default:
		return value.(*haxe__ds__StringMap)
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
