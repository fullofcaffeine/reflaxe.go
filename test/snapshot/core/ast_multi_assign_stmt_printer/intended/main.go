package main

import "snapshot/hxrt"

func main() {
	hxrt.Println(any(hxrt.StringFromLiteral("ok")))
}

func hxrt__test_ast_multi_assign_stmt_smoke() {
	var items map[string]int
	value, found := items["present"]
	_ = value
	_ = found
	var missing int
	missing, found = items["missing"]
	_ = missing
}
