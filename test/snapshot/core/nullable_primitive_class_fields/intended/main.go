package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_Gauge interface {
	current() any
	plainAmount() float64
}

type Gauge struct {
	__hx_this I_Gauge
	reading   any
	amount    any
	count     any
	rate      any
	flag      any
	none      any
}

func (self *Gauge) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Gauge(start float64, whole int) *Gauge {
	return hx_init_Gauge(hx_alloc_Gauge(nil), start, whole)
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

func hx_init_Gauge(self *Gauge, start float64, whole int) *Gauge {
	self.none = nil
	self.flag = true
	self.rate = 2.0
	self.count = 3
	self.amount = start
	self.reading = float64(whole)
	return self
}

func (self *Gauge) current() any {
	return self.amount
}

func (self *Gauge) plainAmount() float64 {
	return func(hx_value_1 any) float64 {
		if hx_value_1 == nil {
			var hx_zero_2 float64
			return hx_zero_2
		}
		return hx_value_1.(float64)
	}(self.amount)
}

func (self *Gauge) HxClassName() string {
	return "Gauge"
}

var Gauge_enabled any = false

var Gauge_hits any = nil

var Gauge_limit any = 4.0

var Gauge_shared any

func describe(value any) *string {
	var hx_if_3 *string
	if value == nil {
		hx_if_3 = hxrt.StringFromLiteral("none")
	} else {
		hx_if_3 = hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("some("), (value.(float64)/float64(4))), hxrt.StringFromLiteral(")"))
	}
	return hx_if_3
}

func half(value float64) float64 {
	return (value / float64(2))
}

