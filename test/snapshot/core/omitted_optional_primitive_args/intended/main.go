package main

import "snapshot/hxrt"

type I_BaseScaler interface {
	scale(factor any, steps any, exact any) *string
}

type BaseScaler struct {
	__hx_this I_BaseScaler
}

func (self *BaseScaler) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_BaseScaler() *BaseScaler {
	return hx_init_BaseScaler(hx_alloc_BaseScaler(nil))
}

func hx_alloc_BaseScaler(owner any) *BaseScaler {
	self := &BaseScaler{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_BaseScaler)
	}
	return self
}

func hx_init_BaseScaler(self *BaseScaler) *BaseScaler {
	return self
}

func (self *BaseScaler) scale(factor any, steps any, exact any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("base("), describeFloat(factor)), hxrt.StringFromLiteral(",")), describeInt(steps)), hxrt.StringFromLiteral(",")), describeBool(exact)), hxrt.StringFromLiteral(")"))
}

func (self *BaseScaler) HxClassName() string {
	return "BaseScaler"
}

type I_Box interface {
}

type Box struct {
	__hx_this I_Box
	label     *string
}

func (self *Box) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Box(width any, count any, open any) *Box {
	return hx_init_Box(hx_alloc_Box(nil), width, count, open)
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

func hx_init_Box(self *Box, width any, count any, open any) *Box {
	self.label = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(describeFloat(width), hxrt.StringFromLiteral("|")), describeInt(count)), hxrt.StringFromLiteral("|")), describeBool(open))
	return self
}

func (self *Box) HxClassName() string {
	return "Box"
}

type I_ChildScaler interface {
	scale(factor any, steps any, exact any) *string
}

type ChildScaler struct {
	*BaseScaler
	__hx_this I_ChildScaler
}

func (self *ChildScaler) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_ChildScaler() *ChildScaler {
	return hx_init_ChildScaler(hx_alloc_ChildScaler(nil))
}

func hx_alloc_ChildScaler(owner any) *ChildScaler {
	self := &ChildScaler{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_ChildScaler)
	}
	self.BaseScaler = hx_alloc_BaseScaler(self.__hx_this)
	return self
}

func hx_init_ChildScaler(self *ChildScaler) *ChildScaler {
	hx_init_BaseScaler(self.BaseScaler)
	return self
}

func (self *ChildScaler) scale(factor any, steps any, exact any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("child("), factor), hxrt.StringFromLiteral(",")), steps), hxrt.StringFromLiteral(",")), hxrt.StdString(exact)), hxrt.StringFromLiteral(")>")), self.BaseScaler.scale(factor, steps, exact))
}

func (self *ChildScaler) HxClassName() string {
	return "ChildScaler"
}

type I_Cursor interface {
	describe() *string
}

type Cursor struct {
	__hx_this I_Cursor
	pos       int
	scale     float64
	live      bool
	pair      *Pair
}

func (self *Cursor) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Cursor(pos any, scale any, live any) *Cursor {
	return hx_init_Cursor(hx_alloc_Cursor(nil), pos, scale, live)
}

func hx_alloc_Cursor(owner any) *Cursor {
	self := &Cursor{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Cursor)
	}
	return self
}

func hx_init_Cursor(self *Cursor, pos any, scale any, live any) *Cursor {
	if pos == nil {
		pos = 1
	}
	if scale == nil {
		scale = 0.5
	}
	if live == nil {
		live = true
	}
	self.pos = hxrt.IntFromNullableAny(pos)
	self.scale = func(hx_value_1 any) float64 {
		if hx_value_1 == nil {
			var hx_zero_2 float64
			return hx_zero_2
		}
		return hx_value_1.(float64)
	}(scale)
	self.live = func(hx_value_3 any) bool {
		if hx_value_3 == nil {
			var hx_zero_4 bool
			return hx_zero_4
		}
		return hx_value_3.(bool)
	}(live)
	self.pair = New_Pair(hxrt.IntFromNullableAny(pos), func(hx_value_5 any) float64 {
		if hx_value_5 == nil {
			var hx_zero_6 float64
			return hx_zero_6
		}
		return hx_value_5.(float64)
	}(scale))
	return self
}

