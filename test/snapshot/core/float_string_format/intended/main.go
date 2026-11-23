package main

import "snapshot/hxrt"

type Holder struct {
	tag    int
	params []any
}

func (self *Holder) HxEnumValue() (string, []any) {
	return [1]string{"Hold"}[self.tag], self.params
}

func Holder_Hold(value float64) *Holder {
	enumValue := &Holder{tag: 0}
	enumValue.params = []any{value}
	return enumValue
}

func main() {
	floats := values()
	_g := 0
	for _g < floats.Len() {
		value := func(hx_value_1 any) float64 {
			if hx_value_1 == nil {
				var hx_zero_2 float64
				return hx_zero_2
			}
			return hx_value_1.(float64)
		}(floats.Get(_g))
		_g = int(int32((_g + 1)))
		var dyn any = value
		var buffer_b *string
		buffer_b = hxrt.StringFromLiteral("")
		buffer_b = hxrt.StringConcatStringPtr(buffer_b, hxrt.StdString(value))
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(value), hxrt.StringFromLiteral(" | ")), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), value)), hxrt.StringFromLiteral(" | ")), hxrt.StringConcatAny(hxrt.StringFromLiteral(""), value)), hxrt.StringFromLiteral(" | ")), buffer_b), hxrt.StringFromLiteral(" | ")), hxrt.StdString(dyn)), hxrt.StringFromLiteral(" | ")), hxrt.StdString(hxrt.NewArray(value))), hxrt.StringFromLiteral(" | ")), hxrt.StdString(Holder_Hold(value))), hxrt.StringFromLiteral(" | ")), hxrt.StdString(func() hxrt.Anonymous {
			hx_obj_3 := map[string]any{}
			hx_obj_3["value"] = value
			return hxrt.AnonymousFromMap(hx_obj_3)
		}())), hxrt.StringFromLiteral(" | ")), haxe__Log_formatOutput(value, nil)))
		hxrt.Println(v)
	}
	hxrt.Println(any(hxrt.StringFromLiteral("--- Sys.println")))
	_g_1 := 0
	for _g_1 < floats.Len() {
		value_1 := func(hx_value_4 any) float64 {
			if hx_value_4 == nil {
				var hx_zero_5 float64
				return hx_zero_5
			}
			return hx_value_4.(float64)
		}(floats.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		hxrt.Println(any(value_1))
	}
	hxrt.Println(any(hxrt.StringFromLiteral("--- typed")))
	count := 2147483647
	widened := float64(count)
	hxrt.Println(any(widened))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("widened="), widened)))
	lowest := -2147483648.0
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("lowest="), lowest)))
	sample := New_Sample(3e8)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("field="), sample.value), hxrt.StringFromLiteral(" half=")), sample.half()))
	hxrt.Println(v_1)
	var maybe any = 1e8
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("nullable="), maybe)))
	floatArray := hxrt.NewArray(1e8, 0.300000000000000044, 1e-7)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("array="), hxrt.StdString(floatArray)))
	hxrt.Println(v_2)
	total := 0.0
	total = (total + 1e8)
	total = (total + 1e8)
	total = (total + 1e8)
	total = (total + 1e8)
	total = (total + 1e8)
	total = (total + 1e8)
	total = (total + 1e8)
	total = (total + 1e8)
	total = (total + 1e8)
	total = (total + 1e8)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("sum="), total)))
	hxrt.Println(any(hxrt.StringFromLiteral("--- json")))
	zero := 0.0
	jsonValues := hxrt.NewArray(1e8, 2147483647.0, 0.300000000000000044, 1e15, 1e21, 1e-7, -zero, 1.5)
	var v_3 any = any(func() *string {
		var space *string = nil
		return hxrt.StdString(hxrt.JsonStringify(any(jsonValues), space))
	}())
	hxrt.Println(v_3)
	var v_4 any = any(func() *string {
		var space_1 *string = nil
		return hxrt.StdString(hxrt.JsonStringify(any(func() hxrt.Anonymous {
			hx_obj_6 := map[string]any{}
			hx_obj_6["big"] = 1e8
			return hxrt.AnonymousFromMap(hx_obj_6)
		}()), space_1))
	}())
	hxrt.Println(v_4)
	nonFinite := hxrt.NewArray(1.5, Math_NaN, Math_POSITIVE_INFINITY)
	var v_5 any = any(func() *string {
		var space_2 *string = nil
		return hxrt.StdString(hxrt.JsonStringify(any(nonFinite), space_2))
	}())
	hxrt.Println(v_5)
}

func values() *hxrt.Array {
	zero := 0.0
	return hxrt.NewArray(0.0, -zero, 1.0, -1.0, 0.5, 1.5, 100.0, 3.14159, 0.1, 0.000123, 0.0001234567, 0.99609375, 1e5, 1e6, 1e7, 12345678.9, 1e8, -1e8, 123456789.0, 999999999.0, 2147483647.0, -2147483648.0, 4294967296.0, 1000000000.5, 123456789.123, 100000000000.5, 999999999999.0, 1e12, -1e12, 1000000000001., 2e12, 1234567890000.0, 99999999999999.0, 100000000000000.25, 999999999999999.0, 1e15, 1000000000000000.5, 9007199254740992.0, 1e16, 1e17, 123456789012345678.0, 1234567890123456789.0, 1e20, 1e21, -1e21, 1e22, 1e23, 1e100, 1.5e300, 1.7976931348623157e308, 2.2250738585072014e-308, 5e-324, 0.300000000000000044, 0.333333333333333315, -0.66666666666666663, 0.300000000000000044, 1e-4, 1e-5, 1e-6, 1e-7, 1.5e-7, 2.5e-5, 9e-6, 1e-10, 123e-20, 1e-300, Math_NaN, Math_POSITIVE_INFINITY, Math_NEGATIVE_INFINITY)
}

type I_Sample interface {
	half() float64
}

type Sample struct {
	__hx_this I_Sample
	value     float64
}

func (self *Sample) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Sample(value float64) *Sample {
	return hx_init_Sample(hx_alloc_Sample(nil), value)
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

func hx_init_Sample(self *Sample, value float64) *Sample {
	self.value = value
	return self
}

func (self *Sample) half() float64 {
	return (self.value / float64(2))
}

func (self *Sample) HxClassName() string {
	return "Sample"
}