func main() {
	g := New_Gauge(1.5, 5)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("init="), g.amount), hxrt.StringFromLiteral(":")), (func(hx_value_8 any) float64 {
		if hx_value_8 == nil {
			var hx_zero_9 float64
			return hx_zero_9
		}
		return hx_value_8.(float64)
	}(func(hx_value_6 any) float64 {
		if hx_value_6 == nil {
			var hx_zero_7 float64
			return hx_zero_7
		}
		return hx_value_6.(float64)
	}(g.reading))/float64(2))), hxrt.StringFromLiteral(":")), g.count), hxrt.StringFromLiteral(":")), (func(hx_value_12 any) float64 {
		if hx_value_12 == nil {
			var hx_zero_13 float64
			return hx_zero_13
		}
		return hx_value_12.(float64)
	}(func(hx_value_10 any) float64 {
		if hx_value_10 == nil {
			var hx_zero_11 float64
			return hx_zero_11
		}
		return hx_value_10.(float64)
	}(g.rate))/float64(4))), hxrt.StringFromLiteral(":")), hxrt.StdString(g.flag)), hxrt.StringFromLiteral(":")), g.none), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.AnyEqualsNull(g.none))))
	hxrt.Println(v)
	g.amount = nil
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("null="), g.amount), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.AnyEqualsNull(g.amount))), hxrt.StringFromLiteral(":")), hxrt.StdString(!hxrt.AnyEqualsNull(g.amount))), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.AnyEqualsNull(g.amount))), hxrt.StringFromLiteral(":")), describe(g.amount)), hxrt.StringFromLiteral(":")), func() float64 {
		var tmp any = g.amount
		var hx_if_18 float64
		if tmp != nil {
			hx_if_18 = tmp.(float64)
		} else {
			hx_if_18 = 0.25
		}
		return hx_if_18
	}()))
	hxrt.Println(v_1)
	g.amount = 2.5
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("value="), g.amount), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.AnyEqualsNull(g.amount))), hxrt.StringFromLiteral(":")), hxrt.StdString((g.amount == 2.5))), hxrt.StringFromLiteral(":")), hxrt.StdString((func(hx_value_21 any) float64 {
		if hx_value_21 == nil {
			var hx_zero_22 float64
			return hx_zero_22
		}
		return hx_value_21.(float64)
	}(g.amount) > 1))), hxrt.StringFromLiteral(":")), describe(g.amount)), hxrt.StringFromLiteral(":")), func() float64 {
		var tmp_1 any = g.amount
		var hx_if_23 float64
		if tmp_1 != nil {
			hx_if_23 = tmp_1.(float64)
		} else {
			hx_if_23 = 0.25
		}
		return hx_if_23
	}()))
	hxrt.Println(v_2)
	n := 3
	g.amount = float64(n)
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("int="), (func(hx_value_26 any) float64 {
		if hx_value_26 == nil {
			var hx_zero_27 float64
			return hx_zero_27
		}
		return hx_value_26.(float64)
	}(func(hx_value_24 any) float64 {
		if hx_value_24 == nil {
			var hx_zero_25 float64
			return hx_zero_25
		}
		return hx_value_24.(float64)
	}(g.amount))/float64(2))), hxrt.StringFromLiteral(":")), half(func(hx_value_28 any) float64 {
		if hx_value_28 == nil {
			var hx_zero_29 float64
			return hx_zero_29
		}
		return hx_value_28.(float64)
	}(g.amount))))
	hxrt.Println(v_3)
	g.amount = 7.0
	g.amount = (func(hx_value_30 any) float64 {
		if hx_value_30 == nil {
			var hx_zero_31 float64
			return hx_zero_31
		}
		return hx_value_30.(float64)
	}(g.amount) + float64(1))
	g.amount = (func(hx_value_32 any) float64 {
		if hx_value_32 == nil {
			var hx_zero_33 float64
			return hx_zero_33
		}
		return hx_value_32.(float64)
	}(g.amount) * 0.5)
	g.amount = (func(hx_value_34 any) float64 {
		if hx_value_34 == nil {
			var hx_zero_35 float64
			return hx_zero_35
		}
		return hx_value_34.(float64)
	}(g.amount) - float64(n))
	g.count = int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(g.count)) + hxrt.Int32Wrap(2)))
	g.count = int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(g.count)) + hxrt.Int32Wrap(1)))
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("compound="), (func(hx_value_38 any) float64 {
		if hx_value_38 == nil {
			var hx_zero_39 float64
			return hx_zero_39
		}
		return hx_value_38.(float64)
	}(func(hx_value_36 any) float64 {
		if hx_value_36 == nil {
			var hx_zero_37 float64
			return hx_zero_37
		}
		return hx_value_36.(float64)
	}(g.amount))/float64(2))), hxrt.StringFromLiteral(":")), g.count), hxrt.StringFromLiteral(":")), twice(hxrt.IntFromNullableAny(g.count))))
	hxrt.Println(v_4)
	var local any = g.amount
	plain := func(hx_value_40 any) float64 {
		if hx_value_40 == nil {
			var hx_zero_41 float64
			return hx_zero_41
		}
		return hx_value_40.(float64)
	}(g.amount)
	g.rate = func(hx_value_42 any) any {
		if hx_value_42 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_42))
	}(g.count)
	g.none = g.amount
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("copy="), (func(hx_value_45 any) float64 {
		if hx_value_45 == nil {
			var hx_zero_46 float64
			return hx_zero_46
		}
		return hx_value_45.(float64)
	}(func(hx_value_43 any) float64 {
		if hx_value_43 == nil {
			var hx_zero_44 float64
			return hx_zero_44
		}
		return hx_value_43.(float64)
	}(local))/float64(2))), hxrt.StringFromLiteral(":")), (plain/float64(4))), hxrt.StringFromLiteral(":")), (func(hx_value_49 any) float64 {
		if hx_value_49 == nil {
			var hx_zero_50 float64
			return hx_zero_50
		}
		return hx_value_49.(float64)
	}(func(hx_value_47 any) float64 {
		if hx_value_47 == nil {
			var hx_zero_48 float64
			return hx_zero_48
		}
		return hx_value_47.(float64)
	}(g.rate))/float64(4))), hxrt.StringFromLiteral(":")), (func(hx_value_53 any) float64 {
		if hx_value_53 == nil {
			var hx_zero_54 float64
			return hx_zero_54
		}
		return hx_value_53.(float64)
	}(func(hx_value_51 any) float64 {
		if hx_value_51 == nil {
			var hx_zero_52 float64
			return hx_zero_52
		}
		return hx_value_51.(float64)
	}(g.none))/float64(2))), hxrt.StringFromLiteral(":")), (func(hx_value_57 any) float64 {
		if hx_value_57 == nil {
			var hx_zero_58 float64
			return hx_zero_58
		}
		return hx_value_57.(float64)
	}(func(hx_value_55 any) float64 {
		if hx_value_55 == nil {
			var hx_zero_56 float64
			return hx_zero_56
		}
		return hx_value_55.(float64)
	}(g.__hx_this.current())) / float64(4))))
	hxrt.Println(v_5)
	g.count = nil
	g.rate = func(hx_value_59 any) any {
		if hx_value_59 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_59))
	}(g.count)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("nullcopy="), g.count), hxrt.StringFromLiteral(":")), g.rate), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.AnyEqualsNull(g.rate))))
	hxrt.Println(v_6)
	g.flag = nil
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bool="), hxrt.StdString(g.flag)), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.AnyEqualsNull(g.flag))))
	hxrt.Println(v_7)
	g.flag = false
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bool="), hxrt.StdString(g.flag)), hxrt.StringFromLiteral(":")), hxrt.StdString((g.flag == false))), hxrt.StringFromLiteral(":")), hxrt.StdString(func() bool {
		var tmp_2 any = g.flag
		var hx_if_62 bool
		if tmp_2 != nil {
			hx_if_62 = tmp_2.(bool)
		} else {
			hx_if_62 = true
		}
		return hx_if_62
	}())))
	hxrt.Println(v_8)
	g.flag = true
	if func(hx_value_63 any) bool {
		if hx_value_63 == nil {
			var hx_zero_64 bool
			return hx_zero_64
		}
		return hx_value_63.(bool)
	}(g.flag) {
		var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bool="), hxrt.StdString(g.flag)), hxrt.StringFromLiteral(":")), (g.__hx_this.plainAmount()/float64(4))), hxrt.StringFromLiteral(":")), hxrt.StdString((g.amount == g.none))))
		hxrt.Println(v_9)
	}
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("static="), Gauge_shared), hxrt.StringFromLiteral(":")), (func(hx_value_69 any) float64 {
		if hx_value_69 == nil {
			var hx_zero_70 float64
			return hx_zero_70
		}
		return hx_value_69.(float64)
	}(func(hx_value_67 any) float64 {
		if hx_value_67 == nil {
			var hx_zero_68 float64
			return hx_zero_68
		}
		return hx_value_67.(float64)
	}(Gauge_limit))/float64(8))), hxrt.StringFromLiteral(":")), Gauge_hits), hxrt.StringFromLiteral(":")), hxrt.StdString(Gauge_enabled)))
	hxrt.Println(v_10)
	Gauge_shared = 1.0
	Gauge_shared = (func(hx_value_71 any) float64 {
		if hx_value_71 == nil {
			var hx_zero_72 float64
			return hx_zero_72
		}
		return hx_value_71.(float64)
	}(Gauge_shared) + 0.5)
	Gauge_hits = 4
	Gauge_hits = int((hxrt.Int32Wrap(hxrt.IntFromNullableAny(Gauge_hits)) + hxrt.Int32Wrap(1)))
	Gauge_limit = nil
	Gauge_enabled = nil
	var v_11 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("static="), (func(hx_value_75 any) float64 {
		if hx_value_75 == nil {
			var hx_zero_76 float64
			return hx_zero_76
		}
		return hx_value_75.(float64)
	}(func(hx_value_73 any) float64 {
		if hx_value_73 == nil {
			var hx_zero_74 float64
			return hx_zero_74
		}
		return hx_value_73.(float64)
	}(Gauge_shared))/float64(2))), hxrt.StringFromLiteral(":")), Gauge_limit), hxrt.StringFromLiteral(":")), Gauge_hits), hxrt.StringFromLiteral(":")), hxrt.StdString(Gauge_enabled)), hxrt.StringFromLiteral(":")), func() float64 {
		var tmp_3 any = Gauge_limit
		var hx_if_79 float64
		if tmp_3 != nil {
			hx_if_79 = tmp_3.(float64)
		} else {
			hx_if_79 = 0.75
		}
		return hx_if_79
	}()))
	hxrt.Println(v_11)
	var slot_count any
	var slot_amount any
	slot_amount = nil
	slot_count = 1
	var v_12 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("anon="), slot_amount), hxrt.StringFromLiteral(":")), slot_count))
	hxrt.Println(v_12)
	slot_amount = float64(n)
	slot_count = nil
	var v_13 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("anon="), (func(hx_value_84 any) float64 {
		if hx_value_84 == nil {
			var hx_zero_85 float64
			return hx_zero_85
		}
		return hx_value_84.(float64)
	}(func(hx_value_82 any) float64 {
		if hx_value_82 == nil {
			var hx_zero_83 float64
			return hx_zero_83
		}
		return hx_value_82.(float64)
	}(slot_amount))/float64(2))), hxrt.StringFromLiteral(":")), slot_count))
	hxrt.Println(v_13)
	view := func(hx_structural_source_86 *Gauge) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_86, func(name string) hxrt.AnonymousField {
			switch name {
			case "reading":
				return hxrt.AnonymousField{Value: hx_structural_source_86.reading, Found: true}
			case "amount":
				return hxrt.AnonymousField{Value: hx_structural_source_86.amount, Found: true}
			case "count":
				return hxrt.AnonymousField{Value: hx_structural_source_86.count, Found: true}
			case "rate":
				return hxrt.AnonymousField{Value: hx_structural_source_86.rate, Found: true}
			case "flag":
				return hxrt.AnonymousField{Value: hx_structural_source_86.flag, Found: true}
			case "none":
				return hxrt.AnonymousField{Value: hx_structural_source_86.none, Found: true}
			case "current":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_93 := hx_structural_source_86
					return hxrt.BoundFunction(func() any {
						return hx_receiver_93.__hx_this.current()
					}, func() any {
						return hx_receiver_93.__hx_this
					}, "current")
				}(), Found: true}
			case "plainAmount":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
					hx_receiver_94 := hx_structural_source_86
					return hxrt.BoundFunction(func() float64 {
						return hx_receiver_94.__hx_this.plainAmount()
					}, func() any {
						return hx_receiver_94.__hx_this
					}, "plainAmount")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "reading":
				hx_structural_source_86.reading = func(hx_value_87 any) any {
					if hx_value_87 == nil {
						return nil
					}
					return hx_value_87.(float64)
				}(incoming)
				return true
			case "amount":
				hx_structural_source_86.amount = func(hx_value_88 any) any {
					if hx_value_88 == nil {
						return nil
					}
					return hx_value_88.(float64)
				}(incoming)
				return true
			case "count":
				hx_structural_source_86.count = func(hx_value_89 any) any {
					if hx_value_89 == nil {
						return nil
					}
					return hx_value_89.(int)
				}(incoming)
				return true
			case "rate":
				hx_structural_source_86.rate = func(hx_value_90 any) any {
					if hx_value_90 == nil {
						return nil
					}
					return hx_value_90.(float64)
				}(incoming)
				return true
			case "flag":
				hx_structural_source_86.flag = func(hx_value_91 any) any {
					if hx_value_91 == nil {
						return nil
					}
					return hx_value_91.(bool)
				}(incoming)
				return true
			case "none":
				hx_structural_source_86.none = func(hx_value_92 any) any {
					if hx_value_92 == nil {
						return nil
					}
					return hx_value_92.(float64)
				}(incoming)
				return true
			}
			return false
		})
	}(g)
	func() any {
		hx_structural_target_95 := view
		hx_structural_value_97 := 4.5
		hxrt.AnonymousSet(hx_structural_target_95, "amount", hx_structural_value_97)
		return hx_structural_value_97
	}()
	func() any {
		hx_structural_target_98 := view
		hx_structural_value_100 := 9
		hxrt.AnonymousSet(hx_structural_target_98, "count", hx_structural_value_100)
		return hx_structural_value_100
	}()
	var v_14 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("view="), g.amount), hxrt.StringFromLiteral(":")), g.count), hxrt.StringFromLiteral(":")), (func(hx_value_106 any) float64 {
		if hx_value_106 == nil {
			var hx_zero_107 float64
			return hx_zero_107
		}
		return hx_value_106.(float64)
	}(func(hx_value_104 any) float64 {
		if hx_value_104 == nil {
			var hx_zero_105 float64
			return hx_zero_105
		}
		return hx_value_104.(float64)
	}(func(hx_value_103 any) any {
		if hx_value_103 == nil {
			return nil
		}
		return hx_value_103.(float64)
	}(hxrt.AnonymousGet(view, "amount")))) / float64(2))))
	hxrt.Println(v_14)
	g.amount = nil
	var v_15 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("view="), g.amount), hxrt.StringFromLiteral(":")), hxrt.StdString((func(hx_value_110 any) any {
		if hx_value_110 == nil {
			return nil
		}
		return hx_value_110.(float64)
	}(hxrt.AnonymousGet(view, "amount")) == nil))))
	hxrt.Println(v_15)
	var v_16 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("reflect="), hxrt.StdString(Reflect_field(g, hxrt.StringFromLiteral("amount")))), hxrt.StringFromLiteral(":")), hxrt.StdString(Reflect_field(g, hxrt.StringFromLiteral("count")))))
	hxrt.Println(v_16)
	Reflect_setField(g, hxrt.StringFromLiteral("amount"), 6.5)
	Reflect_setField(g, hxrt.StringFromLiteral("count"), nil)
	var v_17 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("reflect="), (func(hx_value_113 any) float64 {
		if hx_value_113 == nil {
			var hx_zero_114 float64
			return hx_zero_114
		}
		return hx_value_113.(float64)
	}(func(hx_value_111 any) float64 {
		if hx_value_111 == nil {
			var hx_zero_112 float64
			return hx_zero_112
		}
		return hx_value_111.(float64)
	}(g.amount))/float64(2))), hxrt.StringFromLiteral(":")), g.count), hxrt.StringFromLiteral(":")), hxrt.StdString(Reflect_field(g, hxrt.StringFromLiteral("count")))))
	hxrt.Println(v_17)
	Reflect_setField(g, hxrt.StringFromLiteral("amount"), 3)
	var v_18 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("reflect="), (func(hx_value_117 any) float64 {
		if hx_value_117 == nil {
			var hx_zero_118 float64
			return hx_zero_118
		}
		return hx_value_117.(float64)
	}(func(hx_value_115 any) float64 {
		if hx_value_115 == nil {
			var hx_zero_116 float64
			return hx_zero_116
		}
		return hx_value_115.(float64)
	}(g.amount))/float64(2))), hxrt.StringFromLiteral(":")), hxrt.StdString(Reflect_field(g, hxrt.StringFromLiteral("amount")))))
	hxrt.Println(v_18)
	Reflect_setField(g, hxrt.StringFromLiteral("flag"), false)
	Reflect_setField(g, hxrt.StringFromLiteral("rate"), nil)
	var v_19 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("reflect="), hxrt.StdString(g.flag)), hxrt.StringFromLiteral(":")), g.rate))
	hxrt.Println(v_19)
	Reflect_setField(g, hxrt.StringFromLiteral("amount"), nil)
	var v_20 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("reflect="), g.amount), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.AnyEqualsNull(g.amount))))
	hxrt.Println(v_20)
	sub := New_SubGauge()
	var v_21 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("sub="), sub.amount), hxrt.StringFromLiteral(":")), (func(hx_value_127 any) float64 {
		if hx_value_127 == nil {
			var hx_zero_128 float64
			return hx_zero_128
		}
		return hx_value_127.(float64)
	}(func(hx_value_125 any) float64 {
		if hx_value_125 == nil {
			var hx_zero_126 float64
			return hx_zero_126
		}
		return hx_value_125.(float64)
	}(sub.reading))/float64(2))), hxrt.StringFromLiteral(":")), sub.count), hxrt.StringFromLiteral(":")), sub.doubled()))
	hxrt.Println(v_21)
	options := New_Options(func(hx_value_129 any) any {
		if hx_value_129 == nil {
			return nil
		}
		return float64(hxrt.IntFromNullableAny(hx_value_129))
	}(3), hxrt.StringFromLiteral("w"))
	var v_22 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("struct="), (func(hx_value_132 any) float64 {
		if hx_value_132 == nil {
			var hx_zero_133 float64
			return hx_zero_133
		}
		return hx_value_132.(float64)
	}(func(hx_value_130 any) float64 {
		if hx_value_130 == nil {
			var hx_zero_131 float64
			return hx_zero_131
		}
		return hx_value_130.(float64)
	}(options.weight))/float64(2))), hxrt.StringFromLiteral(":")), options.label))
	hxrt.Println(v_22)
	options.weight = nil
	var v_23 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("struct="), options.weight))
	hxrt.Println(v_23)
}