func (self *Cursor) describe() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(self.pos, hxrt.StringFromLiteral(",")), self.scale), hxrt.StringFromLiteral(",")), hxrt.StdString(self.live)), hxrt.StringFromLiteral(",")), self.pair.count), hxrt.StringFromLiteral(",")), self.pair.ratio)
}

func (self *Cursor) HxClassName() string {
	return "Cursor"
}

func asInt(value int) int {
	return value
}

func attempt(action *hxrt.Function[func() *string]) *string {
	hx_try_return_7 := false
	var hx_try_value_8 *string
	hxrt.TryCatch(func() {
		hx_try_value_8 = func() *string {
			var hx_callable_12 *hxrt.Function[func() *string] = action
			return hxrt.FunctionCall(hx_callable_12)()
		}()
		hx_try_return_7 = true
		return
	}, func(hx_caught_9 any) {
		hx_exception_10 := haxe__Exception_caught(hx_caught_9)
		hx_payload_11 := haxe__Exception_catchValue(hx_caught_9)
		_ = hx_exception_10
		var hx_tmp any = hx_payload_11
		_ = hx_tmp
		hx_try_value_8 = hxrt.StringFromLiteral("threw")
		hx_try_return_7 = true
		return
	})
	if hx_try_return_7 {
		return hx_try_value_8
	}
	var hx_try_zero_13 *string
	return hx_try_zero_13
}

func captured(x any, y any, z any) *string {
	read := hxrt.NewFunction(func() *string {
		return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(x, hxrt.StringFromLiteral("/")), y), hxrt.StringFromLiteral("/")), hxrt.StdString(z))
	})
	return func() *string {
		var hx_callable_14 *hxrt.Function[func() *string] = read
		return hxrt.FunctionCall(hx_callable_14)()
	}()
}

func compound(x any, y any) *string {
	first := attempt(hxrt.NewFunction(func() *string {
		x = (func(hx_value_15 any) float64 {
			if hx_value_15 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_15.(float64)
		}(x) + float64(1))
		return hxrt.StringFromLiteral("stepped")
	}))
	second := attempt(hxrt.NewFunction(func() *string {
		y = int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(y)) * hxrt.Int32Wrap(2)))
		return hxrt.StringFromLiteral("stepped")
	}))
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(first, hxrt.StringFromLiteral(",")), second), hxrt.StringFromLiteral(",")), x), hxrt.StringFromLiteral(",")), y)
}

func defaults(ratio float64, count int, flag bool) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(ratio, hxrt.StringFromLiteral(",")), count), hxrt.StringFromLiteral(",")), hxrt.StdString(flag)), hxrt.StringFromLiteral(",")), (ratio+float64(1))), hxrt.StringFromLiteral(",")), int((hxrt.Int32Wrap(count)*hxrt.Int32Wrap(2)))), hxrt.StringFromLiteral(",")), func() *string {
		var hx_if_16 *string
		if flag {
			hx_if_16 = hxrt.StringFromLiteral("on")
		} else {
			hx_if_16 = hxrt.StringFromLiteral("off")
		}
		return hx_if_16
	}()), hxrt.StringFromLiteral(",")), hxrt.StdString((ratio < float64(count))))
}

func describeBool(x any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString(x)), hxrt.StringFromLiteral(",")), hxrt.StdString((x == nil))), hxrt.StringFromLiteral(",")), hxrt.StdString((x != nil))), hxrt.StringFromLiteral(",")), hxrt.StdString(x)), hxrt.StringFromLiteral(",")), func() *string {
		var hx_if_19 *string
		if func(hx_value_17 any) bool {
			if hx_value_17 == nil {
				var hx_zero_18 bool
				return hx_zero_18
			}
			return hx_value_17.(bool)
		}(x) {
			hx_if_19 = hxrt.StringFromLiteral("yes")
		} else {
			hx_if_19 = hxrt.StringFromLiteral("no")
		}
		return hx_if_19
	}()), hxrt.StringFromLiteral(",")), hxrt.StdString((x == true))), hxrt.StringFromLiteral(",")), hxrt.StdString((x == false)))
}

