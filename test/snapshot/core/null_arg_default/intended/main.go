package main

import "snapshot/hxrt"

type Shade struct {
	tag    int
	params []any
}

func (self *Shade) HxEnumValue() (string, []any) {
	return [2]string{"Light", "Dark"}[self.tag], self.params
}

var Shade_Light *Shade = &Shade{tag: 0}

var Shade_Dark *Shade = &Shade{tag: 1}

type I_BasePricer interface {
	price(rate any, count any, label *string) *string
}

type BasePricer struct {
	__hx_this I_BasePricer
}

func (self *BasePricer) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_BasePricer() *BasePricer {
	return hx_init_BasePricer(hx_alloc_BasePricer(nil))
}

func hx_alloc_BasePricer(owner any) *BasePricer {
	self := &BasePricer{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_BasePricer)
	}
	return self
}

func hx_init_BasePricer(self *BasePricer) *BasePricer {
	return self
}

func (self *BasePricer) price(rate any, count any, label *string) *string {
	if rate == nil {
		rate = 2.5
	}
	if count == nil {
		count = 4
	}
	if label == nil {
		label = hxrt.StringFromLiteral("base")
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("base("), rate.(float64)), hxrt.StringFromLiteral(",")), count.(int)), hxrt.StringFromLiteral(",")), label), hxrt.StringFromLiteral(")"))
}

func (self *BasePricer) HxClassName() string {
	return "BasePricer"
}

type I_ChildPricer interface {
	price(rate any, count any, label *string) *string
}

type ChildPricer struct {
	*BasePricer
	__hx_this I_ChildPricer
}

func (self *ChildPricer) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_ChildPricer() *ChildPricer {
	return hx_init_ChildPricer(hx_alloc_ChildPricer(nil))
}

func hx_alloc_ChildPricer(owner any) *ChildPricer {
	self := &ChildPricer{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_ChildPricer)
	}
	self.BasePricer = hx_alloc_BasePricer(self.__hx_this)
	return self
}

func hx_init_ChildPricer(self *ChildPricer) *ChildPricer {
	hx_init_BasePricer(self.BasePricer)
	return self
}

func (self *ChildPricer) price(rate any, count any, label *string) *string {
	if rate == nil {
		rate = 3.5
	}
	if count == nil {
		count = 6
	}
	if label == nil {
		label = hxrt.StringFromLiteral("child")
	}
	var noRate any = nil
	var noCount any = nil
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("child("), rate.(float64)), hxrt.StringFromLiteral(",")), count.(int)), hxrt.StringFromLiteral(",")), label), hxrt.StringFromLiteral(")>")), self.BasePricer.price(noRate, noCount, nil))
}

func (self *ChildPricer) HxClassName() string {
	return "ChildPricer"
}

type I_Counter interface {
	bump(by int, scale float64, on bool, tag *string) *string
}

type Counter struct {
	__hx_this I_Counter
	text      *string
}

func (self *Counter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Counter(start float64, step int, on bool, label *string) *Counter {
	return hx_init_Counter(hx_alloc_Counter(nil), start, step, on, label)
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

func hx_init_Counter(self *Counter, start float64, step int, on bool, label *string) *Counter {
	if label == nil {
		label = hxrt.StringFromLiteral("ctr")
	}
	self.text = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(start, hxrt.StringFromLiteral(",")), step), hxrt.StringFromLiteral(",")), hxrt.StdString(on)), hxrt.StringFromLiteral(",")), label)
	return self
}

func (self *Counter) bump(by int, scale float64, on bool, tag *string) *string {
	if tag == nil {
		tag = hxrt.StringFromLiteral("tag")
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(by, hxrt.StringFromLiteral(",")), scale), hxrt.StringFromLiteral(",")), hxrt.StdString(on)), hxrt.StringFromLiteral(",")), tag)
}

func (self *Counter) HxClassName() string {
	return "Counter"
}

type I_Holder interface {
}

type Holder struct {
	__hx_this I_Holder
	f         any
	i         any
	b         any
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder() *Holder {
	return hx_init_Holder(hx_alloc_Holder(nil))
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

func hx_init_Holder(self *Holder) *Holder {
	return self
}

func (self *Holder) HxClassName() string {
	return "Holder"
}

func forward(f any, i any, b any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(staticFloat(func(hx_value_1 any) float64 {
		if hx_value_1 == nil {
			return 1.5
		}
		return hx_value_1.(float64)
	}(f)), hxrt.StringFromLiteral(",")), staticInt(func(hx_value_2 any) int {
		if hx_value_2 == nil {
			return 3
		}
		return hxrt.IntFromNullableAny(hx_value_2)
	}(i))), hxrt.StringFromLiteral(",")), hxrt.StdString(staticBool(func(hx_value_3 any) bool {
		if hx_value_3 == nil {
			return true
		}
		return hx_value_3.(bool)
	}(b))))
}

