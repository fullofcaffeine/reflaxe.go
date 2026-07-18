package main

import "snapshot/hxrt"

func main() {
	base := Date_fromString(hxrt.StringFromLiteral("2024-02-29 15:04:05"))
	var v any = any(DateTools_format(base, hxrt.StringFromLiteral("%D|%F|%R|%T|%r|%a|%A|%b|%h|%B|%C|%d|%e|%H|%k|%I|%l|%m|%M|%p|%S|%u|%w|%y|%Y|%%")))
	hxrt.Println(v)
	var v_1 any = any(DateTools_getMonthDays(base))
	hxrt.Println(v_1)
	stamp := DateTools_make(func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["ms"] = 123.0
		hx_obj_1["seconds"] = 5
		hx_obj_1["minutes"] = 4
		hx_obj_1["hours"] = 3
		hx_obj_1["days"] = 2
		return hx_obj_1
	}())
	parsed := DateTools_parse(stamp)
	var v_2 any = any(func(hx_obj_2 map[string]any) int {
		hx_field_3 := hx_obj_2["days"]
		if hx_field_3 == nil {
			var hx_zero_4 int
			return hx_zero_4
		}
		return hx_field_3.(int)
	}(parsed))
	hxrt.Println(v_2)
	var v_3 any = any(func(hx_obj_5 map[string]any) int {
		hx_field_6 := hx_obj_5["hours"]
		if hx_field_6 == nil {
			var hx_zero_7 int
			return hx_zero_7
		}
		return hx_field_6.(int)
	}(parsed))
	hxrt.Println(v_3)
	var v_4 any = any(func(hx_obj_8 map[string]any) int {
		hx_field_9 := hx_obj_8["minutes"]
		if hx_field_9 == nil {
			var hx_zero_10 int
			return hx_zero_10
		}
		return hx_field_9.(int)
	}(parsed))
	hxrt.Println(v_4)
	var v_5 any = any(func(hx_obj_11 map[string]any) int {
		hx_field_12 := hx_obj_11["seconds"]
		if hx_field_12 == nil {
			var hx_zero_13 int
			return hx_zero_13
		}
		return hx_field_12.(int)
	}(parsed))
	hxrt.Println(v_5)
	var v_6 any = any(func(hx_obj_14 map[string]any) float64 {
		hx_field_15 := hx_obj_14["ms"]
		if hx_field_15 == nil {
			var hx_zero_16 float64
			return hx_zero_16
		}
		return hx_field_15.(float64)
	}(parsed))
	hxrt.Println(v_6)
	shifted := Date_fromTime((base.ms + 93784000.))
	var v_7 any = any(DateTools_format(shifted, hxrt.StringFromLiteral("%F %T")))
	hxrt.Println(v_7)
}
