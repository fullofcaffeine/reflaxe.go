package main

import "snapshot/hxrt"

type I_Counter interface {
	add(helper_local int) int
}

type Counter struct {
	__hx_this I_Counter
	base      int
	scale     *hxrt.Function[func(int) int]
}

func (self *Counter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Counter(count_local int) *Counter {
	return hx_init_Counter(hx_alloc_Counter(nil), count_local)
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

func hx_init_Counter(self *Counter, count_local int) *Counter {
	self.scale = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func(count_local_1 int) int {
		return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(count_local_1) * hxrt.Int32Wrap(10)))) + hxrt.Int32Wrap(count)))
	}), func() any {
		return self.__hx_this
	}, "scale")
	self.base = int((hxrt.Int32Wrap(count_local) + hxrt.Int32Wrap(count)))
	return self
}

func (self *Counter) add(helper_local int) int {
	return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(self.base) + hxrt.Int32Wrap(helper_local)))) + hxrt.Int32Wrap(helper(helper_local, 1))))
}

func (self *Counter) HxClassName() string {
	return "Counter"
}

type I_Gauge interface {
}

type Gauge struct {
	__hx_this I_Gauge
	start     int
	read      *hxrt.Function[func(int) int]
}

func (self *Gauge) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Gauge(count_local int) *Gauge {
	return hx_init_Gauge(hx_alloc_Gauge(nil), count_local)
}

func hx_alloc_Gauge(owner any) *Gauge {
	self := &Gauge{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Gauge)
	}
	return self
}

func hx_init_Gauge(self *Gauge, count_local int) *Gauge {
	self.read = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func(step int) int {
		return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(self.start) + hxrt.Int32Wrap(step)))) + hxrt.Int32Wrap(count)))
	}), func() any {
		return self.__hx_this
	}, "read")
	self.start = count_local
	return self
}

func (self *Gauge) HxClassName() string {
	return "Gauge"
}

type I_Holder interface {
}

type Holder struct {
	__hx_this I_Holder
	seen      int
	total     int
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder(count_local int) *Holder {
	return hx_init_Holder(hx_alloc_Holder(nil), count_local)
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

func hx_init_Holder(self *Holder, count_local int) *Holder {
	self.seen = int((hxrt.Int32Wrap(count) * hxrt.Int32Wrap(10)))
	self.total = int((hxrt.Int32Wrap(count_local) + hxrt.Int32Wrap(self.seen)))
	return self
}

func (self *Holder) HxClassName() string {
	return "Holder"
}

func closures() *string {
	count_local := 40
	helper_local := 50
	read := hxrt.NewFunction(func() *string {
		return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(count_local, hxrt.StringFromLiteral(":")), count), hxrt.StringFromLiteral(":")), helper_local), hxrt.StringFromLiteral(":")), helper(count_local, helper_local))
	})
	write := hxrt.NewFunction(func(step int) int {
		count = int((hxrt.Int32Wrap(count) + hxrt.Int32Wrap(step)))
		count_local = int((hxrt.Int32Wrap(count_local) + hxrt.Int32Wrap(step)))
		return func() int {
			helper_local = helper(helper_local, step)
			return helper_local
		}()
	})
	func() int {
		var hx_callable_1 *hxrt.Function[func(int) int] = write
		var hx_argument_2 int = 3
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(func() *string {
		var hx_callable_3 *hxrt.Function[func() *string] = read
		return hxrt.FunctionCall(hx_callable_3)()
	}(), hxrt.StringFromLiteral(" ")), count_local), hxrt.StringFromLiteral(":")), count), hxrt.StringFromLiteral(":")), helper_local)
}

var count int = 1

func countCatch() int {
	hx_try_return_4 := false
	var hx_try_value_5 int
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(7))
	}, func(hx_caught_6 any) {
		hx_exception_7 := haxe__Exception_caught(hx_caught_6)
		hx_payload_8 := haxe__Exception_catchValue(hx_caught_6)
		_ = hx_exception_7
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case int:
				return true
			default:
				return false
			}
		}(hx_payload_8) {
			var count_local int = func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(hx_payload_8)
			count = int((hxrt.Int32Wrap(count) + hxrt.Int32Wrap(count_local)))
			hx_try_value_5 = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(count_local) * hxrt.Int32Wrap(1000)))) + hxrt.Int32Wrap(count)))
			hx_try_return_4 = true
			return
		} else {
			hxrt.Throw(hx_caught_6)
		}
	})
	if hx_try_return_4 {
		return hx_try_value_5
	}
	var hx_try_zero_11 int
	return hx_try_zero_11
}

func countLoop(limit int) int {
	total := 0
	_g := 0
	_g1 := limit
	for _g < _g1 {
		hx_post_12 := _g
		_g = int(int32((_g + 1)))
		count_local := hx_post_12
		total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(count_local) * hxrt.Int32Wrap(10)))) + hxrt.Int32Wrap(count))))))
	}
	return total
}

func helper(a int, b int) int {
	return int((hxrt.Int32Wrap(a) + hxrt.Int32Wrap(b)))
}

func helperCatch() *string {
	hx_try_return_13 := false
	var hx_try_value_14 *string
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("boom")))
	}, func(hx_caught_15 any) {
		hx_exception_16 := haxe__Exception_caught(hx_caught_15)
		hx_payload_17 := haxe__Exception_catchValue(hx_caught_15)
		_ = hx_exception_16
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_17) {
			var helper_local *string = func(hx_value_18 any) *string {
				if hx_value_18 == nil {
					var hx_zero_19 *string
					return hx_zero_19
				}
				return hx_value_18.(*string)
			}(hx_payload_17)
			hx_try_value_14 = hxrt.StringConcatAny(helper_local, helper(2, 3))
			hx_try_return_13 = true
			return
		} else {
			hxrt.Throw(hx_caught_15)
		}
	})
	if hx_try_return_13 {
		return hx_try_value_14
	}
	var hx_try_zero_20 *string
	return hx_try_zero_20
}

