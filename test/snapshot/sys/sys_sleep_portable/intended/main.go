package main

import "snapshot/hxrt"

func main() {
	started := hxrt.ThreadNowSeconds()
	hxrt.SysSleep(0.02)
	elapsed := (hxrt.ThreadNowSeconds() - started)
	hxrt.Println(any((elapsed >= 0.005)))
	hxrt.Println(any((elapsed < 5.0)))
}
