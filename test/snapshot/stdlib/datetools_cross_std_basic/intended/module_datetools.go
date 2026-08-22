package main

import "snapshot/hxrt"

var DateTools_DAYS_OF_MONTH *hxrt.Array = hxrt.NewArray(31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31)

var DateTools_DAY_NAMES *hxrt.Array = hxrt.NewArray(hxrt.StringFromLiteral("Sunday"), hxrt.StringFromLiteral("Monday"), hxrt.StringFromLiteral("Tuesday"), hxrt.StringFromLiteral("Wednesday"), hxrt.StringFromLiteral("Thursday"), hxrt.StringFromLiteral("Friday"), hxrt.StringFromLiteral("Saturday"))

var DateTools_DAY_SHORT_NAMES *hxrt.Array = hxrt.NewArray(hxrt.StringFromLiteral("Sun"), hxrt.StringFromLiteral("Mon"), hxrt.StringFromLiteral("Tue"), hxrt.StringFromLiteral("Wed"), hxrt.StringFromLiteral("Thu"), hxrt.StringFromLiteral("Fri"), hxrt.StringFromLiteral("Sat"))

var DateTools_MONTH_NAMES *hxrt.Array = hxrt.NewArray(hxrt.StringFromLiteral("January"), hxrt.StringFromLiteral("February"), hxrt.StringFromLiteral("March"), hxrt.StringFromLiteral("April"), hxrt.StringFromLiteral("May"), hxrt.StringFromLiteral("June"), hxrt.StringFromLiteral("July"), hxrt.StringFromLiteral("August"), hxrt.StringFromLiteral("September"), hxrt.StringFromLiteral("October"), hxrt.StringFromLiteral("November"), hxrt.StringFromLiteral("December"))

var DateTools_MONTH_SHORT_NAMES *hxrt.Array = hxrt.NewArray(hxrt.StringFromLiteral("Jan"), hxrt.StringFromLiteral("Feb"), hxrt.StringFromLiteral("Mar"), hxrt.StringFromLiteral("Apr"), hxrt.StringFromLiteral("May"), hxrt.StringFromLiteral("Jun"), hxrt.StringFromLiteral("Jul"), hxrt.StringFromLiteral("Aug"), hxrt.StringFromLiteral("Sep"), hxrt.StringFromLiteral("Oct"), hxrt.StringFromLiteral("Nov"), hxrt.StringFromLiteral("Dec"))

func DateTools___format(d *Date, f *string) *string {
	var result_b *string
	result_b = hxrt.StringFromLiteral("")
	pos := 0
	length := hxrt.StringLengthStringPtr(f)
	for pos < length {
		if hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(f, pos), hxrt.StringFromLiteral("%")) {
			x := DateTools___format_get(d, hxrt.StringSubstrStringPtr(f, int(int32((hxrt.Int32Wrap(pos)+hxrt.Int32Wrap(1)))), 1, true))
			result_b = hxrt.StringConcatStringPtr(result_b, hxrt.StdString(x))
			pos = int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(2))))
			continue
		}
		x_1 := hxrt.StringCharAtStringPtr(f, pos)
		result_b = hxrt.StringConcatStringPtr(result_b, hxrt.StdString(x_1))
		pos = int(int32((pos + 1)))
	}
	return result_b
}