func main() {
	var nf any = nil
	var ni any = nil
	var nb any = nil
	var ns *string = nil
	var nshade *Shade = nil
	var nlevel any = nil
	var zeroF any = 0.0
	var zeroI any = 0
	var falseB any = false
	empty := hxrt.StringFromLiteral("")
	row(hxrt.StringFromLiteral("static.float"), hxrt.NewArray(staticFloat(1.5), staticFloat(func(hx_value_4 any) float64 {
		if hx_value_4 == nil {
			return 1.5
		}
		return hx_value_4.(float64)
	}(nf)), staticFloat(func(hx_value_5 any) float64 {
		if hx_value_5 == nil {
			return 1.5
		}
		return hx_value_5.(float64)
	}(zeroF)), staticFloat(2.25)))
	row(hxrt.StringFromLiteral("static.int"), hxrt.NewArray(staticInt(3), staticInt(func(hx_value_6 any) int {
		if hx_value_6 == nil {
			return 3
		}
		return hxrt.IntFromNullableAny(hx_value_6)
	}(ni)), staticInt(func(hx_value_7 any) int {
		if hx_value_7 == nil {
			return 3
		}
		return hxrt.IntFromNullableAny(hx_value_7)
	}(zeroI)), staticInt(5)))
	row(hxrt.StringFromLiteral("static.bool"), hxrt.NewArray(staticBool(true), staticBool(func(hx_value_8 any) bool {
		if hx_value_8 == nil {
			return true
		}
		return hx_value_8.(bool)
	}(nb)), staticBool(func(hx_value_9 any) bool {
		if hx_value_9 == nil {
			return true
		}
		return hx_value_9.(bool)
	}(falseB)), staticBool(false)))
	row(hxrt.StringFromLiteral("static.string"), hxrt.NewArray(staticString(hxrt.StringFromLiteral("dflt")), staticString(nil), staticString(ns), staticString(empty), staticString(hxrt.StringFromLiteral("x"))))
	row(hxrt.StringFromLiteral("static.enum"), hxrt.NewArray(staticShade(Shade_Dark), staticShade(nil), staticShade(nshade), staticShade(Shade_Light)))
	row(hxrt.StringFromLiteral("static.level"), hxrt.NewArray(staticLevel(2), staticLevel(nlevel), staticLevel(any(1))))
	row(hxrt.StringFromLiteral("opt.float"), hxrt.NewArray(optFloat(4.5), optFloat(4.5), optFloat(func(hx_value_10 any) float64 {
		if hx_value_10 == nil {
			return 4.5
		}
		return hx_value_10.(float64)
	}(nf)), optFloat(func(hx_value_11 any) float64 {
		if hx_value_11 == nil {
			return 4.5
		}
		return hx_value_11.(float64)
	}(zeroF)), optFloat(0.75)))
	row(hxrt.StringFromLiteral("opt.int"), hxrt.NewArray(optInt(6), optInt(6), optInt(func(hx_value_12 any) int {
		if hx_value_12 == nil {
			return 6
		}
		return hxrt.IntFromNullableAny(hx_value_12)
	}(ni)), optInt(func(hx_value_13 any) int {
		if hx_value_13 == nil {
			return 6
		}
		return hxrt.IntFromNullableAny(hx_value_13)
	}(zeroI)), optInt(8)))
	row(hxrt.StringFromLiteral("opt.bool"), hxrt.NewArray(optBool(true), optBool(true), optBool(func(hx_value_14 any) bool {
		if hx_value_14 == nil {
			return true
		}
		return hx_value_14.(bool)
	}(nb)), optBool(func(hx_value_15 any) bool {
		if hx_value_15 == nil {
			return true
		}
		return hx_value_15.(bool)
	}(falseB)), optBool(false)))
	row(hxrt.StringFromLiteral("opt.string"), hxrt.NewArray(optString(hxrt.StringFromLiteral("opt")), optString(nil), optString(ns), optString(empty)))
	row(hxrt.StringFromLiteral("nullable.int"), hxrt.NewArray(nullInt(9), nullInt(9), nullInt(func(hx_value_16 any) int {
		if hx_value_16 == nil {
			return 9
		}
		return hxrt.IntFromNullableAny(hx_value_16)
	}(ni)), nullInt(func(hx_value_17 any) int {
		if hx_value_17 == nil {
			return 9
		}
		return hxrt.IntFromNullableAny(hx_value_17)
	}(zeroI)), nullInt(12)))
	row(hxrt.StringFromLiteral("mixed"), hxrt.NewArray(mixed(1, 0.5, true, hxrt.StringFromLiteral("d")), mixed(1, func(hx_value_18 any) float64 {
		if hx_value_18 == nil {
			return 0.5
		}
		return hx_value_18.(float64)
	}(nf), true, hxrt.StringFromLiteral("d")), mixed(1, func(hx_value_19 any) float64 {
		if hx_value_19 == nil {
			return 0.5
		}
		return hx_value_19.(float64)
	}(nf), func(hx_value_20 any) bool {
		if hx_value_20 == nil {
			return true
		}
		return hx_value_20.(bool)
	}(nb), hxrt.StringFromLiteral("d")), mixed(1, func(hx_value_21 any) float64 {
		if hx_value_21 == nil {
			return 0.5
		}
		return hx_value_21.(float64)
	}(nf), func(hx_value_22 any) bool {
		if hx_value_22 == nil {
			return true
		}
		return hx_value_22.(bool)
	}(nb), ns), mixed(1, func(hx_value_23 any) float64 {
		if hx_value_23 == nil {
			return 0.5
		}
		return hx_value_23.(float64)
	}(zeroF), func(hx_value_24 any) bool {
		if hx_value_24 == nil {
			return true
		}
		return hx_value_24.(bool)
	}(falseB), empty)))
	row(hxrt.StringFromLiteral("forward"), hxrt.NewArray(forward(nil, nil, nil), forward(nil, nil, nil), forward(nf, ni, nb), forward(0.0, 0, false)))
	holder := New_Holder()
	var nulls_1 any
	var nulls_0 any
	nulls_0 = nil
	nulls_1 = 0
	row(hxrt.StringFromLiteral("sources"), hxrt.NewArray(staticFloat(func(hx_value_25 any) float64 {
		if hx_value_25 == nil {
			return 1.5
		}
		return hx_value_25.(float64)
	}(holder.f)), staticInt(func(hx_value_26 any) int {
		if hx_value_26 == nil {
			return 3
		}
		return hxrt.IntFromNullableAny(hx_value_26)
	}(holder.i)), staticBool(func(hx_value_27 any) bool {
		if hx_value_27 == nil {
			return true
		}
		return hx_value_27.(bool)
	}(holder.b)), staticFloat(func(hx_value_28 any) float64 {
		if hx_value_28 == nil {
			return 1.5
		}
		return hx_value_28.(float64)
	}(nothingFloat())), staticInt(func(hx_value_29 any) int {
		if hx_value_29 == nil {
			return 3
		}
		return hxrt.IntFromNullableAny(hx_value_29)
	}(nothingInt())), staticInt(func(hx_value_30 any) int {
		if hx_value_30 == nil {
			return 3
		}
		return hxrt.IntFromNullableAny(hx_value_30)
	}(nulls_0)), staticInt(func(hx_value_31 any) int {
		if hx_value_31 == nil {
			return 3
		}
		return hxrt.IntFromNullableAny(hx_value_31)
	}(nulls_1))))
	holder.f = 0.0
	holder.i = 0
	holder.b = false
	row(hxrt.StringFromLiteral("sources.set"), hxrt.NewArray(staticFloat(func(hx_value_32 any) float64 {
		if hx_value_32 == nil {
			return 1.5
		}
		return hx_value_32.(float64)
	}(holder.f)), staticInt(func(hx_value_33 any) int {
		if hx_value_33 == nil {
			return 3
		}
		return hxrt.IntFromNullableAny(hx_value_33)
	}(holder.i)), staticBool(func(hx_value_34 any) bool {
		if hx_value_34 == nil {
			return true
		}
		return hx_value_34.(bool)
	}(holder.b))))
	if nf == nil {
		nf = 6.5
	}
	if ni != nil {
		row(hxrt.StringFromLiteral("narrowed"), hxrt.NewArray(staticInt(ni.(int))))
	}
	row(hxrt.StringFromLiteral("assigned"), hxrt.NewArray(staticFloat(func(hx_value_35 any) float64 {
		if hx_value_35 == nil {
			return 1.5
		}
		return hx_value_35.(float64)
	}(nf))))
	nf = nil
	counter := New_Counter(0.25, 2, true, hxrt.StringFromLiteral("ctr"))
	row(hxrt.StringFromLiteral("method"), hxrt.NewArray(counter.__hx_this.bump(10, 1.5, true, hxrt.StringFromLiteral("tag")), counter.__hx_this.bump(func(hx_value_36 any) int {
		if hx_value_36 == nil {
			return 10
		}
		return hxrt.IntFromNullableAny(hx_value_36)
	}(ni), 1.5, true, hxrt.StringFromLiteral("tag")), counter.__hx_this.bump(func(hx_value_37 any) int {
		if hx_value_37 == nil {
			return 10
		}
		return hxrt.IntFromNullableAny(hx_value_37)
	}(ni), func(hx_value_38 any) float64 {
		if hx_value_38 == nil {
			return 1.5
		}
		return hx_value_38.(float64)
	}(nf), true, hxrt.StringFromLiteral("tag")), counter.__hx_this.bump(func(hx_value_39 any) int {
		if hx_value_39 == nil {
			return 10
		}
		return hxrt.IntFromNullableAny(hx_value_39)
	}(ni), func(hx_value_40 any) float64 {
		if hx_value_40 == nil {
			return 1.5
		}
		return hx_value_40.(float64)
	}(nf), func(hx_value_41 any) bool {
		if hx_value_41 == nil {
			return true
		}
		return hx_value_41.(bool)
	}(nb), hxrt.StringFromLiteral("tag")), counter.__hx_this.bump(func(hx_value_42 any) int {
		if hx_value_42 == nil {
			return 10
		}
		return hxrt.IntFromNullableAny(hx_value_42)
	}(ni), func(hx_value_43 any) float64 {
		if hx_value_43 == nil {
			return 1.5
		}
		return hx_value_43.(float64)
	}(nf), func(hx_value_44 any) bool {
		if hx_value_44 == nil {
			return true
		}
		return hx_value_44.(bool)
	}(nb), ns), counter.__hx_this.bump(func(hx_value_45 any) int {
		if hx_value_45 == nil {
			return 10
		}
		return hxrt.IntFromNullableAny(hx_value_45)
	}(zeroI), func(hx_value_46 any) float64 {
		if hx_value_46 == nil {
			return 1.5
		}
		return hx_value_46.(float64)
	}(zeroF), func(hx_value_47 any) bool {
		if hx_value_47 == nil {
			return true
		}
		return hx_value_47.(bool)
	}(falseB), empty)))
	row(hxrt.StringFromLiteral("ctor"), hxrt.NewArray(New_Counter(0.25, 2, true, hxrt.StringFromLiteral("ctr")).text, New_Counter(func(hx_value_48 any) float64 {
		if hx_value_48 == nil {
			return 0.25
		}
		return hx_value_48.(float64)
	}(nf), 2, true, hxrt.StringFromLiteral("ctr")).text, New_Counter(func(hx_value_49 any) float64 {
		if hx_value_49 == nil {
			return 0.25
		}
		return hx_value_49.(float64)
	}(nf), 2, true, hxrt.StringFromLiteral("ctr")).text, New_Counter(func(hx_value_50 any) float64 {
		if hx_value_50 == nil {
			return 0.25
		}
		return hx_value_50.(float64)
	}(nf), func(hx_value_51 any) int {
		if hx_value_51 == nil {
			return 2
		}
		return hxrt.IntFromNullableAny(hx_value_51)
	}(ni), func(hx_value_52 any) bool {
		if hx_value_52 == nil {
			return true
		}
		return hx_value_52.(bool)
	}(nb), ns).text, New_Counter(func(hx_value_53 any) float64 {
		if hx_value_53 == nil {
			return 0.25
		}
		return hx_value_53.(float64)
	}(zeroF), func(hx_value_54 any) int {
		if hx_value_54 == nil {
			return 2
		}
		return hxrt.IntFromNullableAny(hx_value_54)
	}(zeroI), func(hx_value_55 any) bool {
		if hx_value_55 == nil {
			return true
		}
		return hx_value_55.(bool)
	}(falseB), empty).text))
	row(hxrt.StringFromLiteral("super.ctor"), hxrt.NewArray(New_SubCounter(nil, nil, nil).text, New_SubCounter(nf, nb, ns).text, New_SubCounter(0.0, false, hxrt.StringFromLiteral("")).text))
	var child Pricer = New_ChildPricer()
	base := func(hx_value *ChildPricer) *BasePricer {
		if hx_value == nil {
			return nil
		}
		return hx_value.BasePricer
	}(New_ChildPricer())
	plain := New_BasePricer()
	row(hxrt.StringFromLiteral("interface"), hxrt.NewArray(child.price(nil, nil, nil), child.price(nf, ni, ns), child.price(zeroF, zeroI, empty)))
	row(hxrt.StringFromLiteral("override"), hxrt.NewArray(base.__hx_this.price(nf, ni, ns), plain.__hx_this.price(2.5, 4, hxrt.StringFromLiteral("base")), plain.__hx_this.price(nf, ni, ns)))
	var localScale *hxrt.Function[func(any, any, *string) *string] = hxrt.NewFunction(func(x any, n any, s *string) *string {
		if x == nil {
			x = 8.5
		}
		if n == nil {
			n = 2
		}
		if s == nil {
			s = hxrt.StringFromLiteral("loc")
		}
		return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(x.(float64), hxrt.StringFromLiteral("*")), n.(int)), hxrt.StringFromLiteral("*")), s)
	})
	row(hxrt.StringFromLiteral("local"), hxrt.NewArray(func() *string {
		var hx_callable_56 *hxrt.Function[func(any, any, *string) *string] = localScale
		var hx_argument_57 any = 8.5
		var hx_argument_58 any = 2
		var hx_argument_59 *string = hxrt.StringFromLiteral("loc")
		return hxrt.FunctionCall(hx_callable_56)(hx_argument_57, hx_argument_58, hx_argument_59)
	}(), func() *string {
		var hx_callable_60 *hxrt.Function[func(any, any, *string) *string] = localScale
		var hx_argument_61 any = nf
		var hx_argument_62 any = 2
		var hx_argument_63 *string = hxrt.StringFromLiteral("loc")
		return hxrt.FunctionCall(hx_callable_60)(hx_argument_61, hx_argument_62, hx_argument_63)
	}(), func() *string {
		var hx_callable_64 *hxrt.Function[func(any, any, *string) *string] = localScale
		var hx_argument_65 any = nf
		var hx_argument_66 any = ni
		var hx_argument_67 *string = hxrt.StringFromLiteral("loc")
		return hxrt.FunctionCall(hx_callable_64)(hx_argument_65, hx_argument_66, hx_argument_67)
	}(), func() *string {
		var hx_callable_68 *hxrt.Function[func(any, any, *string) *string] = localScale
		var hx_argument_69 any = nf
		var hx_argument_70 any = ni
		var hx_argument_71 *string = ns
		return hxrt.FunctionCall(hx_callable_68)(hx_argument_69, hx_argument_70, hx_argument_71)
	}(), func() *string {
		var hx_callable_72 *hxrt.Function[func(any, any, *string) *string] = localScale
		var hx_argument_73 any = zeroF
		var hx_argument_74 any = zeroI
		var hx_argument_75 *string = empty
		return hxrt.FunctionCall(hx_callable_72)(hx_argument_73, hx_argument_74, hx_argument_75)
	}()))
	var closure *hxrt.Function[func(*string, any, any) *string] = hxrt.NewFunction(func(s *string, b any, k any) *string {
		if s == nil {
			s = hxrt.StringFromLiteral("cl")
		}
		if b == nil {
			b = false
		}
		if k == nil {
			k = 11
		}
		return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(s, hxrt.StringFromLiteral(":")), hxrt.StdString(b.(bool))), hxrt.StringFromLiteral(":")), k.(int))
	})
	row(hxrt.StringFromLiteral("closure"), hxrt.NewArray(func() *string {
		var hx_callable_76 *hxrt.Function[func(*string, any, any) *string] = closure
		var hx_argument_77 *string = hxrt.StringFromLiteral("cl")
		var hx_argument_78 any = false
		var hx_argument_79 any = 11
		return hxrt.FunctionCall(hx_callable_76)(hx_argument_77, hx_argument_78, hx_argument_79)
	}(), func() *string {
		var hx_callable_80 *hxrt.Function[func(*string, any, any) *string] = closure
		var hx_argument_81 *string = nil
		var hx_argument_82 any = false
		var hx_argument_83 any = 11
		return hxrt.FunctionCall(hx_callable_80)(hx_argument_81, hx_argument_82, hx_argument_83)
	}(), func() *string {
		var hx_callable_84 *hxrt.Function[func(*string, any, any) *string] = closure
		var hx_argument_85 *string = ns
		var hx_argument_86 any = nb
		var hx_argument_87 any = 11
		return hxrt.FunctionCall(hx_callable_84)(hx_argument_85, hx_argument_86, hx_argument_87)
	}(), func() *string {
		var hx_callable_88 *hxrt.Function[func(*string, any, any) *string] = closure
		var hx_argument_89 *string = ns
		var hx_argument_90 any = nb
		var hx_argument_91 any = ni
		return hxrt.FunctionCall(hx_callable_88)(hx_argument_89, hx_argument_90, hx_argument_91)
	}(), func() *string {
		var hx_callable_92 *hxrt.Function[func(*string, any, any) *string] = closure
		var hx_argument_93 *string = empty
		var hx_argument_94 any = falseB
		var hx_argument_95 any = zeroI
		return hxrt.FunctionCall(hx_callable_92)(hx_argument_93, hx_argument_94, hx_argument_95)
	}()))
	hx_callable_96 := hxrt.NamedFunction(staticInt, "Main.staticInt")
	hx_normalized_callable_97 := hxrt.AdaptFunction(hx_callable_96, func(x any) int {
		if x == nil {
			x = 3
		}
		return hxrt.FunctionCall(hx_callable_96)(hxrt.IntFromNullableAny(x))
	})
	var intValue *hxrt.Function[func(any) int] = hx_normalized_callable_97
	hx_callable_98 := hxrt.NamedFunction(staticString, "Main.staticString")
	hx_normalized_callable_99 := hxrt.AdaptFunction(hx_callable_98, func(x *string) *string {
		if x == nil {
			x = hxrt.StringFromLiteral("dflt")
		}
		return hxrt.FunctionCall(hx_callable_98)(x)
	})
	stringValue := hx_normalized_callable_99
	hx_callable_107 := func() *hxrt.Function[func(any, any, any, *string) *string] {
		hx_receiver_100 := counter
		return hxrt.BoundFunction(func(by any, scale any, on any, tag *string) *string {
			if by == nil {
				by = 10
			}
			if scale == nil {
				scale = 1.5
			}
			if on == nil {
				on = true
			}
			if tag == nil {
				tag = hxrt.StringFromLiteral("tag")
			}
			return hx_receiver_100.__hx_this.bump(func(hx_value_101 any) int {
				if hx_value_101 == nil {
					var hx_zero_102 int
					return hx_zero_102
				}
				return hx_value_101.(int)
			}(by), func(hx_value_103 any) float64 {
				if hx_value_103 == nil {
					var hx_zero_104 float64
					return hx_zero_104
				}
				return hx_value_103.(float64)
			}(scale), func(hx_value_105 any) bool {
				if hx_value_105 == nil {
					var hx_zero_106 bool
					return hx_zero_106
				}
				return hx_value_105.(bool)
			}(on), tag)
		}, func() any {
			return hx_receiver_100.__hx_this
		}, "bump")
	}()
	hx_normalized_callable_112 := hxrt.AdaptFunction(hx_callable_107, func(by any, scale any, on any, tag *string) *string {
		if by == nil {
			by = 10
		}
		if scale == nil {
			scale = 1.5
		}
		if on == nil {
			on = true
		}
		if tag == nil {
			tag = hxrt.StringFromLiteral("tag")
		}
		return hxrt.FunctionCall(hx_callable_107)(hxrt.IntFromNullableAny(by), func(hx_value_108 any) float64 {
			if hx_value_108 == nil {
				var hx_zero_109 float64
				return hx_zero_109
			}
			return hx_value_108.(float64)
		}(scale), func(hx_value_110 any) bool {
			if hx_value_110 == nil {
				var hx_zero_111 bool
				return hx_zero_111
			}
			return hx_value_110.(bool)
		}(on), tag)
	})
	var bound *hxrt.Function[func(any, any, any, *string) *string] = hx_normalized_callable_112
	row(hxrt.StringFromLiteral("values"), hxrt.NewArray(func() int {
		var hx_callable_113 *hxrt.Function[func(any) int] = intValue
		var hx_argument_114 any = nil
		return hxrt.FunctionCall(hx_callable_113)(hx_argument_114)
	}(), func() int {
		var hx_callable_115 *hxrt.Function[func(any) int] = intValue
		var hx_argument_116 any = ni
		return hxrt.FunctionCall(hx_callable_115)(hx_argument_116)
	}(), func() int {
		var hx_callable_117 *hxrt.Function[func(any) int] = intValue
		var hx_argument_118 any = zeroI
		return hxrt.FunctionCall(hx_callable_117)(hx_argument_118)
	}(), func() *string {
		var hx_callable_119 *hxrt.Function[func(*string) *string] = stringValue
		var hx_argument_120 *string = ns
		return hxrt.FunctionCall(hx_callable_119)(hx_argument_120)
	}(), func() *string {
		var hx_callable_121 *hxrt.Function[func(any, any, any, *string) *string] = bound
		var hx_argument_122 any = nil
		var hx_argument_123 any = nil
		var hx_argument_124 any = nil
		var hx_argument_125 *string = nil
		return hxrt.FunctionCall(hx_callable_121)(hx_argument_122, hx_argument_123, hx_argument_124, hx_argument_125)
	}(), func() *string {
		var hx_callable_126 *hxrt.Function[func(any, any, any, *string) *string] = bound
		var hx_argument_127 any = ni
		var hx_argument_128 any = nf
		var hx_argument_129 any = nb
		var hx_argument_130 *string = ns
		return hxrt.FunctionCall(hx_callable_126)(hx_argument_127, hx_argument_128, hx_argument_129, hx_argument_130)
	}()))
	hx_callable_131 := hxrt.NamedFunction(staticInt, "Main.staticInt")
	hx_normalized_callable_132 := hxrt.AdaptFunction(hx_callable_131, func(x any) int {
		if x == nil {
			x = 3
		}
		return hxrt.FunctionCall(hx_callable_131)(hxrt.IntFromNullableAny(x))
	})
	var dynamicInt any = hx_normalized_callable_132
	hx_callable_133 := hxrt.NamedFunction(staticString, "Main.staticString")
	hx_normalized_callable_134 := hxrt.AdaptFunction(hx_callable_133, func(x *string) *string {
		if x == nil {
			x = hxrt.StringFromLiteral("dflt")
		}
		return hxrt.FunctionCall(hx_callable_133)(x)
	})
	var dynamicString any = hx_normalized_callable_134
	row(hxrt.StringFromLiteral("dynamic"), hxrt.NewArray(reflaxe__go___internal__CompilerReflect_callMethod(dynamicInt, []any{nil}), reflaxe__go___internal__CompilerReflect_callMethod(dynamicInt, []any{0}), reflaxe__go___internal__CompilerReflect_callMethod(dynamicString, []any{nil})))
	var three any = 3
	row(hxrt.StringFromLiteral("widen"), hxrt.NewArray(staticFloat(func(hx_value_135 any) float64 {
		if hx_value_135 == nil {
			return 1.5
		}
		return float64(hxrt.IntFromNullableAny(hx_value_135))
	}(ni)), staticFloat(func(hx_value_136 any) float64 {
		if hx_value_136 == nil {
			return 1.5
		}
		return float64(hxrt.IntFromNullableAny(hx_value_136))
	}(three)), optFloat(func(hx_value_137 any) float64 {
		if hx_value_137 == nil {
			return 4.5
		}
		return float64(hxrt.IntFromNullableAny(hx_value_137))
	}(ni)), optFloat(func(hx_value_138 any) float64 {
		if hx_value_138 == nil {
			return 4.5
		}
		return float64(hxrt.IntFromNullableAny(hx_value_138))
	}(three)), func() *string {
		var hx_callable_140 *hxrt.Function[func(any, any, *string) *string] = localScale
		var hx_argument_141 any = func(hx_value_139 any) any {
			if hx_value_139 == nil {
				return nil
			}
			return float64(hxrt.IntFromNullableAny(hx_value_139))
		}(ni)
		var hx_argument_142 any = 2
		var hx_argument_143 *string = hxrt.StringFromLiteral("loc")
		return hxrt.FunctionCall(hx_callable_140)(hx_argument_141, hx_argument_142, hx_argument_143)
	}(), child.price(func(hx_value_144 any) any {
		if hx_value_144 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_144))
	}(ni), three, nil)))
	var this1 float64
	this1 = 10
	var meters any = any(this1)
	row(hxrt.StringFromLiteral("abstract"), hxrt.NewArray(_Main__Meters_Impl__plus(func(hx_value_145 any) float64 {
		if hx_value_145 == nil {
			var hx_zero_146 float64
			return hx_zero_146
		}
		return hx_value_145.(float64)
	}(meters), 0.5, 2), _Main__Meters_Impl__plus(func(hx_value_147 any) float64 {
		if hx_value_147 == nil {
			var hx_zero_148 float64
			return hx_zero_148
		}
		return hx_value_147.(float64)
	}(meters), func(hx_value_149 any) float64 {
		if hx_value_149 == nil {
			return 0.5
		}
		return hx_value_149.(float64)
	}(nf), 2), _Main__Meters_Impl__plus(func(hx_value_150 any) float64 {
		if hx_value_150 == nil {
			var hx_zero_151 float64
			return hx_zero_151
		}
		return hx_value_150.(float64)
	}(meters), func(hx_value_152 any) float64 {
		if hx_value_152 == nil {
			return 0.5
		}
		return hx_value_152.(float64)
	}(nf), func(hx_value_153 any) int {
		if hx_value_153 == nil {
			return 2
		}
		return hxrt.IntFromNullableAny(hx_value_153)
	}(ni)), _Main__Meters_Impl__plus(func(hx_value_154 any) float64 {
		if hx_value_154 == nil {
			var hx_zero_155 float64
			return hx_zero_155
		}
		return hx_value_154.(float64)
	}(meters), func(hx_value_156 any) float64 {
		if hx_value_156 == nil {
			return 0.5
		}
		return hx_value_156.(float64)
	}(zeroF), func(hx_value_157 any) int {
		if hx_value_157 == nil {
			return 2
		}
		return hxrt.IntFromNullableAny(hx_value_157)
	}(zeroI))))
}