func twice(value int) int {
	return int((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(2)))
}

type I_Options interface {
}

type Options struct {
	__hx_this I_Options
	weight    any
	label     *string
}

func (self *Options) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Options(weight any, label *string) *Options {
	return hx_init_Options(hx_alloc_Options(nil), weight, label)
}

func hx_alloc_Options(owner any) *Options {
	self := &Options{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Options)
	}
	return self
}

func hx_init_Options(self *Options, weight any, label *string) *Options {
	self.weight = weight
	self.label = label
	return self
}

func (self *Options) HxClassName() string {
	return "Options"
}

type I_SubGauge interface {
	current() any
	plainAmount() float64
	doubled() float64
}

type SubGauge struct {
	*Gauge
	__hx_this I_SubGauge
}

func (self *SubGauge) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SubGauge() *SubGauge {
	return hx_init_SubGauge(hx_alloc_SubGauge(nil))
}

func hx_alloc_SubGauge(owner any) *SubGauge {
	self := &SubGauge{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SubGauge)
	}
	self.Gauge = hx_alloc_Gauge(self.__hx_this)
	return self
}

func hx_init_SubGauge(self *SubGauge) *SubGauge {
	hx_init_Gauge(self.Gauge, 0.5, 7)
	self.count = nil
	return self
}

