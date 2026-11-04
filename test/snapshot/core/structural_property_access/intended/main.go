package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_Counter interface {
	get_value() int
	set_value(value int) int
}

type Counter struct {
	__hx_this I_Counter
	value     int
	reads     int
	writes    int
	stored    int
}

func (self *Counter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Counter() *Counter {
	return hx_init_Counter(hx_alloc_Counter(nil))
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

func hx_init_Counter(self *Counter) *Counter {
	self.stored = 4
	self.writes = 0
	self.reads = 0
	return self
}

func (self *Counter) get_value() int {
	self.reads = int(int32((self.reads + 1)))
	return self.stored
}

func (self *Counter) set_value(value int) int {
	self.writes = int(int32((self.writes + 1)))
	self.stored = value
	return value
}

func main() {
	source := New_Counter()
	var erased any = source
	view := hx_structural_recover(hx_structural_recover(erased))
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("before="), func() int {
		var hx_callable_1 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(view, "get_value"))
		return hxrt.FunctionCall(hx_callable_1)()
	}()))
	hxrt.Println(v)
	assigned := func() int {
		var hx_callable_2 *hxrt.Function[func(int) int] = hx_recover_callable_1(hxrt.AnonymousGet(view, "set_value"))
		var hx_argument_3 int = 9
		return hxrt.FunctionCall(hx_callable_2)(hx_argument_3)
	}()
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("assigned="), assigned)))
	func() int {
		var hx_callable_5 *hxrt.Function[func(int) int] = hx_recover_callable_1(hxrt.AnonymousGet(view, "set_value"))
		var hx_argument_6 int = int((hxrt.Int32Wrap(func() int {
			var hx_callable_4 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(view, "get_value"))
			return hxrt.FunctionCall(hx_callable_4)()
		}()) + hxrt.Int32Wrap(3)))
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("after="), func() int {
		var hx_callable_7 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(view, "get_value"))
		return hxrt.FunctionCall(hx_callable_7)()
	}()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("reads="), source.reads))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("writes="), source.writes))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw-absent="), hxrt.StdString(!Reflect_hasField(view, hxrt.StringFromLiteral("value")))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw-null="), hxrt.StdString(hxrt.AnyEqualsNull(Reflect_field(view, hxrt.StringFromLiteral("value"))))))
	hxrt.Println(v_5)
	raw := hx_structural_recover(hx_structural_recover(erased))
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("plain-null="), hxrt.StdString((func(hx_value_8 any) any {
		if hx_value_8 == nil {
			return nil
		}
		return hx_value_8.(int)
	}(hxrt.AnonymousGet(raw, "value")) == nil))))
	hxrt.Println(v_6)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *Counter:
		return func(hx_structural_source_23 *Counter) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_23, func(name string) hxrt.AnonymousField {
				switch name {
				case "reads":
					return hxrt.AnonymousField{Value: hx_structural_source_23.reads, Found: true}
				case "writes":
					return hxrt.AnonymousField{Value: hx_structural_source_23.writes, Found: true}
				case "stored":
					return hxrt.AnonymousField{Value: hx_structural_source_23.stored, Found: true}
				case "get_value":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_30 := hx_structural_source_23
						return hxrt.BoundFunction(func() int {
							return hx_receiver_30.__hx_this.get_value()
						}, func() any {
							return hx_receiver_30.__hx_this
						}, "get_value")
					}(), Found: true}
				case "set_value":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_31 := hx_structural_source_23
						return hxrt.BoundFunction(func(value int) int {
							return hx_receiver_31.__hx_this.set_value(value)
						}, func() any {
							return hx_receiver_31.__hx_this
						}, "set_value")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "reads":
					hx_structural_source_23.reads = func(hx_value_24 any) int {
						if hx_value_24 == nil {
							var hx_zero_25 int
							return hx_zero_25
						}
						return hx_value_24.(int)
					}(incoming)
					return true
				case "writes":
					hx_structural_source_23.writes = func(hx_value_26 any) int {
						if hx_value_26 == nil {
							var hx_zero_27 int
							return hx_zero_27
						}
						return hx_value_26.(int)
					}(incoming)
					return true
				case "stored":
					hx_structural_source_23.stored = func(hx_value_28 any) int {
						if hx_value_28 == nil {
							var hx_zero_29 int
							return hx_zero_29
						}
						return hx_value_28.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_32 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_32, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_32.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_32.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_32.high = func(hx_value_33 any) int {
						if hx_value_33 == nil {
							var hx_zero_34 int
							return hx_zero_34
						}
						return hx_value_33.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_32.low = func(hx_value_35 any) int {
						if hx_value_35 == nil {
							var hx_zero_36 int
							return hx_zero_36
						}
						return hx_value_35.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *Counter:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Counter:
		return hxrt__generated_method_field__Counter(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Counter(value *Counter, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "get_value":
		return func() *hxrt.Function[func() int] {
			hx_receiver_37 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_37.__hx_this.get_value()
			}, func() any {
				return hx_receiver_37.__hx_this
			}, "get_value")
		}()
	case "set_value":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_38 := value
			return hxrt.BoundFunction(func(value int) int {
				return hx_receiver_38.__hx_this.set_value(value)
			}, func() any {
				return hx_receiver_38.__hx_this
			}, "set_value")
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
	case *Counter:
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
	case *Counter:
		return hxrt__generated_field_lookup__Counter(value, key)
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
	case *Counter:
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
	case *Counter:
		return hxrt__generated_field_has__Counter(value, key)
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
	case *Counter:
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
	case *Counter:
		return hxrt__generated_field_set__Counter(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Counter:
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
	case *Counter:
		return hxrt.NewArray(hxrt.StringFromLiteral("reads"), hxrt.StringFromLiteral("writes"), hxrt.StringFromLiteral("stored"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__Counter(value *Counter, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "reads":
		return hxrt.ReadStoredField(value.reads)
	case "stored":
		return hxrt.ReadStoredField(value.stored)
	case "writes":
		return hxrt.ReadStoredField(value.writes)
	}
	return nil
}

func hxrt__generated_field_has__Counter(value *Counter, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "reads":
		return true
	case "stored":
		return true
	case "writes":
		return true
	}
	return false
}

func hxrt__generated_field_set__Counter(value *Counter, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "reads":
		if incoming == nil {
			var zero int
			value.reads = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.reads = typed
			return true
		default:
			return false
		}
	case "stored":
		if incoming == nil {
			var zero int
			value.stored = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.stored = typed
			return true
		default:
			return false
		}
	case "writes":
		if incoming == nil {
			var zero int
			value.writes = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.writes = typed
			return true
		default:
			return false
		}
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
	key := *hxrt.StdString(field)
	return hxrt__generated_method_field(object, key)
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	return false
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
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_39 any) int {
			if hx_value_39 == nil {
				var hx_zero_40 int
				return hx_zero_40
			}
			return hx_value_39.(int)
		}(any(argument_0)))
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func() int] {
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
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_41 any) int {
				if hx_value_41 == nil {
					var hx_zero_42 int
					return hx_zero_42
				}
				return hx_value_41.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func(int) int] {
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
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_43 any) int {
				if hx_value_43 == nil {
					var hx_zero_44 int
					return hx_zero_44
				}
				return hx_value_43.(int)
			}(any(hxrt.FunctionCall(source)(arg_0)))
		})
	case *hxrt.Function[func(int) int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