func mixed(a int, b float64, c bool, d *string) *string {
	if d == nil {
		d = hxrt.StringFromLiteral("d")
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatAny(a, hxrt.StringFromLiteral("/")), b), hxrt.StringFromLiteral("/")), hxrt.StdString(c)), hxrt.StringFromLiteral("/")), d)
}

func nothingFloat() any {
	return nil
}

func nothingInt() any {
	return nil
}

func nullInt(x int) any {
	return x
}

func optBool(x bool) bool {
	return func(hx_value_158 any) bool {
		if hx_value_158 == nil {
			var hx_zero_159 bool
			return hx_zero_159
		}
		return hx_value_158.(bool)
	}(x)
}

func optFloat(x float64) float64 {
	return func(hx_value_160 any) float64 {
		if hx_value_160 == nil {
			var hx_zero_161 float64
			return hx_zero_161
		}
		return hx_value_160.(float64)
	}(x)
}

func optInt(x int) int {
	return hxrt.IntFromNullableAny(x)
}

func optString(x *string) *string {
	if x == nil {
		x = hxrt.StringFromLiteral("opt")
	}
	return x
}

func row(label *string, values *hxrt.Array) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), hxrt.StringJoinAny(values.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v)
}

func staticBool(x bool) bool {
	return x
}