func DateTools___format_get(d *Date, e *string) *string {
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("%")) {
		return hxrt.StringFromLiteral("%")
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("a")) {
		return func(hx_value_1 any) *string {
			if hx_value_1 == nil {
				var hx_zero_2 *string
				return hx_zero_2
			}
			return hx_value_1.(*string)
		}(DateTools_DAY_SHORT_NAMES.Get(d.__hx_this.getDay()))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("A")) {
		return func(hx_value_3 any) *string {
			if hx_value_3 == nil {
				var hx_zero_4 *string
				return hx_zero_4
			}
			return hx_value_3.(*string)
		}(DateTools_DAY_NAMES.Get(d.__hx_this.getDay()))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("b")) || hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("h")) {
		return func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(DateTools_MONTH_SHORT_NAMES.Get(d.__hx_this.getMonth()))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("B")) {
		return func(hx_value_7 any) *string {
			if hx_value_7 == nil {
				var hx_zero_8 *string
				return hx_zero_8
			}
			return hx_value_7.(*string)
		}(DateTools_MONTH_NAMES.Get(d.__hx_this.getMonth()))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("C")) {
		return StringTools_lpad(hxrt.StdString(func() int {
			v := (float64(d.__hx_this.getFullYear()) / float64(100))
			return hxrt.MathFloorInt(v)
		}()), hxrt.StringFromLiteral("0"), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("d")) {
		return StringTools_lpad(hxrt.StdString(d.__hx_this.getDate()), hxrt.StringFromLiteral("0"), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("D")) {
		return DateTools___format(d, hxrt.StringFromLiteral("%m/%d/%y"))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("e")) {
		return hxrt.StdString(d.__hx_this.getDate())
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("F")) {
		return DateTools___format(d, hxrt.StringFromLiteral("%Y-%m-%d"))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("H")) || hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("k")) {
		return StringTools_lpad(hxrt.StdString(d.__hx_this.getHours()), func() *string {
			var hx_if_9 *string
			if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("H")) {
				hx_if_9 = hxrt.StringFromLiteral("0")
			} else {
				hx_if_9 = hxrt.StringFromLiteral(" ")
			}
			return hx_if_9
		}(), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("I")) || hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("l")) {
		hour := int(int32((hxrt.Int32Wrap(d.__hx_this.getHours()) % hxrt.Int32Wrap(12))))
		return StringTools_lpad(hxrt.StdString(func() any {
			var hx_if_10 any
			if hour == 0 {
				hx_if_10 = 12
			} else {
				hx_if_10 = hour
			}
			return hx_if_10
		}()), func() *string {
			var hx_if_11 *string
			if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("I")) {
				hx_if_11 = hxrt.StringFromLiteral("0")
			} else {
				hx_if_11 = hxrt.StringFromLiteral(" ")
			}
			return hx_if_11
		}(), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("m")) {
		return StringTools_lpad(hxrt.StdString(int(int32((hxrt.Int32Wrap(d.__hx_this.getMonth()) + hxrt.Int32Wrap(1))))), hxrt.StringFromLiteral("0"), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("M")) {
		return StringTools_lpad(hxrt.StdString(d.__hx_this.getMinutes()), hxrt.StringFromLiteral("0"), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("n")) {
		return hxrt.StringFromLiteral("\n")
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("p")) {
		var hx_if_12 *string
		if d.__hx_this.getHours() > 11 {
			hx_if_12 = hxrt.StringFromLiteral("PM")
		} else {
			hx_if_12 = hxrt.StringFromLiteral("AM")
		}
		return hx_if_12
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("r")) {
		return DateTools___format(d, hxrt.StringFromLiteral("%I:%M:%S %p"))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("R")) {
		return DateTools___format(d, hxrt.StringFromLiteral("%H:%M"))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("s")) {
		return hxrt.StdString(func() int {
			v_1 := (d.ms / float64(1000))
			return hxrt.MathFloorInt(v_1)
		}())
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("S")) {
		return StringTools_lpad(hxrt.StdString(d.__hx_this.getSeconds()), hxrt.StringFromLiteral("0"), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("t")) {
		return hxrt.StringFromLiteral("\t")
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("T")) {
		return DateTools___format(d, hxrt.StringFromLiteral("%H:%M:%S"))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("u")) {
		day := d.__hx_this.getDay()
		var hx_if_13 *string
		if day == 0 {
			hx_if_13 = hxrt.StringFromLiteral("7")
		} else {
			hx_if_13 = hxrt.StdString(day)
		}
		return hx_if_13
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("w")) {
		return hxrt.StdString(d.__hx_this.getDay())
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("y")) {
		return StringTools_lpad(hxrt.StdString(int(int32((hxrt.Int32Wrap(d.__hx_this.getFullYear()) % hxrt.Int32Wrap(100))))), hxrt.StringFromLiteral("0"), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("Y")) {
		return hxrt.StdString(d.__hx_this.getFullYear())
	}
	hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Date.format %"), e), hxrt.StringFromLiteral(" not implemented yet.")))
	var hx_throw_zero_14 *string
	return hx_throw_zero_14
}

func DateTools_days(n float64) float64 {
	return ((((n * 24.0) * 60.0) * 60.0) * 1000.0)
}

func DateTools_delta(d *Date, t float64) *Date {
	return Date_fromTime((d.ms + t))
}

func DateTools_format(d *Date, f *string) *string {
	return DateTools___format(d, f)
}

func DateTools_getMonthDays(d *Date) int {
	month := d.__hx_this.getMonth()
	year := d.__hx_this.getFullYear()
	if month != 1 {
		return hxrt.IntFromNullableAny(DateTools_DAYS_OF_MONTH.Get(month))
	}
	isLeap := (((int(int32((hxrt.Int32Wrap(year) % hxrt.Int32Wrap(4)))) == 0) && (int(int32((hxrt.Int32Wrap(year) % hxrt.Int32Wrap(100)))) != 0)) || (int(int32((hxrt.Int32Wrap(year) % hxrt.Int32Wrap(400)))) == 0))
	var hx_if_15 int
	if isLeap {
		hx_if_15 = 29
	} else {
		hx_if_15 = 28
	}
	return hx_if_15
}

func DateTools_hours(n float64) float64 {
	return (((n * 60.0) * 60.0) * 1000.0)
}

func DateTools_make(o map[string]any) float64 {
	return (func(hx_obj_16 map[string]any) float64 {
		hx_field_17 := hx_obj_16["ms"]
		if hx_field_17 == nil {
			var hx_zero_18 float64
			return hx_zero_18
		}
		return hx_field_17.(float64)
	}(o) + (1000.0 * (float64(func(hx_obj_19 map[string]any) int {
		hx_field_20 := hx_obj_19["seconds"]
		if hx_field_20 == nil {
			var hx_zero_21 int
			return hx_zero_21
		}
		return hx_field_20.(int)
	}(o)) + (60.0 * (float64(func(hx_obj_22 map[string]any) int {
		hx_field_23 := hx_obj_22["minutes"]
		if hx_field_23 == nil {
			var hx_zero_24 int
			return hx_zero_24
		}
		return hx_field_23.(int)
	}(o)) + (60.0 * (float64(func(hx_obj_25 map[string]any) int {
		hx_field_26 := hx_obj_25["hours"]
		if hx_field_26 == nil {
			var hx_zero_27 int
			return hx_zero_27
		}
		return hx_field_26.(int)
	}(o)) + (24.0 * float64(func(hx_obj_28 map[string]any) int {
		hx_field_29 := hx_obj_28["days"]
		if hx_field_29 == nil {
			var hx_zero_30 int
			return hx_zero_30
		}
		return hx_field_29.(int)
	}(o))))))))))
}

func DateTools_minutes(n float64) float64 {
	return ((n * 60.0) * 1000.0)
}

func DateTools_parse(t float64) map[string]any {
	s := (t / float64(1000))
	m := (s / float64(60))
	h := (m / float64(60))
	hx_obj_31 := map[string]any{}
	hx_obj_31["ms"] = hxrt.FloatMod(t, float64(1000))
	hx_obj_31["seconds"] = hxrt.MathFloorInt(hxrt.FloatMod(s, float64(60)))
	hx_obj_31["minutes"] = hxrt.MathFloorInt(hxrt.FloatMod(m, float64(60)))
	hx_obj_31["hours"] = hxrt.MathFloorInt(hxrt.FloatMod(h, float64(24)))
	hx_obj_31["days"] = hxrt.MathFloorInt((h / float64(24)))
	return hx_obj_31
}

func DateTools_seconds(n float64) float64 {
	return (n * 1000.0)
}