func describeFloat(x any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral(""), x), hxrt.StringFromLiteral(",")), hxrt.StdString((x == nil))), hxrt.StringFromLiteral(",")), hxrt.StdString((x != nil))), hxrt.StringFromLiteral(",")), hxrt.StdString(x)), hxrt.StringFromLiteral(",")), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_20 any) float64 {
			if hx_value_20 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_20.(float64)
		}(x) + float64(1)))
	}))), hxrt.StringFromLiteral(",")), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (float64(2) * func(hx_value_21 any) float64 {
			if hx_value_21 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_21.(float64)
		}(x)))
	}))), hxrt.StringFromLiteral(",")), hxrt.StdString((func(hx_value_22 any) float64 {
		if hx_value_22 == nil {
			var hx_zero_23 float64
			return (0 / hx_zero_23)
		}
		return hx_value_22.(float64)
	}(x) < 1))), hxrt.StringFromLiteral(",")), hxrt.StdString((func(hx_value_24 any) float64 {
		if hx_value_24 == nil {
			var hx_zero_25 float64
			return (0 / hx_zero_25)
		}
		return hx_value_24.(float64)
	}(x) >= 0))), hxrt.StringFromLiteral(",")), hxrt.StdString((1 > func(hx_value_26 any) float64 {
		if hx_value_26 == nil {
			var hx_zero_27 float64
			return (0 / hx_zero_27)
		}
		return hx_value_26.(float64)
	}(x))))
}

func describeInt(x any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral(""), x), hxrt.StringFromLiteral(",")), hxrt.StdString((x == nil))), hxrt.StringFromLiteral(",")), hxrt.StdString((x != nil))), hxrt.StringFromLiteral(",")), hxrt.StdString(x)), hxrt.StringFromLiteral(",")), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(x))) + hxrt.Int32Wrap(1))))
	}))), hxrt.StringFromLiteral(",")), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(x))) - hxrt.Int32Wrap(2))))
	}))), hxrt.StringFromLiteral(",")), hxrt.StdString((func(hx_value_28 any) float64 {
		if hx_value_28 == nil {
			var hx_zero_29 float64
			return (0 / hx_zero_29)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_28))
	}(x) < 1.5))), hxrt.StringFromLiteral(",")), hxrt.StdString((func(hx_value_30 any) float64 {
		if hx_value_30 == nil {
			var hx_zero_31 float64
			return (0 / hx_zero_31)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_30))
	}(x) >= 0.5))), hxrt.StringFromLiteral(",")), hxrt.StdString((1.5 > func(hx_value_32 any) float64 {
		if hx_value_32 == nil {
			var hx_zero_33 float64
			return (0 / hx_zero_33)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_32))
	}(x))))
}

func forwardBool(x any) *string {
	return describeBool(x)
}

func forwardFloat(x any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(describeFloat(x), hxrt.StringFromLiteral("|")), nullableFloat(x))
}

func forwardInt(x any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(describeInt(x), hxrt.StringFromLiteral("|")), nullableInt(x))
}

func forwardToDefaults(ratio any, count any, flag any) *string {
	return defaults(func(hx_value_34 any) float64 {
		if hx_value_34 == nil {
			var hx_zero_35 float64
			return hx_zero_35
		}
		return hx_value_34.(float64)
	}(ratio), hxrt.IntFromNullableAny(count), func(hx_value_36 any) bool {
		if hx_value_36 == nil {
			var hx_zero_37 bool
			return hx_zero_37
		}
		return hx_value_36.(bool)
	}(flag))
}

func interpolated(x any, y any, z any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("x="), x), hxrt.StringFromLiteral(" y=")), y), hxrt.StringFromLiteral(" z=")), hxrt.StdString(z))
}

var lastPos int = 0

var lastScale float64 = 0

