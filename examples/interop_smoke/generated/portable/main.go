package main

import (
	"context"
	"examples_interop_smoke_portable/hxrt"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	wrappedNow := time.Now()
	wrappedUnixDirect := wrappedNow.Unix()
	wrappedUnixReceiver := wrappedNow.Unix()
	var wrappedCtx context.Context = context.Background()
	wrappedStatusOk := hxrt.StringEqualStringPtr(hxrt.StdString(http.StatusText(200)), hxrt.StringFromLiteral("OK"))
	wrappedOk := ((((wrappedUnixDirect == wrappedUnixReceiver) && (wrappedUnixDirect > 0)) && (wrappedCtx != nil)) && wrappedStatusOk)
	externNow := time.Now()
	externUnixDirect := externNow.Unix()
	externUnixReceiver := externNow.Unix()
	var externCtx context.Context = context.Background()
	externStatusOk := hxrt.StringEqualStringPtr(hxrt.StdString(http.StatusText(200)), hxrt.StringFromLiteral("OK"))
	externOk := ((((externUnixDirect == externUnixReceiver) && (externUnixDirect > 0)) && (externCtx != nil)) && externStatusOk)
	valueErrorOk := go__result_fromValueError(strconv.Atoi(hxrt.StringValue(hxrt.StringFromLiteral("42"))))
	valueErrorErr := go__result_fromValueError(strconv.Atoi(hxrt.StringValue(hxrt.StringFromLiteral("oops"))))
	valueErrorPass := (((((func(hx_value_1 any) bool {
		if hx_value_1 == nil {
			var hx_zero_2 bool
			return hx_zero_2
		}
		return hx_value_1.(bool)
	}(valueErrorOk.isOk()) && !func(hx_value_3 any) bool {
		if hx_value_3 == nil {
			var hx_zero_4 bool
			return hx_zero_4
		}
		return hx_value_3.(bool)
	}(valueErrorOk.isErr())) && (func(hx_value_5 any) int {
		if hx_value_5 == nil {
			var hx_zero_6 int
			return hx_zero_6
		}
		return hx_value_5.(int)
	}(valueErrorOk.unwrap()) == 42)) && !func(hx_value_7 any) bool {
		if hx_value_7 == nil {
			var hx_zero_8 bool
			return hx_zero_8
		}
		return hx_value_7.(bool)
	}(valueErrorErr.isOk())) && func(hx_value_9 any) bool {
		if hx_value_9 == nil {
			var hx_zero_10 bool
			return hx_zero_10
		}
		return hx_value_9.(bool)
	}(valueErrorErr.isErr())) && !hxrt.StringEqualStringPtr(func(hx_value_11 any) *string {
		if hx_value_11 == nil {
			var hx_zero_12 *string
			return hx_zero_12
		}
		return hx_value_11.(*string)
	}(valueErrorErr.error()), nil))
	fmt.Println(func() int {
		var hx_if_13 int
		if (wrappedOk && externOk) && valueErrorPass {
			hx_if_13 = 1
		} else {
			hx_if_13 = 0
		}
		return hx_if_13
	}())
}

func go__result_fromValueError(value any, err error) *go___Result {
	if err != nil {
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral(err.Error())))
	}
	return New_go___Result(value, nil)
}
