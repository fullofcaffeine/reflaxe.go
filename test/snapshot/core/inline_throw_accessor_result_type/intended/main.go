package main

import "snapshot/hxrt"

func boolContext(value *SnapshotInlineThrowAccessors) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bool="), hxrt.StdString(func() bool {
		var hx_if_2 bool
		if value.valid {
			hx_if_2 = value.boolValue
		} else {
			hx_if_2 = func() bool {
				hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("flag")))
				var hx_throw_zero_1 bool
				return hx_throw_zero_1
			}()
		}
		return hx_if_2
	}()))
}

func capture(label *string, action *hxrt.Function[func() *string]) *string {
	hx_try_return_3 := false
	var hx_try_value_4 *string
	hxrt.TryCatch(func() {
		hx_try_value_4 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(":miss:")), func() *string {
			var hx_callable_8 *hxrt.Function[func() *string] = action
			return hxrt.FunctionCall(hx_callable_8)()
		}())
		hx_try_return_3 = true
		return
	}, func(hx_caught_5 any) {
		hx_exception_6 := haxe__Exception_caught(hx_caught_5)
		hx_payload_7 := haxe__Exception_catchValue(hx_caught_5)
		_ = hx_exception_6
		var hx_tmp any = hx_payload_7
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
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("generic="), hxrt.StdString(func(hx_value_9 any) *string {
		if hx_value_9 == nil {
			var hx_zero_10 *string
			return hx_zero_10
		}
		return hx_value_9.(*string)
	}(genericValue(value))))
}

func genericValue(value *SnapshotInlineThrowAccessors) any {
	var hx_if_12 any
	if value.valid {
		hx_if_12 = value.genericValue
	} else {
		hx_if_12 = func() any {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("item")))
			var hx_throw_zero_11 any
			return hx_throw_zero_11
		}()
	}
	return hx_if_12
}

func intContext(value *SnapshotInlineThrowAccessors) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("int="), hxrt.StdString((func() int {
		if !value.valid {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("number")))
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
	var v_6 any = any(capture(hxrt.StringFromLiteral("string"), hxrt.NewFunction(func() *string {
		return stringContext(invalid)
	})))
	hxrt.Println(v_6)
	var v_7 any = any(capture(hxrt.StringFromLiteral("int"), hxrt.NewFunction(func() *string {
		return intContext(invalid)
	})))
	hxrt.Println(v_7)
	var v_8 any = any(capture(hxrt.StringFromLiteral("bool"), hxrt.NewFunction(func() *string {
		return boolContext(invalid)
	})))
	hxrt.Println(v_8)
	var v_9 any = any(capture(hxrt.StringFromLiteral("nullable"), hxrt.NewFunction(func() *string {
		return nullableContext(invalid)
	})))
	hxrt.Println(v_9)
	var v_10 any = any(capture(hxrt.StringFromLiteral("generic"), hxrt.NewFunction(func() *string {
		return genericContext(invalid)
	})))
	hxrt.Println(v_10)
}

func nestedFunctionContext() *string {
	action := hxrt.NewFunction(func() *string {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("nested")))
		var hx_throw_zero_13 *string
		return hx_throw_zero_13
	})
	return capture(hxrt.StringFromLiteral("nested"), action)
}

func nullableContext(value *SnapshotInlineThrowAccessors) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("nullable="), hxrt.StdString((func() *SnapshotInlineThrowBox {
		if !value.valid {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("maybe")))
		}
		return value.nullableValue
	}() != nil)))
}

func stringContext(value *SnapshotInlineThrowAccessors) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("string="), hxrt.StdString(hxrt.StringEqualStringPtr(func() *string {
		if !value.valid {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("text")))
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

func (self *SnapshotInlineThrowAccessors) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotInlineThrowAccessors(valid bool, textValue *string, intValue int, boolValue bool, nullableValue *SnapshotInlineThrowBox, genericValue any) *SnapshotInlineThrowAccessors {
	return hx_init_SnapshotInlineThrowAccessors(hx_alloc_SnapshotInlineThrowAccessors(nil), valid, textValue, intValue, boolValue, nullableValue, genericValue)
}

func hx_alloc_SnapshotInlineThrowAccessors(owner any) *SnapshotInlineThrowAccessors {
	self := &SnapshotInlineThrowAccessors{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SnapshotInlineThrowAccessors)
	}
	return self
}

func hx_init_SnapshotInlineThrowAccessors(self *SnapshotInlineThrowAccessors, valid bool, textValue *string, intValue int, boolValue bool, nullableValue *SnapshotInlineThrowBox, genericValue any) *SnapshotInlineThrowAccessors {
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
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("text")))
	}
	return self.textValue
}

func (self *SnapshotInlineThrowAccessors) get_number() int {
	if !self.valid {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("number")))
	}
	return self.intValue
}

func (self *SnapshotInlineThrowAccessors) get_flag() bool {
	var hx_if_15 bool
	if self.valid {
		hx_if_15 = self.boolValue
	} else {
		hx_if_15 = func() bool {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("flag")))
			var hx_throw_zero_14 bool
			return hx_throw_zero_14
		}()
	}
	return hx_if_15
}

func (self *SnapshotInlineThrowAccessors) get_maybe() *SnapshotInlineThrowBox {
	if !self.valid {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("maybe")))
	}
	return self.nullableValue
}

func (self *SnapshotInlineThrowAccessors) get_item() any {
	var hx_if_17 any
	if self.valid {
		hx_if_17 = self.genericValue
	} else {
		hx_if_17 = func() any {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("item")))
			var hx_throw_zero_16 any
			return hx_throw_zero_16
		}()
	}
	return hx_if_17
}

type I_SnapshotInlineThrowBox interface {
}

type SnapshotInlineThrowBox struct {
	__hx_this I_SnapshotInlineThrowBox
	label     *string
}

func (self *SnapshotInlineThrowBox) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SnapshotInlineThrowBox(label *string) *SnapshotInlineThrowBox {
	return hx_init_SnapshotInlineThrowBox(hx_alloc_SnapshotInlineThrowBox(nil), label)
}

func hx_alloc_SnapshotInlineThrowBox(owner any) *SnapshotInlineThrowBox {
	self := &SnapshotInlineThrowBox{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SnapshotInlineThrowBox)
	}
	return self
}

func hx_init_SnapshotInlineThrowBox(self *SnapshotInlineThrowBox, label *string) *SnapshotInlineThrowBox {
	self.label = label
	return self
}

func hx_recover_class_0(value any) *haxe__ValueException {
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

func hx_recover_class_1(value any) *haxe__Exception {
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
