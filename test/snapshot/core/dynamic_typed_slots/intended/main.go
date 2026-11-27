package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_Box interface {
}

type Box struct {
	__hx_this I_Box
	f         float64
	b         bool
	s         *string
	nf        any
}

func (self *Box) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Box(f float64) *Box {
	return hx_init_Box(hx_alloc_Box(nil), f)
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

func hx_init_Box(self *Box, f float64) *Box {
	self.nf = 0.5
	self.s = hxrt.StringFromLiteral("init")
	self.b = false
	self.f = f
	return self
}

func (self *Box) HxClassName() string {
	return "Box"
}

var always bool = true

func boolOf(d any) bool {
	return hxrt.DynamicToBool(d)
}

func boolSlots(d any) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("-- Bool slots, "), hxrt.StdString(d)))
	hxrt.Println(v)
	local := hxrt.DynamicToBool(d)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("local "), yesNo(local)))
	hxrt.Println(v_1)
	var o_n any
	o_n = d
	field := hxrt.DynamicToBool(o_n)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dynamic field "), yesNo(field)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("argument "), yesNo(hxrt.DynamicToBool(d))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("return "), yesNo(boolOf(d))))
	hxrt.Println(v_4)
	box := New_Box(0.5)
	box.b = hxrt.DynamicToBool(d)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field write "), yesNo(box.b)))
	hxrt.Println(v_5)
	var bs_0 bool
	bs_0 = false
	bs_0 = hxrt.DynamicToBool(d)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array write "), yesNo(bs_0)))
	hxrt.Println(v_6)
	var arr_0 any
	arr_0 = d
	element := hxrt.DynamicToBool(arr_0)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Array<Dynamic> element "), yesNo(element)))
	hxrt.Println(v_7)
}

var consumed *string = hxrt.StringFromLiteral("")

func dyn(value any) any {
	return value
}

func floatOf(d any) float64 {
	return hxrt.DynamicToFloatSlot(d)
}

func floatOfCall(d any) float64 {
	return hxrt.DynamicToFloatSlot(dyn(d))
}

func floatSlots(d any) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("-- Float slots, "), hxrt.StdString(d)))
	hxrt.Println(v)
	local := hxrt.DynamicToFloatSlot(d)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("local "), floatText(local)))
	hxrt.Println(v_1)
	var o_n any
	o_n = d
	field := hxrt.DynamicToFloatSlot(o_n)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dynamic field "), floatText(field)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("argument "), half(hxrt.DynamicToFloatSlot(d))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("return "), floatText(floatOf(d))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("return of call "), floatText(floatOfCall(d))))
	hxrt.Println(v_5)
	box := New_Box(hxrt.DynamicToFloatSlot(d))
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("constructor "), floatText(box.f)))
	hxrt.Println(v_6)
	box.f = 0.5
	box.f = hxrt.DynamicToFloatSlot(d)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field write "), floatText(box.f)))
	hxrt.Println(v_7)
	staticFloat = hxrt.DynamicToFloatSlot(d)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("static write "), floatText(staticFloat)))
	hxrt.Println(v_8)
	reassigned := 0.5
	reassigned = hxrt.DynamicToFloatSlot(d)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("reassign "), floatText(reassigned)))
	hxrt.Println(v_9)
	var fs_0 float64
	fs_0 = 0.5
	fs_0 = hxrt.DynamicToFloatSlot(d)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array write "), floatText(fs_0)))
	hxrt.Println(v_10)
	var literal_1 float64
	_ = literal_1
	var literal_0 float64
	literal_0 = hxrt.DynamicToFloatSlot(d)
	literal_1 = 0.5
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array literal "), floatText(literal_0)))
	hxrt.Println(v_11)
	m := New_haxe__ds__StringMap()
	v_12 := hxrt.DynamicToFloatSlot(d)
	m.__hx_this.set(hxrt.StringFromLiteral("k"), v_12)
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("map write "), floatText(func(hx_value_1 any) float64 {
		if hx_value_1 == nil {
			var hx_zero_2 float64
			return hx_zero_2
		}
		return hx_value_1.(float64)
	}(hxrt.DynamicToNullableFloat(m.__hx_this.get(hxrt.StringFromLiteral("k")))))))
	hxrt.Println(v_13)
	var st_f float64
	st_f = hxrt.DynamicToFloatSlot(d)
	var v_14 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("structure "), floatText(st_f)))
	hxrt.Println(v_14)
	var hx_if_3 float64
	if always {
		hx_if_3 = hxrt.DynamicToFloatSlot(d)
	} else {
		hx_if_3 = 0.5
	}
	picked := hx_if_3
	var v_15 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ternary "), floatText(picked)))
	hxrt.Println(v_15)
	casted := hxrt.DynamicToFloatSlot(d)
	var v_16 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("cast "), floatText(casted)))
	hxrt.Println(v_16)
	var arr_0 any
	arr_0 = d
	element := hxrt.DynamicToFloatSlot(arr_0)
	var v_17 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Array<Dynamic> element "), floatText(element)))
	hxrt.Println(v_17)
	reflected := hxrt.DynamicToFloatSlot(Reflect_field(func() hxrt.Anonymous {
		hx_obj_4 := map[string]any{}
		hx_obj_4["n"] = d
		return hxrt.AnonymousFromMap(hx_obj_4)
	}(), hxrt.StringFromLiteral("n")))
	var v_18 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Reflect.field "), floatText(reflected)))
	hxrt.Println(v_18)
	called := hxrt.DynamicToFloatSlot(dyn(d))
	var v_19 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("call "), floatText(called)))
	hxrt.Println(v_19)
	var nullable any = hxrt.DynamicToNullableFloat(d)
	unwrapped := func(hx_value_5 any) float64 {
		if hx_value_5 == nil {
			var hx_zero_6 float64
			return hx_zero_6
		}
		return hx_value_5.(float64)
	}(nullable)
	var v_20 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Null<Float> then Float "), floatText(unwrapped)))
	hxrt.Println(v_20)
	lambda := hxrt.NewFunction(func(x float64) float64 {
		return (x / float64(2))
	})
	var v_21 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("lambda argument "), hxrt.StdString(func() float64 {
		var hx_callable_7 *hxrt.Function[func(float64) float64] = lambda
		var hx_argument_8 float64 = hxrt.DynamicToFloatSlot(d)
		return hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
	}())))
	hxrt.Println(v_21)
}

