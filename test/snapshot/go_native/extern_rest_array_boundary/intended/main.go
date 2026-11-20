package main

import (
	"fmt"
	"snapshot/hxrt"
)

var evaluations int = 0

func main() {
	var record_number int
	record_number = 7
	var v any = any(hxrt.StdString(fmt.Sprintf(*hxrt.StdString(hxrt.StringFromLiteral("%d")), []any{3}...)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(fmt.Sprintf(*hxrt.StdString(hxrt.StringFromLiteral("%d")), []any{record_number}...)))
	hxrt.Println(v_1)
	local := 11
	var v_2 any = any(hxrt.StdString(fmt.Sprintf(*hxrt.StdString(hxrt.StringFromLiteral("%d")), []any{local}...)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StdString(fmt.Sprintf(*hxrt.StdString(hxrt.StringFromLiteral("%d:%d")), []any{next(), next()}...)))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("evaluations="), evaluations))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StdString(fmt.Sprintf(*hxrt.StdString(hxrt.StringFromLiteral("empty")), []any{}...)))
	hxrt.Println(v_5)
}

func next() int {
	evaluations = int(int32((evaluations + 1)))
	return evaluations
}
