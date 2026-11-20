package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"snapshot/hxrt"
	"strings"
)

func arguments() *hxrt.Array {
	return hxrt.NewArray(hxrt.StringFromLiteral("-n"), name())
}

var calls *hxrt.Array = hxrt.NewArray()

func fileName() *string {
	hx_arr_1 := calls
	hx_arr_1.Push(hxrt.StringFromLiteral("fileName"))
	var this1 *string
	this1 = hxrt.StringFromLiteral("local.yaml")
	return this1
}

func joinAll(parts []*string) *string {
	return hxrt.StdString(filepath.Join(hxrt.StringSlice(parts)...))
}

func main() {
	var v any = any(hxrt.StdString(filepath.Join([]string{*hxrt.StdString(hxrt.StdString(filepath.Dir(*hxrt.StdString(hxrt.StringFromLiteral("/a/b/c.txt"))))), *hxrt.StdString(hxrt.StringFromLiteral("d.txt"))}...)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(filepath.Join([]string{*hxrt.StdString(hxrt.StringFromLiteral("/root")), *hxrt.StdString(name())}...)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StdString(filepath.Join([]string{*hxrt.StdString(hxrt.StringFromLiteral("/root")), *hxrt.StdString(fileName())}...)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StdString(filepath.Join([]string{*hxrt.StdString(name()), *hxrt.StdString(hxrt.StringFromLiteral("x")), *hxrt.StdString(hxrt.StdString(filepath.Base(*hxrt.StdString(hxrt.StringFromLiteral("/q/y")))))}...)))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StdString(filepath.Join([]string{*hxrt.StdString(hxrt.StdString(filepath.Dir(*hxrt.StdString(hxrt.StringFromLiteral("/a/b/c.txt"))))), *hxrt.StdString(name())}...)))
	hxrt.Println(v_4)
	rest := hxrt.NewArray(hxrt.StringFromLiteral("e"), hxrt.StringFromLiteral("f"))
	var v_5 any = any(hxrt.StdString(exec.Command(*hxrt.StdString(hxrt.StdString(filepath.Join([]string{*hxrt.StdString(hxrt.StringFromLiteral("/bin")), *hxrt.StdString(hxrt.StringFromLiteral("echo"))}...))), hxrt.StringSlice(func() []*string {
		var this1 *hxrt.Array
		this1 = rest
		return func(hx_lambda_raw_2 []any) []*string {
			hx_lambda_out_3 := make([]*string, 0, len(hx_lambda_raw_2))
			for _, hx_lambda_item_4 := range hx_lambda_raw_2 {
				hx_lambda_out_3 = append(hx_lambda_out_3, func(hx_value_5 any) *string {
					if hx_value_5 == nil {
						var hx_zero_6 *string
						return hx_zero_6
					}
					return hx_value_5.(*string)
				}(hx_lambda_item_4))
			}
			return hx_lambda_out_3
		}(this1.Values())
	}())...).String()))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StdString(exec.Command(*hxrt.StdString(hxrt.StringFromLiteral("/bin/echo")), hxrt.StringSlice(func() []*string {
		array := arguments()
		var this1_1 *hxrt.Array
		this1_1 = array
		return func(hx_lambda_raw_7 []any) []*string {
			hx_lambda_out_8 := make([]*string, 0, len(hx_lambda_raw_7))
			for _, hx_lambda_item_9 := range hx_lambda_raw_7 {
				hx_lambda_out_8 = append(hx_lambda_out_8, func(hx_value_10 any) *string {
					if hx_value_10 == nil {
						var hx_zero_11 *string
						return hx_zero_11
					}
					return hx_value_10.(*string)
				}(hx_lambda_item_9))
			}
			return hx_lambda_out_8
		}(this1_1.Values())
	}())...).String()))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StdString(exec.Command(*hxrt.StdString(hxrt.StringFromLiteral("/bin/echo")), []string{*hxrt.StdString(hxrt.StdString(filepath.Base(*hxrt.StdString(hxrt.StringFromLiteral("/q/y"))))), *hxrt.StdString(name())}...).String()))
	hxrt.Println(v_7)
	replacer := strings.NewReplacer([]string{*hxrt.StdString(hxrt.StdString(strings.ToUpper(*hxrt.StdString(hxrt.StringFromLiteral("a"))))), *hxrt.StdString(hxrt.StringFromLiteral("1")), *hxrt.StdString(fileName()), *hxrt.StdString(hxrt.StringFromLiteral("2"))}...)
	var v_8 any = any(hxrt.StdString(replacer.Replace(*hxrt.StdString(hxrt.StringFromLiteral("A local.yaml")))))
	hxrt.Println(v_8)
	var v_9 any = any(joinAll(func() []*string {
		array_1 := hxrt.NewArray(hxrt.StringFromLiteral("p"), name(), fileName())
		var this1_2 *hxrt.Array
		this1_2 = array_1
		return func(hx_lambda_raw_12 []any) []*string {
			hx_lambda_out_13 := make([]*string, 0, len(hx_lambda_raw_12))
			for _, hx_lambda_item_14 := range hx_lambda_raw_12 {
				hx_lambda_out_13 = append(hx_lambda_out_13, func(hx_value_15 any) *string {
					if hx_value_15 == nil {
						var hx_zero_16 *string
						return hx_zero_16
					}
					return hx_value_15.(*string)
				}(hx_lambda_item_14))
			}
			return hx_lambda_out_13
		}(this1_2.Values())
	}()))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StdString(fmt.Sprint([]any{*hxrt.StdString(hxrt.StringFromLiteral("<")), *hxrt.StdString(name()), *hxrt.StdString(hxrt.StringFromLiteral("|")), *hxrt.StdString(fileName()), *hxrt.StdString(hxrt.StringFromLiteral(">"))}...)))
	hxrt.Println(v_10)
	calls = hxrt.NewArray()
	hxrt.StdString(filepath.Join([]string{*hxrt.StdString(fileName()), *hxrt.StdString(name()), *hxrt.StdString(fileName())}...))
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("order="), hxrt.StringJoinAny(calls.Values(), hxrt.StringFromLiteral(","))))
	hxrt.Println(v_11)
}

func name() *string {
	hx_arr_17 := calls
	hx_arr_17.Push(hxrt.StringFromLiteral("name"))
	return hxrt.StringFromLiteral("config.yaml")
}

func _Main__FileName_Impl___new(value *string) *string {
	var this1 *string
	this1 = value
	return this1
}
