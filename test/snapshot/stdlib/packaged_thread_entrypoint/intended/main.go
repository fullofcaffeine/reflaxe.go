package main

import "snapshot/hxrt"

func main() {
	defer hxrt.ThreadWaitForAll()
	hxrt.Println(hxrt.StringFromLiteral("packaged-entry"))
	delay := New_sys__thread__Lock()
	sys__thread__Thread_create(func() {
		delay.wait(0.02)
		hxrt.Println(hxrt.StringFromLiteral("packaged-worker"))
	})
}
