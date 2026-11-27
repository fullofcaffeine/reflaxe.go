package main

import "snapshot/hxrt"

type Reading struct {
	tag    int
	params []any
}

func (self *Reading) HxEnumValue() (string, []any) {
	return [2]string{"Value", "Missing"}[self.tag], self.params
}

func Reading_Value(amount float64) *Reading {
	enumValue := &Reading{tag: 0}
	enumValue.params = []any{amount}
	return enumValue
}

var Reading_Missing *Reading = &Reading{tag: 1}

type I_Box interface {
}

type Box struct {
	__hx_this I_Box
	item      any
}

func (self *Box) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Box(item any) *Box {
	return hx_init_Box(hx_alloc_Box(nil), item)
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

func hx_init_Box(self *Box, item any) *Box {
	self.item = item
	return self
}

func (self *Box) HxClassName() string {
	return "Box"
}

type I_Doubler interface {
	scale(x float64) float64
}

type Doubler struct {
	__hx_this I_Doubler
}

func (self *Doubler) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Doubler() *Doubler {
	return hx_init_Doubler(hx_alloc_Doubler(nil))
}

func hx_alloc_Doubler(owner any) *Doubler {
	self := &Doubler{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Doubler)
	}
	return self
}

func hx_init_Doubler(self *Doubler) *Doubler {
	return self
}

func (self *Doubler) scale(x float64) float64 {
	return (x * float64(2))
}

func (self *Doubler) HxClassName() string {
	return "Doubler"
}

type I_FloatBox interface {
}

type FloatBox struct {
	*Box
	__hx_this I_FloatBox
}

func (self *FloatBox) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_FloatBox(n int) *FloatBox {
	return hx_init_FloatBox(hx_alloc_FloatBox(nil), n)
}

func hx_alloc_FloatBox(owner any) *FloatBox {
	self := &FloatBox{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_FloatBox)
	}
	self.Box = hx_alloc_Box(self.__hx_this)
	return self
}

func hx_init_FloatBox(self *FloatBox, n int) *FloatBox {
	hx_init_Box(self.Box, float64(n))
	return self
}

func (self *FloatBox) HxClassName() string {
	return "FloatBox"
}

type I_Holder interface {
	scale(factor float64) float64
}