func staticFloat(x float64) float64 {
	return x
}

func staticInt(x int) int {
	return x
}

func staticLevel(x any) int {
	if x == nil {
		x = 2
	}
	return hxrt.IntFromNullableAny(x)
}

func staticShade(x *Shade) *string {
	if x == nil {
		x = Shade_Dark
	}
	return hxrt.StdString(x)
}

func staticString(x *string) *string {
	if x == nil {
		x = hxrt.StringFromLiteral("dflt")
	}
	return x
}

type Pricer interface {
	price(rate any, count any, label *string) *string
}

type I_SubCounter interface {
	bump(by int, scale float64, on bool, tag *string) *string
}

type SubCounter struct {
	*Counter
	__hx_this I_SubCounter
}

func (self *SubCounter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SubCounter(start any, on any, label *string) *SubCounter {
	return hx_init_SubCounter(hx_alloc_SubCounter(nil), start, on, label)
}

func hx_alloc_SubCounter(owner any) *SubCounter {
	self := &SubCounter{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SubCounter)
	}
	self.Counter = hx_alloc_Counter(self.__hx_this)
	return self
}

func hx_init_SubCounter(self *SubCounter, start any, on any, label *string) *SubCounter {
	hx_init_Counter(self.Counter, func(hx_value_162 any) float64 {
		if hx_value_162 == nil {
			return 0.25
		}
		return hx_value_162.(float64)
	}(start), 2, func(hx_value_163 any) bool {
		if hx_value_163 == nil {
			return true
		}
		return hx_value_163.(bool)
	}(on), label)
	return self
}

