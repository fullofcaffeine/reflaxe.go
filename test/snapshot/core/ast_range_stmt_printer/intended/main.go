package main

import "snapshot/hxrt"

func main() {
	hxrt.Println(any(hxrt.StringFromLiteral("ok")))
}

func hxrt__test_ast_range_stmt_smoke() {
	items := map[string]int{"a": 1, "b": 2}
	for key, value := range items {
		_ = key
		_ = value
	}
	var seenKey string
	for seenKey = range items {
		_ = seenKey
	}
	for index := range []int{1, 2, 3} {
		_ = index
	}
}