type Holder struct {
	__hx_this I_Holder
	value     float64
	rate      float64
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder(value float64) *Holder {
	return hx_init_Holder(hx_alloc_Holder(nil), value)
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

func hx_init_Holder(self *Holder, value float64) *Holder {
	self.rate = 2
	self.value = value
	return self
}

func (self *Holder) scale(factor float64) float64 {
	return ((self.value * factor) * self.rate)
}

func (self *Holder) HxClassName() string {
	return "Holder"
}

var Holder_base float64 = 1

var Holder_total float64 = 0

func amount(reading *Reading) float64 {
	var hx_switch_1 float64
	switch reading.tag {
	case 0:
		_g := reading.params[0].(float64)
		value := _g
		hx_switch_1 = (value / float64(2))
	case 1:
		hx_switch_1 = -1
	}
	return hx_switch_1
}

func describe(value any) *string {
	var hx_if_2 *string
	if value == nil {
		hx_if_2 = hxrt.StringFromLiteral("null")
	} else {
		hx_if_2 = hxrt.StdString((value.(float64) / float64(4)))
	}
	return hx_if_2
}

func literal() float64 {
	return 4
}

func main() {
	n := 3
	half := hxrt.NewFunction(func(k int) float64 {
		return float64(k)
	})
	arrow := hxrt.NewFunction(func(k int) float64 {
		return float64(k)
	})
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("return="), (plain(n)/float64(2))), hxrt.StringFromLiteral(" literal=")), (literal()/float64(8))), hxrt.StringFromLiteral(" if=")), (viaIf(n, true)/float64(2))), hxrt.StringFromLiteral(":")), viaIf(n, false)), hxrt.StringFromLiteral(" switch=")), (viaSwitch(n)/float64(2))), hxrt.StringFromLiteral(":")), viaSwitch(0)), hxrt.StringFromLiteral(" closure=")), (func() float64 {
		var hx_callable_3 *hxrt.Function[func(int) float64] = half
		var hx_argument_4 int = n
		return hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}()/float64(2))), hxrt.StringFromLiteral(" arrow=")), (func() float64 {
		var hx_callable_5 *hxrt.Function[func(int) float64] = arrow
		var hx_argument_6 int = 5
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()/float64(2))), hxrt.StringFromLiteral(" string=")), hxrt.StdString((plain(n) / float64(2)))))
	hxrt.Println(v)
	x := 1.5
	x = float64(n)
	fromLiteral := 1.0
	fromLiteral = (fromLiteral + 0.5)
	fromInt := float64(n)
	Holder_total = float64(n)
	Holder_total = (Holder_total + float64(n))
	Holder_base = 5
	holder := New_Holder(float64(n))
	holder.value = float64(n)
	holder.value = (holder.value + float64(n))
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("assign="), (x/float64(2))), hxrt.StringFromLiteral(" literal=")), fromLiteral), hxrt.StringFromLiteral(" local=")), (fromInt/float64(2))), hxrt.StringFromLiteral(" static=")), (Holder_total/float64(4))), hxrt.StringFromLiteral(":")), (Holder_base/float64(2))), hxrt.StringFromLiteral(" field=")), (holder.value/float64(4))), hxrt.StringFromLiteral(" method=")), (holder.__hx_this.scale(float64(n)) / float64(4))))
	hxrt.Println(v_1)
	scaleBy := hxrt.NewFunction(func(f float64) float64 {
		return (f * 1.5)
	})
	var scaler Scaler = New_Doubler()
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("args="), (sum(float64(n), 2, nil)/float64(2))), hxrt.StringFromLiteral(":")), (sum(float64(n), float64(n), nil)/float64(2))), hxrt.StringFromLiteral(":")), (sum(float64(n), float64(n), float64(n))/float64(2))), hxrt.StringFromLiteral(":")), (sum(1, 2.0, 3.0)/float64(4))), hxrt.StringFromLiteral(" rest=")), total([]float64{float64(n), 1.0, float64(n)})), hxrt.StringFromLiteral(":")), total([]float64{float64(n), 2.0})), hxrt.StringFromLiteral(" value=")), func() float64 {
		var hx_callable_7 *hxrt.Function[func(float64) float64] = scaleBy
		var hx_argument_8 float64 = float64(n)
		return hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
	}()), hxrt.StringFromLiteral(" interface=")), (scaler.scale(float64(n))/float64(4))), hxrt.StringFromLiteral(" math=")), (Math_max(float64(n), 0.5)/float64(2))), hxrt.StringFromLiteral(" nullable=")), describe(float64(n))), hxrt.StringFromLiteral(":")), describe(nil)))
	hxrt.Println(v_2)
	sample := New_Sample(float64(n), hxrt.StringFromLiteral("w"))
	box := New_Box(float64(n))
	boxed := (hxrt.DynamicToFloatSlot(box.item) / float64(2))
	box.item = 5.0
	reading := Reading_Value(float64(n))
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("ctor="), (New_Holder(float64(n)).value/float64(2))), hxrt.StringFromLiteral(" super=")), (New_Weighted(n).value/float64(2))), hxrt.StringFromLiteral(":")), (hxrt.DynamicToFloatSlot(New_FloatBox(n).item)/float64(2))), hxrt.StringFromLiteral(" struct=")), (sample.weight/float64(2))), hxrt.StringFromLiteral(" box=")), boxed), hxrt.StringFromLiteral(":")), (hxrt.DynamicToFloatSlot(box.item)/float64(2))), hxrt.StringFromLiteral(" enum=")), amount(reading)), hxrt.StringFromLiteral(":")), amount(Reading_Value(7))), hxrt.StringFromLiteral(":")), amount(Reading_Missing)))
	hxrt.Println(v_3)
	values := hxrt.NewArray(1.0, float64(n), 2.5)
	hx_array_target_9 := values
	hx_array_index_10 := 0
	hx_array_target_9.Set(hx_array_index_10, float64(n))
	values.Push(float64(n))
	values.Push(4.0)
	values.Insert(0, float64(n))
	func() {
		hx_insert_position_15 := 1
		var hx_insert_value_16 any = 7.0
		values.Insert(hx_insert_position_15, hx_insert_value_16)
	}()
	func() float64 {
		hx_array_target_17 := values
		hx_array_index_18 := 1
		var hx_array_current_19 float64 = hxrt.DynamicToFloatSlot(hx_array_target_17.Get(hx_array_index_18))
		var hx_array_assigned_20 float64 = (hx_array_current_19 + float64(n))
		hx_array_target_17.Set(hx_array_index_18, hx_array_assigned_20)
		return hx_array_assigned_20
	}()
	halves := 0.0
	_g := 0
	for _g < values.Len() {
		value := hxrt.DynamicToFloatSlot(values.Get(_g))
		_g = int(int32((_g + 1)))
		halves = (halves + (value / float64(2)))
	}
	weights := New_haxe__ds__StringMap()
	weights.__hx_this.set(hxrt.StringFromLiteral("a"), float64(n))
	weights.__hx_this.set(hxrt.StringFromLiteral("b"), float64(n))
	weights.__hx_this.set(hxrt.StringFromLiteral("c"), 2.0)
	_g_1 := New_haxe__ds__StringMap()
	_g_1.__hx_this.set(hxrt.StringFromLiteral("d"), float64(n))
	_g_1.__hx_this.set(hxrt.StringFromLiteral("e"), 1.0)
	literalMap := _g_1
	optional := hxrt.NewArray(float64(n), nil)
	optional.Push(2.0)
	var this1 []float64
	this1 = []float64{}
	hx_len_22 := 2
	if hx_len_22 < 0 {
		hx_len_22 = 0
	}
	if hx_len_22 <= len(this1) {
		this1 = this1[:hx_len_22]
	} else {
		var hx_zero_23 float64
		for len(this1) < hx_len_22 {
			this1 = append(this1, hx_zero_23)
		}
	}
	vector := this1
	vector[0] = float64(n)
	vector[1] = 1
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(","))), hxrt.StringFromLiteral(" halves=")), halves), hxrt.StringFromLiteral(" first=")), (hxrt.DynamicToFloatSlot(values.Get(0))/float64(2))), hxrt.StringFromLiteral(" find=")), func() int {
		hx_indexof_target_24 := values
		var hx_indexof_value_25 float64 = float64(n)
		hx_indexof_length_26 := hx_indexof_target_24.Len()
		hx_indexof_index_27 := 0
		for hx_indexof_index_27 < hx_indexof_length_26 {
			hx_indexof_element_28 := hxrt.DynamicToFloatSlot(hx_indexof_target_24.Get(hx_indexof_index_27))
			if hx_indexof_element_28 == hx_indexof_value_25 {
				return hx_indexof_index_27
			}
			hx_indexof_index_27 = (hx_indexof_index_27 + 1)
		}
		return -1
	}()), hxrt.StringFromLiteral(":")), hxrt.StdString(func() bool {
		hx_contains_target_29 := values
		var hx_contains_value_30 float64 = 7.0
		hx_contains_length_31 := hx_contains_target_29.Len()
		hx_contains_index_32 := 0
		for hx_contains_index_32 < hx_contains_length_31 {
			hx_contains_element_33 := hxrt.DynamicToFloatSlot(hx_contains_target_29.Get(hx_contains_index_32))
			if hx_contains_element_33 == hx_contains_value_30 {
				return true
			}
			hx_contains_index_32 = (hx_contains_index_32 + 1)
		}
		return false
	}())), hxrt.StringFromLiteral(" optional=")), (func(hx_value_34 any) float64 {
		if hx_value_34 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_34.(float64)
	}(hxrt.DynamicToNullableFloat(optional.Get(0)))/float64(2))), hxrt.StringFromLiteral(":")), optional.Get(1)), hxrt.StringFromLiteral(":")), (func(hx_value_35 any) float64 {
		if hx_value_35 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_35.(float64)
	}(hxrt.DynamicToNullableFloat(optional.Get(2)))/float64(4))), hxrt.StringFromLiteral(" vector=")), (vector[0]/float64(2))), hxrt.StringFromLiteral(":")), (vector[1]/float64(4))), hxrt.StringFromLiteral(" map=")), (func(hx_value_36 any) float64 {
		if hx_value_36 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_36.(float64)
	}(hxrt.DynamicToNullableFloat(weights.__hx_this.get(hxrt.StringFromLiteral("a"))))/float64(2))), hxrt.StringFromLiteral(":")), (func(hx_value_37 any) float64 {
		if hx_value_37 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_37.(float64)
	}(hxrt.DynamicToNullableFloat(weights.__hx_this.get(hxrt.StringFromLiteral("b"))))*1.5)), hxrt.StringFromLiteral(":")), (func(hx_value_38 any) float64 {
		if hx_value_38 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_38.(float64)
	}(hxrt.DynamicToNullableFloat(weights.__hx_this.get(hxrt.StringFromLiteral("c"))))/float64(4))), hxrt.StringFromLiteral(":")), (func(hx_value_39 any) float64 {
		if hx_value_39 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_39.(float64)
	}(hxrt.DynamicToNullableFloat(literalMap.__hx_this.get(hxrt.StringFromLiteral("d"))))/float64(2))), hxrt.StringFromLiteral(":")), (func(hx_value_40 any) float64 {
		if hx_value_40 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_40.(float64)
	}(hxrt.DynamicToNullableFloat(literalMap.__hx_this.get(hxrt.StringFromLiteral("e")))) / float64(4))))
	hxrt.Println(v_4)
	hx_obj_41 := map[string]any{}
	hx_obj_41["x"] = float64(n)
	hx_obj_41["y"] = 1.0
	p := hxrt.AnonymousFromMap(hx_obj_41)
	made := makePoint(n)
	var p2_y float64
	var p2_x float64
	p2_x = 1
	p2_y = 2
	p2_x = float64(n)
	p2_y = (p2_y + float64(n))
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("anon="), norm(p)), hxrt.StringFromLiteral(":")), norm(func() hxrt.Anonymous {
		hx_obj_42 := map[string]any{}
		hx_obj_42["x"] = 1.0
		hx_obj_42["y"] = float64(n)
		return hxrt.AnonymousFromMap(hx_obj_42)
	}())), hxrt.StringFromLiteral(":")), norm(made)), hxrt.StringFromLiteral(":")), (hxrt.DynamicToFloatSlot(hxrt.AnonymousGet(made, "x"))/float64(2))), hxrt.StringFromLiteral(":")), (p2_x/float64(2))), hxrt.StringFromLiteral(":")), (p2_y / float64(2))))
	hxrt.Println(v_5)
	var hx_if_43 float64
	if n > 1 {
		hx_if_43 = float64(n)
	} else {
		hx_if_43 = 0.5
	}
	ternary := hx_if_43
	var hx_if_44 float64
	if n == 3 {
		hx_if_44 = float64(n)
	} else {
		hx_if_44 = 0.75
	}
	switched := hx_if_44
	var hx_if_45 float64
	if n > 1 {
		hx_if_45 = float64(n)
	} else {
		hx_if_45 = 0.5
	}
	mixed := hx_if_45
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("arms="), (ternary/float64(2))), hxrt.StringFromLiteral(":")), (switched/float64(2))), hxrt.StringFromLiteral(":")), (mixed / float64(2)))))
	var nf any = float64(n)
	var nfLiteral any = 4.0
	var ni any = n
	var nfFromNi any = func(hx_value_46 any) any {
		if hx_value_46 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_46))
	}(ni)
	var reassigned any = nil
	if n > 0 {
		reassigned = int((hxrt.Int32Wrap(n) + hxrt.Int32Wrap(1)))
	}
	var nfFromReassigned any = func(hx_value_47 any) any {
		if hx_value_47 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_47))
	}(reassigned)
	var noInt any = nil
	var nfNull any = func(hx_value_48 any) any {
		if hx_value_48 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_48))
	}(noInt)
	fromNi := float64(hxrt.IntFromNullableAny(ni))
	fromReassigned := float64(hxrt.IntFromNullableAny(reassigned))
	var v_6 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("nullable="), (func(hx_value_49 any) float64 {
		if hx_value_49 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_49.(float64)
	}(nf)+0.5)), hxrt.StringFromLiteral(":")), (func(hx_value_50 any) float64 {
		if hx_value_50 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_50.(float64)
	}(nfLiteral)/float64(8))), hxrt.StringFromLiteral(":")), (func(hx_value_51 any) float64 {
		if hx_value_51 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_51.(float64)
	}(nfFromNi)/float64(2))), hxrt.StringFromLiteral(":")), (func(hx_value_52 any) float64 {
		if hx_value_52 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_52.(float64)
	}(nfFromReassigned)/float64(2))), hxrt.StringFromLiteral(":")), hxrt.StdString((nfNull == nil))), hxrt.StringFromLiteral(":")), (fromNi/float64(2))), hxrt.StringFromLiteral(":")), (fromReassigned/float64(2))), hxrt.StringFromLiteral(":")), (func(hx_value_53 any) float64 {
		if hx_value_53 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_53.(float64)
	}(maybe(n, true))/float64(2))), hxrt.StringFromLiteral(":")), maybe(n, false)), hxrt.StringFromLiteral(":")), (func(hx_value_54 any) float64 {
		if hx_value_54 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_54.(float64)
	}(maybeLiteral()) / float64(4))))
	hxrt.Println(v_6)
}

