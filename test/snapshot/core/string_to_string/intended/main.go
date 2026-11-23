package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_Holder interface {
}

type Holder struct {
	__hx_this I_Holder
	text      *string
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder(text *string) *Holder {
	return hx_init_Holder(hx_alloc_Holder(nil), text)
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

func hx_init_Holder(self *Holder, text *string) *Holder {
	self.text = text
	return self
}

func (self *Holder) HxClassName() string {
	return "Holder"
}

var calls int = 0

func main() {
	s := hxrt.StringFromLiteral("abc")
	var v any = any(func(hx_receiver_1 *string) *string {
		if hx_receiver_1 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_1
	}(s))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringEqualStringPtr(func(hx_receiver_2 *string) *string {
		if hx_receiver_2 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_2
	}(s), s))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringLengthStringPtr(func(hx_receiver_3 *string) *string {
		if hx_receiver_3 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_3
	}(s)), hxrt.StringFromLiteral(" ")), hxrt.StringToUpperCaseStringPtr(func(hx_receiver_4 *string) *string {
		if hx_receiver_4 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_4
	}(s))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(func(hx_receiver_5 *string) *string {
		if hx_receiver_5 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_5
	}(next()), hxrt.StringFromLiteral(" ")), calls))
	hxrt.Println(v_3)
	func(hx_receiver_6 *string) *string {
		if hx_receiver_6 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_6
	}(next())
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("statement "), calls))
	hxrt.Println(v_4)
	var v_5 any = any(func(hx_receiver_8 *string) *string {
		if hx_receiver_8 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_8
	}(func() *string {
		var hx_if_7 *string
		if calls > 1 {
			hx_if_7 = s
		} else {
			hx_if_7 = hxrt.StringFromLiteral("other")
		}
		return hx_if_7
	}()))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), func(hx_receiver_9 *string) *string {
		if hx_receiver_9 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_9
	}(s)), hxrt.StringFromLiteral("!")))
	hxrt.Println(v_6)
	var words_1 *string
	var words_0 *string
	_ = words_0
	words_0 = hxrt.StringFromLiteral("p")
	words_1 = hxrt.StringFromLiteral("q")
	var v_7 any = any(func(hx_receiver_10 *string) *string {
		if hx_receiver_10 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_10
	}(words_1))
	hxrt.Println(v_7)
	holder := New_Holder(hxrt.StringFromLiteral("field"))
	var v_8 any = any(func(hx_receiver_11 *string) *string {
		if hx_receiver_11 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_11
	}(holder.text))
	hxrt.Println(v_8)
	var anon_text *string
	anon_text = hxrt.StringFromLiteral("anon")
	var v_9 any = any(func(hx_receiver_12 *string) *string {
		if hx_receiver_12 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_12
	}(anon_text))
	hxrt.Println(v_9)
	name := hxrt.StringFromLiteral("bob")
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func(hx_receiver_13 *string) *string {
		if hx_receiver_13 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_13
	}(name), hxrt.StringFromLiteral(" ")), hxrt.StringToUpperCaseStringPtr(name)))
	hxrt.Println(v_10)
	var this1 *string
	this1 = hxrt.StringFromLiteral("x")
	label := this1
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(_Main__Label_Impl__toString(label), hxrt.StringFromLiteral(" ")), func() *string {
		var hx_if_14 *string
		if label == nil {
			hx_if_14 = hxrt.StringFromLiteral("null")
		} else {
			hx_if_14 = _Main__Label_Impl__toString(label)
		}
		return hx_if_14
	}()))
	hxrt.Println(v_11)
	color := hxrt.StringFromLiteral("red")
	var v_12 any = any(func(hx_receiver_15 *string) *string {
		if hx_receiver_15 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_15
	}(color))
	hxrt.Println(v_12)
	var dynamicText any = hxrt.StringFromLiteral("dyn")
	var v_13 any = reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(dynamicText, hxrt.StringFromLiteral("toString")), []any{})
	hxrt.Println(v_13)
	typedFromDynamic := hxrt.DynamicToString(dynamicText)
	var v_14 any = any(func(hx_receiver_16 *string) *string {
		if hx_receiver_16 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_16
	}(typedFromDynamic))
	hxrt.Println(v_14)
	var v_15 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(s), hxrt.StringFromLiteral(" ")), hxrt.StdString(func(hx_receiver_17 *string) *string {
		if hx_receiver_17 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_17
	}(s))))
	hxrt.Println(v_15)
	empty := hxrt.StringFromLiteral("")
	var v_16 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), func(hx_receiver_18 *string) *string {
		if hx_receiver_18 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_18
	}(empty)), hxrt.StringFromLiteral("]")))
	hxrt.Println(v_16)
	unicode := hxrt.StringFromLiteral("héllo")
	var v_17 any = any(hxrt.StringLengthStringPtr(func(hx_receiver_19 *string) *string {
		if hx_receiver_19 == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		}
		return hx_receiver_19
	}(unicode)))
	hxrt.Println(v_17)
	missing := nothing()
	var v_18 any = any(hxrt.StdString(missing))
	hxrt.Println(v_18)
	hxrt.TryCatch(func() {
		var v_19 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null toString "), func(hx_receiver_25 *string) *string {
			if hx_receiver_25 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			}
			return hx_receiver_25
		}(missing)))
		hxrt.Println(v_19)
	}, func(hx_caught_20 any) {
		hx_exception_21 := haxe__Exception_caught(hx_caught_20)
		hx_payload_22 := haxe__Exception_catchValue(hx_caught_20)
		_ = hx_exception_21
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_22) {
			var e *string = func(hx_value_23 any) *string {
				if hx_value_23 == nil {
					var hx_zero_24 *string
					return hx_zero_24
				}
				return hx_value_23.(*string)
			}(hx_payload_22)
			hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null toString threw "), e)))
		} else {
			hxrt.Throw(hx_caught_20)
		}
	})
	hxrt.TryCatch(func() {
		func(hx_receiver_31 *string) *string {
			if hx_receiver_31 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			}
			return hx_receiver_31
		}(missing)
		hxrt.Println(any(hxrt.StringFromLiteral("null statement ran")))
	}, func(hx_caught_26 any) {
		hx_exception_27 := haxe__Exception_caught(hx_caught_26)
		hx_payload_28 := haxe__Exception_catchValue(hx_caught_26)
		_ = hx_exception_27
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_28) {
			var e_1 *string = func(hx_value_29 any) *string {
				if hx_value_29 == nil {
					var hx_zero_30 *string
					return hx_zero_30
				}
				return hx_value_29.(*string)
			}(hx_payload_28)
			hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null statement threw "), e_1)))
		} else {
			hxrt.Throw(hx_caught_26)
		}
	})
	optional := nothing()
	hxrt.TryCatch(func() {
		value := func(hx_receiver_37 *string) *string {
			if hx_receiver_37 == nil {
				hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			}
			return hx_receiver_37
		}(optional)
		hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null optional "), value)))
	}, func(hx_caught_32 any) {
		hx_exception_33 := haxe__Exception_caught(hx_caught_32)
		hx_payload_34 := haxe__Exception_catchValue(hx_caught_32)
		_ = hx_exception_33
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_34) {
			var e_2 *string = func(hx_value_35 any) *string {
				if hx_value_35 == nil {
					var hx_zero_36 *string
					return hx_zero_36
				}
				return hx_value_35.(*string)
			}(hx_payload_34)
			hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null optional threw "), e_2)))
		} else {
			hxrt.Throw(hx_caught_32)
		}
	})
}

func next() *string {
	calls = int(int32((calls + 1)))
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("n"), calls)
}

func nothing() *string {
	return nil
}

func _Main__Label_Impl___new(text *string) *string {
	var this1 *string
	this1 = text
	return this1
}

func _Main__Label_Impl__toString(this1 *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("label:"), this1)
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
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
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
	case *Holder:
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
	case *Holder:
		return hxrt__generated_field_lookup__Holder(value, key)
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
	case *Holder:
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
	case *Holder:
		return hxrt__generated_field_has__Holder(value, key)
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
	case *Holder:
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
	case *Holder:
		return hxrt__generated_field_set__Holder(value, key, incoming)
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
	case *Holder:
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
	case *Holder:
		return hxrt.NewArray(hxrt.StringFromLiteral("text"))
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

func hxrt__generated_field_lookup__Holder(value *Holder, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "text":
		return hxrt.ReadStoredField(value.text)
	}
	return nil
}

func hxrt__generated_field_has__Holder(value *Holder, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "text":
		return true
	}
	return false
}

func hxrt__generated_field_set__Holder(value *Holder, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "text":
		if incoming == nil {
			var zero *string
			value.text = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.text = typed
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
	case *hxrt.Function[func() any]:
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