func main() {
	row(hxrt.StringFromLiteral("static.float"), hxrt.NewArray(describeFloat(nil), describeFloat(nil), describeFloat(0.5)))
	row(hxrt.StringFromLiteral("static.int"), hxrt.NewArray(describeInt(nil), describeInt(nil), describeInt(3)))
	row(hxrt.StringFromLiteral("static.bool"), hxrt.NewArray(describeBool(nil), describeBool(nil), describeBool(true), describeBool(false)))
	row(hxrt.StringFromLiteral("nullable.param"), hxrt.NewArray(nullableFloat(nil), nullableFloat(2.5), nullableInt(nil), nullableInt(-2)))
	row(hxrt.StringFromLiteral("forward"), hxrt.NewArray(forwardFloat(nil), forwardFloat(nil), forwardFloat(1.25), forwardInt(nil), forwardInt(7), forwardBool(nil), forwardBool(false)))
	row(hxrt.StringFromLiteral("defaults"), hxrt.NewArray(defaults(1.5, 4, true), defaults(0.25, 4, true), defaults(0.25, 1, false)))
	row(hxrt.StringFromLiteral("defaults.forward"), hxrt.NewArray(forwardToDefaults(3.5, 5, false), forwardToDefaults(0.5, 1, true)))
	row(hxrt.StringFromLiteral("mixed"), hxrt.NewArray(mixed(hxrt.StringFromLiteral("a"), nil, nil), mixed(hxrt.StringFromLiteral("b"), 1.5, nil), mixed(hxrt.StringFromLiteral("c"), nil, 2), mixed(hxrt.StringFromLiteral("d"), 0.5, 6)))
	meter := New_Meter(10)
	row(hxrt.StringFromLiteral("method"), hxrt.NewArray(meter.read(nil, nil), meter.read(nil, nil), meter.read(2.5, nil), meter.read(1.5, 3)))
	row(hxrt.StringFromLiteral("constructor"), hxrt.NewArray(New_Box(nil, nil, nil).label, New_Box(nil, nil, nil).label, New_Box(4.5, 2, true).label))
	var localFloat *hxrt.Function[func(any, any, any) *string] = hxrt.NewFunction(func(x any, y any, z any) *string {
		return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(x, hxrt.StringFromLiteral(",")), y), hxrt.StringFromLiteral(",")), hxrt.StdString(z)), hxrt.StringFromLiteral(",")), hxrt.StdString((x == nil))), hxrt.StringFromLiteral(",")), attempt(hxrt.NewFunction(func() *string {
			return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_38 any) float64 {
				if hx_value_38 == nil {
					hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
				}
				return hx_value_38.(float64)
			}(x) + float64(hxrt.IntFromNullableAny(y))))
		}))), hxrt.StringFromLiteral(",")), hxrt.StdString((func(hx_value_39 any) float64 {
			if hx_value_39 == nil {
				var hx_zero_40 float64
				return (0 / hx_zero_40)
			}
			return hx_value_39.(float64)
		}(x) < func(hx_value_41 any) float64 {
			if hx_value_41 == nil {
				var hx_zero_42 float64
				return (0 / hx_zero_42)
			}
			return float64(hxrt.IntFromNullableAny(hx_value_41))
		}(y))))
	})
	row(hxrt.StringFromLiteral("local"), hxrt.NewArray(func() *string {
		var hx_callable_43 *hxrt.Function[func(any, any, any) *string] = localFloat
		var hx_argument_44 any = nil
		var hx_argument_45 any = nil
		var hx_argument_46 any = nil
		return hxrt.FunctionCall(hx_callable_43)(hx_argument_44, hx_argument_45, hx_argument_46)
	}(), func() *string {
		var hx_callable_47 *hxrt.Function[func(any, any, any) *string] = localFloat
		var hx_argument_48 any = nil
		var hx_argument_49 any = nil
		var hx_argument_50 any = nil
		return hxrt.FunctionCall(hx_callable_47)(hx_argument_48, hx_argument_49, hx_argument_50)
	}(), func() *string {
		var hx_callable_51 *hxrt.Function[func(any, any, any) *string] = localFloat
		var hx_argument_52 any = 1.5
		var hx_argument_53 any = 2
		var hx_argument_54 any = false
		return hxrt.FunctionCall(hx_callable_51)(hx_argument_52, hx_argument_53, hx_argument_54)
	}()))
	var localDefaults *hxrt.Function[func(any, any, any) *string] = hxrt.NewFunction(func(ratio any, count any, flag any) *string {
		if ratio == nil {
			ratio = 1.5
		}
		if count == nil {
			count = 4
		}
		if flag == nil {
			flag = true
		}
		return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(ratio.(float64), hxrt.StringFromLiteral(",")), count.(int)), hxrt.StringFromLiteral(",")), hxrt.StdString(flag.(bool))), hxrt.StringFromLiteral(",")), (ratio.(float64) + float64(count.(int))))
	})
	var relay *hxrt.Function[func(any, any, any) *string] = hxrt.NewFunction(func(ratio any, count any, flag any) *string {
		return func() *string {
			var hx_callable_59 *hxrt.Function[func(any, any, any) *string] = localDefaults
			var hx_argument_60 any = func(hx_value_55 any) float64 {
				if hx_value_55 == nil {
					var hx_zero_56 float64
					return hx_zero_56
				}
				return hx_value_55.(float64)
			}(ratio)
			var hx_argument_61 any = hxrt.IntFromNullableAny(count)
			var hx_argument_62 any = func(hx_value_57 any) bool {
				if hx_value_57 == nil {
					var hx_zero_58 bool
					return hx_zero_58
				}
				return hx_value_57.(bool)
			}(flag)
			return hxrt.FunctionCall(hx_callable_59)(hx_argument_60, hx_argument_61, hx_argument_62)
		}()
	})
	row(hxrt.StringFromLiteral("local.defaults"), hxrt.NewArray(func() *string {
		var hx_callable_63 *hxrt.Function[func(any, any, any) *string] = localDefaults
		var hx_argument_64 any = 1.5
		var hx_argument_65 any = 4
		var hx_argument_66 any = true
		return hxrt.FunctionCall(hx_callable_63)(hx_argument_64, hx_argument_65, hx_argument_66)
	}(), func() *string {
		var hx_callable_67 *hxrt.Function[func(any, any, any) *string] = localDefaults
		var hx_argument_68 any = 0.5
		var hx_argument_69 any = 4
		var hx_argument_70 any = true
		return hxrt.FunctionCall(hx_callable_67)(hx_argument_68, hx_argument_69, hx_argument_70)
	}(), func() *string {
		var hx_callable_71 *hxrt.Function[func(any, any, any) *string] = relay
		var hx_argument_72 any = 2.5
		var hx_argument_73 any = 1
		var hx_argument_74 any = false
		return hxrt.FunctionCall(hx_callable_71)(hx_argument_72, hx_argument_73, hx_argument_74)
	}()))
	var closure *hxrt.Function[func(any, any, any) *string] = hxrt.NewFunction(func(x any, y any, z any) *string {
		return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(x, hxrt.StringFromLiteral(",")), y), hxrt.StringFromLiteral(",")), hxrt.StdString(z)), hxrt.StringFromLiteral(",")), hxrt.StdString((y == nil))), hxrt.StringFromLiteral(",")), attempt(hxrt.NewFunction(func() *string {
			return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(y))) - hxrt.Int32Wrap(1))))
		}))), hxrt.StringFromLiteral(",")), hxrt.StdString((func(hx_value_75 any) float64 {
			if hx_value_75 == nil {
				var hx_zero_76 float64
				return (0 / hx_zero_76)
			}
			return float64(hxrt.IntFromNullableAny(hx_value_75))
		}(y) > 0.5)))
	})
	row(hxrt.StringFromLiteral("closure"), hxrt.NewArray(func() *string {
		var hx_callable_77 *hxrt.Function[func(any, any, any) *string] = closure
		var hx_argument_78 any = nil
		var hx_argument_79 any = nil
		var hx_argument_80 any = nil
		return hxrt.FunctionCall(hx_callable_77)(hx_argument_78, hx_argument_79, hx_argument_80)
	}(), func() *string {
		var hx_callable_81 *hxrt.Function[func(any, any, any) *string] = closure
		var hx_argument_82 any = nil
		var hx_argument_83 any = nil
		var hx_argument_84 any = nil
		return hxrt.FunctionCall(hx_callable_81)(hx_argument_82, hx_argument_83, hx_argument_84)
	}(), func() *string {
		var hx_callable_85 *hxrt.Function[func(any, any, any) *string] = closure
		var hx_argument_86 any = 0.5
		var hx_argument_87 any = 9
		var hx_argument_88 any = true
		return hxrt.FunctionCall(hx_callable_85)(hx_argument_86, hx_argument_87, hx_argument_88)
	}()))
	var scaler Scaler = New_BaseScaler()
	var child Scaler = New_ChildScaler()
	direct := New_ChildScaler()
	row(hxrt.StringFromLiteral("interface"), hxrt.NewArray(scaler.scale(nil, nil, nil), scaler.scale(nil, nil, nil), scaler.scale(2.5, 3, true)))
	row(hxrt.StringFromLiteral("override"), hxrt.NewArray(child.scale(nil, nil, nil), child.scale(1.5, nil, nil), direct.scale(nil, nil, nil), direct.scale(0.5, 1, false)))
	row(hxrt.StringFromLiteral("captured"), hxrt.NewArray(captured(nil, nil, nil), captured(nil, nil, nil), captured(1.5, 2, true)))
	row(hxrt.StringFromLiteral("updated"), hxrt.NewArray(updated(nil, nil), updated(nil, 4), updated(0.5, 1)))
	row(hxrt.StringFromLiteral("compound"), hxrt.NewArray(compound(nil, nil), compound(nil, nil), compound(1.5, 2)))
	row(hxrt.StringFromLiteral("interpolated"), hxrt.NewArray(interpolated(nil, nil, nil), interpolated(1.5, 2, true)))
	row(hxrt.StringFromLiteral("narrowed"), hxrt.NewArray(narrowed(nil, nil), narrowed(1.5, 2)))
	row(hxrt.StringFromLiteral("typed.slots"), hxrt.NewArray(slots(nil, nil, nil), slots(3, 0.25, true)))
	row(hxrt.StringFromLiteral("constructor.fields"), hxrt.NewArray(New_Cursor(nil, nil, nil).describe(), New_Cursor(4, 2.5, false).describe()))
}

