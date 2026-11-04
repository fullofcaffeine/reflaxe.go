package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_CallbackReader interface {
}

type CallbackReader struct {
	__hx_this I_CallbackReader
	read      *hxrt.Function[func(*hxrt.Function[func(any) any]) int]
}

func (self *CallbackReader) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_CallbackReader(read *hxrt.Function[func(*hxrt.Function[func(any) any]) int]) *CallbackReader {
	return hx_init_CallbackReader(hx_alloc_CallbackReader(nil), read)
}

func hx_alloc_CallbackReader(owner any) *CallbackReader {
	self := &CallbackReader{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_CallbackReader)
	}
	return self
}

func hx_init_CallbackReader(self *CallbackReader, read *hxrt.Function[func(*hxrt.Function[func(any) any]) int]) *CallbackReader {
	self.read = read
	return self
}

type I_Holder interface {
}

type Holder struct {
	__hx_this I_Holder
	callback  *hxrt.Function[func(any) any]
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder(callback *hxrt.Function[func(any) any]) *Holder {
	return hx_init_Holder(hx_alloc_Holder(nil), callback)
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

func hx_init_Holder(self *Holder, callback *hxrt.Function[func(any) any]) *Holder {
	self.callback = callback
	return self
}

func main() {
	hx_obj_1 := map[string]any{}
	hx_obj_1["next"] = hxrt.NewFunction(func() int {
		return 3
	})
	source := hxrt.AnonymousFromMap(hx_obj_1)
	replace(source, hx_recover_callable_0(hxrt.NewFunction(func() int {
		return 4
	})))
	var v any = any(func() int {
		var hx_callable_2 *hxrt.Function[func() int] = hx_recover_callable_1(hxrt.AnonymousGet(source, "next"))
		return hxrt.FunctionCall(hx_callable_2)()
	}())
	hxrt.Println(v)
	holder := New_Holder(hx_recover_callable_2(hxrt.NewFunction(func(value int) int {
		return int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(1)))
	})))
	Reflect_setField(holder, hxrt.StringFromLiteral("callback"), hxrt.NewFunction(func(value int) int {
		return int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(2)))
	}))
	var v_1 any = any(func() int {
		var hx_callable_3 *hxrt.Function[func(int) int] = hx_recover_callable_3(holder.callback)
		var hx_argument_4 int = 3
		return hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}())
	hxrt.Println(v_1)
	reader := New_CallbackReader(hx_recover_callable_4(hxrt.NewFunction(func(callback *hxrt.Function[func(int) int]) int {
		return func() int {
			var hx_callable_5 *hxrt.Function[func(int) int] = callback
			var hx_argument_6 int = 3
			return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
		}()
	})))
	var v_2 any = any(func() int {
		var hx_callable_7 *hxrt.Function[func(*hxrt.Function[func(int) int]) int] = hx_recover_callable_5(reader.read)
		var hx_argument_8 *hxrt.Function[func(int) int] = hxrt.NewFunction(func(value int) int {
			return int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(1)))
		})
		return hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
	}())
	hxrt.Println(v_2)
	var empty *hxrt.Function[func(int) int] = nil
	Reflect_setField(holder, hxrt.StringFromLiteral("callback"), empty)
	var v_3 any = any((hx_recover_callable_3(holder.callback) == nil))
	hxrt.Println(v_3)
}

