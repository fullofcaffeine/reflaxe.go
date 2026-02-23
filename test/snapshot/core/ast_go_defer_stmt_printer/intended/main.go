package main

import "snapshot/hxrt"

func main() {
	hxrt.Println(hxrt.StringFromLiteral("ok"))
}

func hxrt__test_ast_go_defer_stmt_smoke() {
	fn := func() {
	}
	defer fn()
	go fn()
}