func makePoint(n int) hxrt.Anonymous {
	hx_obj_55 := map[string]any{}
	hx_obj_55["x"] = float64(n)
	hx_obj_55["y"] = 1.0
	return hxrt.AnonymousFromMap(hx_obj_55)
}

func maybe(n int, present bool) any {
	if !present {
		return nil
	}
	return float64(n)
}

func maybeLiteral() any {
	return 6.0
}

func norm(p hxrt.Anonymous) float64 {
	return ((hxrt.DynamicToFloatSlot(hxrt.AnonymousGet(p, "x")) * hxrt.DynamicToFloatSlot(hxrt.AnonymousGet(p, "x"))) + (hxrt.DynamicToFloatSlot(hxrt.AnonymousGet(p, "y")) / float64(2)))
}

func plain(n int) float64 {
	return float64(n)
}

func sum(a float64, b float64, c any) float64 {
	return ((a + func(hx_value_56 any) float64 {
		if hx_value_56 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_56.(float64)
	}(b)) + func(hx_value_58 any) float64 {
		if hx_value_58 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null"))
		}
		return hx_value_58.(float64)
	}(func() any {
		var hx_if_57 any
		if c == nil {
			hx_if_57 = 0.0
		} else {
			hx_if_57 = c.(float64)
		}
		return hx_if_57
	}()))
}