func mixed(label *string, x any, y any) *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(":")), x), hxrt.StringFromLiteral(":")), y)
}

func narrowed(x any, y any) *string {
	var hx_if_89 *string
	if x != nil {
		hx_if_89 = hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (x.(float64) + float64(1)))
	} else {
		hx_if_89 = hxrt.StringFromLiteral("none")
	}
	left := hx_if_89
	var hx_if_90 *string
	if y == nil {
		hx_if_90 = hxrt.StringFromLiteral("none")
	} else {
		hx_if_90 = hxrt.StringConcatAny(hxrt.StringFromLiteral(""), int((hxrt.Int32Wrap(y.(int)) + hxrt.Int32Wrap(1))))
	}
	right := hx_if_90
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(left, hxrt.StringFromLiteral(",")), right)
}

func nullableFloat(x any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral(""), x), hxrt.StringFromLiteral(",")), hxrt.StdString((x == nil))), hxrt.StringFromLiteral(",")), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (func(hx_value_91 any) float64 {
			if hx_value_91 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_91.(float64)
		}(x) - float64(1)))
	}))), hxrt.StringFromLiteral(",")), hxrt.StdString((func(hx_value_92 any) float64 {
		if hx_value_92 == nil {
			var hx_zero_93 float64
			return (0 / hx_zero_93)
		}
		return hx_value_92.(float64)
	}(x) > 0)))
}

