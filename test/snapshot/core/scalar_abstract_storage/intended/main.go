package main

import "snapshot/hxrt"

type I_Counter interface {
}

type Counter struct {
	__hx_this I_Counter
	count     int
	meters    float64
	flag      bool
	maybe     any
	byKey     *haxe__ds__IntMap
}

func (self *Counter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Counter(count int) *Counter {
	return hx_init_Counter(hx_alloc_Counter(nil), count)
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

func hx_init_Counter(self *Counter, count int) *Counter {
	self.count = count
	self.meters = 1.5
	self.flag = true
	self.maybe = nil
	self.byKey = New_haxe__ds__IntMap()
	return self
}

func (self *Counter) HxClassName() string {
	return "Counter"
}

var Counter_start int = 7

func main() {
	count := 5
	plusOne := int((hxrt.Int32Wrap(count) + hxrt.Int32Wrap(1)))
	low := int((hxrt.Int32Wrap(1) + hxrt.Int32Wrap(2)))
	score := 4
	meters := 1.5
	flag := true
	big := 2147483647
	wrapped := int((hxrt.Int32Wrap(big) + hxrt.Int32Wrap(1)))
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("ops="), plusOne), hxrt.StringFromLiteral(",")), low), hxrt.StringFromLiteral(",")), int((hxrt.Int32Wrap(score)+hxrt.Int32Wrap(score)))), hxrt.StringFromLiteral(",")), _Main__Score_Impl__doubled(score)), hxrt.StringFromLiteral(",")), (meters+meters)), hxrt.StringFromLiteral(",")), hxrt.StdString(!flag)), hxrt.StringFromLiteral(",")), wrapped))
	hxrt.Println(v)
	counter := New_Counter(next(count))
	counter.count = int((hxrt.Int32Wrap(counter.count) + hxrt.Int32Wrap(2)))
	Counter_start = int((hxrt.Int32Wrap(Counter_start) + hxrt.Int32Wrap(1)))
	var this1 haxe__IMap = counter.byKey
	this1.(*haxe__ds__IntMap).__hx_this.set(count, hxrt.StringFromLiteral("five"))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("fields="), counter.count), hxrt.StringFromLiteral(",")), counter.meters), hxrt.StringFromLiteral(",")), hxrt.StdString(counter.flag)), hxrt.StringFromLiteral(",")), counter.maybe), hxrt.StringFromLiteral(",")), Counter_start), hxrt.StringFromLiteral(",")), func() *string {
		var this1_1 haxe__IMap = counter.byKey
		return func(hx_value_1 any) *string {
			if hx_value_1 == nil {
				var hx_zero_2 *string
				return hx_zero_2
			}
			return hx_value_1.(*string)
		}(this1_1.(*haxe__ds__IntMap).__hx_this.get(5))
	}()))
	hxrt.Println(v_1)
	counts := hxrt.NewArray(3, 1, 2)
	counts.Push(count)
	_g := New_haxe__ds__StringMap()
	_g.__hx_this.set(hxrt.StringFromLiteral("top"), 3)
	levels := _g
	inc := hxrt.NewFunction(func(c int) int {
		return int((hxrt.Int32Wrap(c) + hxrt.Int32Wrap(1)))
	})
	box := 9
	var maybe any = nil
	var dyn any = count
	back := hxrt.IntFromNullableAny(dyn)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("values="), hxrt.StringJoinAny(counts.Values(), hxrt.StringFromLiteral("|"))), hxrt.StringFromLiteral(",")), int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(counts.Get(3)))+hxrt.Int32Wrap(1)))), hxrt.StringFromLiteral(",")), func(hx_value_4 any) any {
		if hx_value_4 == nil {
			return nil
		}
		return hx_value_4.(int)
	}(levels.__hx_this.get(hxrt.StringFromLiteral("top")))), hxrt.StringFromLiteral(",")), func() int {
		var hx_callable_5 *hxrt.Function[func(int) int] = inc
		var hx_argument_6 int = count
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()), hxrt.StringFromLiteral(",")), int((hxrt.Int32Wrap(box)+hxrt.Int32Wrap(1)))), hxrt.StringFromLiteral(",")), optional(nil)), hxrt.StringFromLiteral(",")), optional(4)), hxrt.StringFromLiteral(",")), maybe), hxrt.StringFromLiteral(",")), int((hxrt.Int32Wrap(back) + hxrt.Int32Wrap(1)))))
	hxrt.Println(v_2)
	var quarter any = 2.25
	var seven any = 7
	var castMeters any = hxrt.DynamicToNullableFloat(seven)
	var widened any = castMeters
	var nothing any = nil
	var hx_try_7 *string
	hxrt.TryCatch(func() {
		func(hx_cast_8 any) int {
			if hxrt.AnyEqualsNull(hx_cast_8) {
				var hx_zero_9 int
				return hx_zero_9
			}
			if !func(hx_value any) bool {
				switch hx_value.(type) {
				case int:
					return true
				case int8:
					return true
				case int16:
					return true
				case int32:
					return true
				case int64:
					return true
				case uint:
					return true
				case uint8:
					return true
				case uint16:
					return true
				case uint32:
					return true
				case uint64:
					return true
				case uintptr:
					return true
				default:
					return false
				}
			}(hx_cast_8) {
				hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Class cast error")))
			}
			return hxrt.IntFromNullableAny(hx_cast_8)
		}(nothing)
		hx_try_7 = hxrt.StringFromLiteral("passed")
	}, func(hx_caught_10 any) {
		hx_exception_11 := haxe__Exception_caught(hx_caught_10)
		hx_payload_12 := haxe__Exception_catchValue(hx_caught_10)
		_ = hx_exception_11
		var e any = hx_payload_12
		_ = e
		hx_try_7 = hxrt.StringFromLiteral("threw")
	})
	nullCast := hx_try_7
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("casts="), (hxrt.DynamicToFloatSlot(quarter)*float64(2))), hxrt.StringFromLiteral(",")), (func(hx_value_13 any) float64 {
		if hx_value_13 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_13.(float64)
	}(widened)+0.5)), hxrt.StringFromLiteral(",")), nullCast))
	hxrt.Println(v_3)
}

func next(c int) int {
	return int((hxrt.Int32Wrap(c) + hxrt.Int32Wrap(1)))
}

func optional(c any) *string {
	var hx_if_14 *string
	if c == nil {
		hx_if_14 = hxrt.StringFromLiteral("none")
	} else {
		hx_if_14 = hxrt.StdString(c.(int))
	}
	return hx_if_14
}

func _Main__Score_Impl__doubled(this1 int) int {
	return int((hxrt.Int32Wrap(this1) * hxrt.Int32Wrap(2)))
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
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return source
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
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
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
	case *hxrt.Function[func() int]:
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
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_7 any) bool {
				if hx_value_7 == nil {
					var hx_zero_8 bool
					return hx_zero_8
				}
				return hx_value_7.(bool)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_e273723c(value any) *hxrt.Function[func() int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
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

func hx_recover_class_haxe__ds__IntMap(value any) *haxe__ds__IntMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__IntMap:
		return source
	default:
		return value.(*haxe__ds__IntMap)
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