func floatText(f float64) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(f), hxrt.StringFromLiteral(" ")), hxrt.StdString((f/float64(2)))), hxrt.StringFromLiteral(" ")), hxrt.StdString((f+0.25))), hxrt.StringFromLiteral(" ")), hxrt.StdString((f == 3)))
}

func half(f float64) *string {
	return hxrt.StdString((f / float64(2)))
}

func handled(label *string, run *hxrt.Function[func() *string]) {
	hxrt.TryCatch(func() {
		consumed = func() *string {
			var hx_callable_12 *hxrt.Function[func() *string] = run
			return hxrt.FunctionCall(hx_callable_12)()
		}()
	}, func(hx_caught_9 any) {
		hx_exception_10 := haxe__Exception_caught(hx_caught_9)
		hx_payload_11 := haxe__Exception_catchValue(hx_caught_9)
		_ = hx_exception_10
		var e any = hx_payload_11
		_ = e
	})
	hxrt.Println(any(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(" handled"))))
}

func incompatibleValues() {
	hxrt.Println(any(hxrt.StringFromLiteral("-- values that the slot cannot hold")))
	values := hxrt.NewArray(hxrt.StringFromLiteral("x"), true, func() hxrt.Anonymous {
		hx_obj_13 := map[string]any{}
		hx_obj_13["a"] = 1
		return hxrt.AnonymousFromMap(hx_obj_13)
	}())
	_g := 0
	for _g < values.Len() {
		var d any = values.Get(_g)
		_g = int(int32((_g + 1)))
		handled(hxrt.StringFromLiteral("Float local"), hxrt.NewFunction(func() *string {
			f := hxrt.DynamicToFloatSlot(d)
			return hxrt.StdString(f)
		}))
		handled(hxrt.StringFromLiteral("Float argument"), hxrt.NewFunction(func() *string {
			return half(hxrt.DynamicToFloatSlot(d))
		}))
		handled(hxrt.StringFromLiteral("Float return"), hxrt.NewFunction(func() *string {
			return hxrt.StdString(floatOf(d))
		}))
		handled(hxrt.StringFromLiteral("Null<Float> local"), hxrt.NewFunction(func() *string {
			var v any = hxrt.DynamicToNullableFloat(d)
			return hxrt.StdString(v)
		}))
	}
	others := hxrt.NewArray(1, 2.5, hxrt.StringFromLiteral("x"), func() hxrt.Anonymous {
		hx_obj_14 := map[string]any{}
		hx_obj_14["a"] = 1
		return hxrt.AnonymousFromMap(hx_obj_14)
	}())
	_g_1 := 0
	for _g_1 < others.Len() {
		var d_1 any = others.Get(_g_1)
		_g_1 = int(int32((_g_1 + 1)))
		handled(hxrt.StringFromLiteral("Bool local"), hxrt.NewFunction(func() *string {
			b := hxrt.DynamicToBool(d_1)
			return hxrt.StdString(b)
		}))
		handled(hxrt.StringFromLiteral("Bool argument"), hxrt.NewFunction(func() *string {
			return yesNo(hxrt.DynamicToBool(d_1))
		}))
		handled(hxrt.StringFromLiteral("Bool return"), hxrt.NewFunction(func() *string {
			return hxrt.StdString(boolOf(d_1))
		}))
		handled(hxrt.StringFromLiteral("Bool field write"), hxrt.NewFunction(func() *string {
			box := New_Box(0.5)
			box.b = hxrt.DynamicToBool(d_1)
			return hxrt.StdString(box.b)
		}))
		handled(hxrt.StringFromLiteral("Null<Bool> local"), hxrt.NewFunction(func() *string {
			var v any = hxrt.DynamicToNullableBool(d_1)
			return hxrt.StdString(v)
		}))
	}
	hxrt.Println(any(hxrt.StringFromLiteral("done")))
}

