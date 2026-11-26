package main

import "snapshot/hxrt"

type Shape struct {
	tag    int
	params []any
}

func (self *Shape) HxEnumValue() (string, []any) {
	return [3]string{"Circle", "Square", "Empty"}[self.tag], self.params
}

func Shape_Circle(r int) *Shape {
	enumValue := &Shape{tag: 0}
	enumValue.params = []any{r}
	return enumValue
}

func Shape_Square(side int) *Shape {
	enumValue := &Shape{tag: 1}
	enumValue.params = []any{side}
	return enumValue
}

var Shape_Empty *Shape = &Shape{tag: 2}

type I_Box interface {
	size() int
}

type Box struct {
	__hx_this I_Box
	shape     *Shape
}

func (self *Box) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Box(shape *Shape) *Box {
	return hx_init_Box(hx_alloc_Box(nil), shape)
}

func hx_alloc_Box(owner any) *Box {
	self := &Box{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Box)
	}
	return self
}

func hx_init_Box(self *Box, shape *Shape) *Box {
	self.shape = shape
	return self
}

func (self *Box) size() int {
	_g := self.shape
	switch _g.tag {
	case 0:
		_g_1 := _g.params[0].(int)
		r := _g_1
		return int((hxrt.Int32Wrap(r) * hxrt.Int32Wrap(2)))
	case 1:
		_g_2 := _g.params[0].(int)
		side := _g_2
		return side
	case 2:
		return 0
	}
	panic("unreachable")
}

func (self *Box) HxClassName() string {
	return "Box"
}

func area(shape *Shape) int {
	switch shape.tag {
	case 0:
		_g := shape.params[0].(int)
		r := _g
		return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(r) * hxrt.Int32Wrap(r)))) * hxrt.Int32Wrap(3)))
	case 1:
		_g_1 := shape.params[0].(int)
		side := _g_1
		return int((hxrt.Int32Wrap(side) * hxrt.Int32Wrap(side)))
	case 2:
		return 0
	}
	panic("unreachable")
}

func boolSwitch(flag bool) *string {
	if flag {
		return hxrt.StringFromLiteral("yes")
	} else {
		return hxrt.StringFromLiteral("no")
	}
}

func colorName(color int) *string {
	switch color {
	case 0:
		return hxrt.StringFromLiteral("red")
	case 1:
		return hxrt.StringFromLiteral("green")
	}
	panic("unreachable")
}

func guarded(shape *Shape) *string {
	switch shape.tag {
	case 0:
		_g := shape.params[0].(int)
		r := _g
		if r > 10 {
			return hxrt.StringFromLiteral("big circle")
		} else {
			return hxrt.StringFromLiteral("circle")
		}
	case 1:
		_g_1 := shape.params[0].(int)
		side := _g_1
		if side == 0 {
			return hxrt.StringFromLiteral("dot")
		} else {
			return hxrt.StringFromLiteral("other")
		}
	case 2:
		return hxrt.StringFromLiteral("other")
	}
	panic("unreachable")
}

func ifChain(value int) *string {
	if value < 0 {
		return hxrt.StringFromLiteral("negative")
	} else {
		if value == 0 {
			return hxrt.StringFromLiteral("zero")
		} else {
			return hxrt.StringFromLiteral("positive")
		}
	}
}

func insideIf(shape *Shape, flag bool) int {
	if flag {
		switch shape.tag {
		case 0:
			_g := shape.params[0].(int)
			r := _g
			return r
		case 1:
			_g_1 := shape.params[0].(int)
			side := _g_1
			return side
		case 2:
			return 0
		}
	} else {
		return -1
	}
	panic("unreachable")
}

func insideTry(shape *Shape) int {
	hx_try_return_1 := false
	var hx_try_value_2 int
	hxrt.TryCatch(func() {
		switch shape.tag {
		case 0:
			_g := shape.params[0].(int)
			r := _g
			hx_try_value_2 = r
			hx_try_return_1 = true
			return
		case 1:
			_g_1 := shape.params[0].(int)
			side := _g_1
			hx_try_value_2 = side
			hx_try_return_1 = true
			return
		case 2:
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("empty")))
		}
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
			_ = e
			hx_try_value_2 = -1
			hx_try_return_1 = true
			return
		} else {
			hxrt.Throw(hx_caught_3)
		}
	})
	if hx_try_return_1 {
		return hx_try_value_2
	}
	var hx_try_zero_8 int
	return hx_try_zero_8
}

func intDefault(value int) *string {
	if value == 0 {
		return hxrt.StringFromLiteral("zero")
	} else {
		return hxrt.StringFromLiteral("many")
	}
}

func loopUntil(limit int) int {
	i := 0
	for true {
		if int((hxrt.Int32Wrap(i) * hxrt.Int32Wrap(i))) >= limit {
			return i
		}
		i = int(int32((i + 1)))
	}
	panic("unreachable")
}

