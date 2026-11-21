package main

import "snapshot/hxrt"

type Kind struct {
	tag    int
	params []any
}

func (self *Kind) HxEnumValue() (string, []any) {
	return [2]string{"Plain", "Count"}[self.tag], self.params
}

var Kind_Plain *Kind = &Kind{tag: 0}

func Kind_Count(n int) *Kind {
	enumValue := &Kind{tag: 1}
	enumValue.params = []any{n}
	return enumValue
}

func catchReturn(fail bool) int {
	var hx_try_1 int
	hx_try_return_2 := false
	var hx_try_value_3 int
	hxrt.TryCatch(func() {
		hx_try_1 = mayThrow(fail)
	}, func(hx_caught_4 any) {
		hx_exception_5 := haxe__Exception_caught(hx_caught_4)
		hx_payload_6 := haxe__Exception_catchValue(hx_caught_4)
		_ = hx_exception_5
		var e any = hx_payload_6
		_ = e
		hx_try_value_3 = -1
		hx_try_return_2 = true
		return
	})
	if hx_try_return_2 {
		return hx_try_value_3
	}
	value := hx_try_1
	return int((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(10)))
}

func guardedCase(kind *Kind, early bool) *string {
	var hx_switch_7 *string
	switch kind.tag {
	case 0:
		var hx_if_8 *string
		if early {
			return hxrt.StringFromLiteral("early")
		} else {
			hx_if_8 = hxrt.StringFromLiteral("plain")
		}
		hx_switch_7 = hx_if_8
	case 1:
		_g := kind.params[0].(int)
		n := _g
		hx_switch_7 = hxrt.StringConcatAny(hxrt.StringFromLiteral("count"), n)
	}
	picked := hx_switch_7
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("picked:"), picked)
}

func ifReturn(value int) *string {
	var hx_if_9 int
	if value < 0 {
		return hxrt.StringFromLiteral("negative")
	} else {
		hx_if_9 = int((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(2)))
	}
	half := hx_if_9
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("half:"), half)
}

func loopExits() *string {
	out := hxrt.NewArray()
	_g := 0
	for _g < 6 {
		hx_post_10 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_10
		var hx_if_11 int
		if i == 1 {
			continue
		} else {
			hx_if_11 = int((hxrt.Int32Wrap(i) * hxrt.Int32Wrap(2)))
		}
		doubled := hx_if_11
		var hx_if_12 int
		if i == 4 {
			break
		} else {
			hx_if_12 = int((hxrt.Int32Wrap(doubled) + hxrt.Int32Wrap(1)))
		}
		kept := hx_if_12
		out.Push(kept)
	}
	return hxrt.StringJoinAny(out.Values(), hxrt.StringFromLiteral(","))
}

func main() {
	var v any = any(guardedCase(Kind_Plain, true))
	hxrt.Println(v)
	var v_1 any = any(guardedCase(Kind_Plain, false))
	hxrt.Println(v_1)
	var v_2 any = any(guardedCase(Kind_Count(2), true))
	hxrt.Println(v_2)
	var v_3 any = any(catchReturn(false))
	hxrt.Println(v_3)
	var v_4 any = any(catchReturn(true))
	hxrt.Println(v_4)
	var v_5 any = any(tryBodyReturn(true))
	hxrt.Println(v_5)
	var v_6 any = any(tryBodyReturn(false))
	hxrt.Println(v_6)
	var v_7 any = any(ifReturn(-1))
	hxrt.Println(v_7)
	var v_8 any = any(ifReturn(4))
	hxrt.Println(v_8)
	var v_9 any = any(loopExits())
	hxrt.Println(v_9)
	var v_10 any = any(nestedValueTry(true))
	hxrt.Println(v_10)
	var v_11 any = any(nestedValueTry(false))
	hxrt.Println(v_11)
	var v_12 any = any(nestedTryStatements(0))
	hxrt.Println(v_12)
	var v_13 any = any(nestedTryStatements(1))
	hxrt.Println(v_13)
	var v_14 any = any(nestedTryStatements(2))
	hxrt.Println(v_14)
}