func isNull(s *string) bool {
	return hxrt.StringEqualStringPtr(s, nil)
}

func main() {
	floatSlots(3)
	floatSlots(2.5)
	boolSlots(true)
	stringSlots(nil)
	stringSlots(hxrt.StringFromLiteral("x"))
	otherValuesInStringSlots()
	nullableSlots()
	incompatibleValues()
}

func nullableFloatOf(d any) any {
	return hxrt.DynamicToNullableFloat(d)
}

func nullableSlots() {
	hxrt.Println(any(hxrt.StringFromLiteral("-- Null<Float> and Null<Bool> slots")))
	values := hxrt.NewArray(3, 2.5, nil)
	_g := 0
	for _g < values.Len() {
		var d any = values.Get(_g)
		_g = int(int32((_g + 1)))
		var v any = hxrt.DynamicToNullableFloat(d)
		var returned any = nullableFloatOf(d)
		var hx_if_15 *string
		if v == nil {
			hx_if_15 = hxrt.StringFromLiteral("null")
		} else {
			hx_if_15 = hxrt.StdString((v.(float64) + 0.5))
		}
		text := hx_if_15
		var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Null<Float> "), hxrt.StdString(v)), hxrt.StringFromLiteral(" ")), text), hxrt.StringFromLiteral(" ")), hxrt.StdString((returned == nil))))
		hxrt.Println(v_1)
		box := New_Box(0.5)
		box.nf = hxrt.DynamicToNullableFloat(d)
		var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Null<Float> field "), hxrt.StdString(box.nf)), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.AnyEqualsNull(box.nf))))
		hxrt.Println(v_2)
	}
	flags := hxrt.NewArray(true, false, nil)
	_g_1 := 0
	for _g_1 < flags.Len() {
		var d_1 any = flags.Get(_g_1)
		_g_1 = int(int32((_g_1 + 1)))
		var v_3 any = hxrt.DynamicToNullableBool(d_1)
		var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Null<Bool> "), hxrt.StdString(v_3)), hxrt.StringFromLiteral(" ")), hxrt.StdString((v_3 == nil))), hxrt.StringFromLiteral(" ")), func() *string {
			var hx_if_16 *string
			if v_3 == true {
				hx_if_16 = hxrt.StringFromLiteral("yes")
			} else {
				hx_if_16 = hxrt.StringFromLiteral("no")
			}
			return hx_if_16
		}()))
		hxrt.Println(v_4)
	}
}

func otherValuesInStringSlots() {
	hxrt.Println(any(hxrt.StringFromLiteral("-- Int, Float and Bool values in String slots")))
	values := hxrt.NewArray(3, 2.5, true)
	_g := 0
	for _g < values.Len() {
		var d any = values.Get(_g)
		_g = int(int32((_g + 1)))
		s := hxrt.DynamicToString(d)
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(s), hxrt.StringFromLiteral(" ")), hxrt.StringConcatStringPtr(s, hxrt.StringFromLiteral("!"))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.StringEqualStringPtr(s, nil))))
		hxrt.Println(v)
	}
}

var staticFloat float64 = 0.5

var staticString *string = hxrt.StringFromLiteral("init")

