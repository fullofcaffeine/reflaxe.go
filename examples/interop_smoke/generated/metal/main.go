package main

import (
	"context"
	"errors"
	"examples_interop_smoke_metal/hxrt"
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
	valueErrorPass := (((((go__result_isOk__int_95e97e5e(valueErrorOk) && !go__result_isErr__int_95e97e5e(valueErrorOk)) && (go__result_unwrap__int_95e97e5e(valueErrorOk) == 42)) && !go__result_isOk__int_95e97e5e(valueErrorErr)) && go__result_isErr__int_95e97e5e(valueErrorErr)) && !hxrt.StringEqualStringPtr(go__result_error__int_95e97e5e(valueErrorErr), nil))
	fmt.Println(func() int {
		var hx_if_1 int
		if (wrappedOk && externOk) && valueErrorPass {
			hx_if_1 = 1
		} else {
			hx_if_1 = 0
		}
		return hx_if_1
	}())
}

func go__result_fromValueError(value any, err error) *go___Result {
	if err != nil {
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral(err.Error())))
	}
	return New_go___Result(value, nil)
}

func go__result_ok__int_95e97e5e(value int) *go___Result {
	return New_go___Result(value, nil)
}

func go__result_failure__int_95e97e5e(message *string) *go___Result {
	return New_go___Result(nil, New_go___Error(message))
}

func go__result_valueError__int_95e97e5e(result *go___Result) (int, error) {
	var zero int
	if result == nil {
		return zero, errors.New("nil go.Result")
	}
	if result.errorValue != nil {
		return zero, errors.New(*hxrt.StdString(result.errorValue.message))
	}
	if result.value == nil {
		return zero, nil
	}
	return result.value.(int), nil
}

func go__result_isOk__int_95e97e5e(result *go___Result) bool {
	_, err := go__result_valueError__int_95e97e5e(result)
	return (err == nil)
}

func go__result_isErr__int_95e97e5e(result *go___Result) bool {
	_, err := go__result_valueError__int_95e97e5e(result)
	return (err != nil)
}

func go__result_unwrap__int_95e97e5e(result *go___Result) int {
	value, err := go__result_valueError__int_95e97e5e(result)
	if err != nil {
		hxrt.Throw(hxrt.StringFromLiteral(err.Error()))
		var zero int
		return zero
	}
	return value
}

func go__result_error__int_95e97e5e(result *go___Result) *string {
	_, err := go__result_valueError__int_95e97e5e(result)
	if err == nil {
		return nil
	}
	return hxrt.StringFromLiteral(err.Error())
}