func nullableInt(x any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral(""), x), hxrt.StringFromLiteral(",")), hxrt.StdString((x == nil))), hxrt.StringFromLiteral(",")), attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(x))) * hxrt.Int32Wrap(3))))
	}))), hxrt.StringFromLiteral(",")), hxrt.StdString((func(hx_value_94 any) float64 {
		if hx_value_94 == nil {
			var hx_zero_95 float64
			return (0 / hx_zero_95)
		}
		return float64(hxrt.IntFromNullableAny(hx_value_94))
	}(x) <= 0.5)))
}

func row(label *string, values *hxrt.Array) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(":"))))
	hxrt.Println(v)
}

func slots(p any, r any, f any) *string {
	if p == nil {
		p = 2
	}
	if r == nil {
		r = 1.5
	}
	if f == nil {
		f = false
	}
	n := 0
	n = hxrt.IntFromNullableAny(p)
	ratio := 0.0
	ratio = func(hx_value_96 any) float64 {
		if hx_value_96 == nil {
			var hx_zero_97 float64
			return hx_zero_97
		}
		return hx_value_96.(float64)
	}(r)
	flag := true
	flag = func(hx_value_98 any) bool {
		if hx_value_98 == nil {
			var hx_zero_99 bool
			return hx_zero_99
		}
		return hx_value_98.(bool)
	}(f)
	lastPos = hxrt.IntFromNullableAny(p)
	lastScale = func(hx_value_100 any) float64 {
		if hx_value_100 == nil {
			var hx_zero_101 float64
			return hx_zero_101
		}
		return hx_value_100.(float64)
	}(r)
	var record_w float64
	var record_v int
	record_v = 0
	record_w = 0
	record_v = hxrt.IntFromNullableAny(p)
	record_w = func(hx_value_102 any) float64 {
		if hx_value_102 == nil {
			var hx_zero_103 float64
			return hx_zero_103
		}
		return hx_value_102.(float64)
	}(r)
	cells := hxrt.NewArray(0, 0)
	hx_array_target_104 := cells
	hx_array_index_105 := 0
	hx_array_target_104.Set(hx_array_index_105, p)
	hx_array_target_106 := cells
	hx_array_index_107 := 1
	hx_array_target_106.Set(hx_array_index_107, asInt(hxrt.IntFromNullableAny(p)))
	var hx_if_110 int
	if func(hx_value_108 any) bool {
		if hx_value_108 == nil {
			var hx_zero_109 bool
			return hx_zero_109
		}
		return hx_value_108.(bool)
	}(f) {
		hx_if_110 = 0
	} else {
		hx_if_110 = hxrt.IntFromNullableAny(p)
	}
	picked := hx_if_110
	return hxrt.StringJoinAny(hxrt.NewArray(hxrt.StringConcatAny(hxrt.StringFromLiteral(""), n), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), ratio), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString(flag)), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), lastPos), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), lastScale), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), record_v), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), record_w), hxrt.StringJoinAny(cells.Values(), hxrt.StringFromLiteral("-")), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), picked), hxrt.StringCharAtStringPtr(hxrt.StringFromLiteral("abcdef"), p.(int)), hxrt.StringSubstrStringPtr(hxrt.StringFromLiteral("abcdef"), p.(int), p.(int), true), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), hxrt.StringCharCodeAtAnyStringPtr(hxrt.StringFromLiteral("abcdef"), p.(int))), hxrt.StringFromCharCode(int((hxrt.Int32Wrap(64)+hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(p))))))).Values(), hxrt.StringFromLiteral(","))
}

