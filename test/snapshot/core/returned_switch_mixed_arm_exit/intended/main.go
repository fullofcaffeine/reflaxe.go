package main

import "snapshot/hxrt"

type Shape struct {
	tag    int
	params []any
}

func (self *Shape) HxEnumValue() (string, []any) {
	return [3]string{"Dot", "Line", "Box"}[self.tag], self.params
}

var Shape_Dot *Shape = &Shape{tag: 0}

func Shape_Line(length int) *Shape {
	enumValue := &Shape{tag: 1}
	enumValue.params = []any{length}
	return enumValue
}

func Shape_Box(width int, height int) *Shape {
	enumValue := &Shape{tag: 2}
	enumValue.params = []any{width, height}
	return enumValue
}

func elseIfChain(kind int, a int) int {
	var hx_if_4 int
	if kind == 1 {
		var hx_if_3 int
		if a == 0 {
			return 10
		} else {
			var hx_if_2 int
			if a == 1 {
				hx_if_2 = 20
			} else {
				var hx_if_1 int
				if a == 2 {
					return 30
				} else {
					hx_if_1 = 40
				}
				hx_if_2 = hx_if_1
			}
			hx_if_3 = hx_if_2
		}
		hx_if_4 = hx_if_3
	} else {
		hx_if_4 = 0
	}
	return hx_if_4
}

func enumSubject(shape *Shape, flag bool) int {
	var hx_switch_5 int
	switch shape.tag {
	case 0:
		hx_switch_5 = 0
	case 1:
		_g := shape.params[0].(int)
		length := _g
		var hx_if_6 int
		if length > 0 {
			return int((hxrt.Int32Wrap(length) + hxrt.Int32Wrap(10)))
		} else {
			hx_if_6 = int((hxrt.Int32Wrap(length) + hxrt.Int32Wrap(20)))
		}
		hx_switch_5 = hx_if_6
	case 2:
		_g_1 := shape.params[0].(int)
		_g1 := shape.params[1].(int)
		w := _g_1
		h := _g1
		var hx_if_8 int
		if flag {
			var hx_if_7 int
			if w == h {
				return int((hxrt.Int32Wrap(w) * hxrt.Int32Wrap(h)))
			} else {
				hx_if_7 = int((hxrt.Int32Wrap(w) + hxrt.Int32Wrap(h)))
			}
			hx_if_8 = hx_if_7
		} else {
			w_1 := _g_1
			hx_if_8 = w_1
		}
		hx_switch_5 = hx_if_8
	}
	return hx_switch_5
}

func main() {
	var v any = any(mixedArm(1, true))
	hxrt.Println(v)
	var v_1 any = any(mixedArm(1, false))
	hxrt.Println(v_1)
	var v_2 any = any(mixedArm(2, true))
	hxrt.Println(v_2)
	var v_3 any = any(elseIfChain(1, 0))
	hxrt.Println(v_3)
	var v_4 any = any(elseIfChain(1, 1))
	hxrt.Println(v_4)
	var v_5 any = any(elseIfChain(1, 2))
	hxrt.Println(v_5)
	var v_6 any = any(elseIfChain(1, 3))
	hxrt.Println(v_6)
	var v_7 any = any(switchInArm(1, 0))
	hxrt.Println(v_7)
	var v_8 any = any(switchInArm(1, 1))
	hxrt.Println(v_8)
	var v_9 any = any(switchInArm(1, 2))
	hxrt.Println(v_9)
	var v_10 any = any(switchInArm(3, 0))
	hxrt.Println(v_10)
	var v_11 any = any(tryInArm(1, false))
	hxrt.Println(v_11)
	var v_12 any = any(tryInArm(1, true))
	hxrt.Println(v_12)
	var v_13 any = any(tryInArm(2, true))
	hxrt.Println(v_13)
	var v_14 any = any(tryInArm(2, false))
	hxrt.Println(v_14)
	shape := Shape_Line(5)
	var v_15 any = any(enumSubject(shape, true))
	hxrt.Println(v_15)
	shape_1 := Shape_Line(-1)
	var v_16 any = any(enumSubject(shape_1, true))
	hxrt.Println(v_16)
	shape_2 := Shape_Box(3, 3)
	var v_17 any = any(enumSubject(shape_2, true))
	hxrt.Println(v_17)
	shape_3 := Shape_Box(3, 4)
	var v_18 any = any(enumSubject(shape_3, true))
	hxrt.Println(v_18)
	var v_19 any = any(enumSubject(Shape_Box(3, 4), false))
	hxrt.Println(v_19)
	var v_20 any = any(New_Picker(1).pick(true))
	hxrt.Println(v_20)
	var v_21 any = any(New_Picker(1).pick(false))
	hxrt.Println(v_21)
	offset := 5
	local := hxrt.NewFunction(func(name *string, flag bool) *string {
		var hx_if_10 *string
		if hxrt.StringEqualStringPtr(name, hxrt.StringFromLiteral("x")) {
			var hx_if_9 *string
			if flag {
				return hxrt.StringConcatAny(hxrt.StringFromLiteral("X"), offset)
			} else {
				hx_if_9 = hxrt.StringConcatAny(hxrt.StringFromLiteral("x"), offset)
			}
			hx_if_10 = hx_if_9
		} else {
			hx_if_10 = name
		}
		return hx_if_10
	})
	var v_22 any = any(func() *string {
		var hx_callable_11 *hxrt.Function[func(*string, bool) *string] = local
		var hx_argument_12 *string = hxrt.StringFromLiteral("x")
		var hx_argument_13 bool = true
		return hxrt.FunctionCall(hx_callable_11)(hx_argument_12, hx_argument_13)
	}())
	hxrt.Println(v_22)
	var v_23 any = any(func() *string {
		var hx_callable_14 *hxrt.Function[func(*string, bool) *string] = local
		var hx_argument_15 *string = hxrt.StringFromLiteral("x")
		var hx_argument_16 bool = false
		return hxrt.FunctionCall(hx_callable_14)(hx_argument_15, hx_argument_16)
	}())
	hxrt.Println(v_23)
	var v_24 any = any(func() *string {
		var hx_callable_17 *hxrt.Function[func(*string, bool) *string] = local
		var hx_argument_18 *string = hxrt.StringFromLiteral("y")
		var hx_argument_19 bool = true
		return hxrt.FunctionCall(hx_callable_17)(hx_argument_18, hx_argument_19)
	}())
	hxrt.Println(v_24)
}

