package main

import (
	"errors"
	"snapshot/hxrt"
)

var calls *hxrt.Array = hxrt.NewArray()

func describe(error *go___Error) *string {
	var hx_if_1 *string
	if error == nil {
		hx_if_1 = hxrt.StringFromLiteral("null")
	} else {
		hx_if_1 = StringTools_replace(error.message, hxrt.StringFromLiteral("\n"), hxrt.StringFromLiteral("|"))
	}
	return hx_if_1
}

func joinAll(errors_local []*go___Error) *go___Error {
	return func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.Join(hxrt.ExportErrors(errors_local, hxrt.ExportLegacyError[*go___Error])...))
}

func joinWithNull(error *go___Error) *go___Error {
	return func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.Join([]hxrt.RawError{hxrt.ExportLegacyError(error), nil}...))
}

func legacy(message *string) *go___Error {
	hx_arr_2 := calls
	hx_arr_2.Push(message)
	return New_go___Error(message)
}

func main() {
	first := New_go___Error(hxrt.StringFromLiteral("first"))
	second := func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.New(*hxrt.StdString(hxrt.StringFromLiteral("second"))))
	var none *go___Error = nil
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("literal="), describe(func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.Join([]hxrt.RawError{hxrt.ExportLegacyError(first), hxrt.ExportLegacyError(none), hxrt.ExportLegacyError(second)}...)))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("all-null="), describe(func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.Join([]hxrt.RawError{hxrt.ExportLegacyError(none), nil}...)))))
	hxrt.Println(v_1)
	list := hxrt.NewArray(first, none, second)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("spread="), describe(func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.Join(hxrt.ExportErrors(func() []*go___Error {
		var this1 *hxrt.Array
		this1 = list
		return func(hx_lambda_raw_3 []any) []*go___Error {
			hx_lambda_out_4 := make([]*go___Error, 0, len(hx_lambda_raw_3))
			for _, hx_lambda_item_5 := range hx_lambda_raw_3 {
				hx_lambda_out_4 = append(hx_lambda_out_4, hx_recover_class_go___Error(hx_lambda_item_5))
			}
			return hx_lambda_out_4
		}(this1.Values())
	}(), hxrt.ExportLegacyError[*go___Error])...)))))
	hxrt.Println(v_2)
	nulls := hxrt.NewArray(nil, none)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("spread-null="), describe(func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.Join(hxrt.ExportErrors(func() []*go___Error {
		var this1_1 *hxrt.Array
		this1_1 = nulls
		return func(hx_lambda_raw_6 []any) []*go___Error {
			hx_lambda_out_7 := make([]*go___Error, 0, len(hx_lambda_raw_6))
			for _, hx_lambda_item_8 := range hx_lambda_raw_6 {
				hx_lambda_out_7 = append(hx_lambda_out_7, hx_recover_class_go___Error(hx_lambda_item_8))
			}
			return hx_lambda_out_7
		}(this1_1.Values())
	}(), hxrt.ExportLegacyError[*go___Error])...)))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("nested="), describe(func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.Join([]hxrt.RawError{hxrt.ExportLegacyError(func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.Join([]hxrt.RawError{hxrt.ExportLegacyError(first), hxrt.ExportLegacyError(none)}...))), hxrt.ExportLegacyError(second)}...)))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("rest="), describe(joinAll([]*go___Error{first, none, second}))))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("shadow="), describe(joinWithNull(second))))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("unwrap="), describe(func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.Unwrap(hxrt.ExportLegacyError(first))))))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order="), describe(func(err interface{ Error() string }) *go___Error {
		if err == nil {
			return nil
		}
		return New_go___Error(hxrt.StringFromLiteral(err.Error()))
	}(errors.Join([]hxrt.RawError{hxrt.ExportLegacyError(legacy(hxrt.StringFromLiteral("a"))), hxrt.ExportLegacyError(none), hxrt.ExportLegacyError(legacy(hxrt.StringFromLiteral("b")))}...)))), hxrt.StringFromLiteral(" ")), hxrt.StringJoinAny(calls.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_8)
	x := hxrt.ImportNativeError(errors.New(*hxrt.StdString(hxrt.StringFromLiteral("x"))))
	y := hxrt.ImportNativeError(errors.New(*hxrt.StdString(hxrt.StringFromLiteral("y"))))
	var missing *hxrt.NativeErrorValue = nil
	joined := hxrt.ImportNativeError(errors.Join([]hxrt.RawError{hxrt.ExportNativeError(x), hxrt.ExportNativeError(missing), hxrt.ExportNativeError(y)}...))
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native="), StringTools_replace(hxrt.StdString(joined.Text()), hxrt.StringFromLiteral("\n"), hxrt.StringFromLiteral("|"))))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native-is="), hxrt.StdString(errors.Is(hxrt.ExportNativeError(joined), hxrt.ExportNativeError(x)))), hxrt.StringFromLiteral(",")), hxrt.StdString(errors.Is(hxrt.ExportNativeError(joined), hxrt.ExportNativeError(y)))))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native-is-copy="), hxrt.StdString(errors.Is(hxrt.ExportNativeError(joined), hxrt.ExportNativeError(hxrt.ImportNativeError(errors.New(*hxrt.StdString(hxrt.StringFromLiteral("x")))))))))
	hxrt.Println(v_11)
	natives := hxrt.NewArray(x, missing, y)
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native-spread-is="), hxrt.StdString(errors.Is(hxrt.ExportNativeError(hxrt.ImportNativeError(errors.Join(hxrt.ExportErrors(func() []*hxrt.NativeErrorValue {
		var this1_2 *hxrt.Array
		this1_2 = natives
		return func(hx_lambda_raw_9 []any) []*hxrt.NativeErrorValue {
			hx_lambda_out_10 := make([]*hxrt.NativeErrorValue, 0, len(hx_lambda_raw_9))
			for _, hx_lambda_item_11 := range hx_lambda_raw_9 {
				hx_lambda_out_10 = append(hx_lambda_out_10, func(hx_value_12 any) *hxrt.NativeErrorValue {
					if hx_value_12 == nil {
						var hx_zero_13 *hxrt.NativeErrorValue
						return hx_zero_13
					}
					return hx_value_12.(*hxrt.NativeErrorValue)
				}(hx_lambda_item_11))
			}
			return hx_lambda_out_10
		}(this1_2.Values())
	}(), hxrt.ExportNativeError)...))), hxrt.ExportNativeError(y)))))
	hxrt.Println(v_12)
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native-nested-is="), hxrt.StdString(errors.Is(hxrt.ExportNativeError(hxrt.ImportNativeError(errors.Join([]hxrt.RawError{hxrt.ExportNativeError(hxrt.ImportNativeError(errors.Join([]hxrt.RawError{hxrt.ExportNativeError(x)}...))), hxrt.ExportNativeError(y)}...))), hxrt.ExportNativeError(x)))))
	hxrt.Println(v_13)
	var v_14 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native-null="), hxrt.StdString((hxrt.ImportNativeError(errors.Join([]hxrt.RawError{hxrt.ExportNativeError(missing), hxrt.ExportNativeError(nil)}...)) == nil))))
	hxrt.Println(v_14)
}

func hx_recover_class_go___Error(value any) *go___Error {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *go___Error:
		return source
	default:
		return value.(*go___Error)
	}
}