func updated(x any, y any) *string {
	before := hxrt.StringConcatAny(hxrt.StringConcatAny(x, hxrt.StringFromLiteral("/")), y)
	if x == nil {
		x = 2.5
	}
	if y == nil {
		y = 3
	}
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(before, hxrt.StringFromLiteral(">")), x), hxrt.StringFromLiteral("/")), y), hxrt.StringFromLiteral(">")), (func(hx_value_111 any) float64 {
		if hx_value_111 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_111.(float64)
	}(x) + float64(hxrt.IntFromNullableAny(y))))
}

type I_Meter interface {
	read(extra any, times any) *string
}

type Meter struct {
	__hx_this I_Meter
	offset    float64
}

func (self *Meter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Meter(offset float64) *Meter {
	return hx_init_Meter(hx_alloc_Meter(nil), offset)
}

func hx_alloc_Meter(owner any) *Meter {
	self := &Meter{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Meter)
	}
	return self
}

func hx_init_Meter(self *Meter, offset float64) *Meter {
	self.offset = offset
	return self
}

func (self *Meter) read(extra any, times any) *string {
	_gthis := self
	shown := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral(""), extra), hxrt.StringFromLiteral("/")), times)
	total := attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (_gthis.offset + func(hx_value_112 any) float64 {
			if hx_value_112 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
			}
			return hx_value_112.(float64)
		}(extra)))
	}))
	scaled := attempt(hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringFromLiteral(""), (_gthis.offset * float64(hxrt.IntFromNullableAny(times))))
	}))
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(shown, hxrt.StringFromLiteral(":")), total), hxrt.StringFromLiteral(":")), scaled)
}

func (self *Meter) HxClassName() string {
	return "Meter"
}

type I_Pair interface {
}

type Pair struct {
	__hx_this I_Pair
	count     int
	ratio     float64
}

func (self *Pair) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Pair(count int, ratio float64) *Pair {
	return hx_init_Pair(hx_alloc_Pair(nil), count, ratio)
}

func hx_alloc_Pair(owner any) *Pair {
	self := &Pair{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Pair)
	}
	return self
}

func hx_init_Pair(self *Pair, count int, ratio float64) *Pair {
	self.count = count
	self.ratio = ratio
	return self
}

func (self *Pair) HxClassName() string {
	return "Pair"
}

type Scaler interface {
	scale(factor any, steps any, exact any) *string
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