func mayThrow(fail bool) int {
	if fail {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("bad")))
	}
	return 7
}

func mixedArm(kind int, flag bool) int {
	value := 1
	var hx_if_21 int
	if kind == 1 {
		var hx_if_20 int
		if flag {
			return int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(10)))
		} else {
			hx_if_20 = int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(20)))
		}
		hx_if_21 = hx_if_20
	} else {
		hx_if_21 = -1
	}
	return hx_if_21
}

func switchInArm(kind int, inner int) int {
	var hx_if_24 int
	if kind == 1 {
		var hx_switch_22 int
		switch inner {
		case 0:
			return 100
		case 1:
			var hx_if_23 int
			if kind == 1 {
				return 101
			} else {
				hx_if_23 = 0
			}
			hx_switch_22 = hx_if_23
		default:
			hx_switch_22 = int((hxrt.Int32Wrap(inner) * hxrt.Int32Wrap(2)))
		}
		hx_if_24 = hx_switch_22
	} else {
		hx_if_24 = int(-int32(kind))
	}
	return hx_if_24
}

func tryInArm(kind int, fail bool) int {
	var hx_switch_25 int
	switch kind {
	case 1:
		var hx_try_26 int
		hx_try_return_27 := false
		var hx_try_value_28 int
		hxrt.TryCatch(func() {
			hx_try_26 = mayThrow(fail)
		}, func(hx_caught_29 any) {
			hx_exception_30 := haxe__Exception_caught(hx_caught_29)
			hx_payload_31 := haxe__Exception_catchValue(hx_caught_29)
			_ = hx_exception_30
			var e any = hx_payload_31
			_ = e
			hx_try_value_28 = -7
			hx_try_return_27 = true
			return
		})
		if hx_try_return_27 {
			return hx_try_value_28
		}
		hx_switch_25 = hx_try_26
	case 2:
		var hx_try_32 int
		hx_try_return_33 := false
		var hx_try_value_34 int
		hxrt.TryCatch(func() {
			if fail {
				hx_try_value_34 = 200
				hx_try_return_33 = true
				return
			}
			hx_try_32 = int((hxrt.Int32Wrap(mayThrow(false)) + hxrt.Int32Wrap(1)))
		}, func(hx_caught_35 any) {
			hx_exception_36 := haxe__Exception_caught(hx_caught_35)
			hx_payload_37 := haxe__Exception_catchValue(hx_caught_35)
			_ = hx_exception_36
			var e_1 any = hx_payload_37
			_ = e_1
			hx_try_32 = -2
		})
		if hx_try_return_33 {
			return hx_try_value_34
		}
		hx_switch_25 = hx_try_32
	default:
		hx_switch_25 = 0
	}
	return hx_switch_25
}

type I_Picker interface {
	pick(flag bool) int
}

type Picker struct {
	__hx_this I_Picker
	base      int
}

func (self *Picker) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Picker(base int) *Picker {
	return hx_init_Picker(hx_alloc_Picker(nil), base)
}

func hx_alloc_Picker(owner any) *Picker {
	self := &Picker{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Picker)
	}
	return self
}

func hx_init_Picker(self *Picker, base int) *Picker {
	self.base = base
	return self
}

func (self *Picker) pick(flag bool) int {
	_g := self.base
	var hx_if_39 int
	if _g == 0 {
		hx_if_39 = -1
	} else {
		var hx_if_38 int
		if flag {
			return int((hxrt.Int32Wrap(self.base) + hxrt.Int32Wrap(10)))
		} else {
			hx_if_38 = int((hxrt.Int32Wrap(self.base) + hxrt.Int32Wrap(20)))
		}
		hx_if_39 = hx_if_38
	}
	return hx_if_39
}

func (self *Picker) HxClassName() string {
	return "Picker"
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