func (self *SubGauge) doubled() float64 {
	return (func(hx_value_138 any) float64 {
		if hx_value_138 == nil {
			var hx_zero_139 float64
			return hx_zero_139
		}
		return hx_value_138.(float64)
	}(func(hx_value_136 any) float64 {
		if hx_value_136 == nil {
			var hx_zero_137 float64
			return hx_zero_137
		}
		return hx_value_136.(float64)
	}(self.amount)) * float64(2))
}

func (self *SubGauge) HxClassName() string {
	return "SubGauge"
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *Gauge:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *SubGauge:
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
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Gauge:
		return hxrt__generated_method_field__Gauge(value, key)
	case *SubGauge:
		return hxrt__generated_method_field__SubGauge(value, key)
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Gauge(value *Gauge, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "current":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.current()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "current")
		}()
	case "plainAmount":
		return func() *hxrt.Function[func() float64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() float64 {
				return hx_receiver_1.__hx_this.plainAmount()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "plainAmount")
		}()
	}
	return nil
}

func hxrt__generated_method_field__SubGauge(value *SubGauge, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "doubled":
		return func() *hxrt.Function[func() float64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() float64 {
				return hx_receiver_1.__hx_this.doubled()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "doubled")
		}()
	}
	if value.Gauge == nil {
		return nil
	}
	return hxrt__generated_method_field__Gauge(value.Gauge, key)
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
	case *Gauge:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Options:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *SubGauge:
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
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Gauge:
		return hxrt__generated_field_lookup__Gauge(value, key)
	case *Options:
		return hxrt__generated_field_lookup__Options(value, key)
	case *SubGauge:
		return hxrt__generated_field_lookup__SubGauge(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Gauge:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *Options:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *SubGauge:
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
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Gauge:
		return hxrt__generated_field_has__Gauge(value, key)
	case *Options:
		return hxrt__generated_field_has__Options(value, key)
	case *SubGauge:
		return hxrt__generated_field_has__SubGauge(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Gauge:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *Options:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *SubGauge:
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
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Gauge:
		return hxrt__generated_field_set__Gauge(value, key, incoming)
	case *Options:
		return hxrt__generated_field_set__Options(value, key, incoming)
	case *SubGauge:
		return hxrt__generated_field_set__SubGauge(value, key, incoming)
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Gauge:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Options:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *SubGauge:
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
	default:
		return nil
	}
	switch receiver.(type) {
	case *Gauge:
		return hxrt.NewArray(hxrt.StringFromLiteral("reading"), hxrt.StringFromLiteral("amount"), hxrt.StringFromLiteral("count"), hxrt.StringFromLiteral("rate"), hxrt.StringFromLiteral("flag"), hxrt.StringFromLiteral("none"))
	case *Options:
		return hxrt.NewArray(hxrt.StringFromLiteral("weight"), hxrt.StringFromLiteral("label"))
	case *SubGauge:
		return hxrt.NewArray(hxrt.StringFromLiteral("reading"), hxrt.StringFromLiteral("amount"), hxrt.StringFromLiteral("count"), hxrt.StringFromLiteral("rate"), hxrt.StringFromLiteral("flag"), hxrt.StringFromLiteral("none"))
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__Gauge(value *Gauge, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "amount":
		return hxrt.ReadStoredField(value.amount)
	case "count":
		return hxrt.ReadStoredField(value.count)
	case "flag":
		return hxrt.ReadStoredField(value.flag)
	case "none":
		return hxrt.ReadStoredField(value.none)
	case "rate":
		return hxrt.ReadStoredField(value.rate)
	case "reading":
		return hxrt.ReadStoredField(value.reading)
	}
	return nil
}

func hxrt__generated_field_has__Gauge(value *Gauge, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "amount":
		return true
	case "count":
		return true
	case "flag":
		return true
	case "none":
		return true
	case "rate":
		return true
	case "reading":
		return true
	}
	return false
}

