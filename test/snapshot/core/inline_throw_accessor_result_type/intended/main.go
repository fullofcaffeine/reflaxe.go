package main

import "snapshot/hxrt"

func boolContext(value *SnapshotInlineThrowAccessors) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bool="), hxrt.StdString(func() bool {
		var hx_if_2 bool
		if value.valid {
			hx_if_2 = value.boolValue
		} else {
			hx_if_2 = func() bool {
				hxrt.Throw(hxrt.StringFromLiteral("flag"))
				var hx_throw_zero_1 bool
				return hx_throw_zero_1
			}()
		}
		return hx_if_2
	}()))
}

func capture(label *string, action func() *string) *string {
	hx_try_return_3 := false
	var hx_try_value_4 *string
	hxrt.TryCatch(func() {
		hx_try_value_4 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(":miss:")), action())
		hx_try_return_3 = true
		return
	}, func(hx_caught_5 any) {
		hx_tmp := hx_caught_5
		_ = hx_tmp
		hx_try_value_4 = hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(":throw"))
		hx_try_return_3 = true
		return
	})
	if hx_try_return_3 {
		return hx_try_value_4
	}
	return hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(":unreachable"))
}

func genericContext(value *SnapshotInlineThrowAccessors) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic="), hxrt.StdString(genericValue(value)))
}

func genericValue(value *SnapshotInlineThrowAccessors) any {
	var hx_if_8 any
	if value.valid {
		hx_if_8 = value.genericValue
	} else {
		hx_if_8 = func() any {
			hxrt.Throw(hxrt.StringFromLiteral("item"))
			var hx_throw_zero_7 any
			return hx_throw_zero_7
		}()
	}
	return hx_if_8
}

func intContext(value *SnapshotInlineThrowAccessors) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("int="), hxrt.StdString((func() int {
		if !value.valid {
			hxrt.Throw(hxrt.StringFromLiteral("number"))
		}
		return value.intValue
	}() == 7)))
}

func main() {
	valid := New_SnapshotInlineThrowAccessors(true, hxrt.StringFromLiteral("text"), 7, true, New_SnapshotInlineThrowBox(hxrt.StringFromLiteral("box")), hxrt.StringFromLiteral("generic"))
	invalid := New_SnapshotInlineThrowAccessors(false, hxrt.StringFromLiteral("text"), 7, true, New_SnapshotInlineThrowBox(hxrt.StringFromLiteral("box")), hxrt.StringFromLiteral("generic"))
	var v any = any(stringContext(valid))
	hxrt.Println(v)
	var v_1 any = any(intContext(valid))
	hxrt.Println(v_1)
	var v_2 any = any(boolContext(valid))
	hxrt.Println(v_2)
	var v_3 any = any(nullableContext(valid))
	hxrt.Println(v_3)
	var v_4 any = any(genericContext(valid))
	hxrt.Println(v_4)
	var v_5 any = any(nestedFunctionContext())
	hxrt.Println(v_5)
	var v_6 any = any(capture(hxrt.StringFromLiteral("string"), func() *string {
		return stringContext(invalid)
	}))
	hxrt.Println(v_6)
	var v_7 any = any(capture(hxrt.StringFromLiteral("int"), func() *string {
		return intContext(invalid)
	}))
	hxrt.Println(v_7)
	var v_8 any = any(capture(hxrt.StringFromLiteral("bool"), func() *string {
		return boolContext(invalid)
	}))
	hxrt.Println(v_8)
	var v_9 any = any(capture(hxrt.StringFromLiteral("nullable"), func() *string {
		return nullableContext(invalid)
	}))
	hxrt.Println(v_9)
	var v_10 any = any(capture(hxrt.StringFromLiteral("generic"), func() *string {
		return genericContext(invalid)
	}))
	hxrt.Println(v_10)
}

func nestedFunctionContext() *string {
	action := func() *string {
		hxrt.Throw(hxrt.StringFromLiteral("nested"))
		var hx_throw_zero_9 *string
		return hx_throw_zero_9
	}
	return capture(hxrt.StringFromLiteral("nested"), action)
}

func nullableContext(value *SnapshotInlineThrowAccessors) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("nullable="), hxrt.StdString((func() *SnapshotInlineThrowBox {
		if !value.valid {
			hxrt.Throw(hxrt.StringFromLiteral("maybe"))
		}
		return value.nullableValue
	}() != nil)))
}

func stringContext(value *SnapshotInlineThrowAccessors) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("string="), hxrt.StdString(hxrt.StringEqualStringPtr(func() *string {
		if !value.valid {
			hxrt.Throw(hxrt.StringFromLiteral("text"))
		}
		return value.textValue
	}(), hxrt.StringFromLiteral("text"))))
}

type I_SnapshotInlineThrowAccessors interface {
	get_text() *string
	get_number() int
	get_flag() bool
	get_maybe() *SnapshotInlineThrowBox
	get_item() any
}

type SnapshotInlineThrowAccessors struct {
	__hx_this     I_SnapshotInlineThrowAccessors
	valid         bool
	textValue     *string
	intValue      int
	boolValue     bool
	nullableValue *SnapshotInlineThrowBox
	genericValue  any
	text          *string
	number        int
	flag          bool
	maybe         *SnapshotInlineThrowBox
	item          any
}

func New_SnapshotInlineThrowAccessors(valid bool, textValue *string, intValue int, boolValue bool, nullableValue *SnapshotInlineThrowBox, genericValue any) *SnapshotInlineThrowAccessors {
	self := &SnapshotInlineThrowAccessors{}
	self.__hx_this = self
	self.valid = valid
	self.textValue = textValue
	self.intValue = intValue
	self.boolValue = boolValue
	self.nullableValue = nullableValue
	self.genericValue = genericValue
	return self
}

func (self *SnapshotInlineThrowAccessors) get_text() *string {
	if !self.valid {
		hxrt.Throw(hxrt.StringFromLiteral("text"))
	}
	return self.textValue
}

func (self *SnapshotInlineThrowAccessors) get_number() int {
	if !self.valid {
		hxrt.Throw(hxrt.StringFromLiteral("number"))
	}
	return self.intValue
}

func (self *SnapshotInlineThrowAccessors) get_flag() bool {
	var hx_if_11 bool
	if self.valid {
		hx_if_11 = self.boolValue
	} else {
		hx_if_11 = func() bool {
			hxrt.Throw(hxrt.StringFromLiteral("flag"))
			var hx_throw_zero_10 bool
			return hx_throw_zero_10
		}()
	}
	return hx_if_11
}

func (self *SnapshotInlineThrowAccessors) get_maybe() *SnapshotInlineThrowBox {
	if !self.valid {
		hxrt.Throw(hxrt.StringFromLiteral("maybe"))
	}
	return self.nullableValue
}

func (self *SnapshotInlineThrowAccessors) get_item() any {
	var hx_if_13 any
	if self.valid {
		hx_if_13 = self.genericValue
	} else {
		hx_if_13 = func() any {
			hxrt.Throw(hxrt.StringFromLiteral("item"))
			var hx_throw_zero_12 any
			return hx_throw_zero_12
		}()
	}
	return hx_if_13
}

type I_SnapshotInlineThrowBox interface {
}

type SnapshotInlineThrowBox struct {
	__hx_this I_SnapshotInlineThrowBox
	label     *string
}

func New_SnapshotInlineThrowBox(label *string) *SnapshotInlineThrowBox {
	self := &SnapshotInlineThrowBox{}
	self.__hx_this = self
	self.label = label
	return self
}