func stringOf(d any) *string {
	return hxrt.DynamicToString(d)
}

func stringOfCall(d any) *string {
	return hxrt.DynamicToString(dyn(d))
}

func stringSlots(d any) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("-- String slots, "), hxrt.StdString(d)))
	hxrt.Println(v)
	local := hxrt.DynamicToString(d)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("local "), stringText(local)))
	hxrt.Println(v_1)
	var o_n any
	o_n = d
	field := hxrt.DynamicToString(o_n)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dynamic field "), stringText(field)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("argument "), hxrt.StdString(isNull(hxrt.DynamicToString(d)))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("return "), stringText(stringOf(d))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("return of call "), stringText(stringOfCall(d))))
	hxrt.Println(v_5)
	box := New_Box(0.5)
	box.s = hxrt.DynamicToString(d)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field write "), stringText(box.s)))
	hxrt.Println(v_6)
	staticString = hxrt.DynamicToString(d)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("static write "), stringText(staticString)))
	hxrt.Println(v_7)
	reassigned := hxrt.StringFromLiteral("init")
	reassigned = hxrt.DynamicToString(d)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("reassign "), stringText(reassigned)))
	hxrt.Println(v_8)
	var ss_0 *string
	ss_0 = hxrt.StringFromLiteral("init")
	ss_0 = hxrt.DynamicToString(d)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array write "), stringText(ss_0)))
	hxrt.Println(v_9)
	m := New_haxe__ds__StringMap()
	v_10 := hxrt.DynamicToString(d)
	m.__hx_this.set(hxrt.StringFromLiteral("k"), v_10)
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("map write "), stringText(func(hx_value_17 any) *string {
		if hx_value_17 == nil {
			var hx_zero_18 *string
			return hx_zero_18
		}
		return hx_value_17.(*string)
	}(m.__hx_this.get(hxrt.StringFromLiteral("k"))))))
	hxrt.Println(v_11)
	var st_s *string
	st_s = hxrt.DynamicToString(d)
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("structure "), stringText(st_s)))
	hxrt.Println(v_12)
	var hx_if_19 *string
	if always {
		hx_if_19 = hxrt.DynamicToString(d)
	} else {
		hx_if_19 = hxrt.StringFromLiteral("other")
	}
	picked := hx_if_19
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ternary "), stringText(picked)))
	hxrt.Println(v_13)
	casted := hxrt.DynamicToString(d)
	var v_14 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("cast "), stringText(casted)))
	hxrt.Println(v_14)
	var arr_0 any
	arr_0 = d
	element := hxrt.DynamicToString(arr_0)
	var v_15 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Array<Dynamic> element "), stringText(element)))
	hxrt.Println(v_15)
	reflected := hxrt.DynamicToString(Reflect_field(func() hxrt.Anonymous {
		hx_obj_20 := map[string]any{}
		hx_obj_20["n"] = d
		return hxrt.AnonymousFromMap(hx_obj_20)
	}(), hxrt.StringFromLiteral("n")))
	var v_16 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Reflect.field "), stringText(reflected)))
	hxrt.Println(v_16)
	called := hxrt.DynamicToString(dyn(d))
	var v_17 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("call "), stringText(called)))
	hxrt.Println(v_17)
}

func stringText(s *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(hxrt.StringEqualStringPtr(s, nil)), hxrt.StringFromLiteral(" ")), hxrt.StdString(s))
}

