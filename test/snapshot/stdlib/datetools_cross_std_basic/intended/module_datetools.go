package main

import "snapshot/hxrt"

var DateTools_DAYS_OF_MONTH []int = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

var DateTools_DAY_NAMES []*string = []*string{hxrt.StringFromLiteral("Sunday"), hxrt.StringFromLiteral("Monday"), hxrt.StringFromLiteral("Tuesday"), hxrt.StringFromLiteral("Wednesday"), hxrt.StringFromLiteral("Thursday"), hxrt.StringFromLiteral("Friday"), hxrt.StringFromLiteral("Saturday")}

var DateTools_DAY_SHORT_NAMES []*string = []*string{hxrt.StringFromLiteral("Sun"), hxrt.StringFromLiteral("Mon"), hxrt.StringFromLiteral("Tue"), hxrt.StringFromLiteral("Wed"), hxrt.StringFromLiteral("Thu"), hxrt.StringFromLiteral("Fri"), hxrt.StringFromLiteral("Sat")}

var DateTools_MONTH_NAMES []*string = []*string{hxrt.StringFromLiteral("January"), hxrt.StringFromLiteral("February"), hxrt.StringFromLiteral("March"), hxrt.StringFromLiteral("April"), hxrt.StringFromLiteral("May"), hxrt.StringFromLiteral("June"), hxrt.StringFromLiteral("July"), hxrt.StringFromLiteral("August"), hxrt.StringFromLiteral("September"), hxrt.StringFromLiteral("October"), hxrt.StringFromLiteral("November"), hxrt.StringFromLiteral("December")}

