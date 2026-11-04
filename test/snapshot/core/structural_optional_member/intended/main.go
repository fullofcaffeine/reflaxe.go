package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_C interface {
}

type C struct {
	__hx_this I_C
	n         int
}

func (self *C) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_C() *C {
	return hx_init_C(hx_alloc_C(nil))
}

func hx_alloc_C(owner any) *C {
	self := &C{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_C)
	}
	return self
}

func hx_init_C(self *C) *C {
	self.n = 1
	return self
}

func main() {
	var erased any = New_C()
	view := hx_structural_recover(hx_structural_recover(erased))
	var v any = any(func(hx_value_1 any) int {
		if hx_value_1 == nil {
			var hx_zero_2 int
			return hx_zero_2
		}
		return hx_value_1.(int)
	}(hxrt.AnonymousGet(view, "n")))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringEqualStringPtr(func(hx_value_3 any) *string {
		if hx_value_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_value_3.(*string)
	}(hxrt.AnonymousGet(view, "label")), nil))
	hxrt.Println(v_1)
	var v_2 any = any(Reflect_hasField(view, hxrt.StringFromLiteral("label")))
	hxrt.Println(v_2)
	present := func(hx_structural_source_5 *Present) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_5, func(name string) hxrt.AnonymousField {
			switch name {
			case "label":
				return hxrt.AnonymousField{Value: hx_structural_source_5.label, Found: true}
			case "n":
				return hxrt.AnonymousField{Value: hx_structural_source_5.n, Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "label":
				hx_structural_source_5.label = func(hx_value_6 any) *string {
					if hx_value_6 == nil {
						var hx_zero_7 *string
						return hx_zero_7
					}
					return hx_value_6.(*string)
				}(incoming)
				return true
			case "n":
				hx_structural_source_5.n = func(hx_value_8 any) int {
					if hx_value_8 == nil {
						var hx_zero_9 int
						return hx_zero_9
					}
					return hx_value_8.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(New_Present())
	var v_3 any = any(hxrt.StringEqualStringPtr(func(hx_value_10 any) *string {
		if hx_value_10 == nil {
			var hx_zero_11 *string
			return hx_zero_11
		}
		return hx_value_10.(*string)
	}(hxrt.AnonymousGet(present, "label")), nil))
	hxrt.Println(v_3)
	var v_4 any = any(Reflect_hasField(present, hxrt.StringFromLiteral("label")))
	hxrt.Println(v_4)
}

type I_Present interface {
}

type Present struct {
	*C
	__hx_this I_Present
	label     *string
}

func (self *Present) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Present() *Present {
	return hx_init_Present(hx_alloc_Present(nil))
}

func hx_alloc_Present(owner any) *Present {
	self := &Present{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Present)
	}
	self.C = hx_alloc_C(self.__hx_this)
	return self
}

func hx_init_Present(self *Present) *Present {
	hx_init_C(self.C)
	self.label = nil
	return self
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *C:
		return func(hx_structural_source_14 *C) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_14, func(name string) hxrt.AnonymousField {
				switch name {
				case "n":
					return hxrt.AnonymousField{Value: hx_structural_source_14.n, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "n":
					hx_structural_source_14.n = func(hx_value_15 any) int {
						if hx_value_15 == nil {
							var hx_zero_16 int
							return hx_zero_16
						}
						return hx_value_15.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *Present:
		return func(hx_structural_source_17 *Present) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_17, func(name string) hxrt.AnonymousField {
				switch name {
				case "label":
					return hxrt.AnonymousField{Value: hx_structural_source_17.label, Found: true}
				case "n":
					return hxrt.AnonymousField{Value: hx_structural_source_17.n, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "label":
					hx_structural_source_17.label = func(hx_value_18 any) *string {
						if hx_value_18 == nil {
							var hx_zero_19 *string
							return hx_zero_19
						}
						return hx_value_18.(*string)
					}(incoming)
					return true
				case "n":
					hx_structural_source_17.n = func(hx_value_20 any) int {
						if hx_value_20 == nil {
							var hx_zero_21 int
							return hx_zero_21
						}
						return hx_value_20.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_22 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_22, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_22.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_22.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_22.high = func(hx_value_23 any) int {
						if hx_value_23 == nil {
							var hx_zero_24 int
							return hx_zero_24
						}
						return hx_value_23.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_22.low = func(hx_value_25 any) int {
						if hx_value_25 == nil {
							var hx_zero_26 int
							return hx_zero_26
						}
						return hx_value_25.(int)
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
	case *C:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Present:
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
	case *C:
		return hxrt__generated_field_lookup__C(value, key)
	case *Present:
		return hxrt__generated_field_lookup__Present(value, key)
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
	case *C:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *Present:
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
	case *C:
		return hxrt__generated_field_has__C(value, key)
	case *Present:
		return hxrt__generated_field_has__Present(value, key)
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
	case *C:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *Present:
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
	case *C:
		return hxrt__generated_field_set__C(value, key, incoming)
	case *Present:
		return hxrt__generated_field_set__Present(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *C:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Present:
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
	case *C:
		return hxrt.NewArray(hxrt.StringFromLiteral("n"))
	case *Present:
		return hxrt.NewArray(hxrt.StringFromLiteral("n"), hxrt.StringFromLiteral("label"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__C(value *C, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "n":
		return hxrt.ReadStoredField(value.n)
	}
	return nil
}

func hxrt__generated_field_has__C(value *C, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "n":
		return true
	}
	return false
}

func hxrt__generated_field_set__C(value *C, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "n":
		if incoming == nil {
			var zero int
			value.n = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.n = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__Present(value *Present, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "label":
		return hxrt.ReadStoredField(value.label)
	}
	if value.C == nil {
		return nil
	}
	return hxrt__generated_field_lookup__C(value.C, key)
}

func hxrt__generated_field_has__Present(value *Present, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "label":
		return true
	}
	if value.C == nil {
		return false
	}
	return hxrt__generated_field_has__C(value.C, key)
}

func hxrt__generated_field_set__Present(value *Present, key string, incoming any) bool {
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
	}
	if value.C == nil {
		return false
	}
	return hxrt__generated_field_set__C(value.C, key, incoming)
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