func yesNo(b bool) *string {
	var hx_if_21 *string
	if b {
		hx_if_21 = hxrt.StringFromLiteral("yes")
	} else {
		hx_if_21 = hxrt.StringFromLiteral("no")
	}
	return hx_if_21
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
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
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
	case *haxe__ds__StringMap:
		return hxrt__generated_method_field__haxe__ds__StringMap(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value, key)
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

func hxrt__generated_method_field__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.clear()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "clear")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__StringMap] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *haxe__ds__StringMap {
				return hx_receiver_1.__hx_this.copy()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_1.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key *string) bool {
				return hx_receiver_1.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_1.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(*string) any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key *string) any {
				return hx_receiver_1.__hx_this.get(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_1.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.iterator()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.keys()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "keys")
		}()
	case "remove":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key *string) bool {
				return hx_receiver_1.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_1.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "removeIMap")
		}()
	case "set":
		return func() *hxrt.Function[func(*string, any)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key *string, value any) {
				hx_receiver_1.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_1.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setIMap")
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
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	return nil
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
	case *Box:
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
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Box:
		return hxrt__generated_field_lookup__Box(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_lookup__haxe__ds__StringMap(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Box:
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
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Box:
		return hxrt__generated_field_has__Box(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_has__haxe__ds__StringMap(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Box:
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
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Box:
		return hxrt__generated_field_set__Box(value, key, incoming)
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_set__haxe__ds__StringMap(value, key, incoming)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Box:
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
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *Box:
		return hxrt.NewArray(hxrt.StringFromLiteral("f"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("s"), hxrt.StringFromLiteral("nf"))
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__ds__StringMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("h"))
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("keys"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__Box(value *Box, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "b":
		return hxrt.ReadStoredField(value.b)
	case "f":
		return hxrt.ReadStoredField(value.f)
	case "nf":
		return hxrt.ReadStoredField(value.nf)
	case "s":
		return hxrt.ReadStoredField(value.s)
	}
	return nil
}

func hxrt__generated_field_has__Box(value *Box, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "b":
		return true
	case "f":
		return true
	case "nf":
		return true
	case "s":
		return true
	}
	return false
}

func hxrt__generated_field_set__Box(value *Box, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "b":
		if incoming == nil {
			var zero bool
			value.b = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.b = typed
			return true
		default:
			return false
		}
	case "f":
		if incoming == nil {
			var zero float64
			value.f = zero
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.f = typed
			return true
		case int:
			value.f = float64(typed)
			return true
		default:
			return false
		}
	case "nf":
		if incoming == nil {
			value.nf = nil
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.nf = typed
			return true
		case int:
			value.nf = float64(typed)
			return true
		default:
			return false
		}
	case "s":
		if incoming == nil {
			var zero *string
			value.s = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.s = typed
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

func hxrt__generated_field_lookup__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "h":
		return hxrt.ReadStoredField(value.h)
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__StringMap(value *haxe__ds__StringMap, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		if incoming == nil {
			var zero *hxrt.StringMapCell
			value.h = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.StringMapCell:
			value.h = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "keys":
		return hxrt.ReadStoredField(value.keys)
	case "map":
		return hxrt.ReadStoredField(value.map_)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		return true
	case "map":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		if incoming == nil {
			var zero hxrt.Anonymous
			value.keys = zero
			return true
		}
		switch typed := incoming.(type) {
		case hxrt.Anonymous:
			value.keys = typed
			return true
		default:
			return false
		}
	case "map":
		if incoming == nil {
			var zero haxe__IMap
			value.map_ = zero
			return true
		}
		switch typed := incoming.(type) {
		case haxe__IMap:
			value.map_ = typed
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
	case *hxrt.Function[func(float64) float64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_1 any) float64 {
			if hx_value_1 == nil {
				var hx_zero_2 float64
				return hx_zero_2
			}
			return hx_value_1.(float64)
		}(argument_0))
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() bool]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func()]:
		hxrt.FunctionCall(source)()
		return nil
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_3 any) *string {
			if hx_value_3 == nil {
				var hx_zero_4 *string
				return hx_zero_4
			}
			return hx_value_3.(*string)
		}(argument_0))
	case *hxrt.Function[func(any) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(*string) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(argument_0))
	case *hxrt.Function[func(any) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(*string, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_7 any) *string {
			if hx_value_7 == nil {
				var hx_zero_8 *string
				return hx_zero_8
			}
			return hx_value_7.(*string)
		}(argument_0), argument_1)
		return nil
	case *hxrt.Function[func(any, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(argument_0, argument_1)
		return nil
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
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
	case *hxrt.Function[func() *string]:
		return source
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
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(nil)
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
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
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() haxe__IMap]:
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
	case *hxrt.Function[func() bool]:
		return source
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
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_7 any) bool {
				if hx_value_7 == nil {
					var hx_zero_8 bool
					return hx_zero_8
				}
				return hx_value_7.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_9 any) bool {
				if hx_value_9 == nil {
					var hx_zero_10 bool
					return hx_zero_10
				}
				return hx_value_9.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_11 any) bool {
				if hx_value_11 == nil {
					var hx_zero_12 bool
					return hx_zero_12
				}
				return hx_value_11.(bool)
			}(nil)
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_13 any) bool {
				if hx_value_13 == nil {
					var hx_zero_14 bool
					return hx_zero_14
				}
				return hx_value_13.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_15 any) bool {
				if hx_value_15 == nil {
					var hx_zero_16 bool
					return hx_zero_16
				}
				return hx_value_15.(bool)
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
