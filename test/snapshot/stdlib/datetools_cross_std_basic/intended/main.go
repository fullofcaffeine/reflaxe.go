package main

import "snapshot/hxrt"

func main() {
	base := Date_fromString(hxrt.StringFromLiteral("2024-02-29 15:04:05"))
	var v any = any(DateTools_format(base, hxrt.StringFromLiteral("%D|%F|%R|%T|%r|%a|%A|%b|%h|%B|%C|%d|%e|%H|%k|%I|%l|%m|%M|%p|%S|%u|%w|%y|%Y|%%")))
	hxrt.Println(v)
	var v_1 any = any(DateTools_getMonthDays(base))
	hxrt.Println(v_1)
	stamp := DateTools_make(func() hxrt.Anonymous {
		hx_obj_1 := map[string]any{}
		hx_obj_1["ms"] = 123.0
		hx_obj_1["seconds"] = 5
		hx_obj_1["minutes"] = 4
		hx_obj_1["hours"] = 3
		hx_obj_1["days"] = 2
		return hxrt.AnonymousFromMap(hx_obj_1)
	}())
	parsed := DateTools_parse(stamp)
	var v_2 any = any(func(hx_value_2 any) int {
		if hx_value_2 == nil {
			var hx_zero_3 int
			return hx_zero_3
		}
		return hx_value_2.(int)
	}(hxrt.AnonymousGet(parsed, "days")))
	hxrt.Println(v_2)
	var v_3 any = any(func(hx_value_4 any) int {
		if hx_value_4 == nil {
			var hx_zero_5 int
			return hx_zero_5
		}
		return hx_value_4.(int)
	}(hxrt.AnonymousGet(parsed, "hours")))
	hxrt.Println(v_3)
	var v_4 any = any(func(hx_value_6 any) int {
		if hx_value_6 == nil {
			var hx_zero_7 int
			return hx_zero_7
		}
		return hx_value_6.(int)
	}(hxrt.AnonymousGet(parsed, "minutes")))
	hxrt.Println(v_4)
	var v_5 any = any(func(hx_value_8 any) int {
		if hx_value_8 == nil {
			var hx_zero_9 int
			return hx_zero_9
		}
		return hx_value_8.(int)
	}(hxrt.AnonymousGet(parsed, "seconds")))
	hxrt.Println(v_5)
	var v_6 any = any(func(hx_value_10 any) float64 {
		if hx_value_10 == nil {
			var hx_zero_11 float64
			return hx_zero_11
		}
		return hx_value_10.(float64)
	}(hxrt.AnonymousGet(parsed, "ms")))
	hxrt.Println(v_6)
	shifted := Date_fromTime((base.ms + 93784000.))
	var v_7 any = any(DateTools_format(shifted, hxrt.StringFromLiteral("%F %T")))
	hxrt.Println(v_7)
}

func hx_recover_class_0(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_1(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