var DateTools_MONTH_SHORT_NAMES []*string = []*string{hxrt.StringFromLiteral("Jan"), hxrt.StringFromLiteral("Feb"), hxrt.StringFromLiteral("Mar"), hxrt.StringFromLiteral("Apr"), hxrt.StringFromLiteral("May"), hxrt.StringFromLiteral("Jun"), hxrt.StringFromLiteral("Jul"), hxrt.StringFromLiteral("Aug"), hxrt.StringFromLiteral("Sep"), hxrt.StringFromLiteral("Oct"), hxrt.StringFromLiteral("Nov"), hxrt.StringFromLiteral("Dec")}

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
		return DateTools_DAY_SHORT_NAMES[d.getDay()]
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("A")) {
		return DateTools_DAY_NAMES[d.getDay()]
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("b")) || hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("h")) {
		return DateTools_MONTH_SHORT_NAMES[d.getMonth()]
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("B")) {
		return DateTools_MONTH_NAMES[d.getMonth()]
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("C")) {
		return StringTools_lpad(hxrt.StdString(Math_floor((float64(d.getFullYear()) / float64(100)))), hxrt.StringFromLiteral("0"), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("d")) {
		return StringTools_lpad(hxrt.StdString(d.getDate()), hxrt.StringFromLiteral("0"), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("D")) {
		return DateTools___format(d, hxrt.StringFromLiteral("%m/%d/%y"))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("e")) {
		return hxrt.StdString(d.getDate())
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("F")) {
		return DateTools___format(d, hxrt.StringFromLiteral("%Y-%m-%d"))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("H")) || hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("k")) {
		return StringTools_lpad(hxrt.StdString(d.getHours()), func() *string {
			var hx_if_17 *string
			if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("H")) {
				hx_if_17 = hxrt.StringFromLiteral("0")
			} else {
				hx_if_17 = hxrt.StringFromLiteral(" ")
			}
			return hx_if_17
		}(), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("I")) || hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("l")) {
		hour := int(int32((hxrt.Int32Wrap(d.getHours()) % hxrt.Int32Wrap(12))))
		return StringTools_lpad(hxrt.StdString(func() any {
			var hx_if_18 any
			if hour == 0 {
				hx_if_18 = 12
			} else {
				hx_if_18 = hour
			}
			return hx_if_18
		}()), func() *string {
			var hx_if_19 *string
			if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("I")) {
				hx_if_19 = hxrt.StringFromLiteral("0")
			} else {
				hx_if_19 = hxrt.StringFromLiteral(" ")
			}
			return hx_if_19
		}(), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("m")) {
		return StringTools_lpad(hxrt.StdString(int(int32((hxrt.Int32Wrap(d.getMonth()) + hxrt.Int32Wrap(1))))), hxrt.StringFromLiteral("0"), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("M")) {
		return StringTools_lpad(hxrt.StdString(d.getMinutes()), hxrt.StringFromLiteral("0"), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("n")) {
		return hxrt.StringFromLiteral("\n")
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("p")) {
		var hx_if_20 *string
		if d.getHours() > 11 {
			hx_if_20 = hxrt.StringFromLiteral("PM")
		} else {
			hx_if_20 = hxrt.StringFromLiteral("AM")
		}
		return hx_if_20
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("r")) {
		return DateTools___format(d, hxrt.StringFromLiteral("%I:%M:%S %p"))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("R")) {
		return DateTools___format(d, hxrt.StringFromLiteral("%H:%M"))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("s")) {
		return hxrt.StdString(Math_floor((d.getTime() / float64(1000))))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("S")) {
		return StringTools_lpad(hxrt.StdString(d.getSeconds()), hxrt.StringFromLiteral("0"), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("t")) {
		return hxrt.StringFromLiteral("\t")
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("T")) {
		return DateTools___format(d, hxrt.StringFromLiteral("%H:%M:%S"))
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("u")) {
		day := d.getDay()
		var hx_if_21 *string
		if day == 0 {
			hx_if_21 = hxrt.StringFromLiteral("7")
		} else {
			hx_if_21 = hxrt.StdString(day)
		}
		return hx_if_21
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("w")) {
		return hxrt.StdString(d.getDay())
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("y")) {
		return StringTools_lpad(hxrt.StdString(int(int32((hxrt.Int32Wrap(d.getFullYear()) % hxrt.Int32Wrap(100))))), hxrt.StringFromLiteral("0"), 2)
	}
	if hxrt.StringEqualStringPtr(e, hxrt.StringFromLiteral("Y")) {
		return hxrt.StdString(d.getFullYear())
	}
	hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Date.format %"), e), hxrt.StringFromLiteral(" not implemented yet.")))
	var hx_throw_zero_22 *string
	return hx_throw_zero_22
}

func DateTools_days(n float64) float64 {
	return ((((n * 24.0) * 60.0) * 60.0) * 1000.0)
}

func DateTools_delta(d *Date, t float64) *Date {
	return Date_fromTime((d.getTime() + t))
}

func DateTools_format(d *Date, f *string) *string {
	return DateTools___format(d, f)
}

func DateTools_getMonthDays(d *Date) int {
	month := d.getMonth()
	year := d.getFullYear()
	if month != 1 {
		return DateTools_DAYS_OF_MONTH[month]
	}
	isLeap := (((int(int32((hxrt.Int32Wrap(year) % hxrt.Int32Wrap(4)))) == 0) && (int(int32((hxrt.Int32Wrap(year) % hxrt.Int32Wrap(100)))) != 0)) || (int(int32((hxrt.Int32Wrap(year) % hxrt.Int32Wrap(400)))) == 0))
	var hx_if_23 int
	if isLeap {
		hx_if_23 = 29
	} else {
		hx_if_23 = 28
	}
	return hx_if_23
}

func DateTools_hours(n float64) float64 {
	return (((n * 60.0) * 60.0) * 1000.0)
}

func DateTools_make(o map[string]any) float64 {
	return (func(hx_obj_24 map[string]any) float64 {
		hx_field_25 := hx_obj_24["ms"]
		if hx_field_25 == nil {
			var hx_zero_26 float64
			return hx_zero_26
		}
		return hx_field_25.(float64)
	}(o) + (1000.0 * (float64(func(hx_obj_27 map[string]any) int {
		hx_field_28 := hx_obj_27["seconds"]
		if hx_field_28 == nil {
			var hx_zero_29 int
			return hx_zero_29
		}
		return hx_field_28.(int)
	}(o)) + (60.0 * (float64(func(hx_obj_30 map[string]any) int {
		hx_field_31 := hx_obj_30["minutes"]
		if hx_field_31 == nil {
			var hx_zero_32 int
			return hx_zero_32
		}
		return hx_field_31.(int)
	}(o)) + (60.0 * (float64(func(hx_obj_33 map[string]any) int {
		hx_field_34 := hx_obj_33["hours"]
		if hx_field_34 == nil {
			var hx_zero_35 int
			return hx_zero_35
		}
		return hx_field_34.(int)
	}(o)) + (24.0 * float64(func(hx_obj_36 map[string]any) int {
		hx_field_37 := hx_obj_36["days"]
		if hx_field_37 == nil {
			var hx_zero_38 int
			return hx_zero_38
		}
		return hx_field_37.(int)
	}(o))))))))))
}

func DateTools_minutes(n float64) float64 {
	return ((n * 60.0) * 1000.0)
}

func DateTools_parse(t float64) map[string]any {
	s := (t / float64(1000))
	m := (s / float64(60))
	h := (m / float64(60))
	hx_obj_39 := map[string]any{}
	hx_obj_39["ms"] = hxrt.FloatMod(t, float64(1000))
	hx_obj_39["seconds"] = Math_floor(hxrt.FloatMod(s, float64(60)))
	hx_obj_39["minutes"] = Math_floor(hxrt.FloatMod(m, float64(60)))
	hx_obj_39["hours"] = Math_floor(hxrt.FloatMod(h, float64(24)))
	hx_obj_39["days"] = Math_floor((h / float64(24)))
	return hx_obj_39
}

func DateTools_seconds(n float64) float64 {
	return (n * 1000.0)
}
