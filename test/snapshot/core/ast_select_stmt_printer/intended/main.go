package main

import "snapshot/hxrt"

func main() {
	hxrt.Println(hxrt.StringFromLiteral("ok"))
}

func hxrt__test_ast_select_stmt_smoke() {
	in := make(chan int, 1)
	out := make(chan int, 1)
	select {
	case out <- 1:
		_ = 11
	case value := <-in:
		_ = value
	case <-in:
		_ = 22
	default:
		_ = 0
	}
}
