package main

import "snapshot/hxrt"

func main() {
	hxrt.Println(any(hxrt.StringFromLiteral("ok")))
}

type hxrt__test_ast_Point struct {
	X     int
	Y     int
	Label *string
}

func hxrt__test_ast_structured_composite_control_smoke() {
	point := &hxrt__test_ast_Point{X: 1, Y: 2, Label: hxrt.StringFromLiteral("ab")}
	values := []int{1, 2, 3}
	indexed := [3]int{0: 4, 2: 6}
	lookup := map[string]int{"one": 1, "two": 2}
	total := 0
	for index := 0; index < len(values); index++ {
		total += values[index]
	}
	total++
	total--
	_ = point
	_ = indexed
	_ = lookup
}
