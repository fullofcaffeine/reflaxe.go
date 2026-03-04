package main

import (
	"snapshot/hxrt"
	"strconv"
)

func main() {
	ok := go__result_fromValueError(strconv.Atoi(hxrt.StringValue(hxrt.StringFromLiteral("42"))))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ok.isOk="), hxrt.StdString(func(hx_value_1 any) bool {
		if hx_value_1 == nil {
			var hx_zero_2 bool
			return hx_zero_2
		}
		return hx_value_1.(bool)
	}(ok.isOk()))))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ok.isErr="), hxrt.StdString(func(hx_value_3 any) bool {
		if hx_value_3 == nil {
			var hx_zero_4 bool
			return hx_zero_4
		}
		return hx_value_3.(bool)
	}(ok.isErr()))))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("ok.unwrap="), func(hx_value_5 any) int {
		if hx_value_5 == nil {
			var hx_zero_6 int
			return hx_zero_6
		}
		return hx_value_5.(int)
	}(ok.unwrap())))
	err := go__result_fromValueError(strconv.Atoi(hxrt.StringValue(hxrt.StringFromLiteral("x"))))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("err.isOk="), hxrt.StdString(func(hx_value_7 any) bool {
		if hx_value_7 == nil {
			var hx_zero_8 bool
			return hx_zero_8
		}
		return hx_value_7.(bool)
	}(err.isOk()))))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("err.isErr="), hxrt.StdString(func(hx_value_9 any) bool {
		if hx_value_9 == nil {
			var hx_zero_10 bool
			return hx_zero_10
		}
		return hx_value_9.(bool)
	}(err.isErr()))))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("err.hasError="), hxrt.StdString(!hxrt.StringEqualStringPtr(func(hx_value_11 any) *string {
		if hx_value_11 == nil {
			var hx_zero_12 *string
			return hx_zero_12
		}
		return hx_value_11.(*string)
	}(err.error()), nil))))
}

func go__result_fromValueError(value any, err error) *go___Result {
	if err != nil {
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral(err.Error())))
	}
	return New_go___Result(value, nil)
}