func helperLoop(values *hxrt.Array) int {
	total := 0
	_g := 0
	for _g < values.Len() {
		helper_local := hxrt.IntFromNullableAny(values.Get(_g))
		_g = int(int32((_g + 1)))
		total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(helper_local) * hxrt.Int32Wrap(100)))) + hxrt.Int32Wrap(helper(helper_local, helper_local)))))))
	}
	return total
}

func localFunctions() *string {
	count_local := hxrt.NewFunction(func() int {
		return int((hxrt.Int32Wrap(count) * hxrt.Int32Wrap(2)))
	})
	helper_local := hxrt.NewFunction(func(a int) int {
		return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(a) * hxrt.Int32Wrap(10)))) + hxrt.Int32Wrap(helper(a, 1))))
	})
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(func() int {
		var hx_callable_21 *hxrt.Function[func() int] = count_local
		return hxrt.FunctionCall(hx_callable_21)()
	}(), hxrt.StringFromLiteral(":")), count), hxrt.StringFromLiteral(" ")), func() int {
		var hx_callable_22 *hxrt.Function[func(int) int] = helper_local
		var hx_argument_23 int = 2
		return hxrt.FunctionCall(hx_callable_22)(hx_argument_23)
	}()), hxrt.StringFromLiteral(":")), helper(2, 2))
}

func locals() *string {
	count_local := 3
	helper_local := 5
	before := hxrt.StringConcatAny(hxrt.StringConcatAny(count_local, hxrt.StringFromLiteral(":")), count)
	count = int((hxrt.Int32Wrap(count) + hxrt.Int32Wrap(count_local)))
	count = int(int32((count + 1)))
	count_local = int(int32((count_local + 1)))
	add := hxrt.NamedFunction(helper, "Main.helper")
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(before, hxrt.StringFromLiteral(" ")), count_local), hxrt.StringFromLiteral(":")), count), hxrt.StringFromLiteral(" ")), int((hxrt.Int32Wrap(helper_local)+hxrt.Int32Wrap(helper(1, 2))))), hxrt.StringFromLiteral(" ")), func() int {
		var hx_callable_24 *hxrt.Function[func(int, int) int] = add
		var hx_argument_25 int = helper_local
		var hx_argument_26 int = count_local
		return hxrt.FunctionCall(hx_callable_24)(hx_argument_25, hx_argument_26)
	}())
}

func main() {
	var v any = any(locals())
	hxrt.Println(v)
	var v_1 any = any(params(10, 20))
	hxrt.Println(v_1)
	var v_2 any = any(countLoop(3))
	hxrt.Println(v_2)
	var v_3 any = any(helperLoop(hxrt.NewArray(1, 2)))
	hxrt.Println(v_3)
	var v_4 any = any(countCatch())
	hxrt.Println(v_4)
	var v_5 any = any(helperCatch())
	hxrt.Println(v_5)
	var v_6 any = any(localFunctions())
	hxrt.Println(v_6)
	var v_7 any = any(closures())
	hxrt.Println(v_7)
	var v_8 any = any(Other_run(4))
	hxrt.Println(v_8)
	counter := New_Counter(6)
	var v_9 any = any(hxrt.StringConcatAny(hxrt.StringConcatAny(counter.add(7), hxrt.StringFromLiteral(" ")), func() int {
		var hx_callable_27 *hxrt.Function[func(int) int] = counter.scale
		var hx_argument_28 int = 2
		return hxrt.FunctionCall(hx_callable_27)(hx_argument_28)
	}()))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatAny(hxrt.StringConcatAny(func() int {
		var hx_callable_29 *hxrt.Function[func(int) int] = offset
		var hx_argument_30 int = 4
		return hxrt.FunctionCall(hx_callable_29)(hx_argument_30)
	}(), hxrt.StringFromLiteral(" ")), unrelated()))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatAny(hxrt.StringConcatAny(func() int {
		var hx_callable_31 *hxrt.Function[func(int) int] = New_Gauge(100).read
		var hx_argument_32 int = 20
		return hxrt.FunctionCall(hx_callable_31)(hx_argument_32)
	}(), hxrt.StringFromLiteral(" ")), New_Holder(3).total))
	hxrt.Println(v_11)
	var v_12 any = any(count)
	hxrt.Println(v_12)
}

var offset *hxrt.Function[func(int) int] = hxrt.NewFunction(func(helper_local int) int {
	return int((hxrt.Int32Wrap(helper_local) + hxrt.Int32Wrap(helper(helper_local, 1))))
})

func params(count_local int, helper_local int) *string {
	count = int((hxrt.Int32Wrap(count) + hxrt.Int32Wrap(count_local)))
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(count_local, hxrt.StringFromLiteral(",")), helper_local), hxrt.StringFromLiteral(",")), count), hxrt.StringFromLiteral(",")), helper(count_local, helper_local))
}

func unrelated() int {
	count_local := 9
	helper_local := 2
	return int((hxrt.Int32Wrap(count_local) + hxrt.Int32Wrap(helper_local)))
}

func Other_run(count_local int) int {
	helper_local := int((hxrt.Int32Wrap(count_local) * hxrt.Int32Wrap(2)))
	return int((hxrt.Int32Wrap(helper_local) + hxrt.Int32Wrap(helper(count_local, count))))
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
