package main

import "snapshot/hxrt"

func main() {
	hxrt.Println(any(hxrt.StringFromLiteral("ok")))
}

func hxrt__test_ast_send_recv_stmt_smoke() {
	ch := make(chan int, 1)
	ch <- 7
	_ = <-ch
}