func hxrt__generated_field_set__Gauge(value *Gauge, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "amount":
		if incoming == nil {
			value.amount = nil
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.amount = typed
			return true
		case int:
			value.amount = float64(typed)
			return true
		default:
			return false
		}
	case "count":
		if incoming == nil {
			value.count = nil
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.count = typed
			return true
		default:
			return false
		}
	case "flag":
		if incoming == nil {
			value.flag = nil
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.flag = typed
			return true
		default:
			return false
		}
	case "none":
		if incoming == nil {
			value.none = nil
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.none = typed
			return true
		case int:
			value.none = float64(typed)
			return true
		default:
			return false
		}
	case "rate":
		if incoming == nil {
			value.rate = nil
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.rate = typed
			return true
		case int:
			value.rate = float64(typed)
			return true
		default:
			return false
		}
	case "reading":
		if incoming == nil {
			value.reading = nil
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.reading = typed
			return true
		case int:
			value.reading = float64(typed)
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__Options(value *Options, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "label":
		return hxrt.ReadStoredField(value.label)
	case "weight":
		return hxrt.ReadStoredField(value.weight)
	}
	return nil
}

func hxrt__generated_field_has__Options(value *Options, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "label":
		return true
	case "weight":
		return true
	}
	return false
}

func hxrt__generated_field_set__Options(value *Options, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "label":
		if incoming == nil {
			var zero *string
			value.label = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.label = typed
			return true
		default:
			return false
		}
	case "weight":
		if incoming == nil {
			value.weight = nil
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.weight = typed
			return true
		case int:
			value.weight = float64(typed)
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__SubGauge(value *SubGauge, key string) any {
	if value == nil {
		return nil
	}
	if value.Gauge == nil {
		return nil
	}
	return hxrt__generated_field_lookup__Gauge(value.Gauge, key)
}

func hxrt__generated_field_has__SubGauge(value *SubGauge, key string) bool {
	if value == nil {
		return false
	}
	if value.Gauge == nil {
		return false
	}
	return hxrt__generated_field_has__Gauge(value.Gauge, key)
}

func hxrt__generated_field_set__SubGauge(value *SubGauge, key string, incoming any) bool {
	if value == nil {
		return false
	}
	if value.Gauge == nil {
		return false
	}
	return hxrt__generated_field_set__Gauge(value.Gauge, key, incoming)
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
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() float64]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
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