func main() {
	log := hxrt.NewArray()
	_g := 0
	_g1 := hxrt.NewArray(Shape_Circle(11), Shape_Square(0), Shape_Square(3), Shape_Empty)
	for _g < _g1.Len() {
		shape := func(hx_value_9 any) *Shape {
			if hx_value_9 == nil {
				var hx_zero_10 *Shape
				return hx_zero_10
			}
			return hx_value_9.(*Shape)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(area(shape), hxrt.StringFromLiteral(" ")), guarded(shape)), hxrt.StringFromLiteral(" ")), New_Box(shape).size()))
		hxrt.Println(v)
		var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(nested(shape, false), hxrt.StringFromLiteral(" ")), insideIf(shape, true)), hxrt.StringFromLiteral(" ")), insideTry(shape)), hxrt.StringFromLiteral(" ")), tryInsideArm(shape)))
		hxrt.Println(v_1)
		hxrt.TryCatch(func() {
			var v_2 any = any(throwArms(shape))
			hxrt.Println(v_2)
		}, func(hx_caught_11 any) {
			hx_exception_12 := haxe__Exception_caught(hx_caught_11)
			hx_payload_13 := haxe__Exception_catchValue(hx_caught_11)
			_ = hx_exception_12
			if func(hx_value any) bool {
				switch hx_value.(type) {
				case *string:
					return true
				default:
					return false
				}
			}(hx_payload_13) {
				var e *string = func(hx_value_14 any) *string {
					if hx_value_14 == nil {
						var hx_zero_15 *string
						return hx_zero_15
					}
					return hx_value_14.(*string)
				}(hx_payload_13)
				hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("caught:"), e)))
			} else {
				hxrt.Throw(hx_caught_11)
			}
		})
		voidSwitch(shape, log)
	}
	var v_3 any = any(hxrt.StringJoinAny(log.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(colorName(1), hxrt.StringFromLiteral(" ")), loopUntil(10)), hxrt.StringFromLiteral(" ")), intDefault(0)), hxrt.StringFromLiteral(" ")), stringWildcard(hxrt.StringFromLiteral("b"))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(boolSwitch(false), hxrt.StringFromLiteral(" ")), ifChain(-2)), hxrt.StringFromLiteral(" ")), ifChain(0)), hxrt.StringFromLiteral(" ")), ifChain(3)))
	hxrt.Println(v_5)
	pick := hxrt.NewFunction(func(shape_1 *Shape) *string {
		switch shape_1.tag {
		case 0:
			_g_1 := shape_1.params[0].(int)
			_ = _g_1
			return hxrt.StringFromLiteral("lambda-circle")
		case 1:
			_g_2 := shape_1.params[0].(int)
			_ = _g_2
			return hxrt.StringFromLiteral("lambda-other")
		case 2:
			return hxrt.StringFromLiteral("lambda-other")
		}
		panic("unreachable")
	})
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		var hx_callable_16 *hxrt.Function[func(*Shape) *string] = pick
		var hx_argument_17 *Shape = Shape_Circle(1)
		return hxrt.FunctionCall(hx_callable_16)(hx_argument_17)
	}(), hxrt.StringFromLiteral(" ")), func() *string {
		var hx_callable_18 *hxrt.Function[func(*Shape) *string] = pick
		var hx_argument_19 *Shape = Shape_Empty
		return hxrt.FunctionCall(hx_callable_18)(hx_argument_19)
	}()))
	hxrt.Println(v_6)
}

func nested(shape *Shape, flag bool) int {
	switch shape.tag {
	case 0:
		_g := shape.params[0].(int)
		r := _g
		if flag {
			return r
		} else {
			return int(-int32(r))
		}
	case 1:
		_g_1 := shape.params[0].(int)
		side := _g_1
		if side == 0 {
			return 100
		} else {
			return side
		}
	case 2:
		return 7
	}
	panic("unreachable")
}

func stringWildcard(value *string) int {
	if hxrt.StringEqualStringPtr(value, hxrt.StringFromLiteral("a")) {
		return 1
	} else {
		return 2
	}
}

func throwArms(shape *Shape) int {
	switch shape.tag {
	case 0:
		_g := shape.params[0].(int)
		r := _g
		return r
	case 1:
		_g_1 := shape.params[0].(int)
		_ = _g_1
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("square")))
		var hx_throw_zero_20 int
		return hx_throw_zero_20
	case 2:
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("empty")))
		var hx_throw_zero_21 int
		return hx_throw_zero_21
	}
	panic("unreachable")
}

func tryInsideArm(shape *Shape) int {
	switch shape.tag {
	case 0:
		_g := shape.params[0].(int)
		r := _g
		hx_try_return_22 := false
		var hx_try_value_23 int
		hxrt.TryCatch(func() {
			if r < 0 {
				hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("negative")))
			}
			hx_try_value_23 = r
			hx_try_return_22 = true
			return
		}, func(hx_caught_24 any) {
			hx_exception_25 := haxe__Exception_caught(hx_caught_24)
			hx_payload_26 := haxe__Exception_catchValue(hx_caught_24)
			_ = hx_exception_25
			if func(hx_value any) bool {
				switch hx_value.(type) {
				case *string:
					return true
				default:
					return false
				}
			}(hx_payload_26) {
				var e *string = func(hx_value_27 any) *string {
					if hx_value_27 == nil {
						var hx_zero_28 *string
						return hx_zero_28
					}
					return hx_value_27.(*string)
				}(hx_payload_26)
				_ = e
				hx_try_value_23 = 0
				hx_try_return_22 = true
				return
			} else {
				hxrt.Throw(hx_caught_24)
			}
		})
		if hx_try_return_22 {
			return hx_try_value_23
		}
	case 1:
		_g_1 := shape.params[0].(int)
		side := _g_1
		return side
	case 2:
		return -5
	}
	panic("unreachable")
}

func voidSwitch(shape *Shape, log *hxrt.Array) {
	switch shape.tag {
	case 0:
		_g := shape.params[0].(int)
		r := _g
		log.Push(hxrt.StringConcatAny(hxrt.StringFromLiteral("circle"), r))
		return
	case 1:
		_g_1 := shape.params[0].(int)
		_ = _g_1
		return
	case 2:
		log.Push(hxrt.StringFromLiteral("empty"))
		return
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