func total(values []float64) float64 {
	out := 0.0
	var _g_current int
	var _g_args []float64
	_g_current = 0
	_g_args = func(hx_value_59 any) []float64 {
		if hx_value_59 == nil {
			var hx_zero_60 []float64
			return hx_zero_60
		}
		return hx_value_59.([]float64)
	}(any(values))
	for _g_current < len(_g_args) {
		this1 := hxrt.ArrayFromValues(func(hx_sort_src_61 []float64) []any {
			hx_sort_out_63 := make([]any, 0, len(hx_sort_src_61))
			for _, hx_sort_item_62 := range hx_sort_src_61 {
				hx_sort_out_63 = append(hx_sort_out_63, hx_sort_item_62)
			}
			return hx_sort_out_63
		}(_g_args))
		hx_post_64 := _g_current
		_g_current = int(int32((_g_current + 1)))
		index := hx_post_64
		value := hxrt.DynamicToFloatSlot(this1.Get(index))
		out = (out + (value / float64(2)))
	}
	return out
}

func viaIf(n int, flag bool) float64 {
	var hx_if_65 float64
	if flag {
		hx_if_65 = float64(n)
	} else {
		hx_if_65 = 0.5
	}
	return hx_if_65
}

func viaSwitch(n int) float64 {
	var hx_if_66 float64
	if n == 0 {
		hx_if_66 = 0.25
	} else {
		hx_if_66 = float64(n)
	}
	return hx_if_66
}