func replace(value hxrt.Anonymous, next *hxrt.Function[func() any]) {
	Reflect_setField(value, hxrt.StringFromLiteral("next"), next)
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
	case *CallbackReader:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Holder:
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
	case *CallbackReader:
		return hxrt__generated_field_lookup__CallbackReader(value, key)
	case *Holder:
		return hxrt__generated_field_lookup__Holder(value, key)
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
	case *CallbackReader:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *Holder:
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
	case *CallbackReader:
		return hxrt__generated_field_has__CallbackReader(value, key)
	case *Holder:
		return hxrt__generated_field_has__Holder(value, key)
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
	case *CallbackReader:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *Holder:
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
	case *CallbackReader:
		return hxrt__generated_field_set__CallbackReader(value, key, incoming)
	case *Holder:
		return hxrt__generated_field_set__Holder(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *CallbackReader:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Holder:
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
	case *CallbackReader:
		return hxrt.NewArray(hxrt.StringFromLiteral("read"))
	case *Holder:
		return hxrt.NewArray(hxrt.StringFromLiteral("callback"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__CallbackReader(value *CallbackReader, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "read":
		return hxrt.ReadStoredField(value.read)
	}
	return nil
}

func hxrt__generated_field_has__CallbackReader(value *CallbackReader, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "read":
		return true
	}
	return false
}

func hxrt__generated_field_set__CallbackReader(value *CallbackReader, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "read":
		if (hxrt.FunctionOrigin(incoming) == nil) || hxrt.ReflectIsFunction(incoming) {
			value.read = hx_recover_callable_4(incoming)
			return true
		}
		return false
	}
	return false
}

func hxrt__generated_field_lookup__Holder(value *Holder, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "callback":
		return hxrt.ReadStoredField(value.callback)
	}
	return nil
}

func hxrt__generated_field_has__Holder(value *Holder, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "callback":
		return true
	}
	return false
}

func hxrt__generated_field_set__Holder(value *Holder, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "callback":
		if (hxrt.FunctionOrigin(incoming) == nil) || hxrt.ReflectIsFunction(incoming) {
			value.callback = hx_recover_callable_2(incoming)
			return true
		}
		return false
	}
	return false
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
	return nil
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	return false
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	switch source := hxrt.FunctionOrigin(value).(type) {
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_17 any) int {
			if hx_value_17 == nil {
				var hx_zero_18 int
				return hx_zero_18
			}
			return hx_value_17.(int)
		}(any(argument_0)))
	case *hxrt.Function[func(*hxrt.Function[func(int) int]) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_callable_3(argument_0))
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func() any] {
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
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() int] {
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
	case *hxrt.Function[func() int]:
		return source
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_19 any) int {
				if hx_value_19 == nil {
					var hx_zero_20 int
					return hx_zero_20
				}
				return hx_value_19.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func(any) any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any) any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_21 any) int {
				if hx_value_21 == nil {
					var hx_zero_22 int
					return hx_zero_22
				}
				return hx_value_21.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.Function[func(int) int]) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(hx_recover_callable_3(arg_0))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_3(value any) *hxrt.Function[func(int) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(int) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int) int]:
		return source
	case *hxrt.Function[func(*hxrt.Function[func(int) int]) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return hxrt.FunctionCall(source)(hx_recover_callable_3(arg_0))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_23 any) int {
				if hx_value_23 == nil {
					var hx_zero_24 int
					return hx_zero_24
				}
				return hx_value_23.(int)
			}(any(hxrt.FunctionCall(source)(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_4(value any) *hxrt.Function[func(*hxrt.Function[func(any) any]) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(any) any]) int {
			return hxrt.FunctionCall(source)(func(hx_value_25 any) int {
				if hx_value_25 == nil {
					var hx_zero_26 int
					return hx_zero_26
				}
				return hx_value_25.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.Function[func(int) int]) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(any) any]) int {
			return hxrt.FunctionCall(source)(hx_recover_callable_3(arg_0))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(any) any]) int {
			return func(hx_value_27 any) int {
				if hx_value_27 == nil {
					var hx_zero_28 int
					return hx_zero_28
				}
				return hx_value_27.(int)
			}(any(hxrt.FunctionCall(source)(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_5(value any) *hxrt.Function[func(*hxrt.Function[func(int) int]) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(int) int]) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(int) int]) int {
			return hxrt.FunctionCall(source)(func(hx_value_29 any) int {
				if hx_value_29 == nil {
					var hx_zero_30 int
					return hx_zero_30
				}
				return hx_value_29.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.Function[func(int) int]) int]:
		return source
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(int) int]) int {
			return func(hx_value_31 any) int {
				if hx_value_31 == nil {
					var hx_zero_32 int
					return hx_zero_32
				}
				return hx_value_31.(int)
			}(any(hxrt.FunctionCall(source)(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
