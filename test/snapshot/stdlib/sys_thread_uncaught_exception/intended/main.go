package main

import "snapshot/hxrt"

func main() {
	defer hxrt.ThreadWaitForAll()
	started := New_sys__thread__Lock()
	sys__thread__Thread_create(func() {
		started.release()
		New_sys__thread__Lock().wait(0.02)
		hxrt.Println(any(hxrt.StringFromLiteral("child-before-throw")))
		hxrt.Throw(hxrt.StringFromLiteral("child-failure"))
	})
	started.wait(nil)
	hxrt.Println(any(hxrt.StringFromLiteral("main-survived")))
}
