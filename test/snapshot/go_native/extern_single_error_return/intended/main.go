package main

import (
	"context"
	"errors"
	"snapshot/hxrt"
)

func main() {
	noError := func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(context.Background().Err())
	error := func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.New(*hxrt.StdString(hxrt.StringFromLiteral("broken"))))
	func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.New(*hxrt.StdString(hxrt.StringFromLiteral("ignored"))))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("nil="), hxrt.StdString((noError == nil))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("message="), error.toString()))
	hxrt.Println(v_1)
}