func mayThrow(fail bool) int {
	if fail {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("bad")))
	}
	return 7
}

func nestedTryStatements(mode int) *string {
	hx_try_return_14 := false
	var hx_try_value_15 *string
	hxrt.TryCatch(func() {
		hxrt.TryCatch(func() {
			if mode == 0 {
				hx_try_value_15 = hxrt.StringFromLiteral("inner")
				hx_try_return_14 = true
				return
			}
			if mode == 1 {
				hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("boom")))
			}
		}, func(hx_caught_19 any) {
			hx_exception_20 := haxe__Exception_caught(hx_caught_19)
			hx_payload_21 := haxe__Exception_catchValue(hx_caught_19)
			_ = hx_exception_20
			if func(hx_value any) bool {
				switch hx_value.(type) {
				case *string:
					return true
				default:
					return false
				}
			}(hx_payload_21) {
				var e_1 *string = func(hx_value_22 any) *string {
					if hx_value_22 == nil {
						var hx_zero_23 *string
						return hx_zero_23
					}
					return hx_value_22.(*string)
				}(hx_payload_21)
				_ = e_1
				hx_try_value_15 = hxrt.StringFromLiteral("inner-catch")
				hx_try_return_14 = true
				return
			} else {
				hxrt.Throw(hx_caught_19)
			}
		})
		if hx_try_return_14 {
			return
		}
	}, func(hx_caught_16 any) {
		hx_exception_17 := haxe__Exception_caught(hx_caught_16)
		hx_payload_18 := haxe__Exception_catchValue(hx_caught_16)
		_ = hx_exception_17
		var e any = hx_payload_18
		_ = e
		hx_try_value_15 = hxrt.StringFromLiteral("outer-catch")
		hx_try_return_14 = true
		return
	})
	if hx_try_return_14 {
		return hx_try_value_15
	}
	return hxrt.StringFromLiteral("tail")
}

func nestedValueTry(fail bool) *string {
	var hx_try_24 int
	hx_try_return_25 := false
	var hx_try_value_26 *string
	hxrt.TryCatch(func() {
		var hx_try_27 int
		hxrt.TryCatch(func() {
			hx_try_27 = mayThrow(fail)
		}, func(hx_caught_28 any) {
			hx_exception_29 := haxe__Exception_caught(hx_caught_28)
			hx_payload_30 := haxe__Exception_catchValue(hx_caught_28)
			_ = hx_exception_29
			var e any = hx_payload_30
			_ = e
			hx_try_value_26 = hxrt.StringFromLiteral("inner-catch")
			hx_try_return_25 = true
			return
		})
		if hx_try_return_25 {
			return
		}
		inner := hx_try_27
		hx_try_24 = int((hxrt.Int32Wrap(inner) * hxrt.Int32Wrap(2)))
	}, func(hx_caught_31 any) {
		hx_exception_32 := haxe__Exception_caught(hx_caught_31)
		hx_payload_33 := haxe__Exception_catchValue(hx_caught_31)
		_ = hx_exception_32
		var e_1 any = hx_payload_33
		_ = e_1
		hx_try_24 = -1
	})
	if hx_try_return_25 {
		return hx_try_value_26
	}
	outer := hx_try_24
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("outer:"), outer)
}

func tryBodyReturn(early bool) int {
	var hx_try_34 int
	hx_try_return_35 := false
	var hx_try_value_36 int
	hxrt.TryCatch(func() {
		if early {
			hx_try_value_36 = 5
			hx_try_return_35 = true
			return
		}
		hx_try_34 = mayThrow(false)
	}, func(hx_caught_37 any) {
		hx_exception_38 := haxe__Exception_caught(hx_caught_37)
		hx_payload_39 := haxe__Exception_catchValue(hx_caught_37)
		_ = hx_exception_38
		var e any = hx_payload_39
		_ = e
		hx_try_34 = 0
	})
	if hx_try_return_35 {
		return hx_try_value_36
	}
	value := hx_try_34
	return int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(100)))
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