type I_Sample interface {
}

type Sample struct {
	__hx_this I_Sample
	weight    float64
	label     *string
}

func (self *Sample) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Sample(weight float64, label *string) *Sample {
	return hx_init_Sample(hx_alloc_Sample(nil), weight, label)
}

func hx_alloc_Sample(owner any) *Sample {
	self := &Sample{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Sample)
	}
	return self
}

func hx_init_Sample(self *Sample, weight float64, label *string) *Sample {
	self.weight = weight
	self.label = label
	return self
}

func (self *Sample) HxClassName() string {
	return "Sample"
}

type Scaler interface {
	scale(x float64) float64
}

type I_Weighted interface {
	scale(factor float64) float64
}

type Weighted struct {
	*Holder
	__hx_this I_Weighted
}

func (self *Weighted) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Weighted(n int) *Weighted {
	return hx_init_Weighted(hx_alloc_Weighted(nil), n)
}

func hx_alloc_Weighted(owner any) *Weighted {
	self := &Weighted{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Weighted)
	}
	self.Holder = hx_alloc_Holder(self.__hx_this)
	return self
}

func hx_init_Weighted(self *Weighted, n int) *Weighted {
	hx_init_Holder(self.Holder, float64(n))
	return self
}

func (self *Weighted) HxClassName() string {
	return "Weighted"
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