func (self *SubCounter) HxClassName() string {
	return "SubCounter"
}

func _Main__Meters_Impl___new(value float64) any {
	var this1 float64
	this1 = value
	return any(this1)
}

func _Main__Meters_Impl__plus(this1 float64, by float64, times int) float64 {
	return (this1 + (by * float64(times)))
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any, any, *string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(argument_0, argument_1, func(hx_value_1 any) *string {
			if hx_value_1 == nil {
				var hx_zero_2 *string
				return hx_zero_2
			}
			return hx_value_1.(*string)
		}(argument_2))
	case *hxrt.Function[func(*string, any, any) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(func(hx_value_3 any) *string {
			if hx_value_3 == nil {
				var hx_zero_4 *string
				return hx_zero_4
			}
			return hx_value_3.(*string)
		}(argument_0), argument_1, argument_2)
	case *hxrt.Function[func(any) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(*string) *string]:
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
	case *hxrt.Function[func(any, any, any, *string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		var argument_3 any = nil
		if 3 < len(arguments) {
			argument_3 = arguments[3]
		}
		return hxrt.FunctionCall(source)(argument_0, argument_1, argument_2, func(hx_value_7 any) *string {
			if hx_value_7 == nil {
				var hx_zero_8 *string
				return hx_zero_8
			}
			return hx_value_7.(*string)
		}(argument_3))
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_9 any) int {
			if hx_value_9 == nil {
				var hx_zero_10 int
				return hx_zero_10
			}
			return hx_value_9.(int)
		}(argument_0))
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}
