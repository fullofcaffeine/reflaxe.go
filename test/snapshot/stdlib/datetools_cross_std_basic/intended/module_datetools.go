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
		return func(hx_value_19 any) *string {
			if hx_value_19 == nil {
				var hx_zero_20 *string
				return hx_zero_20
			}
			return hx_value_19.(*string)
		}(DateTools_DAY_SHORT_NAMES.Get(d.__hx_this.getDay()))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("A")) {
		return func(hx_value_21 any) *string {
			if hx_value_21 == nil {
				var hx_zero_22 *string
				return hx_zero_22
			}
			return hx_value_21.(*string)
		}(DateTools_DAY_NAMES.Get(d.__hx_this.getDay()))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("b")) || hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("h")) {
		return func(hx_value_23 any) *string {
			if hx_value_23 == nil {
				var hx_zero_24 *string
				return hx_zero_24
			}
			return hx_value_23.(*string)
		}(DateTools_MONTH_SHORT_NAMES.Get(d.__hx_this.getMonth()))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("B")) {
		return func(hx_value_25 any) *string {
			if hx_value_25 == nil {
				var hx_zero_26 *string
				return hx_zero_26
			}
			return hx_value_25.(*string)
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
			var hx_if_27 *string
			if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("H")) {
				hx_if_27 = hxrt.StringFromLiteral("0")
			} else {
				hx_if_27 = hxrt.StringFromLiteral(" ")
			}
			return hx_if_27
		}(), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("I")) || hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("l")) {
		hour := int(int32((hxrt.Int32Wrap(d.__hx_this.getHours()) % hxrt.Int32Wrap(12))))
		return StringTools_lpad(hxrt.StdString(func() any {
			var hx_if_28 any
			if hour == 0 {
				hx_if_28 = 12
			} else {
				hx_if_28 = hour
			}
			return hx_if_28
		}()), func() *string {
			var hx_if_29 *string
			if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("I")) {
				hx_if_29 = hxrt.StringFromLiteral("0")
			} else {
				hx_if_29 = hxrt.StringFromLiteral(" ")
			}
			return hx_if_29
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
		var hx_if_30 *string
		if d.__hx_this.getHours() > 11 {
			hx_if_30 = hxrt.StringFromLiteral("PM")
		} else {
			hx_if_30 = hxrt.StringFromLiteral("AM")
		}
		return hx_if_30
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
		var hx_if_31 *string
		if day == 0 {
			hx_if_31 = hxrt.StringFromLiteral("7")
		} else {
			hx_if_31 = hxrt.StdString(day)
		}
		return hx_if_31
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
	var hx_throw_zero_32 *string
	return hx_throw_zero_32
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
	var hx_if_33 int
	if isLeap {
		hx_if_33 = 29
	} else {
		hx_if_33 = 28
	}
	return hx_if_33
}

func DateTools_hours(n float64) float64 {
	return (((n * 60.0) * 60.0) * 1000.0)
}

func DateTools_make(o map[string]any) float64 {
	return (func(hx_obj_34 map[string]any) float64 {
		hx_field_35 := hx_obj_34["ms"]
		if hx_field_35 == nil {
			var hx_zero_36 float64
			return hx_zero_36
		}
		return hx_field_35.(float64)
	}(o) + (1000.0 * (float64(func(hx_obj_37 map[string]any) int {
		hx_field_38 := hx_obj_37["seconds"]
		if hx_field_38 == nil {
			var hx_zero_39 int
			return hx_zero_39
		}
		return hx_field_38.(int)
	}(o)) + (60.0 * (float64(func(hx_obj_40 map[string]any) int {
		hx_field_41 := hx_obj_40["minutes"]
		if hx_field_41 == nil {
			var hx_zero_42 int
			return hx_zero_42
		}
		return hx_field_41.(int)
	}(o)) + (60.0 * (float64(func(hx_obj_43 map[string]any) int {
		hx_field_44 := hx_obj_43["hours"]
		if hx_field_44 == nil {
			var hx_zero_45 int
			return hx_zero_45
		}
		return hx_field_44.(int)
	}(o)) + (24.0 * float64(func(hx_obj_46 map[string]any) int {
		hx_field_47 := hx_obj_46["days"]
		if hx_field_47 == nil {
			var hx_zero_48 int
			return hx_zero_48
		}
		return hx_field_47.(int)
	}(o))))))))))
}

func DateTools_minutes(n float64) float64 {
	return ((n * 60.0) * 1000.0)
}

func DateTools_parse(t float64) map[string]any {
	s := (t / float64(1000))
	m := (s / float64(60))
	h := (m / float64(60))
	hx_obj_49 := map[string]any{}
	hx_obj_49["ms"] = hxrt.FloatMod(t, float64(1000))
	hx_obj_49["seconds"] = hxrt.MathFloorInt(hxrt.FloatMod(s, float64(60)))
	hx_obj_49["minutes"] = hxrt.MathFloorInt(hxrt.FloatMod(m, float64(60)))
	hx_obj_49["hours"] = hxrt.MathFloorInt(hxrt.FloatMod(h, float64(24)))
	hx_obj_49["days"] = hxrt.MathFloorInt((h / float64(24)))
	return hx_obj_49
}

func DateTools_seconds(n float64) float64 {
	return (n * 1000.0)
}
