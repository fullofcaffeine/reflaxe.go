package main

import "snapshot/hxrt"

func Reflect_callMethod(o any, func_ any, args *hxrt.Array) any {
	return hxrt.ReflectCallMethod(func_, args.ValuesCopy())
}

func Reflect_compare(a any, b any) int {
	return hxrt.ReflectCompare(a, b)
}

func Reflect_compareMethods(f1 any, f2 any) bool {
	return hxrt.ReflectCompareMethods(f1, f2)
}

func Reflect_copy(o any) any {
	return hxrt.ReflectCopy(o)
}

func Reflect_deleteField(o any, field *string) bool {
	return hxrt.ReflectDeleteField(o, field)
}

func Reflect_field(o any, field *string) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	var typeValue any = reflaxe__go___internal__CompilerReflect_typeField(o, field)
	if !hxrt.AnyEqualsNull(typeValue) || reflaxe__go___internal__CompilerReflect_hasTypeField(o, field) {
		return typeValue
	}
	nativeField := hxrt.ReflectLookupField(o, field)
	if nativeField.Found {
		return nativeField.Value
	}
	var generatedField any = reflaxe__go___internal__CompilerReflect_generatedField(o, field)
	if !hxrt.AnyEqualsNull(generatedField) || reflaxe__go___internal__CompilerReflect_hasGeneratedField(o, field) {
		return generatedField
	}
	var generatedMethod any = reflaxe__go___internal__CompilerReflect_generatedMethod(o, field)
	if !hxrt.AnyEqualsNull(generatedMethod) {
		return generatedMethod
	}
	nativeMethod := hxrt.ReflectLookupMethod(o, field)
	var hx_if_4 any
	if nativeMethod.Found {
		hx_if_4 = nativeMethod.Value
	} else {
		hx_if_4 = nil
	}
	return hx_if_4
}

func Reflect_fields(o any) *hxrt.Array {
	generatedFields := reflaxe__go___internal__CompilerReflect_generatedFields(o)
	if generatedFields != nil {
		return generatedFields
	}
	return hxrt.ArrayFromValues(func(hx_sort_src_5 []*string) []any {
		hx_sort_out_7 := make([]any, 0, len(hx_sort_src_5))
		for _, hx_sort_item_6 := range hx_sort_src_5 {
			hx_sort_out_7 = append(hx_sort_out_7, hx_sort_item_6)
		}
		return hx_sort_out_7
	}(hxrt.ReflectFields(o)))
}

func Reflect_getProperty(o any, field *string) any {
	var getter any = Reflect_field(o, hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("get_"), field))
	var hx_if_8 any
	if hxrt.AnyEqualsNull(getter) {
		hx_if_8 = Reflect_field(o, field)
	} else {
		hx_if_8 = Reflect_callMethod(o, getter, hxrt.NewArray())
	}
	return hx_if_8
}

func Reflect_hasField(o any, field *string) bool {
	if hxrt.AnyEqualsNull(o) {
		return false
	}
	if reflaxe__go___internal__CompilerReflect_hasTypeField(o, field) {
		return true
	}
	if hxrt.ReflectLookupField(o, field).Found {
		return true
	}
	if reflaxe__go___internal__CompilerReflect_hasGeneratedField(o, field) {
		return true
	}
	if !hxrt.AnyEqualsNull(reflaxe__go___internal__CompilerReflect_generatedMethod(o, field)) {
		return true
	}
	return hxrt.ReflectLookupMethod(o, field).Found
}

func Reflect_isEnumValue(v any) bool {
	return reflaxe__go___internal__CompilerReflect_isEnumValue(v)
}

func Reflect_isFunction(f any) bool {
	return hxrt.ReflectIsFunction(f)
}

func Reflect_isObject(v any) bool {
	return hxrt.ReflectIsObject(v)
}

func Reflect_makeVarArgs(f func(*hxrt.Array) any) any {
	return hxrt.ReflectMakeVarArgs(f)
}

func Reflect_setField(o any, field *string, value any) {
	if hxrt.AnyEqualsNull(o) {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
	}
	if hxrt.ReflectSetField(o, field, value) {
		return
	}
	reflaxe__go___internal__CompilerReflect_setGeneratedField(o, field, value)
}

func Reflect_setProperty(o any, field *string, value any) {
	var setter any = Reflect_field(o, hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("set_"), field))
	if hxrt.AnyEqualsNull(setter) {
		Reflect_setField(o, field, value)
		return
	}
	Reflect_callMethod(o, setter, hxrt.NewArray(value))
}
